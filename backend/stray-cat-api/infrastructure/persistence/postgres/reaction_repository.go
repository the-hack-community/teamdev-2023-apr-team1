// infrastructure/persistence/postgres/reaction_repository.go
package postgres

import (
	"database/sql"
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type ReactionRepository struct {
	DB *sql.DB
}

func NewReactionRepository(db *sql.DB) repository.ReactionRepository {
	return &ReactionRepository{DB: db}
}

func (r *ReactionRepository) FindByID(reactionID string) (*model.Reaction, error) {
	reaction := &model.Reaction{}
	err := r.DB.QueryRow("SELECT reaction_id, user_id, cat_id, date_time FROM reactions WHERE reaction_id = $1", reactionID).Scan(&reaction.ReactionID, &reaction.UserID, &reaction.CatID, &reaction.DateTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

func (r *ReactionRepository) Store(reaction *model.Reaction) error {
	err := r.DB.QueryRow("INSERT INTO reactions(reaction_id, user_id, cat_id, date_time) VALUES($1, $2, $3, $4) RETURNING reaction_id", reaction.ReactionID, reaction.UserID, reaction.CatID, reaction.DateTime).Scan(&reaction.ReactionID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionRepository) Update(reaction *model.Reaction) error {
	_, err := r.DB.Exec("UPDATE reactions SET user_id = $1, cat_id = $2, date_time = $3 WHERE reaction_id = $4", reaction.UserID, reaction.CatID, reaction.DateTime, reaction.ReactionID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionRepository) Delete(reactionID string) error {
	_, err := r.DB.Exec("DELETE FROM reactions WHERE reaction_id = $1", reactionID)
	if err != nil {
		return err
	}
	return nil
}
