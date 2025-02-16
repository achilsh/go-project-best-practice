package unit_test_demo

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestHttpMock(t *testing.T) {
	// 启用 httpmock
	httpmock.Activate()
	defer httpmock.Deactivate()

	//add new responder, When a request comes in that matches,
	// the responder is called and the response returned to the client.
	httpmock.RegisterResponder("GET", "https://abc.xyz.com/",
		httpmock.NewStringResponder(200, "hell world")) //or `{"id":123, "name":"sz"}`

	// 实际发起请求：
	// 发送一个GET请求
	resp, err := http.Get("https://abc.xyz.com/")
	if err != nil {
		fmt.Println("请求出错:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return
	}
	// 输出响应内容
	fmt.Println("http mock response data: ", string(body))
}

func TestHttpMockWithResponder(t *testing.T) {
	// 启用 httpmock
	httpmock.Activate()
	defer httpmock.Deactivate()

	//add new responder, When a request comes in that matches,
	// the responder is called and the response returned to the client.
	httpmock.RegisterResponder("GET", "https://abc.xyz.com/",

		// 自定义 responder 的处理
		func(in *http.Request) (*http.Response, error) {

			articles := make([]map[string]interface{}, 0)
			// can get json data from file.
			// httpmock.NewJsonResponse(200, httpmock.File("body.json"))
			resp, err := httpmock.NewJsonResponse(200, articles)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			// or get request body from in.Body, eg:
			//article := make(map[string]interface{})
			//if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
			//	return httpmock.NewStringResponse(400, ""), nil
			//}
			//
			//articles = append(articles, article)
			//
			//resp, err := httpmock.NewJsonResponse(200, article)
			//if err != nil {
			//	return httpmock.NewStringResponse(500, ""), nil
			//}
			return resp, nil
		})

	// 实际发起请求：
	// 发送一个GET请求
	resp, err := http.Get("https://abc.xyz.com/")
	if err != nil {
		fmt.Println("请求出错:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return
	}
	// 输出响应内容
	fmt.Println("http mock response data: ", string(body))
}
