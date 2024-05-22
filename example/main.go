package main

import (
	"math/rand"
	"time"

	"github.com/isaqueveras/outis"
)

func main() {
	watch := outis.Watcher("8b1d6a18-5f3d-4482-a574-35d3965c8783", "v1/example")

	go watch.Go(
		outis.WithID(outis.ID("0b2d07ca-e3db-478a-9455-d5f476ac8d37")),
		outis.WithName("Example routine"),
		outis.WithInterval(time.Second),
		outis.WithScript(example),
	)

	watch.Wait()
}

func example(ctx *outis.Context) {
	ctx.Metadata("uuid", outis.Metadata{
		"success": []uint{23423, 1423, 4322},
		"failure": []uint{23423, 4546, 3423},
	})

	totalProblemsResolved := ctx.NewHistogram("problems_resolved")
	for i := 0; i < 5; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		n := rand.Intn(10 + time.Now().Second()*3)
		time.Sleep(time.Duration(n) * time.Millisecond)
		totalProblemsResolved.Add(float64(n))
	}

	ctx.Debug("Hello 02")

	customersNotified := ctx.NewIndicator("customers_notified")
	customersNotified.Add(13.21)

	totalCanceledProcesses := ctx.NewHistogram("canceled_processes")
	for i := 0; i < 3; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		n := rand.Intn(10 + time.Now().Second()*5)
		time.Sleep(time.Duration(n) * time.Millisecond)
		totalCanceledProcesses.Add(float64(n))
	}

	ctx.Metadata("client", outis.Metadata{
		"id":    2134234,
		"name":  "Antonio Francisco da Silva",
		"email": "antonio.francisco.silva@email.com",
		"address": outis.Metadata{
			"id":     2345564,
			"street": "Av. 01",
			"number": "S/N",
		},
	})
}
