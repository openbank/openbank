# Card API v1.0.0

Provides create and read operations on the card resource.

* Host `https://`

* Base Path ``

## Create new card {#method-post-createcard}

Create new card

```sh
curl -X POST \
	https:///v1/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_number": "string",
		"account_id": "string",
		"contact_number": "string",
		"first_name": "string",
		"last_name": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcard}

`POST https:///v1/card`

### Body Parameters {#body-parameters-method-post-createcard}

| Name           | Type   | Description                                                  |
|----------------|--------|--------------------------------------------------------------|
| card_number    | string | CardNumber is the unique identification number of each card. |
| account_id     | string | Account is the ID of the account associated with the card.   |
| contact_number | string | ContactNumber is the contact number of the card owner.       |
| first_name     | string | FirstName is the first name of card owner.                   |
| last_name      | string | LastName is the last name of the card owner.                 |

### Responses {#responses-method-post-createcard}

#### Response body {#response-body-method-post-createcard}

| Name    | Type   | Description |
|---------|--------|-------------|
| card_id | string |             |
| expiry  | string |             |

Example:

```json
{
  "card_id": "string",
  "expiry": "string"
}
```

#### Response codes {#response-codes-method-post-createcard}

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
	https:///v1/card/attribute \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_id": "string",
		"name": "string",
		"type": "string",
		"value": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcardattribute}

`POST https:///v1/card/attribute`

### Body Parameters {#body-parameters-method-post-createcardattribute}

| Name    | Type   | Description |
|---------|--------|-------------|
| card_id | string |             |
| name    | string |             |
| type    | string |             |
| value   | string |             |

### Responses {#responses-method-post-createcardattribute}

#### Response body {#response-body-method-post-createcardattribute}

| Name         | Type   | Description |
|--------------|--------|-------------|
| attribute_id | string |             |
| card_id      | string |             |

Example:

```json
{
  "attribute_id": "string",
  "card_id": "string"
}
```

#### Response codes {#response-codes-method-post-createcardattribute}

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
	https:///v1/card/{CardID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletecard}

`DELETE https:///v1/card/{CardID}`

### Query Parameters {#query-parameters-method-delete-deletecard}

| Name    | Type   | Description |
|---------|--------|-------------|
| card_id | string |             |

### Responses {#responses-method-delete-deletecard}

#### Response body {#response-body-method-delete-deletecard}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletecard}

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
	https:///v1/card/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcard}

`GET https:///v1/card/{CardToken}`

### Query Parameters {#query-parameters-method-get-getcard}

| Name       | Type   | Description                                              |
|------------|--------|----------------------------------------------------------|
| card_token | string | CardToken is the identifier to get the card information. |

### Responses {#responses-method-get-getcard}

#### Response body {#response-body-method-get-getcard}

| Name           | Type             | Description                                                           |
|----------------|------------------|-----------------------------------------------------------------------|
| id             | string           | ID is a unique identifier of a card.                                  |
| account        | string           | Account is the ID of the account associated with the card.            |
| card_number    | string           | CardNumber is the unique identification number of each card.          |
| owner_name     | string           | OwnerName is the name of the card owner.                              |
| contact_number | string           | ContactNumber is the contact number of the card owner.                |
| first_name     | string           | FirstName is the first name of card owner.                            |
| last_name      | string           | LastName is the last name of the card owner.                          |
| expiry         | Timestamp        | Expiry is an expiry date of the card.                                 |
| status         | CardStatus       | Status is the status of the card.                                     |
| access_status  | CardAccessStatus | AccessStatus is the access status of the card.                        |
| amount_due     | Amount           | AmountDue is the the card holder is expected to paid by the due date. |
| credit_limit   | string           | CreditLimit is the allowed credit limit.                              |

##### Objects {#objects-Card}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

##### Enums {#enums-Card}

###### CardStatus

| Value             | Description                             |
|-------------------|-----------------------------------------|
| UnknownCardStatus |                                         |
| Lock              | Lock is the lock status for a card.     |
| Unlock            | Unlock is the unlock status for a card. |
| Active            | Active is the active status for a card. |

###### CardAccessStatus

| Value                   | Description                           |
|-------------------------|---------------------------------------|
| UnknownCardAccessStatus |                                       |
| Often                   | Often indicates a card is used often. |
| Rare                    | Rare indicates a card is used rarely. |

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

#### Response codes {#response-codes-method-get-getcard}

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
	https:///v1/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getusercards}

`GET https:///v1/card`

### Responses {#responses-method-get-getusercards}

#### Response body {#response-body-method-get-getusercards}

| Name   | Type    | Description |
|--------|---------|-------------|
| result | \[]Card |             |

##### Objects {#objects-GetUserCardsResponse}

###### Card

