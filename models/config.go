package models

type Config struct {
	DatabaseName		string	`yaml:"databaseName"`
	DatabaseAddr		string	`yaml:"databaseAddr"`
	Port				int32	`yaml:"port"`
	TaskCollection		string	`yaml:"taskCollection"`
	UserCollection		string	`yaml:"userCollection"`
	NoteCollection		string	`yaml:"noteCollection"`
}