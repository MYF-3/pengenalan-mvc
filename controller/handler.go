package controller

import {
	"https://github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-MVC-Golang-Concept/"
	"github.com/gin-gonic/gin" 
	"net/http"
}

func AddAntrianHandler(c *gin.Context){
	flag,err :=addAntrian()
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	}else{
		c.JSON(http.StatusBadRequest,map[string]interface{
			"status":"failed",
			"error":err,
		})
	}
}


func GetAntrianHandler(c *gin.Context){

	flag,err,resp :=getAntrian()
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
			"error":resp,
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag,err := updateAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"ststus":"failed",
			"error":"err",
		})
	}
}

func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag,err := deleteAntrian(idAntrian)
if flag {
	c.JSON(http.StatusOK,map[string]interface{}{
		"status":"success",
	})
}else {
	c.JSON(http.StatusBadRequest,map[string]interface{}{
		"status":"failed",
		"error":"err",
	})
}
}

func PageAntrianHandler(c *gin.Context) {
	flag,err,result := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len (result)>0{
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	}else {
		c.JSON(http.StatusBadRequest), map[string]interface{}{
			"status":"failed",
			"error":err,
		}
	}
}