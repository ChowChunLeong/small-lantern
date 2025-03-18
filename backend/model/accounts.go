package model

type Account struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;comment:BIGINT=>Scalability&Growth, UNSIGNED=>No negative number"`

	UserID            uint64 `gorm:"unique;not null"` // One-to-One (unique ensures a single user per account)
	Provider          string `gorm:"not null;index;uniqueIndex:idx_provider_account"`
	ProviderAccountID string `gorm:"size:100;not null;uniqueIndex:idx_provider_account"`

	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // One-to-One
}