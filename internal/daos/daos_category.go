package daos

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NamaCategory string
}
