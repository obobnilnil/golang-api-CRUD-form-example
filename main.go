package main

import (
	"mars_git/database"
	"mars_git/handler"
	"mars_git/repository"
	"mars_git/service"
	"mars_git/utility"

	"github.com/gin-gonic/gin"
)

func main() {
	//db := database.Mariadb() // not connected
	db := database.Postgresql()
	defer db.Close()
	utility.CountTables(db)
	r := repository.NewRepositoryAdapter(db)
	// fmt.Println(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHanerhandlerAdapter(s)

	router := gin.Default()
	router.POST("/api/createForm", h.LeaderCreateFormHandler)
	router.POST("/api/submitForm", h.SubmitFormHandler)

	err := router.Run(":5599")
	if err != nil {
		panic(err.Error())
	}
}
