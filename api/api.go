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
	p.PUT("/update/:id", h.UpdatePublic)
	p.DELETE("/delete/:id", h.DeletePublic)
	p.GET("/getall", h.GetAllPublic)
	p.GET("/getbyid/:id", h.GetbyIdPublic)
	p.GET("/cheak/:id", h.CheakPublic)
	
	pa:=r.Group("/party")
	pa.POST("/create", h.CreateParty)
	pa.PUT("/update/:id", h.UpdateParty)
	pa.DELETE("/delete/:id", h.DeleteParty)
	pa.GET("/getall", h.GetAllParty)
	pa.GET("/getbyid/:id", h.GetbyIdParty)

	ca:=r.Group("/candidate")
	ca.POST("/create", h.CreateCandidate)
	ca.PUT("/update/:id", h.UpdateCandidate)
	ca.DELETE("/delete/:id", h.DeleteCandidate)
	ca.GET("/getall", h.GetAllCandidate)
	ca.GET("/getbyid/:id", h.GetByIdCandidate)
	return r
}