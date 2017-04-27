package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Data struct {
	Board [][]int `json:"board"`
}

func parse(w http.ResponseWriter, req *http.Request) {

	if origin := req.Header.Get("Origin"); origin != "" {
		 w.Header().Set("Access-Control-Allow-Origin", origin)
		 w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		 w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		 w.Header().Set("Content-Type", "application/json")
 }

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		return
	}

	tabToSend := getBestState(data.Board)
	jData, err := json.Marshal(tabToSend)

	if err != nil {
		log.Println("err  Marshal",err)
		return
	}
	w.Write(jData)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/aiturn", parse)
	http.ListenAndServe(":5000", nil)
}
