package service

import (
	"github.com/ivansukach/profile-service/repositories"
)

type ProfileService struct {
	r repositories.Repository
}

func (ps *ProfileService) Create(profile repositories.Profile) error {
	return ps.r.InsertIntoDB(profile)
}
func (ps *ProfileService) Update(profile repositories.Profile) error {
	return ps.r.UpdateDB(profile)
}
func (ps *ProfileService) GetByLogin(login string) (repositories.Profile, error) {
	return ps.r.SelectFromDB(login)
}
func (ps *ProfileService) Delete(login string) error {
	return ps.r.DeleteFromDB(login)
}
func (ps *ProfileService) Listing() ([]repositories.Profile, error) {
	return ps.r.SelectAllFromDB()
}
func New(repo repositories.Repository) *ProfileService {
	return &ProfileService{r: repo}
}
