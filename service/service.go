package service

import (
	"fmt"
	"github.com/ivansukach/profile-service/repositories"
	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	r *repositories.ProfileRepository
}

func (ps ProfileService) CreateUser(user repositories.Profile) error {
	log.Info("createUser")
	rows, err := ps.r.SelectFromDB(user.Login)
	if err != nil || rows == nil {
		ps.r.InsertIntoDB(user)
		return nil
	}
	log.Warning("Already exists")
	return err
}
func (ps ProfileService) UpdateUser(user repositories.Profile) error {
	log.Info("updateUser")
	_, err := ps.r.SelectFromDB(user.Login)
	if err == nil {
		ps.r.UpdateDB(user)
		return nil
	}
	return err
}
func (ps ProfileService) DeleteUser(login string) error {
	log.Info("deleteUser")
	err := ps.r.DeleteFromDB(login)
	return err
}
func (ps ProfileService) GetUserByLogin(login string) (*repositories.Profile, error) {
	log.Info("getUserByLogin")
	rows, err := ps.r.SelectFromDB(login)
	if err != nil {
		log.Warning("There is not any record with value " + login + "in field Login")
		return nil, err
	}
	u := repositories.Profile{}
	err = rows.Scan(&u.Login, &u.Name, &u.Surname, &u.Password, &u.Gender, &u.HasAnyPets, &u.Employed, &u.Age)
	if err != nil {
		fmt.Println(err)
		log.Warning("Error when scanning DB")
	}
	return &u, err
}
func (ps ProfileService) Listing() ([]repositories.Profile, error) {
	log.Info("Listing")
	rows, err := ps.r.SelectAllFromDB()
	var list []repositories.Profile
	if err != nil {
		for rows.Next() {
			u := repositories.Profile{}
			err := rows.Scan(&u.Login, &u.Name, &u.Surname, &u.Password, &u.Gender, &u.HasAnyPets, &u.Employed, &u.Age)
			if err != nil {
				fmt.Println(err)
				log.Warning("Error when scanning DB")
				continue
			}
			list[len(list)-1] = u
		}
	} else {
		return nil, err
	}
	return list, nil
}
