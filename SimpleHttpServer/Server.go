package main

import (
	"net/http"
	"runtime"
	"log"
	"fmt"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
			msg_api := []byte("hello world\n")
			w.Write(msg_api)
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	//log.Fatal("Start Server ")
	http.Handle("/", http.StripPrefix("/", http.HandlerFunc(helloworld)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("runing")
}
