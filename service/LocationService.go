package service

import (
	"subasic/models"
)

type LocationService struct { }

func (dao LocationService) Create(l *models.Location) error {
	query := "INSERT INTO \"location\" (address, coordinate, \"googleId\", name) VALUES ($1, $2, $3, $4)"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return err }
	defer stmt.Close()

	_, err = stmt.Exec(l.Address, l.Coordinate, l.GoogleId, l.Name)
	if err != nil { return err }
	return nil
}

func (dao LocationService) GetById(i int)(models.Location, error) {
	var location models.Location
	query := "SELECT * FROM \"location\" WHERE \"locationId\" = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return location, err }
	defer stmt.Close()

	rows, err := stmt.Query(i)
	if err != nil {
		return location, err
	}

	for rows.Next() {
		err = rows.Scan(&location.LocationId, &location.Address, &location.Coordinate, &location.GoogleId, &location.Name)
		if err != nil {
			return location, err
		}
	}
	return location, nil
}

func (dao LocationService) GetAll() ([]models.Location, error) {
	query := "SELECT * FROM \"location\""
	locations := make([]models.Location, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return locations, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return locations, err
	}

	for rows.Next() {
		var location models.Location
		err = rows.Scan(&location.LocationId, &location.Address, &location.Coordinate, &location.GoogleId, &location.Name)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}
	return locations, nil
}

func (dao LocationService) Update(l *models.Location) error {
	query := "UPDATE \"location\" SET \"address\" = $1, \"coordinate\" = $2, \"googleId\" = $3, \"name\" = $4 WHERE \"locationId\" = $5"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return err }
	defer stmt.Close()

	_, err = stmt.Exec(l.Address, l.Coordinate, l.GoogleId, l.Name, l.LocationId)
	if err != nil { return err }
	return nil
}

func (dao LocationService) Delete(i int) error {
	query := "DELETE from \"location\" WHERE \"locationId\" = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil { return err }
	defer stmt.Close()

	_, err = stmt.Exec(i)
	if err != nil { return err }
	return nil
}
