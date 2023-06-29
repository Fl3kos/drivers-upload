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

func AclEndpointCall(usersJson, username, token string) {
	url := fmt.Sprintf(consts.AclEndpointUrl, username)
	logs.Debugln("URL to PUT:", url)
	var jsonStr = []byte(usersJson)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authority", "com.pro.api.dgrp.io")
	req.Header.Set("accept", "application/json, text/plain")
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("client_id", "dKSBbUAiriDZzZoVC9pLqstZHsCD0tJfx6GycX3Ox9FIG4cm")
	req.Header.Set("origin", "https://acl-web.com.pro.webs.dgrp.io")
	req.Header.Set("referer", "https://acl-web.com.pro.webs.dgrp.io")
	req.Header.Set("x-diagroup-application-id", "ACL")

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
