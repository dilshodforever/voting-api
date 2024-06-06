package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

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


func (h *Handler) DeletePublic(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.Public.DeletePublic(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}



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


func (h *Handler) GetbyIdPublic(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.Public.GetByIdPublic(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
