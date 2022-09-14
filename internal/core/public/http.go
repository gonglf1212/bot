package public

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func JsonGet(url string) (rebyte []byte, err error) {
	client := &http.Client{}
	fmt.Println("-dd---", url)
	response, err := client.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return rebyte, errors.New(fmt.Sprintln("StatusCode is ", response.StatusCode))
	}

	if rebyte, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	return
}
