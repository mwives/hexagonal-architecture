package cli

import (
	"fmt"

	"github.com/mwives/hexagonal-architecture/app"
)

func Run(
	service app.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	productPrice float64,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product created: ID=%s, Name=%s, Price=%f",
			product.GetID(), product.GetName(), product.GetPrice(),
		)

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", res.GetID())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", res.GetID())

	case "get":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product: ID=%s, Name=%s, Price=%f, Status=%s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus(),
		)

	default:
		result = "Invalid action"
	}

	return result, nil
}
