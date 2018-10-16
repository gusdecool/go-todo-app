package model

import (
	"time"
)

type Task struct {
	Id int `gorm:"type:int;not null;primary_key"`
	Name string `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
}