package goym

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/oklookat/goym/schema"
)

var (
	_albumIds = [4]schema.ID{
		"3370827",
		"14143149",
		"4979501",
		"389132",
	}
	_artistIds = [4]schema.ID{
		"419326",
		"1813",
		"205640",
		"1053",
	}
	_trackIds = [4]schema.ID{
		"27694817",
		"27694818",
		"27694819",
		"27694820",
	}

	_cachedClient *Client
)

// Получить клиент для запросов к API.
func getClient(t *testing.T) *Client {
	if _cachedClient != nil {
		return _cachedClient
	}

	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cl, err := New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	//cl.Http.SetLogger(loggerDefault{})
	cl.Http.SetRateLimit(1, time.Duration(1)*time.Second)

	_cachedClient = cl
	return cl
}

// type loggerDefault struct {
// }

// func (l loggerDefault) Debugf(msg string, args ...any) {
// 	log.Printf(msg, args...)
// }

// func (l loggerDefault) Err(msg string, err error) {
// 	if err == nil {
// 		log.Printf("%s", msg)
// 		return
// 	}
// 	log.Printf("%s. Err: %s", msg, err.Error())
// }
