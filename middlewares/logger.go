package middlewares

import "github.com/gin-gonic/gin"

func LogMiddleware(g *gin.Engine) {
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	g.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	g.Use(gin.Recovery())
}
