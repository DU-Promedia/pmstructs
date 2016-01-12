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

type ExtraBlock struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Origin          string        `bson:"origin" json:"origin"`
	OriginPlacement int           `bson:"originplacement" json:"placement"`
	Headline        string        `bson:"headline"`
	Pubdate         time.Time     `bson:"pubdate" json:"pubdate"`
	ArticleList     []ArticleRef  `bson:"articles" json:"-"`
	Articles        []Article     `bson:"-" json:"articles"`
}

func (e *ExtraBlock) Save(db *mgo.Database) bool {
	collection := db.C("extrablocks")

	query := bson.M{"origin": e.Origin, "originplacement": e.OriginPlacement}
	res := ExtraBlock{}

	err := collection.Find(query).One(&res)
	if err != nil {
		log.Println("ExtraBlock Save:", err)
	}

	if res.Id.Valid() {
		log.Println("Saving extra block")
		collection.UpdateId(res.Id, e)
	} else {
		log.Println("Inserting extra block")
		collection.Insert(e)
	}

	// Load new data in to ExtraBlock
	collection.Find(query).One(&e)

	return true
}

func (e *ExtraBlock) Remove(db *mgo.Database) bool {
	collection := db.C("extrablocks")

	err := collection.RemoveId(e.Id)
	if err != nil {
		log.Println("ExtraBlock Remove:", err)
		return false
	}

	return true
}

type RightNow struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	Origin    string        `bson:"origin" json:"origin"`
	ArticleID bson.ObjectId `bson:"articleid" json:"id"`
	Headline  string        `bson:"headline" json:"headline"`
	Link      string        `bson:"link" json:"-"`
	Text      string        `bson:"text" json:"text"`
	Pubdate   time.Time     `bson:"pubdate" json:"pubdate"`
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
