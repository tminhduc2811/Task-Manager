package repository
import (
	. "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NoteRepository interface {
	FindAll() ([]*Note, error)
	FindById(id string) (*Note, error)
	Search (query string) ([]*Note, error)
	Update(id string, task *Note) error
	Delete(id string) error
	Create(task *Note) error
}

type MgoNoteRepository struct {
	session			*mgo.Session
	dbName			string
	dbCollection	string
}

func NewMgoNoteRepository(session *mgo.Session, dbName string, dbCollection string) *MgoNoteRepository {
	return &MgoNoteRepository{
		session:		session,
		dbName:			dbName,
		dbCollection:	dbCollection,
	}
}

func (r *MgoNoteRepository) getSession() *mgo.Session {
	return r.session.Copy()
}

func (r *MgoNoteRepository) FindAll() ([]*Note, error) {
	var notes []*Note
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(nil).All(&notes)
	return notes, err
}

func (r *MgoNoteRepository) FindById(id string) (*Note, error) {
	var note *Note
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(bson.M{"_id":id}).One(&note)
	return note, err
}

func (r *MgoNoteRepository) Search (query string) ([]*Note, error) {
	var notes []*Note
	session := r.getSession()
	defer session.Clone()
	err := session.DB(r.dbName).C(r.dbCollection).Find(query).All(&notes)
	return notes, err
}

func (r *MgoNoteRepository) Update(id string, note *Note) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &note)
	return err
}

func (r *MgoNoteRepository) Delete(id string) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Remove(id)
	return err
}

func (r *MgoNoteRepository) Create(note *Note) error {
	session := r.getSession()
	defer session.Close()
	err := session.DB(r.dbName).C(r.dbCollection).Insert(&note)
	return err
}
