package repositories

import (
	"fmt"
	"gin_fleamarket_practice/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository{
	aa := &ItemMemoryRepository{items: items}
	fmt.Println("itemsフィールド")
	fmt.Printf("%+v\n", aa)
	return aa
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	fmt.Println("★★★ FindAll の中 ★★★")
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", r.items)   
	return &r.items, nil
}
