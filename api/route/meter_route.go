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

func NewMeterRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	mr := repository.NewMeterRepository(db, domain.MeterTable)
	mc := &controller.MeterController{
		MeterUsecase: usecase.NewMeterUsecase(mr, timeout),
	}

	group.GET("/meters", mc.Fetch)
	group.GET("/meters/:id", mc.Get)
}
