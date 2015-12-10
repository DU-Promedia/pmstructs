package pmstructs

import (
	"encoding/xml"
	"net/url"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Sections, or lists of content that's parsed
 */
type ArticleListWrapper struct {
	XMLName xml.Name    `xml:"MobileContent"`
	TheList ArticleList `xml:"MobileContentBlocks"`
}

type ArticleList struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Url         string        `json:"url" bson:"url"`
	Articles    []Article     `xml:"MobileContentBlock>StandardArticle" bson:"-" json:"-"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
}

type ArticleContentPlacement struct {
	XMLName     xml.Name      `xml:"ContentPlacement" bson:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Url         string        `json:"url" bson:"url"`
	Articles    []Article     `xml:"StandardArticle" bson:"-"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
}

type ArticleStatisticsList struct {
	XMLName     xml.Name      `xml:"StatisticsList" bson:"-" json:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Url         string        `json:"url" bson:"url"`
	Articles    []Article     `xml:"List>ListItem>StandardArticle" bson:"-" json:"-"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
}

/*
 * ContentPlacements
 */
func (list *ArticleContentPlacement) Save(db *mgo.Database) {
	sectionCollection := db.C("sections")

	// Index should be unique by originid so we should be safe
	parseUrl, _ := url.Parse(list.Url)

	list.Origin = parseUrl.Host
	oldArticles := list.Articles

	if err := sectionCollection.Insert(list); err != nil {
		// Update
		sectionCollection.Update(bson.M{"url": list.Url}, list)
	}
	sectionCollection.Find(bson.M{"url": list.Url}).One(&list)

	list.Articles = oldArticles
}

func (list *ArticleContentPlacement) SaveToDB(db *mgo.Database) {

	// Save to db
	list.Save(db)

	i := 0

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
 * StatisticsList
 */
func (list *ArticleStatisticsList) Save(db *mgo.Database) {
	sectionCollection := db.C("sections")

	// Index should be unique by originid so we should be safe
	parseUrl, _ := url.Parse(list.Url)
	list.Origin = parseUrl.Host
	oldArticles := list.Articles

	if err := sectionCollection.Insert(list); err != nil {
		// Update
		sectionCollection.Update(bson.M{"url": list.Url}, list)
	}
	sectionCollection.Find(bson.M{"url": list.Url}).One(&list)

	list.Articles = oldArticles
}

func (list *ArticleStatisticsList) SaveToDB(db *mgo.Database) {

	// Save section
	list.Save(db)

	i := 0

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

/*
 * Article list
 */
func (list *ArticleList) Save(db *mgo.Database) {
	sectionCollection := db.C("sections")

	// Index should be unique by originid so we should be safe
	parseUrl, _ := url.Parse(list.Url)
	list.Origin = parseUrl.Host
	oldArticles := list.Articles

	if err := sectionCollection.Insert(list); err != nil {
		// Update
		sectionCollection.Update(bson.M{"url": list.Url}, list)
	}
	sectionCollection.Find(bson.M{"url": list.Url}).One(&list)

	list.Articles = oldArticles
}

func (list *ArticleList) SaveToDB(db *mgo.Database) {

	list.Save(db)

	i := 0

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
