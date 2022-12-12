package storage

import (
	"database/sql"
	"fmt"

	"github.com/asadbekGo/golang_crud/models"
)

func Create(db *sql.DB, user models.User) (string, error) {

	var (
		id    string
		query string
	)

	query = `
		INSERT INTO 
			users (first_name, last_name)
		VALUES ( $1, $2 )
		RETURNING id
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (models.User, error) {

	var (
		user  models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
		WHERE id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetList(db *sql.DB) ([]models.User, error) {

	var (
		users []models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Update(db *sql.DB, user models.User, id string) (models.User, error) { 

	
	_, err := db.Query(`
		UPDATE 
			users 
		SET 
			first_name = $2, 
			last_name = $3
		WHERE 
			id = $1
	`, id, user.FirstName, user.LastName) 
	
	if err != nil {
		return models.User{}, err
	} 

	return models.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil

}

func Delete(db *sql.DB, id string) error {

	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func Patch(db *sql.DB, user models.User, id string) (models.User, error) {

	var (
		query      string
		args       = make(map[string]string)
		first_name string
		last_name  string
	)
	fmt.Println("ID++=> ", id)

	err := db.QueryRow(`select first_name, last_name from users where id = $1`, id).Scan(&first_name, &last_name)
	args["first_name"] = first_name
	args["last_name"] = last_name
	if err != nil {
		return models.User{}, err
	}
	if user.FirstName != "" {
		args["first_name"] = user.FirstName
	}

	if user.LastName != "" {
		args["last_name"] = user.LastName
	}
	query = `
		UPDATE 
			users 
		SET 
			first_name = $2, 
			last_name = $3
		WHERE 
			id = $1
	`

	_, err = db.Exec(
		query,
		id,
		args["first_name"],
		args["last_name"],
	)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
