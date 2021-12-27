package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	mock_interfaces "github.com/xStrato/ports-adapters-go-sample/mocks"
	"github.com/xStrato/ports-adapters-go-sample/src/application/entities"
	"github.com/xStrato/ports-adapters-go-sample/src/application/services"
)

func TestProductService(t *testing.T) {
	t.Run("Get_ValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		product := entities.NewProduct("PS5", 5000)
		repository := mock_interfaces.NewMockIProductRepository(ctrl)
		repository.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
		//Act
		service := services.NewProductService(repository)
		result, err := service.Get(uuid.NewV4().String())
		//Assert
		require.Nil(t, err)
		require.Equal(t, product, result)
	})
}
