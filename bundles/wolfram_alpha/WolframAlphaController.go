package wolfram_alpha

import (
	"encoding/json"
	"net/http"
	"net/url"

	"../common"
	"github.com/Lafriakh/env"
)

type WolframAlphaController struct {
	common.Controller
}

func (c *WolframAlphaController) HandleQuery(w http.ResponseWriter, req *http.Request) {
	env.New()
	// map to read request in
	parsed := make(map[string]string)
	// Parse the request
	_ = json.NewDecoder(req.Body).Decode(&parsed)
	// get the query
	query := url.PathEscape(parsed["query"])

	// instance of response
	response := new(Response)
	// send http request to wolfram alpha api
	url := env.GetString("BASE_URL") + "&input=" + query + "&appid=" + env.GetString("APP_ID") + "&includepodid=Result"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	// Decode response to json
	json.NewDecoder(resp.Body).Decode(&response)

	// return json response
	json.NewEncoder(w).Encode(response)

}
