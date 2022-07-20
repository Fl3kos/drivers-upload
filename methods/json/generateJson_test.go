package json_test

import (
	"drivers-create/methods/file"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	"strings"
	"testing"
)

func TestGenerateJson(t *testing.T) {

	log.InitTestLogger()
	var userNames = []string{"Usuario Uno", "Usuario Dos"}
	var userPasswords = []string{"B000001b", "K000002k"}
	var userUsers = []string{"B0000011", "K0000021"}

	result := json.GenerateJson(userNames, userPasswords, userUsers)

	expectResult := file.ReadFile("../../test/json/userCouchbaseTest.json")
	expectResult = strings.TrimSuffix(expectResult, "\n")

	if result == expectResult {
		t.Logf("The user json is correct")
	} else {
		t.Errorf("The user json is incorrect")
	}

}
