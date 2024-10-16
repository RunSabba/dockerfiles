package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

// loggingMiddleware wraps an http.HandlerFunc and logs the request details
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Create a custom ResponseWriter to capture the status code
        lrw := newLoggingResponseWriter(w)

        // Call the next handler
        next.ServeHTTP(lrw, r)

        // Log the request details
        log.Printf(
            "%s %s %s %d %s %s",
            r.RemoteAddr,
            r.Method,
            r.URL.Path,
            lrw.statusCode,
            r.UserAgent(),
            time.Since(start),
        )
    }
}

type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Docker!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Wrap the hello handler with the logging middleware
    http.HandleFunc("/", loggingMiddleware(hello))

    fmt.Printf("Server listening on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
        os.Exit(1)
    }
}