package model

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;comment:BIGINT=>Scalability&Growth, UNSIGNED=>No negative number"`
	Email     string    `gorm:"size:100;not null;unique"`
	Name      string    `gorm:"size:100;not null"`
	Image         string  `gorm:"size:255;not null"`
	
	Account   *Account  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // One-to-One with Account

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
