package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X float64 `json:"longitude"`
	Y float64 `json:"latitude"`
}

func (p *Point) parseBytes(bytes []byte) error {
	utf8String := string(bytes)
	utf8String = strings.Trim(utf8String, "()")
	coordinates := strings.Split(utf8String, ",")

	x, err := strconv.ParseFloat(coordinates[0], 64)
	y, err := strconv.ParseFloat(coordinates[1], 64)
	if err != nil { return err }

	p.X = x
	p.Y = y
	return nil
}

func (p Point) GetBytes() ([]byte, error) {
	utf8String := "(" + fmt.Sprintf("%f", p.X) + "," + fmt.Sprintf("%f", p.Y) + ")"
	return []byte(utf8String), nil
}

// Scan implements the Scanner interface.
func (p *Point) Scan(value interface{}) error {
	return p.parseBytes(value.([]byte))
}

// Value implements the driver Valuer interface.
func (p Point) Value() (driver.Value, error) {
	return p.GetBytes()
}