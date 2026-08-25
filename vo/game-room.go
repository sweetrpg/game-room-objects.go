package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

// GameRoom value object.
// This value object is a serializable representation of the GameRoom model.
type GameRoomVO struct {
	ID     string            `json:"id" jsonapi:"primary,license"`
	UserID string            `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	Notes  string            `json:"notes" jsonapi:"attr,notes"`
	Tags   []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
