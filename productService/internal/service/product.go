package service

import (
	"context"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/domain"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/dto"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/repository"
)

type ProductService interface {
	CreateCategory(ctx context.Context, input *dto.CreateCategory) error
	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
	GetCategoryById(ctx context.Context, id int64) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int64, input *dto.UpdateCategory) error

	CreateProduct(ctx context.Context, input *dto.CreateProduct) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductById(ctx context.Context, id int64) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int64, input *dto.UpdateProduct) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func (p *productService) CreateCategory(ctx context.Context, input *dto.CreateCategory) error {
	category := &domain.Category{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageURL,
		Status:      "publish",
	}

	return p.productRepository.CreateCategory(ctx, category)
}

func (p *productService) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	return p.productRepository.GetAllCategories(ctx)
}

func (p *productService) GetCategoryById(ctx context.Context, id int64) (*domain.Category, error) {
	return p.productRepository.GetCategoryById(ctx, id)
}

func (p *productService) UpdateCategory(ctx context.Context, id int64, input *dto.UpdateCategory) error {
	category, err := p.productRepository.GetCategoryById(ctx, id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		category.Name = *input.Name
	}

	if input.Description != nil {
		category.Description = *input.Description
	}

	if input.Status != nil {
		category.Status = *input.Status
	}

	return p.productRepository.UpdateCategory(ctx, id, category)
}

func (p *productService) CreateProduct(ctx context.Context, input *dto.CreateProduct) error {
	product := &domain.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryId:  input.CategoryId,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageURL:    input.ImageURL,
		Status:      input.Status,
	}

	return p.productRepository.CreateProduct(ctx, product)
}

func (p *productService) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	return p.productRepository.GetAllProducts(ctx)
}

func (p *productService) GetProductById(ctx context.Context, id int64) (*domain.Product, error) {
	return p.productRepository.GetProductById(ctx, id)
}

func (p *productService) UpdateProduct(ctx context.Context, id int64, input *dto.UpdateProduct) error {
	product, err := p.productRepository.GetProductById(ctx, id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		product.Name = *input.Name
	}

	if input.Description != nil {
		product.Description = *input.Description
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.ImageURL != nil {
		product.ImageURL = *input.ImageURL
	}

	if input.Status != nil {
		product.Status = *input.Status
	}

	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	return p.productRepository.UpdateProduct(ctx, id, product)
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}
