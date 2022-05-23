package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	dbInit()

	router := mux.NewRouter()
	router.HandleFunc("/player/count", func(w http.ResponseWriter, r *http.Request) {
		count, err := getCount()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		writeResult(w, http.StatusOK, count)
	}).Methods("GET")

	router.HandleFunc("/player/limit/{limit_size}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sLimitSize := vars["limit_size"]

		limitSize, err := strconv.Atoi(sLimitSize)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		players, err := getPlayerByLimit(limitSize)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeResult(w, http.StatusOK, players)
	}).Methods("GET")

	router.HandleFunc("/player/", func(w http.ResponseWriter, r *http.Request) {
		baBody, err := ioutil.ReadAll(r.Body)
		var players []Player

		err = json.Unmarshal(baBody, &players)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = createPlayers(players)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeResult(w, http.StatusOK, len(players))
	}).Methods("POST")

	router.HandleFunc("/player/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		player, err := getPlayerByID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeResult(w, http.StatusOK, player)
	}).Methods("GET")

	router.HandleFunc("/player/trade", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		form := r.PostForm
		sellID := form.Get("sellID")
		buyID := form.Get("buyID")
		sAmount := form.Get("amount")
		sPrice := form.Get("price")

		amount, err := strconv.Atoi(sAmount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		price, err := strconv.Atoi(sPrice)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = buyGoods(sellID, buyID, amount, price)
		if err != nil {
			writeResult(w, http.StatusBadRequest, false)
			return
		}

		writeResult(w, http.StatusOK, true)
	}).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func writeResult(w http.ResponseWriter, status int, response interface{}) {
	baResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(baResponse)
}
