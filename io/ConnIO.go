package io

import (
    "strings"
    "net"
    "fmt"
)

func ConnRead (conn net.Conn, del string) (string, error) {
    mssgChar := make ([ ]byte, 1)
    var mssg string
    for {
        _, err := conn.Read (mssgChar)
        if err != nil {
            return mssg, err
        }
        mssg += string (mssgChar)
        if strings.Index (mssg, del) >= 0 {
            mssg = strings.Replace (mssg, del, "", -1)
            var err error
            return mssg, err
        }
    }
}

func ConnWrite (conn net.Conn, mssg string) {
    fmt.Fprintf (conn, mssg)
}
