KYC API v1.0.0
==============

Provides create and read operations on the KYC resource.

* Host `https://`

* Base Path ``

Add KYC check {#method-put-addkyccheck}
---------------------------------------

Add KYC check for the customer specified

```sh
curl -X PUT \
	https:///v1/customers/{CustomerID}/kyc_check/{KYCCheckID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"kyc_check_id": "string",
		"customer_id": "string",
		"customer_number": "string",
		"date": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"how": "string",
		"staff_user_id": "string",
		"staff_name": "string",
		"satisfied": "bool",
		"comments": "string"
	}'
```

### HTTP Request {#http-request-method-put-addkyccheck}

`PUT https:///v1/customers/{CustomerID}/kyc_check/{KYCCheckID}`

### Query Parameters {#query-parameters-method-put-addkyccheck}

| Name         | Type   | Description |
|--------------|--------|-------------|
| customer_id  | string |             |
| kyc_check_id | string |             |

### Body Parameters {#body-parameters-method-put-addkyccheck}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| kyc_check_id    | string    |             |
| customer_id     | string    |             |
| customer_number | string    |             |
| date            | Timestamp |             |
| how             | string    |             |
| staff_user_id   | string    |             |
| staff_name      | string    |             |
| satisfied       | bool      |             |
| comments        | string    |             |

##### Objects {#objects-AddKYCCheckRequest}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-put-addkyccheck}

#### Response body {#response-body-method-put-addkyccheck}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| bank_id         | string    |             |
| customer_id     | string    |             |
| id              | string    |             |
| customer_number | string    |             |
| date            | Timestamp |             |
| how             | string    |             |
| staff_user_id   | string    |             |
| staff_name      | string    |             |
| satisfied       | bool      |             |
| comments        | string    |             |

##### Objects {#objects-AddKYCCheckResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "bank_id": "string",
  "customer_id": "string",
  "id": "string",
  "customer_number": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "how": "string",
  "staff_user_id": "string",
  "staff_name": "string",
  "satisfied": "bool",
  "comments": "string"
}
```

#### Response codes {#response-codes-method-put-addkyccheck}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC check added successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Add KYC document {#method-put-addkycdocument}
---------------------------------------------

Add KYC document for the customer specified

```sh
curl -X PUT \
	https:///v1/customers/{CustomerID}/kyc_check/{KYCDocumentID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"kyc_document_id": "string",
		"customer_id": "string",
		"customer_number": "string",
		"type": "string",
		"number": "string",
		"issue_date": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"issue_place": "string",
		"expiry_date": {
			"seconds": "int64",
			"nanos": "int32"
		}
	}'
```

### HTTP Request {#http-request-method-put-addkycdocument}

`PUT https:///v1/customers/{CustomerID}/kyc_check/{KYCDocumentID}`

### Query Parameters {#query-parameters-method-put-addkycdocument}

| Name            | Type   | Description |
|-----------------|--------|-------------|
| customer_id     | string |             |
| kyc_document_id | string |             |

### Body Parameters {#body-parameters-method-put-addkycdocument}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| kyc_document_id | string    |             |
| customer_id     | string    |             |
| customer_number | string    |             |
| type            | string    |             |
| number          | string    |             |
| issue_date      | Timestamp |             |
| issue_place     | string    |             |
| expiry_date     | Timestamp |             |

##### Objects {#objects-AddKYCDocumentRequest}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-put-addkycdocument}

#### Response body {#response-body-method-put-addkycdocument}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| bank_id         | string    |             |
| customer_id     | string    |             |
| id              | string    |             |
| customer_number | string    |             |
| type            | string    |             |
| number          | string    |             |
| issue_date      | Timestamp |             |
| issue_place     | string    |             |
| expiry_date     | Timestamp |             |

