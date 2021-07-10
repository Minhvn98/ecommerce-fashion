package repository

import (
	"fmt"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func GetCategories() []models.Category {
	results, err := db.DbConn.Query("SELECT id, name FROM categories")
	if err != nil {
		fmt.Println(err)
	}
	var categories = make([]models.Category, 0)
	var category models.Category
	defer results.Close()
	for results.Next() {
		err = results.Scan(&category.ID, &category.Name)
		if err != nil {
			fmt.Println(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func GetCategoryById(id int) models.Category {
	results, err := db.DbConn.Query("SELECT id, name FROM categories WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	var category models.Category
	defer results.Close()
	for results.Next() {
		err = results.Scan(&category.ID, &category.Name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return category
}
