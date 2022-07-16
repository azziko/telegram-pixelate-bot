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
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) GetFile(fileID int) ([]byte, error) {
	query := url.Values{}
	query.Add("file_id", strconv.Itoa(fileID))

	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, getFileMethod),
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
		Host:   c.host,
		Path:   path.Join(c.basePath, sendMessageMethod),
	}

	_, err := c.doRequest(query, url)
	if err != nil {
		return fmt.Errorf("failed to sendMessage: %v", err)
	}

	return nil
}

func (c *Client) SendPhoto(chatID int, fileID int, caption string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("photo", strconv.Itoa(fileID)) //figure out how to pass multipart-data
	query.Add("caption", caption)

	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, sendPhotoMethod),
	}

	_, err := c.doRequest(query, url)
	if err != nil {
		return fmt.Errorf("failed to sendPhoto: %v", err)
	}

	return nil
}

func (c *Client) DownloadImage(filePath string) ([]byte, error) {
	query := url.Values{}
	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(receiveFileMethod, c.basePath, filePath),
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
