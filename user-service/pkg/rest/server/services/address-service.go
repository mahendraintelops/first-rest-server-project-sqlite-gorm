package services

import (
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/daos"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/models"
)

type AddressService struct {
	addressDao *daos.AddressDao
}

func NewAddressService() (*AddressService, error) {
	addressDao, err := daos.NewAddressDao()
	if err != nil {
		return nil, err
	}
	return &AddressService{
		addressDao: addressDao,
	}, nil
}

func (addressService *AddressService) CreateAddress(address *models.Address) (*models.Address, error) {
	return addressService.addressDao.CreateAddress(address)
}

func (addressService *AddressService) GetAddress(id int64) (*models.Address, error) {
	return addressService.addressDao.GetAddress(id)
}

func (addressService *AddressService) UpdateAddress(id int64, address *models.Address) (*models.Address, error) {
	return addressService.addressDao.UpdateAddress(id, address)
}

func (addressService *AddressService) DeleteAddress(id int64) error {
	return addressService.addressDao.DeleteAddress(id)
}
