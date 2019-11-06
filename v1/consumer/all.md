# Consumer API v1.0.0

Provides create and read operations on the consumer resource.

*
Host ``
EOL

*
Base Path ``

## Set consumer enable or disable.

Set enable or disable consumer

```sh
curl -X PUT \
	/v1/management/consumers/{ConsumerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"consumer_id": "string",
		"enable": "bool"
	}'
```
{{snippet enableordisable_consumer []}}

### HTTP Request

`PUT /v1/management/consumers/{ConsumerID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |
| Enable     | bool   |             |

### Responses

#### Response body

| Name   | Type | Description |
|--------|------|-------------|
| Enable | bool |             |

Example:

```json
{
  "enable": "bool"
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

## Get consumer

Returns consumer data by it's ID

```sh
curl -X GET \
	/v1/management/consumers/{ConsumerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/management/consumers/{ConsumerID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Consumer | Consumer |             |

##### Objects

###### Consumer

| Name            | Type         | Description |
|-----------------|--------------|-------------|
| ID              | int32        |             |
| AppName         | string       |             |
| AppType         | string       |             |
| Description     | string       |             |
| DeveloperEmail  | string       |             |
| RedirectURL     | string       |             |
| CreatedByUserID | string       |             |
| CreatedByUser   | CreateByUser |             |
| Enabled         | bool         |             |
| Created         | Timestamp    |             |

###### CreateByUser

| Name       | Type   | Description |
|------------|--------|-------------|
| UserID     | string |             |
| Email      | string |             |
| ProviderID | string |             |
| Provider   | string |             |
| Username   | string |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "consumer": {
    "id": "int32",
    "app_name": "string",
    "app_type": "string",
    "description": "string",
    "developer_email": "string",
    "redirect_url": "string",
    "created_by_user_id": "string",
    "created_by_user": {
      "user_id": "string",
      "email": "string",
      "provider_id": "string",
      "provider": "string",
      "username": "string"
    },
    "enabled": "bool",
    "created": {
      "seconds": "int64",
      "nanos": "int32"
    }
  }
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

## Get consumer call lists

Returns a list up to 20 consumer call limits data.

```sh
curl -X GET \
	/v1/management/consumers/{ConsumerID}/consumer/call-limits \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getconsumercall_limits []}}

### HTTP Request

`GET /v1/management/consumers/{ConsumerID}/consumer/call-limits`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |

### Responses

#### Response body

| Name               | Type         | Description |
|--------------------|--------------|-------------|
| PerSecondCallLimit | string       |             |
| PerMinuteCallLimit | string       |             |
| PerHourCallLimit   | string       |             |
| PerDayCallLimit    | string       |             |
| PerWeekCallLimit   | string       |             |
| PerMonthCallLimit  | string       |             |
| CurrentState       | CurrentState |             |

##### Objects

###### CurrentState

| Name      | Type  | Description |
|-----------|-------|-------------|
| PerSecond | State |             |
| PerHour   | State |             |
| PerDay    | State |             |
| PerWeek   | State |             |
| PerMonth  | State |             |
| PerYear   | State |             |

###### State

| Name           | Type  | Description |
|----------------|-------|-------------|
| CallsMade      | int32 |             |
| ResetInSeconds | int32 |             |

Example:

```json
{
  "per_second_call_limit": "string",
  "per_minute_call_limit": "string",
  "per_hour_call_limit": "string",
  "per_day_call_limit": "string",
  "per_week_call_limit": "string",
  "per_month_call_limit": "string",
  "current_state": {
    "per_second": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    },
    "per_hour": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    },
    "per_day": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    },
    "per_week": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    },
    "per_month": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    },
    "per_year": {
      "calls_made": "int32",
      "reset_in_seconds": "int32"
    }
  }
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

## Get consumers

Returns all consumers data up to 20 consumers data.

```sh
curl -X GET \
	/v1/management/consumers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/management/consumers`

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Consumers | []Consumer |             |

##### Objects

###### Consumer

| Name            | Type         | Description |
|-----------------|--------------|-------------|
| ID              | int32        |             |
| AppName         | string       |             |
| AppType         | string       |             |
| Description     | string       |             |
| DeveloperEmail  | string       |             |
| RedirectURL     | string       |             |
| CreatedByUserID | string       |             |
| CreatedByUser   | CreateByUser |             |
| Enabled         | bool         |             |
| Created         | Timestamp    |             |

###### CreateByUser

| Name       | Type   | Description |
|------------|--------|-------------|
| UserID     | string |             |
| Email      | string |             |
| ProviderID | string |             |
| Provider   | string |             |
| Username   | string |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "consumers": [
    {
      "id": "int32",
      "app_name": "string",
      "app_type": "string",
      "description": "string",
      "developer_email": "string",
      "redirect_url": "string",
      "created_by_user_id": "string",
      "created_by_user": {
        "user_id": "string",
        "email": "string",
        "provider_id": "string",
        "provider": "string",
        "username": "string"
      },
      "enabled": "bool",
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

## Get consumers for logged in user

Returns all consumers for logged in user data up to 20 consumers data.

```sh
curl -X GET \
	/v1/management/users/current/consumers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getconsumersforloggedin_user []}}

