package model

import "gorm.io/gorm"

type Container struct {
	gorm.Model
	// container could be under another container, but if the container is at top level, this will be 0
	ContainerID uint
	Name        string
	Description string
}
