package api

import (
	"github.com/banzaicloud/pipeline/ssh"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Get SSH Key
func SshKeyGet(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Get SSH KEY"})
	kubeConfig, ok := GetK8sConfig(c)
	if ok != true {
		return
	}
	response, err := ssh.KeyGet(kubeConfig)
	if err != nil {
		log.Error("Error msg", err.Error())
		//c.JSON(http.StatusBadRequest, htype.ErrorResponse{
		//	Code:    http.StatusBadRequest,
		//	Message: "Error listing deployments",
		//	Error:   err.Error(),
		//})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

// Add SSH Key
func SshKeyAdd(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Get SSH KEY"})
	kubeConfig, ok := GetK8sConfig(c)
	if ok != true {
		return
	}
	response, err := ssh.KeyAdd(kubeConfig)
	if err != nil {
		log.Error("Error msg", err.Error())
		//c.JSON(http.StatusBadRequest, htype.ErrorResponse{
		//	Code:    http.StatusBadRequest,
		//	Message: "Error listing deployments",
		//	Error:   err.Error(),
		//})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

// Delete SSH Key
func SshKeyDelete(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{"tag": "Get SSH KEY"})
	kubeConfig, ok := GetK8sConfig(c)
	if ok != true {
		return
	}
	response, err := ssh.KeyDelete(kubeConfig)
	if err != nil {
		log.Error("Error msg", err.Error())
		//c.JSON(http.StatusBadRequest, htype.ErrorResponse{
		//	Code:    http.StatusBadRequest,
		//	Message: "Error listing deployments",
		//	Error:   err.Error(),
		//})
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
