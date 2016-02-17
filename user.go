package pmstructs

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	OldId            string        `json:"-" bson:"oldid,omitempty"`
	Name             string        `json:"name,omitempty" bson:"name"`
	Email            string        `json:"email,omitempty" bson:"email"`
	AppId            string        `json:"-" bson:"appid"`
	Device           string        `json:"device" bson:"device"`
	Loads            int           `json:"-" bson:"loads"`
	LastSeenTakeover time.Time     `json:"-" bson:"lastseentakeover"`
	CreateDate       time.Time     `json:"created" bson:"created"`
	LastLoad         time.Time     `json:"-" bson:"lastload"`
}

func (u *User) Get(id string, db *mgo.Database) bool {
	c := db.C("users")

	if bson.IsObjectIdHex(id) {
		// Is an object id
		oid := bson.ObjectIdHex(id)
		err := c.FindId(oid).One(&u)
		if err != nil {
			log.Println("User Get:", err)

			return u.Create(db)
		}

		return true
	} else {
		return u.Create(db)
	}

	return false
}

func (u *User) Create(db *mgo.Database) bool {
	c := db.C("users")

	u.Id = bson.NewObjectId()
	u.Loads = 0
	u.CreateDate = time.Now()

	err := c.Insert(u)
	if err != nil {
		log.Println("User Create:", err)
		return false
	}

	return true
}

func (u *User) UpdateTakeover(d time.Time, db *mgo.Database) {
	c := db.C("users")

	if u.Id.Valid() == false {
		log.Println("User UpdateTaker: No update made, no id given:", u.Id.String())
		return
	}

	query := bson.M{"$set": bson.M{"lastseentakeover": d}}
	err := c.UpdateId(u.Id, query)
	if err != nil {
		log.Println("User UpdateTakeover:", err)
	}
}

func (u *User) UpdateLoads(db *mgo.Database) {
	c := db.C("users")

	query := bson.M{"$inc": bson.M{"loads": 1}, "$set": bson.M{"lastload": time.Now()}}
	err := c.UpdateId(u.Id, query)
	if err != nil {
		log.Println("User UpdateLoads:", err)
	}
}
