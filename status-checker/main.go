package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"

)

func Handler() {
	fmt.Println("hello test")

}

func main() {
	lambda.Start(Handler)
}
