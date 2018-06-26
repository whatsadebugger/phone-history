package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// PanicIfError will panic if err != nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func publicError(c *gin.Context, statusCode int, err error) {
	c.Error(err).SetType(gin.ErrorTypePublic)
	c.JSON(statusCode, gin.H{"errors": []string{err.Error()}})
}

func (app *Application) findAddress(id string) (address, error) {
	var ad address
	aid, err := strconv.Atoi(id)
	if err != nil {
		return ad, err
	}
	err = app.Database.One("ID", aid, &ad)
	return ad, err
}

func mustNotErr(err error) {
	if err != nil {
		panic(err)
	}
}
