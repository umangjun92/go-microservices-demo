package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enableMocks = true
}

func StopMockups() {
	enableMocks = false
}

func AddMockup(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		// return local mock withour calling any external reosurce
		mock := mocks[getMockId("POST", url)]
		// jsonM, _ := json.Marshal(mock)
		// fmt.Printf("all mocks %s", jsonM)
		if mock == nil {
			return nil, errors.New("no mockup found for given rqst")
		}
		// s, _ := json.Marshal(mock.Response)
		// fmt.Printf("mocks %s", s)
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	fmt.Printf("in here nOW >>>>> %s", string(jsonBytes))

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	return client.Do(request)
}
