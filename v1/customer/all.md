# Customer API v1.0.0

Provides create and read operations on the customer resource.

* Host `https://`

* Base Path ``

## Add address to the customer. {#method-post-addaddresstocustomer}

Add address to the customer.

```sh
curl -X POST \
	https:///v1/customers/{CustomerID}/address \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"customer_id": "string",
		"line_1": "string",
		"line_2": "string",
		"line_3": "string",
		"city": "string",
		"county": "string",
		"state": "string",
		"postcode": "string",
		"country_code": "string",
		"tags": "[]string",
		"status": "string"
	}'
```

### HTTP Request {#http-request-method-post-addaddresstocustomer}

`POST https:///v1/customers/{CustomerID}/address`

### Query Parameters {#query-parameters-method-post-addaddresstocustomer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Body Parameters {#body-parameters-method-post-addaddresstocustomer}

| Name         | Type      | Description |
|--------------|-----------|-------------|
| customer_id  | string    |             |
| line_1       | string    |             |
| line_2       | string    |             |
| line_3       | string    |             |
| city         | string    |             |
| county       | string    |             |
| state        | string    |             |
| postcode     | string    |             |
| country_code | string    |             |
| tags         | \[]string |             |
| status       | string    |             |

### Responses {#responses-method-post-addaddresstocustomer}

#### Response body {#response-body-method-post-addaddresstocustomer}

| Name                | Type      | Description |
|---------------------|-----------|-------------|
| customer_address_id | string    |             |
| customer_id         | string    |             |
| line_1              | string    |             |
| line_2              | string    |             |
| line_3              | string    |             |
| city                | string    |             |
| county              | string    |             |
| state               | string    |             |
| postcode            | string    |             |
| country_code        | string    |             |
| tags                | \[]string |             |
| status              | string    |             |
| insert_date         | Timestamp |             |

##### Objects {#objects-AddAddressToCustomerResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "customer_address_id": "string",
  "customer_id": "string",
  "line_1": "string",
  "line_2": "string",
  "line_3": "string",
  "city": "string",
  "county": "string",
  "state": "string",
  "postcode": "string",
  "country_code": "string",
  "tags": "[]string",
  "status": "string",
  "insert_date": {
    "seconds": "int64",
    "nanos": "int32"
  }
}
```

#### Response codes {#response-codes-method-post-addaddresstocustomer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add social media handle {#method-post-addsocialmediahandle}

Add social media handle for the customer

```sh
curl -X POST \
	https:///v1/customers/{CustomerID}/social_media_handles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"customer_id": "string",
		"customer_number": "string",
		"type": "string",
		"handle": "string",
		"date_added": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"date_activated": {
			"seconds": "int64",
			"nanos": "int32"
		}
	}'
```

### HTTP Request {#http-request-method-post-addsocialmediahandle}

`POST https:///v1/customers/{CustomerID}/social_media_handles`

### Query Parameters {#query-parameters-method-post-addsocialmediahandle}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |

### Body Parameters {#body-parameters-method-post-addsocialmediahandle}

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| customer_id     | string    |             |
| customer_number | string    |             |
| type            | string    |             |
| handle          | string    |             |
| date_added      | Timestamp |             |
| date_activated  | Timestamp |             |

##### Objects {#objects-AddSocialMediaHandleRequest}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-post-addsocialmediahandle}

#### Response body {#response-body-method-post-addsocialmediahandle}

| Name    | Type   | Description |
|---------|--------|-------------|
| success | string |             |

Example:

```json
{
  "success": "string"
}
```

#### Response codes {#response-codes-method-post-addsocialmediahandle}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Social media added successfully.                                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create credit limit order {#method-post-createcreditlimitorder}

Create Credit limit order for the customer on specific bank

