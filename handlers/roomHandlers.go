package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/IgorMing/english_city/database"
	ec "github.com/IgorMing/english_city"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to index")
	database.GetRooms()
}

func RoomsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	rooms, err := database.GetRooms()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		panic(err)
	}
}

func RoomsByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "Room Id: ", id)
}

func AddRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room ec.Room
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF=8")

	if err = json.Unmarshal(body, &room); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	roomResponse, err := database.InsertRoom(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(roomResponse); err != nil {
		panic(err)
	}
}

func DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity)
		return
	}

	err = database.DeleteRoom(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room ec.Room
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusUnprocessableEntity),
			http.StatusUnprocessableEntity)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF=8")

	if err = json.Unmarshal(body, &room); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	updateResponse, err := database.UpdateRoom(id, room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updateResponse); err != nil {
		panic(err)
	}
}
