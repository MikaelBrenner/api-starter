package service

import (
	"api-starter/models/db"
	"api-starter/models/entities"
	"errors"
	"github.com/goonode/mogo"
	"labix.org/v2/mgo/bson"
)

type UserService struct{}

func (service UserService) Create(user *entities.User) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entities.User{}).(*entities.User)
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err == nil {
		return errors.New("user already exists")
	}
	userModel := mogo.NewDoc(user).(*entities.User)
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

func (service UserService) Update(user *entities.User) error {
	conn := db.GetConnection()
	defer conn.Session.Close()

	userToUpdate := mogo.NewDoc(entities.User{}).(*entities.User)
	err := userToUpdate.FindOne(bson.M{"email": user.Email}, userToUpdate)
	if err != nil {
		return err
	}
	userToUpdate.Name = user.Name
	userToUpdate.Info = user.Info
	err = mogo.Save(userToUpdate)
	return err
}

func (service UserService) Find(user *entities.User) (*entities.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(entities.User{}).(*entities.User)
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (service UserService) FindByEmail(email string) (*entities.User, error) {
	conn := db.GetConnection()
	defer conn.Session.Close()

	user := new(entities.User)
	user.Email = email
	return service.Find(user)
}
