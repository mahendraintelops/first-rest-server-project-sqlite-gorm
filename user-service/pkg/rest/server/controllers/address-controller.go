package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/models"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type AddressController struct {
	addressService *services.AddressService
}

func NewAddressController() (*AddressController, error) {
	addressService, err := services.NewAddressService()
	if err != nil {
		return nil, err
	}
	return &AddressController{
		addressService: addressService,
	}, nil
}

func (addressController *AddressController) CreateAddress(context *gin.Context) {
	// validate input
	var input models.Address
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger address creation
	addressCreated, err := addressController.addressService.CreateAddress(&input)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, addressCreated)
}

func (addressController *AddressController) FetchAddress(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger address fetching
	address, err := addressController.addressService.GetAddress(id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceName := os.Getenv("SERVICE_NAME")
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// get the current span by the request context
		currentSpan := trace.SpanFromContext(context.Request.Context())
		currentSpan.SetAttributes(attribute.String("address.id", strconv.FormatInt(address.Id, 10)))
	}

	context.JSON(http.StatusOK, address)
}

func (addressController *AddressController) UpdateAddress(context *gin.Context) {
	// validate input
	var input models.Address
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger address update
	if _, err := addressController.addressService.UpdateAddress(id, &input); err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (addressController *AddressController) DeleteAddress(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger address deletion
	if err := addressController.addressService.DeleteAddress(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