```sh
curl -X POST \
	https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"requested_current_rate_amount1": "string",
		"requested_current_rate_amount2": "string",
		"requested_current_valid_end_date": "string",
		"current_credit_documentation": "string",
		"temporary_requested_current_amount": "string",
		"requested_temporary_valid_end_date": "string",
		"temporary_credit_documentation": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcreditlimitorder}

`POST https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests`

### Query Parameters {#query-parameters-method-post-createcreditlimitorder}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-post-createcreditlimitorder}

| Name                               | Type   | Description |
|------------------------------------|--------|-------------|
| bank_id                            | string |             |
| customer_id                        | string |             |
| requested_current_rate_amount1     | string |             |
| requested_current_rate_amount2     | string |             |
| requested_current_valid_end_date   | string |             |
| current_credit_documentation       | string |             |
| temporary_requested_current_amount | string |             |
| requested_temporary_valid_end_date | string |             |
| temporary_credit_documentation     | string |             |

### Responses {#responses-method-post-createcreditlimitorder}

#### Response body {#response-body-method-post-createcreditlimitorder}

| Name                  | Type   | Description |
|-----------------------|--------|-------------|
| credit_limit_order_id | string |             |

Example:

```json
{
  "credit_limit_order_id": "string"
}
```

#### Response codes {#response-codes-method-post-createcreditlimitorder}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a customer {#method-post-createcustomer}

Creates a new customer

```sh
curl -X POST \
	https:///v1/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"legal_name": "string",
		"mobile_phone_number": "string",
		"email": "string",
		"face_image": {
			"url": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			}
		},
		"date_of_birth": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"relationship_status": "string",
		"dob_of_dependents": [
			{
				"seconds": "int64",
				"nanos": "int32"
			}
		],
		"credit_rating": {
			"rating": "string",
			"source": "string"
		},
		"credit_limit": {
			"cur": "string",
			"num": "string"
		},
		"highest_education_attained": "string",
		"employment_status": "string",
		"last_ok_date": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"title": "string",
		"branchId": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcustomer}

`POST https:///v1/customers`

### Body Parameters {#body-parameters-method-post-createcustomer}

| Name                       | Type         | Description |
|----------------------------|--------------|-------------|
| bank_id                    | string       |             |
| legal_name                 | string       |             |
| mobile_phone_number        | string       |             |
| email                      | string       |             |
| face_image                 | FaceImage    |             |
| date_of_birth              | Timestamp    |             |
| relationship_status        | string       |             |
| dob_of_dependents          | \[]Timestamp |             |
| credit_rating              | CreditRating |             |
| credit_limit               | Amount       |             |
| highest_education_attained | string       |             |
| employment_status          | string       |             |
| last_ok_date               | Timestamp    |             |
| title                      | string       |             |
| branchId                   | string       |             |

##### Objects {#objects-CreateCustomerRequest}

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-post-createcustomer}

#### Response body {#response-body-method-post-createcustomer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| customer_id | string |             |
| bank_id     | string |             |

Example:

```json
{
  "customer_id": "string",
  "bank_id": "string"
}
```

#### Response codes {#response-codes-method-post-createcustomer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create user customer link {#method-post-createusercustomerlink}

Creare user customer link

```sh
curl -X POST \
	https:///v1/banks/{BankID}/user_customer_links \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"user_id": "string",
		"customer_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createusercustomerlink}

`POST https:///v1/banks/{BankID}/user_customer_links`

### Query Parameters {#query-parameters-method-post-createusercustomerlink}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Body Parameters {#body-parameters-method-post-createusercustomerlink}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| user_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-post-createusercustomerlink}

#### Response body {#response-body-method-post-createusercustomerlink}

| Name                  | Type      | Description |
|-----------------------|-----------|-------------|
| user_customer_link_id | string    |             |
| customer_id           | string    |             |
| user_id               | string    |             |
| date_inserted         | Timestamp |             |
| is_active             | bool      |             |

##### Objects {#objects-CreateUserCustomerLinkResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "user_customer_link_id": "string",
  "customer_id": "string",
  "user_id": "string",
  "date_inserted": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "is_active": "bool"
}
```

#### Response codes {#response-codes-method-post-createusercustomerlink}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete customer address {#method-delete-deletecustomeraddress}

Permanently delete customer address.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletecustomeraddress}

`DELETE https:///v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID}`

### Query Parameters {#query-parameters-method-delete-deletecustomeraddress}

| Name                | Type   | Description |
|---------------------|--------|-------------|
| bank_id             | string |             |
| customer_id         | string |             |
| customer_address_id | string |             |

### Responses {#responses-method-delete-deletecustomeraddress}

#### Response body {#response-body-method-delete-deletecustomeraddress}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletecustomeraddress}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Customer address successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete tax residence {#method-delete-deletetaxresidence}

Permanently delete tax residence.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/customers/{CustomerID}/tax_residencies/{TaxResidenceID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletetaxresidence}

`DELETE https:///v1/banks/{BankID}/customers/{CustomerID}/tax_residencies/{TaxResidenceID}`

