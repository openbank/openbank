# Entitlement API v1.0.0

Provides CRUD operations on the entitlement part resource.

* Host ``

* Base Path ``

## Add an entitlement request for current user {#method-post-addentitlementrequestforcurrentuser}

Add an entitlement request for current user

```sh
curl -X POST \
	/v1/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"role_name": "string"
	}'
```
{{snippet addentitlementrequestforcurrent_user []}}

### HTTP Request

`POST /v1/entitlement-requests`

### Body Parameters

| Name     | Type   | Description |
|----------|--------|-------------|
| BankID   | string |             |
| RoleName | string |             |

### Responses

#### Response body

| Name                 | Type      | Description |
|----------------------|-----------|-------------|
| EntitlementRequestID | string    |             |
| User                 | User      |             |
| RoleName             | string    |             |
| BankID               | string    |             |
| Created              | Timestamp |             |

##### Objects

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| UserID       | string       |             |
| Email        | string       |             |
| ProviderID   | string       |             |
| Provider     | string       |             |
| Username     | string       |             |
| Entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "entitlement_request_id": "string",
  "user": {
    "user_id": "string",
    "email": "string",
    "provider_id": "string",
    "provider": "string",
    "username": "string",
    "entitlements": {
      "list": [
        {
          "entitlement_id": "string",
          "role_name": "string",
          "bank_id": "string"
        }
      ]
    }
  },
  "role_name": "string",
  "bank_id": "string",
  "created": {
    "seconds": "int64",
    "nanos": "int32"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Add an entitlement request for current user                                            |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add the entitlement request for user {#method-post-addentitlementrequestforuser}

Add the entitlement entitlement request for user

```sh
curl -X POST \
	/v1/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"user_id": "string",
		"bank_id": "string",
		"role_name": "string"
	}'
```
{{snippet addentitlementrequestforuser []}}

### HTTP Request

`POST /v1/users/{UserID}/entitlements`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| UserID | string |             |

### Body Parameters

| Name     | Type   | Description |
|----------|--------|-------------|
| UserID   | string |             |
| BankID   | string |             |
| RoleName | string |             |

### Responses

#### Response body

| Name          | Type   | Description |
|---------------|--------|-------------|
| UserID        | string |             |
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "user_id": "string",
  "entitlement_id": "string",
  "role_name": "string",
  "bank_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Add the entitlement request for user                                                   |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete an entitlement {#method-delete-deleteentitlement}

Permanently delete an entitlement.

```sh
curl -X DELETE \
	/v1/users/{UserID}/entitlements/{EntitlementID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/users/{UserID}/entitlements/{EntitlementID}`

### Query Parameters

| Name          | Type   | Description |
|---------------|--------|-------------|
| UserID        | string |             |
| EntitlementID | string |             |

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
| 204    | Entitlement successfully deleted.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a entitlement request {#method-delete-deleteentitlementrequest}

Permanently delete an entitlement request.

```sh
curl -X DELETE \
	/v1/entitlement-requests/{EntitlementRequestID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteentitlementrequest []}}

### HTTP Request

`DELETE /v1/entitlement-requests/{EntitlementRequestID}`

### Query Parameters

| Name                 | Type   | Description |
|----------------------|--------|-------------|
| EntitlementRequestID | string |             |

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
| 204    | Entitlement request successfully deleted.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all entitlements {#method-get-getallentitlementrequests}

Returns a list containing up to 20 entitlements. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getallentitlement_requests []}}

### HTTP Request

`GET /v1/entitlement-requests`

### Responses

#### Response body

| Name                | Type                 | Description |
|---------------------|----------------------|-------------|
| EntitlementRequests | []EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                 | Type      | Description |
|----------------------|-----------|-------------|
| EntitlementRequestID | string    |             |
| User                 | User      |             |
| RoleName             | string    |             |
| BankID               | string    |             |
| Created              | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| UserID       | string       |             |
| Email        | string       |             |
| ProviderID   | string       |             |
| Provider     | string       |             |
| Username     | string       |             |
| Entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "entitlement_requests": [
    {
      "entitlement_request_id": "string",
      "user": {
        "user_id": "string",
        "email": "string",
        "provider_id": "string",
        "provider": "string",
        "username": "string",
        "entitlements": {
          "list": [
            {
              "entitlement_id": "string",
              "role_name": "string",
              "bank_id": "string"
            }
          ]
        }
      },
      "role_name": "string",
      "bank_id": "string",
      "created": {
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
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all entitlements {#method-get-getallentitlements}

Returns a list containing up to 20 entitlements. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getallentitlements []}}

### HTTP Request

`GET /v1/entitlements`

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

##### Objects

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "list": [
    {
      "entitlement_id": "string",
      "role_name": "string",
      "bank_id": "string"
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

## Retrieve entitlement information {#method-get-getentitlementforcurrentuser}

Retrieve information about the entitlement specified for current user

```sh
curl -X GET \
	/v1/users/current/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getentitlementforcurrentuser []}}

### HTTP Request

`GET /v1/users/current/entitlements`

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

##### Objects

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "list": [
    {
      "entitlement_id": "string",
      "role_name": "string",
      "bank_id": "string"
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

## Retrieve entitlement information {#method-get-getentitlementforuser}

Retrieve information about the entitlement specified by the User ID

```sh
curl -X GET \
	/v1/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getentitlementfor_user []}}

### HTTP Request

`GET /v1/users/{UserID}/entitlements`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| UserID | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

##### Objects

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "list": [
    {
      "entitlement_id": "string",
      "role_name": "string",
      "bank_id": "string"
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

## Retrieve entitlement information {#method-get-getentitlementforuseratbank}

Retrieve information about the entitlement specified by the User ID at bank

```sh
curl -X GET \
	/v1/banks/{BankID}/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getentitlementforuserat_bank []}}

### HTTP Request

`GET /v1/banks/{BankID}/users/{UserID}/entitlements`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |
| UserID | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

##### Objects

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "list": [
    {
      "entitlement_id": "string",
      "role_name": "string",
      "bank_id": "string"
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

## Retrieve entitlement request information {#method-get-getentitlementrequestforcurrentuser}

Retrieve information about the entitlement request specified for current user

```sh
curl -X GET \
	/v1/users/current/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getentitlementrequestforcurrent_user []}}

### HTTP Request

`GET /v1/users/current/entitlement-requests`

### Responses

#### Response body

| Name                | Type                 | Description |
|---------------------|----------------------|-------------|
| EntitlementRequests | []EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                 | Type      | Description |
|----------------------|-----------|-------------|
| EntitlementRequestID | string    |             |
| User                 | User      |             |
| RoleName             | string    |             |
| BankID               | string    |             |
| Created              | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| UserID       | string       |             |
| Email        | string       |             |
| ProviderID   | string       |             |
| Provider     | string       |             |
| Username     | string       |             |
| Entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "entitlement_requests": [
    {
      "entitlement_request_id": "string",
      "user": {
        "user_id": "string",
        "email": "string",
        "provider_id": "string",
        "provider": "string",
        "username": "string",
        "entitlements": {
          "list": [
            {
              "entitlement_id": "string",
              "role_name": "string",
              "bank_id": "string"
            }
          ]
        }
      },
      "role_name": "string",
      "bank_id": "string",
      "created": {
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
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve entitlement request information {#method-get-getentitlementrequestforuser}

Retrieve information about the entitlement request specified by the User ID

```sh
curl -X GET \
	/v1/users/{UserID}/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getentitlementrequestforuser []}}

### HTTP Request

`GET /v1/users/{UserID}/entitlement-requests`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| UserID | string |             |

### Responses

#### Response body

| Name                | Type                 | Description |
|---------------------|----------------------|-------------|
| EntitlementRequests | []EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                 | Type      | Description |
|----------------------|-----------|-------------|
| EntitlementRequestID | string    |             |
| User                 | User      |             |
| RoleName             | string    |             |
| BankID               | string    |             |
| Created              | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| UserID       | string       |             |
| Email        | string       |             |
| ProviderID   | string       |             |
| Provider     | string       |             |
| Username     | string       |             |
| Entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type   | Description |
|------|--------|-------------|
| List | []List |             |

###### List

| Name          | Type   | Description |
|---------------|--------|-------------|
| EntitlementID | string |             |
| RoleName      | string |             |
| BankID        | string |             |

Example:

```json
{
  "entitlement_requests": [
    {
      "entitlement_request_id": "string",
      "user": {
        "user_id": "string",
        "email": "string",
        "provider_id": "string",
        "provider": "string",
        "username": "string",
        "entitlements": {
          "list": [
            {
              "entitlement_id": "string",
              "role_name": "string",
              "bank_id": "string"
            }
          ]
        }
      },
      "role_name": "string",
      "bank_id": "string",
      "created": {
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
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all entitlements {#method-get-getroles}

Returns a list containing up to 20 entitlements. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/roles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/roles`

### Responses

#### Response body

| Name  | Type    | Description |
|-------|---------|-------------|
| Roles | []Roles |             |

##### Objects

###### Roles

| Name           | Type   | Description |
|----------------|--------|-------------|
| Role           | string |             |
| RequiresBankID | bool   |             |

Example:

```json
{
  "roles": [
    {
      "role": "string",
      "requires_bank_id": "bool"
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
