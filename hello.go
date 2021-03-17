package main

import (
    "fmt"

    "sesquipedalian-dev.github.io/go_vanity_2"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}