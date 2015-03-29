package webdriver

import (
  "encoding/json"
)

type (

  Sessions struct {

    List []*Session

    *Wire

  }

  Session struct {

    ActualCapabilities *ActualCapabilities

    *Wire

  }

)

func (s *Sessions) SetDefaults() (err error) {

  s.Wire = &Wire{}
  s.Wire.SetDefaults()

  return err
}

func (s *Session) SetDefaults() (err error) {

  s.Wire = &Wire{}
  s.Wire.SetDefaults()

  return err
}

// Creates a new session for a Client.  An JsonWireProtocol call is made
// to establish a session with a server.  The new session is added to the
// list of active sessions and returned to the caller.
//
// Capabilities are optional, however, it you define them, then, you must
// pass them in a specific order to this method.  Desired first, then, Required.
// Capabilities are currently implemented as a simple map and quite frankly
// I high doubt there will be much need to even support passing capabilities.
// However, it is in the spec, so, there is minimal support for it.
//
//      session, err := client.NewSession(
//               &webdriver.Capabilities{"Platform": "Linux"}, // desired
//               &webdriver.Capabilities{})                    // required
//
// When a new session is created, the server will return the actual capabilities
// currently supported.  An ActualCapabilities struct is created and attached
// to the returned session.
func (s *Sessions) NewSession(capabilities ...*Capabilities) (session *Session, err error) {

  // make the API call to establish a new session
  // if there is an error it is returned
  var wireResponse *WireResponse
  if wireResponse, err = s.Session(capabilities...); err == nil {

    // seems like everything went as planned.
    // create a new session and initialize it.
    session = &Session{}
    session.SetDefaults()

    // setting the BaseUrl on the session is critical for http requests
    session.BaseUrl = s.BaseUrl

    // the Session ID returned during the API call.
    session.SessionID = wireResponse.SessionID

    // add the newly created session to the list of sessions
    s.List = append(s.List, session)

    // extract the actual capabilities from the response and attach
    // them to the session
    capabilities := &ActualCapabilities{}
    if err = json.Unmarshal(wireResponse.Value, capabilities); err == nil {

      session.ActualCapabilities = capabilities

    }


  }

  return session, err
}


