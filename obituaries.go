package pmstructs

import (
	// "encoding/json"
	// "log"
	"time"

	//"github.com/simplereach/timeutils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Obituary list for importing
 */
type ObituaryImportList struct {
	Hits []ObituaryImportItem
}

type ObituaryImportItem struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Origin    string        `bson:"origin"`
	OriginId  int           `bson:"originid" json:"Id"`
	Newspaper struct {
		Id        int    `bson:"id" json:"Id"`
		Name      string `bson:"name" json:"Name"`
		Shortname string `bson:"shortname" json:"ShortName"`
		Url       string `bson:"url" json:"Url"`
	} `bson:"-" json:"Newspaper"`
	Category struct {
		Name string `bson:"name" json:"Name"`
	} `bson:"category" json:"Category"`
	ValidFromRaw time.Time `bson:"-" json:"ValidFrom"`
	ValidToRaw   time.Time `bson:"-" json:"ValidTo"`
	// ValidFrom    time.Time      `bson:"validfrom"`
	// ValidTo      time.Time      `bson:"validto"`
	Headline string `bson:"headline" json:"Headline"`
	Images   struct {
		ThumbUrl  string `bson:"thumb"`
		MediumUrl string `bson:"medium"`
		LargeUrl  string `bson:"large"`
	} `bson:"images"`
	ExternalAd    bool   `bson:"externalad"`
	HasVideo      bool   `bson:"hasvideo"`
	CanonicalUrl  string `bson:"canonicalurl"`
	ClientViewUrl string `bson:"clientviewurl"`
	CandelsCount  int    `bson:"candelscount"`
	MemoriesCount int    `bson:"memoriescount"`
}

type ObituaryItem struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	Origin   string        `bson:"origin" json:"origin"`
	OriginId int           `bson:"originid" json:"originid"`
}

func (o *ObituaryImportItem) Save(db *mgo.Database) {
	collection := db.C("obituaries")

	// o.ValidFrom = o.ValidFromRaw.Time
	// o.ValidTo = o.ValidToRaw.Time

	find := bson.M{"originid": o.OriginId}
	collection.Upsert(find, o)
}

// type ObituaryItem struct {
// 	Id        bson.ObjectId `bson:"_id,omitempty" json:"mid"`
// 	OriginId  int           `json:"Id"`
// 	Newspaper struct {
// 		Id        int    `json:"Id"`
// 		Name      string `json:"Name"`
// 		Shortname string `json:"ShortName"`
// 		Url       string `json:"Url"`
// 	} `json:"Newspaper"`
// 	Category struct {
// 		Name string `json:"Name"`
// 	} `json:"Category"`
// 	ValidFrom string `json:"ValidFrom"`
// 	ValidTo   string `json:"ValidTo"`
// 	Headline  string `json:"Headline"`
// 	Images    struct {
// 		ThumbUrl  string
// 		MediumUrl string
// 		LargeUrl  string
// 	}
// 	ExternalAd    bool
// 	HasVideo      bool
// 	CanonicalUrl  string
// 	ClientViewUrl string
// 	CandelsCount  int
// 	MemoriesCount int
// }

// func (o *ObituaryItem) UnmarshalJSON(b []byte) error {
// 	abc := ObituaryImportItem{}

// 	err := json.Unmarshal(b, &abc)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return err
// }
