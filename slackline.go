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
  if no cred or channels, display link to get token from slack
  message format
*/

func main() {
	config, err := conf.Load()
	if err != nil {
		fmt.Print(err)
    return
	}

  errs := config.Valid() 
  if len(errs) > 0 {
    for _, err = range errs {
      fmt.Println(err)
    }
    return
  }

	if config.Message != "" {
		err = slack.JoinChannel(config)
		if err != nil {
			fmt.Println(err)
      return
		}

		err = slack.PostMessage(config)
		if err != nil {
			fmt.Println(err)
		}
	}

  if config.Save {
		err = conf.Save(config)
    if err == nil { fmt.Println("Your settings were saved, happy slacklining.") }
  } else if config.Init {
    err = conf.Init(config)
    if err == nil { fmt.Println("New settings were saved, happy slacklining.") }
  } else {
		err = fmt.Errorf("You must provide a message to post.")
	}

  if err != nil {
   fmt.Println(err) 
  }
}

