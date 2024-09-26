package app

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

var (
	ErrInvalidStatus          = errors.New("the status must be either enabled or disabled")
	ErrNegativePrice          = errors.New("the price must be greater than or equal to zero")
	ErrZeroPriceForEnable     = errors.New("the price must be greater than zero to enable the product")
	ErrNonZeroPriceForDisable = errors.New("the price must be zero to disable the product")
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.New().String(),
		Status: DISABLED,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, ErrInvalidStatus
	}
	if p.Price < 0 {
		return false, ErrNegativePrice
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return ErrZeroPriceForEnable
	}

	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	if p.Price > 0 {
		return ErrNonZeroPriceForDisable
	}

	p.Status = DISABLED
	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}
