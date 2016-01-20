package pmstructs

import (
	"log"
	"net/url"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * Complete article
 * Fimpa sections

 */
type Article struct {
	Id              bson.ObjectId      `bson:"_id,omitempty" json:"mid"`                       // OUT
	OriginID        string             `xml:"id,attr" json:"id"`                               // OUT
	OriginalLink    string             `xml:"StandardArticleOriginalLink" json:"originallink"` // OUT
	OriginSource    string             `bson:"originsource" json:"originsource"`               // OUT
	Title           string             `xml:"StandardArticleTitle" json:"title"`               // OUT
	Subtitle        string             `xml:"StandardArticleSubTitle" json:"subtitle,omitempty"`
	Supertitle      string             `xml:"StandardArticleSuperTitle" json:"supertitle,omitempty"`
	Preamble        string             `xml:"StandardArticlePreamble" json:"preamble"` // OUT
	Body            string             `xml:"StandardArticleBody" json:"content"`      // OUT
	BodyParts       []string           `bson:"contentparts,omitempty" json:"contentparts,omitempty"`
	Image           string             `xml:"StandardArticleImage>StandardArticleImagePath" json:"image,omitempty" bson:"image"`
	ImageByline     string             `xml:"StandardArticleImage>StandardArticlePhotographer" json:"imagebyline,omitempty" bson:"imagebyline"`
	ArticleImages   []ArticleImage     `xml:"ArticleImages>ArticleImage" json:"articleimages,omitempty" bson:"articleimages"`
	ImageAlbum      ArticleImageAlbum  `xml:"StandardArticleTopImageAlbum>ImageAlbum" json:"imagealbum,omitempty" bson:"imagealbum"`
	Category        string             `xml:"StandardArticleCategory" json:"category,omitempty"` // OUT
	ArticleType     string             `xml:"StandardArticleType" bson:"articletype" json:"-"`
	ArticleInfo     string             `xml:"StandardArticleInfo" bson:"articleinfo" json:"-"`
	PubdateRaw      string             `xml:"StandardArticlePubDate" bson:"-" json:"-"`
	ModdateRaw      string             `xml:"StandardArticlePubModDate" bson:"-" json:"-"`
	Pubdate         time.Time          `json:"pubdate" bson:"pubdate"`
	Moddate         time.Time          `json:"moddate" bson:"moddate"`
	Location        string             `xml:"Location" json:"location,omitempty"`
	Latitude        string             `xml:"StandardArticleGeo>StandardArticleLatitude" json:"latitude,omitempty" bson:"latitude"`    // OUT
	Longitude       string             `xml:"StandardArticleGeo>StandardArticleLongitude" json:"longitude,omitempty" bson:"longitude"` // OUT
	Department      string             `xml:"ArticleDepartment" json:"department,omitempty"`                                           // OUT
	Teaser          ArticleTeaser      `xml:"StandardArticleTeaser" json:"teaser"`                                                     // OUT
	ExtraTeaser     ArticleExtraTeaser `xml:"StandardArticleExtraTeaser" json:"extrateaser"`
	Byline          []ArticleByline    `xml:"StandardArticleBylines>StandardArticleByline" json:"bylines"`
	Links           []ArticleLinks     `xml:"StandardArticleLinks>Link" json:"articlelinks,omitempty"`
	CommentCount    int                `bson:"commentcount,omitempty" json:"commentcount"`
	CommentsEnabled bool               `xml:"StandardArticleArticleCommentsEnabled" json:"commentsenabled"`
	CommentsTitle   string             `xml:"StandardArticleArticleComments>DiscusstionTitle" json:"commenttitle,omitempty"`
	Comments        []ArticleComments  `xml:"StandardArticleArticleComments>StandardArticleArticleComment" json:"comments,omitempty"`
	Facts           []ArticleFact      `xml:"StandardArticleFacts>StandardArticleFact" json:"facts,omitempty"`
	BackgroundFacts []ArticleFact      `xml:"StandardArticleBackgroundFacts>StandardArticleBackgroundFact" json:"backgroundfacts,omitempty"`
	Theme           string             `xml:"StandardArticleTheme" json:"-" bson:"theme"`
	LastMod         time.Time          `json:"lastmod" bson:"lastmod"` // OUT
	ArticleTags     []string           `xml:"StandardArticleKeyWords>StandardArticleKeyWord" json:"articletags,omitempty"`
	Tags            []string           `json:"-" bson:"tags,omitempty"`
	Video           ArticleVideo       `xml:"PicSearchVideo" bson:"video" json:"video,omitempty"` // OUT
	TopContent      string             `xml:"HandeMadeTopContent" bson:"topcontent" json:"topcontent,omitempty"`
	Sections        []ArticleSection   `bson:"sections" json:"sections,omitempty"`
	Shares          ArticleShares      `bson:"shares" json:"shares,omitempty"`                                               // OUT
	Serie           ArticleSerie       `xml:"StandardArticleArticleSeries>ArticleSerie" bson:"serie" json:"serie,omitempty"` // OUT
}

/*
 * Article Serie
 */
type ArticleSerie struct {
	Title    string                `xml:"Title" bson:"title" json:"title,omitempty"`
	Articles []ArticleSerieArticle `xml:"Articles>Article" bson:"articles" json:"articles,omitempty"`
}

type ArticleSerieArticle struct {
	Title      string    `xml:"Title" bson:"title" json:"title,omitempty"`
	Preamble   string    `xml:"Preamble" bson:"preamble" json:"preamble,omitempty"`
	Image      string    `xml:"Image" bson:"image" json:"image,omitempty"`
	Pubdate    time.Time `bson:"pubdate" json:"pubdate,omitempty"`
	RawPubdate string    `xml:"PubDate" bson:"-" json:"-"`
	Location   string    `xml:"Location" bson:"location" json:"location,omitempty"`
	Internal   bool      `xml:"Internal" bson:"internal" json:"internal,omitempty"`
}

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
 * Article references used in sections
 */
type ArticleRef struct {
	ArticleID bson.ObjectId `bson:"articleid" json:"articleid"`
}

/* TODO: Fix WebPolls from XML
type ArticlePoll struct {
}

/*
 * Saving Article Ref and shares
*/

type ArticleShares struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	ArticleID bson.ObjectId `bson:"articleid,omitempty" json:"articleid"`
	Origin    string        `bson:"origin" json:"origin"`
	Date      time.Time     `bson:"date" json:"date"`
	FB        struct {
		Shares int `bson:"shares" json:"shares"`
	} `bson:"fb" json:"fb"`
	Twitter struct {
		Shares int `bson:"shares" json:"shares"`
	} `bson:"twitter" json:"twitter"`
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
	Image       string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePath" json:"image,omitempty"`
	ImageByline string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePhotographer" json:"imagebyline,omitempty"`
	Title       string `xml:"StandardArticleTeaserTitle" json:"title,omitempty"`
	Body        string `xml:"StandardArticleTeaserBody" json:"body,omitempty"`
}

