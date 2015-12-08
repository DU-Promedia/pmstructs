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

type Config struct {
	ID        bson.ObjectId    `bson:"_id,omitempty" json:"id"`
	AppID     string           `json:"appid"`
	Origin    string           `json:"origin" bson:"originhost"`
	SectionID bson.ObjectId    `bson:"sectionid,omitempty" json:"mid"`
	Sections  []ConfigSections `json:"sections"`
}

type ConfigSections struct {
	Name        string           `json:"name"`
	Url         string           `json:"url"`
	Type        string           `json:"type"`
	SectionID   bson.ObjectId    `json:"mid" bson:"sectionid,omitempty"`
	Subsections []ConfigSections `json:"subsections" bson:"subsections,omitempty"`
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

func (c *Config) Save(db *mgo.Database) {
	collection := db.C("configs")
	sectionsCollection := db.C("sections")

	// Find created section based on url
	for ix, x := range c.Sections {
		findSect := bson.M{"url": x.Url}
		sect := ArticleListCommon{}

		err := sectionsCollection.Find(findSect).One(&sect)
		if err != nil {
			if debugMode {
				log.Println("Could not find", findSect, err)
				log.Println("Inserting ...")
			}

			insect := ArticleListCommon{}
			insect.Origin = c.Origin
			insect.Url = x.Url

			err = sectionsCollection.Insert(insect)
			if err != nil {
				log.Println(err)
				continue
			}

			sectionsCollection.Find(findSect).One(&sect)
		}

		c.Sections[ix].SectionID = sect.ID

		if len(x.Subsections) > 0 {
			for ixb, y := range x.Subsections {
				findSect = bson.M{"url": y.Url}

				err = sectionsCollection.Find(findSect).One(&sect)
				if err != nil {
					if debugMode {
						log.Println("Could not find", findSect, err)
						log.Println("Inserting ...")
					}

					insect := ArticleListCommon{}
					insect.Origin = c.Origin
					insect.Url = y.Url

					err = sectionsCollection.Insert(insect)
					if err != nil {
						log.Println(err)
						continue
					}

					sectionsCollection.Find(findSect).One(&sect)

					continue
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
