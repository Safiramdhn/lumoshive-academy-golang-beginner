package repositories

import (
	"database/sql"
	"golang-beginner-19/models"
)

type MaterialRepositoryDB struct {
	DB *sql.DB
}

func NewMaterialRepositoryDB(db *sql.DB) *MaterialRepositoryDB {
	return &MaterialRepositoryDB{DB: db}
}

func (repo *MaterialRepositoryDB) Create(materialInput *models.Material) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
			return
		}
	}()

	sqlStatement := `INSERT INTO material (title, description, media_url, added_by) VALUES ($1, $2, $3, $4) RETURNING id;`
	_, err = tx.Exec(sqlStatement, materialInput.Title, materialInput.Description, materialInput.MediaURL, materialInput.AddedBy)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MaterialRepositoryDB) GetById(id int) (*models.Material, error) {
	sqlStatement := `SELECT id, title, description, media_url, added_by FROM material WHERE id = $1;`
	var material models.Material
	err := repo.DB.QueryRow(sqlStatement, id).Scan(&material.ID, &material.Title, &material.Description, &material.MediaURL, &material.AddedBy)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &material, nil
}

func (repo *MaterialRepositoryDB) GetAll() (*[]models.Material, error) {
	sqlStatement := `SELECT id, title, description, media_url, added_by FROM material;`

	var materials []models.Material
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var material models.Material
		err = rows.Scan(&material.ID, &material.Title, &material.Description, &material.MediaURL, &material.AddedBy)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return &materials, nil
}
