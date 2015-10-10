package storage

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
)

// sorting
type byID []model.Chocolate

func (c byID) Len() int {
	return len(c)
}

func (c byID) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c byID) Less(i, j int) bool {
	return c[i].ID < c[j].ID
}

// NewChocolateStorage initializes the storage
func NewChocolateStorage() *ChocolateStorage {
	return &ChocolateStorage{make(map[int64]*model.Chocolate), 1}
}

// ChocolateStorage stores all of the tasty chocolate, needs to be injected into
// User and Chocolate Resource. In the real world, you would use a database for that.
type ChocolateStorage struct {
	chocolates map[int64]*model.Chocolate
	idCount    int64
}

// GetAll of the chocolate
func (s ChocolateStorage) GetAll() []model.Chocolate {
	result := []model.Chocolate{}
	for key := range s.chocolates {
		result = append(result, *s.chocolates[key])
	}

	sort.Sort(byID(result))
	return result
}

// GetOne tasty chocolate
func (s ChocolateStorage) GetOne(id string) (model.Chocolate, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return model.Chocolate{}, fmt.Errorf("Chocolate id must be integer: %s", id)
	}
	choc, ok := s.chocolates[intID]
	if ok {
		return *choc, nil
	}

	return model.Chocolate{}, fmt.Errorf("Chocolate for id %s not found", id)
}

// Insert a fresh one
func (s *ChocolateStorage) Insert(c model.Chocolate) string {
	c.ID = s.idCount
	s.chocolates[c.ID] = &c
	s.idCount++
	return c.GetID()
}

// Delete one :(
func (s *ChocolateStorage) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("Chocolate id must be integer: %s", id)
	}
	_, exists := s.chocolates[intID]
	if !exists {
		return fmt.Errorf("Chocolate with id %s does not exist", id)
	}
	delete(s.chocolates, intID)

	return nil
}

// Update updates an existing chocolate
func (s *ChocolateStorage) Update(c model.Chocolate) error {
	_, exists := s.chocolates[c.ID]
	if !exists {
		return fmt.Errorf("Chocolate with id %s does not exist", c.ID)
	}
	s.chocolates[c.ID] = &c

	return nil
}
