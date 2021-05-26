package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TestMiddleware struct {
}

func (l TestMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/api/v2/status") {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println("Request Received")
		var buf bytes.Buffer
		next.ServeHTTP(&responseWriterTest{
			w:      w,
			buffer: &buf,
		}, r)

		fmt.Printf("Request Received -> Response: %d %s\n", buf.Len(), buf.String())
		buf.Write([]byte("ROY is a nice guy"))
		w.Write(buf.Bytes())
	})
}

type responseWriterTest struct {
	w      http.ResponseWriter
	buffer io.Writer
}

func (r responseWriterTest) Header() http.Header {
	return r.w.Header()
}

func (r responseWriterTest) Write(bytes []byte) (int, error) {
	return r.buffer.Write(bytes)
}

func (r responseWriterTest) WriteHeader(statusCode int) {
	r.w.WriteHeader(statusCode)
}
