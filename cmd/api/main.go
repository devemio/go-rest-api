package main

import (
	"encoding/json"
	"github.com/devemio/go-rest-api/internal/infrastructure/middleware"
	"log"
	"net/http"
)

func main() {
	log.Println("Start application")
	http.HandleFunc("/", middleware.Cors(middleware.ContentType(handler)))

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	//app := application.New(config.New())
	//if err := app.Start(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//c := collection.New()
	//c.Put("test", 1)
	//fmt.Println(c.Contains("test"))
}

type DtoOut struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("API call")
	dto := &DtoOut{
		Message: "success",
	}
	json.NewEncoder(w).Encode(dto)
	log.Println("API call END")
}
