package main

import (
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
	ctx.Metric("uuid", outis.Metric{
		"success": []uint{23423, 1423, 4322},
		"failure": []uint{23423, 4546, 3423},
	})

	ctx.Debug("Hello 02")

	ctx.Metric("client", outis.Metric{
		"id":    2134234,
		"name":  "Antonio Francisco da Silva",
		"email": "antonio.francisco.silva@email.com",
		"address": outis.Metric{
			"id":     2345564,
			"street": "Av. 01",
			"number": "S/N",
		},
	})
}
