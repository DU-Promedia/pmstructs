package pmstructs

import (
	"encoding/xml"
	"log"
	"net/url"

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
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"MobileContentBlock>StandardArticle" bson:"-" json:"-"`
}

type ArticleContentPlacement struct {
	XMLName     xml.Name      `xml:"ContentPlacement" bson:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"StandardArticle" bson:"-"`
}

type ArticleStatisticsList struct {
	XMLName     xml.Name      `xml:"StatisticsList" bson:"-" json:"-"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID    string        `xml:"id,attr" bson:"originid" json:"id"`
	Origin      string        `bson:"origin" json:"origin"`
	Url         string        `json:"url" bson:"url"`
	ArticleList []ArticleRef  `bson:"articlelist" json:"articlelist"`
	Articles    []Article     `xml:"List>ListItem>StandardArticle" bson:"-" json:"-"`
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
	oldReflist := list.ArticleList

	// Find in DB
	err := sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
	if err != nil {
		// Found no section for articlelist
		log.Println("Inserting section")
		sectionCollection.Insert(list)
		err = sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
		if err != nil {
			log.Println("ArticleContentPlacement:", err)
			return
		}
	}

	if len(oldReflist) > 0 {
		list.ArticleList = oldReflist
	}

	if len(oldArticles) > 0 {
		list.Articles = oldArticles
	}

	sectionCollection.Update(bson.M{"url": list.Url}, list)

	//sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
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
 * Article list
 */
func (list *ArticleList) Save(db *mgo.Database) {
	sectionCollection := db.C("sections")

	// Index should be unique by originid so we should be safe
	parseUrl, _ := url.Parse(list.Url)
	list.Origin = parseUrl.Host

	oldArticles := list.Articles
	oldReflist := list.ArticleList

	// Find in DB
	err := sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
	if err != nil {
		// Found no section for articlelist
		log.Println("Inserting section")
		sectionCollection.Insert(list)
		err = sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
		if err != nil {
			log.Println("ArticleContentPlacement:", err)
			return
		}
	}

	if len(oldReflist) > 0 {
		list.ArticleList = oldReflist
	}

	if len(oldArticles) > 0 {
		list.Articles = oldArticles
	}

	sectionCollection.Update(bson.M{"url": list.Url}, list)

	//sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
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

/*
 * StatisticsList
 */
func (list *ArticleStatisticsList) Save(db *mgo.Database) {
	sectionCollection := db.C("sections")

	// Index should be unique by originid so we should be safe
	parseUrl, _ := url.Parse(list.Url)
	list.Origin = parseUrl.Host

	oldArticles := list.Articles
	oldReflist := list.ArticleList
	savedList := ArticleListCommon{}

	// Find in DB
	err := sectionCollection.Find(bson.M{"url": list.Url}).One(&savedList)
	if err != nil {
		sectionCollection.Insert(list)
		err = sectionCollection.Find(bson.M{"url": list.Url}).One(&savedList)
		if err != nil {
			log.Println("ArticleContentPlacement:", err)
			return
		}
	}

	if len(oldReflist) > 0 {
		list.ArticleList = oldReflist
	}

	if len(oldArticles) > 0 {
		list.Articles = oldArticles
	}

	sectionCollection.Update(bson.M{"url": list.Url}, list)

	//sectionCollection.Find(bson.M{"url": list.Url}).One(&list)
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
