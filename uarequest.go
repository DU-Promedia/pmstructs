package pmstructs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var UACredentials = map[string]UACrendtial{
	"na.se":  UACrendtial{"6ezSTID3RBG2f1sXoCK7yw", "3FJwMEpqSM2cMZVXLbtzzA"},
	"vlt.se": UACrendtial{"vBGrk0m5R3e-Hc0WxxH_YA", "duZqa-7KSDKGK2nPEdkHnQ"},
	"lt.se":  UACrendtial{"zH1U18UwQh6Ofo6dmdMuRA", "Jp82jgk6SqWj1gKQwoNrZg"},
}

type UACrendtial struct {
	User string
	Pass string
}

func UAInit() {
	// c := UACrendtial{"6ezSTID3RBG2f1sXoCK7yw", "3FJwMEpqSM2cMZVXLbtzzA"}
	// UACredentials["na.se"] = c

	// ua := UARequest{}
	// ua.Init()
	// ua.Endpoint = "reports/responses/list"
	// ua.Get()
}

type UAResponse struct {
	Pushes []UAPush `json:"pushes"`
}

/* PUSH */
type UAPush struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UAID            string        `json:"push_uuid" bson:"uaid"`
	Origin          string        `json:"origin" bson:"origin"`
	Sent            UADate        `json:"push_time" bson:"sent"`
	Type            string        `json:"push_type" bson:"type"`
	DirectResponses int           `json:"direct_responses" bson:"direct_responses"`
	Sends           int           `json:"sends" bson:"sends"`
	Link            string        `json:"link" bson:"link"`
	Text            string        `json:"text" bson:"text"`
}

type UADate struct {
	time.Time
}

func (u *UADate) GetTime() time.Time {
	return u.Time
}

func (u *UADate) UnmarshalJSON(buf []byte) error {
	const ua_dateform = "2006-01-02 15:04:05"
	tt, err := time.Parse(ua_dateform, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	u.Time = tt
	return nil
}

func (u *UAPush) GetDetails() {
	request := UARequest{}
	request.Init(u.Origin)
	request.Endpoint = "reports/responses/" + u.UAID
	err := request.GetSingle()
	if err != nil {
		log.Println(err)
	}
}

/* REQUEST */
type UARequest struct {
	Base        string
	Client      *http.Client
	Request     *http.Request
	Response    *http.Response
	Body        []byte
	Endpoint    string
	Querystring string
	Origin      string
	Pushes      []UAPush
}

func (u *UARequest) Save(db *mgo.Database) bool {
	collection := db.C("pushes")

	// Loop pushes
	for _, row := range u.Pushes {
		// Set origin of push
		row.Origin = u.Origin
		log.Println(row)
		// Get more info on the push
		row.GetDetails()
		continue

		savedPush := UAPush{}
		query := bson.M{"uaid": row.UAID}

		err := collection.Find(query).One(&savedPush)
		if err != nil {
			log.Println("Could not find", row.UAID)
			// Insert

			err = collection.Insert(row)
			if err != nil {
				log.Println("UAPush Save: Could not insert")
			}
		} else {
			err = collection.UpdateId(savedPush.Id, row)
			if err != nil {
				log.Println("UAPush Save: Could not update")
			}
		}
	}

	return true
}

func (u *UARequest) Init(origin string) {
	u.Base = "https://go.urbanairship.com/api/"
	u.Origin = origin
}

func (u *UARequest) Bake() error {
	var err error

	// Credentials
	creds := UACredentials[u.Origin]

	// Check stuff we need
	if len(u.Origin) == 0 {
		return errors.New("UARequest Get: No origin set")
	}
	if len(u.Endpoint) == 0 {
		return errors.New("UARequest Get: No endpoint set")
	}
	if len(creds.User) == 0 {
		return errors.New("UARequest Get: No credentials found")
	}

	// Create a http client
	u.Client = &http.Client{}

	uri := u.Base + u.Endpoint
	if len(u.Querystring) > 0 {
		uri = uri + "?" + u.Querystring
	}

	u.Request, err = http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("No request made:", err)
		return err
	}
	u.Request.SetBasicAuth(creds.User, creds.Pass)
	u.Request.Header.Set("Content-type", "application/json")
	u.Request.Header.Set("Accept", "application/vnd.urbanairship+json; version=3;")

	u.Response, err = u.Client.Do(u.Request)
	if err != nil {
		log.Println("No response:", err)
		return err
	}

	defer u.Response.Body.Close()

	if u.Response.StatusCode != 200 {
		return errors.New(u.Origin + " UARequest Bake returned: " + u.Response.Request.URL.String() + " " + u.Response.Status)
	} else {
		u.Body, err = ioutil.ReadAll(u.Response.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UARequest) GetSingle() error {
	var err error

	err = u.Bake()
	if err != nil {
		return err
	}

	log.Println(string(u.Body))

	uaresponse := UAPush{}

	err = json.Unmarshal(u.Body, &uaresponse)
	if err != nil {
		log.Println("UA GetSingle:", err)
		return err
	}

	u.Pushes = append(u.Pushes, uaresponse)

	return nil
}

func (u *UARequest) GetList() error {
	var err error

	// Dates
	const dateform = "2006-01-02T15:04:05Z"
	const ua_dateform = "2006-01-02 15:04:05"

	loc, _ := time.LoadLocation("Europe/Stockholm")
	now := time.Now()
	enddate := now.Format(dateform)

	startnow := time.Date(now.Year(), now.Month()-1, now.Day(), now.Hour(), now.Minute(), 0, 0, loc)
	startdate := startnow.Format(dateform)

	// URL
	u.Querystring = "start=" + startdate + "&end=" + enddate + "&limit=25"

	err = u.Bake()
	if err != nil {
		return err
	}

	uaresponse := UAResponse{}

	err = json.Unmarshal(u.Body, &uaresponse)
	if err != nil {
		log.Println("json marshal error:", err)
		return err
	}

	u.Pushes = uaresponse.Pushes

	return nil
}
