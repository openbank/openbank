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

### HTTP Request {#http-request-method-post-answerconsentchallenge}

`POST https:///v1/banks/{BankID}/consents/{ConsentID}/challenge`

### Query Parameters {#query-parameters-method-post-answerconsentchallenge}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Body Parameters {#body-parameters-method-post-answerconsentchallenge}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |
| answer     | string |             |

### Responses {#responses-method-post-answerconsentchallenge}

#### Response body {#response-body-method-post-answerconsentchallenge}

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

##### Enums {#enums-Consent}

###### Status

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

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```

#### Response codes {#response-codes-method-post-answerconsentchallenge}

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

### HTTP Request {#http-request-method-post-createconsentemail}

`POST https:///v1/banks/{BankID}/consents/email`

### Query Parameters {#query-parameters-method-post-createconsentemail}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Body Parameters {#body-parameters-method-post-createconsentemail}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |
| for     | string |             |
| view    | string |             |
| email   | string |             |

### Responses {#responses-method-post-createconsentemail}

#### Response body {#response-body-method-post-createconsentemail}

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

##### Enums {#enums-Consent}

###### Status

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

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```

#### Response codes {#response-codes-method-post-createconsentemail}

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

### HTTP Request {#http-request-method-post-createconsentsms}

`POST https:///v1/banks/{BankID}/consents/sms`

### Query Parameters {#query-parameters-method-post-createconsentsms}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Body Parameters {#body-parameters-method-post-createconsentsms}

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| for          | string |             |
| view         | string |             |
| phone_number | string |             |

### Responses {#responses-method-post-createconsentsms}

#### Response body {#response-body-method-post-createconsentsms}

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

##### Enums {#enums-Consent}

###### Status

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

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```

#### Response codes {#response-codes-method-post-createconsentsms}

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

### HTTP Request {#http-request-method-get-getconsents}

`GET https:///v1/banks/{BankID}/consents`

### Query Parameters {#query-parameters-method-get-getconsents}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getconsents}

#### Response body {#response-body-method-get-getconsents}

| Name     | Type       | Description |
|----------|------------|-------------|
| consents | \[]Consent |             |

##### Objects {#objects-GetConsentsResponse}

###### Consent

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

##### Enums {#enums-GetConsentsResponse}

###### Status

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

#### Response codes {#response-codes-method-get-getconsents}

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

### HTTP Request {#http-request-method-post-revokeconsent}

`POST https:///v1/banks/{BankID}/consents/{ConsentID}/revoke`

### Query Parameters {#query-parameters-method-post-revokeconsent}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Body Parameters {#body-parameters-method-post-revokeconsent}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| consent_id | string |             |

### Responses {#responses-method-post-revokeconsent}

#### Response body {#response-body-method-post-revokeconsent}

| Name       | Type   | Description |
|------------|--------|-------------|
| consent_id | string |             |
| jwt        | string |             |
| status     | Status |             |

##### Enums {#enums-Consent}

###### Status

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

Example:

```json
{
  "consent_id": "string",
  "jwt": "string",
  "status": "Status"
}
```

#### Response codes {#response-codes-method-post-revokeconsent}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Consent revoked successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
