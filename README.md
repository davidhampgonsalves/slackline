# Slackline
> Slackline lets you post updates/messages to slack with minimal disruption from the cmd line.

[[https://github.com/davidhampgonsalves/slackline/blob/master/slackline.gif|width:600px]]

## Usage
**Post a message.** `slackline Hello world from slackline!`

**Setup slack token and channel via prompts. Settings saved to ~/.slackline.yml** `slackline --init`

**Set your slack token / channel.** `slackline -t <slack-token> -c "#channel" message to post`

**Flags can be saved to ~/.slackline.yml. Your token will be encrypted.**
``` sh
slackline -s <slack-token> -c "#channel" --save message to post
```

## Setup 
```sh 
go get https://github.com/davidhampgonsalves/slackline
```
