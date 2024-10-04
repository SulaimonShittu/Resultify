package log

import (
	"Resultify/registry"
	"bytes"
	"fmt"
	sllog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	sllog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	sllog.SetFlags(0)
	sllog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log message. Service responded with %v - %v", res.StatusCode, res.Status)
	}
	return len(data), nil
}
