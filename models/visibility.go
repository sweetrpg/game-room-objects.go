package models

// Visibility controls who may see a Shelf collection or entry.
type Visibility string

const (
	VisibilityPrivate          Visibility = "private"
	VisibilityFriends          Visibility = "friends"
	VisibilityFriendsOfFriends Visibility = "friends-of-friends"
	VisibilityPublic           Visibility = "public"
)

// visibilityRank orders levels from least to most exposed. Used only to
// determine whether a change makes an entry more visible - never to clamp
// overrides, which may move in either direction relative to a default.
var visibilityRank = map[Visibility]int{
	VisibilityPrivate:          0,
	VisibilityFriends:          1,
	VisibilityFriendsOfFriends: 2,
	VisibilityPublic:           3,
}

// MoreExposedThan reports whether v is visible to a strictly larger audience than other.
func (v Visibility) MoreExposedThan(other Visibility) bool {
	return visibilityRank[v] > visibilityRank[other]
}