type ArticleExtraTeaser struct {
	Title string `xml:"StandardArticleExtraTeaserTitle" json:"title,omitempty"`
	Body  string `xml:"StandardArticleExtraTeaserBody" json:"body,omitempty"`
}

type ArticleByline struct {
	Name       string `xml:"Name" json:"name,omitempty"`
	Email      string `xml:"Email" json:"email,omitempty"`
	Phone      string `xml:"Phone" json:"phone,omitempty"`
	Role       string `xml:"OccupationalRole" json:"role,omitempty"`
	Image      string `xml:"ImagePath" json:"image,omitempty"`
	Department string `xml:"Department" json:"department,omitempty"`
}

type ArticleLinks struct {
	Title    string `xml:"Title" json:"title,omitempty"`
	Internal bool   `xml:"Internal" json:"is_internal,omitempty"`
	Url      string `xml:"Url" json:"url,omitempty"`
	Image    string `xml:"Image" json:"image,omitempty"`
	Pubdate  string `xml:"PubDate" json:"pubdate,omitempty"`
	Preamble string `xml:"Preamble" json:"preamble,omitempty"`
	Category string `xml:"Category" json:"category,omitempty"`
}

type ArticleComments struct {
	ID         string    `xml:"id,attr" json:"id"`
	Title      string    `xml:"Title" json:"title"`
	Body       string    `xml:"Body" json:"body"`
	Author     string    `xml:"Author>AliasOrFullName" json:"author"`
	Pubdate    time.Time `bson:"pubdate" json:"pubdate"`
	RawPubdate string    `xml:"PublicationDate" bson:"-" json:"-"`
	BodyQuote  string    `xml:"BodyQuote" json:"quoted,omitempty"`
}

