package models

import (
	"encoding/json"
	"testing"
)

func TestNewWishlistDefaultsPrivate(t *testing.T) {
	wl := NewWishlist("wl-1", "user-1")
	if wl.Visibility != VisibilityPrivate {
		t.Errorf("Visibility = %s, want %s", wl.Visibility, VisibilityPrivate)
	}
}

func TestWishlistRoundTrip(t *testing.T) {
	wl := Wishlist{
		ID:         "wl-1",
		UserID:     "user-1",
		Visibility: VisibilityPublic,
		Entries:    []WishlistEntry{{VolumeID: "vol-1"}},
	}

	data, err := json.Marshal(wl)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Wishlist
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != wl.ID || got.Visibility != wl.Visibility || len(got.Entries) != 1 || got.Entries[0].VolumeID != "vol-1" {
		t.Fatalf("round-trip mismatch: got %+v, want %+v", got, wl)
	}
}
