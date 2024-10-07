package v1

import (
	"product/repository"
	"product/repository/repo"
	"sync"
)

type handler struct {
	productRep repository.ProductRepo
}

func NewProductHandler() *handler {
	return &handler{
		productRep: repo.NewProductRepo(),
	}
}

func Test(wg *sync.WaitGroup, mu *sync.Mutex) {
	wg.Add()
	wg.Done()
	wg.Wait()
	mu.Lock()
	mu.Unlock()
}
