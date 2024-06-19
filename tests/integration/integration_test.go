package integration_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pradiptarana/book-online-store/app"
	"github.com/stretchr/testify/assert"
)

func runTestServer() *httptest.Server {
	return httptest.NewServer(app.SetupServer())
}

func Test_registration_endpoint(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()

	t.Run("it should return ok when registration", func(t *testing.T) {
		buf := new(bytes.Buffer)
		// create teardown

		io := bytes.NewBuffer([]byte(`{"username": "testing@gmail.com", "password": "12345"}`))
		resp, _ := http.Post(fmt.Sprintf("%s/api/v1/signup", ts.URL), "application/json", io)

		_, _ = buf.ReadFrom(resp.Body)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "{\n    \"message\": \"sign up success\"\n}", buf.String())
	})
}
