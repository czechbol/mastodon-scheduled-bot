package main

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/adhocore/gronx"
	"github.com/adhocore/gronx/pkg/tasker"
	"github.com/mattn/go-mastodon"
)

var (
	POST_TEXT     = ""
	TZ            = ""
	CRON_SCHEDULE = ""
)

type MastodonBot struct {
	client *mastodon.Client
}

func NewMastodonBot() *MastodonBot {
	server := os.Getenv("SERVER")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")

	log.Info("Server:", server)
	log.Info("Client ID:", clientID)
	log.Info("Client Secret:", clientSecret)
	log.Info("Access Token:", accessToken)
	return &MastodonBot{
		client: mastodon.NewClient(&mastodon.Config{
			Server:       server,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			AccessToken:  accessToken,
		}),
	}
}

// a function that creates a new mastodon post for a bot with given credentials
func (b *MastodonBot) Post(ctx context.Context, text string) (*mastodon.Status, error) {
	// post the status
	return b.client.PostStatus(ctx, &mastodon.Toot{
		Status: text,
	})
}

func main() {
	// logging using logrus
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	POST_TEXT = os.Getenv("POST_TEXT")
	TZ = os.Getenv("TZ")
	CRON_SCHEDULE = os.Getenv("CRON_SCHEDULE")
	log.Info("Post text: ", POST_TEXT)
	log.Info("Timezone: '", TZ, "'")
	log.Info("Cron trigger: ", CRON_SCHEDULE)

	err := App()

	log.Error(err)
}

func App() error {
	taskr := tasker.New(tasker.Option{
		Verbose: false,
		Tz:      TZ,
	})

	nextRun, err := NextRun(true)
	if err != nil {
		return fmt.Errorf("Error getting next run time: %v", err)
	}
	log.Info("Next run: ", nextRun)

	// add task to run every minute
	taskr.Task(CRON_SCHEDULE, func(ctx context.Context) (int, error) {
		_, err := postToMastodonJob(ctx)
		if err != nil {
			return 1, err
		}

		nextRun, err := NextRun(true)
		if err != nil {
			return 1, err
		}
		log.Info("Next run: ", nextRun)
		return 0, nil
	})

	// finally run the tasker, it ticks sharply on every minute and runs all the tasks due on that time!
	// it exits gracefully when ctrl+c is received making sure pending tasks are completed.
	taskr.Run()
	return nil
}

func postToMastodonJob(ctx context.Context) (*mastodon.Status, error) {
	bot := NewMastodonBot()
	log := log.WithFields(log.Fields{
		"function": "postToMastodonJob",
	})
	log.Info("Posting to Mastodon: ", POST_TEXT)
	result, err := bot.Post(ctx, POST_TEXT)
	if err != nil {
		log.WithError(err).Error("Error posting to Mastodon")
		return nil, fmt.Errorf("Error posting to Ma@caustodon: %v", err)
	}
	log.Info("Successfully posted to Mastodon: ", result.Content)
	return result, nil
}

func NextRun(allowCurrent bool) (time.Time, error) {
	return gronx.NextTick(CRON_SCHEDULE, allowCurrent)
}
