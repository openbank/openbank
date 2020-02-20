# Card API v1.0.0

Provides create and read operations on the card resource.

* Host ``

* Base Path ``

## Create new card {#method-post-createcard}

Create new card

```sh
curl -X POST \
	/v1/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_number": "string",
		"account_id": "string",
		"contact_number": "string",
		"first_name": "string",
		"last_name": "string"
	}'
```
### HTTP Request

`POST /v1/card`

### Body Parameters

| Name          | Type   | Description                                                  |
|---------------|--------|--------------------------------------------------------------|
| CardNumber    | string | CardNumber is the unique identification number of each card. |
| AccountID     | string | Account is the ID of the account associated with the card.   |
| ContactNumber | string | ContactNumber is the contact number of the card owner.       |
| FirstName     | string | FirstName is the first name of card owner.                   |
| LastName      | string | LastName is the last name of the card owner.                 |

### Responses

#### Response body

| Name   | Type   | Description |
|--------|--------|-------------|
| CardID | string |             |
| Expiry | string |             |

Example:

```json
{
  "card_id": "string",
  "expiry": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Card Created                                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create new card attribute {#method-post-createcardattribute}

Create new card attribute

```sh
curl -X POST \
	/v1/card/attribute \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_id": "string",
		"name": "string",
		"type": "string",
		"value": "string"
	}'
```
{{snippet createcardattribute []}}

### HTTP Request

`POST /v1/card/attribute`

### Body Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| CardID | string |             |
| Name   | string |             |
| Type   | string |             |
| Value  | string |             |

### Responses

#### Response body

| Name        | Type   | Description |
|-------------|--------|-------------|
| AttributeID | string |             |
| CardID      | string |             |

Example:

```json
{
  "attribute_id": "string",
  "card_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Card Attribute Created                                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a card {#method-delete-deletecard}

Permanently delete a card.

```sh
curl -X DELETE \
	/v1/card/{CardID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/card/{CardID}`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| CardID | string |             |

### Responses

#### Response body

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Card successfully deleted.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve card information {#method-get-getcard}

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

| Name          | Type             | Description                                                           |
|---------------|------------------|-----------------------------------------------------------------------|
| ID            | string           | ID is a unique identifier of a card.                                  |
| Account       | string           | Account is the ID of the account associated with the card.            |
| CardNumber    | string           | CardNumber is the unique identification number of each card.          |
| OwnerName     | string           | OwnerName is the name of the card owner.                              |
| ContactNumber | string           | ContactNumber is the contact number of the card owner.                |
| FirstName     | string           | FirstName is the first name of card owner.                            |
| LastName      | string           | LastName is the last name of the card owner.                          |
| Expiry        | Timestamp        | Expiry is an expiry date of the card.                                 |
| Status        | CardStatus       | Status is the status of the card.                                     |
| AccessStatus  | CardAccessStatus | AccessStatus is the access status of the card.                        |
| AmountDue     | Amount           | AmountDue is the the card holder is expected to paid by the due date. |
| CreditLimit   | string           | CreditLimit is the allowed credit limit.                              |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

Example:

```json
{
  "id": "string",
  "account": "string",
  "card_number": "string",
  "owner_name": "string",
  "contact_number": "string",
  "first_name": "string",
  "last_name": "string",
  "expiry": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "CardStatus",
  "access_status": "CardAccessStatus",
  "amount_due": {
    "cur": "string",
    "num": "string"
  },
  "credit_limit": "string"
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

## Retrieves all available cards {#method-get-getusercards}

Retrieves all available cards for specific user, selected by the user_id

```sh
curl -X GET \
	/v1/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getusercards []}}

### HTTP Request

`GET /v1/card`

### Responses

#### Response body

| Name   | Type   | Description |
|--------|--------|-------------|
| Result | []Card |             |

##### Objects

###### Card

| Name          | Type             | Description                                                           |
|---------------|------------------|-----------------------------------------------------------------------|
| ID            | string           | ID is a unique identifier of a card.                                  |
| Account       | string           | Account is the ID of the account associated with the card.            |
| CardNumber    | string           | CardNumber is the unique identification number of each card.          |
| OwnerName     | string           | OwnerName is the name of the card owner.                              |
| ContactNumber | string           | ContactNumber is the contact number of the card owner.                |
| FirstName     | string           | FirstName is the first name of card owner.                            |
| LastName      | string           | LastName is the last name of the card owner.                          |
| Expiry        | Timestamp        | Expiry is an expiry date of the card.                                 |
| Status        | CardStatus       | Status is the status of the card.                                     |
| AccessStatus  | CardAccessStatus | AccessStatus is the access status of the card.                        |
| AmountDue     | Amount           | AmountDue is the the card holder is expected to paid by the due date. |
| CreditLimit   | string           | CreditLimit is the allowed credit limit.                              |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

Example:

```json
{
  "result": [
    {
      "id": "string",
      "account": "string",
      "card_number": "string",
      "owner_name": "string",
      "contact_number": "string",
      "first_name": "string",
      "last_name": "string",
      "expiry": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "status": "CardStatus",
      "access_status": "CardAccessStatus",
      "amount_due": {
        "cur": "string",
        "num": "string"
      },
      "credit_limit": "string"
    }
  ]
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

## Update card access status. {#method-put-updatecardaccessstatus}

Update card access status.

```sh
curl -X PUT \
	/v1/card/access_status/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_token": "string",
		"access_status": "CardAccessStatus"
	}'
```
{{snippet updatecardaccess_status []}}

### HTTP Request

`PUT /v1/card/access_status/{CardToken}`

### Query Parameters

| Name      | Type   | Description                                              |
|-----------|--------|----------------------------------------------------------|
| CardToken | string | CardToken is the identifier to get the card information. |

### Body Parameters

| Name         | Type             | Description                                              |
|--------------|------------------|----------------------------------------------------------|
| CardToken    | string           | CardToken is the identifier to get the card information. |
| AccessStatus | CardAccessStatus | AccessStatus is the new card access status.              |

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

## Update card status {#method-put-updatecardstatus}

Update Card status.

```sh
curl -X PUT \
	/v1/card/status/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_token": "string",
		"status": "CardStatus"
	}'
```
{{snippet updatecardstatus []}}

### HTTP Request

`PUT /v1/card/status/{CardToken}`

### Query Parameters

| Name      | Type   | Description                                              |
|-----------|--------|----------------------------------------------------------|
| CardToken | string | CardToken is the identifier to get the card information. |

### Body Parameters

| Name      | Type       | Description                                              |
|-----------|------------|----------------------------------------------------------|
| CardToken | string     | CardToken is the identifier to get the card information. |
| Status    | CardStatus | Status is the new card status.                           |

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
