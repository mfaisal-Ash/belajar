package belajar

// import "time"

// User represents the user structure
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
}

// Credential represents the response for credentials
type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

// ResponseDataUser represents the response data for users
type ResponseDataUser struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty"`
}

// Response represents a general response
type Response struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
}

// ResponseEncode represents the response after encoding
type ResponseEncode struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

// QRCodeResponse represents the response after scanning the QR code to get the OTP
type QRCodeResponse struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	OTP     string `json:"otp,omitempty" bson:"otp,omitempty"`
}
