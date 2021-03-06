package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Profile struct {
	Login      string
	Password   string
	Name       string
	Surname    string
	Age        int32
	Gender     bool
	HasAnyPets bool
	Employed   bool
}

func OpenPostgreSQLRepository() (Repository, error) {
	db, err := sql.Open("postgres", "su:su@/profiles")
	if err != nil {
		log.Println(err)
	}
	return &ProfileRepository{db: db}, nil
}

type Repository interface {
	InsertIntoDB(user Profile) error
	DeleteFromDB(login string) error
	SelectFromDB(login string) (Profile, error)
	SelectAllFromDB() ([]Profile, error)
	UpdateDB(user Profile) error
	CloseDB() error
}
type ProfileRepository struct {
	db *sql.DB
}

func (r *ProfileRepository) InsertIntoDB(user Profile) error {
	_, err := r.db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Login, user.Name, user.Surname,
		user.Password, user.Gender, user.HasAnyPets, user.Employed, user.Age)
	if err != nil {
		log.Error(err)
	}
	return err
}
func (r *ProfileRepository) DeleteFromDB(login string) error {
	_, err := r.db.Query("DELETE * FROM products WHERE id=?", login)
	if err != nil {
		log.Error(err)
	}
	return err
}
func (r *ProfileRepository) SelectFromDB(login string) (Profile, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE id=?", login)
	if err != nil {
		log.Error(err)
	}
	p := new(Profile)
	for rows.Next() {
		err := rows.Scan(p.Login, p.Name, p.Surname, p.Password, p.Gender, p.HasAnyPets, p.Employed, p.Age)
		if err != nil {
			log.Error(err)
		}
	}
	return *p, err
}
func (r *ProfileRepository) SelectAllFromDB() ([]Profile, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		log.Error(err)
	}
	profiles := make([]Profile, 0)
	p := new(Profile)
	for rows.Next() {
		err := rows.Scan(p.Login, p.Name, p.Surname, p.Password, p.Gender, p.HasAnyPets, p.Employed, p.Age)
		if err != nil {
			log.Error(err)
		}
		profiles = append(profiles, *p)
	}
	return profiles, err
}
func (r *ProfileRepository) UpdateDB(user Profile) error {
	_, err := r.db.Query("UPDATE users SET name=? surname=? password=? age=? gender=? hasanypets=? employed=? WHERE id=?",
		user.Name, user.Surname, user.Password, user.Age, user.Gender, user.HasAnyPets, user.Employed, user.Login)
	if err != nil {
		log.Error(err)
	}
	return err
}
func (r *ProfileRepository) CloseDB() error {
	r.db.Close()
	log.Info("Database is closed")
	return nil
}
