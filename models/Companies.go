package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Company retorna o modelo para dentro do banco para a compania
type Company struct {
	gorm.Model
	ID        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name      string    `gorm:"type:varchar(100); not null"`
	Zipcode   string    `gorm:"type:varchar(5);NOT NULL"`
	Website   string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// TableName retorna nome fisico que quero para minha tabela
func (u *Company) TableName() string {
	return "companies"
}
