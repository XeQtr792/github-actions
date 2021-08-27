package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Client *http.Client
}



func (a *Api) GetStuff(uURL string) {
	resp, err := a.Client.Get(uURL) // fixed
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
