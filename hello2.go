package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

// type MyEvent struct {
//         Name string `json:"name"`
// }

// func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
func HandleRequest(ctx context.Context) (string, error) {        
        return fmt.Sprintf("Hello!"), nil
}

func main() {
        lambda.Start(HandleRequest)
}