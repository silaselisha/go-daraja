package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/silaselisha/go-daraja/pkg/handler"
)

func dukaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		darajaClient, err := handler.NewDarajaClient("./..")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := darajaClient.NIPush("test", "0792918261", 1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		buff, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		_, err = w.Write(buff)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("%+v method not implemented\n", r.Method)))
	}
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("headers: ", r.Header)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", string(d))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(r.Body)))
}

func main() {
	http.HandleFunc("/api/v1/duka", dukaHandler)
	http.HandleFunc("/api/v1/checkout", checkoutHandler)

	fmt.Printf("server starting...\n")
	fmt.Print("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
