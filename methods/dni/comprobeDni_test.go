package dni

import (
	"drivers-create/methods/log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestComprobeDniAndNie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AllDnis are Correct")
}

var _ = Describe("AllDnis are Correct", func() {
	Context("AllDnis are Correct", func() {
		It("All correct", func() {
			log.InitTestLogger()
			documents := []string{"00000011B", "00000021K", "X0000001R", "Y0000001S"}

			_, err := ComprobeAllDnis(documents)

			//Expect(err).ToNot(nil)
			Expect(err).To(BeNil())
		})
	})
})
