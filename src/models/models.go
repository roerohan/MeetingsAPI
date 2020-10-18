package models

// Participant is a struct defining the participant object
type Participant struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	RSVP  string `bson:"rsvp" json:"rsvp"`
}

// Meeting is a struct defining the meeting object
type Meeting struct {
	ID           string        `bson:"id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Participants []Participant `bson:"participants" json:"participants"`
	StartTime    int64         `bson:"startTime" json:"startTime"`
	EndTime      int64         `bson:"endTime" json:"endTime"`
	CreatedAt    int64         `bson:"createdAt" json:"createdAt"`
}
