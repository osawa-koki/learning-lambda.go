package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("===== ===== Test starting ===== =====")
	code := m.Run()
	fmt.Println("===== ===== Test ended ===== =====")
	os.Exit(code)
}

func TestHandler(t *testing.T) {
	t.Run("Main handler is working.", func(t *testing.T) {
		res, err := handler(events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/hello",
		})
		if err != nil {
			t.Errorf("Error is not nil: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Status code is not 200: %d", res.StatusCode)
		}
	})
}
