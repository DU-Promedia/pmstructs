package pmstructs

import (
	"encoding/xml"
	"log"
	"net/url"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Teaser article
 */
type TeaserArticle struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"mid"`
	OriginID string        `xml:"id,attr" json:"id"`
	Image    string        `xml:"TeaserArticleImage>TeaserArticleImagePath" json:"image"`
	Title    string        `xml:"TeaserArticleTitle" json:"title"`
	Subtitle string        `xml:"TeaserArticleSubTitle" json:"subtitle"`
	Preamble string        `xml:"TeaserArticlePreamble" json:"preamble"`
	Body     string        `xml:"TeaserArticleBody" json:"body"`
	Internal string        `xml:"Internal" json:"internal"`
	Link     string        `xml:"TeaserArticleExternal>TeaserArticleExternalLink" json:"link"`
	Linktext string        `xml:"TeaserArticleExternal>TeaserArticleExternalLinkName" json:"linktext"`
}

/*
 * Complete article
 */
type Article struct {
	Id              bson.ObjectId      `bson:"_id,omitempty" json:"mid"`
	OriginID        string             `xml:"id,attr" json:"id"`
	OriginalLink    string             `xml:"StandardArticleOriginalLink" json:"originallink"`
	OriginSource    string             `bson:"originsource" json:"originsource"`
	Title           string             `xml:"StandardArticleTitle" json:"title"`
	Subtitle        string             `xml:"StandardArticleSubTitle" json:"subtitle"`
	Supertitle      string             `xml:"StandardArticleSuperTitle" json:"supertitle"`
	Preamble        string             `xml:"StandardArticlePreamble" json:"preamble"`
	Body            string             `xml:"StandardArticleBody" json:"content"`
	Image           string             `xml:"StandardArticleImage>StandardArticleImagePath" json:"image" bson:"image"`
	ImageByline     string             `xml:"StandardArticleImage>StandardArticlePhotographer" json:"imagebyline" bson:"imagebyline"`
	ArticleImages   []ArticleImage     `xml:"ArticleImages>ArticleImage" json:"articleimages" bson:"articleimages"`
	ImageAlbum      ArticleImageAlbum  `xml:"StandardArticleTopImageAlbum>ImageAlbum" json:"imagealbum" bson:"imagealbum"`
	Category        string             `xml:"StandardArticleCategory" json:"category"`
	ArticleType     string             `xml:"StandardArticleType" json:"articletype"`
	ArticleInfo     string             `xml:"StandardArticleInfo" json:"articleinfo"`
	PubdateRaw      string             `xml:"StandardArticlePubDate"`
	ModdateRaw      string             `xml:"StandardArticlePubModDate"`
	Pubdate         time.Time          `json:"pubdate" bson:"pubdate"`
	Moddate         time.Time          `json:"moddate" bson:"moddate"`
	Location        string             `xml:"Location" json:"location"`
	Latitude        string             `xml:"StandardArticleGeo>StandardArticleLatitude" json:"latitude" bson:"latitude"`
	Longitude       string             `xml:"StandardArticleGeo>StandardArticleLongitude" json:"longitude" bson:"longitude"`
	Department      string             `xml:"ArticleDepartment" json:"department"`
	Teaser          ArticleTeaser      `xml:"StandardArticleTeaser" json:"teaser"`
	ExtraTeaser     ArticleExtraTeaser `xml:"StandardArticleExtraTeaser" json:"extrateaser"`
	Byline          []ArticleByline    `xml:"StandardArticleBylines>StandardArticleByline" json:"bylines"`
	Links           []ArticleLinks     `xml:"StandardArticleLinks>Link" json:"articlelinks"`
	CommentsEnabled string             `xml:"StandardArticleArticleCommentsEnabled" json:"commentsenabled"`
	CommentsTitle   string             `xml:"StandardArticleArticleComments>DiscusstionTitle" json:"commenttitle"`
	Comments        []ArticleComments  `xml:"StandardArticleArticleComments>StandardArticleArticleComment" json:"comments"`
	Facts           []ArticleFact      `xml:"StandardArticleFacts>StandardArticleFact" json:"facts"`
	BackgroundFacts []ArticleFact      `xml:"StandardArticleBackgroundFacts>StandardArticleBackgroundFact" json:"backgroundfacts"`
	Theme           string             `xml:"StandardArticleTheme" json:"-" bson:"theme"`
	LastMod         time.Time          `json:"lastmod" bson:"lastmod"`
	ArticleTags     []string           `xml:"StandardArticleKeyWords>StandardArticleKeyWord" json:"articletags"`
	Tags            []string           `json:"-" bson:"tags,omitempty"`
	Video           ArticleVideo       `xml:"PicSearchVideo" bson:"video" json:"video"`
	TopContent      string             `xml:"HandeMadeTopContent" bson:"topcontent" json:"topcontent"`
}

type ArticleRef struct {
	ArticleID bson.ObjectId `bson:"articleid" json:"articleid"`
}

/* TODO: Fix WebPolls from XML
type ArticlePoll struct {
}

/*
 * Articles section struct
*/
type ArticleSection struct {
	SectionID bson.ObjectId `json:"id" bson:"sectionid"`
	Placement int           `json:"placement" bson:"placement"`
}

/*
 * Article parts
 */
type ArticleTeaser struct {
	Image       string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePath" json:"image"`
	ImageByline string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePhotographer" json:"imagebyline"`
	Title       string `xml:"StandardArticleTeaserTitle" json:"title"`
	Body        string `xml:"StandardArticleTeaserBody" json:"body"`
}

