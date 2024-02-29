package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ContainerID uint
	Name        string
	Description string
}
