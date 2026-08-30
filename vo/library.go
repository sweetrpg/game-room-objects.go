package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Library value object.
// This value object is a serializable representation of the Library model.
type LibraryVO struct {
	ID                string           `json:"id" jsonapi:"primary,library"`
	UserID            string           `json:"user_id" jsonapi:"relation,user"`
	DefaultVisibility string           `json:"default_visibility" jsonapi:"attr,default_visibility"`
	Entries           []LibraryEntryVO `json:"entries" jsonapi:"attr,entries"`
	modelcore.AuditableVO
}

// LibraryEntry value object.
type LibraryEntryVO struct {
	VolumeID           string    `json:"volume_id"`
	VolumeTitle        string    `json:"volume_title"`
	VisibilityOverride *string   `json:"visibility_override,omitempty"`
	AddedAt            time.Time `json:"added_at"`
}
