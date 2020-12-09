package main

import (
	"assessment-online-store/entity"
	"assessment-online-store/http/api"
	"assessment-online-store/router"
	"assessment-online-store/usecase"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var TimeFlashSale time.Time

func init()  {
	entity.SeedData()
	TimeFlashSale = time.Now().Local().Add(time.Minute * 1)
	log.Printf("Time flash sale at %v", TimeFlashSale.Format("Jan-02-06 15:04:05"))
}

func main() {
	var httpAddr = flag.String("http", ":"+"8080", "HTTP Listen address")
	//init context
	ctx := context.Background()
	uc := usecase.NewUseCase(ctx, entity.Inventories, entity.Carts, TimeFlashSale)

	//make error channel
	errs := make(chan error)

	confHandler := &api.Handler{Usecase: uc}
	routeHttp := router.NewHttpServer(ctx, confHandler)

	go func() {
		fmt.Println("listening on port ", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, routeHttp)
	}()

	//Print if conditional error
	log.Printf("error %v", <-errs)
}
