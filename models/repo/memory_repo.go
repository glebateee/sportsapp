package repo

import (
	"fmt"
	"math/rand"

	"github.com/glebateee/core/services"
	"github.com/glebateee/sportsapp/models"
)

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

func RegisterMemoryRepoService() {
	if err := services.AddSingleton(func() models.Repository {
		repo := &MemoryRepo{}
		repo.Seed()
		return repo
	}); err != nil {
		panic(err)
	}
}

func (r *MemoryRepo) GetProductPage(offset int, limit int) ([]models.Product, int) {
	return getPage(r.products, offset, limit), len(r.products)
}

func getPage(src []models.Product, offset, limit int) []models.Product {
	start := (offset - 1) * limit
	if offset > 0 && len(src) > start {
		end := min(len(src), start+limit)
		return src[start:end]
	}
	return []models.Product{}
}
func (r *MemoryRepo) GetProduct(id int) models.Product {
	for _, p := range r.products {
		if p.ID == id {
			return p
		}
	}
	return models.Product{}
}

func (r *MemoryRepo) GetProducts() []models.Product {
	return r.products
}
func (r *MemoryRepo) GetCategories() []models.Category {
	return r.categories
}
func (r *MemoryRepo) Seed() {
	r.categories = make([]models.Category, 3)
	for i := 0; i < 3; i++ {
		catName := fmt.Sprintf("Category_%v", i+1)
		r.categories[i] = models.Category{ID: i + 1, CategoryName: catName}
	}

	for i := 0; i < 21; i++ {
		name := fmt.Sprintf("Product_%v", i+1)
		price := rand.Float64() * float64(rand.Intn(500))
		cat := &r.categories[rand.Intn(len(r.categories))]
		r.products = append(r.products, models.Product{
			ID:   i + 1,
			Name: name, Price: price,
			Description: fmt.Sprintf("%v (%v)", name, cat.CategoryName),
			Category:    cat,
		})

	}
}
