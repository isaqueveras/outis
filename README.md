# Outis
The outis library helps you create and manage routines, schedule execution time, and control competition between other processes.

## Example of how to use.

```go
package main

import ...

func main() {
  // Initialize Outis to be able to add routines
  watch := outis.Watcher(utils.NewOutis())

  go watch.Go(
    // Routine identifier to perform concurrency control
    outis.WithID("422138b3-c721-4021-97ab-8cf7e174fb4f"),

    outis.WithName("Here is the name of my routine"),
    outis.WithDesc("Here is the description of my routine"),

    // It will run every 10 second
    outis.WithInterval(time.Second * 10),

    // It will run from 12pm to 4pm.
    // by default, there are no time restrictions.
    outis.WithHours(12, 16),

    // Time when routine information will be updated
    outis.WithLoadInterval(time.Second * 30),

    // Here the script function that will be executed will be passed
    outis.WithRoutine(func(ctx *outis.Context) {
      ctx.Info("this is an information message")
      ctx.Error(errors.New("this is an error message"))

      ctx.Metric("client_ids", []int64{234234})
      ctx.Metric("notification", outis.Metric{
        "client_id": 234234,
        "message":   "Hi, we are notifying you.",
        "fcm":       "3p2okrmionfiun2uni3nfin2i3f",
      })
    }),
  )

  // Method that maintains routine in the process
  watch.Wait()
}

```

## Implements the Interface
It is necessary to implement the Interface interface to initialize Watcher.

```go
package utils

import ...

type repo struct{}

// NewOutis initializes the implementation of the Interface interface.
//
// It is recommended to create a role for each environment,
// one for development and one for production.
//
// For development you can implement the Lock and Unlock method
// to use in-memory control and production can be done using
// a redis database or a central server.
func NewOutis() outis.Interface {
  return &repo{}
}

// Lock defines the method by which concurrency
// blocking will be implemented
func (*repo) Lock(outis.ID) (outis.ID, error) {}

// Unlock defines the method by which concurrency
// unblocking will be implemented
func (*repo) Unlock(outis.ID) error {}

// Store defines the method for saving
// the routine's initial information
func (*repo) Store(*outis.Context) error {}

// Load defines the method to fetch updated
// information from the routine
func (*repo) Load(*outis.Context) error {}

// Event defines the method for handling events
func (*repo) Event(outis.Event) {}
```
