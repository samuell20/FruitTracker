package main

import(
    "log"

    "github.com/samuell20/FruitTracker/cmd/api/bootstrap"
)



func main(){
    if error := bootstrap.Run(); error != nil {
        log.Fatal(error)
    }
}
