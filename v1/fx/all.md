# FX API v1.0.0

Provides create and read operations on the foreign exchange resource.

*
Host ``
EOL

*
Base Path ``

## Create a foreign exchange quote

Creates a new foreign exchange quote and returns it.

```sh
curl -X POST \
	/v1/fx \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"fx": {
			"bank_id": "string",
			"from_currency_code": "string",
			"to_currency_code": "string",
			"rate": "string",
			"inverse_rate": "string",
			"effective_date": "string"
		}
	}'
```
{{snippet createfx []}}

### HTTP Request

`POST /v1/fx`

### Body Parameters

| Name | Type | Description                                       |
|------|------|---------------------------------------------------|
| FX   | FX   | FX is the foreign exchange information to create. |

##### Objects

###### FX

| Name             | Type   | Description                                                              |
|------------------|--------|--------------------------------------------------------------------------|
| BankID           | string | BankID is an identifier for the bank on this fx transaction.             |
| FromCurrencyCode | string | FromCurrencyCode is the currency to transfer from.                       |
| ToCurrencyCode   | string | ToCurrencyCode is the currency that we are transferring to.              |
| Rate             | string | Rate is the exchange rate of the foreign exchange.                       |
| InverseRate      | string | InverseRate is the inverse of the exchange rate of the foreign exchange. |
| EffectiveDate    | string | EffectiveDate is the effective date of the foreign exchange quote.       |

### Responses

#### Response body

| Name             | Type   | Description                                                              |
|------------------|--------|--------------------------------------------------------------------------|
| BankID           | string | BankID is an identifier for the bank on this fx transaction.             |
| FromCurrencyCode | string | FromCurrencyCode is the currency to transfer from.                       |
| ToCurrencyCode   | string | ToCurrencyCode is the currency that we are transferring to.              |
| Rate             | string | Rate is the exchange rate of the foreign exchange.                       |
| InverseRate      | string | InverseRate is the inverse of the exchange rate of the foreign exchange. |
| EffectiveDate    | string | EffectiveDate is the effective date of the foreign exchange quote.       |

Example:

```json
{
  "bank_id": "string",
  "from_currency_code": "string",
  "to_currency_code": "string",
  "rate": "string",
  "inverse_rate": "string",
  "effective_date": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Foreign Exchange created successfully.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve exchange rate between two currencies

Retrieve the exchange rate from a currency to another

```sh
curl -X GET \
	/v1/fx/{FromCurrencyCode}/{ToCurrencyCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcurrentfxrate []}}

### HTTP Request

`GET /v1/fx/{FromCurrencyCode}/{ToCurrencyCode}`

### Query Parameters

| Name             | Type   | Description                                                 |
|------------------|--------|-------------------------------------------------------------|
| FromCurrencyCode | string | FromCurrencyCode is the currency to transfer from.          |
| ToCurrencyCode   | string | ToCurrencyCode is the currency that we are transferring to. |

### Responses

#### Response body

| Name             | Type   | Description                                                              |
|------------------|--------|--------------------------------------------------------------------------|
| BankID           | string | BankID is an identifier for the bank on this fx transaction.             |
| FromCurrencyCode | string | FromCurrencyCode is the currency to transfer from.                       |
| ToCurrencyCode   | string | ToCurrencyCode is the currency that we are transferring to.              |
| Rate             | string | Rate is the exchange rate of the foreign exchange.                       |
| InverseRate      | string | InverseRate is the inverse of the exchange rate of the foreign exchange. |
| EffectiveDate    | string | EffectiveDate is the effective date of the foreign exchange quote.       |

Example:

```json
{
  "bank_id": "string",
  "from_currency_code": "string",
  "to_currency_code": "string",
  "rate": "string",
  "inverse_rate": "string",
  "effective_date": "string"
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

## Update a foreign exchange quote

Updates a foreign exchange quote

```sh
curl -X PUT \
	/v1/fx \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"fx": {
			"bank_id": "string",
			"from_currency_code": "string",
			"to_currency_code": "string",
			"rate": "string",
			"inverse_rate": "string",
			"effective_date": "string"
		}
	}'
```
{{snippet updatefx []}}

### HTTP Request

`PUT /v1/fx`

### Body Parameters

| Name | Type | Description                                       |
|------|------|---------------------------------------------------|
| FX   | FX   | FX is the foreign exchange information to update. |

##### Objects

###### FX

| Name             | Type   | Description                                                              |
|------------------|--------|--------------------------------------------------------------------------|
| BankID           | string | BankID is an identifier for the bank on this fx transaction.             |
| FromCurrencyCode | string | FromCurrencyCode is the currency to transfer from.                       |
| ToCurrencyCode   | string | ToCurrencyCode is the currency that we are transferring to.              |
| Rate             | string | Rate is the exchange rate of the foreign exchange.                       |
| InverseRate      | string | InverseRate is the inverse of the exchange rate of the foreign exchange. |
| EffectiveDate    | string | EffectiveDate is the effective date of the foreign exchange quote.       |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Foreign Exchange successfully updated.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
