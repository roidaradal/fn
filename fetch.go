package fn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchObject[T any](req *http.Request) (*T, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response not OK: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var item T
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response to struct: %w", err)
	}
	return &item, nil
}
