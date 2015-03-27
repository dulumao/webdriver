# Webdriver Golang style

# Running the examples

The examples subdirectory contains many examples.  All of them are part of package main and all of a main function, therefore, using the "go build" will fail.  The intent is to use "go build <file_name>" or "go run <file_name>" to run an example.

The examples were written simultaneously with the code and tests.  Nothing special, just proof.

# Firefox
One feature of this library is the ability to point at an existing firefox profile simply by providing a directory point to it.  When the Firefox client starts up, it will attempt to examine the specified profile
directory and only make the changes it needs.  The idea being that you may want or need to perform some types
of tests over a period of time.  This should leave things like cookies, cache, etc. in tact allowing you to test
a profile over time.  At least, I think so :)

