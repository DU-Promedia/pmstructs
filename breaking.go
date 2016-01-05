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

type RightNow struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	Origin    string        `bson:"origin"`
	ArticleID bson.ObjectId `bson:"articleid" json:"id"`
	Headline  string        `bson:"headline"`
	Link      string        `bson:"link"`
	Text      string        `bson:"text"`
	Pubdate   time.Time     `bson:"pubdate"`
}

func (r *RightNow) Save(db *mgo.Database) bool {
	collection := db.C("rightnow")

	query := bson.M{"origin": r.Origin, "link": r.Link}
	res := RightNow{}

	err := collection.Find(query).One(&res)
	if err != nil {
		log.Println("RightNow Save:", err)
	}

	if res.Id.Valid() {
		err = collection.UpdateId(res.Id, r)
		if err != nil {
			log.Println("RightNow Save:", err)
			return false
		}
	} else {
		err = collection.Insert(r)
		if err != nil {
			log.Println("RightNow Save:", err)
			return false
		}
	}

	return true
}

func (r *RightNow) Remove(db *mgo.Database) bool {
	collection := db.C("rightnow")
	// log.Println("Will remove", r.Id)
	// return false

	err := collection.RemoveId(r.Id)
	if err != nil {
		log.Println("RightNow Remove:", err)
		return false
	}

	return true
}
