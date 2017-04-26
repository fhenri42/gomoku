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
  isWin bool   `json:"isWin,string"`
}

func sendBack(tab [20][20]int, isWin bool) {

	request := gorequest.New()
  var sendTab sendBackStruct
  sendTab.board = tab
  sendTab.isWin = isWin
	_,_,err := request.Post("https://localhost:8080").Send(sendTab).End()
	if err != nil {
		log.Println(err)
		return
	}
  fmt.Print("all good")
	// log.Println("\nrep = ", rep)
	// log.Println("\nbody = ", body)

}

func parse(w http.ResponseWriter, req *http.Request) {

		fmt.Printf("in the post\n")
		fmt.Printf("%s\n", req)
		body, err := ioutil.ReadAll(req.Body)
    fmt.Print(body)
		if err != nil {
			log.Println(err)
			return
		}
		var data sendBackStruct
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("err",err)
			return
		}
		sendBack(data.board, false)

	w.WriteHeader(200)
}

func main() {

	http.HandleFunc("/aiturn", parse)
	http.ListenAndServe(":5000", nil)
}
