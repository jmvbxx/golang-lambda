package main

import (
    "context"
    "net/http"
    "log"
    // "fmt"
    // "io/ioutil"

    "github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) *http.Response {        
    response, err := http.Get("http://httpbin.org/status/418")
    if err != nil {
        log.Fatal(err)
    } 
    return response
}

func main() {
    lambda.Start(HandleRequest)
}
