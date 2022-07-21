package sql

import (
	"drivers-create/consts"
	"drivers-create/methods/file"
	"drivers-create/methods/log"
	"strings"
	"testing"
)

func TestGenerateSqlLiteInsertDriversTable(t *testing.T) {
	log.InitTestLogger()
	allNames := []string{"Usuario Uno", "Usuario Dos"}
	allUsers := []string{"B0000011", "K0000021"}
	allDnis := []string{"00000011B", "00000021K"}
	allPhones := []string{"666777888", "722444777"}

	query := GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)

	expected := file.ReadFile(consts.DriverInsertTestRoute)
	expected = strings.TrimSuffix(expected, "\n")

	if query != expected {
		t.Errorf("Expected: %v, got: %v", expected, query)
	}
}

func TestGenerateSqlLiteInsertRelation(t *testing.T) {
	log.InitTestLogger()
	allDnis := []string{"00000011B", "00000021K", "", "00000011B", "00000021K"}
	shopCodes := []string{"1", "2"}

	query := GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)

	expected := file.ReadFile(consts.DriversShopTestRoute)
	expected = strings.TrimSuffix(expected, "\n")

	if query != expected {
		t.Errorf("Expected:\n %v, got:\n %v", expected, query)
	}
}

func TestGenerateSqlLiteInsertShop(t *testing.T) {
	log.InitTestLogger()
	shopCodes := []string{"1", "2"}
	shopNames := []string{"Shop One", "Shop Two"}

	query := GenerateSqlLiteInsertShopTable(shopCodes, shopNames)

	expected := file.ReadFile(consts.ShopInsertTestRoute)
	expected = strings.TrimSuffix(expected, "\n")

	if query != expected {
		t.Errorf("Expected:\n %v, got:\n %v", expected, query)
	}
}
