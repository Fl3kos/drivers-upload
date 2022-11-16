package json_test

import (
	"drivers-create/consts"
	"drivers-create/methods/file"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGenerateJson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate Json")
}

var _ = Describe("Generate Json", func() {
	Context("GenerateCouchbaseJson", func() {
		It("GenerateJson", func() {
			log.InitTestLogger()
			var userNames = []string{"Usuario Uno", "Usuario Dos"}
			var userPasswords = []string{"B000001b", "K000002k"}
			var userUsers = []string{"B0000011", "K0000021"}

			result := json.GenerateJson(userNames, userPasswords, userUsers)

			expectResult := file.ReadFile(consts.UserCouchbaseRoute)
			expectResult = strings.TrimSuffix(expectResult, "\n")

			Expect(result).To(BeEquivalentTo(expectResult))
		})
	})

	Context("GenerateAclJson", func() {
		It("GenerateACLJson", func() {
			log.InitTestLogger()

			var userNames = []string{"Usuario Uno", "Usuario Dos"}
			var userPasswords = []string{"B000001b", "K000002k"}
			var userUsers = []string{"B0000011", "K0000021"}
			var phoneNumbers = []string{"666777888", "666888777"}

			result := json.GenerateAclJson(userNames, userPasswords, userUsers, phoneNumbers)

			expectResult := file.ReadFile(consts.AclCouchbaseRoute)
			expectResult = strings.TrimSuffix(expectResult, "\n")

			Expect(result).To(BeEquivalentTo(expectResult))
		})
	})
})
