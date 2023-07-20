package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/bia-test/api/middleware"
	"github.com/jordanlanch/bia-test/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewMeterRouter(env, timeout, db, protectedRouter)
	NewConsumptionRouter(env, timeout, db, protectedRouter)
}
