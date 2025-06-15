package main

import (
	"aia_backend/global"
	"aia_backend/handler"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	global.MustSetup(ctx)
	addr := ":9096"
	baseRoute := mux.NewRouter()
	router := baseRoute.PathPrefix("/v1").Subrouter()
	router.HandleFunc("/hello", handler.PingHandler)
	// 常见问题 crud
	router.Path("/frequent_questions").Methods(http.MethodPost).HandlerFunc(handler.UploadFrequentQuestionHandler)
	router.Path("/frequent_questions").Methods(http.MethodGet).HandlerFunc(handler.ListFrequentQuestionsHandler)
	//router.Path("/frequent_questions").Methods(http.MethodDelete).HandlerFunc(handler.PingHandler)
	//router.Path("/frequent_questions").Methods(http.MethodPut).HandlerFunc(handler.PingHandler)

	//router.Path("/product_poster").Methods(http.MethodPost).HandlerFunc(handler.PingHandler)
	//router.Path("/product_poster").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	//router.Path("/products/{category}").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	router.Path("/products").Methods(http.MethodPost).HandlerFunc(handler.UploadProductFileHandler)
	router.Path("/products").Methods(http.MethodGet).HandlerFunc(handler.ListProductsHandler)
	router.Path("/products/{product_id}").Methods(http.MethodGet).HandlerFunc(handler.DownloadProductFileHandler)

	// 核保和理赔应该怎么弄 也放进 products 里面？
	router.Path("/videos").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	router.Path("/videos").Methods(http.MethodPost).HandlerFunc(handler.PingHandler)
	router.Path("/video/{video_id}").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	router.Path("/video/{video_id}").Methods(http.MethodDelete).HandlerFunc(handler.PingHandler)

	router.Path("/team_members").Methods(http.MethodPost).HandlerFunc(handler.PingHandler)
	router.Path("/team_members/{id}").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	router.Path("/team_members/{id}").Methods(http.MethodDelete).HandlerFunc(handler.PingHandler)
	router.Path("/team_members/{id}").Methods(http.MethodPut).HandlerFunc(handler.PingHandler)

	router.Path("/team_stars").Methods(http.MethodGet).HandlerFunc(handler.PingHandler)
	router.Path("/team_stars").Methods(http.MethodPut).HandlerFunc(handler.PingHandler)

	server := http.Server{
		Addr:    addr,
		Handler: baseRoute,
	}

	go func() {
		fmt.Println("start server")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case s := <-stop:
		log.Println("Recv signal %s shut down...", s)
		_ = server.Shutdown(ctx)
	}
}