### HTTP Request

`GET /v1/management/users/current/consumers`

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Consumers | []Consumer |             |

##### Objects

###### Consumer

| Name            | Type         | Description |
|-----------------|--------------|-------------|
| ID              | int32        |             |
| AppName         | string       |             |
| AppType         | string       |             |
| Description     | string       |             |
| DeveloperEmail  | string       |             |
| RedirectURL     | string       |             |
| CreatedByUserID | string       |             |
| CreatedByUser   | CreateByUser |             |
| Enabled         | bool         |             |
| Created         | Timestamp    |             |

###### CreateByUser

| Name       | Type   | Description |
|------------|--------|-------------|
| UserID     | string |             |
| Email      | string |             |
| ProviderID | string |             |
| Provider   | string |             |
| Username   | string |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "consumers": [
    {
      "id": "int32",
      "app_name": "string",
      "app_type": "string",
      "description": "string",
      "developer_email": "string",
      "redirect_url": "string",
      "created_by_user_id": "string",
      "created_by_user": {
        "user_id": "string",
        "email": "string",
        "provider_id": "string",
        "provider": "string",
        "username": "string"
      },
      "enabled": "bool",
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

## Set consumer call limits.

Set consumer call limits.

```sh
curl -X PUT \
	/v1/management/consumers/{ConsumerID}/consumer/call-limits \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"consumer_id": "string",
		"per_second_call_limit": "string",
		"per_minute_call_limit": "string",
		"per_hour_call_limit": "string",
		"per_day_call_limit": "string",
		"per_week_call_limit": "string",
		"per_month_call_limit": "string"
	}'
```
{{snippet setconsumercalls_limit []}}

### HTTP Request

`PUT /v1/management/consumers/{ConsumerID}/consumer/call-limits`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |

### Body Parameters

| Name               | Type   | Description |
|--------------------|--------|-------------|
| ConsumerID         | string |             |
| PerSecondCallLimit | string |             |
| PerMinuteCallLimit | string |             |
| PerHourCallLimit   | string |             |
| PerDayCallLimit    | string |             |
| PerWeekCallLimit   | string |             |
| PerMonthCallLimit  | string |             |

### Responses

#### Response body

| Name               | Type   | Description |
|--------------------|--------|-------------|
| PerSecondCallLimit | string |             |
| PerMinuteCallLimit | string |             |
| PerHourCallLimit   | string |             |
| PerDayCallLimit    | string |             |
| PerWeekCallLimit   | string |             |
| PerMonthCallLimit  | string |             |

Example:

```json
{
  "per_second_call_limit": "string",
  "per_minute_call_limit": "string",
  "per_hour_call_limit": "string",
  "per_day_call_limit": "string",
  "per_week_call_limit": "string",
  "per_month_call_limit": "string"
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

## Set consumer redirect url.

Set consumer redirect url.

```sh
curl -X PUT \
	/v1/management/consumers/{ConsumerID}/consumer/redirect_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"consumer_id": "string",
		"redirect_url": "string"
	}'
```
{{snippet updateconsumerredirect_url []}}

### HTTP Request

`PUT /v1/management/consumers/{ConsumerID}/consumer/redirect_url`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| ConsumerID | string |             |

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| ConsumerID  | string |             |
| RedirectURL | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Consumer | Consumer |             |

##### Objects

###### Consumer

| Name            | Type         | Description |
|-----------------|--------------|-------------|
| ID              | int32        |             |
| AppName         | string       |             |
| AppType         | string       |             |
| Description     | string       |             |
| DeveloperEmail  | string       |             |
| RedirectURL     | string       |             |
| CreatedByUserID | string       |             |
| CreatedByUser   | CreateByUser |             |
| Enabled         | bool         |             |
| Created         | Timestamp    |             |

###### CreateByUser

| Name       | Type   | Description |
|------------|--------|-------------|
| UserID     | string |             |
| Email      | string |             |
| ProviderID | string |             |
| Provider   | string |             |
| Username   | string |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "consumer": {
    "id": "int32",
    "app_name": "string",
    "app_type": "string",
    "description": "string",
    "developer_email": "string",
    "redirect_url": "string",
    "created_by_user_id": "string",
    "created_by_user": {
      "user_id": "string",
      "email": "string",
      "provider_id": "string",
      "provider": "string",
      "username": "string"
    },
    "enabled": "bool",
    "created": {
      "seconds": "int64",
      "nanos": "int32"
    }
  }
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
