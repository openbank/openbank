Annex
=====

APIError
--------

APIError defines the error model returned by an API in case something went wrong

| Name    | Type      | Description                                                                               |
|---------|-----------|-------------------------------------------------------------------------------------------|
| Code    | int32     | Code of the error returned, corresponding to gRPC or HTTP code, depending on your client. |
| Type    | ErrorType | Type is an enum to help handling errors programatically.                                  |
| Message | string    | Message is a human readable message providing more details about the error.               |
| DocURL  | string    | DocURL is a direct link to the specific error type documentation, when applicable.        |

ErrorType
---------

ErrorType defines the type of the error.

| Value                | Description                                                                           |
|----------------------|---------------------------------------------------------------------------------------|
| UnknownErrorType     |                                                                                       |
| AccountAlreadyExists | ErrorType_AccountAlreadyExists is returned when trying to create a duplicate account. |