##### Objects {#objects-AddKYCDocumentResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "bank_id": "string",
  "customer_id": "string",
  "id": "string",
  "customer_number": "string",
  "type": "string",
  "number": "string",
  "issue_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "issue_place": "string",
  "expiry_date": {
    "seconds": "int64",
    "nanos": "int32"
  }
}
```

#### Response codes {#response-codes-method-put-addkycdocument}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC Document added successfully.                                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Add KYC media {#method-put-addkycmedia}
---------------------------------------

Add KYC media for the customer specified

```sh
curl -X PUT \
	https:///v1/customers/{CustomerID}/kyc_check/{KYCMediaID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"kyc_media_id": "string",
		"customer_id": "string",
		"customer_number": "string",
		"type": "string",
		"url": "string",
		"date": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"relates_to_kyc_document_id": "string",
		"relates_to_kyc_check_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-addkycmedia}

`PUT https:///v1/customers/{CustomerID}/kyc_check/{KYCMediaID}`

### Query Parameters {#query-parameters-method-put-addkycmedia}

| Name         | Type   | Description |
|--------------|--------|-------------|
| customer_id  | string |             |
| kyc_media_id | string |             |

### Body Parameters {#body-parameters-method-put-addkycmedia}

| Name                       | Type      | Description |
|----------------------------|-----------|-------------|
| kyc_media_id               | string    |             |
| customer_id                | string    |             |
| customer_number            | string    |             |
| type                       | string    |             |
| url                        | string    |             |
| date                       | Timestamp |             |
| relates_to_kyc_document_id | string    |             |
| relates_to_kyc_check_id    | string    |             |

##### Objects {#objects-AddKYCMediaRequest}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-put-addkycmedia}

#### Response body {#response-body-method-put-addkycmedia}

| Name                       | Type      | Description |
|----------------------------|-----------|-------------|
| bank_id                    | string    |             |
| customer_id                | string    |             |
| id                         | string    |             |
| customer_number            | string    |             |
| type                       | string    |             |
| url                        | string    |             |
| date                       | Timestamp |             |
| relates_to_kyc_document_id | string    |             |
| relates_to_kyc_check_id    | string    |             |

##### Objects {#objects-AddKYCMediaResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "bank_id": "string",
  "customer_id": "string",
  "id": "string",
  "customer_number": "string",
  "type": "string",
  "url": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "relates_to_kyc_document_id": "string",
  "relates_to_kyc_check_id": "string"
}
```

#### Response codes {#response-codes-method-put-addkycmedia}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC Media added successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Add KYC status {#method-put-addkycstatus}
-----------------------------------------

Add KYC status for the customer specified

```sh
curl -X PUT \
	https:///v1/customers/{CustomerID}/kyc_statuses \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"customer_id": "string",
		"customer_number": "string",
		"ok": "bool",
		"date": {
			"seconds": "int64",
			"nanos": "int32"
		}
	}'
```

### HTTP Request {#http-request-method-put-addkycstatus}

`PUT https:///v1/customers/{CustomerID}/kyc_statuses`

### Query Parameters {#query-parameters-method-put-addkycstatus}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-addkycstatus}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| customer_id     | string    |             |
| customer_number | string    |             |
| ok              | bool      |             |
| date            | Timestamp |             |

##### Objects {#objects-AddKYCStatusRequest}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-put-addkycstatus}

#### Response body {#response-body-method-put-addkycstatus}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| customer_id     | string    |             |
| customer_number | string    |             |
| ok              | bool      |             |
| date            | Timestamp |             |

##### Objects {#objects-AddKYCStatusResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "customer_id": "string",
  "customer_number": "string",
  "ok": "bool",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  }
}
```

#### Response codes {#response-codes-method-put-addkycstatus}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC status added successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get Customer KYC check {#method-get-getcustomerkyccheck}
--------------------------------------------------------

Get Customer KYC check fot the customer

```sh
curl -X GET \
	https:///v1/customers/{CustomerID}/kyc_checks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomerkyccheck}

`GET https:///v1/customers/{CustomerID}/kyc_checks`

### Query Parameters {#query-parameters-method-get-getcustomerkyccheck}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Responses {#responses-method-get-getcustomerkyccheck}

#### Response body {#response-body-method-get-getcustomerkyccheck}

| Name   | Type     | Description |
|--------|----------|-------------|
| checks | \[]Check |             |

