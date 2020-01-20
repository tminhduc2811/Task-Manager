package controllers

import (
	. "../models"
	. "../repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	repo UserRepository
}

func NewUserController(r UserRepository) *UserController {
	return &UserController{
		repo:r,
	}
}

func (s *UserController) Create(c *gin.Context) {
	var userModel *User
	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userExists, err := s.repo.Exists(bson.M{"email": userModel.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userExists {
		c.JSON(http.StatusForbidden, gin.H{"error": "User name already existed"})
		return
	}

	userModel.Id = bson.NewObjectId()
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(userModel.PassWord), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	userModel.HashPassword = hashpassword
	userModel.PassWord = ""
	if err := s.repo.Create(userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userModel.HashPassword = nil
	c.JSON(http.StatusOK, userModel)
}

func (s *UserController) FindAll(c *gin.Context){
	resp, err := s.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = json.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})

}

func (s *UserController) FindById(c *gin.Context) {
	var userModel *User
	id := c.Param("id")
	userModel, err := s.repo.FindOne(bson.M{"_id":id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userModel})
}

func (s *UserController) Search(c *gin.Context) {
	//TODO: Create search model later
}

func (s *UserController) Update(c *gin.Context) {
	var userModel *User
	id := c.Param("id")
	err := c.ShouldBindJSON(&userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = s.repo.Update(id, userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "Resource updated successfully"})
}

func (s *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := s.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "Resource deleted successfully"})
}
