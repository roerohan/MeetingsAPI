package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/roerohan/MeetingsAPI/src/models"
	"github.com/roerohan/MeetingsAPI/src/utils"
)

// RouteHandler stores mongo client
type RouteHandler struct {
	Meeting *mongo.Collection
}

// NewRouteHandler returns a RouteHandler object
func NewRouteHandler(Meeting *mongo.Collection) *RouteHandler {
	return &RouteHandler{
		Meeting: Meeting,
	}
}

func (router *RouteHandler) getMeeting(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	participantArr, noParticipant := query["participant"]

	if !noParticipant {
		startArr, noStart := query["start"]
		endArr, noEnd := query["end"]

		if !noStart || !noEnd {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		start, _ := strconv.Atoi(startArr[0])
		end, _ := strconv.Atoi(endArr[0])

		meetingsCursor, err := router.Meeting.Find(context.TODO(), bson.D{{
			"startTime",
			bson.D{{
				"$gte", start,
			}}}, {
			"endTime",
			bson.D{{
				"$lte", end,
			}},
		}})

		if err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		var meets []bson.M

		if err = meetingsCursor.All(context.TODO(), &meets); err != nil {
			log.Fatal(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(meets)
		return
	}

	participant := participantArr[0]

	unwindStage := bson.D{{
		"$unwind", "$participants",
	}}

	fmt.Println(participant)

	matchStage := bson.D{{
		"$match", bson.M{
			"participants.email": participant,
		},
	}}

	meetingsCursor, err := router.Meeting.Aggregate(context.TODO(), mongo.Pipeline{unwindStage, matchStage})

	var meetings []bson.M
	if err = meetingsCursor.All(context.TODO(), &meetings); err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if len(meetings) != 0 {
		json.NewEncoder(w).Encode(meetings)
		return
	}
}

func (router *RouteHandler) scheduleMeeting(w http.ResponseWriter, r *http.Request) {
	var m models.Meeting

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&m)

	m.ID = utils.RandStringRunes(32)
	m.CreatedAt = time.Now().Unix()

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if m.StartTime > m.EndTime {
		http.Error(w, "startTime be lesser than endTime", http.StatusBadRequest)
	}

	// Aggregation Pipeline
	// 1. Filter out everything that lies in the same time slot
	// currentMeeting.endTime > newMeeting.endTime > currentMeeting.startTime
	// OR
	// currentMeeting.endTime > newMeeting.startTime > currentMeeting.startTime
	emails := make(bson.A, len(m.Participants))
	for i := 0; i < len(m.Participants); i++ {
		emails[i] = m.Participants[i].Email
	}

	unwindStage := bson.D{{
		"$unwind", "$participants",
	}}

	matchStage := bson.D{{
		"$match", bson.M{
			"participants.email": bson.M{
				"$in": emails,
			},
			"participants.rsvp": bson.M{
				"$in": bson.A{"Yes", "Maybe"},
			},
		},
	}}

	checkTimeStage := bson.D{{
		"$match", bson.M{
			"$or": bson.A{
				bson.M{
					"startTime": bson.M{"$lte": m.EndTime},
					"endTime":   bson.M{"$gte": m.EndTime},
				},
				bson.M{
					"startTime": bson.M{"$lte": m.StartTime},
					"endTime":   bson.M{"$gte": m.StartTime},
				},
			},
		},
	}}

	meetingsCursor, err := router.Meeting.Aggregate(context.TODO(), mongo.Pipeline{unwindStage, matchStage, checkTimeStage})

	var meetings []bson.M
	if err = meetingsCursor.All(context.TODO(), &meetings); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if len(meetings) != 0 {
		fmt.Println(meetings)
		json.NewEncoder(w).Encode(meetings)
		return
	}

	_, err = router.Meeting.InsertOne(context.TODO(), m)

	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(m)
}

// MeetingsController handles requests to /meetings
func (router *RouteHandler) MeetingsController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		router.getMeeting(w, r)

	case "POST":
		router.scheduleMeeting(w, r)
	}
}