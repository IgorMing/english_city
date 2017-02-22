package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to index")
	foo()
}

func roomsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		panic(err)
	}
}

func roomsByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "Room Id: ", id)
}

func addRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room Room
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &room); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF=8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	roomResponse, err := addRoom(room.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(roomResponse); err != nil {
		panic(err)
	}
}

func deleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity)
		return
	}

	_err := deleteRoom(id)
	if _err != nil {
		http.Error(w, _err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
