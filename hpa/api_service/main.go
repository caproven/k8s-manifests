package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.Logger

var hostname string

func main() {
	logger, _ = zap.NewDevelopment()

	port := os.Getenv("PORT")
	if port == "" {
		logger.Info("PORT env not set, using default port")
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	hostname, _ = os.Hostname()

	router := mux.NewRouter()
	router.HandleFunc("/hello", sayHello).Methods("GET")
	router.HandleFunc("/hello/{name}", sayHelloTo).Methods("GET")

	logger.Info("Starting api",
		zap.String("addr", addr),
	)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Fatal("Failed to serve http",
			zap.Error(err),
		)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(helloStr("world")))
}

func sayHelloTo(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(helloStr(name)))
}

func helloStr(name string) string {
	return fmt.Sprintf("hello %s, from %s", name, hostname)
}
