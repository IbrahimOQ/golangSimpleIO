# golangSimpleIO - by Ibrahim Oladipupo Qamardeen (kensharval@gmail.com)

DEVELOPER INFO
# I'm Ibrahim, a server-side systems engineer. This will be my first open-source project. I work with Golang and Java.

DESCRITPION
# This Go package provides a straightforward way to read messages from socket connections and command line interface.
# It also provides functions for writing to consoles and socket connections.

---------------------------------------------------------------------------------------------------------------------------------------
USAGE: READING FROM CONNECTION
---------------------------------------------------------------------------------------------------------------------------------------

# input, response := io.ConnRead (connRes, "\n")

# Reading from a network connection requires a connection resource.
# It can be provided by some procedures like net.Dial (...).

# "ConnRead" also optionally requires a delimiter, as a string. If no delimiter is provided, "\n" will be used.

# The procedure returns two data. The first one is the input taken.
# The second is an error concerning the input process.
# If there is no error, nothing will be provided.

# Providing a string that can not be found in a message received, will help return the whole message, once the connection is closed by the other-party.

# input, response := io.ConnRead (connRes, "##end")

# Assuming a connection is to an HTTP server, the example above will return the full HTTP response, once the connection is closed by the server.

---------------------------------------------------------------------------------------------------------------------------------------
WRITING TO CONNECTION
---------------------------------------------------------------------------------------------------------------------------------------

# io.ConnWrite (connRes, "This is some message")

# Writing to a network connection requires a connection resource and a message to be sent.
# Connection resource can be provided by some procedures like net.Dial (...).

# This function is based on fmt.FPrintf (), and was provided for programmers who desire all-in-one.

---------------------------------------------------------------------------------------------------------------------------------------
READING FROM CONSOLE
---------------------------------------------------------------------------------------------------------------------------------------

# input, response := io.ConsoleRead ("\n", "Some input: ")

# "ConsoleRead" requires at least one parameter. The first is a single-character string that will mark the end of an input process.
# The second parameter is what should be displayed as label when trying to take input from the console.
# If a label is not needed, this parameter can be omitted

# The procedure returns two data. The first one is the input taken. The second is an error concerning the input process.
# If there is no error, nothing will be provided.

---------------------------------------------------------------------------------------------------------------------------------------
WRITING TO CONSOLE
---------------------------------------------------------------------------------------------------------------------------------------

# io.ConsoleWrite ("Hello world!", "\n", "This is Golang")

# "ConsoleWrite" allows writing of multiple strings to the console, in one call.
# It is based on fmt.Println ().
