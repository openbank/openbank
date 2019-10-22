# Account Application API v1.0.0

Provides create and read operations on the account application resource.

*
Host ``
EOL

*
Base Path ``

## Create an account application

Creates a new account application

```sh
curl -X POST \
	/v1/account-applications \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"product_code": "string",
		"user_id": "string",
		"profile_id": "string"
	}'
```
{{snippet createaccountapplication []}}

### HTTP Request

`POST /v1/account-applications`

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| ProductCode | string |             |
| UserID      | string |             |
| ProfileID   | string |             |

### Responses

#### Response body

| Name              | Type                     | Description |
|-------------------|--------------------------|-------------|
| ID                | string                   |             |
| ProductCode       | string                   |             |
| UserID            | string                   |             |
| CustomerProfileID | string                   |             |
| DateOfApplication | Timestamp                |             |
| Status            | AccountApplicationStatus |             |

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
  "product_code": "string",
  "user_id": "string",
  "customer_profile_id": "string",
  "date_of_application": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "AccountApplicationStatus"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve an account application

Retrieves all data from an account application selected by the supplied accountapplicationid.

```sh
curl -X GET \
	/v1/account-applications/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getaccountapplication []}}

### HTTP Request

`GET /v1/account-applications/{ID}`

### Query Parameters

| Name | Type   | Description |
|------|--------|-------------|
| ID   | string |             |

### Responses

#### Response body

| Name              | Type                     | Description |
|-------------------|--------------------------|-------------|
| ID                | string                   |             |
| ProductCode       | string                   |             |
| UserID            | string                   |             |
| CustomerProfileID | string                   |             |
| DateOfApplication | Timestamp                |             |
| Status            | AccountApplicationStatus |             |

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
  "product_code": "string",
  "user_id": "string",
  "customer_profile_id": "string",
  "date_of_application": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "AccountApplicationStatus"
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

## List all accounts

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/account-applications \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getaccountapplications []}}

### HTTP Request

`GET /v1/account-applications`

### Responses

#### Response body

| Name    | Type                 | Description                                             |
|---------|----------------------|---------------------------------------------------------|
| Result  | []AccountApplication | Result is a list containing up to 20 Accounts.          |
| HasMore | bool                 | HasMore indicates if there are more accounts available. |

##### Objects

###### AccountApplication

| Name              | Type                     | Description |
|-------------------|--------------------------|-------------|
| ID                | string                   |             |
| ProductCode       | string                   |             |
| UserID            | string                   |             |
| CustomerProfileID | string                   |             |
| DateOfApplication | Timestamp                |             |
| Status            | AccountApplicationStatus |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "result": [
    {
      "id": "string",
      "product_code": "string",
      "user_id": "string",
      "customer_profile_id": "string",
      "date_of_application": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "status": "AccountApplicationStatus"
    }
  ],
  "has_more": "bool"
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

## Update an account application

Updates an account application with the given status

```sh
curl -X PUT \
	/v1/account-applications/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"id": "string",
		"status": "AccountApplicationStatus"
	}'
```
{{snippet updateaccountapplication_status []}}

### HTTP Request

`PUT /v1/account-applications/{ID}`

### Query Parameters

| Name | Type   | Description                                                       |
|------|--------|-------------------------------------------------------------------|
| ID   | string | ID is the unique identifier of the account application to update. |

### Body Parameters

| Name   | Type                     | Description                                                       |
|--------|--------------------------|-------------------------------------------------------------------|
| ID     | string                   | ID is the unique identifier of the account application to update. |
| Status | AccountApplicationStatus | Status of the account application                                 |

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
| 204    | Account Application status updated successfully.                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  AccountApplicationStatus

Entity type defines the type of account application

| Value                    | Description                     |
|--------------------------|---------------------------------|
| UnknownApplicationStatus |                                 |
| Requested                | Account application is created  |
| Accepted                 | Account application is accepted |
