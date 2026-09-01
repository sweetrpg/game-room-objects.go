package data

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UpdateVolumeTitleByVolumeID sets the VolumeTitle on every library entry
// referencing the given volumeID across all user libraries in a single bulk operation.
// Returns the list of affected user IDs and any error encountered.
//
// This is a system-triggered update driven by trusted catalog events
// (e.g., volume title changes), not a user-initiated request.
// Owner-scoped access rules apply to user requests; this event-driven
// denormalization refresh deliberately spans all users.
func UpdateVolumeTitleByVolumeID(ctx context.Context, collection *mongo.Collection, volumeID string, newTitle string) (affectedUserIDs []string, err error) {
	// Update all library documents that have an entry with the matching volume_id.
	// Use arrayFilters to target only the matching array element within each document.
	updateOpts := options.UpdateMany().SetArrayFilters([]interface{}{
		bson.M{"elem.volume_id": volumeID},
	})
	result, err := collection.UpdateMany(
		ctx,
		bson.M{"entries": bson.M{"$elemMatch": bson.M{"volume_id": volumeID}}},
		bson.A{
			bson.M{
				"$set": bson.M{
					"entries.$[elem].volume_title": newTitle,
				},
			},
		},
		updateOpts,
	)
	if err != nil {
		return nil, err
	}

	// If no documents were modified, return an empty list.
	if result.ModifiedCount == 0 {
		return []string{}, nil
	}

	// Query for all documents that have entries with the target volume_id
	// to get the affected user IDs.
	findOpts := options.Find().SetProjection(bson.M{"user_id": 1})
	cursor, err := collection.Find(
		ctx,
		bson.M{"entries": bson.M{"$elemMatch": bson.M{"volume_id": volumeID}}},
		findOpts,
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Extract user IDs from the matched documents.
	var results []struct {
		UserID string `bson:"user_id"`
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	userIDs := make([]string, len(results))
	for i, r := range results {
		userIDs[i] = r.UserID
	}

	return userIDs, nil
}
