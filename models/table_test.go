package models

import (
	"encoding/json"
	"testing"
)

func TestNewTableDefaultsPrivate(t *testing.T) {
	tbl := NewTable("tbl-1", "user-1", "Friday Night Game")
	if tbl.Visibility != VisibilityPrivate {
		t.Errorf("Visibility = %s, want %s", tbl.Visibility, VisibilityPrivate)
	}
}

func TestTableRoundTrip(t *testing.T) {
	tbl := Table{
		ID:         "tbl-1",
		UserID:     "user-1",
		Name:       "Friday Night Game",
		VolumeIDs:  []string{"vol-1", "vol-2"},
		Visibility: VisibilityFriendsOfFriends,
	}

	data, err := json.Marshal(tbl)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Table
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != tbl.ID || got.Name != tbl.Name || got.Visibility != tbl.Visibility || len(got.VolumeIDs) != 2 {
		t.Fatalf("round-trip mismatch: got %+v, want %+v", got, tbl)
	}
}
