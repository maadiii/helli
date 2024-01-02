package event

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
)

var dir = "event_logs/"

func logEvent(ctx context.Context, event Event) {
	switch event.Action {
	case Create:
		logCreate(ctx, event)
	case Increase:
		logIncrease(ctx, event)
	case Decrease:
		logDecrease(ctx, event)
	case Transfer:
		logTransfer(ctx, event)
	}
}

func logCreate(ctx context.Context, event Event) {
	file, err := os.OpenFile(dir+event.Username, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("[ERROR] log file creation: %s", event.Username)

		return
	}

	opts := levelHandler(Create)
	logger := slog.New(slog.NewJSONHandler(file, opts))

	logger.Log(
		ctx,
		LevelCreate,
		fmt.Sprintf("create account with fund %f", event.Amount),
		"account_id", event.AccountID,
		"amount", event.Amount,
	)
}

func logIncrease(ctx context.Context, event Event) {
	file, err := os.OpenFile(dir+event.Username, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("[ERROR] log file reading: %s", event.Username)

		return
	}

	opts := levelHandler(Increase)
	logger := slog.New(slog.NewJSONHandler(file, opts))

	logger.Log(
		ctx,
		LevelIncrease,
		fmt.Sprintf("increase fund by %f", event.Amount),
		"account_id", event.AccountID,
		"amount", event.Amount,
	)
}

func logDecrease(ctx context.Context, event Event) {
	file, err := os.OpenFile(dir+event.Username, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("[ERROR] log file reading: %s", event.Username)

		return
	}
	defer file.Close()

	opts := levelHandler(Decrease)
	logger := slog.New(slog.NewJSONHandler(file, opts))

	logger.Log(
		ctx,
		LevelDecrease,
		fmt.Sprintf("decrease fund by %f", event.Amount),
		"account_id", event.AccountID,
		"amount", event.Amount,
	)
}

func logTransfer(ctx context.Context, event Event) {
	file, err := os.OpenFile(dir+event.SourceUsername, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("[ERROR] reading source log file to transfer: %s", event.SourceUsername)

		return
	}

	opts := levelHandler(Transfer)
	logger := slog.New(slog.NewJSONHandler(file, opts))

	logger.Log(
		ctx,
		LevelTransfer,
		fmt.Sprintf("send fund by %f", event.Amount),
		"source_account_id", event.SourceAccountID,
		"destination_username", event.DestinationUsername,
		"destination_account_id", event.DestinationAccountID,
		"amount", event.Amount,
	)

	file.Close()

	file, err = os.OpenFile(dir+event.DestinationUsername, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("[ERROR] reading destination log file to transfer: %s", event.SourceUsername)

		return
	}
	defer file.Close()

	opts = levelHandler(receive)
	logger = slog.New(slog.NewJSONHandler(file, opts))

	logger.Log(
		ctx,
		levelReceive,
		fmt.Sprintf("receive fund by %f", event.Amount),
		"source_username", event.SourceUsername,
		"source_account_id", event.SourceAccountID,
		"destination_account_id", event.DestinationAccountID,
		"amount", event.Amount,
	)
}

func levelHandler(action action) *slog.HandlerOptions {
	var level slog.Level

	switch action {
	case Create:
		level = LevelCreate
	case Increase:
		level = LevelIncrease
	case Decrease:
		level = LevelDecrease
	case Transfer:
		level = LevelTransfer
	case receive:
		level = levelReceive
	}

	return &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLable, exists := levelNames[level]
				if !exists {
					levelLable = level.String()
				}

				a.Value = slog.StringValue(levelLable)
			}

			return a
		},
	}
}

const (
	LevelCreate   = slog.Level(-1)
	LevelIncrease = slog.Level(1)
	LevelDecrease = slog.Level(2)
	LevelTransfer = slog.Level(3)
	levelReceive  = slog.Level(5)
)

var levelNames = map[slog.Leveler]string{
	LevelCreate:   "CREATE",
	LevelIncrease: "INCREASE",
	LevelDecrease: "DECREASE",
	LevelTransfer: "SEND",
	levelReceive:  "RECEIVE",
}
