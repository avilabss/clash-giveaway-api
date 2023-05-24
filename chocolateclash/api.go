package chocolateclash

import "fmt"

type Api struct {
	BaseUrl string
}

func (api *Api) Init(league string) {
	api.BaseUrl = fmt.Sprintf("https://%s.chocolateclash.com/cc_n", league)
}
