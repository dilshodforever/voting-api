package handler
import pb "root/genprotos"
type Handler struct{
	Public pb.PublicServiceClient
	Party pb.PartyServiceClient
}

func NewHandler(pu pb.PublicServiceClient, pa pb.PartyServiceClient) *Handler{
	return &Handler{pu, pa}
}