### Query Parameters {#query-parameters-method-delete-deletetaxresidence}

| Name             | Type   | Description |
|------------------|--------|-------------|
| bank_id          | string |             |
| customer_id      | string |             |
| tax_residence_id | string |             |

### Responses {#responses-method-delete-deletetaxresidence}

#### Response body {#response-body-method-delete-deletetaxresidence}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletetaxresidence}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Tax residence successfully deleted.                                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all CRM events {#method-get-getcrmevents}

Returns a list containing up to 20 crm events.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/crm-events \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcrmevents}

`GET https:///v1/banks/{BankID}/crm-events`

### Query Parameters {#query-parameters-method-get-getcrmevents}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getcrmevents}

#### Response body {#response-body-method-get-getcrmevents}

| Name       | Type        | Description |
|------------|-------------|-------------|
| crm_events | \[]CRMEvent |             |
| has_more   | bool        |             |

##### Objects {#objects-GetCRMEventsResponse}

###### CRMEvent

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| id              | string    |             |
| bank_id         | string    |             |
| customer_name   | string    |             |
| customer_number | string    |             |
| category        | string    |             |
| detail          | string    |             |
| channel         | string    |             |
| scheduled_date  | Timestamp |             |
| actual_date     | Timestamp |             |
| result          | string    |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "crm_events": [
    {
      "id": "string",
      "bank_id": "string",
      "customer_name": "string",
      "customer_number": "string",
      "category": "string",
      "detail": "string",
      "channel": "string",
      "scheduled_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "actual_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "result": "string"
    }
  ],
  "has_more": "bool"
}
```

#### Response codes {#response-codes-method-get-getcrmevents}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all credit limit order {#method-get-getcreditlimitorder}

Returns a list containing up to 20 CreditLimitOrder.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcreditlimitorder}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests`

### Query Parameters {#query-parameters-method-get-getcreditlimitorder}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-get-getcreditlimitorder}

#### Response body {#response-body-method-get-getcreditlimitorder}

| Name   | Type                | Description |
|--------|---------------------|-------------|
| result | \[]CreditLimitOrder |             |

##### Objects {#objects-GetCreditLimitOrderResponse}

###### CreditLimitOrder

| Name               | Type   | Description |
|--------------------|--------|-------------|
| rank_amount_1      | string |             |
| nominal_interest_1 | string |             |
| rank_amount_2      | string |             |
| nominal_interest_2 | string |             |

Example:

```json
{
  "result": [
    {
      "rank_amount_1": "string",
      "nominal_interest_1": "string",
      "rank_amount_2": "string",
      "nominal_interest_2": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcreditlimitorder}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get credit limit order by id {#method-get-getcreditlimitorderbyid}

Returns Credit limit order by id

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/request/{RequestID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcreditlimitorderbyid}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/credit_limit/request/{RequestID}`

### Query Parameters {#query-parameters-method-get-getcreditlimitorderbyid}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |
| request_id  | string |             |

### Responses {#responses-method-get-getcreditlimitorderbyid}

#### Response body {#response-body-method-get-getcreditlimitorderbyid}

| Name               | Type             | Description |
|--------------------|------------------|-------------|
| credit_limit_order | CreditLimitOrder |             |

##### Objects {#objects-GetCreditLimitOrderByIDResponse}

###### CreditLimitOrder

| Name               | Type   | Description |
|--------------------|--------|-------------|
| rank_amount_1      | string |             |
| nominal_interest_1 | string |             |
| rank_amount_2      | string |             |
| nominal_interest_2 | string |             |

Example:

```json
{
  "credit_limit_order": {
    "rank_amount_1": "string",
    "nominal_interest_1": "string",
    "rank_amount_2": "string",
    "nominal_interest_2": "string"
  }
}
```

#### Response codes {#response-codes-method-get-getcreditlimitorderbyid}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all customer adddresses {#method-get-getcustomeraddresses}

