PushOnAuth
============

A service that will send you a "Push" notification when any kind of auth happens on a Linux system.

![Say it with pixels](http://i.imgur.com/5VSbFC9.png)

## Don't want to go though the setup with go?

I can relate with that. Just download a precompiled version for your OS/Arch in the releases and run like so

`./PushOnAuth`

After it makes you a first time config, edit that config and test that it works by running it again.

After that you can run it "forever" by doing

```bash
nohup ./PushOnAuth &
```

There are also two scripts included for initd and systemd that allow you to easilly daemonize the process.

##Sample setup

```json
{
   "Notifications":{
      "PushOver":{
         "UserToken":"token",
         "AppToken":"token"
      },
      "PushAlot":{
         "Token":"token"
      },
      "Pushjet": {
         "Secret":"secret"
      }
   },
   "Watches":[
      {
         "Path":"/var/log/auth.log",
         "TriggerWords":[
            "Accepted publickey",
            "Accepted password"
         ]
      }
   ]
}

```

##Note
Please notify me if any push services that you would like to see implemented are not implemented yet
