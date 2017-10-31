package wolfram_alpha

import (
	"encoding/json"
	"net/http"
	"os"

	"Prime/chatbot/bundles/common"
)

type WolframAlphaController struct {
	common.Controller
}

func (c *WolframAlphaController) HandleQuery(message string) (*Response, error) {

	// instance of response
	response := new(Response)
	// send http request to wolfram alpha api
	url := os.Getenv("BASE_URL") + "&input=" + message + "&appid=" + os.Getenv("APP_ID") + "&includepodid=Result"
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