type ArticleVideo struct {
	Title        string                `xml:"PicSearchVideoTitle" json:"title,omitempty"`
	Description  string                `xml:"PicSearchVideoDescription" json:"description,omitempty"`
	Category     string                `xml:"PicSearchVideoCategory" json:"category,omitempty"`
	ThumbnailUrl string                `xml:"PicSearchVideoThumbNail" json:"tumbnailurl,omitempty"`
	Publishdate  string                `xml:"PicSearchVideoPublishDate" json:"publishdate,omitempty"`
	MediaId      string                `xml:"PicSearchVideoMediaId" json:"mediaid,omitempty"`
	Streams      []ArticleVideoStreams `xml:"PicSearchMediaStreams>PicSearchMediaStream" json:"streams,omitempty"`
}

type ArticleVideoStreams struct {
	StreamFormat string `xml:"PicSearchMediaStreamFormat" json:"format"`
	StreamUri    string `xml:"PicSearchMediaStreamUri" json:"uri"`
}

type ArticleImage struct {
	Url          string `xml:"ImageUrl" json:"url,omitempty"`
	Title        string `xml:"ImageTitle" json:"title,omitempty"`
	Text         string `xml:"ImageText" json:"text,omitempty"`
	Photographer string `xml:"ImagePhotographer" json:"photographer,omitempty"`
	ShowTitle    string `xml:"ImageShowTitle" json:"showtitle,omitempty"`
}

type ArticleImageAlbum struct {
	Id                 string                   `xml:"id,attr" json:"id,omitempty"`
	Title              string                   `xml:"ImageAlbumTitle" json:"title,omitempty"`
	Description        string                   `xml:"ImageAlbumDescription" json:"description,omitempty"`
	Images             []ArticleImageAlbumImage `xml:"ImageAlbumImages>ImageAlbumImage" json:"images,omitempty"`
	AllowComments      string                   `xml:"ImageAlbumAllowComments" json:"allowcomments,omitempty"`
	JsonUrl            string                   `xml:"ImageAlbumJsonUrl" json:"jsonurl,omitempty"`
	SharingLink        string                   `xml:"ImageAlbumSharingLink" json:"sharinglink,omitempty"`
	TeaserTitle        string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImageTitle" json:"teasertitle,omitempty"`
	TeaserDesc         string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImageDescription" json:"teaserdesc,omitempty"`
	TeaserPhotographer string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImagePhotographer" json:"teaserphotographer,omitempty"`
	TeaserImagePath    string                   `xml:"ImageAlbumTeaserImage>ImageAlbumImage>ImageAlbumImagePath" json:"teaserimagepath,omitempty"`
}

type ArticleImageAlbumImage struct {
	Id           string `xml:"id,attr" json:"id,omitempty"`
	Title        string `xml:"ImageAlbumImageTitle" json:"title,omitempty"`
	Description  string `xml:"ImageAlbumImageDescription" json:"description,omitempty"`
	Photographer string `xml:"ImageAlbumImagePhotographer" json:"photographer,omitempty"`
	ImagePath    string `xml:"ImageAlbumImagePath" json:"imagepath,omitempty"`
}

type ArticleFact struct {
	Title string `xml:"Title" json:"title,omitempty"`
	Body  string `xml:"Body" json:"body,omitempty"`
}

/*
 * Article
 */
