package main

import (
	"encoding/json"
	"github.com/iluxaorlov/covid2.0/method"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	dataList := method.TakeData()

	js, err := json.Marshal(dataList)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(js)
	if err != nil {
		panic(err.Error())
	}
}
