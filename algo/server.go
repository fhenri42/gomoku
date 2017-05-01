package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Data struct {
	Board [][]int `json:"board"`
	PlayerScore int `json:"playerScore"`
	AiScore int `json:"aiScore"`
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

	log.Println(data)
	if len(data.Board) == 19 {
		var response Data
		response.Board, response.AiScore, response.PlayerScore = getBestState(data.Board, data.AiScore, data.PlayerScore)

		jData, err := json.Marshal(response)

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
		initSdl()
	//	http.HandleFunc("/aiturn", parse)
	//	http.ListenAndServe(":5000", nil)
	}