type ArticleExtraTeaser struct {
	Title string `xml:"StandardArticleExtraTeaserTitle" json:"title"`
	Body  string `xml:"StandardArticleExtraTeaserBody" json:"body"`
}

type ArticleByline struct {
	Name       string `xml:"Name" json:"name"`
	Email      string `xml:"Email" json:"email"`
	Phone      string `xml:"Phone" json:"phone"`
	Role       string `xml:"OccupationalRole" json:"role"`
	Image      string `xml:"ImagePath" json:"image"`
	Department string `xml:"Department" json:"department"`
}

type ArticleLinks struct {
	Title    string `xml:"Title" json:"title"`
	Internal bool   `xml:"Internal" json:"is_internal"`
	Url      string `xml:"Url" json:"url"`
	Image    string `xml:"Image" json:"image"`
	Pubdate  string `xml:"PubDate" json:"pubdate"`
	Preamble string `xml:"Preamble" json:"preamble"`
	Category string `xml:"Category" json:"category"`
}

type ArticleComments struct {
	Title     string `xml:"Title" json:"title"`
	Body      string `xml:"Body" json:"body"`
	Author    string `xml:"Author>AliasOrFullName" json:"author"`
	Pubdate   string `xml:"PublicationDate" json:"pubdate"`
	BodyQuote string `xml:"BodyQuote" json:"quoted"`
}

type ArticleVideo struct {
	Title        string                `xml:"PicSearchVideoTitle" json:"title"`
	Description  string                `xml:"PicSearchVideoDescription" json:"description"`
	Category     string                `xml:"PicSearchVideoCategory" json:"category"`
	ThumbnailUrl string                `xml:"PicSearchVideoThumbNail" json:"tumbnailurl"`
	Publishdate  string                `xml:"PicSearchVideoPublishDate" json:"publishdate"`
	MediaId      string                `xml:"PicSearchVideoMediaId" json:"mediaid"`
	Streams      []ArticleVideoStreams `xml:"PicSearchMediaStreams>PicSearchMediaStream" json:"streams"`
}

type ArticleVideoStreams struct {
	StreamFormat string `xml:"PicSearchMediaStreamFormat" json:"format"`
	StreamUri    string `xml:"PicSearchMediaStreamUri" json:"uri"`
}

