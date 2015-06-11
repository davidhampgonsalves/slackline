# Slackline
> Slackline lets you post updates/messages to slack with minimal disruption from the cmd line.

![gif](https://github.com/davidhampgonsalves/slackline/blob/master/slackline.gif)

## Usage
** Post a message. **
``` sh
slackline Hello world from slackline!
```

**Setup your slack token and channel via prompts. Settings are save at ~/.slackline.yml**
``` sh
slackline --init
```

**Set your slack token / channel.**
``` sh
slackline -t <slack-token> -c #channel message to post
```

**Flags can be saved to ~/.slackline.yml. Your token will be encrypted.**
``` sh
slackline -s <slack-token> -c #channel --save message to post
```

## Setup 
```sh 
go get https://github.com/davidhampgonsalves/slackline
```
