package api

import(
	"assessment-online-store/usecase"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase usecase.InterfaceUseCase
}

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