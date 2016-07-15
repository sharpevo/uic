package models

import (
	"encoding/hex"
	"github.com/agnivade/easy-scrypt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"time"
	"uic/mongo"
)

type UserInfo struct {
	Id    string
	Name  string
	Email string
	Roles string
}

type User struct {
	Id              bson.ObjectId     `json:"id" bson:"_id"`
	Name            string            `json:"name"`
	Email           string            `json:"Email"`
	Password        string            `json:"password`
	Roles           map[string]bool   `json:"roles"`
	Tokens          map[string]string `json:"-"`
	DateCreated     time.Time
	DateLastLogin   time.Time
	DateResetPasswd time.Time
	Enabled         bool
	Deleted         bool
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

func CheckEmail(email string) (existed bool, err error) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return true, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("user")
	count, err := collection.Find(bson.M{"email": email}).Count()
	if count != 0 {
		return true, err
	}
	return false, err
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
	user.DateLastLogin = time.Now()
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

func (user *User) FindById(id string) (code int, err error) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return ERROR_DATABASE, err
	}

	collection := session.DB(mongo.MongoConfig.Database).C("user")

	if !bson.IsObjectIdHex(id) {
		return ERROR_INPUT, err
	}

	err = collection.FindId(bson.ObjectIdHex(id)).One(user)
	if err != nil {
		return ERROR_NOT_FOUND, err
	}

	return 0, nil
}

func (user *User) HasToken(iss string, token string) bool {
	return user.Tokens[iss] == token
}

func (user *User) RemoveToken(iss string) bool {
	delete(user.Tokens, iss)
	user.Update()
	_, ok := user.Tokens[iss]
	return !ok
}

func (user *User) RemoveAllTokens() error {
	user.Tokens = map[string]string{}
	_, err := user.Update()
	return err
}

func ParseUser(form *RegisterForm) (u *User, err error) {
	passwordKey, _ := scrypt.DerivePassphrase(form.Password, 32)
	user := User{
		Name:     form.Name,
		Email:    form.Email,
		Password: hex.EncodeToString(passwordKey),
	}
	return &user, nil
}

func (user *User) EncryptPassword(password string) (encpasswd string) {
	passwordKey, _ := scrypt.DerivePassphrase(password, 32)
	return hex.EncodeToString(passwordKey)

}

func (user *User) Create() (code int, err error) {
	user.Id = bson.NewObjectId()
	user.DateCreated = time.Now()
	user.Enabled = true
	user.Deleted = false
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return ERROR_DATABASE, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("user")
	count, _ := collection.Find(bson.M{"email": user.Email}).Count()
	if count != 0 {
		return ERROR_DUPLICATE, err
	}
	err = collection.Insert(user)
	if err != nil {
		return ERROR_DATABASE, err
	}
	return 0, nil
}

func GetUsersSortByEmail() ([]User, error) {
	session, err := mongo.CopyMasterSession()
	users := []User{}
	if err != nil {
		return users, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("user")
	iter := collection.Find(nil).Sort("email").Iter()
	err = iter.All(&users)
	return users, err
}

func (user *User) AddRole(roleName string) error {
	if user.Roles == nil {
		user.Roles = make(map[string]bool)
	}
	user.Roles[roleName] = true
	_, err := user.Update()
	return err
}

func (user *User) RemoveRole(roleName string) error {
	delete(user.Roles, roleName)
	_, err := user.Update()
	return err
}
