[![Build Status](https://circleci.com/gh/jduckett/webdriver.png?circle-token=:circle-token)](https://circleci.com/gh/jduckett/webdriver)

# Webdriver Golang Style
A pure Golang webdriver library that supports the Json Wire Protocol.

# Firefox
One feature of this library is the ability to use an existing firefox profile simply by providing a directory that points to it.  When the Firefox client starts up, it will attempt to examine the specified profile
directory and only make the changes it needs.  The idea being that you may want or need to perform some types
of tests over a period of time.  This should leave things like cookies, cache, etc. in tact allowing you to test
a profile over time.  At least, I think so :)

