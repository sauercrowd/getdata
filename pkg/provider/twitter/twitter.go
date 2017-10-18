package twitter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Twitter struct {
	token string
}

//Login with API key and API secret
func Login(key, secret string) (*Twitter, error) {
	req, err := createRequest(key, secret)
	if err != nil {
		return nil, err
	}
	token, err := getToken(req)
	return &Twitter{token: token}, nil
}

// createRequest encods key and secret and creates the correct request to get tweets
func createRequest(key, secret string) (*http.Request, error) {
	auth := string(base64.StdEncoding.EncodeToString([]byte(key + ":" + secret)))
	bodyStr := []byte("grant_type=client_credentials")
	u := "https://api.twitter.com/oauth2/token"
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, fmt.Errorf("Could not create request for twitter login: %v", err)
	}
	req.Header.Set("User-Agent", "JonasDataCollector")
	req.Header.Set("grant_type", "client_credentials")
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	return req, nil
}

//get a token
func getToken(req *http.Request) (string, error) {
	var client http.Client
	//do request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Could not send request to get login token for twitter: %v", err)
	}

	//read body and unmarshal json
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Could not read response to get login token for twitter: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("Could not unmarshal response to get login token for twitter: %v", err)
	}

	//check if an error occured
	if val, ok := data["errors"]; ok {
		return "", fmt.Errorf("Could not get token for twitter: %v", val)
	}

	// check if the access token field is present
	if token, ok := data["access_token"]; ok {
		tokenStr, ok := token.(string)
		if ok {
			return tokenStr, nil
		}
		return "", fmt.Errorf("Could not convert twitter token to string")
	}
	return "", fmt.Errorf("Could not find token at twitter response")
}
