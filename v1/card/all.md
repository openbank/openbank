# Card API v1.0.0

Provides create and read operations on the card resource.

* Host ``

* Base Path ``

## Retrieve card information

Retrieves all data from a card, selected by the card_token you supplied.

```sh
curl -X GET \
	/v1/card/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/card/{CardToken}`

### Query Parameters

| Name      | Type   | Description                                              |
|-----------|--------|----------------------------------------------------------|
| CardToken | string | CardToken is the identifier to get the card information. |

### Responses

#### Response body

| Name          | Type             | Description                                                |
|---------------|------------------|------------------------------------------------------------|
| ID            | string           | ID is a unique identifier of a card.                       |
| Account       | string           | Account is the ID of the account associated with the card. |
| OwnerName     | string           | OwnerName is the name of the card owner.                   |
| ContactNumber | string           | ContactNumber is the contact number of the card owner.     |
| FirstName     | string           | FirstName is the first name of card owner.                 |
| LastName      | string           | LastName is the last name of the card owner.               |
| Expiry        | Timestamp        | Expiry is an expiry date of the card.                      |
| Status        | CardStatus       | Status is the status of the card.                          |
| AccessStatus  | CardAccessStatus | AccessStatus is the access status of the card.             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "id": "string",
  "account": "string",
  "owner_name": "string",
  "contact_number": "string",
  "first_name": "string",
  "last_name": "string",
  "expiry": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "CardStatus",
  "access_status": "CardAccessStatus"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update card access status

Update card access status.

```sh
curl -X POST \
	/v1/card/access_status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet updatecardaccess_status []}}

### HTTP Request

`POST /v1/card/access_status`

### Responses

#### Response body

| Name         | Type   | Description                                                   |
|--------------|--------|---------------------------------------------------------------|
| Success      | bool   | Success is a boolean indicating the success of the operation. |
| ErrorCode    | string | ErrorCode is the code of the error.                           |
| ErrorMessage | string | ErrorMessage is the message of the error.                     |

Example:

```json
{
  "success": "bool",
  "error_code": "string",
  "error_message": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update card status

Update Card status.

```sh
curl -X POST \
	/v1/card/status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet updatecardstatus []}}

### HTTP Request

`POST /v1/card/status`

### Responses

#### Response body

| Name         | Type   | Description                                                   |
|--------------|--------|---------------------------------------------------------------|
| Success      | bool   | Success is a boolean indicating the success of the operation. |
| ErrorCode    | string | ErrorCode is the code of the error.                           |
| ErrorMessage | string | ErrorMessage is the message of the error.                     |

Example:

```json
{
  "success": "bool",
  "error_code": "string",
  "error_message": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
