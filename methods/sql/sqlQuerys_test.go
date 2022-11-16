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
	Context("GenerateSqlLiteInsertDriversTable", func() {
		It("GenerateSqlLiteInsertDriversTable", func() {
			log.InitTestLogger()
			allNames := []string{"Usuario Uno", "Usuario Dos"}
			allUsers := []string{"B0000011", "K0000021"}
			allDnis := []string{"00000011B", "00000021K"}
			allPhones := []string{"666777888", "722444777"}

			query := GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)

			expected := file.ReadFile(consts.DriverInsertTestRoute)
			expected = strings.TrimSuffix(expected, "\n")

			Expect(query).To(BeEquivalentTo(expected))
		})
	})
	Context("GenerateSqlLiteInsertRelation", func() {
		It("GenerateSqlLiteInsertRelation", func() {
			log.InitTestLogger()
			allDnis := []string{"00000011B", "00000021K", "", "00000011B", "00000021K"}
			shopCodes := []string{"1", "2"}

			query := GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)

			expected := file.ReadFile(consts.DriversShopTestRoute)
			expected = strings.TrimSuffix(expected, "\n")

			Expect(query).To(BeEquivalentTo(expected))
		})
	})
	Context("GenerateSqlLiteInsertShop", func() {
		It("GenerateSqlLiteInsertShop", func() {
			log.InitTestLogger()
			shopCodes := []string{"1", "2"}
			shopNames := []string{"Shop One", "Shop Two"}

			query := GenerateSqlLiteInsertShopTable(shopCodes, shopNames)

			expected := file.ReadFile(consts.ShopInsertTestRoute)
			expected = strings.TrimSuffix(expected, "\n")

			Expect(query).To(BeEquivalentTo(expected))
		})
	})

	Context("GenerateACLInsert", func() {
		It("GenerateACLInsert", func() {
			log.InitTestLogger()

			users := []string{"B0000011", "K0000021"}

			query := GenerateAclInsert(users)

			expectedResult := file.ReadFile(consts.AclSqlRoute)
			expectedResult = strings.TrimSuffix(expectedResult, "\n")

			Expect(query).To(BeEquivalentTo(expectedResult))
		})
	})
})
