package web

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/szlaci83/HumanityCoins/web/controllers"
)

func Serve(app *controllers.Application) {
    router := mux.NewRouter()
    router.HandleFunc("/register", app.AddUser).Methods("POST")
    router.HandleFunc("/thank/{id}", app.AddThanks).Methods("POST")
    router.HandleFunc("/user/{id}", app.GetByID).Methods("GET")
    router.HandleFunc("/user", app.GetKeys).Methods("GET")
    fmt.Println("Listening (http://localhost:8000/) ...")
    log.Fatal(http.ListenAndServe(":8000", router))
}
