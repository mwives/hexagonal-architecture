package cli_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mwives/hexagonal-architecture/adapters/cli"
	mock_app "github.com/mwives/hexagonal-architecture/app/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productID := "1"
	productName := "Product 1"
	productPrice := 9.99
	productStatus := "enabled"

	productMock := mock_app.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_app.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	t.Run("Create", func(t *testing.T) {
		result, err := cli.Run(service, "create", "", productName, productPrice)

		assert.Nil(t, err)
		assert.Equal(t, "Product created: ID=1, Name=Product 1, Price=9.990000", result)
	})

	t.Run("Enable", func(t *testing.T) {
		result, err := cli.Run(service, "enable", productID, "", 0)

		assert.Nil(t, err)
		assert.Equal(t, "Product 1 has been enabled", result)
	})

	t.Run("Disable", func(t *testing.T) {
		result, err := cli.Run(service, "disable", productID, "", 0)

		assert.Nil(t, err)
		assert.Equal(t, "Product 1 has been disabled", result)
	})

	t.Run("Get", func(t *testing.T) {
		result, err := cli.Run(service, "get", productID, "", 0)

		assert.Nil(t, err)
		assert.Equal(t, "Product: ID=1, Name=Product 1, Price=9.990000, Status=enabled", result)
	})
}
