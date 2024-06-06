package api

import (
	"root/api/handler"
	_ "root/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
func NewGin(h *handler.Handler)*gin.Engine{
	r:=gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	p:=r.Group("/public")
	p.POST("/create", h.CreatePublic)
	p.PUT("/update", h.UpdatePublic)
	p.DELETE("/delete", h.DeletePublic)
	p.GET("/getall", h.GetAllPublic)
	p.GET("/getbyid", h.GetbyIdPublic)
	
	pa:=r.Group("/party")
	pa.POST("/create", h.CreateParty)
	pa.PUT("/update", h.UpdateParty)
	pa.DELETE("/delete", h.DeleteParty)
	pa.GET("/getall", h.GetAllParty)
	pa.GET("/getbyid", h.GetbyIdParty)
	return r
}