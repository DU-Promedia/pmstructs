package pmstructs

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	//"sync"
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
	Pubdate         time.Time          `json:"pubdate,omitempty" bson:"pubdate"`
	Moddate         time.Time          `json:"moddate,omitempty" bson:"moddate"`
	Location        string             `xml:"Location" json:"location,omitempty"`
	Latitude        string             `xml:"StandardArticleGeo>StandardArticleLatitude" json:"latitude,omitempty" bson:"latitude"`    // OUT
	Longitude       string             `xml:"StandardArticleGeo>StandardArticleLongitude" json:"longitude,omitempty" bson:"longitude"` // OUT
	Department      string             `xml:"ArticleDepartment" json:"department,omitempty"`                                           // OUT
	Teaser          ArticleTeaser      `xml:"StandardArticleTeaser" json:"teaser"`                                                     // OUT
	ExtraTeaser     ArticleExtraTeaser `xml:"StandardArticleExtraTeaser" bson:"extrateaser" json:"-"`
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
	Shares          ArticleShares      `bson:"shares" json:"shares,omitempty"`                                               // OUT
	Serie           ArticleSerie       `xml:"StandardArticleArticleSeries>ArticleSerie" bson:"serie" json:"serie,omitempty"` // OUT
	Poll            ArticlePoll        `xml:"WebPolls>WebPoll" bson:"poll" json:"poll,omitempty"`                            // OUT
	//Sections        []ArticleSection   `bson:"sections" json:"sections,omitempty"`
}

/*
 * Webpoll
 */
type ArticlePoll struct {
	PollId         string `xml:"id,attr" bson:"pollid" json:"pollid,omitempty"`
	RefId          string `xml:"refId,attr" bson:"refid" json:"refid,omitempty"`
	ContentId      string `xml:"contentId,attr" bson:"contentid" json:"id,omitempty"`
	Header         string `xml:"Header" bson:"header" json:"header,omitempty"`
	Question       string `xml:"Question" bson:"question" json:"question,omitempty"`
	TotalVoteCount int    `xml:"TotalVoteCount" bson:"totalvotecount" json:"totalvotecount,omitempty"`
	InActive       bool   `xml:"InActive" bson:"inactive" json:"inactive,omitempty"`
	Options        []struct {
		Id     string `xml:"id,attr" bson:"id" json:"id,omitempty"`
		RefId  string `xml:"refId,attr" bson:"refid" json:"refid,omitempty"`
		Label  string `xml:"Label" bson:"label" json:"label,omitempty"`
		Weight int    `xml:"Weight" bson:"weight" json:"weight,omitempty"`
	} `xml:"Options>Option" bson:"options" json:"options,omitempty"`
	Results []struct {
		Id            string  `xml:"Id" bson:"id" json:"id,omitempty"`
		Label         string  `xml:"Label" bson:"label" json:"label,omitempty"`
		Weight        int     `xml:"Weight" bson:"weight" json:"weight,omitempty"`
		Offset        int     `xml:"Offset" bson:"offset" json:"offset,omitempty"`
		PercentageInt int     `xml:"PercentageInt" bson:"percentageint" json:"percentageint,omitempty"`
		Value         int     `xml:"Value" bson:"value" json:"value,omitempty"`
		Percentage    float32 `xml:"Percentage" bson:"percentage" json:"percentage,omitempty"`
		RawValue      int     `xml:"RawValue" bson:"rawvalue" json:"rawvalue,omitempty"`
	} `xml:"Results>Result" bson:"results" json:"results,omitempty"`
}

/*
 * Article Serie
 */
type ArticleSerie struct {
	Title    string                `xml:"Title" bson:"title" json:"title,omitempty"`
	Articles []ArticleSerieArticle `xml:"Articles>Article" bson:"articles" json:"articles,omitempty"`
}

func (a *ArticleSerie) TrigUpdateOfSiblings(db *mgo.Database) {
	// var wg sync.WaitGroup
	// wg.Add(len(a.Articles))

	// for _, art := range a.Articles {
	// 	log.Println("Updating sibling ...", art.OriginID)

	// 	go func(a ArticleSerieArticle) {
	// 		a.UpdateFromSource(db)

	// 		wg.Done()
	// 	}(art)
	// }

	// log.Println("TrigUpdateOfSiblings Waiting ...")
	// wg.Wait()
}

type ArticleSerieArticle struct {
	OriginID   string    `bson:"originid" json:"id"`
	Title      string    `xml:"Title" bson:"title" json:"title,omitempty"`
	Preamble   string    `xml:"Preamble" bson:"preamble" json:"preamble,omitempty"`
	Link       string    `xml:"Link" bson:"link" json:"link,omitempty"`
	Image      string    `xml:"Image" bson:"image" json:"image,omitempty"`
	Pubdate    time.Time `bson:"pubdate" json:"pubdate,omitempty"`
	RawPubdate string    `xml:"PubDate" bson:"-" json:"-"`
	Location   string    `xml:"Location" bson:"location" json:"location,omitempty"`
	Internal   bool      `xml:"Internal" bson:"internal" json:"internal,omitempty"`
}

