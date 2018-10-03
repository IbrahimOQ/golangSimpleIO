package io

import (
    "bufio"
    "fmt"
    "os"
)

func ConsoleRead (args ... string) (string, error) {
    delimiter := byte ('\n')
    if len (args) > 0 && args [0] != "" {
        delimiter = byte (args [0][0])
    }
    if len (args) > 1 && args [1] != "" {
        fmt.Print (args [1])
    }
    input, status := bufio.NewReader (os.Stdin).ReadString (delimiter)
    return input, status
}

func ConsoleWrite (args ... string) {
    for count := 1; count <= len (args); count ++ {
        fmt.Print (args [count - 1])
    }
}