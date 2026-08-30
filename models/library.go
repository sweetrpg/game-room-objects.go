package models

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// Library is a user's collection of owned catalog volumes.
type Library struct {
	ID                string         `bson:"_id" json:"id" jsonapi:"primary,library"`
	UserID            string         `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	DefaultVisibility Visibility     `bson:"default_visibility" json:"default_visibility" jsonapi:"attr,default_visibility"`
	Entries           []LibraryEntry `bson:"entries" json:"entries" jsonapi:"attr,entries"`
	modelcore.Auditable
}

// LibraryEntry links a library to one owned catalog volume, with a denormalized
// volume title snapshot and an optional per-entry visibility override.
type LibraryEntry struct {
	VolumeID           string      `bson:"volume_id" json:"volume_id"`
	VolumeTitle        string      `bson:"volume_title" json:"volume_title"`
	VisibilityOverride *Visibility `bson:"visibility_override,omitempty" json:"visibility_override,omitempty"`
	AddedAt            time.Time   `bson:"added_at" json:"added_at"`
}

// NewLibrary creates a library defaulting to private visibility, per the
// "new library defaults to private" requirement.
func NewLibrary(id, userID string) Library {
	return Library{
		ID:                id,
		UserID:            userID,
		DefaultVisibility: VisibilityPrivate,
	}
}

// EffectiveVisibility resolves the entry's visibility: its own override when
// present, otherwise the library's default.
func (e LibraryEntry) EffectiveVisibility(libraryDefault Visibility) Visibility {
	if e.VisibilityOverride != nil {
		return *e.VisibilityOverride
	}
	return libraryDefault
}
