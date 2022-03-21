package fs_test

import (
	"testing"

	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/stretchr/testify/assert"
)

func TestConcatDirPath(t *testing.T) {
	t.Run("WithBackslash", func(t *testing.T) {
		appName := fs.ConcatDirPath("github.com/nurcahyaari/library", "routes")

		assert.Equal(t, appName, "github.com/nurcahyaari/library/routes")
	})

	t.Run("WithouteBackslash", func(t *testing.T) {
		appName := fs.ConcatDirPath("github.com/nurcahyaari/lib/", "routes")

		assert.Equal(t, appName, "github.com/nurcahyaari/lib/routes")
	})
}

func TestGetAppNameBasedOnGoMod(t *testing.T) {
	t.Run("Based On Go Mod", func(t *testing.T) {
		appName := fs.GetAppNameBasedOnGoMod("github.com/nurcahyaari/library")

		assert.Equal(t, appName, "library")
	})

	t.Run("Based On App Name", func(t *testing.T) {
		appName := fs.GetAppNameBasedOnGoMod("library")

		assert.Equal(t, appName, "library")
	})
}

func TestReadImportedPackages(t *testing.T) {
	t.Run("Multiple import", func(t *testing.T) {
		dependency := `
		package services

		import (
			"context"
			"golang-starter/src/modules/product/dto"
			"golang-starter/src/modules/product/repositories"

			"github.com/rs/zerolog/log"
		)

		type ProductService interface {
			GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
			GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
			CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
			DeleteProduct(ctx context.Context, productID int) error
		}

		type ProductServiceImpl struct {
			ProductRepository repositories.Repositories
		}

		func NewProductService(
			productRepository repositories.Repositories,
		) *ProductServiceImpl {
			return &ProductServiceImpl{
				ProductRepository: productRepository,
			}
		}

		func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
			productList, err := s.ProductRepository.GetProductsList(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productsResp := dto.CreateProductsListResponse(productList)
			return productsResp, nil
		}

		func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
			product, err := s.ProductRepository.
				FilterProducts(
					repositories.
						NewProductsFilter("AND").
						SetFilterByProductId(productID, "="),
				).
				GetProducts(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productResp := dto.CreateProductsResponse(*product)
			return productResp, nil
		}

		func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
			product := data.ToProductEntities()

			// start repository
			var err error
			tx := s.ProductRepository.StartTx()

			defer func() {
				s.ProductRepository.CloseTx()
				if err != nil {
					log.Err(err).Msg("an error occured")
					tx.Rollback()
				}
			}()

			res, err := s.ProductRepository.InsertProducts(ctx, product)
			if err != nil {
				return nil, err
			}

			lastInsertedId, err := res.LastInsertId()

			productImages := data.ToProductImagesEntities(lastInsertedId)

			_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
			if err != nil {
				return nil, err
			}

			if err = tx.Commit(); err != nil {
				return nil, err
			}

			return nil, nil
		}

		func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
			err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

			if err != nil {
				log.Err(err).Msg("Error deleting product")
				return err
			}

			return nil
		}
		`

		expect := []fs.ImportedPackages{
			{
				FilePath: "context",
			},
			{
				FilePath: "golang-starter/src/modules/product/dto",
			},
			{
				FilePath: "golang-starter/src/modules/product/repositories",
			},
			{
				FilePath: "github.com/rs/zerolog/log",
			},
		}

		actual := fs.ReadImportedPackages(dependency)

		assert.Equal(t, expect, actual)
	})
	t.Run("Multiple import 2", func(t *testing.T) {
		dependency := `package http

import (
	productsvc "test1/src/modules/product/service"
	usersvc "test1/src/modules/user/service"

	"github.com/go-chi/chi/v5"
)

type HttpHandler interface {
	Router(r *chi.Mux)
}

type HttpHandlerImpl struct {
	userSvc    usersvc.UserService
	productSvc productsvc.ProductService
}

func NewHttpHandler(
	userSvc usersvc.UserService,
	productSvc productsvc.ProductService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}
func (h *HttpHandlerImpl) Router(r *chi.Mux) {}
		`
		expect := []fs.ImportedPackages{
			{
				Alias:    "productsvc",
				FilePath: "test1/src/modules/product/service",
			},
			{
				Alias:    "usersvc",
				FilePath: "test1/src/modules/user/service",
			},
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		}

		actual := fs.ReadImportedPackages(dependency)

		assert.Equal(t, expect, actual)
	})
	t.Run("Multiple import 3", func(t *testing.T) {
		dependency := `package http

import (
	"github.com/go-chi/chi/v5"
	videos1svc "test1/src/modules/videos1/service"
)

type HttpHandler interface {
Router(r *chi.Mux)
}

type HttpHandlerImpl struct {
	videos1Svc videos1svc.Videos1Service
}

func NewHttpHandler(
	videos1Svc videos1svc.Videos1Service,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		videos1Svc: videos1Svc,
	}
}
func (h *HttpHandlerImpl) Router(r *chi.Mux) {}
		`
		expect := []fs.ImportedPackages{
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
			{
				Alias:    "videos1svc",
				FilePath: "test1/src/modules/videos1/service",
			},
		}

		actual := fs.ReadImportedPackages(dependency)

		assert.Equal(t, expect, actual)
	})
	t.Run("Multiple import 4", func(t *testing.T) {
		dependency := `
		
package http

import (
	"github.com/go-chi/chi/v5"
	videos1svc "test1/src/modules/videos1/service"
)

type HttpHandler interface {
Router(r *chi.Mux)
}

type HttpHandlerImpl struct {
	videos1Svc videos1svc.Videos1Service
}

func NewHttpHandler(
	videos1Svc videos1svc.Videos1Service,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		videos1Svc: videos1Svc,
	}
}
func (h *HttpHandlerImpl) Router(r *chi.Mux) {}



		`
		expect := []fs.ImportedPackages{
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
			{
				Alias:    "videos1svc",
				FilePath: "test1/src/modules/videos1/service",
			},
		}

		actual := fs.ReadImportedPackages(dependency)

		assert.Equal(t, expect, actual)
	})
	t.Run("Single import", func(t *testing.T) {
		dependency := `
		package services

		import "context"

		type ProductService interface {
			GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
			GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
			CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
			DeleteProduct(ctx context.Context, productID int) error
		}

		type ProductServiceImpl struct {
			ProductRepository repositories.Repositories
		}

		func NewProductService(
			productRepository repositories.Repositories,
		) *ProductServiceImpl {
			return &ProductServiceImpl{
				ProductRepository: productRepository,
			}
		}

		func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
			productList, err := s.ProductRepository.GetProductsList(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productsResp := dto.CreateProductsListResponse(productList)
			return productsResp, nil
		}

		func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
			product, err := s.ProductRepository.
				FilterProducts(
					repositories.
						NewProductsFilter("AND").
						SetFilterByProductId(productID, "="),
				).
				GetProducts(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productResp := dto.CreateProductsResponse(*product)
			return productResp, nil
		}

		func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
			product := data.ToProductEntities()

			// start repository
			var err error
			tx := s.ProductRepository.StartTx()

			defer func() {
				s.ProductRepository.CloseTx()
				if err != nil {
					log.Err(err).Msg("an error occured")
					tx.Rollback()
				}
			}()

			res, err := s.ProductRepository.InsertProducts(ctx, product)
			if err != nil {
				return nil, err
			}

			lastInsertedId, err := res.LastInsertId()

			productImages := data.ToProductImagesEntities(lastInsertedId)

			_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
			if err != nil {
				return nil, err
			}

			if err = tx.Commit(); err != nil {
				return nil, err
			}

			return nil, nil
		}

		func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
			err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

			if err != nil {
				log.Err(err).Msg("Error deleting product")
				return err
			}

			return nil
		}
		`

		expect := []fs.ImportedPackages{
			{
				FilePath: "context",
			},
		}

		actual := fs.ReadImportedPackages(dependency)

		assert.Equal(t, expect, actual)
	})
}

