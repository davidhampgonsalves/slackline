# Slackline
> Slackline lets you post updates/messages to slack with minimal disruption from the cmd line.

<p style="text-align:center">
  <img src="https://github.com/davidhampgonsalves/slackline/blob/master/slackline.gif" width=1200px> 
</p>

## Usage
**Post a message.**
``` sh
slackline Hello world from slackline!
```

**Setup slack token and channel via prompts. Settings saved to ~/.slackline.yml**
``` sh
slackline --init
```

**Set your slack token / channel.**
``` sh
slackline -t <slack-token> -c "#channel" message to post
```

**Flags can be saved to ~/.slackline.yml. Your token will be encrypted.**
``` sh
slackline -s <slack-token> -c "#channel" --save message to post
```

## Setup 

###Homebrew
```sh
brew install https://raw.githubusercontent.com/davidhampgonsalves/slackline/master/slackline.rb
```

##GO(golang)
```sh 
go get github.com/davidhampgonsalves/slackline
```
