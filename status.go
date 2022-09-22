package sock

// 7.4.  Status Codes

// When closing an established connection (e.g., when sending a Close
// frame, after the opening handshake has completed), an endpoint MAY
// indicate a reason for closure.  The interpretation of this reason by
// an endpoint, and the action an endpoint should take given this
// reason, are left undefined by this specification.  This specification
// defines a set of pre-defined status codes and specifies which ranges
// may be used by extensions, frameworks, and end applications.  The
// status code and any associated textual message are optional
// components of a Close frame.

// Endpoints MAY use the following pre-defined status codes when sending
// a Close frame.
const (
	// 1000 indicates a normal closure, meaning that the purpose for
	// which the connection was established has been fulfilled.
	StatusNormal int = 1000

	// 1001 indicates that an endpoint is "going away", such as a server
	// going down or a browser having navigated away from a page.
	StatusGoingAway int = 1001

	// 1002 indicates that an endpoint is terminating the connection due
	// to a protocol error.
	StatusProtocolError int = 1002

	// 1003 indicates that an endpoint is terminating the connection
	// because it has received a type of data it cannot accept (e.g., an
	// endpoint that understands only text data MAY send this if it
	// receives a binary message).
	StatusCannotAccept int = 1003

	// 1005 is a reserved value and MUST NOT be set as a status code in a
	// Close control frame by an endpoint.  It is designated for use in
	// applications expecting a status code to indicate that no status
	// code was actually present.
	StatusNoStatus int = 1005

	// 1006 is a reserved value and MUST NOT be set as a status code in a
	// Close control frame by an endpoint.  It is designated for use in
	// applications expecting a status code to indicate that the
	// connection was closed abnormally, e.g., without sending or
	// receiving a Close control frame.
	StatusClosedAbnormally int = 1006

	// 1007 indicates that an endpoint is terminating the connection
	// because it has received data within a message that was not
	// consistent with the type of the message (e.g., non-UTF-8 [RFC3629]
	// data within a text message).
	StatusNotConsistent int = 1007

	// 1008 indicates that an endpoint is terminating the connection
	// because it has received a message that violates its policy.  This
	// is a generic status code that can be returned when there is no
	// other more suitable status code (e.g., 1003 or 1009) or if there
	// is a need to hide specific details about the policy.
	StatusPolicyViolation int = 1008

	// 1009 indicates that an endpoint is terminating the connection
	// because it has received a message that is too big for it to
	// process.
	StatusTooBig int = 1009

	// 1010 indicates that an endpoint (client) is terminating the
	// connection because it has expected the server to negotiate one or
	// more extension, but the server didn't return them in the response
	// message of the WebSocket handshake.  The list of extensions that
	// are needed SHOULD appear in the /reason/ part of the Close frame.
	// Note that this status code is not used by the server, because it
	// can fail the WebSocket handshake instead.
	StatusExtensionMissing int = 1010

	// 1011 indicates that a server is terminating the connection because
	// it encountered an unexpected condition that prevented it from
	// fulfilling the request.
	StatusUnexpectedCondition int = 1011

	// 1015 is a reserved value and MUST NOT be set as a status code in a
	// Close control frame by an endpoint.  It is designated for use in
	// applications expecting a status code to indicate that the
	// connection was closed due to a failure to perform a TLS handshake
	// (e.g., the server certificate can't be verified).
	StatusTlsFailure int = 1015
)
