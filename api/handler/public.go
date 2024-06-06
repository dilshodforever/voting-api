package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// CreatePublic handles the creation of a new Public
// @Summary Create Public
// @Description Create page
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.Public     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /public/create [post]
func (h *Handler) CreatePublic(ctx *gin.Context){
		arr:=pb.Public{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.Public.CreatePublic(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Sucsess!!!")
}

// UpdatePublic handles the creation of a new Public
// @Summary Update Public
// @Description Update page
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   Update     body    pb.Public     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /public/update [put]
func (h *Handler) UpdatePublic(ctx *gin.Context){
	arr:=pb.Public{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.Public.UpdatePublic(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}

// DeletePublic handles the creation of a new Public
// @Summary Delete Public
// @Description Delete page
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   Delete     body    pb.ById    true        "Delete"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /public/delete [delete]
func (h *Handler) DeletePublic(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.Public.DeletePublic(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}


// GetAllPublic handles the creation of a new Public
// @Summary GetAll Public
// @Description GetAll page
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   GetAll     body    pb.Void     true        "GetAll"
// @Success 200 {object} pb.GetAllPublic   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /public/getall [get]
func (h *Handler) GetAllPublic(ctx *gin.Context){
	arr:=pb.Public{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	res, err:=h.Public.GetAllPublics(ctx, &pb.Void{})
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdPublic handles the creation of a new Public
// @Summary GetById Public
// @Description GetById page
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   GetById     body    pb.ById     true        "GetById"
// @Success 200 {object} pb.Public   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /public/GetById [get]
func (h *Handler) GetbyIdPublic(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.Public.GetByIdPublic(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
