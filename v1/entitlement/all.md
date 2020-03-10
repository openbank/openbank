Entitlement API v1.0.0
======================

Provides CRUD operations on the entitlement part resource.

* Host `https://`

* Base Path ``

Add an entitlement request for current user {#method-post-addentitlementrequestforcurrentuser}
----------------------------------------------------------------------------------------------

Add an entitlement request for current user

```sh
curl -X POST \
	https:///v1/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"role_name": "string"
	}'
```

### HTTP Request

`POST https:///v1/entitlement-requests`

### Body Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| bank_id   | string |             |
| role_name | string |             |

### Responses

#### Response body

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| entitlement_request_id | string    |             |
| user                   | User      |             |
| role_name              | string    |             |
| bank_id                | string    |             |
| created                | Timestamp |             |

##### Objects

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| user_id      | string       |             |
| email        | string       |             |
| provider_id  | string       |             |
| provider     | string       |             |
| username     | string       |             |
| entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Add the entitlement request for user {#method-post-addentitlementrequestforuser}
--------------------------------------------------------------------------------

Add the entitlement entitlement request for user

```sh
curl -X POST \
	https:///v1/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"user_id": "string",
		"bank_id": "string",
		"role_name": "string"
	}'
```

### HTTP Request

`POST https:///v1/users/{UserID}/entitlements`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| user_id | string |             |

### Body Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| user_id   | string |             |
| bank_id   | string |             |
| role_name | string |             |

### Responses

#### Response body

| Name           | Type   | Description |
|----------------|--------|-------------|
| user_id        | string |             |
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Delete an entitlement {#method-delete-deleteentitlement}
--------------------------------------------------------

Permanently delete an entitlement.

```sh
curl -X DELETE \
	https:///v1/users/{UserID}/entitlements/{EntitlementID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/users/{UserID}/entitlements/{EntitlementID}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| user_id        | string |             |
| entitlement_id | string |             |

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

Delete a entitlement request {#method-delete-deleteentitlementrequest}
----------------------------------------------------------------------

Permanently delete an entitlement request.

```sh
curl -X DELETE \
	https:///v1/entitlement-requests/{EntitlementRequestID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/entitlement-requests/{EntitlementRequestID}`

### Query Parameters

| Name                   | Type   | Description |
|------------------------|--------|-------------|
| entitlement_request_id | string |             |

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

List all entitlements {#method-get-getallentitlementrequests}
-------------------------------------------------------------

Returns a list containing up to 20 entitlements. `after_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/entitlement-requests`

### Responses

#### Response body

| Name                 | Type                  | Description |
|----------------------|-----------------------|-------------|
| entitlement_requests | \[]EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| entitlement_request_id | string    |             |
| user                   | User      |             |
| role_name              | string    |             |
| bank_id                | string    |             |
| created                | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| user_id      | string       |             |
| email        | string       |             |
| provider_id  | string       |             |
| provider     | string       |             |
| username     | string       |             |
| entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

List all entitlements {#method-get-getallentitlements}
------------------------------------------------------

Returns a list containing up to 20 entitlements. `after_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/entitlements`

### Responses

#### Response body

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

##### Objects

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Retrieve entitlement information {#method-get-getentitlementforcurrentuser}
---------------------------------------------------------------------------

Retrieve information about the entitlement specified for current user

```sh
curl -X GET \
	https:///v1/users/current/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/users/current/entitlements`

### Responses

#### Response body

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

##### Objects

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Retrieve entitlement information {#method-get-getentitlementforuser}
--------------------------------------------------------------------

Retrieve information about the entitlement specified by the User ID

```sh
curl -X GET \
	https:///v1/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/users/{UserID}/entitlements`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| user_id | string |             |

### Responses

#### Response body

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

##### Objects

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Retrieve entitlement information {#method-get-getentitlementforuseratbank}
--------------------------------------------------------------------------

Retrieve information about the entitlement specified by the User ID at bank

```sh
curl -X GET \
	https:///v1/banks/{BankID}/users/{UserID}/entitlements \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/users/{UserID}/entitlements`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |
| user_id | string |             |

### Responses

#### Response body

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

##### Objects

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Retrieve entitlement request information {#method-get-getentitlementrequestforcurrentuser}
------------------------------------------------------------------------------------------

Retrieve information about the entitlement request specified for current user

```sh
curl -X GET \
	https:///v1/users/current/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/users/current/entitlement-requests`

### Responses

#### Response body

| Name                 | Type                  | Description |
|----------------------|-----------------------|-------------|
| entitlement_requests | \[]EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| entitlement_request_id | string    |             |
| user                   | User      |             |
| role_name              | string    |             |
| bank_id                | string    |             |
| created                | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| user_id      | string       |             |
| email        | string       |             |
| provider_id  | string       |             |
| provider     | string       |             |
| username     | string       |             |
| entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

Retrieve entitlement request information {#method-get-getentitlementrequestforuser}
-----------------------------------------------------------------------------------

Retrieve information about the entitlement request specified by the User ID

```sh
curl -X GET \
	https:///v1/users/{UserID}/entitlement-requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/users/{UserID}/entitlement-requests`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| user_id | string |             |

### Responses

#### Response body

| Name                 | Type                  | Description |
|----------------------|-----------------------|-------------|
| entitlement_requests | \[]EntitlementRequest |             |

##### Objects

###### EntitlementRequest

| Name                   | Type      | Description |
|------------------------|-----------|-------------|
| entitlement_request_id | string    |             |
| user                   | User      |             |
| role_name              | string    |             |
| bank_id                | string    |             |
| created                | Timestamp |             |

###### User

| Name         | Type         | Description |
|--------------|--------------|-------------|
| user_id      | string       |             |
| email        | string       |             |
| provider_id  | string       |             |
| provider     | string       |             |
| username     | string       |             |
| entitlements | Entitlements |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Entitlements

| Name | Type    | Description |
|------|---------|-------------|
| list | \[]List |             |

###### List

| Name           | Type   | Description |
|----------------|--------|-------------|
| entitlement_id | string |             |
| role_name      | string |             |
| bank_id        | string |             |

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

List all entitlements {#method-get-getroles}
--------------------------------------------

Returns a list containing up to 20 entitlements. `after_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/roles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/roles`

### Responses

#### Response body

| Name  | Type     | Description |
|-------|----------|-------------|
| roles | \[]Roles |             |

##### Objects

###### Roles

| Name             | Type   | Description |
|------------------|--------|-------------|
| role             | string |             |
| requires_bank_id | bool   |             |

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
