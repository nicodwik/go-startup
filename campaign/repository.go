package campaign

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Campaign, error)
	GetByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Campaign, error) {
	var campaings []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaings).Error
	if err != nil {
		return campaings, err
	}

	return campaings, nil
}

func (r *repository) GetByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}