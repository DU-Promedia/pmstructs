package pmstructs

import (
	"log"
	// "net/url"
	// "strings"
	"time"

	//"github.com/PuerkitoBio/goquery"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PushParser struct {
	Payload  Push `json:"payload"`
	Response struct {
		Ok           bool     `json:"ok"`
		OperationsID string   `json:"operation_id"`
		PushID       []string `json:"push_ids"`
	} `json:"response"`
}

type Push struct {
	Id              bson.ObjectId `bson:"_id,omitempty"`
	Origin          string        `bson:"origin" json:"origin"`
	OriginId        string        `bson:"originid" json:"originid"`
	Tags            []string      `bson:"tags" json:"tags"`
	Date            time.Time     `bson:"date" json:"date"`
	Alert           string        `bson:"alert" json:"alert"`
	Url             string        `bson:"url" json:"url"`
	ArticleRef      ArticleRef    `bson:"articleref" json:"-"`
	DeviceTypes     []string      `json:"device_types"`
	Sends           int           `bson:"sends" json:"sends"`
	DirectResponses int           `bson:"direct_responses" json:"direct_responses"`

	Audience     interface{} `bson:"-" json:"audience"`
	Notification struct {
		Alert string `bson:"-" json:"alert"`
		Ios   struct {
			Extra struct {
				Url string `bson:"-" json:"url"`
			} `bson:"-" json:"extra"`
		} `bson:"-" json:"ios"`
		Android struct {
			Extra struct {
				Url string `bson:"-" json:"url"`
			} `bson:"-" json:"extra"`
		} `vjson:"android"`
	} `bson:"-" json:"notification"`
}

func (p *Push) Save(db *mgo.Database) {
	collection := db.C("notifications")
	var err error

	if len(p.Notification.Alert) > 0 {
		p.Alert = p.Notification.Alert
	}

	if len(p.Notification.Ios.Extra.Url) > 0 && (p.Notification.Android.Extra.Url == p.Notification.Ios.Extra.Url) {
		p.Url = p.Notification.Ios.Extra.Url
	} else if len(p.Notification.Android.Extra.Url) > 0 && len(p.Notification.Ios.Extra.Url) == 0 {
		p.Url = p.Notification.Android.Extra.Url
	} else if len(p.Notification.Ios.Extra.Url) > 0 && len(p.Notification.Android.Extra.Url) == 0 {
		p.Url = p.Notification.Ios.Extra.Url
	}

	// Find in DB
	result := Push{}
	find := bson.M{"originid": p.OriginId}
	err = collection.Find(find).One(&result)

	if result.Id.Valid() {
		p.Id = result.Id
	}

	if len(p.Id) > 0 {
		// Update
		collection.UpsertId(p.Id, p)
		return
	} else {
		articleid := GetOriginIdFromUrl(p.Url)
		if len(articleid) > 0 {
			article := Article{}
			article.LoadArticleByOriginId(articleid, db)

			if article.Id.Valid() {
				articleref := ArticleRef{}
				articleref.ArticleID = article.Id

				p.ArticleRef = articleref
			}
		}

		err = collection.Insert(p)
		if err != nil {
			log.Println("Push Save:", err)
			return
		}
	}
}

func (p *Push) ParseAudience() {
	var isTags bool

	switch p.Audience.(type) {
	case string:
		isTags = false
	default:
		isTags = true
	}

	if isTags {
		tagg := p.Audience.(map[string]interface{})

		for _, cont := range tagg {
			var tags []interface{}
			tags = cont.([]interface{})

			for _, a := range tags {
				st := a.(string)
				p.Tags = append(p.Tags, st)
			}
		}
	} else {
		all := p.Audience.(string)
		p.Tags = append(p.Tags, all)
	}
}
