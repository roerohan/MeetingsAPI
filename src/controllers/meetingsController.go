package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/roerohan/MeetingsAPI/src/models"
	"github.com/roerohan/MeetingsAPI/src/utils"
)

func (router *RouteHandler) getMeeting(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	participantArr, participantExists := query["participant"]
	pageArr, pageExists := query["page"]

	if !participantExists || len(participantArr) < 1 {
		startArr, startExists := query["start"]
		endArr, endExists := query["end"]

		if !startExists || !endExists || len(startArr) < 1 || len(endArr) < 1 {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
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
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		var meetings []bson.M

		if err = meetingsCursor.All(context.TODO(), &meetings); err != nil {
			log.Fatal(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		if pageExists {
			itemsPerPage := 100

			page, _ := strconv.Atoi(pageArr[0])

			if page < 1 {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}
			
			startIdx := (page - 1) * itemsPerPage
			endIdx := startIdx + itemsPerPage

			if startIdx >= len(meetings) {
				json.NewEncoder(w).Encode(models.Meeting{})
				return
			}

			if endIdx > len(meetings) {
				endIdx = len(meetings)
			}

			json.NewEncoder(w).Encode(meetings[startIdx: endIdx])
			return
		}

		json.NewEncoder(w).Encode(meetings)
		return
	}

	participant := participantArr[0]

	unwindStage := bson.D{{
		"$unwind", "$participants",
	}}

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
	// 1. Unwind participants
	// 2. Check if the email matches and RSVP is "Yes" or "Maybe"
	// 3. Filter out everything that lies in the same time slot
	// currentMeeting.endTime > newMeeting.endTime > currentMeeting.startTime
	// OR
	// currentMeeting.endTime > newMeeting.startTime > currentMeeting.startTime

	emails := make(bson.A, len(m.Participants))
	for i := 0; i < len(m.Participants); i++ {
		emails[i] = m.Participants[i].Email

		_, found := utils.Find([]string{"Yes", "No", "Maybe", "Not Answered"}, m.Participants[i].RSVP)
		if !found {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}
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
		log.Println("Conflict occured with the meetings returned in the response.")
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
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
