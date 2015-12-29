package pmstructs

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name,omitempty" bson:"name"`
	Email  string        `json:"email,omitempty" bson:"email"`
	Device string        `json:"device" bson:"device"`
}

func (u *User) Get(id string, db *mgo.Database) bool {
	c := db.C("users")

	if bson.IsObjectIdHex(id) {
		// Is an object id
		err := c.FindId(id).One(&u)
		if err != nil {
			log.Println("User Get:", err)
			return false
		}

		return true
	} else {
		if u.Create(db) {
			return true
		} else {
			return false
		}
	}

	return false
}

func (u *User) Create(db *mgo.Database) bool {
	c := db.C("users")

	u.Id = bson.NewObjectId()

	err := c.Insert(u)
	if err != nil {
		log.Println("User Create:", err)
		return false
	}

	return true
}
