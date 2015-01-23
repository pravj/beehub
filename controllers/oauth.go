package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

const (
	TOKEN_ENDPOINT string = "https://github.com/login/oauth/access_token"
)

type OauthController struct {
	beego.Controller
}

type payload struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type response struct {
	AccessToken string `json:"access_token"`
}

func (this *OauthController) ParseCode() {
	token := AccessToken(this.GetString("code"), beego.AppConfig.String("client_id"), beego.AppConfig.String("client_secret"))

	this.Data["Website"] = token
	this.Data["Email"] = "hackpravj@gmail.com"
	this.TplNames = "index.tpl"
}

func AccessToken(Code, ClientId, ClientSecret string) string {
	payloadJson, _ := json.Marshal(payload{Code, ClientId, ClientSecret})
	payloadReader := bytes.NewReader(payloadJson)

	req, _ := http.NewRequest("POST", TOKEN_ENDPOINT, payloadReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var resp response
	json.Unmarshal(body, &resp)

	return resp.AccessToken
}
