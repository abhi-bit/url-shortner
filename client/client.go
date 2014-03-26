package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/abhi-bit/url-shortner"
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/go-couchbase"
)

type KeyStruct struct {
	originalURL string
	shortURL    string
	ID          int64
}

const (
	bname = "default"
	url   = "http://localhost:8091"
)

func mf(err error) {
	if err != nil {
		clog.Error(err)
	}
}

func main() {

	generator := urlshortner.NewGenerator()

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		Handler(w, r, generator)
	})
	http.ListenAndServe(":8080", nil)

}

func Handler(w http.ResponseWriter, r *http.Request, generator *urlshortner.Generator) {

	generator.Start()
	cb, err := couchbase.Connect(url)
	mf(err)

	pool, err := cb.GetPool("default")
	mf(err)

	bucket, err := pool.GetBucket(bname)
	mf(err)

	var id int64
	var hash, origURL, short string

	id = generator.GetID()

	hash = urlshortner.Dehydrate(id)
	short = "http://localhost:8080/" + hash

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		mf(err)
		origURL = strings.Split(string(body), "=")[1]
	}

	key := &KeyStruct{originalURL: origURL, shortURL: short, ID: id}
	fmt.Println("orig:", origURL, "short:", short, "ID:", id, "hash:", hash)
	fmt.Printf("%#v\n", key)
	err = bucket.Set(strconv.Itoa(int(id)), 0, key)

}
