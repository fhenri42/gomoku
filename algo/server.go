package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type sendBackStruct struct {
	board [20][20]int `json:"board"`
}

func parse(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("in the post\n")

	if origin := req.Header.Get("Origin"); origin != "" {
			 w.Header().Set("Access-Control-Allow-Origin", origin)
			 w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			 w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			 w.Header().Set("Content-Type", "application/json")
	 }

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Println("err Request = ", err)
			return
		}

		var data sendBackStruct
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("err Unmrshall",err)
			return
		}
		fmt.Print(data)
		//data.board = getBestState(data.board)

		jData, err := json.Marshal(data.board)
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
