package api

import (
	"go-sysview/internal/pkg/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOperatinSystem Returns the system information like operating system, ip address and memory
// @Summary 	Returns some system information
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} controller.OSInformation "Returned object has system information like operating system, ip address and memory"
// @Router 		/os/ [get]
func GetOperatinSystem(c *gin.Context) {
	c.JSON(http.StatusOK, controller.GetOSInformation())
}
