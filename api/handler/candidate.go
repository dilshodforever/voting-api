package handler

import (
	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateCandidate handles the creation of a new Candidate
// @Summary Create Candidate
// @Description Create page
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Param   Create     body    pb.Candidate     true        "Create"
// @Success 200 {string} string  "Create Successful"
// @Failure 401 {string} string  "Error while Created"
// @Router /candidate/create [post]
func (h *Handler) CreateCandidate(ctx *gin.Context){
		arr:=pb.Candidate{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.Candidate.CreateCandidate(ctx, &arr)
		if err!=nil{
			panic(err)
		}
		ctx.JSON(200, "Success!!!")
}

// UpdateCandidate handles the creation of a new Candidate
// @Summary Update Candidate
// @Description Update page
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Param     id path string true "Candidate ID"
// @Param   Update     body    pb.Candidate     true        "Update"
// @Success 200 {string} string  "Update Successful"
// @Failure 401 {string} string  "Error while created"
// @Router /candidate/update/{id} [put]
func (h *Handler) UpdateCandidate(ctx *gin.Context){
	arr:=pb.Candidate{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.Candidate.UpdateCandidate(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeleteCandidate handles the creation of a new Candidate
// @Summary Delete Candidate
// @Description Delete page
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Param     id path string true "Candidate ID"
// @Success 200 {string} string  "Delete Successful"
// @Failure 401 {string} string  "Error while Deleted"
// @Router /candidate/delete/{id} [delete]
func (h *Handler) DeleteCandidate(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.Candidate.DeleteCandidate(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllCandidate handles the creation of a new Candidate
// @Summary GetAll Candidate
// @Description GetAll page
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} pb.GetAllCandidate   "GetAll Successful"
// @Failure 401 {string} string  "Error while GetAlld"
// @Router /candidate/getall [get]
func (h *Handler) GetAllCandidate(ctx *gin.Context){
	can:=&pb.Candidate{}
	err:=ctx.BindJSON(&can)
	if err!=nil{
		panic(err)
	}
	res, err:=h.Candidate.GetAllCandidates(ctx, can)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdCandidate handles the creation of a new Candidate
// @Summary GetById Candidate
// @Description GetById page
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Param     id path string true "Candidate ID"
// @Success 200 {object} pb.Candidate   "GetById Successful"
// @Failure 401 {string} string "Error while GetByIdd"
// @Router /candidate/getbyid/{id} [get]
func (h *Handler) GetByIdCandidate(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.Candidate.GetByIdCandidate(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}