func TestReadInterfaceWithMethod(t *testing.T) {
	t.Run("Get Method List", func(t *testing.T) {
		dependency := `
		package services

		import (
			"context"
			"golang-starter/src/modules/product/dto"
			"golang-starter/src/modules/product/repositories"

			"github.com/rs/zerolog/log"
		)

		type ProductService interface {
			GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
			GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
			CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
			DeleteProduct(ctx context.Context, productID int) error
		}

		type ProductServiceImpl struct {
			ProductRepository repositories.Repositories
		}

		func NewProductService(
			productRepository repositories.Repositories,
		) *ProductServiceImpl {
			return &ProductServiceImpl{
				ProductRepository: productRepository,
			}
		}

		func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
			productList, err := s.ProductRepository.GetProductsList(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productsResp := dto.CreateProductsListResponse(productList)
			return productsResp, nil
		}

		func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
			product, err := s.ProductRepository.
				FilterProducts(
					repositories.
						NewProductsFilter("AND").
						SetFilterByProductId(productID, "="),
				).
				GetProducts(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productResp := dto.CreateProductsResponse(*product)
			return productResp, nil
		}

		func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
			product := data.ToProductEntities()

			// start repository
			var err error
			tx := s.ProductRepository.StartTx()

			defer func() {
				s.ProductRepository.CloseTx()
				if err != nil {
					log.Err(err).Msg("an error occured")
					tx.Rollback()
				}
			}()

			res, err := s.ProductRepository.InsertProducts(ctx, product)
			if err != nil {
				return nil, err
			}

			lastInsertedId, err := res.LastInsertId()

			productImages := data.ToProductImagesEntities(lastInsertedId)

			_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
			if err != nil {
				return nil, err
			}

			if err = tx.Commit(); err != nil {
				return nil, err
			}

			return nil, nil
		}

		func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
			err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

			if err != nil {
				log.Err(err).Msg("Error deleting product")
				return err
			}

			return nil
		}
		`

		expect := []fs.DependencyInterface{
			{
				Method: "GetProducts(ctx context.Context) (dto.ProductsListResponse, error)",
			},
			{
				Method: "GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)",
			},
			{
				Method: "CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)",
			},
			{
				Method: "DeleteProduct(ctx context.Context, productID int) error",
			},
		}

		actual := fs.ReadInterfaceWithMethod(dependency)

		assert.Equal(t, expect, actual)
	})

	t.Run("Get unstructure method", func(t *testing.T) {
		dependency := `
package http

import (
	"github.com/go-chi/chi/v5"
	videos1svc "test1/src/modules/videos1/service"
)

type HttpHandler interface {
Router(r *chi.Mux)
}

type HttpHandlerImpl struct {
	videos1Svc videos1svc.Videos1Service
}

func NewHttpHandler(
	videos1Svc videos1svc.Videos1Service,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		videos1Svc: videos1Svc,
	}
}
func (h *HttpHandlerImpl) Router(r *chi.Mux) {}`

		expect := []fs.DependencyInterface{
			{
				Method: "Router(r *chi.Mux)",
			},
		}

		actual := fs.ReadInterfaceWithMethod(dependency)

		assert.Equal(t, expect, actual)
	})
}

