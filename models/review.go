package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
)

type Review struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Title     string       `json:"title" db:"title"`
	Content   nulls.String `json:"content" db:"content"`
	MovieID   uuid.UUID    `json:"movie_id" db:"movie_id"`
}

// String is not required by pop and may be deleted
func (r Review) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Reviews is not required by pop and may be deleted
type Reviews []Review

// String is not required by pop and may be deleted
func (r Reviews) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (r *Review) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (r *Review) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (r *Review) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
