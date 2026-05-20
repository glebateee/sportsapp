package store

import (
	"math"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/http/handling"
	"github.com/glebateee/sportsapp/models"
)

type ProductHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}

type ProductTemplateContext struct {
	Products    []models.Product
	Page        int
	PageCount   int
	PageNumbers []int
	PageUrlFunc func(int) string
}

const limit = 4

func (h ProductHandler) GetProducts(offset int) actionresults.ActionResult {
	prods, total := h.Repository.GetProductPage(offset, limit)
	pageCount := int(math.Ceil(float64(total) / float64(limit)))
	return actionresults.NewTemplateAction("product_list.html", ProductTemplateContext{
		Products:    prods,
		Page:        offset,
		PageCount:   pageCount,
		PageNumbers: h.generatePageNumbers(pageCount),
		PageUrlFunc: h.createPageUrlFunction(),
	})
}

func (h ProductHandler) createPageUrlFunction() func(int) string {
	return func(next int) string {
		url, err := h.URLGenerator.GenerateURL(ProductHandler.GetProducts, next)
		if err != nil {
			panic(err)
		}
		return url
	}
}

func (handler ProductHandler) generatePageNumbers(pageCount int) []int {
	pages := make([]int, pageCount)
	for i := range pageCount {
		pages[i] = i + 1
	}
	return pages
}
