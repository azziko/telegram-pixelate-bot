package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	sendMessageMethod = "sendMessage"
	sendPhotoMethod   = "sendPhoto"
	getFileMethod     = "getFile"
	receiveFileMethod = "file"
)

type Client struct {
	Host     string
	BasePath string
	client   http.Client
}

func NewClient(Host string, token string) *Client {
	return &Client{
		Host:     Host,
		BasePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) GetFile(fileID string) ([]byte, error) {
	query := url.Values{}
	query.Add("file_id", fileID)

	url := url.URL{
		Scheme: "https",
		Host:   c.Host,
		Path:   path.Join(c.BasePath, getFileMethod),
	}

	body, err := c.doRequest(query, url)
	if err != nil {
		return nil, fmt.Errorf("failed to getFile: %v", err)
	}

	return body, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	url := url.URL{
		Scheme: "https",
		Host:   c.Host,
		Path:   path.Join(c.BasePath, sendMessageMethod),
	}

	_, err := c.doRequest(query, url)
	if err != nil {
		return fmt.Errorf("failed to sendMessage: %v", err)
	}

	return nil
}

func (c *Client) SendPhoto(chatID int, filepath string, caption string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("photo", filepath)
	query.Add("caption", caption)

	url := url.URL{
		Scheme: "https",
		Host:   c.Host,
		Path:   path.Join(c.BasePath, sendPhotoMethod),
	}

	_, err := c.doRequest(query, url)
	if err != nil {
		return fmt.Errorf("failed to sendPhoto: %v", err)
	}

	return nil
}

func (c *Client) DownloadImage(filepath string) ([]byte, error) {
	query := url.Values{}
	url := url.URL{
		Scheme: "https",
		Host:   c.Host,
		Path:   path.Join(receiveFileMethod, c.BasePath, filepath),
	}

	body, err := c.doRequest(query, url)
	if err != nil {
		return nil, fmt.Errorf("DownloadImage failed: %v", err)
	}

	return body, nil
}

func (c *Client) doRequest(query url.Values, u url.URL) ([]byte, error) {
	var errMessage = "doRequest failed"

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errMessage, err)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errMessage, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errMessage, err)
	}

	return body, nil
}
