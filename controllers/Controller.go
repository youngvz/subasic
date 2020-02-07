package controllers

import "fmt"

type Controller struct { }

func (c Controller) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
}

func (c Controller) createPayload(code int, message string) map[string]interface{} {
	var payload map[string]interface{}
	payload["statusCode"] = code
	payload["message"] = message
	return payload
}
