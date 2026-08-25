package models

import (
	"encoding/json"
	"testing"
)

func TestNewLibraryDefaultsPrivate(t *testing.T) {
	lib := NewLibrary("lib-1", "user-1")
	if lib.DefaultVisibility != VisibilityPrivate {
		t.Errorf("DefaultVisibility = %s, want %s", lib.DefaultVisibility, VisibilityPrivate)
	}
}

func TestLibraryEntryEffectiveVisibility(t *testing.T) {
	override := VisibilityPublic
	overridden := LibraryEntry{VolumeID: "vol-1", VisibilityOverride: &override}
	if got := overridden.EffectiveVisibility(VisibilityPrivate); got != VisibilityPublic {
		t.Errorf("EffectiveVisibility with override = %s, want %s", got, VisibilityPublic)
	}

	inherited := LibraryEntry{VolumeID: "vol-2"}
	if got := inherited.EffectiveVisibility(VisibilityFriends); got != VisibilityFriends {
		t.Errorf("EffectiveVisibility without override = %s, want %s", got, VisibilityFriends)
	}
}

func TestLibraryRoundTrip(t *testing.T) {
	override := VisibilityPrivate
	lib := Library{
		ID:                "lib-1",
		UserID:            "user-1",
		DefaultVisibility: VisibilityPublic,
		Entries: []LibraryEntry{
			{VolumeID: "vol-1"},
			{VolumeID: "vol-2", VisibilityOverride: &override},
		},
	}

	data, err := json.Marshal(lib)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Library
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != lib.ID || got.DefaultVisibility != lib.DefaultVisibility {
		t.Fatalf("round-trip mismatch: got %+v, want %+v", got, lib)
	}
	if len(got.Entries) != 2 || got.Entries[1].VisibilityOverride == nil || *got.Entries[1].VisibilityOverride != VisibilityPrivate {
		t.Fatalf("entries did not round-trip: %+v", got.Entries)
	}
}
