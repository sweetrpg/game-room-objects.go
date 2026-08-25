package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// Table is a named play collection grouping multiple catalog volumes, with
// its own visibility independent of the owner's library or wishlist.
type Table struct {
	ID         string     `bson:"_id" json:"id" jsonapi:"primary,table"`
	UserID     string     `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	Name       string     `bson:"name" json:"name" jsonapi:"attr,name"`
	VolumeIDs  []string   `bson:"volume_ids" json:"volume_ids" jsonapi:"relation,volume"`
	Visibility Visibility `bson:"visibility" json:"visibility" jsonapi:"attr,visibility"`
	modelcore.Auditable
}

// NewTable creates a table defaulting to private visibility, per the "new
// tables default to private" requirement.
func NewTable(id, userID, name string) Table {
	return Table{
		ID:         id,
		UserID:     userID,
		Name:       name,
		Visibility: VisibilityPrivate,
	}
}
