package dni

import (
	"drivers-create/methods/log"
	"testing"
)

func TestComprobeDniAndNie(t *testing.T) {
	log.InitTestLogger()
	documents := []string{"00000011B", "00000021K", "X0000001R", "Y0000001S"}

	_, err := ComprobeAllDnis(documents)

	if err != nil {
		t.Errorf("Expected: %v, got: %v", nil, err)
	}
}
