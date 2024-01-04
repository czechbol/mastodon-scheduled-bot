# Mastodon Bot

This is a very simple bot for Mastodon, written in Go. It uses the `gronx` package for scheduling posts and the `go-mastodon` package for interacting with the Mastodon API.

It is currently only capable of posting the same message whenever a time defined by CRON_SCHEDULE occurs

## Features

- Posts the same message to Mastodon at scheduled interval
- Configurable via environment variables

## Environment Variables

| Variable        | Description                              | Example                   |
| --------------- | ---------------------------------------- | ------------------------- |
| `SERVER`        | The Mastodon instance URL                | https://mastodon.social   |
| `CLIENT_ID`     | The client ID for your Mastodon app      | your_client_id            |
| `CLIENT_SECRET` | The client secret for your Mastodon app  | your_client_secret        |
| `ACCESS_TOKEN`  | The access token for your Mastodon app   | your_access_token         |
| `POST_TEXT`     | The text to be posted by the bot         |  Hello, Mastodon!         |
| `TZ`            | The timezone for scheduling posts        | America/New_York          |
| `CRON_SCHEDULE` | The cron schedule for posts              | * * * * *                 |

## Usage

1. Clone the repository
2. Set the environment variables
3. Run `go run main.go`

## Example

```bash
export SERVER="https://mastodon.social"
export CLIENT_ID="your_client_id"
export CLIENT_SECRET="your_client_secret"
export ACCESS_TOKEN="your_access_token"
export POST_TEXT="Hello, Mastodon!"
export TZ="America/New_York"
export CRON_SCHEDULE="* * * * *"
go run main.go
```

## Docker Compose

You can also run this bot using Docker Compose. Here's an example of how to do it:

1. Create a `docker-compose.yml` file in your project root with the content from `default.docker-compose.yml`
2. Replace the placeholder environment variables
3. Run the command `docker compose up` in your terminal
