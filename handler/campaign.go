package handler

import (
	"govue/campaign"
	"govue/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.Query("user_id"))

	campains, err := h.service.GetCampaigns(UserID)
	if err != nil {
		response := helpers.APIresponse("Error get Campaigns!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIresponse("List Of Campaigns!", http.StatusOK, "success", campains)
	c.JSON(http.StatusOK, response)
}
