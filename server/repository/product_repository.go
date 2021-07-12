package repository

import (
	"fmt"

	db "github.com/Minhvn98/ecommerce-fashion/database"
	"github.com/Minhvn98/ecommerce-fashion/models"
)

func GetProductsByTextSearch(text string, limit int, offset int) []models.Product {
	results, err := db.DbConn.Query("SELECT id, name, description, price, sale_percent, category_id, quantity FROM products WHERE MATCH(name) against ((?) IN natural language mode) LIMIT ? OFFSET ?", text, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	var products []models.Product
	var product models.Product
	var categoryId int
	defer results.Close()
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SalePercent, &categoryId, &product.Quantity)
		if err != nil {
			fmt.Println(err)
		}
		product.Category = GetCategoryById(categoryId)
		product.ProductImages = GetImagesProduct(product.ID)
		product.ProductProperties = GetPropertiesProduct(product.ID)

		products = append(products, product)
	}

	return products
}

func GetProductsByCategory(categoryId int, limit int, offset int) []models.Product {
	results, err := db.DbConn.Query("SELECT id, name, description, price, sale_percent, category_id, quantity FROM products WHERE category_id = ? LIMIT ? OFFSET ?", categoryId, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	var products []models.Product
	var product models.Product
	var cateId int
	defer results.Close()
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SalePercent, &cateId, &product.Quantity)
		if err != nil {
			fmt.Println(err)
		}
		product.Category = GetCategoryById(categoryId)
		product.ProductImages = GetImagesProduct(product.ID)
		product.ProductProperties = GetPropertiesProduct(product.ID)

		products = append(products, product)
	}

	return products
}

func GetProducts(limit int, offset int) []models.Product {
	results, err := db.DbConn.Query("SELECT id, name, description, price, sale_percent, category_id, quantity FROM products LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	var products []models.Product
	var product models.Product
	var categoryId int
	defer results.Close()
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SalePercent, &categoryId, &product.Quantity)
		if err != nil {
			fmt.Println(err)
		}
		product.Category = GetCategoryById(categoryId)
		product.ProductImages = GetImagesProduct(product.ID)
		product.ProductProperties = GetPropertiesProduct(product.ID)

		products = append(products, product)
	}

	return products
}

func GetProductById(id int) models.Product {

	results, err := db.DbConn.Query("SELECT id, name, description, price, sale_percent, category_id, quantity FROM products WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	var product models.Product
	var categoryId int
	defer results.Close()
	for results.Next() {
		err = results.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SalePercent, &categoryId, &product.Quantity)
		if err != nil {
			fmt.Println(err)
		}
	}

	if categoryId == 0 {
		return product
	}

	category := GetCategoryById(categoryId)
	product.Category = category
	product.ProductImages = GetImagesProduct(product.ID)
	product.ProductProperties = GetPropertiesProduct(product.ID)

	return product
}

func GetImagesProduct(productId int) []models.ProductImage {
	results, err := db.DbConn.Query("SELECT id, product_id, uri FROM product_images WHERE product_id = ?", productId)
	if err != nil {
		fmt.Println(err)
	}
	var productImages = make([]models.ProductImage, 0)
	var productImage models.ProductImage

	defer results.Close()
	for results.Next() {
		err = results.Scan(&productImage.ID, &productImage.ProductId, &productImage.Uri)
		if err != nil {
			fmt.Println(err)
		}
		productImages = append(productImages, productImage)
	}
	return productImages
}

func GetPropertiesProduct(productId int) []models.ProductProperty {
	results, err := db.DbConn.Query("SELECT id, product_id, `key`, value FROM product_properties WHERE product_id = ?", productId)
	if err != nil {
		fmt.Println(err)
	}
	var productProperties = make([]models.ProductProperty, 0)
	var productProperty models.ProductProperty

	defer results.Close()
	for results.Next() {
		err = results.Scan(&productProperty.ID, &productProperty.ProductId, &productProperty.Key, &productProperty.Value)
		if err != nil {
			fmt.Println(err)
		}
		productProperties = append(productProperties, productProperty)
	}
	return productProperties
}
