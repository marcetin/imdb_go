package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Movie struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Name        string       `json:"name" db:"name"`
	Description nulls.String `json:"description" db:"description"`
	Url         string       `json:"url" db:"url"`
	Ratings     int          `json:"ratings" db:"ratings"`
}

// String is not required by pop and may be deleted
func (m Movie) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Movies is not required by pop and may be deleted
type Movies []Movie

// String is not required by pop and may be deleted
func (m Movies) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (m *Movie) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Name, Name: "Name"},
		//&validators.BytesArePresent{Field: m.Description, Name: "Description"},
		&validators.StringIsPresent{Field: m.Url, Name: "Url"},
		&validators.IntIsPresent{Field: m.Ratings, Name: "Ratings"},
	), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (m *Movie) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (m *Movie) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (m *Movie) Reviews(tx *pop.Connection) (*Reviews, error) {
	reviews := &Reviews{}

	err := tx.BelongsTo(m).All(reviews)
	if err != nil {
		return reviews, errors.WithStack(err)
	}
	return reviews, nil
}
