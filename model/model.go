package model

// HTTPError
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError
type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

// HTTPError
type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

type Character struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Characters struct {
	Id int `json:"id"`
}

type Response struct {
	Code int    `json:"code"`
	Status string  `json:"status"`
	Data Data `json:"data"`

}

type Data struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Character `json:"results"`
}

type Output struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Characters `json:"results"`
}

type Config struct {
	Redis struct {
		Url string `yaml:"url"`
	} `yaml:"redis"`
	Marvel struct {
		Timestamp string `yaml:"timestamp"`
		Apikey string `yaml:"apikey"`
		Privatekey string `yaml:"privatekey"`
		Limit string `yaml:"limit"`
	} `yaml:"marvel"`
}
