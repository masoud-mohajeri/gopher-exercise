package apiCall

type Post struct {
	UserId uint32 `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ApiCall() {
	//NetHttp()
	Heimdall()
}
