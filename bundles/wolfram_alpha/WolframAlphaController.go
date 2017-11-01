package wolfram_alpha

import (
	"Prime/bundles/common"
	"encoding/json"
	"net/http"
	url "net/url"
	"os"
)

type WolframAlphaController struct {
	common.Controller
}

func (c *WolframAlphaController) HandleQuery(message string) (*Response, error) {

	// instance of response
	response := new(Response)
	encoded := url.QueryEscape(message)
	// send http request to wolfram alpha api
	url := os.Getenv("BASE_URL") + "&input=" + encoded + "&appid=" + os.Getenv("APP_ID") + "&includepodid=Result"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	// Decode response to json
	json.NewDecoder(resp.Body).Decode(&response)
	return response, nil
	// // return json response
	// json.NewEncoder(w).Encode(response)

}
