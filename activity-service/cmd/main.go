package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "activity-service/internal/activity"
    "activity-service/internal/config"
)

func main() {
    cfg := config.LoadConfig()
    r := mux.NewRouter()

    activityService := activity.NewService()
    activityHandler := activity.NewHandler(activityService)

    r.HandleFunc("/activity", activityHandler.CreatActivity).Methods("POST")
    r.HandleFunc("/activity/{id:[0-9]+}", activityHandler.GetActivity).Methods("GET")
    r.HandleFunc("/activity/{id:[0-9]+}", activityHandler.UpdateActivity).Methods("PUT")
    r.HandleFunc("/activity/{id:[0-9]+}", activityHandler.DeleteActivity).Methods("DELETE")

    log.Fatal(http.ListenAndServe(cfg.ServerAddress, r))
}