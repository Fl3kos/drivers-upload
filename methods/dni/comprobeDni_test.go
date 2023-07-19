package dni

import (
	"support-utils/methods/log"
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
		It("Happy Path Only One Shop DNI", func() {
			log.InitTestLogger("AllDnis are Correct")
			documents := []string{"00000011B", "00000021K"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).To(BeNil())
		})

		It("Happy Path Only One Shop NIE", func() {
			log.InitTestLogger("Happy Path Only One Shop NIE")
			documents := []string{"X0000001R", "Y0000001S", "Z0000001Y"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).To(BeNil())
		})

		It("Happy Path Various Shop DNI", func() {
			log.InitTestLogger("Happy Path Various Shop DNI")
			documents := []string{"00000011B", "", "00000021K"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).To(BeNil())
		})

		It("Happy Path Various Shop NIE", func() {
			log.InitTestLogger("Happy Path Various Shop NIE")
			documents := []string{"X0000001R", "", "Y0000001S"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).To(BeNil())
		})

		It("Error Path Only One Shop DNI", func() {
			log.InitTestLogger("Error Path Only One Shop DNI")
			documents := []string{"00000011A", "00000021K"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).ToNot(BeNil())
		})

		It("Error Path Only One Shop NIE", func() {
			log.InitTestLogger("Error Path Only One Shop NIE")
			documents := []string{"X0000001I", "Y0000001S"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).ToNot(BeNil())
		})

		It("Error Path Various Shop DNI", func() {
			log.InitTestLogger("Error Path Various Shop DNI")
			documents := []string{"00000011A", "", "00000021K"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).ToNot(BeNil())
		})

		It("Error Path Various Shop NIE", func() {
			log.InitTestLogger("Error Path Various Shop NIE")
			documents := []string{"X0000001R", "", "Y0000001O"}

			_, err := ComprobeAllDnis(documents)

			Expect(err).ToNot(BeNil())
		})
	})
})
