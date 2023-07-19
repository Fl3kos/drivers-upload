package dniToUser

import (
	"support-utils/methods/log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConvertAllDnisToUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert Dnis to Users")
}

var _ = Describe("Convert Dnis to Users", func() {
	Context("Convert Dnis toUsers", func() {
		It("All Dnis To User", func() {
			log.InitTestLogger("All Dnis To User")

			dnis := []string{"00000011B", "00000021K", "X0000001R", "Y0000001S"}
			expected := []string{"B0000011", "K0000021", "XR000001", "YS000001"}
			actual := ConvertAllDnisToUsers(dnis)

			Expect(actual).To(BeEquivalentTo(expected))
		})
	})
})
