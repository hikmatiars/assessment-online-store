package api

import (
	"assessment-online-store/http/request"
	"assessment-online-store/usecase"
	"assessment-online-store/util"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase usecase.InterfaceUseCase
}

var (
	InputAddCart request.AddToCart
	InputCheckout request.Checkout
)

func (h *Handler) GetListProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := h.Usecase.GetListProductUseCase()

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success",
		"code_type"    : util.CodeHttp( http.StatusOK ),
		"data"         : resp,
	})
}

func (h *Handler) AddCartHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		code int
	)

	w.Header().Set("Content-Type", "application/json")

	InputAddCart = request.AddToCart{}
	err = json.NewDecoder(r.Body).Decode(&InputAddCart)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code" : http.StatusBadRequest,
			"code_message" : "Success",
			"code_type"    : util.CodeHttp( http.StatusBadRequest ),
			"data"         : nil,
		})
		return
	}

	code, err = h.Usecase.AddCartUseCase( InputAddCart )
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code" : code,
			"code_message" : err.Error(),
			"code_type"    : util.CodeHttp( code ),
			"data"         : nil,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success add to cart",
		"code_type"    : util.CodeHttp( http.StatusOK ),
		"data"         : nil,
	})
}

func (h *Handler) GetListCartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	code, resp := h.Usecase.GetListCartUseCase()
	if code == http.StatusNoContent {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code" : http.StatusNoContent,
			"code_message" : "Empty Cart",
			"code_type"    : util.CodeHttp( http.StatusNoContent ),
			"data"         : []map[string]interface{}{},
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success",
		"code_type"    : util.CodeHttp( http.StatusOK ),
		"data"         : resp,
	})
}

func (h *Handler) CheckoutItemCartHandler( w http.ResponseWriter, r *http.Request ) {
	var (
		err error
		resp []map[string]interface{}
		code int
	)

	w.Header().Set("Content-Type", "application/json")

	InputCheckout = request.Checkout{}
	err = json.NewDecoder(r.Body).Decode(&InputCheckout)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code" : http.StatusBadRequest,
			"code_message" : "Success",
			"code_type"    : util.CodeHttp( http.StatusBadRequest ),
			"data"         : nil,
		})
		return
	}

	code, resp, err = h.Usecase.CheckoutItemUseCase( InputCheckout )
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code"         : code,
			"code_message" : err.Error(),
			"code_type"    : util.CodeHttp( code ),
			"data"         : []map[string]interface{}{},
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" 		   : code,
		"code_message" : "Success",
		"code_type"    : util.CodeHttp( http.StatusOK ),
		"data"         : resp,
	})
}