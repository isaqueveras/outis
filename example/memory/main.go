package main

import (
	"errors"
	"time"

	"github.com/isaqueveras/outis"
	"github.com/isaqueveras/outis/memory"
)

func main() {
	watch := outis.Watcher("example/memory", memory.NewOutis())

	go watch.Go(
		outis.WithID("b7504beb-1132-4ced-8813-3525523cac1d"),
		outis.WithName("Routine 01"),
		outis.WithInterval(time.Microsecond),
		outis.WithHours(8, 20),
		outis.WithLoadInterval(time.Second*30),
		outis.WithRoutine(sendNotification),
	)

	watch.Wait()
}

func sendNotification(ctx *outis.Context) {
	ctx.Info("this is an information message")
	ctx.Error(errors.New("this is an error message"))

	ctx.Metric("client_ids", []int64{234234})
	ctx.Metric("notification", outis.Metric{
		"client_id": 234234,
		"message":   "Hi, we are notifying you.",
		"fcm":       "3p2okrmionfiun2uni3nfin2i3f",
	})
}
