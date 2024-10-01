package log

import (
	"io/ioutil"
	sllog "log"
	"net/http"
	"os"
)

var x *sllog.Logger

type outputLog string

func (fx outputLog) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(string(fx), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(p)
}

func Run(destination string) {
	x = sllog.New(outputLog(destination), "", sllog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}

func write(msg string) {
	x.Printf("%s\n", msg)
}
