package routes

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DeepanshuMishraa/conduit.git/models"
	"github.com/DeepanshuMishraa/conduit.git/repository"
	"github.com/DeepanshuMishraa/conduit.git/types"
	"github.com/gin-gonic/gin"
)

func SlowRoute(instanceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		slow := time.Duration(time.Second * 10)
		time.Sleep(slow)

		c.JSON(http.StatusOK, gin.H{
			"message":  "Its a Slow Route",
			"instance": instanceName,
		})
	}
}

func FastRoute(instanceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Its a Fast Route",
			"instance": instanceName,
		})
	}
}

func DBRoute(db *sql.DB, instanceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.JobRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}

		job := &models.Job{
			Title: req.Title,
		}

		createdJob, err := repository.CreateJob(*job, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, types.JobResponse{
			ID:           createdJob.ID,
			Message:      "Job Created Successfully",
			InstanceName: instanceName,
		})
	}
}
