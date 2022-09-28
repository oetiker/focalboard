package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/mattermost/focalboard/server/utils"
)

const (
	CategoryTypeSystem = "system"
	CategoryTypeCustom = "custom"
)

// Category is a board category
// swagger:model
type Category struct {
	// The id for this category
	// required: true
	ID string `json:"id"`

	// The name for this category
	// required: true
	Name string `json:"name"`

	// The user's id for this category
	// required: true
	UserID string `json:"userID"`

	// The team id for this category
	// required: true
	TeamID string `json:"teamID"`

	// The creation time in miliseconds since the current epoch
	// required: true
	CreateAt int64 `json:"createAt"`

	// The last modified time in miliseconds since the current epoch
	// required: true
	UpdateAt int64 `json:"updateAt"`

	// The deleted time in miliseconds since the current epoch. Set to indicate this category is deleted
	// required: false
	DeleteAt int64 `json:"deleteAt"`

	// Category's state in client side
	// required: true
	Collapsed bool `json:"collapsed"`

	// Inter-category sort order per user
	// required: true
	SortOrder int `json:"sortOrder"`

	// Category's type
	// required: true
	Type string `json:"type"`
}

func (c *Category) Hydrate() {
	c.ID = utils.NewID(utils.IDTypeNone)
	c.CreateAt = utils.GetMillis()
	c.UpdateAt = c.CreateAt
	c.SortOrder = 0
	if c.Type == "" {
		c.Type = CategoryTypeCustom
	}
}

func (c *Category) IsValid() error {
	if strings.TrimSpace(c.ID) == "" {
		return NewErrInvalidCategory("category ID cannot be empty")
	}

	if strings.TrimSpace(c.Name) == "" {
		return NewErrInvalidCategory("category name cannot be empty")
	}

	if strings.TrimSpace(c.UserID) == "" {
		return NewErrInvalidCategory("category user ID cannot be empty")
	}

	if strings.TrimSpace(c.TeamID) == "" {
		return NewErrInvalidCategory("category team id ID cannot be empty")
	}

	if c.Type != CategoryTypeCustom && c.Type != CategoryTypeSystem {
		return NewErrInvalidCategory(fmt.Sprintf("category type is invalid. Allowed types: %s and %s", CategoryTypeSystem, CategoryTypeCustom))
	}

	return nil
}

func CategoryFromJSON(data io.Reader) *Category {
	var category *Category
	_ = json.NewDecoder(data).Decode(&category)
	return category
}
