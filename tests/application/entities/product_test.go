package entities_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xStrato/ports-adapters-go-sample/src/application/entities"
)

func TestProduct(t *testing.T) {
	t.Run("Enable_ValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		product := entities.NewProduct("PS5", 5000)
		//Act
		err := product.Enable()
		//Assert
		require.Nil(t, err)
	})

	t.Run("Enable_InvalidProduct_ShouldReturnError", func(t *testing.T) {
		//Arrange
		expectedError := "the price must be greater than zero to enable the product"
		product := entities.NewProduct("PS5", 0)
		//Act
		err := product.Enable()
		//Assert
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("Disable_ValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		product := entities.NewProduct("PS5", 0)
		//Act
		err := product.Disable()
		//Assert
		require.Nil(t, err)
	})

	t.Run("Disable_InvalidProduct_ShouldReturnError", func(t *testing.T) {
		//Arrange
		expectedError := "the price must be zero in order to have the product disabled"
		product := entities.NewProduct("PS5", 5000)
		//Act
		err := product.Disable()
		//Assert
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("IsValid_ValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		product := entities.NewProduct("PS5", 5000)
		//Act
		valid, err := product.IsValid()
		//Assert
		require.Nil(t, err)
		require.True(t, valid)
	})

	t.Run("IsValid_InvalidProduct_ShouldReturnError", func(t *testing.T) {
		//Arrange
		product := entities.NewProduct("PS5", -5000)
		//Act
		valid, err := product.IsValid()
		//Assert
		require.NotNil(t, err)
		require.False(t, valid)
		require.Error(t, err)
	})
}
