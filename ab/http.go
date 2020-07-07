package ab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(method, url string) (string, error) {
	client := &http.Client{}
	method = strings.ToUpper(method)
	resquest, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("new request error ", err.Error())
		return "", err
	}
	resp, err := client.Do(resquest)
	if err != nil {
		fmt.Println("client do error ", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
