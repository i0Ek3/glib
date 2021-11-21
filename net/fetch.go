package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "io/ioutil"
    "net/http"
    "strings"
    "time"

    //"github.com/i0Ek3/noerr"
)

func main() {
    start := time.Now()
    ch := make(chan string) // no buffered
    defa := "links.conf"

    for _, url := range os.Args[1:] {
        reading(url, defa, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }

    elapsed := time.Since(start).Seconds()
    fmt.Println("%.2fs elapsed\n", elapsed)
}

func reading(url, defa string, ch chan string) {
    if url == "" {
        f, err := os.Open(defa)
        defer f.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "failed: %v\n", err)
        }

        r := bufio.NewReader(f)
        buf := make([]byte, 1024)
        for {
            n, err := r.Read(buf)
            if err != nil && err != io.EOF {
                panic(err)
            }
            if n == 0 {
                break
            }
            // FIXME
            for i := 0; i < n; i++ {
                go fetch(string(buf[i]), ch)
            }
        }
    }
    go fetch(url, ch)

}

func fetch(url string, ch chan<- string) {
    start := time.Now()

    prefix := []string{"http://", "https://"}
    if !strings.HasPrefix(url, prefix[0]) && !strings.HasPrefix(url, prefix[1]) {
        url = prefix[1] + url
    }

    resp, err := http.Get(url)
    //noerr.Xerr(err)
    if err != nil {
        //fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        ch <- fmt.Sprint(err)
        //os.Exit(1)
        return
    }

    //b, err := ioutil.ReadAll(resp.Body)
    //b, err := io.Copy(os.Stdout, resp.Body)
    b, err := io.Copy(ioutil.Discard, resp.Body)
    defer resp.Body.Close()
    if resp.Status != "200" {
        fmt.Println(resp.Status)
    }
    //noerr.Xerr(err)
    if err != nil {
        //fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        ch <- fmt.Sprint(err)
        //os.Exit(1)
        return
    }

    elapsed := time.Since(start).Seconds()

    //fmt.Printf("%s", b)
    ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, b, url)
}
