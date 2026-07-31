package models

import (
	modelcore "github.com/sweetrpg/model-core/models"
)

// Shelf model.
// This model represents a shelf of RPG resources that belongs to a user.
type Shelf struct {
	ID     string          `bson:"_id" json:"id" jsonapi:"primary,shelf"`
	UserID string          `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	Notes  string          `json:"notes" jsonapi:"attr,notes"`
	Tags   []modelcore.Tag `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
