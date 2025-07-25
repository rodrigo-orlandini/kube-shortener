package models

type URL struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ShortURL string `gorm:"not null;uniqueIndex"`
	Visits   int    `gorm:"not null;default:0"`
}
