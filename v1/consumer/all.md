Consumer API v1.0.0
===================

Provides create and read operations on the consumer resource.

* Host `https://`

* Base Path ``

Set consumer enable or disable. {#method-put-enableordisableconsumer}
---------------------------------------------------------------------

Set enable or disable consumer

```sh
curl -X PUT \
	https:///v1/management/consumers/{ConsumerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"consumer_id": "string",
		"enable": "bool"
	}'
```

### HTTP Request {#http-request-method-put-enableordisableconsumer}

`PUT https:///v1/management/consumers/{ConsumerID}`

### Query Parameters {#query-parameters-method-put-enableordisableconsumer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |

### Body Parameters {#body-parameters-method-put-enableordisableconsumer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |
| enable      | bool   |             |

### Responses {#responses-method-put-enableordisableconsumer}

#### Response body {#response-body-method-put-enableordisableconsumer}

| Name   | Type | Description |
|--------|------|-------------|
| enable | bool |             |

Example:

```json
{
  "enable": "bool"
}
```

#### Response codes {#response-codes-method-put-enableordisableconsumer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get consumer {#method-get-getconsumer}
--------------------------------------

Returns consumer data by it's ID

```sh
curl -X GET \
	https:///v1/management/consumers/{ConsumerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getconsumer}

`GET https:///v1/management/consumers/{ConsumerID}`

### Query Parameters {#query-parameters-method-get-getconsumer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |

### Responses {#responses-method-get-getconsumer}

#### Response body {#response-body-method-get-getconsumer}

| Name     | Type     | Description |
|----------|----------|-------------|
| consumer | Consumer |             |

##### Objects {#objects-GetConsumerResponse}

###### Consumer

| Name               | Type         | Description |
|--------------------|--------------|-------------|
| id                 | int32        |             |
| app_name           | string       |             |
| app_type           | string       |             |
| description        | string       |             |
| developer_email    | string       |             |
| redirect_url       | string       |             |
| created_by_user_id | string       |             |
| created_by_user    | CreateByUser |             |
| enabled            | bool         |             |
| created            | Timestamp    |             |

###### CreateByUser

| Name        | Type   | Description |
|-------------|--------|-------------|
| user_id     | string |             |
| email       | string |             |
| provider_id | string |             |
| provider    | string |             |
| username    | string |             |

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

#### Response codes {#response-codes-method-get-getconsumer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get consumer call lists {#method-get-getconsumercalllimits}
-----------------------------------------------------------

Returns a list up to 20 consumer call limits data.

```sh
curl -X GET \
	https:///v1/management/consumers/{ConsumerID}/consumer/call-limits \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getconsumercalllimits}

`GET https:///v1/management/consumers/{ConsumerID}/consumer/call-limits`

### Query Parameters {#query-parameters-method-get-getconsumercalllimits}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |

### Responses {#responses-method-get-getconsumercalllimits}

#### Response body {#response-body-method-get-getconsumercalllimits}

| Name                  | Type         | Description |
|-----------------------|--------------|-------------|
| per_second_call_limit | string       |             |
| per_minute_call_limit | string       |             |
| per_hour_call_limit   | string       |             |
| per_day_call_limit    | string       |             |
| per_week_call_limit   | string       |             |
| per_month_call_limit  | string       |             |
| current_state         | CurrentState |             |

##### Objects {#objects-GetConsumerCallLimitsResponse}

###### CurrentState

| Name       | Type  | Description |
|------------|-------|-------------|
| per_second | State |             |
| per_hour   | State |             |
| per_day    | State |             |
| per_week   | State |             |
| per_month  | State |             |
| per_year   | State |             |

###### State

| Name             | Type  | Description |
|------------------|-------|-------------|
| calls_made       | int32 |             |
| reset_in_seconds | int32 |             |

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

#### Response codes {#response-codes-method-get-getconsumercalllimits}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get consumers {#method-get-getconsumers}
----------------------------------------

Returns all consumers data up to 20 consumers data.

```sh
curl -X GET \
	https:///v1/management/consumers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getconsumers}

`GET https:///v1/management/consumers`

### Responses {#responses-method-get-getconsumers}

#### Response body {#response-body-method-get-getconsumers}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| consumers | \[]Consumer |             |

##### Objects {#objects-GetConsumersResponse}

###### Consumer

| Name               | Type         | Description |
|--------------------|--------------|-------------|
| id                 | int32        |             |
| app_name           | string       |             |
| app_type           | string       |             |
| description        | string       |             |
| developer_email    | string       |             |
| redirect_url       | string       |             |
| created_by_user_id | string       |             |
| created_by_user    | CreateByUser |             |
| enabled            | bool         |             |
| created            | Timestamp    |             |

###### CreateByUser

| Name        | Type   | Description |
|-------------|--------|-------------|
| user_id     | string |             |
| email       | string |             |
| provider_id | string |             |
| provider    | string |             |
| username    | string |             |

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

#### Response codes {#response-codes-method-get-getconsumers}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get consumers for logged in user {#method-get-getconsumersforloggedinuser}
--------------------------------------------------------------------------

Returns all consumers for logged in user data up to 20 consumers data.

```sh
curl -X GET \
	https:///v1/management/users/current/consumers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getconsumersforloggedinuser}

`GET https:///v1/management/users/current/consumers`

### Responses {#responses-method-get-getconsumersforloggedinuser}

#### Response body {#response-body-method-get-getconsumersforloggedinuser}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| consumers | \[]Consumer |             |

##### Objects {#objects-GetConsumersForLoggedInUserResponse}

###### Consumer

| Name               | Type         | Description |
|--------------------|--------------|-------------|
| id                 | int32        |             |
| app_name           | string       |             |
| app_type           | string       |             |
| description        | string       |             |
| developer_email    | string       |             |
| redirect_url       | string       |             |
| created_by_user_id | string       |             |
| created_by_user    | CreateByUser |             |
| enabled            | bool         |             |
| created            | Timestamp    |             |

###### CreateByUser

| Name        | Type   | Description |
|-------------|--------|-------------|
| user_id     | string |             |
| email       | string |             |
| provider_id | string |             |
| provider    | string |             |
| username    | string |             |

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

#### Response codes {#response-codes-method-get-getconsumersforloggedinuser}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Set consumer call limits. {#method-put-setconsumercallslimit}
-------------------------------------------------------------

Set consumer call limits.

```sh
curl -X PUT \
	https:///v1/management/consumers/{ConsumerID}/consumer/call-limits \
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

### HTTP Request {#http-request-method-put-setconsumercallslimit}

`PUT https:///v1/management/consumers/{ConsumerID}/consumer/call-limits`

### Query Parameters {#query-parameters-method-put-setconsumercallslimit}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |

### Body Parameters {#body-parameters-method-put-setconsumercallslimit}

| Name                  | Type   | Description |
|-----------------------|--------|-------------|
| consumer_id           | string |             |
| per_second_call_limit | string |             |
| per_minute_call_limit | string |             |
| per_hour_call_limit   | string |             |
| per_day_call_limit    | string |             |
| per_week_call_limit   | string |             |
| per_month_call_limit  | string |             |

### Responses {#responses-method-put-setconsumercallslimit}

#### Response body {#response-body-method-put-setconsumercallslimit}

| Name                  | Type   | Description |
|-----------------------|--------|-------------|
| per_second_call_limit | string |             |
| per_minute_call_limit | string |             |
| per_hour_call_limit   | string |             |
| per_day_call_limit    | string |             |
| per_week_call_limit   | string |             |
| per_month_call_limit  | string |             |

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

#### Response codes {#response-codes-method-put-setconsumercallslimit}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Set consumer redirect url. {#method-put-updateconsumerredirecturl}
------------------------------------------------------------------

Set consumer redirect url.

```sh
curl -X PUT \
	https:///v1/management/consumers/{ConsumerID}/consumer/redirect_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"consumer_id": "string",
		"redirect_url": "string"
	}'
```

### HTTP Request {#http-request-method-put-updateconsumerredirecturl}

`PUT https:///v1/management/consumers/{ConsumerID}/consumer/redirect_url`

### Query Parameters {#query-parameters-method-put-updateconsumerredirecturl}

| Name        | Type   | Description |
|-------------|--------|-------------|
| consumer_id | string |             |

### Body Parameters {#body-parameters-method-put-updateconsumerredirecturl}

| Name         | Type   | Description |
|--------------|--------|-------------|
| consumer_id  | string |             |
| redirect_url | string |             |

### Responses {#responses-method-put-updateconsumerredirecturl}

#### Response body {#response-body-method-put-updateconsumerredirecturl}

| Name     | Type     | Description |
|----------|----------|-------------|
| consumer | Consumer |             |

##### Objects {#objects-UpdateConsumerRedirectURLResponse}

###### Consumer

| Name               | Type         | Description |
|--------------------|--------------|-------------|
| id                 | int32        |             |
| app_name           | string       |             |
| app_type           | string       |             |
| description        | string       |             |
| developer_email    | string       |             |
| redirect_url       | string       |             |
| created_by_user_id | string       |             |
| created_by_user    | CreateByUser |             |
| enabled            | bool         |             |
| created            | Timestamp    |             |

###### CreateByUser

| Name        | Type   | Description |
|-------------|--------|-------------|
| user_id     | string |             |
| email       | string |             |
| provider_id | string |             |
| provider    | string |             |
| username    | string |             |

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

#### Response codes {#response-codes-method-put-updateconsumerredirecturl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
