package repository

import (
	"github.com/f1rdavsi/reporter/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (c *CategoryRepository) CreateCategory(cat models.Category) (models.Category, error) {
	if err := c.db.Create(&cat).Error; err != nil {
		return cat, err
	}

	return cat, nil
}

func (c *CategoryRepository) GetCategories() ([]models.Category, error) {
	var cats []models.Category

	if err := c.db.Model(&models.Category{}).Find(&cats).Error; err != nil {
		return cats, err
	}

	return cats, nil
}

func (c *CategoryRepository) GetCategory(id int) (models.Category, error) {
	var cat models.Category
	if err := c.db.Where("id = ?", id).First(&cat).Error; err != nil {
		return cat, err
	}

	return cat, nil
}

func (c *CategoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	if err := c.db.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (c *CategoryRepository) DeleteCategory(id int) error {
	return c.db.
		Model(&models.Category{}).
		Where("id = ? AND is_active = FALSE", id).
		Update("is_active", false).
		Error
}

func (c *CategoryRepository) RestoreCategory(id int) error {
	return c.db.
		Model(&models.Category{}).
		Where("id = ? AND is_active = TRUE", id).
		Update("is_active", true).
		Error
}
