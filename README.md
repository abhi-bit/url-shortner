# url-shortner

URL shortening as a service

## Installation 

    `go get github.com/abhi-bit/url-shortner`

Also make sure Couchbase Server is running on your box. In this case I used local couchbase instance `http://localhost:8091` and bucket `default`

## Initialisation

Run the `client.go` inside `client` dir. It will start web service on 8080

```
$ go run client.go 
Starting web service on :8080
```

## API

Just 2 API endpoints available currently.

### POST /encode

Pass the url to encode using POST and response will be encoded URL.

```
$ curl -d "http://google.com" http://localhost:8080/encode
http://localhost:8080/2bJ
```

### POST /decode

Pass the url to decode using POST and response is original URL.

```
$ curl -d "http://localhost:8080/2bL" http://localhost:8080/decode
http://google.com
```

## Blob Structure

Below is the blob structure inside Couchbase Server

```
{
    "ID": 10001,
    "origURL": "http://google.com/",
    "shortURL": "http://localhost:8080/2bJ"
}
```
