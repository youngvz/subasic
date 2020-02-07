package service

import (
	"subasic/models"
)

type UserService struct { }

func (dao UserService) Create(u *models.User) error {
	query := "INSERT INTO \"user\" (name, email, admin, \"groupId\") VALUES ($1, $2, $3, $4)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil { return err }

	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Email, false, u.GroupID)
	if err != nil { return err }
	return nil
}

func (dao UserService) GetById(i int)(models.User, error) {
	var user models.User
	query := "SELECT * FROM \"user\" WHERE \"userId\" = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return user, err }
	defer stmt.Close()

	rows, err := stmt.Query(i)
	if err != nil { return user, err }

	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Admin, &user.GroupID)
		if err != nil { return user, err }
	}
	return user, nil
}

func (dao UserService) GetAll() ([]models.User, error) {
	query := "SELECT * FROM \"user\""
	users := make([]models.User, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return users, err }
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil { return users, err }

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Admin, &user.GroupID)
		if err != nil { return nil, err }

		users = append(users, user)
	}
	return users, nil
}

func (dao UserService) Update(u *models.User) error {
	query := "UPDATE \"user\" SET name = $1, email = $2, \"groupId\" = $3 WHERE \"userId\" = $4"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil { return err }

	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Email, u.GroupID, u.UserID)
	if err != nil { return err }
	return nil
}

func (dao UserService) Delete(i int) error {
	query := "DELETE from \"user\" WHERE \"userId\" = $1"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil { return err }

	defer stmt.Close()
	_, err = stmt.Exec(i)
	if err != nil { return err }
	return nil
}
