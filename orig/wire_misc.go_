package webdriver

import (
)

// // POST /session/:sessionid/timeouts
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts
// //
// // Configure the amount of time that a particular type of operation can execute for before they are aborted and a |Timeout| error is returned to the client.
// //
// //     type - {string} The type of operation to set the timeout for. Valid values are: "script" for script timeouts, "implicit" for modifying the implicit wait timeout and "page load" for setting a page load timeout.
// //     ms - {number} The amount of time, in milliseconds, that time-limited commands are permitted to run.
// func (s *Wire) Timeouts(type_value string, ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts",
//                               &Params{"type": type_value, "ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST /session/:sessionid/timeouts/async_script
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts/async_script
// //
// // Set the amount of time, in milliseconds, that asynchronous scripts executed by /session/:sessionId/execute_async are permitted to run before they are aborted and a |Timeout| error is returned to the client.
// //
// //     ms - {number} The amount of time, in milliseconds, that time-limited commands are permitted to run.
// func (s *Wire) TimeoutsAsyncScript(ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts/async_script", &Params{"ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST /session/:sessionid/timeouts/implicit_wait
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/timeouts/implicit_wait
// //
// // Set the amount of time the driver should wait when searching for elements. When searching for a single element, the driver should poll the page until an element is found or the timeout expires, whichever occurs first. When searching for multiple elements, the driver should poll the page until at least one element is found or the timeout expires, at which point it should return an empty list.
// //
// // If this command is never sent, the driver should default to an implicit wait of 0ms.
// //
// //     ms - {number} The amount of time to wait, in milliseconds. This value has a lower bound of 0.
// func (s *Wire) TimeoutsImplicitWait(ms float64) (wireResponse *WireResponse, err error) {

//   if req, err := s.PostRequest("/session/:sessionid/timeouts/implicit_wait", &Params{"ms": ms}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }

// // POST  /session/:sessionId/execute
// //
// // https://code.google.com/p/selenium/wiki/JsonWireProtocol#/session/:sessionId/execute
// //
// // Inject a snippet of JavaScript into the page for execution in the context of the currently selected
// // frame. The executed script is assumed to be synchronous and the result of evaluating the script is
// // returned to the client.
// //
// // The script argument defines the script to execute in the form of a function body. The value returned
// // by that function will be returned to the client. The function will be invoked with the provided args
// // array and the values may be accessed via the arguments object in the order specified.
// //
// // Arguments may be any JSON-primitive, array, or JSON object. JSON objects that define a WebElement
// // reference will be converted to the corresponding DOM element. Likewise, any WebElements in the
// // script result will be returned to the client as WebElement JSON objects.
// //
// //      JSON Parameters:
// //        script - {string} The script to execute.
// //        args - {Array.<*>} The script arguments.
// //
// //      Returns:
// //        {*} The script result.
// func (s *Wire) Execute(script string, args string) (wireResponse *WireResponse, err error) {

//   var req *http.Request
//   if req, err = s.PostRequest("/session/:sessionid/execute", &Params{"script": script, "args": args}); err == nil {

//     wireResponse, err = s.Do(req)

//   }

//   return wireResponse, err
// }











