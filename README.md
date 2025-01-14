# Larry 🐦
[![Go](https://github.com/ezeoleaf/larry/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ezeoleaf/larry/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ezeoleaf/larry)](https://goreportcard.com/report/github.com/ezeoleaf/larry)
[![MIT License](https://img.shields.io/github/license/ezeoleaf/larry?style=flat-square)](https://github.com/ezeoleaf/larry/blob/main/LICENSE)

Larry is a Golang cli bot that tweets random Github repositories.

## Disclaimer

I hold no liability for what you do with this bot or what happens to you by using this bot. Abusing this bot *can* get you banned from Twitter, so make sure to read up on [proper usage](https://support.twitter.com/articles/76915-automation-rules-and-best-practices) of the Twitter API.

## Running bots

- [GolangRepos](https://twitter.com/GolangRepos): Tweets repositories from Github that contain the "golang" topic
- [RustRepos](https://twitter.com/RustRepos): Tweets repositories from Github that contain the "rust" topic
- [MLRepositories](https://twitter.com/MLRepositories): Tweets repositories from Github that contain the "machine-learning" topic
- [CryptoRepos](https://twitter.com/CryptoRepos): Tweets repositories from Github that contain the "crypto" topic

## Installation

You can install Larry by cloning the repo and using `go install`

```bash
git clone https://github.com/ezeoleaf/larry.git
cd larry
go install
```

Or you can run it on the go
```bash
git clone https://github.com/ezeoleaf/larry.git
cd larry
go run . [options]
```

## Usage

### Configuring the bot

Before running the bot, you must first set it up so it can connect to Github and Twitter API. Create a config.json (or rename the config.example.json) file and fill in the following information:
```json
{
    "github_access_token": "xxxxx",
    "twitter_consumer_key": "xxxxx",
    "twitter_consumer_secret": "xxxxx",
    "twitter_access_token": "xxxxx",
    "twitter_access_secret": "xxxxx"
}
```

For generating Github access token you can follow this [guide](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)

For getting Twitter keys and secrets you can follow this [guide](https://developer.twitter.com/en/docs/twitter-api/getting-started/guide)

### Running the bot

To run the bot, you have two ways.

If you have installed the bot, you can run it using
  `larry [options]`

If you want to run it without installing it globally you can use
  `go run . [options]`

Example:

`larry -h`

As a response you will see the entire options available

```
NAME:
   larry - Twitter bot that tweets random repositories

USAGE:
   larry [global options] command [command options] [arguments...]

AUTHOR:
   Ezequiel Olea figueroa <ezeoleaf@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --topic value, -t value         topic for searching repos
   --lang value, -l value          language for searching repos
   --config value, -c value        path to config file (default: "./config.json")
   --time value, -x value          periodicity of tweet in minutes (default: 15)
   --cache value, -r value         size of cache for no repeating repositories (default: 50)
   --hashtag value, --ht value     list of comma separated hashtags
   --tweet-language, --tl          bool for allowing twetting the language of the repo (default: false)
   --safe-mode, --sf               bool for safe mode. If safe mode is enabled, no repository is published (default: false)
   --provider value, --pr value    provider where publishable content comes from (default: "github")
   --publisher value, --pub value  list of comma separared publishers (default: "twitter")
   --help, -h                      show help (default: false)
```

For running the bot, the command will depend on whatever you want to tweet, but, for tweeting about React repositories every 30 minutes, you could use

&nbsp;&nbsp;`larry --topic react --config "path_to_react_config.json" --time 30`

For running the bot for Rust tweets every 15 minutes

&nbsp;&nbsp;`larry --lang rust --config "path_to_rust_config.json" --time 15`


## Have questions? Need help with the bot?

If you're having issues with or have questions about the bot, [file an issue](https://github.com/ezeoleaf/larry/issues) in this repository so anyone can get back to you.

Or feel free to contact me <ezeoleaf@gmail.com> :)
