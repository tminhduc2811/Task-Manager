package repository


import (
	. "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository interface {
	FindAll() ([]*User, error)
	FindOne(query interface{}) (*User, error)
	Search (query interface{}) ([]*User, error)
	Update(id string, task *User) error
	Delete(id string) error
	Create(task *User) error
	Exists(query interface{}) (bool, error)
}

type MgoUserRepository struct {
	session			*mgo.Session
	dbName			string
	dbCollection	string
}

func NewMgoUserRepository(session *mgo.Session, dbName string, dbCollection string) *MgoUserRepository {
	return &MgoUserRepository{
		session:		session,
		dbName:			dbName,
		dbCollection:	dbCollection,
	}
}

func (r *MgoUserRepository) getSession() *mgo.Session {
	return r.session.Copy()
}

func (r *MgoUserRepository) FindAll() ([]*User, error) {
	var user []*User
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(nil).All(&user)
	return user, err
}

func (r *MgoUserRepository) FindOne(query interface{}) (*User, error) {
	var user *User
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).One(&user)
	return user, err
}

func (r *MgoUserRepository) Search (query interface{}) ([]*User, error) {
	var users []*User
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).All(&users)
	return users, err
}

func (r *MgoUserRepository) Update(id string, user *User) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &user)
	return err
}

func (r *MgoUserRepository) Delete(id string) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Remove(id)
	return err
}

func (r *MgoUserRepository) Create(user *User) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Insert(&user)
	return err
}

func (r *MgoUserRepository) Exists(query interface{}) (bool, error) {
	session := r.getSession()
	defer session.Close()
	count, err := session.DB(r.dbName).C(r.dbCollection).Find(query).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}