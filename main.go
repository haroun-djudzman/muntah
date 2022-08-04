package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/haroun-djudzman/muntah/grpc/server"
)

type mulut struct {
	muntahSize [][]byte
}

func (m *mulut) muntah(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")

	params := r.URL.Query()
	indexParam := params.Get("index")
	index, _ := strconv.Atoi(indexParam)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", m.muntahSize[index-1])
}

func main() {
	size := []int{10000, 25000, 50000, 75000, 100000, 200000}
	muntah := [6][]byte{}

	for i, v := range size {
		token := make([]byte, v)
		rand.Read(token)

		muntah[i] = token
	}

	mulut := mulut{
		muntahSize: muntah[:],
	}

	// set up server
	mux := http.NewServeMux()
	mux.HandleFunc("/", mulut.muntah)

	srv := &http.Server{

		Addr:         ":6666",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	srv.SetKeepAlivesEnabled(false)

	go server.StartServer()

	srv.ListenAndServe()
}