func TestReadStructWithObject(t *testing.T) {
	t.Run("Get Struct Obj List", func(t *testing.T) {
		dependency := `
		package services

		import (
			"context"
			"golang-starter/src/modules/product/dto"
			"golang-starter/src/modules/product/repositories"

			"github.com/rs/zerolog/log"
		)

		type ProductService interface {
			GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
			GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
			CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
			DeleteProduct(ctx context.Context, productID int) error
		}

		type ProductServiceImpl struct {
			ProductRepository repositories.Repositories
		}

		func NewProductService(
			productRepository repositories.Repositories,
		) *ProductServiceImpl {
			return &ProductServiceImpl{
				ProductRepository: productRepository,
			}
		}

		func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
			productList, err := s.ProductRepository.GetProductsList(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productsResp := dto.CreateProductsListResponse(productList)
			return productsResp, nil
		}

		func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
			product, err := s.ProductRepository.
				FilterProducts(
					repositories.
						NewProductsFilter("AND").
						SetFilterByProductId(productID, "="),
				).
				GetProducts(ctx)
			if err != nil {
				log.Err(err).Msg("Error fetch productList from DB")
			}
			productResp := dto.CreateProductsResponse(*product)
			return productResp, nil
		}

		func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
			product := data.ToProductEntities()

			// start repository
			var err error
			tx := s.ProductRepository.StartTx()

			defer func() {
				s.ProductRepository.CloseTx()
				if err != nil {
					log.Err(err).Msg("an error occured")
					tx.Rollback()
				}
			}()

			res, err := s.ProductRepository.InsertProducts(ctx, product)
			if err != nil {
				return nil, err
			}

			lastInsertedId, err := res.LastInsertId()

			productImages := data.ToProductImagesEntities(lastInsertedId)

			_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
			if err != nil {
				return nil, err
			}

			if err = tx.Commit(); err != nil {
				return nil, err
			}

			return nil, nil
		}

		func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
			err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

			if err != nil {
				log.Err(err).Msg("Error deleting product")
				return err
			}

			return nil
		}
		`

		expect := []fs.DependencyStruct{
			{
				ObjectName:     "ProductRepository",
				ObjectDataType: "repositories.Repositories",
			},
		}

		actual := fs.ReadStructWithObject(dependency)

		assert.Equal(t, expect, actual)
	})

	t.Run("Get Struct Obj List 1", func(t *testing.T) {
		dependency := `
		package http

import (
	productsvc "test1/src/modules/product/service"
	usersvc "test1/src/modules/user/service"

	"github.com/go-chi/chi/v5"
)

type HttpHandler interface {
	Router(r *chi.Mux)
}

type HttpHandlerImpl struct {
	userSvc    usersvc.UserService
	productSvc productsvc.ProductService
}

func NewHttpHandler(
	userSvc usersvc.UserService,
	productSvc productsvc.ProductService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}
func (h *HttpHandlerImpl) Router(r *chi.Mux) {}
		`

		expect := []fs.DependencyStruct{
			{
				ObjectName:     "userSvc",
				ObjectDataType: "usersvc.UserService",
			},
			{
				ObjectName:     "productSvc",
				ObjectDataType: "productsvc.ProductService",
			},
		}

		actual := fs.ReadStructWithObject(dependency)

		assert.Equal(t, expect, actual)
	})
}

