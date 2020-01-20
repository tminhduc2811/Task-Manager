package controllers

import (
	. "../models"
	. "../repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TaskController struct {
	repo TaskRepository
}

func NewTaskController(r TaskRepository) *TaskController {
	return &TaskController{
		repo:r,
	}
}

func (s *TaskController) Create(c *gin.Context) {
	var taskModel *Task
	if err := c.ShouldBindJSON(&taskModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskModel.CreatedOn = time.Now()
	err := s.repo.Create(taskModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, taskModel)
	return
}

func (s *TaskController) FindAll(c *gin.Context){
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

func (s *TaskController) FindById(c *gin.Context) {
	var taskModel *Task
	id := c.Param("id")
	taskModel, err := s.repo.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": taskModel})
}

func (s *TaskController) Search(c *gin.Context) {
	//TODO: Create search model later
}

func (s *TaskController) Update(c *gin.Context) {
	var taskModel *Task
	id := c.Param("id")
	err := c.ShouldBindJSON(&taskModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = s.repo.Update(id, taskModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "Resource updated successfully"})
}

func (s *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := s.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "Resource deleted successfully"})
}
