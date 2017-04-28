package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Data struct {
	Board [][]int `json:"board"`
	PlayerScore int
	AiScore int
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
	log.Println(data.AiScore)

	if len(data.Board) == 20 {
		tabToSend := getBestState(data.Board, data.AiScore, data.PlayerScore)

		jData, err := json.Marshal(tabToSend)

		if err != nil {
			log.Println("err  Marshal",err)
			return
		}

		w.Write(jData)

		} else {
			w.WriteHeader(503)
		}
	}

	func main() {
		http.HandleFunc("/aiturn", parse)
		http.ListenAndServe(":5000", nil)
	}
