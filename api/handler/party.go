package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateParty handles the creation of a new party
// @Summary Create Party
// @Description Create page
// @Tags Party
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.Party     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /party/create [post]


func (h *Handler) CreateParty(ctx *gin.Context){
		arr:=pb.Party{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.Party.CreateParty(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Sucsess!!!")
}

// UpdateParty handles the creation of a new party
// @Summary Update Party
// @Description Update page
// @Tags Party
// @Accept  json
// @Produce  json
// @Param   Update     body    pb.Party     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /party/update [put]


func (h *Handler) UpdateParty(ctx *gin.Context){
	arr:=pb.Party{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.Party.UpdateParty(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}
// DeleteParty handles the creation of a new party
// @Summary Delete Party
// @Description Delete page
// @Tags Party
// @Accept  json
// @Produce  json
// @Param   Delete     body    pb.ById    true        "Delete"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /party/delete [delete]


func (h *Handler) DeleteParty(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.Party.DeleteParty(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}

// GetAllParty handles the creation of a new party
// @Summary GetAll Party
// @Description GetAll page
// @Tags Party
// @Accept  json
// @Produce  json
// @Param   GetAll     body    pb.Void     true        "GetAll"
// @Success 200 {object} pb.GetAllParty   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /party/getall [get]


func (h *Handler) GetAllParty(ctx *gin.Context){
	res, err:=h.Party.GetAllPartys(ctx, &pb.Void{})
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
// GetByIdParty handles the creation of a new party
// @Summary GetById Party
// @Description GetById page
// @Tags Party
// @Accept  json
// @Produce  json
// @Param   GetById     body    pb.ById     true        "GetById"
// @Success 200 {object} pb.Party   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /party/GetById [get]


func (h *Handler) GetbyIdParty(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.Party.GetByIdParty(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}