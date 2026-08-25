package models

import (
	modelcore "github.com/sweetrpg/model-core/models"
)

// GameRoom model.
// This model represents a game room of RPG resources that belongs to a user.
type GameRoom struct {
	ID     string          `bson:"_id" json:"id" jsonapi:"primary,game-room"`
	UserID string          `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	Notes  string          `json:"notes" jsonapi:"attr,notes"`
	Tags   []modelcore.Tag `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
