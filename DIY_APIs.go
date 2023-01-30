package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *App) productList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	list, err := getList(a.DB, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	fmt.Println(list)
	respondWithJSON(w, http.StatusOK, list)
}

func (a *App) addProduct(w http.ResponseWriter, r *http.Request) {
	var p store
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &p)
	}

	defer r.Body.Close()

	if err := p.addProduct(a.DB, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, p)
}
