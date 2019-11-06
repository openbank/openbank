# Bank API v1.0.0

Provides create and read operations on the bank resource.

*
Host ``
EOL

*
Base Path ``

## Create a bank

Creates a new bank and returns its id.

```sh
curl -X POST \
	/v1/banks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank": {
			"id": "string",
			"full_name": "string",
			"short_name": "string",
			"logo_url": "string",
			"website_url": "string",
			"swift_bic": "string",
			"national_identifier": "string",
			"bank_routing": {
				"scheme": "string",
				"address": "string"
			}
		}
	}'
```
### HTTP Request

`POST /v1/banks`

### Body Parameters

| Name | Type | Description                                   |
|------|------|-----------------------------------------------|
| Bank | Bank | Bank is the related information about a bank. |

##### Objects

###### Bank

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

### Responses

#### Response body

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

Example:

```json
{
  "id": "string",
  "full_name": "string",
  "short_name": "string",
  "logo_url": "string",
  "website_url": "string",
  "swift_bic": "string",
  "national_identifier": "string",
  "bank_routing": {
    "scheme": "string",
    "address": "string"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Bank created successfully.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a new transaction type at a bank

Creates a new transaction type at a bank and returns its transaction type response.

```sh
curl -X POST \
	/v1/banks/transaction-types \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"transaction_type": {
			"id": "string",
			"bankId": "string",
			"short_code": "string",
			"summary": "string",
			"description": "string",
			"currency": "string",
			"amount": "string"
		}
	}'
```
{{snippet createtranscationtypeatbank []}}

### HTTP Request

`POST /v1/banks/transaction-types`

### Body Parameters

| Name            | Type            | Description                                                                    |
|-----------------|-----------------|--------------------------------------------------------------------------------|
| TransactionType | TransactionType | TransactionType is the related information about a transaction type at a bank. |

##### Objects

###### TransactionType

| Name        | Type   | Description                                                               |
|-------------|--------|---------------------------------------------------------------------------|
| ID          | string | ID is an identifier for the transaction type.                             |
| BankID      | string | BankID is an identifier for the bank for the particular transaction type. |
| ShortCode   | string | ShortCode is the short code of the transaction type.                      |
| Summary     | string | Summary is the summary of the trasnaction type.                           |
| Description | string | Description is the description of the transaction type.                   |
| Currency    | string | Currency is the currency of the transaction type.                         |
| Amount      | string | Amount is the amount of the transaction type.                             |

### Responses

#### Response body

| Name        | Type   | Description                                                               |
|-------------|--------|---------------------------------------------------------------------------|
| ID          | string | ID is an identifier for the transaction type.                             |
| BankID      | string | BankID is an identifier for the bank for the particular transaction type. |
| ShortCode   | string | ShortCode is the short code of the transaction type.                      |
| Summary     | string | Summary is the summary of the trasnaction type.                           |
| Description | string | Description is the description of the transaction type.                   |
| Currency    | string | Currency is the currency of the transaction type.                         |
| Amount      | string | Amount is the amount of the transaction type.                             |

Example:

```json
{
  "id": "string",
  "bankId": "string",
  "short_code": "string",
  "summary": "string",
  "description": "string",
  "currency": "string",
  "amount": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | TransactionType created successfully.                                                  |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a bank

Permanently delete a bank.

```sh
curl -X DELETE \
	/v1/banks/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/banks/{ID}`

### Query Parameters

| Name | Type   | Description                       |
|------|--------|-----------------------------------|
| ID   | string | ID is the bank unique identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Bank successfully deleted.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve bank information

Retrieve information about the bank specified by the ID

```sh
curl -X GET \
	/v1/banks/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks/{ID}`

### Query Parameters

| Name | Type   | Description                       |
|------|--------|-----------------------------------|
| ID   | string | ID is the bank unique identifier. |

### Responses

#### Response body

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

Example:

```json
{
  "id": "string",
  "full_name": "string",
  "short_name": "string",
  "logo_url": "string",
  "website_url": "string",
  "swift_bic": "string",
  "national_identifier": "string",
  "bank_routing": {
    "scheme": "string",
    "address": "string"
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

## Retrieve all available banks

Retrieve information regarding all available banks.

```sh
curl -X GET \
	/v1/banks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks`

### Responses

#### Response body

| Name  | Type   | Description                     |
|-------|--------|---------------------------------|
| Banks | []Bank | Banks is the list of the banks. |

##### Objects

###### Bank

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

Example:

```json
{
  "banks": [
    {
      "id": "string",
      "full_name": "string",
      "short_name": "string",
      "logo_url": "string",
      "website_url": "string",
      "swift_bic": "string",
      "national_identifier": "string",
      "bank_routing": {
        "scheme": "string",
        "address": "string"
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

## Update a bank

Updates a bank's information

```sh
curl -X PUT \
	/v1/banks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank": {
			"id": "string",
			"full_name": "string",
			"short_name": "string",
			"logo_url": "string",
			"website_url": "string",
			"swift_bic": "string",
			"national_identifier": "string",
			"bank_routing": {
				"scheme": "string",
				"address": "string"
			}
		}
	}'
```
### HTTP Request

`PUT /v1/banks`

### Body Parameters

| Name | Type | Description                                   |
|------|------|-----------------------------------------------|
| Bank | Bank | Bank is the related information about a bank. |

##### Objects

###### Bank

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

### Responses

#### Response body

| Name               | Type        | Description                                          |
|--------------------|-------------|------------------------------------------------------|
| ID                 | string      | ID is an identifier for the bank.                    |
| FullName           | string      | FullName is the full name of the bank.               |
| ShortName          | string      | ShortName is the short name of the bank.             |
| LogoURL            | string      | LogoURL is the url for the bank's logo.              |
| WebsiteURL         | string      | WebsiteURL is the url for the bank's website.        |
| SwiftBIC           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| NationalIdentifier | string      | NationalIdentifier is the national identifier code.  |
| BankRouting        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| Scheme  | string | Scheme is the routing scheme.   |
| Address | string | Address is the routing address. |

Example:

```json
{
  "id": "string",
  "full_name": "string",
  "short_name": "string",
  "logo_url": "string",
  "website_url": "string",
  "swift_bic": "string",
  "national_identifier": "string",
  "bank_routing": {
    "scheme": "string",
    "address": "string"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Bank successfully updated.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
