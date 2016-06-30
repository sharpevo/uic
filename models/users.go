package models

import (
	"encoding/hex"
	"github.com/agnivade/easy-scrypt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"ssologin/mongo"
)

type User struct {
	Id       bson.ObjectId     `json:"id" bson:"_id"`
	Name     string            `json:"name"`
	Email    string            `json:"Email"`
	Password string            `json:"password`
	Roles    map[string]bool   `json:"roles"`
	Tokens   map[string]string `json:"-"`
}

func (user *User) FindByEmail(email string) (code int, err error) {
	session, err := mongo.CopyMasterSession()

	if err != nil {
		return ERROR_DATABASE, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("user")
	err = collection.Find(bson.M{"email": email}).One(user)
	if err != nil {
		return ERROR_NOT_FOUND, err
	}

	return 0, nil

}

func (user *User) CheckPass(password string) (ok bool, err error) {
	passwordKey, err := hex.DecodeString(user.Password)
	if err != nil {
		beego.Error("DecodePassword:", err)
	}

	correct, err := scrypt.VerifyPassphrase(password, passwordKey)

	if err != nil {
		beego.Error("VerifyPassphrase:", err)
	}

	return correct, err
}

func (user *User) AddToken(iss string, token string) error {
	if user.Tokens == nil {
		user.Tokens = make(map[string]string)
	}

	user.Tokens[iss] = token
	_, err := user.Update()
	return err
}

func (user *User) Update() (code int, err error) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return ERROR_DATABASE, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("user")
	err = collection.Update(bson.M{"_id": user.Id}, user)
	beego.Debug("UpdateUser:", user)
	if err != nil {
		beego.Error(err)
	}
	return ERROR_DATABASE, err
}
