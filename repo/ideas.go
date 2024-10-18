package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tarunngusain08/Culture-Hub/models"
)

// IdeaRepo handles database operations related to ideas.
type IdeaRepo struct {
	db *gorm.DB
}

// NewIdeaRepo creates a new instance of IdeaRepo.
func NewIdeaRepo(db *gorm.DB) *IdeaRepo {
	return &IdeaRepo{db: db}
}

// CreateIdea saves a new idea to the database.
func (r *IdeaRepo) CreateIdea(idea *models.Idea) error {
	if err := r.db.Create(idea).Error; err != nil {
		return err
	}
	return nil
}

// GetIdeas fetches ideas with pagination.
func (r *IdeaRepo) GetIdeas(page, limit int) ([]models.Idea, error) {
	var ideas []models.Idea
	offset := (page - 1) * limit
	if err := r.db.Offset(offset).Limit(limit).Find(&ideas).Error; err != nil {
		return nil, err
	}
	return ideas, nil
}

// GetIdeaByID fetches an idea by its ID.
func (r *IdeaRepo) GetIdeaByID(id uint) (*models.Idea, error) {
	var idea models.Idea
	if err := r.db.First(&idea, id).Error; err != nil {
		return nil, err
	}
	return &idea, nil
}
