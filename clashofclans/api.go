package clashofclans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Api struct {
	BaseUrl string
	JWT     string
}

func (api *Api) GetClanInfo(clanTag string) (*Clan, error) {
	if strings.HasPrefix(clanTag, "#") {
		clanTag = url.QueryEscape(clanTag)
	}

	uri := api.BaseUrl + "/clans/" + clanTag

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return nil, ErrFailedToBuildRequest
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.JWT))

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrRequestFailed.Error(), err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrFailedToReadResponseBody.Error(), err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status: %s\n%s", response.Status, bytes.NewBuffer(body).String())
	}

	var clan Clan

	err = json.Unmarshal(body, &clan)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrFailedToParseResponseBody.Error(), err)
	}

	return &clan, nil
}

func (api *Api) GetGoldPassInfo() (*GoldPass, error) {
	uri := api.BaseUrl + "/goldpass/seasons/current"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return nil, ErrFailedToBuildRequest
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.JWT))

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrRequestFailed.Error(), err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrFailedToReadResponseBody.Error(), err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status: %s\n%s", response.Status, bytes.NewBuffer(body).String())
	}

	var goldPass GoldPass

	err = json.Unmarshal(body, &goldPass)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrFailedToParseResponseBody.Error(), err)
	}

	return &goldPass, nil
}
