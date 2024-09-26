package app_test

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/mwives/hexagonal-architecture/app"
	"github.com/stretchr/testify/assert"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = app.DISABLED

	err := product.Enable()

	assert.Nil(t, err)
}

func TestProduct_EnableWithNegativePrice(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Price = -10
	product.Status = app.DISABLED

	err := product.Enable()

	assert.NotNil(t, err)
	assert.Equal(t, app.ErrZeroPriceForEnable.Error(), err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Price = 0
	product.Status = app.ENABLED

	err := product.Disable()

	assert.Nil(t, err)
}

func TestProduct_DisableWithPositivePrice(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = app.ENABLED

	err := product.Disable()

	assert.NotNil(t, err)
	assert.Equal(t, app.ErrNonZeroPriceForDisable.Error(), err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Product 1"
	product.Price = 10
	product.Status = app.ENABLED

	isValid, err := product.IsValid()

	assert.True(t, isValid)
	assert.Nil(t, err)
}

func TestProduct_IsValidWithWrongStatus(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Product 1"
	product.Price = 10
	product.Status = "INVALID_STATUS"

	isValid, err := product.IsValid()

	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Equal(t, app.ErrInvalidStatus.Error(), err.Error())
}

func TestProduct_IsValidWithNegativePrice(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Product 1"
	product.Price = -10
	product.Status = app.ENABLED

	isValid, err := product.IsValid()

	assert.False(t, isValid)
	assert.NotNil(t, err)
	assert.Equal(t, app.ErrNegativePrice.Error(), err.Error())
}
