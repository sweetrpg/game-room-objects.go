package data

import (
	"context"
	"testing"
	"time"

	"github.com/sweetrpg/game-room-objects.go/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestUpdateVolumeTitleByVolumeID(t *testing.T) {
	ctx := context.Background()

	// Connect to MongoDB. Expects a local instance (e.g., via docker or testcontainers).
	// Set MONGO_URI env var to override the default mongodb://localhost:27017
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		t.Skipf("MongoDB not available at %s: %v", uri, err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("game_room_test")
	collection := db.Collection("libraries")
	defer collection.Drop(ctx)

	// Seed test data: 3 users, each with entries including the target volume.
	targetVolID := "vol-target-1"
	otherVolID := "vol-other-1"
	oldTitle := "Old Title"
	newTitle := "New Title"

	docs := []interface{}{
		models.Library{
			ID:                "lib-user-1",
			UserID:            "user-1",
			DefaultVisibility: models.VisibilityPrivate,
			Entries: []models.LibraryEntry{
				{VolumeID: targetVolID, VolumeTitle: oldTitle, AddedAt: time.Now()},
				{VolumeID: otherVolID, VolumeTitle: "Other Title", AddedAt: time.Now()},
			},
		},
		models.Library{
			ID:                "lib-user-2",
			UserID:            "user-2",
			DefaultVisibility: models.VisibilityPrivate,
			Entries: []models.LibraryEntry{
				{VolumeID: targetVolID, VolumeTitle: oldTitle, AddedAt: time.Now()},
			},
		},
		models.Library{
			ID:                "lib-user-3",
			UserID:            "user-3",
			DefaultVisibility: models.VisibilityPrivate,
			Entries: []models.LibraryEntry{
				{VolumeID: targetVolID, VolumeTitle: oldTitle, AddedAt: time.Now()},
				{VolumeID: otherVolID, VolumeTitle: "Other Title", AddedAt: time.Now()},
			},
		},
		// User without the target volume (should not be affected)
		models.Library{
			ID:                "lib-user-4",
			UserID:            "user-4",
			DefaultVisibility: models.VisibilityPrivate,
			Entries: []models.LibraryEntry{
				{VolumeID: otherVolID, VolumeTitle: "Other Title", AddedAt: time.Now()},
			},
		},
	}

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		t.Fatalf("InsertMany failed: %v", err)
	}

	// Call the function
	affectedUserIDs, err := UpdateVolumeTitleByVolumeID(ctx, collection, targetVolID, newTitle)
	if err != nil {
		t.Fatalf("UpdateVolumeTitleByVolumeID failed: %v", err)
	}

	// Assert correct user IDs were affected
	expectedUsers := []string{"user-1", "user-2", "user-3"}
	if len(affectedUserIDs) != len(expectedUsers) {
		t.Errorf("affected user count = %d, want %d", len(affectedUserIDs), len(expectedUsers))
	}

	// Check all expected users are in the result
	userMap := make(map[string]bool)
	for _, uid := range affectedUserIDs {
		userMap[uid] = true
	}
	for _, expected := range expectedUsers {
		if !userMap[expected] {
			t.Errorf("expected user %s not in results", expected)
		}
	}

	// Verify all target entries were updated
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		t.Fatalf("Find failed: %v", err)
	}
	defer cursor.Close(ctx)

	var libraries []models.Library
	if err = cursor.All(ctx, &libraries); err != nil {
		t.Fatalf("All failed: %v", err)
	}

	for _, lib := range libraries {
		for i, entry := range lib.Entries {
			if entry.VolumeID == targetVolID {
				if entry.VolumeTitle != newTitle {
					t.Errorf("user %s entry %d: VolumeTitle = %q, want %q", lib.UserID, i, entry.VolumeTitle, newTitle)
				}
			} else if entry.VolumeID == otherVolID {
				// Other volumes should not be changed
				if entry.VolumeTitle != "Other Title" {
					t.Errorf("user %s entry %d: VolumeTitle = %q, want %q (should be unchanged)", lib.UserID, i, entry.VolumeTitle, "Other Title")
				}
			}
		}
	}

	// Test idempotence: calling again with the same title should report no further changes
	affectedUserIDsSecond, err := UpdateVolumeTitleByVolumeID(ctx, collection, targetVolID, newTitle)
	if err != nil {
		t.Fatalf("Second UpdateVolumeTitleByVolumeID failed: %v", err)
	}

	// Second call should still return the same affected user IDs (they have the entries)
	// but the update operation won't modify any documents since the title is already set.
	// However, we still query and return the users who have the entry.
	if len(affectedUserIDsSecond) != len(expectedUsers) {
		t.Errorf("second call affected user count = %d, want %d", len(affectedUserIDsSecond), len(expectedUsers))
	}

	// Verify data is still correct
	cursor, err = collection.Find(ctx, bson.M{})
	if err != nil {
		t.Fatalf("Find failed: %v", err)
	}
	defer cursor.Close(ctx)

	libraries = nil
	if err = cursor.All(ctx, &libraries); err != nil {
		t.Fatalf("All failed: %v", err)
	}

	for _, lib := range libraries {
		for i, entry := range lib.Entries {
			if entry.VolumeID == targetVolID {
				if entry.VolumeTitle != newTitle {
					t.Errorf("after second call, user %s entry %d: VolumeTitle = %q, want %q", lib.UserID, i, entry.VolumeTitle, newTitle)
				}
			}
		}
	}
}

func TestUpdateVolumeTitleByVolumeID_NoMatches(t *testing.T) {
	ctx := context.Background()

	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		t.Skipf("MongoDB not available at %s: %v", uri, err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("game_room_test")
	collection := db.Collection("libraries")
	defer collection.Drop(ctx)

	// Seed a library without the target volume
	docs := []interface{}{
		models.Library{
			ID:                "lib-user-1",
			UserID:            "user-1",
			DefaultVisibility: models.VisibilityPrivate,
			Entries: []models.LibraryEntry{
				{VolumeID: "vol-other", VolumeTitle: "Other", AddedAt: time.Now()},
			},
		},
	}

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		t.Fatalf("InsertMany failed: %v", err)
	}

	// Call for a non-existent volume
	affectedUserIDs, err := UpdateVolumeTitleByVolumeID(ctx, collection, "vol-nonexistent", "New Title")
	if err != nil {
		t.Fatalf("UpdateVolumeTitleByVolumeID failed: %v", err)
	}

	// Should return an empty list
	if len(affectedUserIDs) != 0 {
		t.Errorf("affected user count = %d, want 0", len(affectedUserIDs))
	}
}
