package repository

import (
	. "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskRepository interface {
	FindAll() ([]*Task, error)
	FindById(id string) (*Task, error)
	Search (query string) ([]*Task, error)
	Update(id string, task *Task) error
	Delete(id string) error
	Create(task *Task) error
}

type MgoTaskRepository struct {
	session			*mgo.Session
	dbName			string
	dbCollection	string
}

func NewMgoTaskRepository(session *mgo.Session, dbName string, dbCollection string) *MgoTaskRepository {
	return &MgoTaskRepository{
		session:		session,
		dbName:			dbName,
		dbCollection:	dbCollection,
	}
}

func (r *MgoTaskRepository) getSession() *mgo.Session {
	return r.session.Copy()
}

func (r *MgoTaskRepository) FindAll() ([]*Task, error) {
	var tasks []*Task
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(nil).All(&tasks)
	return tasks, err
}

func (r *MgoTaskRepository) FindById(id string) (*Task, error) {
	var result *Task
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(bson.M{"_id":id}).One(&result)
	return result, err
}

func (r *MgoTaskRepository) Search (query string) ([]*Task, error) {
	var tasks []*Task
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).All(&tasks)
	return tasks, err
}

func (r *MgoTaskRepository) Update(id string, task *Task) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &task)
	return err
}

func (r *MgoTaskRepository) Delete(id string) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Remove(id)
	return err
}

func (r *MgoTaskRepository) Create(task *Task) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Insert(&task)
	return err
}
