package custom_models

import (
	"graphql/internal/user"
)

type Catalog struct {
	ID       uint     `json:"id" gorm:"primarykey"`
	Name     string   `json:"name" gorm:"not null"`
	ParentID uint     `json:"parent_id" gorm:"index"`
	Parent   *Catalog `json:"parent" gorm:"foreignKey:ParentID;references:ID"`
	// Дочерние разделы
	Childs []*Catalog `json:"childs" gorm:"foreignKey:ID;references:ParentID"`
	// Товары в разделе
	Items []*Item `json:"items" gorm:"foreignKey:ID;references:ParentID"`
}

type Seller struct {
	ID    uint    `json:"id" gorm:"primarykey"`
	Name  string  `json:"name" gorm:"not null"`
	Deals uint    `json:"deals" gorm:"not null"`
	Items []*Item `json:"items" gorm:"foreignKey:ID"`
}

type CartItem struct {
	ID       uint       `json:"id" gorm:"primarykey"`
	ItemID   uint       `gorm:"primarykey" gorm:"index:idx_member"`
	Item     *Item      `json:"item" gorm:"foreignKey:ItemID"`
	UserID   uint       `gorm:"primarykey" gorm:"index:idx_member"`
	User     *user.User `json:"user" gorm:"foreignKey:ID"`
	Quantity int        `json:"quantity" gorm:"not null"`
}
