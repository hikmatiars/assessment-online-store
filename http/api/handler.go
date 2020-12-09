package api

import (
	"assessment-online-store/http/request"
	"assessment-online-store/usecase"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase usecase.InterfaceUseCase
}

var (
	InputAddCart request.AddToCart
)

func (h *Handler) GetListProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := h.Usecase.GetListProductUseCase()

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success",
		"code_type"    : "success",
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
			"code_type"    : "success",
			"data"         : nil,
		})
		return
	}

	code, err = h.Usecase.AddCartUseCase( InputAddCart )
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code" : code,
			"code_message" : err.Error(),
			"code_type"    : "success",
			"data"         : nil,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success add to cart",
		"code_type"    : "success",
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
			"code_type"    : "empty",
			"data"         : []map[string]interface{}{},
		})
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusOK,
		"code_message" : "Success",
		"code_type"    : "success",
		"data"         : resp,
	})
}