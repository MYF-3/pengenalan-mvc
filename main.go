package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

var client *db.Cllient
var ctx context.Context

func init (){
	ctx = context.Backround()
	conf := &firebase.Config{
		DatabaseURL: "https://{URL_FIREBASE}.firebase.io.com",
	}
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")
	
	app,err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
}

var data []Antrian

type Antrian struct {
	Id String `json:"id"`
	Status bool `json:"status"`
}

func main() {
	router := gin.Default()
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/: idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/: idAntrian/delete", DeleteAntrianHandler)
	router.Run(":8080")
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

func AddAntrian() (bool,error){
	_,_,dataAntrian := getAntrian()
	var Id string
	var antrianRef *db.Ref
	Ref := client.NewRef("antrian")

	if dataAntrian == nil{
		Id = fmt.Sprint("B-0")
		antrianRef = ref.Child("0")
	}else {
		Id = fmt.Sprint("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		Id: Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false,err
	}
	return true,nil
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

func GetAntrian() (bool,error,[]map[string]interface{}){
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln ("Error reading from database:", err)
		return false, err, nil
	}
	return true,nil,data
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

func updateAntrian(idAntrian string) (bool,error) {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian(
		Id: idAntrian,
		Status: true,
	)
	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false,err
	}
	return true,nil
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

func DeleteAntrian (idAntrian string) (bool,error){

	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child (id[1])
	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false,err
	}
	return true,nil
}