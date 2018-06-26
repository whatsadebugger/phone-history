package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddressController is the controller responsible for all address requests
type AddressController struct {
	App *Application
}

func (ac *AddressController) createAddress(c *gin.Context) {
	ad := address{}

	if err := c.ShouldBindJSON(&ad); err != nil {
		publicError(c, 400, err)
	} else if err = ac.App.Database.Save(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
}

func (ac *AddressController) getAddress(c *gin.Context) {
	id := c.Param("id")
	var ad address
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if err := ac.App.Database.One("ID", aid, &ad); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
	} else if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
}

func (ac *AddressController) getEveryAddress(c *gin.Context) {
	var book []address
	if err := ac.App.Database.All(&book); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
	} else if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, book)
	}
}

func (ac *AddressController) updateAddress(c *gin.Context) {
	var ad address

	if err := c.ShouldBindJSON(&ad); err != nil {
		publicError(c, 400, err)
	} else if err = ac.App.Database.Update(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, ad)
	}
	fmt.Println(ad)
}

func (ac *AddressController) deleteAddress(c *gin.Context) {
	id := c.Param("id")

	if ad, err := ac.App.findAddress(id); err == storm.ErrNotFound {
		c.AbortWithError(404, errors.New("address entry not found"))
	} else if err != nil {
		c.AbortWithError(500, err)
	} else if err = ac.App.Database.DeleteStruct(&ad); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.String(200, "Success")
	}
}

func (ac *AddressController) importAddressBook(c *gin.Context) {

	body := c.Request.Body

	var addressbook []address

	lines, err := csv.NewReader(body).ReadAll()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	for _, line := range lines {
		addressbook = append(addressbook,
			address{FirstName: line[0],
				LastName: line[1],
				Email:    line[2],
				Phone:    line[3],
			})
	}
	count := 0
	for _, ad := range addressbook {
		err = ac.App.Database.Save(&ad)
		count++
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	c.String(200, fmt.Sprintf("imported %d records", count))

}

func (ac *AddressController) exportAddressBook(c *gin.Context) {
	var book []address
	if err := ac.App.Database.All(&book); err == storm.ErrNotFound {
		publicError(c, 404, storm.ErrNotFound)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	header := c.Writer.Header()
	header["Content-type"] = []string{"text/csv"}
	header["Content-Disposition"] = []string{"attachment; filename=backup.csv"}

	wr := csv.NewWriter(c.Writer)

	for _, v := range book {
		if err := wr.Write([]string{v.FirstName, v.LastName, v.Email, v.Phone}); err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	wr.Flush()
}
