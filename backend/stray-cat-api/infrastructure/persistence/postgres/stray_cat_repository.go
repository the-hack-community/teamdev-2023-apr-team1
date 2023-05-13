// infrastructure/persistence/postgres/stray_cat_repository.go
package postgres

import (
	"database/sql"
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type StrayCatRepository struct {
	DB *sql.DB
}

func NewStrayCatRepository(db *sql.DB) repository.StrayCatRepository {
	return &StrayCatRepository{DB: db}
}

func (r *StrayCatRepository) FindAll() ([]*model.StrayCat, error) {
	rows, err := r.DB.Query(`SELECT sc.cat_id, sc.user_id, sc.photo_data, sc.capture_date_time, l.lat, l.long, sc.name, sc.features, sc.condition 
							  FROM stray_cats sc JOIN locations l ON sc.location_id = l.location_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var strayCats []*model.StrayCat

	for rows.Next() {
		var cat model.StrayCat
		var location model.Location

		err := rows.Scan(&cat.CatID, &cat.UserID, &cat.PhotoData, &cat.CaptureDateTime, &location.Lat, &location.Long, &cat.Name, &cat.Features, &cat.Condition)
		if err != nil {
			return nil, err
		}

		cat.Location = location
		strayCats = append(strayCats, &cat)
	}

	return strayCats, nil
}

func (r *StrayCatRepository) FindByID(catID int) (*model.StrayCat, error) {
	row := r.DB.QueryRow(`SELECT sc.cat_id, sc.user_id, sc.photo_data, sc.capture_date_time, l.lat, l.long, sc.name, sc.features, sc.condition
						  FROM stray_cats sc JOIN locations l ON sc.location_id = l.location_id WHERE sc.cat_id = $1`, catID)

	var cat model.StrayCat
	var location model.Location

	err := row.Scan(&cat.CatID, &cat.UserID, &cat.PhotoData, &cat.CaptureDateTime, &location.Lat, &location.Long, &cat.Name, &cat.Features, &cat.Condition)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	cat.Location = location
	return &cat, nil
}

func (r *StrayCatRepository) Store(strayCat *model.StrayCat) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	locationID := 0
	err = tx.QueryRow("INSERT INTO locations (lat, long) VALUES ($1, $2) RETURNING location_id", strayCat.Location.Lat, strayCat.Location.Long).Scan(&locationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO stray_cats (user_id, photo_data, capture_date_time, location_id, name, features, condition) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		strayCat.UserID, strayCat.PhotoData, strayCat.CaptureDateTime, locationID, strayCat.Name, strayCat.Features, strayCat.Condition)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *StrayCatRepository) Update(strayCat *model.StrayCat) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	locationID := 0
	err = tx.QueryRow("SELECT location_id FROM stray_cats WHERE cat_id = $1", strayCat.CatID).Scan(&locationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE locations SET lat = $1, long = $2 WHERE location_id = $3", strayCat.Location.Lat, strayCat.Location.Long, locationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE stray_cats SET user_id = $1, photo_data = $2, capture_date_time = $3, name = $4, features = $5, condition = $6 WHERE cat_id = $7",
		strayCat.UserID, strayCat.PhotoData, strayCat.CaptureDateTime, strayCat.Name, strayCat.Features, strayCat.Condition, strayCat.CatID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *StrayCatRepository) Delete(catID int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	locationID := 0
	err = tx.QueryRow("SELECT location_id FROM stray_cats WHERE cat_id = $1", catID).Scan(&locationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM stray_cats WHERE cat_id = $1", catID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM locations WHERE location_id = $1", locationID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
