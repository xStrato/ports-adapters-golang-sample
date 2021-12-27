package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/ports-adapters-go-sample/mocks"
	"github.com/xStrato/ports-adapters-go-sample/src/adapters/cli"
	"github.com/xStrato/ports-adapters-go-sample/src/application/entities"
)

func TestProduct(t *testing.T) {

	t.Run("Run_CreateValidProduct_ShouldReturnNil", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "cba"
		name := "PS6"
		status := "Disabled"
		price := 59.99

		product := entities.NewProductFrom(id, name, status, float32(price))

		service := mocks.NewMockIProductService(ctrl)
		service.EXPECT().Create(name, float32(price)).Return(product, nil).AnyTimes()
		service.EXPECT().Get(id).Return(product, nil).AnyTimes()
		service.EXPECT().Enable(gomock.Any()).Return(product, nil).AnyTimes()
		service.EXPECT().Disable(gomock.Any()).Return(product, nil).AnyTimes()

		expedtedResult := fmt.Sprintf("Product created!! Id: %s, Name: %s, Price: %f, Status: %s", id, name, float32(price), status)
		//Act
		result, err := cli.Run(service, "create", "", name, float32(price))
		//Assert
		require.Nil(t, err)
		require.Equal(t, expedtedResult, result)
	})
}
