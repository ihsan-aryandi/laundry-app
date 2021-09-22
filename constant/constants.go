package constant

import (
	"time"
)

const DefaultTimeFormat = time.RFC3339

/*
	User Profile
*/
const DefaultTotalReputation = 1
const DefaultUserPhoto = "default.jpg"
const DefaultUserRole = "user"

/*
	Authentication
*/
const AuthorizationCookieName = "Authorization"
const UserPayloadContextName = "user_payload"

/*
	Registration Link
*/
const RegistrationLinkLaundryType = "laundry"
const RegistrationLinkCustomerType = "customer"