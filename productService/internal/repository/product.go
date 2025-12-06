package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/domain"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
	GetCategoryById(ctx context.Context, id int64) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int64, category *domain.Category) error

	CreateProduct(ctx context.Context, product *domain.Product) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductById(ctx context.Context, id int64) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int64, product *domain.Product) error
}

type productRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func (p *productRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	return nil
}

func (p *productRepository) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	return nil, nil
}

func (p *productRepository) GetCategoryById(ctx context.Context, id int64) (*domain.Category, error) {
	return nil, nil
}

func (p *productRepository) UpdateCategory(ctx context.Context, id int64, category *domain.Category) error {
	return nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return nil
}

func (p *productRepository) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	return nil, nil
}

func (p *productRepository) GetProductById(ctx context.Context, id int64) (*domain.Product, error) {
	return nil, nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, id int64, product *domain.Product) error {
	return nil
}

func NewProductRepository(dbWrite, dbRead *sql.DB) ProductRepository {
	return &productRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
