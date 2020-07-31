package webserver

import (
	"go-sysview/docs"
	"go-sysview/internal/pkg/api"
	"go-sysview/internal/pkg/binfs"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouter creates the new gin instance
func SetupRouter(port string, debug bool) *gin.Engine {
	// Set the release mode if debug flag is false.
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set Swagger information
	docs.SwaggerInfo.Host = "localhost:" + port

	// Disable log output colors
	gin.DisableConsoleColor()

	// Creates a router with logger and recovery middleware
	router := gin.Default()

	// Register all routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/os", api.GetOperatinSystem)
	}

	// Register swagger route
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	// Register static frontend files
	router.Use(static.Serve("/", binfs.BinaryFileSystem("resources/app")))

	return router
}

// Start the gin webserver with the given port and mode
func Start(port string, debug bool) {
	err := SetupRouter(port, debug).Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start webserver")
	}
}
