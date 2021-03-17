package main

import (
    "fmt"

    greetings "sesquipedalian-dev.github.io/go_vanity_1"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}