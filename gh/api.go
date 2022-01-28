package gh

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// GithubApiEndPoint is the base URL for the GitHub Graphql API.
const GithubApiEndPoint = "https://api.github.com/graphql"

type Client struct {
	User
	Token string
}

// NewClient create new Client
func NewClient(token string) (c *Client, err error) {
	c = &Client{Token: token}
	u, err := c.ViewerLogin()
	if err != nil {
		return c, err
	}
	c.User = u.Data.Viewer
	return c, err
}

// Request send request to github
func (c *Client) Request(s string) (body []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", GithubApiEndPoint, strings.NewReader(s))
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

// ViewerLogin get user login info
func (c *Client) ViewerLogin() (data UserData, err error) {
	b, _ := json.Marshal(Graphql{Query: GetUserData})
	body, err := c.Request(string(b))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if len(data.Errors) != 0 {
		return data, fmt.Errorf(data.Errors[0].Message)
	}
	if data.Message != "" {
		return data, fmt.Errorf(data.Message)
	}
	return data, err
}

// GetUserStatus get user status
func (c *Client) GetUserStatus(username string) (data StatusData, err error) {
	b, err := json.Marshal(Graphql{Query: fmt.Sprintf(GetUserStatusQuery, username)})
	if err != nil {
		return data, err
	}
	body, err := c.Request(string(b))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if len(data.Errors) != 0 {
		return data, fmt.Errorf(data.Errors[0].Message)
	}
	if data.Message != "" {
		return data, fmt.Errorf(data.Message)
	}
	return data, err
}

// ClearUserStatus clear user status
func (c *Client) ClearUserStatus() (err error) {
	b, _ := json.Marshal(Graphql{Query: ClearUserStatusQuery})
	body, err := c.Request(string(b))
	var data StatusData
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if len(data.Errors) != 0 {
		return fmt.Errorf(data.Errors[0].Message)
	}
	if data.Message != "" {
		return fmt.Errorf(data.Message)
	}
	return err
}

// SetUserStatus set user status
func (c *Client) SetUserStatus(emoji string, message string) (err error) {
	b, _ := json.Marshal(Graphql{Query: SetUserStatusQuery, Variables: Status{Emoji: Emojis.Emoji2Shortname(emoji), Message: message}})
	body, err := c.Request(string(b))
	var data StatusData
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if len(data.Errors) != 0 {
		return fmt.Errorf(data.Errors[0].Message)
	}
	if data.Message != "" {
		return fmt.Errorf(data.Message)
	}
	return err
}
