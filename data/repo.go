package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func CreateMeeting(meeting *Meeting) error {
	endpoint, db := getConnectionInfo()

	session, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//Collection Meetings
	c := session.DB(db).C("meetings")

	// Index
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := c.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}

	//Insert
	meeting.Id = bson.NewObjectId()
	err = c.Insert(meeting)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetMeeting(id string) (Meeting, error) {
	endpoint, db := getConnectionInfo()

	session, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB(db).C("meetings")
	var meeting Meeting
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&meeting)
	if err != nil {
		log.Fatal(err)
	}
	return meeting, nil
}

func UpdateMeeting(meeting *Meeting) error {
	endpoint, db := getConnectionInfo()

	session, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//Collection Meetings
	c := session.DB(db).C("meetings")

	// Index
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := c.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}

	// Update
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(meeting.Id.Hex())}, meeting)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func getConnectionInfo() (string, string) {
	endpoint := "mongodb://dev:briefly123!@ds015584.mlab.com:15584/briefly"
	db := "briefly"
	return endpoint, db
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
