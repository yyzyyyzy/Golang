package main

import (
	"encoding/json"
	"fmt"
)

type ImCloud struct {
	Addr                   string `json:"addr"`
	Orig                   string `json:"orig"`
	Name                   string `json:"name"`
	WsUri                  string `json:"ws_uri"`
	HeartbeatCheckInterval int    `json:"heartbeat_check_interval"`
}
type ImServer struct {
	Addr  string `json:"addr"`
	Orig  string `json:"orig"`
	Name  string `json:"name"`
	WsUri string `json:"ws_uri"`
}
type Keys struct {
	ImClient string `json:"im-client"`
	ImServer string `json:"im-server"`
	Jwt      string `json:"jwt"`
}
type Redis struct {
	Addr string `json:"addr"`
}

type Config struct {
	*ImCloud  `json:"im-cloud"`
	ImServerS map[string]*ImServer `json:"im-server-s"`
	*Keys     `json:"keys"`
	*Redis    `json:"redis"`
}

func main() {
	jsonstr := `{
   "im-cloud": {
     "addr": "127.0.0.1:8000",
     "orig": "http://127.0.0.1",
     "name": "im-cloud",
     "ws_uri": "/ws",
     "heartbeat_check_interval": 5
   },
   "im-server-s": {
     "im-server-7000": {
       "addr": "127.0.0.1:7000",
       "orig": "http://127.0.0.1",
       "name": "im-server-7000",
       "ws_uri": "/ws"
     },
     "im-server-9000": {
       "addr": "127.0.0.1:9000",
       "orig": "http://127.0.0.1",
       "name": "im-server-9000",
       "ws_uri": "/ws"
     }
   },
   "redis": {
     "addr": "127.0.0.1:6379"
   },
   "keys": {
     "im-server": "im-server",
     "jwt": "imstar",
     "im-client": "im-client"
   }
}`
	Cfg := new(Config)
	err := json.Unmarshal([]byte(jsonstr), Cfg)
	fmt.Println(Cfg.ImServerS["im-server-9000"].Addr)
	if err != nil {
		fmt.Println(jsonstr)
		fmt.Println(Cfg)
		fmt.Println("===>>> imstart config set err : ", err)
	}
}
