package main

import (
	"github.com/lactobasilusprotectus/go-template/config"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	rootHanlder "github.com/lactobasilusprotectus/go-template/pkg/domain/root/delivery"
	server "github.com/lactobasilusprotectus/go-template/pkg/util/http"
)

func main() {
	//ambil nilai .env
	env := os.Getenv("APP_ENV")

	config, err := config.Read()

	if err != nil {
		log.Fatal("read config err: ", err)
	}

	//
	httpHandler := initHttpHandler(env)

	//
	utility := initUtility(config)

	//
	registerHttpHandler(utility.Server, httpHandler)

	utility.Server.Run(env)

	//Menangkap sinyal dari system
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-c

	utility.Server.Stop()

	log.Println("shutting down server")
	os.Exit(0)
}

func initHttpHandler(env string) AppHttpHandler {
	return AppHttpHandler{
		RootHttpHandler: rootHanlder.NewRootHandler(env),
	}
}

func initUtility(cfg config.Config) AppUtility {
	return AppUtility{
		Server: server.NewServer(cfg.Http),
	}
}

//registerHttpHandler => fungsi untuk mendaftarkan api
func registerHttpHandler(srv *server.Server, handlers AppHttpHandler) {
	h := reflect.ValueOf(handlers)

	for i := 0; i < h.NumField(); i++ {
		srv.RegisterHandler(h.Field(i).Interface().(server.RouteHandler))
	}
}

//AppUtility bungkus semua utility app di struct ini
type AppUtility struct {
	Server *server.Server
}

//AppHttpHandler bungkus semua handler http kedalam struct ini
type AppHttpHandler struct {
	RootHttpHandler *rootHanlder.RootHandler
}
