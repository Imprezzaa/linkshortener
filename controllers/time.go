package controllers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get time is a helper function that grabs the current time and date and converts it to a
// mongoDB primitive. Timestamp will be int64 unix timestamp similar to time.Now().Unix()
func GetTime() primitive.DateTime {
	t := time.Now()

	return primitive.NewDateTimeFromTime(t)
}

// ReturnUserTime takes in a unix timestamp from the DB and returns it in a human readable
// format based on the users location or returns it in UTC if no location found
// TODO: implement function
func ReturnUserTime(p primitive.DateTime, l time.Location) {
	//
}