Returns a list containing up to 20 customer addresses.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/addresses \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomeraddresses}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/addresses`

### Query Parameters {#query-parameters-method-get-getcustomeraddresses}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-get-getcustomeraddresses}

#### Response body {#response-body-method-get-getcustomeraddresses}

| Name      | Type       | Description |
|-----------|------------|-------------|
| addresses | \[]Address |             |

##### Objects {#objects-GetCustomerAddressesResponse}

###### Address

| Name                | Type      | Description |
|---------------------|-----------|-------------|
| customer_address_id | string    |             |
| customer_id         | string    |             |
| line_1              | string    |             |
| line_2              | string    |             |
| line_3              | string    |             |
| city                | string    |             |
| county              | string    |             |
| state               | string    |             |
| postcode            | string    |             |
| country_code        | string    |             |
| tags                | \[]string |             |
| status              | string    |             |
| insert_date         | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "addresses": [
    {
      "customer_address_id": "string",
      "customer_id": "string",
      "line_1": "string",
      "line_2": "string",
      "line_3": "string",
      "city": "string",
      "county": "string",
      "state": "string",
      "postcode": "string",
      "country_code": "string",
      "tags": "[]string",
      "status": "string",
      "insert_date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomeraddresses}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer By ID {#method-get-getcustomerbycustomerid}

Returns the customer data based on it's id

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomerbycustomerid}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}`

### Query Parameters {#query-parameters-method-get-getcustomerbycustomerid}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-get-getcustomerbycustomerid}

#### Response body {#response-body-method-get-getcustomerbycustomerid}

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

##### Objects {#objects-Customer}

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "id": "string",
  "bank_id": "string",
  "customer_number": "string",
  "legal_name": "string",
  "phone_number": "string",
  "email": "string",
  "face_image": {
    "url": "string",
    "date": {
      "seconds": "int64",
      "nanos": "int32"
    }
  },
  "date_of_birth": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "relationship_status": "string",
  "credit_rating": {
    "rating": "string",
    "source": "string"
  },
  "credit_limit": {
    "cur": "string",
    "num": "string"
  },
  "highest_educational_attained": "string",
  "employment_status": "string",
  "kyc_status": "bool",
  "last_ok_data": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "title": "string",
  "branch_id": "string"
}
```

#### Response codes {#response-codes-method-get-getcustomerbycustomerid}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer By CustomerNumber {#method-get-getcustomerbycustomernumber}

Returns the customer data based on the customer number

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/customer-number/{CustomerNumber} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomerbycustomernumber}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/customer-number/{CustomerNumber}`

### Query Parameters {#query-parameters-method-get-getcustomerbycustomernumber}

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| customer_id     | string |             |
| customer_number | string |             |

### Responses {#responses-method-get-getcustomerbycustomernumber}

#### Response body {#response-body-method-get-getcustomerbycustomernumber}

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

##### Objects {#objects-Customer}

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "id": "string",
  "bank_id": "string",
  "customer_number": "string",
  "legal_name": "string",
  "phone_number": "string",
  "email": "string",
  "face_image": {
    "url": "string",
    "date": {
      "seconds": "int64",
      "nanos": "int32"
    }
  },
  "date_of_birth": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "relationship_status": "string",
  "credit_rating": {
    "rating": "string",
    "source": "string"
  },
  "credit_limit": {
    "cur": "string",
    "num": "string"
  },
  "highest_educational_attained": "string",
  "employment_status": "string",
  "kyc_status": "bool",
  "last_ok_data": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "title": "string",
  "branch_id": "string"
}
```

#### Response codes {#response-codes-method-get-getcustomerbycustomernumber}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all customer social media handles {#method-get-getcustomersocialmediahandles}

Returns a list containing up to 20 customer social media handles

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/social_media_handles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomersocialmediahandles}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/social_media_handles`

### Query Parameters {#query-parameters-method-get-getcustomersocialmediahandles}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-get-getcustomersocialmediahandles}

#### Response body {#response-body-method-get-getcustomersocialmediahandles}

| Name   | Type     | Description |
|--------|----------|-------------|
| checks | \[]Check |             |

##### Objects {#objects-GetCustomerSocialMediaHandlesResponse}

