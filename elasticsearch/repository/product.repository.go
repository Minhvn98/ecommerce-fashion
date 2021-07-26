package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	elastic "github.com/olivere/elastic/v7"

	"es-server/models"
	"es-server/pkg"
)

const (
	indexName = "ecommerce-fashion"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ProductManager struct {
	esClient *pkg.ESClient
}

func NewProductManager(es *pkg.ESClient) *ProductManager {
	return &ProductManager{esClient: es}
}

func (pm *ProductManager) SearchProducts(name string) []*models.Product {
	ctx := context.Background()

	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}
	// build query to search for name
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("name", name))

	// get search's service
	searchService := pm.esClient.Search().Index(indexName).SearchSource(query)

	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}

	// get result
	var products []*models.Product

	for _, hit := range searchResult.Hits.Hits {
		var product models.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		fmt.Println(&product)
		products = append(products, &product)
	}

	return products
}

func (pm *ProductManager) AddProduct(product *models.Product) error {
	ctx := context.Background()

	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	// convert id product
	idProduct := strconv.Itoa(product.ID)

	_, err := pm.esClient.Index().
		Index(indexName).
		BodyJson(product.String()).
		Id(idProduct).
		Do(ctx)

	// call to flush data to disk for search. if no call --> need to wait for 5s to search since inserted
	pm.esClient.Refresh(indexName).Do(ctx)

	return err
}

func (pm *ProductManager) UpdateProduct(product *models.Product) error {
	ctx := context.Background()

	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	// convert id product
	idProduct := strconv.Itoa(product.ID)

	update, err := pm.esClient.Update().Index(indexName).Id(idProduct).
		Upsert(product.String()).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)

	// call to flush data to disk for search. if no call --> need to wait for 5s to search since inserted
	pm.esClient.Refresh(indexName).Do(ctx)

	return err
}

func (pm *ProductManager) DeleteProduct(idProduct string) error {
	ctx := context.Background()

	if pm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	res, err := pm.esClient.Delete().
		Index(indexName).
		Id(idProduct).
		Do(ctx)

	if res.Shards.Successful > 0 {
		fmt.Println("Document deleted from from index: ", idProduct)
	}
	return err
}
