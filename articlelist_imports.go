package pmstructs

import (
	"encoding/xml"
	"log"
	//	"net/url"

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
	Articles    []Article     `xml:"MobileContentBlock>StandardArticle" bson:"-" json:"-"`
}

type ArticleContentPlacement struct {
	XMLName     xml.Name      `xml:"ContentPlacement" bson:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"StandardArticle" bson:"-"`
}

type ArticleStatisticsList struct {
	XMLName     xml.Name      `xml:"StatisticsList" bson:"-" json:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Type        string        `bson:"type" json:"type"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"List>ListItem>StandardArticle" bson:"-" json:"-"`
}

/*
 * ContentPlacements
 */
func (list *ArticleContentPlacement) Save(db *mgo.Database) {
	log.Println("Saving ArticleContentPlacement:", list.Url)

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

	i := 0

	list.ArticleList = []ArticleRef{}

	for _, a := range list.Articles {
		i++

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)

		a.SaveToDB(db)
	}

	list.Save(db)
}

func (list *ArticleContentPlacement) GetArticles() []Article {
	listOfArticles := make([]Article, 0)

	for _, a := range list.Articles {
		target := make([]Article, len(listOfArticles)+1)
		copy(target, listOfArticles)
		listOfArticles = append(target, a)
	}

	return listOfArticles
}

/*
 * Article list
 */
func (list *ArticleList) Save(db *mgo.Database) {
	log.Println("Saving ArticleList:", list.Url)

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

	for _, a := range list.Articles {
		i++

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)

		a.SaveToDB(db)
	}

	list.Save(db)
}

func (list *ArticleList) GetArticles() []Article {
	listOfArticles := make([]Article, 0)

	for _, a := range list.Articles {
		listOfArticles = append(listOfArticles, a)
	}

	return listOfArticles
}

/*
 * StatisticsList
 */
func (list *ArticleStatisticsList) Save(db *mgo.Database) {
	log.Println("Saving StatisticsList:", list.Url)

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

	for _, a := range list.Articles {
		i++

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)

		a.SaveToDB(db)
	}

	list.Save(db)
}

func (list *ArticleStatisticsList) GetArticles() []Article {
	listOfArticles := make([]Article, 0)

	for _, a := range list.Articles {
		target := make([]Article, len(listOfArticles)+1)
		copy(target, listOfArticles)
		listOfArticles = append(target, a)
	}

	return listOfArticles
}
