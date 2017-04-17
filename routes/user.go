package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/mike-dunton/menuCreator/controllers"
	"github.com/mike-dunton/menuCreator/models/user"
	"github.com/mike-dunton/menuCreator/mongo"
	"gopkg.in/validator.v2"
)

// user middlewear
func addUser(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var user userModel.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Unable to unmarshal your requst")
		c.JSON(400, "Unable to unmarshal your requst")
		return
	}
	err = validator.Validate(user)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Validation Error")
		c.JSON(400, err.Error())
		return
	}
	userController := &controllers.UserController{}
	err = userController.NewController()
	defer mongo.CloseSession(userController.Service.MongoSession)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	code, respBody, _ := userController.NewUser(user)
	c.JSON(code, respBody)
}
