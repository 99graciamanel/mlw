package ddos

import (
    "fmt"
    "sync"
)

// Hello returns a greeting for the named person.
func Hello(wg *sync.WaitGroup, name string) {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    fmt.Println(message)
    wg.Done()
}