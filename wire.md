# JsonWireProtocol Commands

I decided to break up the code into more managable chunks.  Below is a checklist.  A hyphen means
it has been completed.

wire.go
- /status
- /session
- /sessions
- /session/:sessionId
- /session/:sessionId/back
/session/:sessionId/cookie
/session/:sessionId/cookie/:name
- /session/:sessionId/forward
/session/:sessionId/keys
/session/:sessionId/location
- /session/:sessionId/refresh
- /session/:sessionId/source
- /session/:sessionId/title
- /session/:sessionId/url

wire_action.go
/session/:sessionId/alert_text
/session/:sessionId/accept_alert
/session/:sessionId/dismiss_alert
/session/:sessionId/moveto
/session/:sessionId/click
/session/:sessionId/buttondown
/session/:sessionId/buttonup
/session/:sessionId/doubleclick

wire_misc.go
/session/:sessionId/execute
/session/:sessionId/execute_async
/session/:sessionId/orientation
/session/:sessionId/screenshot
/session/:sessionId/timeouts
/session/:sessionId/timeouts/async_script
/session/:sessionId/timeouts/implicit_wait

wire_ime.go
/session/:sessionId/ime/available_engines
/session/:sessionId/ime/active_engine
/session/:sessionId/ime/activated
/session/:sessionId/ime/deactivate
/session/:sessionId/ime/activate

wire_window.go
- /session/:sessionId/window_handle
- /session/:sessionId/window_handles
/session/:sessionId/frame
/session/:sessionId/frame/parent
- /session/:sessionId/window
/session/:sessionId/window/:windowHandle/size
/session/:sessionId/window/:windowHandle/position
/session/:sessionId/window/:windowHandle/maximize

wire_element.go
- /session/:sessionId/element
- /session/:sessionId/elements
- /session/:sessionId/element/active
/session/:sessionId/element/:id     (json wire doc says command reserved for future use.  not implementing)
- /session/:sessionId/element/:id/element
- /session/:sessionId/element/:id/elements
- /session/:sessionId/element/:id/click
- /session/:sessionId/element/:id/submit
- /session/:sessionId/element/:id/text
- /session/:sessionId/element/:id/value
- /session/:sessionId/element/:id/name
- /session/:sessionId/element/:id/clear
- /session/:sessionId/element/:id/selected
- /session/:sessionId/element/:id/enabled
- /session/:sessionId/element/:id/attribute/:name
/session/:sessionId/element/:id/equals/:other
- /session/:sessionId/element/:id/displayed
- /session/:sessionId/element/:id/location
- /session/:sessionId/element/:id/location_in_view
- /session/:sessionId/element/:id/size
- /session/:sessionId/element/:id/css/:propertyName


wire_storage.go
/session/:sessionId/application_cache/status
/session/:sessionId/local_storage
/session/:sessionId/local_storage/key/:key
/session/:sessionId/local_storage/size
/session/:sessionId/log
/session/:sessionId/log/types
/session/:sessionId/session_storage
/session/:sessionId/session_storage/key/:key
/session/:sessionId/session_storage/size

wire_touch.go
/session/:sessionId/touch/click
/session/:sessionId/touch/down
/session/:sessionId/touch/up
/session/:sessionId/touch/move
/session/:sessionId/touch/scroll
/session/:sessionId/touch/scroll
/session/:sessionId/touch/doubleclick
/session/:sessionId/touch/longclick
/session/:sessionId/touch/flick
/session/:sessionId/touch/flick