type ArticleImage struct {
	Url          string `xml:"ImageUrl" json:"url"`
	Title        string `xml:"ImageTitle" json:"title"`
	Text         string `xml:"ImageText" json:"text"`
	Photographer string `xml:"ImagePhotographer" json:"photographer"`
	ShowTitle    string `xml:"ImageShowTitle" json:"showtitle"`
}

type ArticleImageAlbum struct {
	Id                 string                   `xml:"id,attr" json:"id"`
	Title              string                   `xml:"ImageAlbumTitle" json:"title"`
	Description        string                   `xml:"ImageAlbumDescription" json:"description"`
	Images             []ArticleImageAlbumImage `xml:"ImageAlbumImages>ImageAlbumImage" json:"images"`
	AllowComments      string                   `xml:"ImageAlbumAllowComments" json:"allowcomments"`
	JsonUrl            string                   `xml:"ImageAlbumJsonUrl" json:"jsonurl"`
	SharingLink        string                   `xml:"ImageAlbumSharingLink" json:"sharinglink"`
	TeaserTitle        string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImageTitle" json:"teasertitle"`
	TeaserDesc         string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImageDescription" json:"teaserdesc"`
	TeaserPhotographer string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImagePhotographer" json:"teaserphotographer"`
	TeaserImagePath    string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImagePath" json:"teaserimagepath"`
}

type ArticleImageAlbumImage struct {
	Id           string `xml:"id,attr" json:"id"`
	Title        string `xml:"ImageAlbumImageTitle" json:"title"`
	Description  string `xml:"ImageAlbumImageDescription" json:"description"`
	Photographer string `xml:"ImageAlbumImagePhotographer" json:"photographer"`
	ImagePath    string `xml:"ImageAlbumImagePath" json:"imagepath"`
}

type ArticleFact struct {
	Title string `xml:"Title" json:"title"`
	Body  string `xml:"Body" json:"body"`
}

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

		sect := ArticleSection{}
		sect.SectionID = list.ID
		sect.Placement = i

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)
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

		sect := ArticleSection{}
		sect.SectionID = list.ID
		sect.Placement = i

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)
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

		sect := ArticleSection{}
		sect.SectionID = list.ID
		sect.Placement = i

		a.SaveToDB(db)

		artRef := ArticleRef{}
		artRef.ArticleID = a.Id
		list.ArticleList = append(list.ArticleList, artRef)
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
 * Article
 */
func (a *Article) SaveToDB(db *mgo.Database) {
	collection := db.C("articles")

	a.LastMod = time.Now()
	a.Pubdate, _ = time.Parse(time.RFC1123Z, a.PubdateRaw)
	a.Moddate, _ = time.Parse(time.RFC1123Z, a.ModdateRaw)

	arturl, _ := url.Parse(a.OriginalLink)
	a.OriginSource = arturl.Host

	// Find Document
	docToUpdate := bson.M{"originid": a.OriginID}
	savedArticle := Article{}

	err := collection.Find(docToUpdate).One(&savedArticle)
	if err != nil {
		if debugMode {
			log.Println("Found no document to update, inserting")
		}
		collection.Insert(a)
	}

	// Fields that we set somewhere else ...
	a.Tags = savedArticle.Tags

	if err = collection.Update(docToUpdate, a); err != nil {
		log.Println("Article SaveToDB: Could not update:", err)
		return
	}

	collection.Find(docToUpdate).One(&a)
}

func (a *Article) UpdateTags(db *mgo.Database) {
	collection := db.C("articles")

	a.LastMod = time.Now()

	// Find document
	docToUpdate := bson.M{"originid": a.OriginID}

	if err := collection.Update(docToUpdate, a); err != nil {
		log.Println("Article UpdateTags: Could not update:", err)
	}
}
