package main

import (
	"fmt"
	"github.com/davidhampgonsalves/slackline/conf"
)

/*

flags
 -c = channel names (slack channel names are single word seperated by comma)
 -t = set token
 --save = persist flags to config
 -i = init, ask for channels and token
handle piping output

TODO:
  validate flags / get message
  get cred token, user name from .slk-line or -t
    if no cred, display link to get token from slack
      - allow user to enter token which will be saved to .slk-line encrypted with mac, cpu, etc as key
  make request
  display errors

*/

func main() {
	conf, err := conf.Load()

	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("hello, world %+v", conf)
}
