package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/roerohan/MeetingsAPI/src/controllers"
	"github.com/roerohan/MeetingsAPI/src/models"
)

func httpListener(port int) {
	portStr := strconv.Itoa(port)

	client := models.MongoConnect("mongodb://localhost:27017")
	router := controllers.NewRouteHandler(client.Database("MeetingsAPI").Collection("meetings"))

	http.HandleFunc("/meetings", router.MeetingsController)
	http.HandleFunc("/meeting/", router.MeetingController)

	log.Println("[INFO] Listening on PORT " + portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}

func main() {
	httpListener(3000)
}
