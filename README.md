# Go OpenAI Wrapper

Endpoints supported:
 - Chat completion

## Installation

As a library

```shell
go get github.com/kfatehi/go-openai
```

## Example Usage

Here is an example you can use.

1. Create a directory for your project and `cd` into it

2. Setup your go.mod as main: `go mod init main`

3. Get this library: `go get github.com/kfatehi/go-openai`

4. Put the following code in `main.go` and set your API key:

```
package main

import (
	"log"

	openai "github.com/kfatehi/go-openai"
)

func main() {
	apiKey := "YOUR API KEY"

	req := openai.ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatMessage{
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}

	resp, err := openai.ChatCompletion(apiKey, req)
	if err != nil {
		panic(err)
	}
	log.Println(resp.Choices[0].Content)
}
```

5. Run it: `go run main.go`

## Testing

The test will make a real API call.

1. Create a .env file

2. Run the test

```
go test -modfile go_test.mod
```

## Generated by ChatGPT

this was created by taking the curl request and response from this page https://platform.openai.com/docs/api-reference/chat/create and using a prompt like this:

> write a Go wrapper for the following HTTP API. I will be providing the curl request and response:

followed by:

> it should take the api key and request as parameters and return a tuple of error and response as is the standard idiom in go

then a bit of manual cleanup. i also had it generate a test which i cleaned up.