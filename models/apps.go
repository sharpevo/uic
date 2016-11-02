package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"uic/mongo"
)

type App struct {
	Id      bson.ObjectId   `json:"id" bson:"_id"`
	Name    string          `json:"name"`
	Domain  string          `json:"domain"`
	Enabled bool            `json:"enabled"`
	Remark  string          `json:"remark"`
	Users   map[string]bool `json:"users"`
}

func (app *App) Create() (code int, err error) {
	app.Id = bson.NewObjectId()
	app.DateCreated = time.Now()
	app.Enabled = true
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return ERROR_DATABASE, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("app")
	count, _ := collection.Find(bson.M{"domain": app.Domain}).Count()
	if count != 0 {
		return ERROR_DUPLICATE, err
	}
	err = collection.Insert(app)
	if err != nil {
		return ERROR_DATABASE, err
	}
	return 0, nil
}

func (app *App) FindById(appId string) (app App, err error) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return app, err
	}

	if !bson.IsObjectIdHex(userId) { // panic, objectIdHex ""
		return app, errors.New("user id is invalid")
	}

	collection := session.DB(mongo.MongoConfig.Database).C("app")
	err = collection.FindId(bson.ObjectIdHex(appId)).One(&app)
	return app, err

}

func GetAllApps() (apps []App) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return apps, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("app")
	iter := collection.Find(nil).Sort("name").Iter()
	err = iter.All(&apps)
	return apps, err
}

func (app *App) Update() (code int, err error) {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return ERROR_DATABASE, err
	}
	collection := session.DB(mongo.MongoConfig.Database).C("app")
	err = collection.Update(bson.M{"_id": app.Id}, user)
	beego.Debug("UpdateApp:", app)
	if err != nil {
		beego.Error(err)
	}
	return ERROR_DATABASE, err
}

func (app *App) AddUser(userId string) {
	if app.Users == nil {
		app.Users = make(map[string]bool)
	}
	app.Users[userId] = true
	_, err := app.Update()
	user := User{}
	user.FindById(userId)
	user.AddApp(app.Id.Hex())
	return err
}

func (app *App) RemoveUser(userId string) {
	delete(app.Users, userId)
	_, err := app.Update()
	user := User{}
	user.FindById(userId)
	user.RemoveApp(app.Id.Hex())
	return err
}

func (app *App) Delete() error {
	session, err := mongo.CopyMasterSession()
	if err != nil {
		return err
	}

	for userId := range app.Users {
		user := User{}
		user.FindById(userId)
		user.RemoveApp(app.Id.Hex())
	}

	collection := session.DB(mongo.MongoConfig.Database).C("app")
	err = collection.RemoveId(app.Id)
	return err
}
