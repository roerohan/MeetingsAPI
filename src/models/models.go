package models

// Participant is a struct defining the participant object
type Participant struct {
	Name  string
	Email string
	RSVP  string
}

// Meeting is a struct defining the meeting object
type Meeting struct {
	ID           string
	Title        string
	Participants []Participant
	StartTime    int64
	EndTime      int64
	CreatedAt    int64
}
