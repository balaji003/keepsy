package categories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(ctx context.Context, category *Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockRepo) List(ctx context.Context) ([]*Category, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Category), args.Error(1)
}

func TestListCategories(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo)

		expected := []*Category{{ID: 1, Name: "C1"}, {ID: 2, Name: "C2"}}
		mockRepo.On("List", mock.Anything).Return(expected, nil)

		categories, err := service.ListCategories(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expected, categories)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo := new(MockRepo)
		service := NewService(mockRepo) // Create new service/mock to reset expectations or use independent test

		mockRepo.On("List", mock.Anything).Return(nil, assert.AnError)

		_, err := service.ListCategories(context.Background())
		assert.Error(t, err)
	})
}