##### Objects {#objects-GetCustomerKYCCheckResponse}

###### Check

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| bank_id         | string    |             |
| customer_id     | string    |             |
| id              | string    |             |
| customer_number | string    |             |
| date            | Timestamp |             |
| how             | string    |             |
| staff_user_id   | string    |             |
| staff_name      | string    |             |
| satisfied       | bool      |             |
| comments        | string    |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "checks": [
    {
      "bank_id": "string",
      "customer_id": "string",
      "id": "string",
      "customer_number": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "how": "string",
      "staff_user_id": "string",
      "staff_name": "string",
      "satisfied": "bool",
      "comments": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomerkyccheck}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get Customer KYC documents {#method-get-getcustomerkycdocument}
---------------------------------------------------------------

Get Customer KYC documents fot the customer

```sh
curl -X GET \
	https:///v1/customers/{CustomerID}/kyc_documents \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomerkycdocument}

`GET https:///v1/customers/{CustomerID}/kyc_documents`

### Query Parameters {#query-parameters-method-get-getcustomerkycdocument}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Responses {#responses-method-get-getcustomerkycdocument}

#### Response body {#response-body-method-get-getcustomerkycdocument}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| documents | \[]Document |             |

##### Objects {#objects-GetCustomerKYCDocumentResponse}

###### Document

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| bank_id         | string    |             |
| customer_id     | string    |             |
| id              | string    |             |
| customer_number | string    |             |
| type            | string    |             |
| number          | string    |             |
| issue_date      | Timestamp |             |
| issue_place     | string    |             |
| expiry_date     | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "documents": [
    {
      "bank_id": "string",
      "customer_id": "string",
      "id": "string",
      "customer_number": "string",
      "type": "string",
      "number": "string",
      "issue_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "issue_place": "string",
      "expiry_date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomerkycdocument}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get Customer KYC statuses {#method-get-getcustomerkycstatus}
------------------------------------------------------------

Get Customer KYC statuses fot the customer

```sh
curl -X GET \
	https:///v1/customers/{CustomerID}/kyc_statuses \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomerkycstatus}

`GET https:///v1/customers/{CustomerID}/kyc_statuses`

### Query Parameters {#query-parameters-method-get-getcustomerkycstatus}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Responses {#responses-method-get-getcustomerkycstatus}

#### Response body {#response-body-method-get-getcustomerkycstatus}

| Name     | Type      | Description |
|----------|-----------|-------------|
| statuses | \[]Status |             |

##### Objects {#objects-GetCustomerKYCStatusResponse}

###### Status

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| customer_id     | string    |             |
| customer_number | string    |             |
| ok              | bool      |             |
| date            | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "statuses": [
    {
      "customer_id": "string",
      "customer_number": "string",
      "ok": "bool",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomerkycstatus}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get Customer KYC media {#method-get-getkycmedia}
------------------------------------------------

Get Customer KYC media fot the customer

```sh
curl -X GET \
	https:///v1/customers/{CustomerID}/kyc_media \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getkycmedia}

`GET https:///v1/customers/{CustomerID}/kyc_media`

### Query Parameters {#query-parameters-method-get-getkycmedia}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Responses {#responses-method-get-getkycmedia}

#### Response body {#response-body-method-get-getkycmedia}

| Name   | Type     | Description |
|--------|----------|-------------|
| medias | \[]Media |             |

##### Objects {#objects-GetKYCMediaResponse}

###### Media

| Name                       | Type      | Description |
|----------------------------|-----------|-------------|
| bank_id                    | string    |             |
| customer_id                | string    |             |
| id                         | string    |             |
| customer_number            | string    |             |
| type                       | string    |             |
| url                        | string    |             |
| date                       | Timestamp |             |
| relates_to_kyc_document_id | string    |             |
| relates_to_kyc_check_id    | string    |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "medias": [
    {
      "bank_id": "string",
      "customer_id": "string",
      "id": "string",
      "customer_number": "string",
      "type": "string",
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "relates_to_kyc_document_id": "string",
      "relates_to_kyc_check_id": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getkycmedia}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
