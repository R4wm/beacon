# beacon
simple app to curl base with ip and geo

---

## TODO
- Create Makefile to build and setup cron

## Usage
Client
```bash
ᚱ@cmd $ go get -u github.com/r4wm/beacon 
ᚱ@cmd $ cd $GOPATH/src/github.com/r4wm/beacon/cmd
ᚱ@cmd $ 
ᚱ@cmd $ go build -o $GOBIN/beacon 
ᚱ@cmd $ 
ᚱ@cmd $ 
ᚱ@cmd $ beacon -help
Usage of beacon:
  -hostname string
    	hostname or ip of where to send beacon
ᚱ@cmd $ beacon -hostname https://mintz5.com/index.html
ᚱ@cmd $ date
Sun 19 Jan 2020 08:43:27 PM PST
ᚱ@cmd $ 
```

Server
```
$ tail -1 /var/log/nginx/access.log
104.175.64.217 - - [20/Jan/2020:04:43:16 +0000] "GET /index.html HTTP/1.1" 200 905 "-" "Beacon:Glendora,California"
voidconf@web001:~$ 
```
