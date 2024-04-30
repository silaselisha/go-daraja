package example

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/silaselisha/go-daraja/pkg/handler"
)

type Duka struct {
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
}

func dukaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		duka := new(Duka)
		err = json.Unmarshal(data, duka)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		amount, err := strconv.ParseFloat(duka.Amount, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		darajaClient, err := handler.NewDarajaClient("./..")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := darajaClient.NIPush(duka.Description, duka.PhoneNumber, amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_, err = w.Write(res)
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

func Server() {
	http.HandleFunc("/api/v1/duka", dukaHandler)

	fmt.Printf("server starting...\n")
	fmt.Print("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
