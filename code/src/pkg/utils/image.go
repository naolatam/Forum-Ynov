package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchImage(imageURL string) ([]byte, error) {
	// Make an HTTP GET request to fetch the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: received status code %d", resp.StatusCode)
	}

	// Read the image data as a blob
	imageBlob, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %v", err)
	}

	return imageBlob, nil
}

func ConvertBytesToBase64(blob []byte, mime string) string {
	encoded := base64.StdEncoding.EncodeToString(blob)
	return fmt.Sprintf("data:%s;base64,%s", mime, encoded)
}

func GetDefaultAvatar() []byte {
	file, err := os.Open("internal/frontEnd/static/imgs/forums_background.jpg")
	if err != nil {
		return nil
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}
