package checkout

import (
	"io"
	"net/http"
)

type HttpClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body any) ([]byte, error)
}

type HttpNativeAdapter struct {
}

func (hna *HttpNativeAdapter) Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (hna *HttpNativeAdapter) Post(url string, body any) ([]byte, error) {
	return nil, nil
}
