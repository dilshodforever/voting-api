package handler

import pb "root/genprotos"

type Handler struct {
	Public    pb.PublicServiceClient
	Party     pb.PartyServiceClient
	Candidate pb.CandidateServiceClient
	Election  pb.ElectionServiceClient
	PublicVote   pb.PublicVoteServiceClient
}

func NewHandler(pu pb.PublicServiceClient, pa pb.PartyServiceClient, ca pb.CandidateServiceClient, el  pb.ElectionServiceClient, pub pb.PublicVoteServiceClient) *Handler {
	return &Handler{pu, pa, ca, el, pub}
}
