package controllers

import (
	"asset_management/config"
	"asset_management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceController struct{}

func NewResourceController() *ResourceController {
	return &ResourceController{}
}

func (rc *ResourceController) GetResources(c *gin.Context) {
	var resources []models.Resource
	if err := config.DB.Find(&resources).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": resources})
}

func (rc *ResourceController) AddResource(c *gin.Context) {
	var resource models.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"resource": resource})
}

func (rc *ResourceController) UpdateResource(c *gin.Context) {
	var resource models.Resource
	id := c.Param("id")
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&resource, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&resource).Updates(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resource": resource})
}

func (rc *ResourceController) DeleteResource(c *gin.Context) {
	id := c.Param("id")
	var resource models.Resource
	if err := config.DB.First(&resource, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