func (a *ArticleSerieArticle) UpdateFromSource(db *mgo.Database) {

	if len(a.OriginID) == 0 {
		return
	}

	parsed, err := url.Parse(a.Link)
	if err != nil {
		log.Println("ArticleSerieArticle UpdateFromSource:", err)
		return
	}

	art := Article{}
	art.LoadArticleByOriginId(a.OriginID, db)
	log.Println("UpdateFromSource")
	err = art.GetArticleFromUrl(parsed.Host, a.OriginID, db)
	if err != nil {
		log.Println("UpdateFromSource ERROR:", err)
		return
	}

	log.Println("UpdateFromSource DONE!")
	return
}

/*
 * Article references used in sections
 */
type ArticleRef struct {
	ArticleID bson.ObjectId `bson:"articleid" json:"articleid"`
}

/*
 * Saving Article Ref and shares
 */

type ArticleShares struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ArticleID bson.ObjectId `bson:"articleid,omitempty" json:"articleid,omitempty"`
	Origin    string        `bson:"origin" json:"origin,omitempty"`
	Date      time.Time     `bson:"date" json:"date,omitempty"`
	FB        struct {
		Shares int `bson:"shares" json:"shares,omitempty"`
	} `bson:"fb" json:"fb,omitempty"`
	Twitter struct {
		Shares int `bson:"shares" json:"shares,omitempty"`
	} `bson:"twitter" json:"twitter,omitempty"`
}

/*
 * Articles section struct
 */
type ArticleSection struct {
	SectionID bson.ObjectId `json:"id" bson:"sectionid"`
	Placement int           `json:"placement" bson:"placement"`
}

/*
 * Teaser article
 */
type TeaserArticle struct {
	OriginID string `xml:"id,attr" bson:"originid" json:"id"`
	Image    string `xml:"TeaserArticleImage>TeaserArticleImagePath" bson:"image" json:"image,omitempty"`
	Title    string `xml:"TeaserArticleTitle" bson:"title" json:"title"`
	Body     string `xml:"TeaserArticlePreamble" bson:"body" json:"body,omitempty"`
	Link     string `xml:"TeaserArticleExternal>TeaserArticleExternalLink" bson:"link" json:"link,omitempty"`
}

/*
 * Article parts
 */
type ArticleTeaser struct {
	Image       string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePath" json:"image,omitempty"`
	ImageByline string `xml:"StandardArticleTeaserImage>StandardArticleTeaserImagePhotographer" json:"imagebyline,omitempty"`
	Title       string `xml:"StandardArticleTeaserTitle" json:"title,omitempty"`
	Body        string `xml:"StandardArticleTeaserBody" json:"body,omitempty"`
	Link        string `bson:"link" json:"link,omitempty"`
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
				a.Serie.Articles[ix].OriginID = GetOriginIdFromUrl(art.Link)
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
	//a.Shares = savedArticle.Shares

	if len(a.Id) == 0 {
		if len(savedArticle.Id) > 0 {
			a.Id = savedArticle.Id
		}
	}

	if err = collection.Update(docToUpdate, a); err != nil {
		log.Println("Article SaveToDB: Could not update:", err)
		return
	}

	collection.Find(docToUpdate).One(&a)

	go a.UpdateShares(db)

	return
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
		log.Println("LoadArticleByOriginId: Could not load article:", err)
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
			//log.Println("No share id thats valid found")
			return
		}

		if shares.FB.Shares > 0 {
			// Update article with share count
			set := bson.M{"shares": shares}
			updateQuery := bson.M{"$set": set}

			err = collection.Update(bson.M{"_id": a.Id}, updateQuery)
			if err != nil {
				log.Println("Could not update shares for", a.Id)
				return
			}

			a.Shares = shares
		}
	}
}

func (a *Article) GetArticleFromUrl(host string, id string, db *mgo.Database) error {
	// Create URL
	uri := "http://" + host + "/" + id + "?m=mobile"

	// Do the response
	if len(host) == 0 || len(id) == 0 {
		return errors.New("No host and ID given correctly.")
	}

	time.Sleep(50 * time.Millisecond)

	response, err := http.Get(uri)
	if err != nil {
		return errors.New("Could not load URL")
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println("Source did not answer correclty. " + response.Status)
		return errors.New("HTTP Error: " + response.Status)
	}

	// Read contents
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Could not read body")
	}

	if a.Id.Valid() == false {
		a.Id = bson.NewObjectId()
	}

	// Unmarshal article
	err = xml.Unmarshal(content, &a)
	if err != nil {
		return errors.New("Could not unmarshal")
	}

	a.SaveToDB(db)

	return nil
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

func GetOriginIdFromUrl(s string) string {
	r, _ := regexp.Compile(`(1.[0-9])\w+`)
	match := r.FindString(s)

	return match
}
