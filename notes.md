# Testing

// Design decisions
//
// The webdriver package contains data and methods to communicate with a
// webdriver http server as defined by the JsonWireProtocol spec.
//
// Subpackages firefox, chrome, etc. contain data and methods to physically
// fire up a browser and establish a running JsonWireProtocol server via a
// plugin or extension.  The implementation is different across browsers hence
// the decicion to create subpackages.  Each subpackage utilizes the code in the
// main webdriver package as sort of an API to the JsonWireProtocol API.
// Each subpackage utilizes specifically, but, not limited to the Client
// struct and methods.
//
// Have a look at the documentation for each individual struct throughout the
// code to gain insight as to the reasons why something exists where it does.
// and the intent behind it.
//
// All of the code needed to make http calls to a server is contained in
// wire_http.go.  The code was written prior to writing any of the actual
// code to make JsonWireProtocol API calls and is written in such a way as to
// provide as much flexibility as possible during the development and
// (more importantly) maintenance of the code base.  Instead of making a single
// or couple of "end all / be all" methods with a bunch of conditional logic,
// the approach is a bit more fragmented.  There are a few core functions for
// constructing and executing http requests and several wrapper functions
// to provide convenience for the API while still allowing the ability to break
// an API call down into smaller pieces using the same core functions.
//
// The reasoning behind this approach is that the JsonWireProtocol API has over
// 100 methods defined and I am developing this code base on-the-fly.  So, many
// things are unknown at this point and the idea is to provide as much flexibility
// upfront and avoid headaches later down the line...
