package custom_models

type Item struct {
	ID           uint    `json:"id" gorm:"primarykey"`
	Name         string  `json:"name,omitempty"`
	InStockValue int     `json:"in_stock_val"`
	SellerID     uint    `json:"seller_id"`
	Seller       Seller  `json:"seller" gorm:"foreignKey:SellerID"`
	ParentID     uint    `json:"parent_id"`
	Catalog      Catalog `gorm:"foreignKey:ParentID"`
}

func (item *Item) InStockText() string {
	if item.InStockValue <= 1 {
		return "мало"
	} else if item.InStockValue >= 2 && item.InStockValue <= 3 {
		return "хватает"
	}
	return "много"
}
