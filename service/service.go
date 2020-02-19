package service

import (
	"fmt"
	"github.com/ivansukach/profile-service/repositories"
	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	r *repositories.ProfileRepository
}
type User struct {
	Password   string
	Name       string
	Surname    string
	Age        int32
	Gender     bool
	HasAnyPets bool
	Employed   bool
}
type Profile struct {
	Login string
	User  User
}

func (ps ProfileService) createUser(user User, login string) error {
	log.Info("createUser")
	rows, err := ps.r.SelectFromDB(login)
	if err != nil || rows == nil {
		ps.r.InsertIntoDB(login, user)
		return nil
	}
	log.Warning("Already exists")
	return err
}
func (ps ProfileService) updateUser(user User, login string) error {
	log.Info("updateUser")
	_, err := ps.r.SelectFromDB(login)
	if err == nil {
		ps.r.UpdateDB(login, user)
		return nil
	}
	return err
}
func (ps ProfileService) deleteUser(login string) error {
	log.Info("deleteUser")
	err := ps.r.DeleteFromDB(login)
	return err
}
func (ps ProfileService) getUserByLogin(login string) *User {
	log.Info("getUserByLogin")
	rows, err := ps.r.SelectFromDB(login)
	if err != nil {
		log.Warning("There is not any record with value " + login + "in field Login")
		return nil
	}
	u := User{}
	err = rows.Scan(&login, &u.Name, &u.Surname, &u.Password, &u.Gender, &u.HasAnyPets, &u.Employed, &u.Age)
	if err != nil {
		fmt.Println(err)
		log.Warning("Error when scanning DB")
	}
	return &u
}
func (ps ProfileService) listing() (*[]Profile, error) {
	log.Info("Listing")
	rows, err := ps.r.SelectAllFromDB()
	var list []Profile
	if err != nil {
		for rows.Next() {
			u := User{}
			login := ""
			err := rows.Scan(&login, &u.Name, &u.Surname, &u.Password, &u.Gender, &u.HasAnyPets, &u.Employed, &u.Age)
			if err != nil {
				fmt.Println(err)
				log.Warning("Error when scanning DB")
				continue
			}
			list[len(list)-1].User = u
			list[len(list)-1].Login = login
		}
	} else {
		return nil, err
	}
	return &list, nil
}
