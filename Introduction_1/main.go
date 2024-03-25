package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/WayneMusungu/Microservices.git/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.SetFlags())
	hh := handlers.NewHello(l)

	http.ListenAndServe(":9090", nil)
}