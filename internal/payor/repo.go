package payor

import (
	"context"

	"github.com/jinzhu/gorm"
)

// repo is db repo
type repo struct {
	DB *gorm.DB
}

func (r *repo) findUserByUsername(ctx context.Context, username string) (record User, err error) {
	err = r.DB.Where("username = ?", username).First(&record).Error
	return
}
