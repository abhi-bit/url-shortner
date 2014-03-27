package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/abhi-bit/url-shortner"
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/go-couchbase"
)

const (
	bname          = "default"
	cbURL          = "http://localhost:8091"
	shortURLPrefix = "http://localhost:8080/"
)

func mf(err error) {
	if err != nil {
		clog.Error(err)
	}
}

func main() {

	generator := urlshortner.NewGenerator()

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		EncodeHandler(w, r, generator)
	})
	http.HandleFunc("/decode", DecodeHandlder)
	fmt.Println("Starting web service on :8080")
	http.ListenAndServe(":8080", nil)

}

func EncodeHandler(w http.ResponseWriter, r *http.Request, generator *urlshortner.Generator) {

	generator.Start()
	cb, err := couchbase.Connect(cbURL)
	mf(err)

	pool, err := cb.GetPool("default")
	mf(err)

	bucket, err := pool.GetBucket(bname)
	mf(err)

	var id int64
	var hash, origURL, short string

	id = generator.GetID()

	hash = urlshortner.Dehydrate(id)
	short = shortURLPrefix + hash

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		mf(err)
		origURL = string(body)
	}

	io.Copy(w, strings.NewReader(short))

	err = bucket.Set(strconv.Itoa(int(id)), 0, map[string]interface{}{"origURL": origURL, "shortURL": short, "ID": id})

}

func DecodeHandlder(w http.ResponseWriter, r *http.Request) {

	cb, err := couchbase.Connect(cbURL)
	mf(err)

	pool, err := cb.GetPool("default")
	mf(err)

	bucket, err := pool.GetBucket(bname)
	mf(err)

	var hash, shortURL string
	var id int64

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		mf(err)
		shortURL = string(body)
	}

	hash = strings.Split(shortURL, "/")[3]
	id = urlshortner.Saturate(hash)

	ob := map[string]interface{}{}
	mf(bucket.Get(strconv.Itoa(int(id)), &ob))

	io.Copy(w, strings.NewReader(ob["origURL"].(string)))
}
