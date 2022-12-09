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
	Context("GenerateJson", func() {
		It("GenerateJson", func() {
			log.InitTestLogger("GenerateACLJson")

			var userNames = []string{"Usuario Uno", "Usuario Dos"}
			var userPasswords = []string{"B000001b", "K000002k"}
			var userUsers = []string{"B0000011", "K0000021"}
			var phoneNumbers = []string{"666777888", "666888777"}

			result := json.GenerateJson(userNames, userPasswords, userUsers, phoneNumbers)

			expectResult := file.ReadFile(consts.UserCouchbaseRoute)
			expectResult = strings.TrimSuffix(expectResult, "\n")

			Expect(result).To(BeEquivalentTo(expectResult))
		})
	})

	Context("GenerateEndpointJson", func() {
		It("GenerateEndpointJson", func() {
			log.InitTestLogger("GenerateACLJson")

			var userNames = []string{"Usuario Uno", "Usuario Dos"}
			var userPasswords = []string{"B000001b", "K000002k"}
			var userUsers = []string{"B0000011", "K0000021"}
			var phoneNumbers = []string{"666777888", "666888777"}

			result := json.GenerateEndpointJson(userNames, userPasswords, userUsers, phoneNumbers)
			expectResult := file.ReadFile(consts.AclEPCouchbaseRoute)
			expectResult = strings.TrimSuffix(expectResult, "\n")
			log.Debugln("RESULT:\n" + result)
			log.Debugln("EXPECTED:\n" + expectResult)

			Expect(result).To(BeEquivalentTo(expectResult))
		})
	})
})
