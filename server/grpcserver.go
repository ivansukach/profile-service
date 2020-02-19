package server

import (
	"context"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/ivansukach/profile-service/repositories"
	"github.com/ivansukach/profile-service/service"
	log "github.com/sirupsen/logrus"
)

//Server ...
type Server struct {
	ps service.ProfileService
}

//Create ...
func (s *Server) Create(ctx context.Context, req *protocol.CreateRequest) (*protocol.SuccessResponse, error) {
	var user repositories.Profile
	user.Name = req.Profile.Name
	user.Surname = req.Profile.Surname
	user.Age = req.Profile.Age
	user.Password = req.Profile.Password
	user.Gender = req.Profile.Gender
	user.Employed = req.Profile.Employed
	user.HasAnyPets = req.Profile.HasAnyPets
	err := s.ps.CreateUser(user)
	if err != nil {
		log.Error(err)
		return &protocol.SuccessResponse{Success: false}, err
	}
	return &protocol.SuccessResponse{Success: true}, nil
}

func (s *Server) Update(ctx context.Context, req *protocol.UpdateRequest) (*protocol.SuccessResponse, error) {
	var user repositories.Profile
	user.Name = req.Profile.Name
	user.Surname = req.Profile.Surname
	user.Age = req.Profile.Age
	user.Password = req.Profile.Password
	user.Gender = req.Profile.Gender
	user.Employed = req.Profile.Employed
	user.HasAnyPets = req.Profile.HasAnyPets
	err := s.ps.UpdateUser(user)
	if err != nil {
		log.Error(err)
		return &protocol.SuccessResponse{Success: false}, err
	}
	return &protocol.SuccessResponse{Success: true}, nil
}
func (s *Server) Delete(ctx context.Context, req *protocol.DeleteRequest) (*protocol.SuccessResponse, error) {
	err := s.ps.DeleteUser(req.Login)
	if err != nil {
		log.Error(err)
		return &protocol.SuccessResponse{Success: false}, err
	}
	return &protocol.SuccessResponse{Success: true}, nil
}
func (s *Server) GetByLogin(ctx context.Context, req *protocol.GetByLoginRequest) (*protocol.GetByLoginResponse, error) {
	var resp protocol.GetByLoginResponse
	user, err := s.ps.GetUserByLogin(req.Login)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp.Profile.Name = user.Name
	resp.Profile.Surname = user.Surname
	resp.Profile.Age = user.Age
	resp.Profile.Password = user.Password
	resp.Profile.Gender = user.Gender
	resp.Profile.Employed = user.Employed
	resp.Profile.HasAnyPets = user.HasAnyPets
	return &protocol.GetByLoginResponse{Profile: resp.Profile}, nil
}
func (s *Server) Listing(ctx context.Context, req *protocol.ListingRequest) (*protocol.ListingResponse, error) {
	var resp protocol.ListingResponse
	users, err := s.ps.Listing()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for i := range users {
		resp.Profiles[i].Name = users[i].Name
		resp.Profiles[i].Surname = users[i].Surname
		resp.Profiles[i].Age = users[i].Age
		resp.Profiles[i].Password = users[i].Password
		resp.Profiles[i].Gender = users[i].Gender
		resp.Profiles[i].Employed = users[i].Employed
		resp.Profiles[i].HasAnyPets = users[i].HasAnyPets

	}
	return &protocol.ListingResponse{Profiles: resp.Profiles}, nil
}
