package converts

import "testing"

func TestConvertAllDnisToUsers(t *testing.T) {
	dnis := []string{"00000011B", "00000021K", "X0000001R", "Y0000001S"}
	expected := []string{"B00000011", "K00000021", "XR0000001", "YS0000001"}
	actual := ConvertAllDnisToUsers(dnis)
	if len(actual) != len(expected) {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}
