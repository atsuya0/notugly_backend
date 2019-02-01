package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/tayusa/notugly_backend/app/domain"
)

type coordinateRepository struct {
	conn *sql.DB
}

func (c *coordinateRepository) FindById(
	coordinateId int) (domain.Coordinate, error) {

	var coordinate domain.Coordinate
	var createdAt time.Time

	err := c.conn.QueryRow(
		`SELECT id, image, created_at, count(id)
			FROM coordinates LEFT OUTER JOIN favorites
				ON coordinates.id = favorites.coordinate_id
			WHERE coordinates.id = ?
			GROUP BY id`, coordinateId).Scan(
		&coordinate.Id,
		&coordinate.Image,
		&createdAt,
		&coordinate.Favorites)
	if err != nil {
		return domain.Coordinate{}, err
	}

	coordinate.CreatedAt = domain.JsonTime{Data: createdAt}

	return coordinate, nil
}

func (c *coordinateRepository) GetRandom() (domain.Coordinate, error) {
	var coordinate domain.Coordinate
	var createdAt time.Time

	// SELECT id ... RAND() ... AS rand -> SELECT * ... where id = rand.id
	err := c.conn.QueryRow(
		`SELECT id, image, created_at, count(id)
			FROM coordinates LEFT OUTER JOIN favorites
				ON coordinates.id = favorites.coordinate_id
			GROUP BY id ORDER BY RAND() LIMIT 0, 1`).Scan(
		&coordinate.Id,
		&coordinate.Image,
		&createdAt,
		&coordinate.Favorites)
	if err != nil {
		return domain.Coordinate{}, err
	}

	coordinate.CreatedAt = domain.JsonTime{Data: createdAt}

	return coordinate, nil
}

func (c *coordinateRepository) FindByUserId(
	uid string) ([]domain.Coordinate, error) {

	rows, err := c.conn.Query(
		`SELECT id, image, created_at, count(id)
			FROM coordinates LEFT OUTER JOIN favorites
				ON coordinates.id = favorites.coordinate_id
			WHERE coordinates.user_id = ?
			GROUP BY id`, uid)
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return []domain.Coordinate{}, err
	}

	coordinates := make([]domain.Coordinate, 0)
	var coordinate domain.Coordinate
	var createdAt time.Time

	for rows.Next() {
		err = rows.Scan(
			&coordinate.Id,
			&coordinate.Image,
			&createdAt,
			&coordinate.Favorites)
		if err != nil {
			return []domain.Coordinate{}, err
		}
		coordinate.CreatedAt = domain.JsonTime{Data: createdAt}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (c *coordinateRepository) Store(
	coordinate domain.Coordinate) (int64, error) {

	result, err := c.conn.Exec(
		"INSERT INTO coordinates(id, image, user_id, created_at) VALUES(?, ?, ?, ?)",
		coordinate.Id, coordinate.Image, coordinate.UserId, coordinate.CreatedAt.Data)
	if err != nil {
		return -1, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func (c *coordinateRepository) Delete(coordinateId int) (err error) {
	_, err = c.conn.Exec(
		"DELETE FROM coordinates WHERE id = ?", coordinateId)
	return
}

func NewCoordinateRepository(conn *sql.DB) *coordinateRepository {
	return &coordinateRepository{conn}
}
