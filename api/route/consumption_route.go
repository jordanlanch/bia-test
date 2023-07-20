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

func NewConsumptionRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NewConsumptionRepository(db, domain.ConsumptionTable)
	cc := &controller.ConsumptionController{
		ConsumptionUsecase: usecase.NewConsumptionUsecase(cr, timeout),
	}

	group.GET("/consumption", cc.GetConsumptionsByPeriod)
}
