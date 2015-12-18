package model

import "strconv"

// Chocolate is the chocolate that a user consumes in order to get fat and happy
type Chocolate struct {
	ID    int64  `json:"-"`
	Name  string `json:"name"`
	Taste string `json:"taste"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Chocolate) GetID() string {
	return strconv.FormatInt(c.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Chocolate) SetID(id string) error {
	var err error
	c.ID, err = strconv.ParseInt(id, 10, 64)
	return err
}
