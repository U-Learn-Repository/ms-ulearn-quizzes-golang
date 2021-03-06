package database

import (
	"fmt"
	"time"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	mgo "gopkg.in/mgo.v2"
)

const IP = "192.168.1.73";
const PORT = "7778";

type MongoDB struct {
	Host             string
	Port             string
	Addrs            string
	Database         string
	EventTTLAfterEnd time.Duration
	StdEventTTL      time.Duration
	Info             *mgo.DialInfo
	Session          *mgo.Session
}

func (mongo *MongoDB) SetDefault() {
	mongo.Host = "mongodb"
	// mongo.Addrs = "ulearn-quizzes-db:27017"
	mongo.Addrs = IP+":"+PORT;
	mongo.Database = "quices"

	mongo.EventTTLAfterEnd = 1 * time.Second
	mongo.StdEventTTL = 20 * time.Minute
	mongo.Info = &mgo.DialInfo{
		Addrs:    []string{mongo.Addrs},
		Timeout:  60 * time.Second,
		Database: mongo.Database,
	}
}

func (mongo *MongoDB) Drop() (err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.DB(mongo.Database).DropDatabase()
	if err != nil {
		return err
	}
	return nil
}

func (mongo *MongoDB) Init() (err error) {
	err = mongo.Drop()
	if err != nil {
		fmt.Printf("\nDrop database error: %v\n", err)
	}

	question := models.Question{}
	question.Statement = "Esta es un ejemplo de una pregunta"
	question.Score = 100
	question.UserId = 1
	err = mongo.InsertQuestion(&question)

	return err
}

func (mongo *MongoDB) SetSession() (err error) {
	mongo.Session, err = mgo.DialWithInfo(mongo.Info)
	if err != nil {
		mongo.Session, err = mgo.Dial(mongo.Host)
		if err != nil {
			return err
		}
	}
	return err
}

