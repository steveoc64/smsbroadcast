package smsbroadcast

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type SMS struct {
	API         string
	Username    string
	Password    string
	Destination string
	Source      string
}

func New(api, user, pass, dest, source string) *SMS {
	return &SMS{
		API:         api,
		Username:    user,
		Password:    pass,
		Destination: dest,
		Source:      source,
	}
}

func (s *SMS) Send(msg string) (string, error) {
	rsp, err := http.PostForm(s.API, url.Values{
		"username": {s.Username},
		"password": {s.Password},
		"to":       {s.Destination},
		"from":     {s.Source},
		"message":  {msg},
		"maxsplit": {"2"},
	})
	if err != nil {
		return "error", err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return "error", err
	}
	return string(body), nil
}
