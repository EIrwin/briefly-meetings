package services

import (
	"github.com/eirwin/briefly-meetings/data"
	"log"
	"time"
)

func Create(req CreateMeeting) (data.Meeting, error) {
	var meeting data.Meeting
	err := data.CreateMeeting(&meeting)
	if err != nil {
		log.Fatal(err)
	}
	return meeting, nil
}

func Get(req GetMeeting) (data.Meeting, error) {
	meeting, err := data.GetMeeting(req.MeetingId)
	if err != nil {
		log.Fatal(err)
	}

	if meeting.Active {

		meeting, err = updateTotal(meeting)
		if err != nil {
			log.Fatal(err)
		}
	}
	return meeting, nil
}

func Start(req StartMeeting) (data.Meeting, error) {
	meeting, err := data.GetMeeting(req.MeetingId)
	if err != nil {
		log.Fatal(err)
	}

	meeting.Start = time.Now().UTC()
	meeting.Active = true
	err = data.UpdateMeeting(&meeting)
	if err != nil {
		log.Fatal(err)
	}
	return meeting, nil
}

func End(req EndMeeting) (data.Meeting, error) {
	meeting, err := data.GetMeeting(req.MeetingId)
	if err != nil {
		log.Fatal(err)
	}

	meeting.End = time.Now().UTC()
	meeting.Active = false
	err = data.UpdateMeeting(&meeting)
	if err != nil {
		log.Fatal(err)
	}
	return meeting, nil
}

func Join(req JoinMeeting) (data.Meeting, error) {

	//get user using user id

	meeting, err := data.GetMeeting(req.MeetingId)
	if err != nil {
		log.Fatal(err)
	}

	//add user to meeting
	//meeting.Users = append(meeting.Users)

	err = data.UpdateMeeting(&meeting)
	if err != nil {
		log.Fatal(err)
	}

	return meeting, nil
}

func Leave(req LeaveMeeting) (data.Meeting, error) {

	//get user using user id

	//	var user = await _userService.GetUserById(new GetUserById() { Id = leaveMeeting.UserId });
	//
	//var meeting = _meetingRepository.AsQueryable().FirstOrDefault(p => p.Id == leaveMeeting.MeetingId);
	//
	//var attendees = new List<User>();
	//foreach (User attendee in meeting.Attendees)
	//{
	//if (attendee.Id != user.Id)
	//attendees.Add(new User()
	//{
	//Id = attendee.Id, HourlyRate = attendee.HourlyRate, Salary = attendee.Salary
	//});
	//}
	//meeting.Attendees = attendees;
	//meeting = _meetingRepository.Save(meeting);
	return data.Meeting{}, nil
}

func updateTotal(m data.Meeting) (data.Meeting, error) {
	total := m.Total
	minutesPassed := time.Now().UTC().Sub(m.Start).Minutes()
	hourlyPercent := minutesPassed / 60

	for _, val := range m.Users {
		hourlyRate, err := getHourlyRate(val)
		if err != nil {
			log.Fatal(err)
		}
		total += hourlyRate * float32(hourlyPercent)
		m.Total = total
		data.UpdateMeeting(&m)
	}

	return m, nil
}

func getHourlyRate(u data.MeetingUser) (float32, error) {
	const hours = 2080 //work hours per year
	if u.Salary > 0 {
		return (u.Salary / hours), nil
	}
	return u.HourlyRate, nil
}

type CreateMeeting struct {
}

type GetMeeting struct {
	MeetingId string
}

type StartMeeting struct {
	MeetingId string
}

type EndMeeting struct {
	MeetingId string
}

type JoinMeeting struct {
	MeetingId string
	UserId    string
}

type LeaveMeeting struct {
	MeetingId string
	UserId    string
}
