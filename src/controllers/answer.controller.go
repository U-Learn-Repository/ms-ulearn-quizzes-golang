package controllers

import (
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func InsertAnswer(c *gin.Context) {

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}

	req := models.Answer{}

	err := c.Bind(&req)

	req.Id = bson.NewObjectId()

	if err != nil {
		c.JSON(400, gin.H{"message": "Incorrect data", "status": 400, "body": nil})
		return
	} else {
		err := mongo.InsertAnswer(&req)

		if err != nil {
			c.JSON(400, gin.H{"message": "error post to db", "status": 400, "body": nil})
		}
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": req})
	}
}

// UPDATE: /api/v1/answer/:id
func UpdateAnswer(c *gin.Context) {
	id := c.Params.ByName("id")

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	req := models.Answer{}

	err := c.Bind(&req)

	data, err := mongo.UpdateAnswer(id, req)

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": data})
	}
}
