package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(engine *gin.Engine, app *Application) {

	ac := AddressController{app}
	engine.POST("/address", ac.createAddress)
	engine.GET("/address", ac.getEveryAddress)
	engine.GET("/address/:id", ac.getAddress)
	engine.PUT("/address", ac.updateAddress)
	engine.DELETE("/address/:id", ac.deleteAddress)
	engine.POST("/address/upload", ac.importAddressBook)
	engine.GET("/addressbook", ac.exportAddressBook)
}

func setupRouter(app *Application) *gin.Engine {
	engine := gin.Default()
	initializeRoutes(engine, app)
	return engine
}
