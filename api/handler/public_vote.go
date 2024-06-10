package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// CreatePublic_Vote handles the creation of a new Public_Vote
// @Summary Create Public_Vote
// @Description Create page
// @Tags PublicVote
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.PublicVote     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /public_vote/create [post]
func (h *Handler) CreatePublicVote(ctx *gin.Context){
		arr:=pb.PublicVote{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.PublicVote.CreatePublicVote(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Success!!!")
}

// UpdatePublicVote handles the creation of a new PublicVOTE
// @Summary Update PublicVOTE
// @Description Update page
// @Tags PublicVote
// @Accept  json
// @Produce  json
// @Param     id path string true "PublicVOTE ID"
// @Param   Update     body    pb.PublicVote     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /publicvote/update/{id} [put]
func (h *Handler) UpdatePublicVote(ctx *gin.Context){
	arr:=pb.PublicVote{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.PublicVote.UpdatePublicVote(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeletePublicVote handles the creation of a new PublicVote
// @Summary Delete PublicVote
// @Description Delete page
// @Tags PublicVote
// @Accept  json
// @Produce  json
// @Param     id path string true "PublicVote ID"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /publicvote/delete/{id} [delete]
func (h *Handler) DeletePublicVote(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.PublicVote.DeletePublicVote(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllPublicVote handles the creation of a new PublicVote
// @Summary GetAll PublicVote
// @Description GetAll page
// @Tags PublicVote
// @Accept  json
// @Produce  json
// @Success 200 {object} pb.GetAllPublicVote   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /publicvote/getall [get]
func (h *Handler) GetAllPublicVote(ctx *gin.Context){
	arr:=&pb.PublicVote{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	res, err:=h.PublicVote.GetAllPublucVotes(ctx, arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdPublicVote handles the creation of a new PublicVote
// @Summary GetById PublicVote
// @Description GetById page
// @Tags PublicVote
// @Accept  json
// @Produce  json
// @Param     id path string true "PublicVote ID"
// @Success 200 {object} pb.PublicVote   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /PublicVote/getbyid/{id} [get]
func (h *Handler) GetByIdPublicVote(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.PublicVote.GetByIdPublicVote(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
