package products

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(ctx context.Context, product *Product) error {
	args := m.Called(ctx, product)
	if args.Get(0) == nil {
		product.ID = 1
		return nil
	}
	return args.Error(0)
}

func (m *MockRepo) GetByID(ctx context.Context, id int) (*Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Product), args.Error(1)
}

func (m *MockRepo) ListByUserID(ctx context.Context, userID int) ([]*Product, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		req := CreateProductRequest{
			UserID: 1,
			Name:   "Test Product",
		}

		mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(p *Product) bool {
			return p.UserID == req.UserID && p.Name == req.Name
		})).Return(nil)

		product, err := service.CreateProduct(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, 1, product.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("MissingUserID", func(t *testing.T) {
		service := NewService(nil)
		_, err := service.CreateProduct(context.Background(), CreateProductRequest{Name: "P"})
		assert.Error(t, err)
		assert.Equal(t, "user ID is required", err.Error())
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		expected := &Product{ID: 1, Name: "P"}
		mockRepo.On("GetByID", mock.Anything, 1).Return(expected, nil)

		product, err := service.GetProduct(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, expected, product)
	})

	t.Run("InvalidID", func(t *testing.T) {
		service := NewService(nil)
		_, err := service.GetProduct(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, "invalid product ID", err.Error())
	})
}

func TestListProducts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		expected := []*Product{{ID: 1}, {ID: 2}}
		mockRepo.On("ListByUserID", mock.Anything, 1).Return(expected, nil)

		products, err := service.ListProducts(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, expected, products)
	})
}
