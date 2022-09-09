package converts

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsersToPasswords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert Users to Password")
}

var _ = Describe("Convert Users to Password", func() {
	Context("Users to Password", func() {
		It("User To Password", func() {
			users := []string{"B0000011", "K0000021", "XR000001", "YS000001"}
			expected := []string{"B000001b", "K000002k", "XR00000r", "YS00000s"}
			actual := ConvertAllUsersToPasswords(users)

			Expect(actual).To(BeEquivalentTo(expected))
		})
	})
})
