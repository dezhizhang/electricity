package utils

import "gorm.io/gorm"

func Paginate(page,pageSize int) func(db *gorm.DB) *gorm.DB  {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch  {

		}
	}
}