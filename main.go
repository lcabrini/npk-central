package main

import (
    "os"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lcabrini/npk-common"
)

func main() {
    if _, err := npk.DBConnection(config); err != nil {
        log.Printf("database: %v", err)
        os.Exit(1)
    }

    log.Printf("connected to database")
    npk.SetupSessionStore(config)
    r := mux.NewRouter()
    npk.SetupRoutes(r)
    http.ListenAndServe(config.Listen, r)
}
