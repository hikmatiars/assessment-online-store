package router

import (
	"assessment-online-store/http/api"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, apiHandler *api.Handler) http.Handler {
	//init route
	route := mux.NewRouter()

	route.HandleFunc("/api/check-health", func(w http.ResponseWriter, request *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"ok": true,
		})
	})

	route.HandleFunc("/api/list-product", apiHandler.GetListProductHandler).Methods("GET")

	return route
}
