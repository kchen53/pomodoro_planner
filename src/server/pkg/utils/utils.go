package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ParseBodyTest(r *httptest.ResponseRecorder, x interface{}) {
	if err := json.Unmarshal(r.Body.Bytes(), x); err != nil {
		return
	}
}
