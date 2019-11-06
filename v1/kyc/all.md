# KYC API v1.0.0

Provides create and read operations on the KYC resource.

*
Host ``
EOL

*
Base Path ``

## Add KYC check

Add KYC check for the customer specified

```sh
curl -X PUT \
	/v1/customers/{CustomerID}/kyc_check/{KYCCheckID} \
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
{{snippet addkyccheck []}}

### HTTP Request

`PUT /v1/customers/{CustomerID}/kyc_check/{KYCCheckID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |
| KYCCheckID | string |             |

### Body Parameters

| Name           | Type      | Description |
|----------------|-----------|-------------|
| KYCCheckID     | string    |             |
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Date           | Timestamp |             |
| How            | string    |             |
| StaffUserID    | string    |             |
| StaffName      | string    |             |
| Satisfied      | bool      |             |
| Comments       | string    |             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name           | Type      | Description |
|----------------|-----------|-------------|
| BankID         | string    |             |
| CustomerID     | string    |             |
| ID             | string    |             |
| CustomerNumber | string    |             |
| Date           | Timestamp |             |
| How            | string    |             |
| StaffUserID    | string    |             |
| StaffName      | string    |             |
| Satisfied      | bool      |             |
| Comments       | string    |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC check added successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add KYC document

Add KYC document for the customer specified

```sh
curl -X PUT \
	/v1/customers/{CustomerID}/kyc_check/{KYCDocumentID} \
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
{{snippet addkycdocument []}}

### HTTP Request

`PUT /v1/customers/{CustomerID}/kyc_check/{KYCDocumentID}`

### Query Parameters

| Name          | Type   | Description |
|---------------|--------|-------------|
| CustomerID    | string |             |
| KYCDocumentID | string |             |

### Body Parameters

| Name           | Type      | Description |
|----------------|-----------|-------------|
| KYCDocumentID  | string    |             |
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Type           | string    |             |
| Number         | string    |             |
| IssueDate      | Timestamp |             |
| IssuePlace     | string    |             |
| ExpiryDate     | Timestamp |             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name           | Type      | Description |
|----------------|-----------|-------------|
| BankID         | string    |             |
| CustomerID     | string    |             |
| ID             | string    |             |
| CustomerNumber | string    |             |
| Type           | string    |             |
| Number         | string    |             |
| IssueDate      | Timestamp |             |
| IssuePlace     | string    |             |
| ExpiryDate     | Timestamp |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC Document added successfully.                                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add KYC media

Add KYC media for the customer specified

```sh
curl -X PUT \
	/v1/customers/{CustomerID}/kyc_check/{KYCMediaID} \
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
{{snippet addkycmedia []}}

### HTTP Request

`PUT /v1/customers/{CustomerID}/kyc_check/{KYCMediaID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |
| KYCMediaID | string |             |

### Body Parameters

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| KYCMediaID             | string    |             |
| CustomerID             | string    |             |
| CustomerNumber         | string    |             |
| Type                   | string    |             |
| URL                    | string    |             |
| Date                   | Timestamp |             |
| RelatesToKycDocumentID | string    |             |
| RelatesToKycCheckID    | string    |             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| BankID                 | string    |             |
| CustomerID             | string    |             |
| ID                     | string    |             |
| CustomerNumber         | string    |             |
| Type                   | string    |             |
| URL                    | string    |             |
| Date                   | Timestamp |             |
| RelatesToKycDocumentID | string    |             |
| RelatesToKycCheckID    | string    |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC Media added successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add KYC status

Add KYC status for the customer specified

```sh
curl -X PUT \
	/v1/customers/{CustomerID}/kyc_statuses \
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
{{snippet addkycstatus []}}

### HTTP Request

`PUT /v1/customers/{CustomerID}/kyc_statuses`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Body Parameters

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Ok             | bool      |             |
| Date           | Timestamp |             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Ok             | bool      |             |
| Date           | Timestamp |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | KYC status added successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer KYC check

Get Customer KYC check fot the customer

```sh
curl -X GET \
	/v1/customers/{CustomerID}/kyc_checks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomerkyc_check []}}

### HTTP Request

`GET /v1/customers/{CustomerID}/kyc_checks`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Responses

#### Response body

| Name   | Type    | Description |
|--------|---------|-------------|
| Checks | []Check |             |

##### Objects

###### Check

| Name           | Type      | Description |
|----------------|-----------|-------------|
| BankID         | string    |             |
| CustomerID     | string    |             |
| ID             | string    |             |
| CustomerNumber | string    |             |
| Date           | Timestamp |             |
| How            | string    |             |
| StaffUserID    | string    |             |
| StaffName      | string    |             |
| Satisfied      | bool      |             |
| Comments       | string    |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer KYC documents

Get Customer KYC documents fot the customer

```sh
curl -X GET \
	/v1/customers/{CustomerID}/kyc_documents \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomerkyc_document []}}

### HTTP Request

`GET /v1/customers/{CustomerID}/kyc_documents`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Documents | []Document |             |

##### Objects

###### Document

| Name           | Type      | Description |
|----------------|-----------|-------------|
| BankID         | string    |             |
| CustomerID     | string    |             |
| ID             | string    |             |
| CustomerNumber | string    |             |
| Type           | string    |             |
| Number         | string    |             |
| IssueDate      | Timestamp |             |
| IssuePlace     | string    |             |
| ExpiryDate     | Timestamp |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer KYC statuses

Get Customer KYC statuses fot the customer

```sh
curl -X GET \
	/v1/customers/{CustomerID}/kyc_statuses \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomerkyc_status []}}

### HTTP Request

`GET /v1/customers/{CustomerID}/kyc_statuses`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Statuses | []Status |             |

##### Objects

###### Status

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Ok             | bool      |             |
| Date           | Timestamp |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer KYC media

Get Customer KYC media fot the customer

```sh
curl -X GET \
	/v1/customers/{CustomerID}/kyc_media \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getkycmedia []}}

### HTTP Request

`GET /v1/customers/{CustomerID}/kyc_media`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Responses

#### Response body

| Name   | Type    | Description |
|--------|---------|-------------|
| Medias | []Media |             |

##### Objects

###### Media

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| BankID                 | string    |             |
| CustomerID             | string    |             |
| ID                     | string    |             |
| CustomerNumber         | string    |             |
| Type                   | string    |             |
| URL                    | string    |             |
| Date                   | Timestamp |             |
| RelatesToKycDocumentID | string    |             |
| RelatesToKycCheckID    | string    |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