| Name           | Type             | Description                                                           |
|----------------|------------------|-----------------------------------------------------------------------|
| id             | string           | ID is a unique identifier of a card.                                  |
| account        | string           | Account is the ID of the account associated with the card.            |
| card_number    | string           | CardNumber is the unique identification number of each card.          |
| owner_name     | string           | OwnerName is the name of the card owner.                              |
| contact_number | string           | ContactNumber is the contact number of the card owner.                |
| first_name     | string           | FirstName is the first name of card owner.                            |
| last_name      | string           | LastName is the last name of the card owner.                          |
| expiry         | Timestamp        | Expiry is an expiry date of the card.                                 |
| status         | CardStatus       | Status is the status of the card.                                     |
| access_status  | CardAccessStatus | AccessStatus is the access status of the card.                        |
| amount_due     | Amount           | AmountDue is the the card holder is expected to paid by the due date. |
| credit_limit   | string           | CreditLimit is the allowed credit limit.                              |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

##### Enums {#enums-GetUserCardsResponse}

###### CardStatus

| Value             | Description                             |
|-------------------|-----------------------------------------|
| UnknownCardStatus |                                         |
| Lock              | Lock is the lock status for a card.     |
| Unlock            | Unlock is the unlock status for a card. |
| Active            | Active is the active status for a card. |

###### CardAccessStatus

| Value                   | Description                           |
|-------------------------|---------------------------------------|
| UnknownCardAccessStatus |                                       |
| Often                   | Often indicates a card is used often. |
| Rare                    | Rare indicates a card is used rarely. |

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

#### Response codes {#response-codes-method-get-getusercards}

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
	https:///v1/card/access_status/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_token": "string",
		"access_status": "CardAccessStatus"
	}'
```

### HTTP Request {#http-request-method-put-updatecardaccessstatus}

`PUT https:///v1/card/access_status/{CardToken}`

### Query Parameters {#query-parameters-method-put-updatecardaccessstatus}

| Name       | Type   | Description                                              |
|------------|--------|----------------------------------------------------------|
| card_token | string | CardToken is the identifier to get the card information. |

### Body Parameters {#body-parameters-method-put-updatecardaccessstatus}

| Name          | Type             | Description                                              |
|---------------|------------------|----------------------------------------------------------|
| card_token    | string           | CardToken is the identifier to get the card information. |
| access_status | CardAccessStatus | AccessStatus is the new card access status.              |

##### Enums {#enums-UpdateCardAccessStatusRequest}

###### CardAccessStatus

| Value                   | Description                           |
|-------------------------|---------------------------------------|
| UnknownCardAccessStatus |                                       |
| Often                   | Often indicates a card is used often. |
| Rare                    | Rare indicates a card is used rarely. |

### Responses {#responses-method-put-updatecardaccessstatus}

#### Response body {#response-body-method-put-updatecardaccessstatus}

| Name          | Type   | Description                                                   |
|---------------|--------|---------------------------------------------------------------|
| success       | bool   | Success is a boolean indicating the success of the operation. |
| error_code    | string | ErrorCode is the code of the error.                           |
| error_message | string | ErrorMessage is the message of the error.                     |

Example:

```json
{
  "success": "bool",
  "error_code": "string",
  "error_message": "string"
}
```

#### Response codes {#response-codes-method-put-updatecardaccessstatus}

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
	https:///v1/card/status/{CardToken} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"card_token": "string",
		"status": "CardStatus"
	}'
```

### HTTP Request {#http-request-method-put-updatecardstatus}

`PUT https:///v1/card/status/{CardToken}`

### Query Parameters {#query-parameters-method-put-updatecardstatus}

| Name       | Type   | Description                                              |
|------------|--------|----------------------------------------------------------|
| card_token | string | CardToken is the identifier to get the card information. |

### Body Parameters {#body-parameters-method-put-updatecardstatus}

| Name       | Type       | Description                                              |
|------------|------------|----------------------------------------------------------|
| card_token | string     | CardToken is the identifier to get the card information. |
| status     | CardStatus | Status is the new card status.                           |

##### Enums {#enums-UpdateCardStatusRequest}

###### CardStatus

| Value             | Description                             |
|-------------------|-----------------------------------------|
| UnknownCardStatus |                                         |
| Lock              | Lock is the lock status for a card.     |
| Unlock            | Unlock is the unlock status for a card. |
| Active            | Active is the active status for a card. |

### Responses {#responses-method-put-updatecardstatus}

#### Response body {#response-body-method-put-updatecardstatus}

| Name          | Type   | Description                                                   |
|---------------|--------|---------------------------------------------------------------|
| success       | bool   | Success is a boolean indicating the success of the operation. |
| error_code    | string | ErrorCode is the code of the error.                           |
| error_message | string | ErrorMessage is the message of the error.                     |

Example:

```json
{
  "success": "bool",
  "error_code": "string",
  "error_message": "string"
}
```

#### Response codes {#response-codes-method-put-updatecardstatus}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
