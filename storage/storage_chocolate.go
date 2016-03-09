package storage

import (
	"fmt"
	"strconv"

	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
	"github.com/jinzhu/gorm"
)

// NewChocolateStorage initializes the storage
func NewChocolateStorage(db *gorm.DB) *ChocolateStorage {
	return &ChocolateStorage{db}
}

// ChocolateStorage stores all of the tasty chocolate, needs to be injected into
// User and Chocolate Resource. In the real world, you would use a database for that.
type ChocolateStorage struct {
	db *gorm.DB
}

// GetAll of the chocolate
func (s ChocolateStorage) GetAll() []model.Chocolate {
	var chocolates []model.Chocolate
	s.db.Order("id").Find(&chocolates)
	return chocolates
}

// GetOne tasty chocolate
func (s ChocolateStorage) GetOne(id string) (model.Chocolate, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return model.Chocolate{}, fmt.Errorf("Chocolate id must be integer: %s", id)
	}
	var choc model.Chocolate
	s.db.First(&choc, intID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return model.Chocolate{}, fmt.Errorf("Chocolate for id %s not found", id)
	}
	return choc, nil
}

// Insert a fresh one
func (s *ChocolateStorage) Insert(c model.Chocolate) string {
	s.db.Create(&c)
	return c.GetID()
}

// Delete one :(
func (s *ChocolateStorage) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("Chocolate id must be integer: %s", id)
	}

	var choc model.Chocolate
	s.db.First(&choc, intID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Chocolate with id %s does not exist", id)
	}
	s.db.Delete(&choc)

	return s.db.Error
}

// Update updates an existing chocolate
func (s *ChocolateStorage) Update(c model.Chocolate) error {
	var choc model.Chocolate
	s.db.First(&choc, c.ID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Chocolate with id %s does not exist", c.ID)
	} else if err != nil {
		return err
	}
	choc.Name = c.Name
	choc.Taste = c.Taste
	s.db.Save(&choc)
	return s.db.Error
}
