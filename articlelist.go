package pmstructs

import (
	//"encoding/json"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * AKA: Sections
 */
type ArticleListCommon struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	OriginID    string        `bson:"originid,omitempty" json:"mid,omitempty"`
	Origin      string        `bson:"origin" json:"origin"`
	OriginApp   string        `bson:"originapp" json:"-"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	Articles    []Article     `json:"articles,omitempty" bson:"-"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"-"`
}

func (a *ArticleListCommon) Save(db *mgo.Database) {
	if len(a.Url) == 0 {
		log.Println("ArticleListCommon Save: No url given. Needs it to find")
		return
	}

	coll := db.C("sections")

	findQuery := bson.M{"url": a.Url}
	savedList := ArticleListCommon{}

	err := coll.Find(findQuery).One(&savedList)
	if err != nil {
		// Insert it
		err = coll.Insert(a)
		if err != nil {
			log.Println("ArticleListCommon Save: No insert:", err)
		}
	} else {
		if len(a.ArticleList) == 0 && len(savedList.ArticleList) > 0 {
			a.ArticleList = savedList.ArticleList
		}

		if len(a.Type) == 0 && len(savedList.Type) > 0 {
			a.Type = savedList.Type
		}

		if len(a.Origin) == 0 && len(savedList.Origin) > 0 {
			a.Origin = savedList.Origin
		}

		if len(a.OriginApp) == 0 && len(savedList.OriginApp) > 0 {
			a.OriginApp = savedList.OriginApp
		}

		if len(a.OriginID) == 0 && len(savedList.OriginID) > 0 {
			a.OriginID = savedList.OriginID
		}

		_, err = coll.Upsert(findQuery, a)
		//err = coll.Update(findQuery, a)
		if err != nil {
			log.Println("ArticleListCommon Save: No update:", err)
		}
	}

	return
}

func (a *ArticleListCommon) LoadArticles(db *mgo.Database) {
	articleCollection := db.C("articles")

	for _, ref := range a.ArticleList {
		art := Article{}

		err := articleCollection.FindId(ref.ArticleID).One(&art)
		if err != nil {
			log.Println("... ArticleListCommon LoadArticles: Could not load article:", err)
			return
		}

		a.Articles = append(a.Articles, art)
	}
}

func (a *ArticleListCommon) GetLatestArticles(db *mgo.Database, limit int) error {
	artCol := db.C("articles")

	findQuery := bson.M{"originsource": a.Origin}

	err := artCol.Find(findQuery).Sort("-pubdate").Limit(limit).All(&a.Articles)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleListCommon) LoadFromDB(db *mgo.Database) {
	if len(a.Url) == 0 {
		log.Println("ArticleListCommon LoadFromDB: Add an url")
		return
	}

	secCol := db.C("sections")

	findQuery := bson.M{"url": a.Url}
	err := secCol.Find(findQuery).One(&a)
	if err != nil {
		log.Println("ArticleListCommon LoadFromDB: Could not load from db:", err)
		return
	}

	return
}

func (a *ArticleListCommon) LoadFromDBByID(db *mgo.Database) {
	if a.ID.Valid() == false {
		log.Println("ArticleListCommon LoadFromDBById: Not an objectid")
		return
	}

	secCol := db.C("sections")

	err := secCol.FindId(a.ID).One(&a)
	if err != nil {
		log.Println("ArticleListCommon LoadFromDBById:", err)
		return
	}
}
