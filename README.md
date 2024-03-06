# Outis
The outis library helps you create and manage routines, schedule execution time, and control competition between other processes.

## Example of how to use.
  
```go
package main

import ...

func main() {
  // Initialize Outis to be able to add routines
  watch := outis.Watcher(utils.NewOutisProduction())
  
  go watch.Go(
    // Routine identifier to perform concurrency control 
    outis.WithID("422138b3-c721-4021-97ab-8cf7e174fb4f"),

    outis.WithName("Here is the name of my routine"),
    outis.WithDesc("Here is the description of my routine"),
    
    // It will run every 10 second
    outis.WithInterval(10),
  
    // It will run from 12pm to 4pm.
    // by default, there are no time restrictions.
    outis.WithHours(12, 16),
    
    // Here the script function that will be executed will be passed
    outis.WithRoutine(func(c outis.Channel) {
      log.Println("Hi, I'm a script!")

      ...
      
      if err != nil {
        c <- errors.New("This is an error message")
        return
      }
    }),
  )

  // Method that maintains routine in the process
  watch.Wait()
}

```

## Implements the IOutis interface
It is necessary to implement the IOutis interface to initialize Watcher.

```go
package utils

import ...

type repo struct{}

// NewOutisProduction initializes the implementation of the Outis interface.
// 
// It is recommended to create a role for each environment, 
// one for development and one for production.
// 
// For development you can implement the Lock and Unlock method 
// to use in-memory control and production can be done using 
// a redis database or a central server.
func NewOutisProduction() outis.IOutis {
  return &repo{}
}

// Lock defines the method by which concurrency 
// blocking will be implemented
func (*repo) Lock(id string) (string, error) {}

// Unlock defines the method by which concurrency
// unblocking will be implemented
func (*repo) Unlock(id string) error {}

// Store defines the method for saving 
// the routine's initial information
func (*repo) Store(*outis.Outis) error {}

// Load defines the method to fetch updated 
// information from the routine
func (*repo) Load(*outis.Outis) error {}

// Handling defines the method for handling errors
func (*repo) Handling(error) {}
```
