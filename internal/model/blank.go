package model

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Tags []string

func (t *Tags) Scan(val interface{}) error {
	s := val.(string)
	if len(s) <= 2 {
		return nil
	}
	ss := strings.Split(string(s)[1:len(s)-1], ",")
	*t = ss
	return nil
}

func (t Tags) Value() (driver.Value, error) {
	return "{" + strings.Join(t, ",") + "}", nil
}

type Blank struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Type      string         `json:"type"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Tags      Tags           `json:"tags" gorm:"type:text[];column:tags"`
	Matters   []*Matter      `json:"matters" gorm:"many2many:blank_matters;"`
	UserID    uint           `json:"user_id"`
	User      *User          `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CreateBlankTemp struct {
	Type    string `json:"type"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	Matters string `json:"matters"`
}

func (c CreateBlankTemp) CreateBlank() (*CreateBlank, error) {
	tags := strings.Split(c.Tags, ";")
	var matters []uint
	for _, id := range strings.Split(c.Matters, ";") {
		matterID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return nil, errors.New("文件id有误")
		}
		matters = append(matters, uint(matterID))
	}

	return &CreateBlank{
		Type:    c.Type,
		Title:   c.Title,
		Content: c.Content,
		Tags:    tags,
		Matters: matters,
	}, nil
}

type CreateBlank struct {
	Type    string   `json:"type"`
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Matters []uint   `json:"matters"`
}
