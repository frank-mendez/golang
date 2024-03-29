package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, header http.Header) (response *http.Response, err error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}
	return client.Do(request)
}
