package data

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func CreateMeeting(meeting *Meeting) error {
	return nil
}

func GetMeeting(id string) error {
	return nil
}

func UpdateMeeting(Meeting *Meeting) error {
	return nil
}

type Meeting struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Start  time.Time     `json:"start"`
	End    time.Time     `json:"end"`
	Users  []MeetingUser `json:"users"`
	Active bool          `json:"active"`
	Total  float32       `json:"total"`
}

type MeetingUser struct {
	HourlyRate  float32 `json:"hourlyRate"`
	Salary      float32 `json:"salary"`
	IntervalAmt float32 `json:intervalAmt"`
}
