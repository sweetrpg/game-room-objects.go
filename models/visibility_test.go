package models

import "testing"

func TestMoreExposedThan(t *testing.T) {
	cases := []struct {
		a, b Visibility
		want bool
	}{
		{VisibilityPrivate, VisibilityFriends, false},
		{VisibilityFriends, VisibilityPrivate, true},
		{VisibilityFriends, VisibilityFriendsOfFriends, false},
		{VisibilityFriendsOfFriends, VisibilityFriends, true},
		{VisibilityFriendsOfFriends, VisibilityPublic, false},
		{VisibilityPublic, VisibilityFriendsOfFriends, true},
		{VisibilityPublic, VisibilityPublic, false},
	}
	for _, c := range cases {
		if got := c.a.MoreExposedThan(c.b); got != c.want {
			t.Errorf("%s.MoreExposedThan(%s) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
