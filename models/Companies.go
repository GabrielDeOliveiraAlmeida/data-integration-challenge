package models

import (
	"time"
)

// Company retorna o modelo para dentro do banco para a compania
type Company struct {
	ID          int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	CompanyName string `gorm:"type:varchar(100); not null"`
	Zipcode     string `gorm:"type:varchar(5);NOT NULL"`
	Website     string `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// TableName retorna nome fisico que quero para minha tabela
func (u *Company) TableName() string {
	return "companies"
}
