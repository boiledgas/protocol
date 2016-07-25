package telematics

// Response flags
const (
	RESPONSE_FLAGS_OK            byte = 0x00 // no problem
	RESPONSE_FLAGS_AUTHORIZATION byte = 0x01 // request for identification
	RESPONSE_FLAGS_DESCRIPTION   byte = 0x02 // response for configuration
	RESPONSE_FLAGS_ERROR         byte = 0x80 // error
)

// Packet types
const (
	// Default response
	RESPONSE_OK byte = 0x00
	// Need authorization
	RESPONSE_AUTHORIZATION = 0x01
	// Need description
	RESPONSE_DESCRIPTION = 0x02
	// Unhandled error
	RESPONSE_ERROR = 0x80
)
