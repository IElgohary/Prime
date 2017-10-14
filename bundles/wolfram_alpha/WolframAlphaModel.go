package wolfram_alpha

type SubPod struct {
	Title string `json:"title"`
	Img   struct {
		Src    string `json:"src"`
		Title  string `json:"title"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"img"`
	PlainText string `json:"plaintext"`
}

type Pod struct {
	Title      string   `json:"title"`
	ID         string   `json:"id"`
	Error      bool     `json:"error"`
	NumSubPods int      `json:"numsubpods"`
	SubPods    []SubPod `json:"subpods"`
}

type Response struct {
	QueryResult struct {
		Success bool  `json:"success"`
		Error   bool  `json:"error"`
		NumPods int   `json:"numpods"`
		Pods    []Pod `json:"pods"`
	} `json:"queryresult"`
}
