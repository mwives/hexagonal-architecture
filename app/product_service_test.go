package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mwives/hexagonal-architecture/app"
	mock_app "github.com/mwives/hexagonal-architecture/app/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	productPersistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Get(gomock.All()).Return(product, nil).AnyTimes()

	service := app.ProductService{Persistence: productPersistence}

	result, err := service.Get("123")

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	productPersistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{Persistence: productPersistence}

	result, err := service.Create("product", 1.0)

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	productPersistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{Persistence: productPersistence}

	result, err := service.Enable(product)

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()
	productPersistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{Persistence: productPersistence}

	result, err := service.Disable(product)

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
