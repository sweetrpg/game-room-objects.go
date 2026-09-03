package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewWishlistDefaultsPrivate(t *testing.T) {
	wl := NewWishlist("wl-1", "user-1", "Birthday")
	if wl.Visibility != VisibilityPrivate {
		t.Errorf("Visibility = %s, want %s", wl.Visibility, VisibilityPrivate)
	}
	if wl.Name != "Birthday" {
		t.Errorf("Name = %s, want Birthday", wl.Name)
	}
}

func TestWishlistRoundTrip(t *testing.T) {
	addedAt := time.Date(2026, 8, 28, 12, 0, 0, 0, time.UTC)
	wl := Wishlist{
		ID:         "wl-1",
		UserID:     "user-1",
		Name:       "Con haul",
		Visibility: VisibilityPublic,
		Entries:    []WishlistEntry{{VolumeID: "vol-1", VolumeTitle: "Pathfinder Core", AddedAt: addedAt}},
	}

	data, err := json.Marshal(wl)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Wishlist
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != wl.ID || got.Name != wl.Name || got.Visibility != wl.Visibility || len(got.Entries) != 1 || got.Entries[0].VolumeID != "vol-1" {
		t.Fatalf("round-trip mismatch: got %+v, want %+v", got, wl)
	}
	if got.Entries[0].VolumeTitle != "Pathfinder Core" {
		t.Fatalf("VolumeTitle did not round-trip: got %q", got.Entries[0].VolumeTitle)
	}
	if !got.Entries[0].AddedAt.Equal(addedAt) {
		t.Fatalf("AddedAt did not round-trip: got %v, want %v", got.Entries[0].AddedAt, addedAt)
	}
}
