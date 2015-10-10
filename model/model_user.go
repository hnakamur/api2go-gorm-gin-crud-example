package model

import (
	"errors"
	"strconv"

	"github.com/manyminds/api2go/jsonapi"
)

// User is a generic database user
type User struct {
	ID int64 `jsonapi:"-"`
	//rename the username field to user-name.
	Username      string      `jsonapi:"name=user-name"`
	PasswordHash  string      `jsonapi:"-"`
	Chocolates    []Chocolate `jsonapi:"-" gorm:"many2many:user_chocolates"`
	ChocolatesIDs []string    `jsonapi:"-" sql:"-"`
	exists        bool        `sql:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u User) GetID() string {
	return strconv.FormatInt(u.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *User) SetID(id string) error {
	var err error
	u.ID, err = strconv.ParseInt(id, 10, 64)
	return err
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u User) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "chocolates",
			Name: "sweets",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u User) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, chocolate := range u.Chocolates {
		result = append(result, jsonapi.ReferenceID{
			ID:   chocolate.GetID(),
			Type: "chocolates",
			Name: "sweets",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u User) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.Chocolates {
		result = append(result, u.Chocolates[key])
	}

	return result
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *User) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "sweets" {
		u.ChocolatesIDs = IDs
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new sweets that a users loves so much
func (u *User) AddToManyIDs(name string, IDs []string) error {
	if name == "sweets" {
		u.ChocolatesIDs = append(u.ChocolatesIDs, IDs...)
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (u *User) DeleteToManyIDs(name string, IDs []string) error {
	if name == "sweets" {
		for _, ID := range IDs {
			for pos, oldID := range u.ChocolatesIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.ChocolatesIDs = append(u.ChocolatesIDs[:pos], u.ChocolatesIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
