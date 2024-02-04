package service

import (
	"context"

	database "github.com/alanfranciscos/gRPC-GO/internal/databases"
	"github.com/alanfranciscos/gRPC-GO/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDb,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c * CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()

	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category
	
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categoryList,
	}, nil
}