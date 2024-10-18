package models

import (
	"time"
	"fmt"
	"encoding/json"
	"database/sql/driver"
	"gorm.io/gorm"
)

type IdeaDao struct {
	baseDAO
}

type Idea struct {
    ID               uint      `gorm:"primaryKey" json:"id"`
    Title            string    `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
    Description      string    `gorm:"type:text;not null" json:"description" binding:"required"`
    Tags             string    `gorm:"type:text" json:"tags" binding:"required"`
    Timeline         CustomDate `json:"timeline" binding:"required"`
    ImpactEstimation string    `json:"impact_estimation" binding:"required"`
    UserID           uint      `json:"user_id" binding:"required"`
    Status           string    `gorm:"type:idea_status;default:'Submitted'" json:"status"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
    UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
    VoteCount        int       `gorm:"default:0" json:"vote_count"`
}

type CustomDate time.Time

// UnmarshalJSON handles the custom parsing for the date format
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1]) // Strip the quotes
	t, err := time.Parse("2006-01-02", str) // Ensure format is correct
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}
	*cd = CustomDate(t) // Assign the parsed date
	return nil
}

// MarshalJSON handles the custom formatting when serializing
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(cd) // Convert CustomDate back to time.Time
	return json.Marshal(t.Format("2006-01-02")) // Format and marshal to JSON
}

// Scan implements the sql.Scanner interface for database scanning
func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDate(time.Time{}) // Assign zero time if value is nil
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan type %T into CustomDate", value)
	}
	*cd = CustomDate(t) // Assign the scanned time
	return nil
}

// Value implements the driver.Valuer interface for database storage
func (cd CustomDate) Value() (driver.Value, error) {
	if cd.IsZero() {
		return nil, nil // Return nil if the time is zero
	}
	return time.Time(cd), nil // Convert CustomDate back to time.Time
}

// IsZero checks if the CustomDate is zero
func (cd CustomDate) IsZero() bool {
	return time.Time(cd).IsZero()
}

func (IdeaDao) TableName() string {
	return "ideas"
}

func (i IdeaDao) ByID(id string) (*Idea, error) {
	var idea Idea
	err := i.db.Where("id =?", id).First(&idea).Error
	return &idea, err
}

func (i IdeaDao) VoteUpdate(id string, vote int) error {
	var err error
	if vote > 0 {
		err = i.db.Model(&Idea{}).Where("id = ?", id).Update("vote_count", gorm.Expr("vote_count + ?", 1)).Error
	} else {
		err = i.db.Model(&Idea{}).Where("id = ?", id).Update("vote_count", gorm.Expr("vote_count - ?", 1)).Error
	}
	return err
}

func (i IdeaDao) Create(idea *Idea) error {
	return i.baseDAO.db.Create(idea).Error
}

func (i IdeaDao) GetCount() (count int64, err error) {
	err = i.baseDAO.db.Model(&Idea{}).Count(&count).Error
	return
}

func (i IdeaDao) GetPaginated(offset, limit int) ([]Idea, error) {
	var ideas []Idea
	err := i.db.Offset(offset).Limit(limit).Find(&ideas).Error
	return ideas, err
}
