package pmstructs

import (
	"log"
	//"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TV struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Origin   string        `bson:"origin" json:"origin"`
	Articles []Article     `bson:"articles" json:"articles"`
}

func (t *TV) SaveToDB(db *mgo.Database) bool {
	coll := db.C("tvs")

	savedTv := TV{}
	savedTv.LoadByOrigin(t.Origin, db)

	t.ID = savedTv.ID

	err := coll.Update(bson.M{"origin": t.Origin}, t)
	if err != nil {
		log.Println("TV SaveToDB no update can be done:", err)
		return false
	}

	return true
}

func (t *TV) LoadByOrigin(o string, db *mgo.Database) {
	if len(o) > 0 {
		coll := db.C("tvs")
		findQuery := bson.M{"origin": o}

		err := coll.Find(findQuery).One(&t)
		if err != nil {
			log.Println("No can to load TV:", err)

			err = coll.Insert(t)
			if err != nil {
				log.Println("No can insert TV:", err)
				return
			}

			err = coll.Find(findQuery).One(&t)
			if err != nil {
				log.Println("No can load TV:", err)
			}

			return
		}
	}
}
