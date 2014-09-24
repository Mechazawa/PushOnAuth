package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var cfgfile string = "./pushonauthcfg.json"

type GConfig struct {
	Notifications Notifiers
	Watches   []WatchFile
}

type WatchFile struct {
	Path         string
	TriggerWords []string
}

type Notifiers struct {
	PushOver NPushOver
	PushAlot NPushAlot
	Azise NAzise
}

/* Notifiers */
type NPushOver struct {
	UserToken string
	AppToken  string
}

type NPushAlot struct {
	Token string
}

type NAzise struct {
	Secret string
}

func GetDefaultConfig() GConfig {
	var tfg GConfig
	tfg.Watches = make([]WatchFile, 0)
	defaultwatch := WatchFile{
		Path: "/var/log/auth.log",
		TriggerWords: []string{
			"Accepted publickey",
			"Accepted password",
		},
	}
	tfg.Watches = append(tfg.Watches, defaultwatch)

	var nfr Notifiers
	var pusho NPushOver
	var pusha NPushAlot
	var azise NAzise

	pusho.UserToken = "token"
	pusho.AppToken = "token"
	nfr.PushOver = pusho

	pusha.Token = "token"
	nfr.PushAlot = pusha

	azise.Secret = "secret"
	nfr.Azise = azise

	tfg.Notifications = nfr
	return tfg
}

func CheckIfResetConfig(args []string) {
	if len(args) == 2 {
		if args[1] == "reset" {
			e := os.Remove(cfgfile)
			if e != nil {
				log.Fatal("Could not remove current config file. Permissions issue?")
			}
			Default := GetDefaultConfig()
			out, e := json.Marshal(Default)
			e = ioutil.WriteFile(cfgfile, out, 600)
			if e != nil {
				log.Fatal("cannot open settings file :(")
			}
			log.Fatal("Built config file. please fill it in.")
		}
	}
}

func GetCFG() GConfig {
	b, e := ioutil.ReadFile(cfgfile)
	tfg := GetDefaultConfig()
	if e != nil {
		out, e := json.Marshal(tfg)
		e = ioutil.WriteFile(cfgfile, out, 600)
		if e != nil {
			log.Fatal("cannot open settings file :(")
		}
		log.Fatal("Built config file. please fill it in.")
	}

	e = json.Unmarshal(b, &tfg)
	if e != nil {
		log.Fatalf("Could not parse config settings. You may have to remove %s", cfgfile)
	}

	var nfc = tfg.Notifications
	if !CanPushOver(nfc) && !CanPushAlot(nfc) {
		log.Fatalf("Please fill in at least one push service in %s", cfgfile)
	}

	return tfg
}
