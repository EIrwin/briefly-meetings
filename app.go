package main

import (
	"github.com/eirwin/briefly-meetings/services"
	"log"
)

func main() {
	log.Print("starting meeting service...")
	//router := api.NewRouter()
	//log.Fatal(http.ListenAndServe(":8282", router))
	req := services.CreateMeeting{}
	meeting, err := services.Create(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(meeting)

	get := services.GetMeeting{MeetingId: meeting.Id.Hex()}
	getResult, err := services.Get(get)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(getResult)

	start := services.StartMeeting{MeetingId: meeting.Id.Hex()}
	startResult, err := services.Start(start)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(startResult)

	end := services.EndMeeting{MeetingId: meeting.Id.Hex()}
	endResult, err := services.End(end)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(endResult)

}