// func TestMethodImpl(t *testing.T) {
// 	t.Run("Get Method List", func(t *testing.T) {
// 		dependency := `
// package services

// import (
// 	"context"
// 	"golang-starter/src/modules/product/dto"
// 	"golang-starter/src/modules/product/repositories"

// 	"github.com/rs/zerolog/log"
// )

// type ProductService interface {
// 	GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
// 	GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
// 	CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
// 	DeleteProduct(ctx context.Context, productID int) error
// }

// type ProductServiceImpl struct {
// 	ProductRepository repositories.Repositories
// }

// func NewProductService(
// 	productRepository repositories.Repositories,
// ) *ProductServiceImpl {
// 	return &ProductServiceImpl{
// 		ProductRepository: productRepository,
// 	}
// }

// func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
// 	productList, err := s.ProductRepository.GetProductsList(ctx)
// 	if err != nil {
// 		log.Err(err).Msg("Error fetch productList from DB")
// 	}
// 	productsResp := dto.CreateProductsListResponse(productList)
// 	return productsResp, nil
// }

// func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
// 	product, err := s.ProductRepository.
// 		FilterProducts(
// 			repositories.
// 				NewProductsFilter("AND").
// 				SetFilterByProductId(productID, "="),
// 		).
// 		GetProducts(ctx)
// 	if err != nil {
// 		log.Err(err).Msg("Error fetch productList from DB")
// 	}
// 	productResp := dto.CreateProductsResponse(*product)
// 	return productResp, nil
// }

// func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
// 	product := data.ToProductEntities()

// 	// start repository
// 	var err error
// 	tx := s.ProductRepository.StartTx()

// 	defer func() {
// 		s.ProductRepository.CloseTx()
// 		if err != nil {
// 			log.Err(err).Msg("an error occured")
// 			tx.Rollback()
// 		}
// 	}()

// 	res, err := s.ProductRepository.InsertProducts(ctx, product)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lastInsertedId, err := res.LastInsertId()

// 	productImages := data.ToProductImagesEntities(lastInsertedId)

// 	_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return nil, err
// 	}

// 	return nil, nil
// }

// func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
// 	err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

// 	if err != nil {
// 		log.Err(err).Msg("Error deleting product")
// 		return err
// 	}

// 	return nil
// }
// 		`

// 		expect := []string{
// 			`func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
// 	productList, err := s.ProductRepository.GetProductsList(ctx)
// 	if err != nil {
// 		log.Err(err).Msg("Error fetch productList from DB")
// 	}
// 	productsResp := dto.CreateProductsListResponse(productList)
// 	return productsResp, nil
// }`,
// 			`func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
// 	product, err := s.ProductRepository.
// 		FilterProducts(
// 			repositories.
// 				NewProductsFilter("AND").
// 				SetFilterByProductId(productID, "="),
// 		).
// 		GetProducts(ctx)
// 	if err != nil {
// 		log.Err(err).Msg("Error fetch productList from DB")
// 	}
// 	productResp := dto.CreateProductsResponse(*product)
// 	return productResp, nil
// }`,
// 			`func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
// 	product := data.ToProductEntities()

// 	// start repository
// 	var err error
// 	tx := s.ProductRepository.StartTx()

// 	defer func() {
// 		s.ProductRepository.CloseTx()
// 		if err != nil {
// 			log.Err(err).Msg("an error occured")
// 			tx.Rollback()
// 		}
// 	}()

// 	res, err := s.ProductRepository.InsertProducts(ctx, product)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lastInsertedId, err := res.LastInsertId()

// 	productImages := data.ToProductImagesEntities(lastInsertedId)

// 	_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return nil, err
// 	}

// 	return nil, nil
// }`,
// 			`func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
// 	err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

// 	if err != nil {
// 		log.Err(err).Msg("Error deleting product")
// 		return err
// 	}

// 	return nil
// }`,
// 		}

// 		actual := fs.ReadMethodImpl(dependency)

// 		assert.Equal(t, expect, actual)
// 	})
// }

func TestReadFile(t *testing.T) {
	t.Run("test read file", func(t *testing.T) {
		res, err := fs.ReadFile("./../../LICENSE")
		exp := `MIT License

Copyright (c) 2022 Ari Nurcahya

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
		assert.NoError(t, err)
		assert.Equal(t, exp, res)
	})
}
