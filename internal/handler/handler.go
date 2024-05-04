package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proximity_service_go/internal/model"
	"proximity_service_go/internal/service"
	"proximity_service_go/util"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/v1/", welcomeHandler)
	router.GET("/v1/business/:id", getBusinessHandler)
	router.POST("/v1/business", addBusinessHandler)
	router.GET("/v1/business", getAllBusinessesHandler)
	router.GET("/v1/proximity", getNearbyBusinessesHandler)

}

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API!",
	})
}

func getBusinessHandler(c *gin.Context) {
	id := c.Param("id")
	business, err := service.GetBusinessByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Business not found",
		})
		return
	}
	c.JSON(http.StatusOK, business)
}

func getNearbyBusinessesHandler(c *gin.Context) {
	// Get the values of the query parameters
	lat, latErr := util.ConvertToFloat(c.Query("lat"))
	lng, lngErr := util.ConvertToFloat(c.Query("lng"))
	radius, radiusErr := util.ConvertToFloat(c.Query("radius"))
	resolution, resolutionErr := util.ConvertToInt(c.Query("resolution"))

	switch {
	case latErr != nil:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid latitude parameter",
		})
		return
	case lngErr != nil:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid longitude parameter",
		})
		return
	case radiusErr != nil:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid radius parameter",
		})
		return
	case resolutionErr != nil:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid resolution parameter",
		})
	}
	businesses, err := service.FindNearbyBusinesses(lat, lng, radius, resolution)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Business not found",
		})
		return
	}

	c.JSON(http.StatusOK, businesses)
}

func addBusinessHandler(c *gin.Context) {
	var businessRequest model.BusinessRequest

	if err := c.ShouldBindJSON(&businessRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := service.AddBusiness(&businessRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error occurred while adding business",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Business created successfully",
	})
}

func getAllBusinessesHandler(c *gin.Context) {
	results, err := service.GetAllBusinesses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error occurred while fetching all businesses",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"businesses": results,
	})
}
