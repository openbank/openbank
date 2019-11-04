# Consent API v1.0.0

Provides CRUD operations on the consent part resource.

*
Host ``
EOL

*
Base Path ``

## Answer the consent reqeust challenge

Answer the consent request challenge

```sh
curl -X POST \
	/v1/banks/{BankID}/consents/{ConsentID}/challenge \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"consent_id": "string",
		"answer": "string"
	}'
```
{{snippet answerconsentchallenge []}}

### HTTP Request

`POST /v1/banks/{BankID}/consents/{ConsentID}/challenge`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| ConsentID | string |             |

### Body Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| ConsentID | string |             |
| Answer    | string |             |

### Responses

#### Response body

| Name      | Type   | Description |
|-----------|--------|-------------|
| ConsentID | string |             |
| Jwt       | string |             |
| Status    | Status |             |

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Answered the consent request challenge                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an email consent

Creates a new email consent

```sh
curl -X POST \
	/v1/banks/{BankID}/consents/email \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"for": "string",
		"view": "string",
		"email": "string"
	}'
```
{{snippet createconsentemail []}}

### HTTP Request

`POST /v1/banks/{BankID}/consents/email`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Body Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |
| For    | string |             |
| View   | string |             |
| Email  | string |             |

### Responses

#### Response body

| Name      | Type   | Description |
|-----------|--------|-------------|
| ConsentID | string |             |
| Jwt       | string |             |
| Status    | Status |             |

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Email consent created successfully.                                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an sms consent

Creates a new sms consent

```sh
curl -X POST \
	/v1/banks/{BankID}/consents/sms \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"for": "string",
		"view": "string",
		"phone_number": "string"
	}'
```
{{snippet createconsentsms []}}

### HTTP Request

`POST /v1/banks/{BankID}/consents/sms`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| For         | string |             |
| View        | string |             |
| PhoneNumber | string |             |

### Responses

#### Response body

| Name      | Type   | Description |
|-----------|--------|-------------|
| ConsentID | string |             |
| Jwt       | string |             |
| Status    | Status |             |

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | SMS consent created successfully.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all accounts

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/banks/{BankID}/consents \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks/{BankID}/consents`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name     | Type      | Description |
|----------|-----------|-------------|
| Consents | []Consent |             |

##### Objects

###### Consent

| Name      | Type   | Description |
|-----------|--------|-------------|
| ConsentID | string |             |
| Jwt       | string |             |
| Status    | Status |             |

Example:

```json
{
  "consents": [
    {
      "consent_id": "string",
      "jwt": "string",
      "status": "Status"
    }
  ]
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Revoke the consent

Revoke the consent

```sh
curl -X POST \
	/v1/banks/{BankID}/consents/{ConsentID}/revoke \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"consent_id": "string"
	}'
```
### HTTP Request

`POST /v1/banks/{BankID}/consents/{ConsentID}/revoke`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| ConsentID | string |             |

### Body Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| ConsentID | string |             |

### Responses

#### Response body

| Name      | Type   | Description |
|-----------|--------|-------------|
| ConsentID | string |             |
| Jwt       | string |             |
| Status    | Status |             |

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Consent revoked successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  Status

| Value           | Description |
|-----------------|-------------|
| UnknownStatus   |             |
| INITIATED       |             |
| ACCEPTED        |             |
| REJECTED        |             |
| REVOKED         |             |
| RECEIVED        |             |
| VALID           |             |
| REVOKEDBYPSU    |             |
| EXPIRED         |             |
| TERMINATEDBYTPP |             |
