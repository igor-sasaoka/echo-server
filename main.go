package main

import (
    "bufio"
    "log"
    "net"
)

func echo (c net.Conn) {
    defer c.Close()

    for {
        r := bufio.NewReader(c)
        s, err := r.ReadString('\n')
        if err != nil {
            log.Fatalln("Unable to read data")
        }

        log.Printf("%d bytes read: %s\n", len(s), s)

        log.Printf("Writing data\n")

        w := bufio.NewWriter(c)
        if _, err := w.WriteString(s + "\n> "); err != nil {
            log.Fatalln("Unable to write data")
        }
        w.Flush()
    }
}

func main() {
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalln("Unable to bind to port :8080")
    }

    log.Println("Listening to port :8080")

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatalln("Connection refused")
        } else {
            log.Println("Connection accepted")
        }

        go echo(c)
    }
}
