package storage

import (
	"fmt"
	"strconv"

	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
	"github.com/jinzhu/gorm"
)

// NewUserStorage initializes the storage
func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db}
}

// UserStorage stores all users
type UserStorage struct {
	db *gorm.DB
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() (map[int64]*model.User, error) {
	var users []model.User
	s.db.Find(&users)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	userMap := make(map[int64]*model.User)
	for i, _ := range users {
		u := &users[i]
		userMap[u.ID] = u
	}
	return userMap, nil
}

// GetOne user
func (s UserStorage) GetOne(id string) (model.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return model.User{}, fmt.Errorf("User id must be integer: %s", id)
	}
	return s.getOneWithAssociations(intID)
}

func (s UserStorage) getOneWithAssociations(id int64) (model.User, error) {
	var user model.User
	s.db.First(&user, id)
	s.db.Model(&user).Related(&user.Chocolates, "Chocolates")
	if err := s.db.Error; err == gorm.RecordNotFound {
		return model.User{}, fmt.Errorf("User for id %d not found", id)
	} else if err != nil {
		return model.User{}, err
	}
	user.ChocolatesIDs = make([]string, len(user.Chocolates))
	for i, choc := range user.Chocolates {
		user.ChocolatesIDs[i] = choc.GetID()
	}
	return user, nil
}

// Insert a user
func (s *UserStorage) Insert(c model.User) (string, error) {
	c.Chocolates = make([]model.Chocolate, len(c.ChocolatesIDs))
	err := s.updateChocolatesByChocolatesIDs(&c)
	if err != nil {
		return "", err
	}
	s.db.Create(&c)
	if s.db.Error != nil {
		return "", s.db.Error
	}
	return c.GetID(), nil
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("User id must be integer: %s", id)
	}

	var user model.User
	s.db.First(&user, intID)
	if err := s.db.Error; err == gorm.RecordNotFound {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	s.db.Delete(&user)

	return s.db.Error
}

// Update a user
func (s *UserStorage) Update(c model.User) error {
	user, err := s.getOneWithAssociations(c.ID)
	if err != nil {
		return err
	}

	user.Username = c.Username
	user.PasswordHash = c.PasswordHash
	var chocAssocsToDelete []model.Chocolate
	for i, chocID := range user.ChocolatesIDs {
		if indexOf(chocID, c.ChocolatesIDs) == -1 {
			chocAssocsToDelete = append(chocAssocsToDelete, user.Chocolates[i])
		}
	}
	if len(chocAssocsToDelete) > 0 {
		s.db.Model(&user).Association("Chocolates").Delete(chocAssocsToDelete)
	}

	user.ChocolatesIDs = c.ChocolatesIDs
	err = s.updateChocolatesByChocolatesIDs(&user)
	if err != nil {
		return err
	}
	s.db.Save(&user)
	return s.db.Error
}

func indexOf(s string, items []string) int {
	for i, item := range items {
		if s == item {
			return i
		}
	}
	return -1
}

func (s *UserStorage) updateChocolatesByChocolatesIDs(u *model.User) error {
	u.Chocolates = make([]model.Chocolate, len(u.ChocolatesIDs))
	for i, chocID := range u.ChocolatesIDs {
		intID, err := strconv.ParseInt(chocID, 10, 64)
		if err != nil {
			return err
		}
		s.db.First(&u.Chocolates[i], intID)
	}
	return nil
}
