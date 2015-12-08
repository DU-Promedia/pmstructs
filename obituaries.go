package pmstructs

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

/*
 * Obituary list for importing
 */
type ObituaryImportList struct {
	List []ObituaryImport `json:"Hits"`
}

func (o *ObituaryImportList) Save() {
	for _, item := range o.List {
		log.Println(item.OriginID)
	}
}

/*
 * Obituary entity for importing
 */
type ObituaryImport struct {
	ID            bson.ObjectId  `bson:"_id,omitempty" json:"mid"`
	OriginID      string         `bson:"originid" json:"id"`
	Newspaper     string         `bson:"newspaper_name" json:"Newspaper>Name"`
	NewspaperUrl  string         `bson:"newspaper_url" json:"Newspaper>Url"`
	Category      string         `bson:"category" json:"Category>Name"`
	County        string         `bson:"county" json:"County>Name"`
	City          string         `bson:"city" json:"County>City>Name"`
	ValidFromRaw  string         `bson:"-" json:"ValidFrom"`
	ValidToRaw    string         `bson:"-" json:"ValidTo"`
	Headline      string         `bson:"headline" json:"Headline"`
	Images        ObituaryImages `bson:"images" json:"Images"`
	CanonicalUrl  string         `bson:"canonical" json:"CanonicalUrl"`
	ClientviewUrl string         `bson:"clientviewurl" json:"ClientViewUrl"`
}

type ObituaryImages struct {
	Thumb  string `bson:"thumb" json:"ThumbUrl"`
	Medium string `bson:"medium" json:"MediumUrl"`
	Large  string `bson:"large" json:"LargeUrl"`
}
