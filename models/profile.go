package models

import (
	"albert/utils"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID     int       `gorm:"type:int; not null; index:,unique" json:"user_id"`
	FName      string    `gorm:"type:varchar(50);not null" json:"fname"`
	LName      string    `gorm:"type:varchar(50)" json:"lname"`
	Bio        string    `gorm:"type:varchar(200)" json:"bio"`
	Location   string    `gorm:"type:varchar(50);not null" json:"location"`
	Points     int64     `gorm:"type:bigint;default:0" json:"points"`
	OpenToWork bool      `gorm:"type:boolean;default:false" json:"otw"`
	DOB        time.Time `gorm:"not null" json:"dob" `
	Avatar     string    `gorm:"type:varchar(100)" json:"avatar"`
	Skills     []string  `gorm:"type:text[]" json:"skills"`
}

func (profile *Profile) GetProfile(userId int) (*Profile, error) {
	if err := DB.Model(Profile{}).Where("user_id = ?", userId).Take(&profile).Error; err != nil {
		return &Profile{}, err
	}

	return profile, nil
}

func (profile *Profile) CreateProfile() (*Profile, error) {
	if err := DB.Create(&profile).Error; err != nil {

		return &Profile{}, utils.ParseError(err)
	}

	return profile, nil
}

func (profile *Profile) UpdateProfile() (*Profile, error) {
	if err := DB.Model(Profile{}).Where("user_id = ?", profile.UserID).Updates(&profile).Error; err != nil {
		return &Profile{}, err
	}

	return profile, nil
}
