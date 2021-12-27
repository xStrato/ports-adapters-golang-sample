package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xStrato/ports-adapters-go-sample/src/application"
)

func TestProduct(t *testing.T) {
	t.Run("Enable_ValidProduct_ShouldPass", func(t *testing.T) {
		//Arrange
		product := application.Product{}
		//Act
		//Assert
		require.Nil(t, product)
	})
}
