package main

import (
    "log"
    "os/exec"
)

func main() {

    cmd := exec.Command("mkdir", "/home/kali/Desktop/prova2")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}
