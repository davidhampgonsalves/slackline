package main

import (
	"fmt"
	"github.com/davidhampgonsalves/slackline/conf"
	"github.com/davidhampgonsalves/slackline/slack"
)

/*

flags
 -c = channel names (slack channel names are single word seperated by comma)
 -t = set token
 --save = persist flags to config
 -i = init, ask for channels and token
handle piping input?

TODO:
  validate flags / get message
  get cred token, user name from .slk-line or -t
    if no cred, display link to get token from slack
      - allow user to enter token which will be saved to .slk-line encrypted with mac, cpu, etc as key
  make request
  display errors

*/

func main() {
	config, err := conf.Load()
	fmt.Println(*config)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = slack.JoinChannel(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = slack.PostMessage(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.Save {
		conf.Save(config)
	}
}
