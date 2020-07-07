package master

import (
	"database/sql"
	"liveCodeAPI/api-master/master/controllers"
	"liveCodeAPI/api-master/master/repositories/goodsrepository"
	"liveCodeAPI/api-master/master/repositories/userrepository"
	"liveCodeAPI/api-master/master/repositories/warehouserepository"
	"liveCodeAPI/api-master/master/usecases/goodsusecase"
	"liveCodeAPI/api-master/master/usecases/userusecase"
	"liveCodeAPI/api-master/master/usecases/warehouseusecase"
	"liveCodeAPI/middleware"

	"github.com/gorilla/mux"
)

func InitData(r *mux.Router, db *sql.DB) {

	whRepo := warehouserepository.InitWHRepoImpl(db)
	whUsecase := warehouseusecase.InitWHUsecase(whRepo)
	controllers.WarehouseController(r, whUsecase)

	goodsRepo := goodsrepository.InitGoodsRepoImpl(db)
	goodsUsecase := goodsusecase.InitGoodsUsecase(goodsRepo)
	controllers.GoodsController(r, goodsUsecase)

	userRepo := userrepository.InitUserRepoImpl(db)
	userUsecase := userusecase.InitUserUsecase(userRepo)
	controllers.UserController(r, userUsecase)
	r.Use(middleware.ActivityLogMiddleware)
}
