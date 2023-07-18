package sql

import (
	"drivers-create/consts"
	"drivers-create/methods/file"
	"drivers-create/methods/log"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SqlTest")
}

var _ = Describe("Sql Test", func() {
	Context("GenerateACLInsert", func() {
		It("GenerateACLInsert", func() {
			log.InitTestLogger("GenerateACLInsert")

			users := []string{"B0000011", "K0000021"}

			query := GenerateAclInsert(users, "ROLE_APPTMS_DRIVER")

			expectedResult := file.ReadFile(consts.AclSqlRoute)
			expectedResult = strings.TrimSuffix(expectedResult, "\n")

			Expect(query).To(BeEquivalentTo(expectedResult))
		})
	})
})
