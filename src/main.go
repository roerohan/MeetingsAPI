package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/roerohan/MeetingsAPI/src/models"
	"github.com/roerohan/MeetingsAPI/src/controllers"
)


func httpListener(port int) {
	portStr := strconv.Itoa(port)

	client := models.MongoConnect("mongodb://localhost:27017")
	router := routes.NewRouteHandler(client.Database("MeetingsAPI").Collection("meetings"))
	
	http.HandleFunc("/meetings", router.MeetingsController)

	fmt.Println("[INFO] Listening on PORT " + portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}


func main() {
	httpListener(3000)
}
