package main
//

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
  "github.com/parnurzeal/gorequest"
)

type sendBackStruct struct {
	board [20][20]int `json:"board"`
  isWin bool `json:"isWin,string"`
}
type test struct {
	board string `json:"board"`
}

func sendBack(tab [20][20]int, isWin bool) {

	request := gorequest.New()
  var sendTab sendBackStruct
  sendTab.board = tab
  sendTab.isWin = isWin
	_, _, err := request.Post("http://localhost:8080").Send(sendTab).End()
	if err != nil {
		log.Println("err =",err)
		return
	}
  fmt.Print("all good")
}

func parse(w http.ResponseWriter, req *http.Request) {

	if origin := req.Header.Get("Origin"); origin != "" {
			 w.Header().Set("Access-Control-Allow-Origin", origin)
			 w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			 w.Header().Set("Access-Control-Allow-Headers",
					 "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	 }
		fmt.Printf("in the post\n")
		fmt.Printf("%s\n", req)
		body, err := ioutil.ReadAll(req.Body)
    fmt.Print(body)
		if err != nil {
			log.Println("err Request = ", err)
			return
		}
		var data test

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("err MArshall",err)
			return
		}
		fmt.Print(data)
  //  board, iswin := resolution(data.board)
  //  sendBack(board,iswin)

	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/aiturn", parse)
	http.ListenAndServe(":5000", nil)
}
