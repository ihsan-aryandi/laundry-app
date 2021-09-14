package constant

/*
	Status Code : 400
*/
const ErrInvalidFieldsCode = "INVALID_FIELDS"
const ErrInvalidFieldsMessage = "Invalid fields"

const ErrInvalidRequestBodyCode = "INVALID_REQUESTED_BODY"
const ErrInvalidRequestBodyMessage = "Your requested body is invalid"

const ErrDataNotFoundCode = "DATA_NOT_FOUND"
const ErrDataNotFoundMessage = "%s not found"

const ErrDataIsExistCode = "DATA_ALREADY_EXISTS"
const ErrDataIsExistMessage = "%s already exists"

const ErrUsernameIsAlreadyRegisteredCode = "USERNAME_ALREADY_REGISTERED"
const ErrUsernameIsAlreadyRegisteredMessage = "This username is already registered"

const ErrLoginCode = "LOGIN_ERROR"
const ErrLoginMessage = "Username or password is incorrect"

const ErrForbiddenCode = "FORBIDDEN"
const ErrForbiddenMessage = "You are not allowed to access this route"

const ErrUnauthorizedCode = "UNAUTHORIZED"
const ErrUnauthorizedMessage = "You are unauthorized"

const ErrInvalidPaginationCode = "INVALID_PAGINATION"
const ErrInvalidPaginationMessage = "Invalid pagination format"

/*
	Status Code : 500
*/
const ErrInternalServerCode = "INTERNAL_SERVER_ERROR"
const ErrInternalServerMessage = "Something is wrong with server"