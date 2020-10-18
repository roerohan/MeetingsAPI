package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/roerohan/MeetingsAPI/src/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (router *RouteHandler) getMeetingByID(w http.ResponseWriter, r *http.Request) {
	id := utils.ExtractParam(r)

	var meeting bson.M

	err := router.Meeting.FindOne(context.TODO(), bson.D{{
		"id", id,
	}}).Decode(&meeting)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

// MeetingController handles requests to /meetings
func (router *RouteHandler) MeetingController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		router.getMeetingByID(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
