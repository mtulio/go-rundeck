package rundeck

import (
	"bytes"
	"encoding/json"
	"fmt"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Token is a token
type Token struct {
	responses.TokenResponse
}

// TokenOption is a type for options in creating new tokens
type TokenOption func(*Token) error

// TokenDuration is an option for setting the duration of a new token
func TokenDuration(duration string) TokenOption {
	return func(t *Token) error {
		t.Duration = duration
		return nil
	}
}

// TokenRoles is an option to set the roles for a new token
func TokenRoles(roles ...string) TokenOption {
	return func(t *Token) error {
		t.Roles = append(t.Roles, roles...)
		return nil
	}
}

// GetTokens gets all tokens
func (c *Client) GetTokens() ([]*Token, error) {
	tokens := make([]*Token, 0)
	data, err := c.httpGet("tokens", requestJSON())
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(data, &tokens)
	return tokens, jsonErr
}

// GetUserTokens gets the api tokens for a user
func (c *Client) GetUserTokens(user string) ([]*Token, error) {
	data, err := c.httpGet("tokens/"+user, requestJSON())
	if err != nil {
		return nil, err
	}
	tokens := make([]*Token, 0)
	jsonErr := json.Unmarshal(data, &tokens)
	return tokens, jsonErr
}

// GetToken gets a token
func (c *Client) GetToken(tokenID string) (*Token, error) {
	data, err := c.httpGet("token/"+tokenID, requestJSON())
	if err != nil {
		return nil, err
	}
	token := &Token{}
	jsonErr := json.Unmarshal(data, &token)
	return token, jsonErr
}

// CreateToken creates a token
func (c *Client) CreateToken(u string, opts ...TokenOption) (*Token, error) {
	tokenRequest := &Token{}
	for _, opt := range opts {
		if err := opt(tokenRequest); err != nil {
			return nil, err
		}
	}
	if len(tokenRequest.Roles) == 0 {
		tokenRequest.Roles = []string{"*"}
	}
	newToken, newErr := json.Marshal(tokenRequest)
	if newErr != nil {
		return nil, newErr
	}
	url := fmt.Sprintf("tokens/%s", u)
	data, err := c.httpPost(url, requestJSON(), withBody(bytes.NewReader(newToken)))
	if err != nil {
		return nil, err
	}
	token := &Token{}
	jsonErr := json.Unmarshal(data, &token)
	return token, jsonErr
}

// DeleteToken deletes a token
func (c *Client) DeleteToken(token string) error {
	url := fmt.Sprintf("token/%s", token)
	return c.httpDelete(url, requestJSON())
}