package main

import (
	"net/http"
	"net/url"
)

func Push(ntf Notifiers, Title string, Message string) {
	if CanPushOver(ntf) {
		SendPushOver(Title, Message, ntf.PushOver.AppToken, ntf.PushOver.UserToken)
	}
	if CanPushAlot(ntf) {
		SendPushAlot(Title, Message, ntf.PushAlot.Token)
	}
	if CanAzise(ntf) {
		SendAzise(Title, Message, ntf.Azise.Secret)
	}
}

func IsTokenSet(token string) bool {
	return token != "" && token != "token" && token != "secret"
}


func CanPushOver(ntf Notifiers) bool {
	return IsTokenSet(ntf.PushOver.UserToken) && IsTokenSet(ntf.PushOver.AppToken)
}
func SendPushOver(Title string, Message string, Token string, User string) {
	http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{
			"token":   {Token},
			"user":    {User},
			"message": {Message},
			"Title":   {Title},
		})
}

func CanAzise(ntf Notifiers) bool {
	return IsTokenSet(ntf.Azise.Secret)
}
func SendAzise(Title string, Message string, Secret string) {
        http.PostForm("https://api.azise.net/message",
                url.Values{
                        "secret":  {Secret},
                        "message": {Message},
                        "Title":   {Title},
                })
}

func CanPushAlot(ntf Notifiers) bool {
	return IsTokenSet(ntf.PushAlot.Token)
}
func SendPushAlot(Title string, Body string, Token string) {
	http.PostForm("https://pushalot.com/api/sendmessage",
		url.Values{
			"AuthorizationToken": {Token},
			"Body": {Body},
			"Title": {Title},
		})
}
