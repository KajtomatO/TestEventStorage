package main

import (
	"fmt"
	"net/http"
	"strings"
)

type DataStore interface {
	GetDataRecord(id string) (string, error)
	SetDataRecord(id string, data string)
}

type DataServer struct {
	store DataStore
}

func (p *DataServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")
	switch r.Method {
	case http.MethodPost:
		p.writeData(w, user)
	case http.MethodGet:
		p.readData(w, user)
	}
}

func (p *DataServer) readData(w http.ResponseWriter, user string) {
	data, err := p.store.GetDataRecord(user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, data)
}

func (p *DataServer) writeData(w http.ResponseWriter, input string) {
	inputSplited := strings.SplitAfter(input, "/")
	if len(inputSplited) == 2 {
		user := strings.TrimSuffix(inputSplited[0], "/")
		data := inputSplited[1]
		p.store.SetDataRecord(user, data)
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
