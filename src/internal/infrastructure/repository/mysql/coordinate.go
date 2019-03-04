package mysql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type coordinateRepository struct {
	db *sql.DB
}

func (c *coordinateRepository) FindById(
	ctx context.Context, coordinateId int) (domain.Coordinate, error) {

	var coordinate domain.Coordinate
	var createdAt time.Time

	err := c.db.QueryRowContext(ctx,
		`SELECT
			id, image_name, coordinates.user_id, created_at, COUNT(coordinate_id)
			FROM coordinates LEFT OUTER JOIN favorites
				ON coordinates.id = favorites.coordinate_id
			WHERE coordinates.id = ?
			GROUP BY id`, coordinateId).Scan(
		&coordinate.Id,
		&coordinate.ImageName,
		&coordinate.UserId,
		&createdAt,
		&coordinate.Favorites)
	if err != nil {
		return domain.Coordinate{}, err
	}

	coordinate.CreatedAt = domain.JsonTime{Data: createdAt}

	return coordinate, nil
}

func (c *coordinateRepository) GetAtRandom(
	ctx context.Context) (domain.Coordinate, error) {

	var coordinate domain.Coordinate
	var createdAt time.Time

	// SELECT id ... RAND() ... AS rand -> SELECT * ... where id = rand.id
	err := c.db.QueryRowContext(ctx,
		`SELECT
			id, image_name, coordinates.user_id, created_at, COUNT(coordinate_id)
			FROM coordinates LEFT OUTER JOIN favorites
				ON coordinates.id = favorites.coordinate_id
			GROUP BY id ORDER BY RAND() LIMIT 0, 1`).Scan(
		&coordinate.Id,
		&coordinate.ImageName,
		&coordinate.UserId,
		&createdAt,
		&coordinate.Favorites)
	if err != nil {
		return domain.Coordinate{}, err
	}

	coordinate.CreatedAt = domain.JsonTime{Data: createdAt}

	return coordinate, nil
}

func (c *coordinateRepository) FindFavoriteByCoordinateIdAndUserId(
	ctx context.Context, coordinateId int, uid string) (domain.Favorite, error) {

	var favorite domain.Favorite
	err := c.db.QueryRowContext(ctx,
		`SELECT coordinate_id, user_id FROM favorites
			WHERE coordinate_id = ? AND user_id = ?`,
		coordinateId, uid).Scan(
		&favorite.CoordinateId,
		&favorite.UserId)
	if err != nil {
		return domain.Favorite{}, err
	}

	return favorite, nil
}

func (c *coordinateRepository) FindByUserId(
	ctx context.Context, uid string) ([]domain.Coordinate, error) {

	rows, err := c.db.QueryContext(ctx,
		`SELECT
			id, image_name, coordinates.user_id, created_at, COUNT(coordinate_id)
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
			&coordinate.ImageName,
			&coordinate.UserId,
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
	ctx context.Context, coordinate domain.Coordinate) (int64, error) {

	result, err := c.db.ExecContext(ctx,
		`INSERT INTO coordinates(image_name, user_id, created_at)
			VALUES(?, ?, ?)`,
		coordinate.ImageName,
		coordinate.UserId,
		coordinate.CreatedAt.Data)
	if err != nil {
		return -1, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func (c *coordinateRepository) Delete(
	ctx context.Context, coordinateId int) (err error) {

	_, err = c.db.ExecContext(
		ctx, "DELETE FROM coordinates WHERE id = ?", coordinateId)
	return
}

func NewCoordinateRepository(db *sql.DB) *coordinateRepository {
	return &coordinateRepository{db}
}
