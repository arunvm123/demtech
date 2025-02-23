package main

import (
	"log"
	"net/http"

	"github.com/arunvm123/demtech/model"
	"github.com/gin-gonic/gin"
)

type handlerGetLogsAggregatesResponse struct {
	UserName string `json:"UserName,omitempty"`
	Scenario string `json:"Scenario"`
	Count    int    `json:"Count"`
}

// @Summary Get aggregated logs
// @Description Retrieves aggregated logs with optional filtering by username
// @Accept json
// @Produce json
// @Param username query string false "Filter logs by username"
// @Success 200 {array} handlerGetLogsAggregatesResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /logs [get]
func (server *server) handlerGetLogsAggregates(c *gin.Context) {

	username := c.Query("username")

	dbArgs := model.GetAggregatedLogsArgs{
		UserName: nil,
	}
	if len(username) > 0 {
		dbArgs.UserName = &username
	}

	response, err := server.db.GetAggregatedLogs(dbArgs)
	if err != nil {
		log.Printf("Error while calling GetAggregatedLogs: %v \n", err.Error())
		c.JSON(http.StatusInternalServerError, "Internal Sever Error")
		return
	}

	c.JSON(http.StatusOK, mapAggregatedLogsToResponse(response))
	return
}

func mapAggregatedLogsToResponse(logs []model.AggregatedLog) []handlerGetLogsAggregatesResponse {
	response := make([]handlerGetLogsAggregatesResponse, len(logs))

	for i, log := range logs {
		response[i] = handlerGetLogsAggregatesResponse{
			UserName: log.UserName,
			Scenario: log.Scenario,
			Count:    log.Count,
		}
	}

	return response
}
