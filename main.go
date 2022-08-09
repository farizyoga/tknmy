package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tokenomy/helper"
	"tokenomy/repository"
)

func main() {
	http.HandleFunc("/", HandleRoot)

	fmt.Print("Web Running on Port 8000")
	http.ListenAndServe(":8000", nil)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStrings := query["id"]
	idInts := []int{}

	if idStrings == nil {
		helper.CreateResponse(w, 200, repository.BookResource.FindAll(), "")
		return
	}

	for _, id := range idStrings {
		for _, num := range strings.Split(id, ",") {
			i, err := strconv.Atoi(num)

			if err != nil {
				helper.CreateResponse(w, 400, nil, "invalid or empty ID "+idStrings[0])
				return
			}

			idInts = append(idInts, i)
		}
	}

	books := repository.BookResource.Finds(idInts)

	if books == nil {
		helper.CreateResponse(w, 404, nil, "resource with ID "+idStrings[0]+" not exist")
		return
	}

	helper.CreateResponse(w, 200, books, "")
}
