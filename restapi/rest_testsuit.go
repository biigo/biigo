package restapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	restful "github.com/emicklei/go-restful"
)

//RestTestSuite 接口测试套件
type RestTestSuite struct {
}

//Request 请求接口
func (suite RestTestSuite) Request(method, url string, body interface{}) *httptest.ResponseRecorder {
	var b io.Reader
	if body != nil {
		bodyByte, _ := json.Marshal(body)
		b = bytes.NewReader(bodyByte)
	}
	req, _ := http.NewRequest(method, url, b)
	req.Header.Set("Content-Type", restful.MIME_JSON)
	resp := httptest.NewRecorder()
	restful.DefaultContainer.ServeHTTP(resp, req)
	return resp
}
