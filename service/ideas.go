package service

import (
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/repo"
)

// IdeaService handles the business logic for ideas.
type IdeaService struct {
	IdeaRepo *repo.IdeaRepo
}

// CreateIdea validates and creates a new idea.
func (s *IdeaService) CreateIdea(idea *models.Idea) error {
	// Add business logic for creating an idea here (e.g., validation)
	return s.IdeaRepo.CreateIdea(idea)
}

// GetIdeas retrieves ideas with pagination.
func (s *IdeaService) GetIdeas(page, limit int) ([]models.Idea, error) {
	return s.IdeaRepo.GetIdeas(page, limit)
}

// GetIdeaByID retrieves an idea by its ID.
func (s *IdeaService) GetIdeaByID(id uint) (*models.Idea, error) {
	return s.IdeaRepo.GetIdeaByID(id)
}
