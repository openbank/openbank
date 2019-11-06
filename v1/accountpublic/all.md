# Account Public API v1.0.0

Provides create and read operations on the account public resource.

*
Host ``
EOL

*
Base Path ``

## Get bank public account.

Return customer public account for specific bank.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/public \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getbankpublic_account []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/public`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name     | Type            | Description |
|----------|-----------------|-------------|
| Accounts | []PublicAccount |             |

##### Objects

###### PublicAccount

| Name           | Type            | Description |
|----------------|-----------------|-------------|
| ID             | string          |             |
| Label          | string          |             |
| BankID         | string          |             |
| ViewsAvailable | []ViewAvailable |             |

###### ViewAvailable

| Name      | Type   | Description |
|-----------|--------|-------------|
| ID        | string |             |
| ShortName | string |             |
| IsPublic  | bool   |             |

Example:

```json
{
  "accounts": [
    {
      "id": "string",
      "label": "string",
      "bank_id": "string",
      "views_available": [
        {
          "id": "string",
          "short_name": "string",
          "is_public": "bool"
        }
      ]
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

## Get public account at all banks.

Return public account at all banks.

```sh
curl -X GET \
	/v1/accounts/public \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getpublicaccountatall_banks []}}

### HTTP Request

`GET /v1/accounts/public`

### Responses

#### Response body

| Name     | Type            | Description |
|----------|-----------------|-------------|
| Accounts | []PublicAccount |             |

##### Objects

###### PublicAccount

| Name           | Type            | Description |
|----------------|-----------------|-------------|
| ID             | string          |             |
| Label          | string          |             |
| BankID         | string          |             |
| ViewsAvailable | []ViewAvailable |             |

###### ViewAvailable

| Name      | Type   | Description |
|-----------|--------|-------------|
| ID        | string |             |
| ShortName | string |             |
| IsPublic  | bool   |             |

Example:

```json
{
  "accounts": [
    {
      "id": "string",
      "label": "string",
      "bank_id": "string",
      "views_available": [
        {
          "id": "string",
          "short_name": "string",
          "is_public": "bool"
        }
      ]
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

## Get Public Account By ID.

Return customer public account for specific bank.

```sh
curl -X GET \
	/v1/banks/{BankID}/public/accounts/{AccountID}/{ViewID}/account \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getpublicaccountbyid []}}

### HTTP Request

`GET /v1/banks/{BankID}/public/accounts/{AccountID}/{ViewID}/account`

### Query Parameters

| Name      | Type  | Description |
|-----------|-------|-------------|
| BankID    | int32 |             |
| AccountID | int32 |             |
| ViewID    | int32 |             |

### Responses

#### Response body

| Name            | Type             | Description |
|-----------------|------------------|-------------|
| ID              | string           |             |
| BankID          | string           |             |
| Label           | string           |             |
| Number          | string           |             |
| Owners          | Owner            |             |
| Type            | string           |             |
| Balance         | Amount           |             |
| AccountRoutings | []AccountRouting |             |
| AccountRules    | []AccountRule    |             |

##### Objects

###### Owner

| Name        | Type   | Description |
|-------------|--------|-------------|
| ID          | string |             |
| Provider    | string |             |
| DisplayName | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### AccountRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| Scheme  | string |             |
| Address | string |             |

###### AccountRule

| Name   | Type   | Description |
|--------|--------|-------------|
| Scheme | string |             |
| Value  | string |             |

Example:

```json
{
  "id": "string",
  "bank_id": "string",
  "label": "string",
  "number": "string",
  "owners": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  },
  "type": "string",
  "balance": {
    "cur": "string",
    "num": "string"
  },
  "account_routings": [
    {
      "scheme": "string",
      "address": "string"
    }
  ],
  "account_rules": [
    {
      "scheme": "string",
      "value": "string"
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
