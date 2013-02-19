package coverage

import (
	"labix.org/v2/mgo/bson"
	"net/url"
	"time"
)

type Feed struct {
	ID       bson.ObjectId
	Title    string
	URL      url.URL
	Articles []Article `bson:-`
	Logs     LogEntries
	Times    struct {
		Added     time.Time
		Updated   time.Time
		LastCheck time.Time
	}
}
