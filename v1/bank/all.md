Bank API v1.0.0
===============

Provides create and read operations on the bank resource.

* Host `https://`

* Base Path ``

Create a bank {#method-post-createbank}
---------------------------------------

Creates a new bank and returns its id.

```sh
curl -X POST \
	https:///v1/banks \
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

### HTTP Request {#http-request-method-post-createbank}

`POST https:///v1/banks`

### Body Parameters {#body-parameters-method-post-createbank}

| Name | Type | Description                                   |
|------|------|-----------------------------------------------|
| bank | Bank | Bank is the related information about a bank. |

##### Objects {#objects-CreateBankRequest}

###### Bank

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

### Responses {#responses-method-post-createbank}

#### Response body {#response-body-method-post-createbank}

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects {#objects-Bank}

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

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

#### Response codes {#response-codes-method-post-createbank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Bank created successfully.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a new transaction type at a bank {#method-post-createtranscationtypeatbank}
----------------------------------------------------------------------------------

Creates a new transaction type at a bank and returns its transaction type response.

```sh
curl -X POST \
	https:///v1/banks/transaction-types \
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

### HTTP Request {#http-request-method-post-createtranscationtypeatbank}

`POST https:///v1/banks/transaction-types`

### Body Parameters {#body-parameters-method-post-createtranscationtypeatbank}

| Name             | Type            | Description                                                                    |
|------------------|-----------------|--------------------------------------------------------------------------------|
| transaction_type | TransactionType | TransactionType is the related information about a transaction type at a bank. |

##### Objects {#objects-CreateTransactionTypeAtBankRequest}

###### TransactionType

| Name        | Type   | Description                                                               |
|-------------|--------|---------------------------------------------------------------------------|
| id          | string | ID is an identifier for the transaction type.                             |
| bankId      | string | BankID is an identifier for the bank for the particular transaction type. |
| short_code  | string | ShortCode is the short code of the transaction type.                      |
| summary     | string | Summary is the summary of the trasnaction type.                           |
| description | string | Description is the description of the transaction type.                   |
| currency    | string | Currency is the currency of the transaction type.                         |
| amount      | string | Amount is the amount of the transaction type.                             |

### Responses {#responses-method-post-createtranscationtypeatbank}

#### Response body {#response-body-method-post-createtranscationtypeatbank}

| Name        | Type   | Description                                                               |
|-------------|--------|---------------------------------------------------------------------------|
| id          | string | ID is an identifier for the transaction type.                             |
| bankId      | string | BankID is an identifier for the bank for the particular transaction type. |
| short_code  | string | ShortCode is the short code of the transaction type.                      |
| summary     | string | Summary is the summary of the trasnaction type.                           |
| description | string | Description is the description of the transaction type.                   |
| currency    | string | Currency is the currency of the transaction type.                         |
| amount      | string | Amount is the amount of the transaction type.                             |

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

#### Response codes {#response-codes-method-post-createtranscationtypeatbank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | TransactionType created successfully.                                                  |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a bank {#method-delete-deletebank}
-----------------------------------------

Permanently delete a bank.

```sh
curl -X DELETE \
	https:///v1/banks/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletebank}

`DELETE https:///v1/banks/{ID}`

### Query Parameters {#query-parameters-method-delete-deletebank}

| Name | Type   | Description                       |
|------|--------|-----------------------------------|
| id   | string | ID is the bank unique identifier. |

### Responses {#responses-method-delete-deletebank}

#### Response codes {#response-codes-method-delete-deletebank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Bank successfully deleted.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve bank information {#method-get-getbank}
-----------------------------------------------

Retrieve information about the bank specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getbank}

`GET https:///v1/banks/{ID}`

### Query Parameters {#query-parameters-method-get-getbank}

| Name | Type   | Description                       |
|------|--------|-----------------------------------|
| id   | string | ID is the bank unique identifier. |

### Responses {#responses-method-get-getbank}

#### Response body {#response-body-method-get-getbank}

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects {#objects-Bank}

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

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

#### Response codes {#response-codes-method-get-getbank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available banks {#method-get-getbanks}
---------------------------------------------------

Retrieve information regarding all available banks.

```sh
curl -X GET \
	https:///v1/banks \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getbanks}

`GET https:///v1/banks`

### Responses {#responses-method-get-getbanks}

#### Response body {#response-body-method-get-getbanks}

| Name  | Type    | Description                     |
|-------|---------|---------------------------------|
| banks | \[]Bank | Banks is the list of the banks. |

##### Objects {#objects-GetBanksResponse}

###### Bank

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

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

#### Response codes {#response-codes-method-get-getbanks}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a bank {#method-put-updatebank}
--------------------------------------

Updates a bank's information

```sh
curl -X PUT \
	https:///v1/banks \
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

### HTTP Request {#http-request-method-put-updatebank}

`PUT https:///v1/banks`

### Body Parameters {#body-parameters-method-put-updatebank}

| Name | Type | Description                                   |
|------|------|-----------------------------------------------|
| bank | Bank | Bank is the related information about a bank. |

##### Objects {#objects-UpdateBankRequest}

###### Bank

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

### Responses {#responses-method-put-updatebank}

#### Response body {#response-body-method-put-updatebank}

| Name                | Type        | Description                                          |
|---------------------|-------------|------------------------------------------------------|
| id                  | string      | ID is an identifier for the bank.                    |
| full_name           | string      | FullName is the full name of the bank.               |
| short_name          | string      | ShortName is the short name of the bank.             |
| logo_url            | string      | LogoURL is the url for the bank's logo.              |
| website_url         | string      | WebsiteURL is the url for the bank's website.        |
| swift_bic           | string      | SwiftBIC is the SWIFT bank identifier code.          |
| national_identifier | string      | NationalIdentifier is the national identifier code.  |
| bank_routing        | BankRouting | BankRouting is the routing information for the bank. |

##### Objects {#objects-Bank}

###### BankRouting

| Name    | Type   | Description                     |
|---------|--------|---------------------------------|
| scheme  | string | Scheme is the routing scheme.   |
| address | string | Address is the routing address. |

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

#### Response codes {#response-codes-method-put-updatebank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Bank successfully updated.                                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
