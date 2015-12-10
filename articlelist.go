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
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	OriginApp   string        `bson:"originapp" json:"-"`
	Url         string        `json:"url" bson:"url"`
	Articles    []Article     `json:"articles" bson:"-"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"-"`
}

type ArticleListCommonCache struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"cacheid"`
	SectionID bson.ObjectId `bson:"sectiond" json:"mid"`
	OriginID  string        `bson:"originid" json:"id"`
	Origin    string        `bson:"origin" json:"origin"`
	OriginApp string        `bson:"originapp" json:"originapp"`
	Url       string        `json:"url" bson:"url"`
	Articles  []Article     `json:"articles" bson:"articles"`
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
		err = coll.Update(findQuery, a)
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

func (a *ArticleListCommon) SaveCached(db *mgo.Database) {
	a.LoadArticles(db)

	if len(a.Articles) == 0 {
		log.Println("ArticleListCommon SaveCached: No articles in list to cache")
		return
	}

	cacheCol := db.C("sections_cache")

	cache := ArticleListCommonCache{}
	cache.SectionID = a.ID
	cache.OriginID = a.OriginID
	cache.Origin = a.Origin
	cache.OriginApp = a.OriginApp
	cache.Url = a.Url
	cache.Articles = a.Articles

	findQuery := bson.M{"sectionid": cache.SectionID}
	savedList := ArticleListCommon{}
	err := cacheCol.Find(findQuery).One(&savedList)
	if err != nil {
		// Do an insert
		err = cacheCol.Insert(cache)
		if err != nil {
			log.Println("ArticleListCommon SaveCached: No insert:", err)
		}
		return
	}

	err = cacheCol.Update(findQuery, cache)
	if err != nil {
		log.Println("ArticleListCommon SaveCached: Could not insert or update:", err)
	}

}
