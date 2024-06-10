package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateElection handles the creation of a new Election
// @Summary Create Election
// @Description Create page
// @Tags Election
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.Election     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /election/create [post]
func (h *Handler) CreateElection(ctx *gin.Context){
		arr:=pb.Election{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.Election.CreateElection(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Success!!!")
}

// UpdateElection handles the creation of a new Election
// @Summary Update Election
// @Description Update page
// @Tags Election
// @Accept  json
// @Produce  json
// @Param     id path string true "Election ID"
// @Param   Update     body    pb.Election     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /election/update/{id} [put]
func (h *Handler) UpdateElection(ctx *gin.Context){
	arr:=pb.Election{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.Election.UpdateElection(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeleteElection handles the creation of a new Election
// @Summary Delete Election
// @Description Delete page
// @Tags Election
// @Accept  json
// @Produce  json
// @Param     id path string true "Election ID"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /election/delete/{id} [delete]
func (h *Handler) DeleteElection(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.Election.DeleteElection(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllElection handles the creation of a new Election
// @Summary GetAll Election
// @Description GetAll page
// @Tags Election
// @Accept  json
// @Produce  json
// @Success 200 {object} pb.GetAllElection   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /election/getall [get]
func (h *Handler) GetAllElection(ctx *gin.Context){
	arr:=&pb.Election{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	res, err:=h.Election.GetAllElections(ctx, arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdElection handles the creation of a new Election
// @Summary GetById Election
// @Description GetById page
// @Tags Election
// @Accept  json
// @Produce  json
// @Param     id path string true "Election ID"
// @Success 200 {object} pb.Election   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /election/getbyid/{id} [get]
func (h *Handler) GetByIdElection(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.Election.GetByIdElection(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
