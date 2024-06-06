package blueprints

import (
	dbcontroller "rfid_payment/db/db_controller"
	"rfid_payment/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	control *dbcontroller.HandlersController
}

func ServiceRouther(control *dbcontroller.HandlersController) *Router {
	return &Router{control: control}
}

func (r *Router) Start(port string) {
	router := gin.New()

	router.POST("/user/add_new_user", services.CreateNewUser)

	router.POST("/user/delete_user", services.DeleteUser)

	router.POST("/admin/new_admin", services.NewAdmin)

	router.POST("/admin/delete_admin", services.DeleteAdmin)

	router.POST("/transaction/topup", services.TopUp)

	router.POST("/transaction/pay", services.Transaction)

	// Card-Tracker
	router.POST("/tracker/add_new_user", services.TrackerCreateNewUser)

	router.Run(port)
}
