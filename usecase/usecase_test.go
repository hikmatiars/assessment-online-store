package usecase

import (
	"assessment-online-store/entity"
	"assessment-online-store/http/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type TestSuite struct {
	suite.Suite
	Usecase     InterfaceUseCase
	Inventories []*entity.ProductInventory
	Cart        []*entity.Cart
	FlashTime   time.Time
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestGetListProductSuccess() {
	s.Inventories = entity.SeedData()
	s.Usecase = NewUseCase(nil, s.Inventories, s.Cart, s.FlashTime)

	products := s.Usecase.GetListProductUseCase()
	s.Suite.NotEmpty(products)
}

func (s *TestSuite) TestGetListProductFailed() {
	flash := time.Time{}

	s.Usecase = NewUseCase(nil, nil, nil, flash)
	products := s.Usecase.GetListProductUseCase()

	s.Suite.Empty(products)
}

func (s *TestSuite) TestAddCartSuccess() {
	s.Inventories = entity.SeedData()
	s.Usecase = NewUseCase(nil, s.Inventories, s.Cart, s.FlashTime)

	req := request.AddToCart{
		ProductId: 1,
		Quantity:  1,
	}

	_, err  := s.Usecase.AddCartUseCase(req)
	_, cart := s.Usecase.GetListCartUseCase()

	assert.NotEmpty(s.T(),cart)
	assert.NoError(s.T(), err)
}

func (s *TestSuite) TestAddCartFailed() {
	s.Inventories = entity.SeedData()
	s.Usecase = NewUseCase(nil, s.Inventories, s.Cart, s.FlashTime)
	req := request.AddToCart{
		ProductId: 100,
		Quantity: 10,
	}

	code, err := s.Usecase.AddCartUseCase(req)
	_, cart := s.Usecase.GetListCartUseCase()

	//check error if code other than status no content
	if code != http.StatusNoContent {
		assert.NoError(s.T(), err)
	}

	assert.Empty(s.T(),cart)
}

func (s *TestSuite) TestCheckoutItemSuccess() {
	var (
		products []request.ObjProductCheckout
		data []map[string]interface{}
		err error
	)

	s.Inventories = entity.SeedData()
	s.Usecase = NewUseCase(nil, s.Inventories, s.Cart, s.FlashTime)
	req := request.AddToCart{
		ProductId: 1,
		Quantity: 1,
	}

	_, err = s.Usecase.AddCartUseCase(req)
	assert.NoError(s.T(), err)

	products = append(products, request.ObjProductCheckout{
		ProductId: 1,
		Quantity:  1,
	})

	reqCheckout := request.Checkout{
		Products: products,
	}

	_, data, err = s.Usecase.CheckoutItemUseCase( reqCheckout )
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), data)
}