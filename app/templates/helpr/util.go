package helpr

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"time"
)

// CustomLog global
var CustomLog *log.Logger

// SetupLogger setup logger
func SetupLogger(customLog *log.Logger) {
	CustomLog = customLog
}

// Debug for custom logging
func Debug(v interface{}) {
	if os.Getenv("STAGE") != "production" {
		color := "\n\033[1;36m------------ DEBUG ------------\033[0m\n\033[1;33m%+v\033[0m\n\033[1;36m-------------------------------\033[0m\n"
		k := reflect.TypeOf(v).Kind()
		t := fmt.Sprintf("%T", v)

		out := v
		if k == reflect.Ptr && t == "*errors.errorString" {
			color = "\n\033[1;36m------------ DEBUG ------------\033[0m\n\033[1;31m%+v\033[0m\n\033[1;36m-------------------------------\033[0m\n"
			out = v
		} else if k == reflect.Ptr {
			b, _ := json.MarshalIndent(v, "", "\t")
			out = string(b)
		}
		if flag.Lookup("test.v") == nil {
			CustomLog.Output(2, fmt.Sprintf(color, out))
		}
	}
}

// NetClient - custom http.Client with transport configs
func NetClient() *http.Client {
	netTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	netClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	defer netTransport.CloseIdleConnections()
	return netClient
}

// GenerateChallenge generator
func GenerateChallenge(l int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	id, err := randASCIIBytes(l)
	if err != nil {
		return "", fmt.Errorf("Error generating ID")
	}
	return string(id), nil
}

// randASCIIBytes - A helper function create and fill a slice of length n with characters from a-zA-Z0-9_-.
// https://github.com/kpbird/golang_random_string
func randASCIIBytes(n int) ([]byte, error) {
	letterBytes := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	output := make([]byte, n)
	randomness := make([]byte, n)
	_, err := rand.Read(randomness)
	if err != nil {
		return nil, err
	}

	l := len(letterBytes)
	for pos := range output {
		random := uint8(randomness[pos])
		randomPos := random % uint8(l)
		output[pos] = letterBytes[randomPos]
	}

	return output, nil
}
