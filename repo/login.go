package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginRepo struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) *LoginRepo {
	return &LoginRepo{db: db}
}

func (l *LoginRepo) Login(userDetails *models.User) error {
	var user models.User

	// Check if the input is an email using a regex pattern
	isEmail := utils.IsValidEmail(userDetails.Email)

	if isEmail {
		// Search by email if it's an email
		if err := l.db.Where("email = ?", userDetails.Email).First(&user).Error; err != nil {
			return err
		}
	} else {
		// Search by username if it's not an email
		if err := l.db.Where("username = ?", userDetails.Username).First(&user).Error; err != nil {
			return err
		}
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password)); err != nil {
		return err
	}

	return nil
}
