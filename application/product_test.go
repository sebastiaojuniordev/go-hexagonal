package application_test

import (
	"github.com/satori/go.uuid"
	"github.com/sebastiaojuniordev/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	require.Empty(t, product.ID)

	productId := uuid.NewV4().String()
	product.ID = productId
	require.Equal(t, productId, product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	require.Empty(t, product.Name)

	product.Name = "Product One"
	require.Equal(t, "Product One", product.Name)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	require.Empty(t, product.Status)

	product.Status = application.DISABLED
	require.Equal(t, application.DISABLED, product.Status)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	require.Equal(t, float64(0), product.Price)

	product.Price = 10
	require.Equal(t, float64(10), product.Price)
}

func TestNewProduct(t *testing.T) {
	product := application.NewProduct()
	require.NotEmpty(t, product.ID)
	require.Equal(t, application.DISABLED, product.Status)
}
