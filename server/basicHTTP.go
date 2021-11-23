package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)

	portNum := ":3000"
	fmt.Println("server is listening on port:", portNum)
	err := http.ListenAndServe(portNum, nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("w:", w)
	fmt.Println("req:", req)
}
