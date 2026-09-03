package models

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// Wishlist is a user's collection of wanted catalog volumes, with its own
// visibility independent of the library.
type Wishlist struct {
	ID         string          `bson:"_id" json:"id" jsonapi:"primary,wishlist"`
	UserID     string          `bson:"user_id" json:"user_id" jsonapi:"relation,user"`
	Name       string          `bson:"name" json:"name" jsonapi:"attr,name"`
	Visibility Visibility      `bson:"visibility" json:"visibility" jsonapi:"attr,visibility"`
	Entries    []WishlistEntry `bson:"entries" json:"entries" jsonapi:"attr,entries"`
	modelcore.Auditable
}

// WishlistEntry links a wishlist to one wanted catalog volume, with a
// denormalized volume title snapshot captured at add time and kept current by
// the volume-updated consumer (mirrors LibraryEntry).
type WishlistEntry struct {
	VolumeID    string    `bson:"volume_id" json:"volume_id"`
	VolumeTitle string    `bson:"volume_title" json:"volume_title"`
	AddedAt     time.Time `bson:"added_at" json:"added_at"`
}

// NewWishlist creates a wishlist defaulting to private visibility, per the
// "new wishlists default to private" requirement.
func NewWishlist(id, userID, name string) Wishlist {
	return Wishlist{
		ID:         id,
		UserID:     userID,
		Name:       name,
		Visibility: VisibilityPrivate,
	}
}
