package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Wishlist value object.
// This value object is a serializable representation of the Wishlist model.
type WishlistVO struct {
	ID         string            `json:"id" jsonapi:"primary,wishlist"`
	UserID     string            `json:"user_id" jsonapi:"relation,user"`
	Visibility string            `json:"visibility" jsonapi:"attr,visibility"`
	Entries    []WishlistEntryVO `json:"entries" jsonapi:"attr,entries"`
	modelcore.AuditableVO
}

// WishlistEntry value object.
type WishlistEntryVO struct {
	VolumeID string    `json:"volume_id"`
	AddedAt  time.Time `json:"added_at"`
}
