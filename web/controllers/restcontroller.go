package controllers

import (
	"fmt"
	"github.com/szlaci83/HumanityCoins/blockchain"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type Application struct {
 	Fabric *blockchain.FabricSetup
}

type thank struct{
	Thanker    string `json:"name"`		  // person who gives the "thank"
	ThankType  string `json:"type"`		  // number of points given
	Message    string `json:"message"`	// the reason for giving thanks
}

type entity struct {
	UserID    string `json:"name"`
	Balance   string`json:"balance"`		// points received by thanks
}

func (app *Application) AddUser(w http.ResponseWriter, r *http.Request) {
	var person entity
	_ = json.NewDecoder(r.Body).Decode(&person)
	result, err := app.Fabric.InvokeAddUser(person.UserID, person.Balance)
	if err != nil {
		http.Error(w, "Unable to write to blockchain", 500)
	}
	fmt.Print(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}


func (app *Application)  AddThanks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    var tx thank
	_ = json.NewDecoder(r.Body).Decode(&tx)

	result, err := app.Fabric.InvokeAddThanks(id, tx.Thanker, tx.ThankType, tx.Message)
	if err != nil {
		http.Error(w, "Unable to write the blockchain", 500)
	}
	fmt.Print(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (app *Application)  GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := app.Fabric.InvokeGetByID(id)
	if err != nil {
		http.Error(w, "Unable to read the blockchain", 500)
	}
	fmt.Print(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func (app *Application)  GetKeys(w http.ResponseWriter, r *http.Request) {
	result, err := app.Fabric.InvokeGetKeys()
	if err != nil {
		http.Error(w, "Unable to read the blockchain", 500)
	}
	fmt.Print(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
