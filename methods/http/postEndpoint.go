package http

import (
	"bytes"
	"drivers-create/consts"
	logs "drivers-create/methods/log"
	"fmt"
	"net/http"
)

func AuthEndpointCall(usersJson string) {
	logs.Debugln("URL to POST:", consts.AuthEndpointUrl)

	var jsonStr = []byte(usersJson)
	req, err := http.NewRequest("POST", consts.AuthEndpointUrl, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logs.Debugln("response Status:", resp.Status)
	if resp.Status != "204 No Content" {
		fmt.Println("Error calling endpoint, check logs")
		logs.Errorln("Error sending the data, check ACL and couchbase, after publish with postman")
		return
	}
}
