Account Public API v1.0.0
=========================

Provides create and read operations on the account public resource.

* Host `https://`

* Base Path ``

Get bank public account. {#method-get-getbankpublicaccount}
-----------------------------------------------------------

Return customer public account for specific bank.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/public \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getbankpublicaccount}

`GET https:///v1/banks/{BankID}/accounts/public`

### Query Parameters {#query-parameters-method-get-getbankpublicaccount}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getbankpublicaccount}

#### Response body {#response-body-method-get-getbankpublicaccount}

| Name     | Type             | Description |
|----------|------------------|-------------|
| accounts | \[]PublicAccount |             |

##### Objects {#objects-GetBankPublicAccountResponse}

###### PublicAccount

| Name            | Type             | Description |
|-----------------|------------------|-------------|
| id              | string           |             |
| label           | string           |             |
| bank_id         | string           |             |
| views_available | \[]ViewAvailable |             |

###### ViewAvailable

| Name       | Type   | Description |
|------------|--------|-------------|
| id         | string |             |
| short_name | string |             |
| is_public  | bool   |             |

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

#### Response codes {#response-codes-method-get-getbankpublicaccount}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get public account at all banks. {#method-get-getpublicaccountatallbanks}
-------------------------------------------------------------------------

Return public account at all banks.

```sh
curl -X GET \
	https:///v1/accounts/public \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getpublicaccountatallbanks}

`GET https:///v1/accounts/public`

### Responses {#responses-method-get-getpublicaccountatallbanks}

#### Response body {#response-body-method-get-getpublicaccountatallbanks}

| Name     | Type             | Description |
|----------|------------------|-------------|
| accounts | \[]PublicAccount |             |

##### Objects {#objects-GetPublicAccountAtAllBanksResponse}

###### PublicAccount

| Name            | Type             | Description |
|-----------------|------------------|-------------|
| id              | string           |             |
| label           | string           |             |
| bank_id         | string           |             |
| views_available | \[]ViewAvailable |             |

###### ViewAvailable

| Name       | Type   | Description |
|------------|--------|-------------|
| id         | string |             |
| short_name | string |             |
| is_public  | bool   |             |

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

#### Response codes {#response-codes-method-get-getpublicaccountatallbanks}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Get Public Account By ID. {#method-get-getpublicaccountbyid}
------------------------------------------------------------

Return customer public account for specific bank.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/public/accounts/{AccountID}/{ViewID}/account \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getpublicaccountbyid}

`GET https:///v1/banks/{BankID}/public/accounts/{AccountID}/{ViewID}/account`

### Query Parameters {#query-parameters-method-get-getpublicaccountbyid}

| Name       | Type  | Description |
|------------|-------|-------------|
| bank_id    | int32 |             |
| account_id | int32 |             |
| int32      |       |             |

### Responses {#responses-method-get-getpublicaccountbyid}

#### Response body {#response-body-method-get-getpublicaccountbyid}

| Name             | Type              | Description |
|------------------|-------------------|-------------|
| id               | string            |             |
| bank_id          | string            |             |
| label            | string            |             |
| number           | string            |             |
| owners           | Owner             |             |
| type             | string            |             |
| balance          | Amount            |             |
| account_routings | \[]AccountRouting |             |
| account_rules    | \[]AccountRule    |             |

##### Objects {#objects-GetPublicAccountByIDResponse}

###### Owner

| Name         | Type   | Description |
|--------------|--------|-------------|
| id           | string |             |
| provider     | string |             |
| display_name | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### AccountRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| scheme  | string |             |
| address | string |             |

###### AccountRule

| Name   | Type   | Description |
|--------|--------|-------------|
| scheme | string |             |
| value  | string |             |

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

#### Response codes {#response-codes-method-get-getpublicaccountbyid}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
