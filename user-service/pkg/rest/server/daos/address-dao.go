package daos

import (
	"errors"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/first-rest-server-project-sqlite-gorm/user-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressDao struct {
	db *gorm.DB
}

func NewAddressDao() (*AddressDao, error) {
	sqlClient, err := sqls.InitGORMSQLiteDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Address{})
	if err != nil {
		return nil, err
	}
	return &AddressDao{
		db: sqlClient.DB,
	}, nil
}

func (addressDao *AddressDao) CreateAddress(m *models.Address) (*models.Address, error) {
	if err := addressDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create address: %v", err)
		return nil, err
	}

	log.Debugf("address created")
	return m, nil
}

func (addressDao *AddressDao) GetAddress(id int64) (*models.Address, error) {
	var m *models.Address
	if err := addressDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get address: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("address retrieved")
	return m, nil
}

func (addressDao *AddressDao) UpdateAddress(id int64, m *models.Address) (*models.Address, error) {
	if id == 0 {
		return nil, errors.New("invalid address ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var address *models.Address
	if err := addressDao.db.Where("id = ?", id).First(&address).Error; err != nil {
		log.Debugf("failed to find address for update: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	if err := addressDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update address: %v", err)
		return nil, err
	}
	log.Debugf("address updated")
	return m, nil
}

func (addressDao *AddressDao) DeleteAddress(id int64) error {
	var m *models.Address
	if err := addressDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete address: %v", err)
		return err
	}

	log.Debugf("address deleted")
	return nil
}
