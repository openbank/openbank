Consent API v1.0.0
==================

Provides CRUD operations on the consent part resource.

* Host `https://`

* Base Path ``

Answer the consent reqeust challenge {#method-post-answerconsentchallenge}
--------------------------------------------------------------------------

Answer the consent request challenge

```sh
curl -X POST \
	https:///v1/banks/{BankID}/consents/{ConsentID}/challenge \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"consent_id": "string",
		"answer": "string"
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/consents/{ConsentID}/challenge`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |
| answer     | string |             |

### Responses

#### Response body

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

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

Create an email consent {#method-post-createconsentemail}
---------------------------------------------------------

Creates a new email consent

```sh
curl -X POST \
	https:///v1/banks/{BankID}/consents/email \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"for": "string",
		"view": "string",
		"email": "string"
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/consents/email`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Body Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |
| for     | string |             |
| view    | string |             |
| email   | string |             |

### Responses

#### Response body

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

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

Create an sms consent {#method-post-createconsentsms}
-----------------------------------------------------

Creates a new sms consent

```sh
curl -X POST \
	https:///v1/banks/{BankID}/consents/sms \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"for": "string",
		"view": "string",
		"phone_number": "string"
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/consents/sms`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Body Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| for          | string |             |
| view         | string |             |
| phone_number | string |             |

### Responses

#### Response body

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

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

List all consents {#method-get-getconsents}
-------------------------------------------

Returns a list containing up to 20 consents. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/consents \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/consents`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses

#### Response body

| Name     | Type       | Description |
|----------|------------|-------------|
| consents | \[]Consent |             |

##### Objects

###### Consent

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

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

Revoke the consent {#method-post-revokeconsent}
-----------------------------------------------

Revoke the consent

```sh
curl -X POST \
	https:///v1/banks/{BankID}/consents/{ConsentID}/revoke \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"consent_id": "string"
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/consents/{ConsentID}/revoke`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Responses

#### Response body

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

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

Annex
-----

#### Status

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
