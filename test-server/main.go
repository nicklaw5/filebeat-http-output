package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

func main() {
	port := 8888
	host := "0.0.0.0"

	flag.IntVar(&port, "port", port, "the server port")
	flag.StringVar(&host, "host", host, "the server host")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandleFunc)

	s := http.Server{
		Addr:         net.JoinHostPort("", strconv.Itoa(port)),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func testHandleFunc(rw http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dump))
	rw.Write([]byte(`{"success":"true"}`))
}
