package repositories

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func OpenPostgreSQLRepository() (Repository, error) {
	db, err := sql.Open("pgx", "su:su@/profiles")
	if err != nil {
		log.Println(err)
	}
	return &repository{db: db}, nil
}

type Repository interface {
	createUser(user User, login string) error
	updateUser(user User, login string) error
	deleteUser(login string) error
	getUserByLogin(login string) *User
	listing() (*[]Profile, error)
	closeDB() error
}
type ProfileRepository struct {
	db *sql.DB
}

func (r repository) InsertIntoDB(login string, user User) error {
	_, err := r.db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?)", login, user.Name, user.Surname,
		user.Password, user.Gender, user.HasAnyPets, user.Employed, user.Age)
	if err != nil {
		log.Errorf("There are some problems during insertion in table")
		log.Println(err)
	}
	return err
}
func (r repository) DeleteFromDB(login string) error {
	_, err := r.db.Query("DELETE * FROM products WHERE id=?", login)
	if err != nil {
		log.Errorf("There are some problems during deletion from table")
		log.Println(err)
	}
	return err
}
func (r repository) SelectFromDB(login string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE id=?", login)
	if err != nil {
		log.Errorf("There are some problems during selection in table")
		log.Println(err)
	}
	return rows, err
}
func (r repository) SelectAllFromDB() (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		log.Errorf("There are some problems during selection in table")
		log.Println(err)
	}
	return rows, err
}
func (r repository) UpdateDB(login string, user User) (*sql.Rows, error) {
	rows, err := r.db.Query("UPDATE users SET name=? surname=? password=? age=? gender=? hasanypets=? employed=? WHERE id=?",
		user.Name, user.Surname, user.Password, user.Age, user.Gender, user.HasAnyPets, user.Employed, login)
	if err != nil {
		log.Errorf("There are some problems during update table")
		log.Println(err)
	}
	return rows, err
}
func (r repository) closeDB() error {
	r.db.Close()
	log.Info("Database is closed")
	return nil
}
