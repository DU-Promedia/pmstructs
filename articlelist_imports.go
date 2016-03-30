package pmstructs

import (
	"encoding/xml"
	//"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Sections, or lists of content that's parsed
 */

type ArticleListRootElement struct {
	XMLName xml.Name
}

type ArticleListWrapper struct {
	XMLName xml.Name    `xml:"MobileContent"`
	TheList ArticleList `xml:"MobileContentBlocks"`
}

type ArticleList struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `bson:"-" json:"-"`
	AllArticles []struct {
		Article Article       `xml:"StandardArticle" bson:"-" json:"-"`
		Simple  TeaserArticle `xml:"TeaserArticle" bson:"-" json:"-"`
	} `xml:"MobileContentBlock" bson:"-" json:"-"`
}

type ArticleContentPlacement struct {
	XMLName     xml.Name      `xml:"ContentPlacement" bson:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"StandardArticle" bson:"-" json:"-"`
	// Articles []struct {
	// 	Article Article       `xml:"StandardArticle" bson:"-" json:"-"`
	// 	Simple  TeaserArticle `xml:"TeaserArticle" bson:"-" json:"-"`
	// } `xml:"StandardArticle"`
}

type ArticleStatisticsList struct {
	XMLName     xml.Name      `xml:"StatisticsList" bson:"-" json:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `bson:"-" json:"-"`
	AllArticles []struct {
		Article Article       `xml:"StandardArticle" bson:"-" json:"-"`
		Simple  TeaserArticle `xml:"TeaserArticle" bson:"-" json:"-"`
	} `xml:"List>ListItem" bson:"-" json:"-"`
}

/*
 * ContentPlacements
 */
func (list *ArticleContentPlacement) Save(db *mgo.Database) {
	common := ArticleListCommon{}
	common.Url = list.Url
	common.Type = list.Type
	common.Origin = list.Origin
	common.Type = list.Type
	common.ArticleList = list.ArticleList
	common.Articles = list.Articles

	common.Save(db)
}

func (list *ArticleContentPlacement) SaveToDB(db *mgo.Database) {
	// Save to db
	list.Save(db)

	list.ArticleList = []ArticleRef{}

	for _, a := range list.Articles {
		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)

		a.SaveToDB(db)

		if len(a.Serie.Articles) > 0 {
			a.Serie.TrigUpdateOfSiblings(db)
		}
	}

	list.Save(db)
}

/*
 * Article list
 */
func (list *ArticleList) Save(db *mgo.Database) {
	common := ArticleListCommon{}
	common.Url = list.Url
	common.Type = list.Type
	common.Origin = list.Origin
	common.Type = list.Type
	common.ArticleList = list.ArticleList
	common.Articles = list.Articles

	common.Save(db)
}

func (list *ArticleList) SaveToDB(db *mgo.Database) {
	list.Save(db)

	i := 0

	list.ArticleList = []ArticleRef{}

	for _, list_a := range list.AllArticles {
		i++

		if len(list_a.Article.OriginID) > 0 {
			a := list_a.Article
			a.SaveToDB(db)

			artRef := ArticleRef{}
			artRef.ArticleID = a.Id
			list.ArticleList = append(list.ArticleList, artRef)
			list.Articles = append(list.Articles, list_a.Article)

			a.SaveToDB(db)

			if len(a.Serie.Articles) > 0 {
				a.Serie.TrigUpdateOfSiblings(db)
			}
		} else if len(list_a.Simple.Title) > 0 {
			a := Article{}
			b := list_a.Simple

			a.OriginID = b.OriginID
			a.Teaser = ArticleTeaser{
				b.Image,
				"",
				b.Title,
				b.Body,
				b.Link,
			}

			a.SaveToDB(db)

			artRef := ArticleRef{}
			artRef.ArticleID = a.Id
			list.ArticleList = append(list.ArticleList, artRef)
			//list.Articles = append(list.Articles, list_a.Simple)

			a.SaveToDB(db)
		}
	}

	// Save to cache

	list.Save(db)
}

/*
 * StatisticsList
 */
func (list *ArticleStatisticsList) Save(db *mgo.Database) {
	common := ArticleListCommon{}
	common.Url = list.Url
	common.Type = list.Type
	common.Origin = list.Origin
	common.Type = list.Type
	common.ArticleList = list.ArticleList
	common.Articles = list.Articles

	common.Save(db)
}

func (list *ArticleStatisticsList) SaveToDB(db *mgo.Database) {
	// Save section
	list.Save(db)

	i := 0

	list.ArticleList = []ArticleRef{}

	for _, list_a := range list.AllArticles {
		i++

		if len(list_a.Article.OriginID) > 0 {
			a := list_a.Article
			a.SaveToDB(db)

			artRef := ArticleRef{}
			artRef.ArticleID = a.Id
			list.ArticleList = append(list.ArticleList, artRef)
			list.Articles = append(list.Articles, list_a.Article)

			a.SaveToDB(db)

			if len(a.Serie.Articles) > 0 {
				a.Serie.TrigUpdateOfSiblings(db)
			}
		} else if len(list_a.Simple.Title) > 0 {
			a := Article{}
			b := list_a.Simple

			a.OriginID = b.OriginID
			a.Teaser = ArticleTeaser{
				b.Image,
				"",
				b.Title,
				b.Body,
				b.Link,
			}

			a.SaveToDB(db)

			artRef := ArticleRef{}
			artRef.ArticleID = a.Id
			list.ArticleList = append(list.ArticleList, artRef)
			//list.Articles = append(list.Articles, list_a.Simple)

			a.SaveToDB(db)
		}
	}

	list.Save(db)
}
