package storage

import (
	"fmt"
	"strconv"

	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
)

// NewUserStorage initializes the storage
func NewUserStorage() *UserStorage {
	return &UserStorage{make(map[int64]*model.User), 1}
}

// UserStorage stores all users
type UserStorage struct {
	users   map[int64]*model.User
	idCount int64
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() map[int64]*model.User {
	return s.users
}

// GetOne user
func (s UserStorage) GetOne(id string) (model.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return model.User{}, fmt.Errorf("User id must be integer: %s", id)
	}
	user, ok := s.users[intID]
	if ok {
		return *user, nil
	}

	return model.User{}, fmt.Errorf("User for id %s not found", id)
}

// Insert a user
func (s *UserStorage) Insert(c model.User) string {
	c.ID = s.idCount
	s.users[c.ID] = &c
	s.idCount++
	return c.GetID()
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("User id must be integer: %s", id)
	}
	_, exists := s.users[intID]
	if !exists {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	delete(s.users, intID)

	return nil
}

// Update a user
func (s *UserStorage) Update(c model.User) error {
	_, exists := s.users[c.ID]
	if !exists {
		return fmt.Errorf("User with id %s does not exist", c.ID)
	}
	s.users[c.ID] = &c

	return nil
}
