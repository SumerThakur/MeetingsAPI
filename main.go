/* Meetings API developed by - Sumer Thakur*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type participant struct {
	participantName  string `json:"name" bson:"name,omitempty"`
	participantEmail string `json:"email" bson:"email,omitempty"`
	participantRsvp  string `json:"rsvp" bson:"rsvp,omitempty"`
}

type meeting struct {
	meetingId           string `json:"id" bson:"id,omitempty"`
	meetingTitle        string `json:"title" bson:"Title,omitempty"`
	meetingParticipants string `json:"participants" bson:"Participants,omitempty"`
	meetingDate         string `json:"date" bson:"date,omitempty"`
	meetingStartTime    string `json:"startTime" bson:"startTime,omitempty"`
	meetingEndTime      string `json:"endTime" bson:"endTime,omitempty"`
	meetingTimeStamp    string `json:"timeStamp" bson:"timeStamp,omitempty"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Method Restricted"}`))
	}
}

func scheduleMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Method Restricted"}`))
	}
}

func meetingHandler(w http.ResponseWriter, r *http.Request) {
	titleId := r.URL.Path[len("/meeting/"):]
	meetingDatabaseSearch(titleId);
}

func meetingTimeHandler(w http.ResponseWriter, r *http.Request) {
	titleStartTime := r.URL.Path[len("meetings?start=/"):("&end=")]
	titleEndTime := r.URL.Path[len("&end"):]
	meetingTimeSearch(titleStartTime, titleEndTime string);
}

func meetingEmailHandler(w http.ResponseWriter, r *http.Request) {
	titleEmail := r.URL.Path[len("/meetings?participant="):]
	meetingEmailSearch(titleEmail string)
}

func meetingDatabaseSearch(titleId string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("mongodb+srv://user_1:SumerMeetings@meetings.yyoqg.mongodb.net/Meetings?retryWrites=true&w=majority")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	meetingDatabase := client.Database("Meetings")
	meetingInfoDatabase := Meetings.Collection("meetingInfo")

	var meetingIdSearchResult []meeting
	cursor, err := meeetingInfoDatabase.Find(ctx, bson.M{"meetingId": bson.D{titleId}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &meetingIdSearchResult); err != nil {
		panic(err)
	}
	fmt.Println(meetingIdSearchResult)

}

func meetingTimeSearch(titleStartTime, titleEndTime string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("mongodb+srv://user_1:SumerMeetings@meetings.yyoqg.mongodb.net/Meetings?retryWrites=true&w=majority")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	meetingDatabase := client.Database("Meetings")
	meetingInfoDatabase := Meetings.Collection("meetingInfo")

	var meetingTimeSearchResult []meeting
	cursor, err := meeetingInfoDatabase.Find(ctx, bson.M{"meetingStartTime": bson.D{titleStartTime}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &meetingIdSearchResult); err != nil {
		panic(err)
	}
	

}



func meetingEmailSearch(titleEmail string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("mongodb+srv://user_1:SumerMeetings@meetings.yyoqg.mongodb.net/Meetings?retryWrites=true&w=majority")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	meetingDatabase := client.Database("Meetings")
	participantInfoDatabase := Meetings.Collection("participantInfo")

	var meetingEmailSearchResult []meeting
	cursor, err := participantInfoDatabase.Find(ctx, bson.M{"participantEmail": bson.D{titleEmail}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &meetingIdSearchResult); err != nil {
		panic(err)
	}
	

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/meetings", scheduleMeeting)
	http.HandleFunc("/meeting/", meetingHandler)
	http.HandleFunc("/meetings?start=<start time here>&end=<end time here>", meetingTimeHandler)
	http.HandleFunc("/meetings?participant=<email id>", meetingEmailHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
	
}
