package converts

import "testing"

func TestUsersToPasswords(t *testing.T) {
	users := []string{"B00000011", "K00000021", "XR0000001", "YS0000001"}
	expected := []string{"B0000001b", "K0000002k", "XR000000x", "YS000000y"}
	actual := ConvertAllUsersToPasswords(users)
	if len(actual) != len(expected) {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}
