package pmstructs

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var remoteTypes = map[string]bool{
	"main":       true,
	"other":      true,
	"subsection": true,
}

type ConfigList struct {
	Configs []Config
}

type Config struct {
	ID           bson.ObjectId    `bson:"_id,omitempty" json:"mid"`
	AppID        string           `json:"appid"`
	Origin       string           `json:"origin" bson:"originhost"`
	PrimaryColor string           `json:"primarycolor" bson:"primarycolor"`
	SectionID    bson.ObjectId    `bson:"sectionid,omitempty" json:"-"`
	PushTags     []ConfigPushTags `bson:"push_tags" json:"push_tags"`
	Sections     []ConfigSections `json:"sections"`
}

type ConfigSections struct {
	Name        string           `json:"name"`
	Url         string           `json:"url"`
	Type        string           `json:"type,omitempty"`
	SectionID   bson.ObjectId    `json:"id,omitempty" bson:"sectionid,omitempty"`
	Icon        string           `bson:"-" json:"icon,omitempty"`
	HeaderColor string           `bson:"headercolor" json:"headercolor"`
	Action      string           `bson:"-" json:"action,omitempty"`
	Subsections []ConfigSections `json:"subsections" bson:"subsections,omitempty"`
}

type ConfigPushTags struct {
	Name  string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
}

func (c *Config) GetSections() []ConfigSections {
	return c.Sections
}

func (c *Config) GetRemoteSections() []ConfigSections {
	remoteSections := []ConfigSections{}

	for _, row := range c.Sections {
		if remoteTypes[row.Type] == true {
			remoteSections = append(remoteSections, row)
		}

		for _, subrow := range row.Subsections {
			if remoteTypes[subrow.Type] == true {
				remoteSections = append(remoteSections, subrow)
			}
		}
	}

	return remoteSections
}

func (c *Config) GetSectionsWithTypeOf(typeof string) []ConfigSections {
	sects := []ConfigSections{}

	for _, sect := range c.Sections {
		if sect.Type == typeof {
			sects = append(sects, sect)
		}

		for _, subsect := range sect.Subsections {
			if subsect.Type == typeof {
				sects = append(sects, subsect)
			}
		}
	}

	return sects
}

func (c *Config) Save(db *mgo.Database) {
	collection := db.C("configs")
	sectionsCollection := db.C("sections")

	// Find created section based on url
	for ix, x := range c.Sections {
		findSect := bson.M{"url": x.Url}

		sect := ArticleListCommon{}
		sect.Url = x.Url
		sect.Origin = c.Origin
		sect.OriginApp = c.AppID
		sect.Url = x.Url
		sect.Type = x.Type
		sect.Save(db)

		err := sectionsCollection.Find(findSect).One(&sect)
		if err != nil {
			if debugMode {
				log.Println("Could not find", findSect, err)
				continue
			}
		}

		c.Sections[ix].SectionID = sect.ID

		if len(x.Subsections) > 0 {
			for ixb, y := range x.Subsections {
				findSect = bson.M{"url": y.Url}

				subsect := ArticleListCommon{}
				subsect.Origin = c.Origin
				subsect.OriginApp = c.AppID
				subsect.Url = y.Url
				subsect.Type = y.Type
				subsect.Save(db)

				err = sectionsCollection.Find(findSect).One(&sect)
				if err != nil {
					if debugMode {
						log.Println("Could not find", findSect, err)
						continue
					}
				}

				c.Sections[ix].Subsections[ixb].SectionID = sect.ID
			}
		}
	}

	find := bson.M{"appid": c.AppID}
	savedConfig := Config{}

	err := collection.Find(find).One(&savedConfig)
	if err != nil {
		err = collection.Insert(c)
		if err != nil {
			log.Println(c, err)
			return
		}

		collection.Find(find).One(&c)
		return
	}

	collection.Update(bson.M{"_id": savedConfig.ID}, c)
}

func (c *Config) GetDUAdminName() string {
	switch c.Origin {
	case "na.se":
		return "na"
	case "lt.se":
		return "lt"
	case "norrteljetidning.se":
		return "nt"
	case "vlt.se":
		return "vlt"
	case "nynashamnsposten.se":
		return "np"
	case "bblat.se":
		return "bblat"
	case "fagersta-posten.se":
		return "fp"
	case "salaallehanda.com":
		return "sa"
	case "avestatidning.com":
		return "at"
	}

	return ""
}
