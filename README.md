# twitchSendMessage

Send anything to twitch chat!

## Installation
```bash
go install github.com/ottojo/twitchSendMessage
```

## Usage
1. Obtain twitch oauth token: https://twitchapps.com/tmi/
2. Send a message:
```bash
echo "Hello!" | twitchSendMessage -o oauth:123... -c someTwitchChannel -u yourUserName
```