###### Check

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| customer_number | string    |             |
| type            | string    |             |
| handle          | string    |             |
| date_added      | Timestamp |             |
| date_activated  | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "checks": [
    {
      "customer_number": "string",
      "type": "string",
      "handle": "string",
      "date_added": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "date_activated": {
        "seconds": "int64",
        "nanos": "int32"
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomersocialmediahandles}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers data for current user {#method-get-getcustomersforcurrentuser}

Returns the current user customers data

```sh
curl -X GET \
	https:///v1/users/current/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomersforcurrentuser}

`GET https:///v1/users/current/customers`

### Responses {#responses-method-get-getcustomersforcurrentuser}

#### Response body {#response-body-method-get-getcustomersforcurrentuser}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| customers | \[]Customer |             |

##### Objects {#objects-GetCustomersForCurrentUserResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customers": [
    {
      "id": "string",
      "bank_id": "string",
      "customer_number": "string",
      "legal_name": "string",
      "phone_number": "string",
      "email": "string",
      "face_image": {
        "url": "string",
        "date": {
          "seconds": "int64",
          "nanos": "int32"
        }
      },
      "date_of_birth": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "relationship_status": "string",
      "credit_rating": {
        "rating": "string",
        "source": "string"
      },
      "credit_limit": {
        "cur": "string",
        "num": "string"
      },
      "highest_educational_attained": "string",
      "employment_status": "string",
      "kyc_status": "bool",
      "last_ok_data": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "title": "string",
      "branch_id": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomersforcurrentuser}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers data by bank id {#method-get-getcustomersforcurrentuseratbank}

Returns the current user customers data by bank ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcustomersforcurrentuseratbank}

`GET https:///v1/banks/{BankID}/customers`

### Query Parameters {#query-parameters-method-get-getcustomersforcurrentuseratbank}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getcustomersforcurrentuseratbank}

#### Response body {#response-body-method-get-getcustomersforcurrentuseratbank}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| customers | \[]Customer |             |

##### Objects {#objects-GetCustomersForCurrentUserAtBankResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customers": [
    {
      "id": "string",
      "bank_id": "string",
      "customer_number": "string",
      "legal_name": "string",
      "phone_number": "string",
      "email": "string",
      "face_image": {
        "url": "string",
        "date": {
          "seconds": "int64",
          "nanos": "int32"
        }
      },
      "date_of_birth": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "relationship_status": "string",
      "credit_rating": {
        "rating": "string",
        "source": "string"
      },
      "credit_limit": {
        "cur": "string",
        "num": "string"
      },
      "highest_educational_attained": "string",
      "employment_status": "string",
      "kyc_status": "bool",
      "last_ok_data": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "title": "string",
      "branch_id": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getcustomersforcurrentuseratbank}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Return customer firehose {#method-get-getfirehosecustomer}

Returns the list of customer firehose

```sh
curl -X GET \
	https:///v1/banks/{BankID}/firehose/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getfirehosecustomer}

`GET https:///v1/banks/{BankID}/firehose/customers`

### Query Parameters {#query-parameters-method-get-getfirehosecustomer}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getfirehosecustomer}

#### Response body {#response-body-method-get-getfirehosecustomer}

| Name      | Type        | Description |
|-----------|-------------|-------------|
| customers | \[]Customer |             |

##### Objects {#objects-GetFirehoseCustomerResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customers": [
    {
      "id": "string",
      "bank_id": "string",
      "customer_number": "string",
      "legal_name": "string",
      "phone_number": "string",
      "email": "string",
      "face_image": {
        "url": "string",
        "date": {
          "seconds": "int64",
          "nanos": "int32"
        }
      },
      "date_of_birth": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "relationship_status": "string",
      "credit_rating": {
        "rating": "string",
        "source": "string"
      },
      "credit_limit": {
        "cur": "string",
        "num": "string"
      },
      "highest_educational_attained": "string",
      "employment_status": "string",
      "kyc_status": "bool",
      "last_ok_data": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "title": "string",
      "branch_id": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getfirehosecustomer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers tax_residencies {#method-get-gettaxresidenceofcustomer}

Returns the user tax_residencies

```sh
curl -X GET \
	https:///v1/banks/{BankID}/customers/{CustomerID}/tax-residences \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-gettaxresidenceofcustomer}

`GET https:///v1/banks/{BankID}/customers/{CustomerID}/tax-residences`

### Query Parameters {#query-parameters-method-get-gettaxresidenceofcustomer}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Responses {#responses-method-get-gettaxresidenceofcustomer}

#### Response body {#response-body-method-get-gettaxresidenceofcustomer}

| Name   | Type            | Description |
|--------|-----------------|-------------|
| result | \[]TaxResidence |             |

##### Objects {#objects-GetTaxResidenceOfCustomerResponse}

###### TaxResidence

| Name             | Type   | Description |
|------------------|--------|-------------|
| domain           | string |             |
| tax_number       | string |             |
| tax_residence_id | string |             |

Example:

```json
{
  "result": [
    {
      "domain": "string",
      "tax_number": "string",
      "tax_residence_id": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-gettaxresidenceofcustomer}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer addresses. {#method-put-updatecustomeraddress}

Update customer addresses information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"customer_address_id": "string",
		"line_1": "string",
		"line_2": "string",
		"line_3": "string",
		"city": "string",
		"county": "string",
		"state": "string",
		"postcode": "string",
		"country_code": "string",
		"tags": "[]string",
		"status": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomeraddress}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID}`

### Query Parameters {#query-parameters-method-put-updatecustomeraddress}

| Name                | Type   | Description |
|---------------------|--------|-------------|
| bank_id             | string |             |
| customer_id         | string |             |
| customer_address_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomeraddress}

| Name                | Type      | Description |
|---------------------|-----------|-------------|
| bank_id             | string    |             |
| customer_id         | string    |             |
| customer_address_id | string    |             |
| line_1              | string    |             |
| line_2              | string    |             |
| line_3              | string    |             |
| city                | string    |             |
| county              | string    |             |
| state               | string    |             |
| postcode            | string    |             |
| country_code        | string    |             |
| tags                | \[]string |             |
| status              | string    |             |

### Responses {#responses-method-put-updatecustomeraddress}

#### Response body {#response-body-method-put-updatecustomeraddress}

| Name                | Type      | Description |
|---------------------|-----------|-------------|
| customer_address_id | string    |             |
| customer_id         | string    |             |
| line_1              | string    |             |
| line_2              | string    |             |
| line_3              | string    |             |
| city                | string    |             |
| county              | string    |             |
| state               | string    |             |
| postcode            | string    |             |
| country_code        | string    |             |
| tags                | \[]string |             |
| status              | string    |             |
| insert_date         | Timestamp |             |

##### Objects {#objects-UpdateCustomerAddressResponse}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "customer_address_id": "string",
  "customer_id": "string",
  "line_1": "string",
  "line_2": "string",
  "line_3": "string",
  "city": "string",
  "county": "string",
  "state": "string",
  "postcode": "string",
  "country_code": "string",
  "tags": "[]string",
  "status": "string",
  "insert_date": {
    "seconds": "int64",
    "nanos": "int32"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomeraddress}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer branch. {#method-put-updatecustomerbranch}

Update customer branch information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/branch \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"branch_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomerbranch}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/branch`

### Query Parameters {#query-parameters-method-put-updatecustomerbranch}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomerbranch}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |
| branch_id   | string |             |

### Responses {#responses-method-put-updatecustomerbranch}

#### Response body {#response-body-method-put-updatecustomerbranch}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerBranchResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomerbranch}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer credit limit. {#method-put-updatecustomercreditlimit}

Update customer credit limit information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/credit-limit \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"credit_limit": {
			"cur": "string",
			"num": "string"
		}
	}'
```

### HTTP Request {#http-request-method-put-updatecustomercreditlimit}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/credit-limit`

### Query Parameters {#query-parameters-method-put-updatecustomercreditlimit}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomercreditlimit}

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| customer_id  | string |             |
| credit_limit | Amount |             |

##### Objects {#objects-UpdateCustomerCreditLimitRequest}

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-put-updatecustomercreditlimit}

#### Response body {#response-body-method-put-updatecustomercreditlimit}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerCreditLimitResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomercreditlimit}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer credut rating and source. {#method-put-updatecustomercreditratingandsource}

Update customer credit rating and source information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/credit-rating-and-source \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"credit_rating": {
			"rating": "string",
			"source": "string"
		}
	}'
```

### HTTP Request {#http-request-method-put-updatecustomercreditratingandsource}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/credit-rating-and-source`

### Query Parameters {#query-parameters-method-put-updatecustomercreditratingandsource}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomercreditratingandsource}

| Name          | Type         | Description |
|---------------|--------------|-------------|
| bank_id       | string       |             |
| customer_id   | string       |             |
| credit_rating | CreditRating |             |

##### Objects {#objects-UpdateCustomerCreditRatingAndSourceRequest}

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

### Responses {#responses-method-put-updatecustomercreditratingandsource}

#### Response body {#response-body-method-put-updatecustomercreditratingandsource}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerCreditRatingAndSourceResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomercreditratingandsource}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer email. {#method-put-updatecustomeremail}

Update customer email information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/email \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"email": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomeremail}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/email`

### Query Parameters {#query-parameters-method-put-updatecustomeremail}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomeremail}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |
| email       | string |             |

### Responses {#responses-method-put-updatecustomeremail}

#### Response body {#response-body-method-put-updatecustomeremail}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerEmailResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomeremail}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer identity. {#method-put-updatecustomeridentity}

Update customer identity information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/identity \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"legal_name": "string",
		"date_of_birth": "string",
		"title": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomeridentity}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/identity`

### Query Parameters {#query-parameters-method-put-updatecustomeridentity}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomeridentity}

| Name          | Type   | Description |
|---------------|--------|-------------|
| bank_id       | string |             |
| customer_id   | string |             |
| legal_name    | string |             |
| date_of_birth | string |             |
| title         | string |             |

### Responses {#responses-method-put-updatecustomeridentity}

#### Response body {#response-body-method-put-updatecustomeridentity}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerIdentityResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomeridentity}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer mobile number. {#method-put-updatecustomermobilenumber}

Update customer mobile number information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/mobile-number \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"mobile_phone_number": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomermobilenumber}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/mobile-number`

### Query Parameters {#query-parameters-method-put-updatecustomermobilenumber}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomermobilenumber}

| Name                | Type   | Description |
|---------------------|--------|-------------|
| bank_id             | string |             |
| customer_id         | string |             |
| mobile_phone_number | string |             |

### Responses {#responses-method-put-updatecustomermobilenumber}

#### Response body {#response-body-method-put-updatecustomermobilenumber}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerMobileNumberResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomermobilenumber}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer number. {#method-put-updatecustomernumber}

Update customer number information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/number \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"customer_number": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomernumber}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/number`

### Query Parameters {#query-parameters-method-put-updatecustomernumber}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomernumber}

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| customer_id     | string |             |
| customer_number | string |             |

### Responses {#responses-method-put-updatecustomernumber}

#### Response body {#response-body-method-put-updatecustomernumber}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerNumberResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomernumber}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer data. {#method-put-updatecustomerotherdata}

Update customer data information.

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/customers/{CustomerID}/data \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"face_image": {
			"url": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			}
		},
		"relationship_status": "string",
		"dependants": "int32",
		"highest_education_attained": "string",
		"employment_status": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecustomerotherdata}

