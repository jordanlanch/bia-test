package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/bia-test/api/controller"
	"github.com/jordanlanch/bia-test/bootstrap"
	"github.com/jordanlanch/bia-test/domain"
	"github.com/jordanlanch/bia-test/repository"
	"github.com/jordanlanch/bia-test/usecase"
	"gorm.io/gorm"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserTable)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
