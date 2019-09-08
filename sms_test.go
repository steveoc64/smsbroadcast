package smsbroadcast

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	// get the params to setup the server
	config := &SMS{}
	c, err := ioutil.ReadFile("config.json")
	assert.Nil(t, err)
	assert.Nil(t, json.Unmarshal(c, config))
	spew.Dump(config)

	// setup an SMSer
	sms := New(config.API, config.Username, config.Password, config.Destination, config.Source)
	rsp, err := sms.Send("this is a test message")
	assert.Nil(t, err)

	t.Log("Response", rsp)
}
