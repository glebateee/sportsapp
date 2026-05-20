package main

import (
	"sync"

	"github.com/glebateee/core/http"

	"github.com/glebateee/core/http/handling"
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/pipeline/basic"
	"github.com/glebateee/core/services"
	"github.com/glebateee/sportsapp/models/repo"
	"github.com/glebateee/sportsapp/store"
)

func registerServices() {
	services.RegisterDefaultServices()
	repo.RegisterMemoryRepoService()
}
func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: store.ProductHandler{}},
		).AddMethodAlias("/", store.ProductHandler.GetProducts, 1).
			AddMethodAlias("/products", store.ProductHandler.GetProducts, 1),
	)
}
func main() {
	registerServices()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
