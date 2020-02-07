package service

import (
	"fmt"
	"subasic/models"
)

type GroupService struct { }

func (dao GroupService) Create(g *models.Group) error {
	query := "INSERT INTO \"group\"(name, \"userId\", \"photoId\") VALUES($1, $2, $3)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(g.Name, g.UserID, g.PhotoID)
	if err != nil {
		return err
	}
	return nil
}

func (dao GroupService) GetById(i int)(models.Group, error) {
	var group models.Group
	query := "SELECT * FROM \"group\" WHERE \"groupId\" = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return group, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(i)
	if err != nil {
		return group, err
	}

	for rows.Next() {
		err = rows.Scan(&group.GroupID, &group.Name, &group.UserID, &group.PhotoID)
		if err != nil {
			fmt.Println(err)
			return group, err
		}
	}

	return group, nil
}

func (dao GroupService) GetAll() ([]models.Group, error) {
	query := "SELECT * FROM \"group\""
	groups := make([]models.Group, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return groups, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return groups, err
	}

	for rows.Next() {
		var row models.Group
		err = rows.Scan(&row.GroupID, &row.Name, &row.UserID, &row.PhotoID)
		if err != nil {
			return nil, err
		}

		groups = append(groups, row)
	}

	return groups, nil
}

func (dao GroupService) Update(g *models.Group) error {
	query := "UPDATE \"group\" SET \"name\" = $1, \"userId\" = $2, \"photoId\" = $3, WHERE \"groupId\" = $4"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(g.Name, g.UserID, g.PhotoID, g.GroupID)
	if err != nil {
		return err
	}
	return nil
}

func (dao GroupService) Delete(i int) error {
	query := "DELETE from \"group\" WHERE id = $1"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(i)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil
}