func (a *Article) SaveToDB(db *mgo.Database) {
	collection := db.C("articles")
	sectCol := db.C("sections")

	a.LastMod = time.Now()
	if len(a.PubdateRaw) > 0 {
		a.Pubdate, _ = time.Parse(time.RFC1123Z, a.PubdateRaw)
	}

	if len(a.ModdateRaw) > 0 {
		a.Moddate, _ = time.Parse(time.RFC1123Z, a.ModdateRaw)
	}

	if a.Moddate.Before(a.Pubdate) {
		a.Moddate = a.Pubdate
	}

	arturl, _ := url.Parse(a.OriginalLink)
	a.OriginSource = arturl.Host

	// Comments
	a.CommentCount = len(a.Comments)

	if len(a.Comments) > 0 {
		for ix, com := range a.Comments {
			if len(com.RawPubdate) > 0 {
				a.Comments[ix].Pubdate, _ = time.Parse(time.RFC1123Z, com.RawPubdate)
			}
		}
	}

	// Article serie
	if len(a.Serie.Articles) > 0 {
		for ix, art := range a.Serie.Articles {
			if len(art.RawPubdate) > 0 {
				a.Serie.Articles[ix].Pubdate, _ = time.Parse(time.RFC1123Z, art.RawPubdate)
			}
		}
	}

	// Find Document
	docToUpdate := bson.M{"originid": a.OriginID}
	savedArticle := Article{}

	err := collection.Find(docToUpdate).One(&savedArticle)
	if err != nil {
		log.Println("Found no document to update, inserting:")

		err = collection.Insert(a)
		if err != nil {
			log.Println("Article SaveToDB: On insert error:", err)
			return
		}

		// Populate from db (so we get the id)
		err = collection.Find(docToUpdate).One(&savedArticle)
		if err != nil {
			log.Println("Found no document to work with:", err)
			return
		}
	}

	// Fields that we set somewhere else ...
	a.Tags = savedArticle.Tags

	if len(a.Id) == 0 {
		if len(savedArticle.Id) > 0 {
			a.Id = savedArticle.Id
		}
	}

	if len(a.Id) > 0 {
		// Find sections with article in it
		findSections := bson.M{"articlelist": bson.M{"$elemMatch": bson.M{"articleid": a.Id}}}
		sects := []ArticleListCommon{}

		err = sectCol.Find(findSections).All(&sects)
		if err != nil {
			log.Println("Article SaveToDB: Found no sections:", err)
		} else {
			// Reset sections, janitor should remove old articles
			a.Sections = []ArticleSection{}

			// Loop sections
			for _, s := range sects {
				// Create new Article sections and set section id as reference
				aSect := ArticleSection{}
				aSect.SectionID = s.ID

				// Loop all articles in sections article list
				for placement, sp := range s.ArticleList {
					// Check if sections article list have the article
					if sp.ArticleID.String() == a.Id.String() {
						aSect.Placement = placement
						break
					}
				}

				a.Sections = append(a.Sections, aSect)
			}
		}
	}

	// Parse body!
	// rr := strings.NewReader(a.Body)
	// doc, err := goquery.NewDocumentFromReader(rr)
	// if err != nil {
	// 	log.Println("Could not read article body")
	// } else {
	// 	bodyParts := []string{}
	// 	doc.Find("p").Each(func(i int, s *goquery.Selection) {
	// 		html, _ := s.Html()
	// 		bodyParts = append(bodyParts, html)
	// 	})

	// 	if len(bodyParts) > 0 {
	// 		a.BodyParts = bodyParts
	// 	}
	// }

	if err = collection.Update(docToUpdate, a); err != nil {
		log.Println("Article SaveToDB: Could not update:", err)
		return
	}

	collection.Find(docToUpdate).One(&a)
	go a.UpdateShares(db)
}

func (a *Article) LoadArticleById(id bson.ObjectId, db *mgo.Database) bool {
	if !id.Valid() {
		log.Println("No valid ID")
		return false
	}

	collection := db.C("articles")

	err := collection.FindId(id).One(&a)
	if err != nil {
		log.Println("LoadArticleById: Could not load article:", err)
		return false
	}

	return true
}

func (a *Article) LoadArticleByOriginId(id string, db *mgo.Database) bool {
	if len(id) == 0 {
		log.Println("No valid ID")
		return false
	}

	collection := db.C("articles")
	query := bson.M{"originid": id}

	err := collection.Find(query).One(&a)
	if err != nil {
		log.Println("LoadArticleById: Could not load article:", err)
		return false
	}

	return true
}

func (a *Article) UpdateShares(db *mgo.Database) {
	if len(a.Id) > 0 {
		collection := db.C("articles")
		sharesCollection := db.C("shares")

		// Find latest share count
		shares := ArticleShares{}
		findQuery := bson.M{"articleid": a.Id}

		err := sharesCollection.Find(findQuery).Sort("-date").Limit(1).One(&shares)
		if err != nil {
			//log.Println("Found no shares for article", a.Id)
			return
		}

		// Check the ID
		if shares.Id.Valid() == false {
			log.Println("No share id thats calid found")
			return
		}

		// Update article with share count
		updateQuery := bson.M{"$set": bson.M{"shares": shares}}

		err = collection.UpdateId(a.Id, updateQuery)
		if err != nil {
			log.Println("Could not update shares for", a.Id)
		}
	}
}

func (a *Article) Update(db *mgo.Database) {
	if a.Id.Valid() {
		collection := db.C("articles")

		err := collection.UpdateId(a.Id, a)
		if err != nil {
			log.Println("Article Update:", err)
			return
		}
	} else {
		log.Println("No ID given. Article not updated")
	}
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
