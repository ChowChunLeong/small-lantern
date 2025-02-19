package models

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;comment:BIGINT=>Scalability&Growth, UNSIGNED=>No negative number"`
	Email     string    `gorm:"size:100;not null;unique"`
	Name      string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