`PUT https:///v1/banks/{BankID}/customers/{CustomerID}/data`

### Query Parameters {#query-parameters-method-put-updatecustomerotherdata}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| customer_id | string |             |

### Body Parameters {#body-parameters-method-put-updatecustomerotherdata}

| Name                       | Type      | Description |
|----------------------------|-----------|-------------|
| bank_id                    | string    |             |
| customer_id                | string    |             |
| face_image                 | FaceImage |             |
| relationship_status        | string    |             |
| dependants                 | int32     |             |
| highest_education_attained | string    |             |
| employment_status          | string    |             |

##### Objects {#objects-UpdateCustomerOtherDataRequest}

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-put-updatecustomerotherdata}

#### Response body {#response-body-method-put-updatecustomerotherdata}

| Name     | Type     | Description |
|----------|----------|-------------|
| customer | Customer |             |

##### Objects {#objects-UpdateCustomerOtherDataResponse}

###### Customer

| Name                         | Type         | Description |
|------------------------------|--------------|-------------|
| id                           | string       |             |
| bank_id                      | string       |             |
| customer_number              | string       |             |
| legal_name                   | string       |             |
| phone_number                 | string       |             |
| email                        | string       |             |
| face_image                   | FaceImage    |             |
| date_of_birth                | Timestamp    |             |
| relationship_status          | string       |             |
| credit_rating                | CreditRating |             |
| credit_limit                 | Amount       |             |
| highest_educational_attained | string       |             |
| employment_status            | string       |             |
| kyc_status                   | bool         |             |
| last_ok_data                 | Timestamp    |             |
| title                        | string       |             |
| branch_id                    | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| url  | string    |             |
| date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "customer": {
    "id": "string",
    "bank_id": "string",
    "customer_number": "string",
    "legal_name": "string",
    "phone_number": "string",
    "email": "string",
    "face_image": {
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    },
    "date_of_birth": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "relationship_status": "string",
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_educational_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "last_ok_data": {
      "seconds": "int64",
      "nanos": "int32"
    },
    "title": "string",
    "branch_id": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecustomerotherdata}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
