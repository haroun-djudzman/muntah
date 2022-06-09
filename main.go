package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type mulut struct {
	muntahSize [][]byte
}

func (m *mulut) muntah(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	indexParam := params.Get("index")
	index, _ := strconv.Atoi(indexParam)

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

	http.HandleFunc("/", mulut.muntah)
	http.ListenAndServe(":6666", nil)
}
