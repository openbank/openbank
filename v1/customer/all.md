# Customer API v1.0.0

Provides create and read operations on the customer resource.

*
Host ``
EOL

*
Base Path ``

## Add address to the customer.

Add address to the customer.

```sh
curl -X POST \
	/v1/customers/{CustomerID}/address \
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
{{snippet addaddressto_customer []}}

### HTTP Request

`POST /v1/customers/{CustomerID}/address`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Body Parameters

| Name        | Type     | Description |
|-------------|----------|-------------|
| CustomerID  | string   |             |
| Line1       | string   |             |
| Line2       | string   |             |
| Line3       | string   |             |
| City        | string   |             |
| County      | string   |             |
| State       | string   |             |
| Postcode    | string   |             |
| CountryCode | string   |             |
| Tags        | []string |             |
| Status      | string   |             |

### Responses

#### Response body

| Name              | Type      | Description |
|-------------------|-----------|-------------|
| CustomerAddressID | string    |             |
| CustomerID        | string    |             |
| Line1             | string    |             |
| Line2             | string    |             |
| Line3             | string    |             |
| City              | string    |             |
| County            | string    |             |
| State             | string    |             |
| Postcode          | string    |             |
| CountryCode       | string    |             |
| Tags              | []string  |             |
| Status            | string    |             |
| InsertDate        | Timestamp |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Add social media handle

Add social media handle for the customer

```sh
curl -X POST \
	/v1/customers/{CustomerID}/social_media_handles \
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
{{snippet addsocialmedia_handle []}}

### HTTP Request

`POST /v1/customers/{CustomerID}/social_media_handles`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |

### Body Parameters

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CustomerID     | string    |             |
| CustomerNumber | string    |             |
| Type           | string    |             |
| Handle         | string    |             |
| DateAdded      | Timestamp |             |
| DateActivated  | Timestamp |             |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name    | Type   | Description |
|---------|--------|-------------|
| Success | string |             |

Example:

```json
{
  "success": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Social media added successfully.                                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create credit limit order

Create Credit limit order for the customer on specific bank

```sh
curl -X POST \
	/v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests \
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
{{snippet createcreditlimit_order []}}

### HTTP Request

`POST /v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name                            | Type   | Description |
|---------------------------------|--------|-------------|
| BankID                          | string |             |
| CustomerID                      | string |             |
| RequestedCurrentRateAmount1     | string |             |
| RequestedCurrentRateAmount2     | string |             |
| RequestedCurrentValidEndDate    | string |             |
| CurrentCreditDocumentation      | string |             |
| TemporaryRequestedCurrentAmount | string |             |
| RequestedTemporaryValidEndDate  | string |             |
| TemporaryCreditDocumentation    | string |             |

### Responses

#### Response body

| Name               | Type   | Description |
|--------------------|--------|-------------|
| CreditLimitOrderID | string |             |

Example:

```json
{
  "credit_limit_order_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a customer

Creates a new customer

```sh
curl -X POST \
	/v1/customers \
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
### HTTP Request

`POST /v1/customers`

### Body Parameters

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| BankID                   | string       |             |
| LegalName                | string       |             |
| MobilePhoneNumber        | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| DobOfDependents          | []Timestamp  |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

##### Objects

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name       | Type   | Description |
|------------|--------|-------------|
| CustomerID | string |             |
| BankID     | string |             |

Example:

```json
{
  "customer_id": "string",
  "bank_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create user customer link

Creare user customer link

```sh
curl -X POST \
	/v1/banks/{BankID}/user_customer_links \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"user_id": "string",
		"customer_id": "string"
	}'
```
{{snippet createusercustomer_link []}}

### HTTP Request

`POST /v1/banks/{BankID}/user_customer_links`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| UserID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name               | Type      | Description |
|--------------------|-----------|-------------|
| UserCustomerLinkID | string    |             |
| CustomerID         | string    |             |
| UserID             | string    |             |
| DateInserted       | Timestamp |             |
| IsActive           | bool      |             |

##### Objects

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete customer address

Permanently delete customer address.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletecustomeraddress []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID}`

### Query Parameters

| Name              | Type   | Description |
|-------------------|--------|-------------|
| BankID            | string |             |
| CustomerID        | string |             |
| CustomerAddressID | string |             |

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
| 204    | Customer address successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete tax residence

Permanently delete tax residence.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/customers/{CustomerID}/tax_residencies/{TaxResidenceID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletetaxresidence []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/customers/{CustomerID}/tax_residencies/{TaxResidenceID}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CustomerID     | string |             |
| TaxResidenceID | string |             |

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
| 204    | Tax residence successfully deleted.                                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all CRM events

Returns a list containing up to 20 crm events.

```sh
curl -X GET \
	/v1/banks/{BankID}/crm-events \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcrmevents []}}

### HTTP Request

`GET /v1/banks/{BankID}/crm-events`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| CRMEvents | []CRMEvent |             |
| HasMore   | bool       |             |

##### Objects

###### CRMEvent

| Name           | Type      | Description |
|----------------|-----------|-------------|
| ID             | string    |             |
| BankID         | string    |             |
| CustomerName   | string    |             |
| CustomerNumber | string    |             |
| Category       | string    |             |
| Detail         | string    |             |
| Channel        | string    |             |
| ScheduledDate  | Timestamp |             |
| ActualDate     | Timestamp |             |
| Result         | string    |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all credit limit order

Returns a list containing up to 20 CreditLimitOrder.

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcreditlimit_order []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/credit_limit/requests`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name   | Type               | Description |
|--------|--------------------|-------------|
| Result | []CreditLimitOrder |             |

##### Objects

###### CreditLimitOrder

| Name             | Type   | Description |
|------------------|--------|-------------|
| RankAmount1      | string |             |
| NominalInterest1 | string |             |
| RankAmount2      | string |             |
| NominalInterest2 | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get credit limit order by id

Returns Credit limit order by id

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/credit_limit/request/{RequestID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcreditlimitorderby_id []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/credit_limit/request/{RequestID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |
| RequestID  | string |             |

### Responses

#### Response body

| Name             | Type             | Description |
|------------------|------------------|-------------|
| CreditLimitOrder | CreditLimitOrder |             |

##### Objects

###### CreditLimitOrder

| Name             | Type   | Description |
|------------------|--------|-------------|
| RankAmount1      | string |             |
| NominalInterest1 | string |             |
| RankAmount2      | string |             |
| NominalInterest2 | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all customer adddresses

Returns a list containing up to 20 customer addresses.

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/addresses \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomeraddresses []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/addresses`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name      | Type      | Description |
|-----------|-----------|-------------|
| Addresses | []Address |             |

##### Objects

###### Address

| Name              | Type      | Description |
|-------------------|-----------|-------------|
| CustomerAddressID | string    |             |
| CustomerID        | string    |             |
| Line1             | string    |             |
| Line2             | string    |             |
| Line3             | string    |             |
| City              | string    |             |
| County            | string    |             |
| State             | string    |             |
| Postcode          | string    |             |
| CountryCode       | string    |             |
| Tags              | []string  |             |
| Status            | string    |             |
| InsertDate        | Timestamp |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer By ID

Returns the customer data based on it's id

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomerbycustomerid []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

##### Objects

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get Customer By CustomerNumber

Returns the customer data based on the customer number

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/customer-number/{CustomerNumber} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomerbycustomernumber []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/customer-number/{CustomerNumber}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CustomerID     | string |             |
| CustomerNumber | string |             |

### Responses

#### Response body

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

##### Objects

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all customer social media handles

Returns a list containing up to 20 customer social media handles

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/social_media_handles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomersocialmediahandles []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/social_media_handles`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name   | Type    | Description |
|--------|---------|-------------|
| Checks | []Check |             |

##### Objects

###### Check

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CustomerNumber | string    |             |
| Type           | string    |             |
| Handle         | string    |             |
| DateAdded      | Timestamp |             |
| DateActivated  | Timestamp |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers data for current user

Returns the current user customers data

```sh
curl -X GET \
	/v1/users/current/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomersforcurrentuser []}}

### HTTP Request

`GET /v1/users/current/customers`

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Customers | []Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers data by bank id

Returns the current user customers data by bank ID

```sh
curl -X GET \
	/v1/banks/{BankID}/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcustomersforcurrentuseratbank []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Customers | []Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Return customer firehose

Returns the list of customer firehose

```sh
curl -X GET \
	/v1/banks/{BankID}/firehose/customers \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getfirehosecustomer []}}

### HTTP Request

`GET /v1/banks/{BankID}/firehose/customers`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name      | Type       | Description |
|-----------|------------|-------------|
| Customers | []Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Get customers tax_residencies

Returns the user tax_residencies

```sh
curl -X GET \
	/v1/banks/{BankID}/customers/{CustomerID}/tax-residences \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet gettaxresidenceofcustomer []}}

### HTTP Request

`GET /v1/banks/{BankID}/customers/{CustomerID}/tax-residences`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Responses

#### Response body

| Name   | Type           | Description |
|--------|----------------|-------------|
| Result | []TaxResidence |             |

##### Objects

###### TaxResidence

| Name           | Type   | Description |
|----------------|--------|-------------|
| Domain         | string |             |
| TaxNumber      | string |             |
| TaxResidenceID | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update customer addresses.

Update customer addresses information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID} \
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
{{snippet updatecustomeraddress []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/addresses/{CustomerAddressID}`

### Query Parameters

| Name              | Type   | Description |
|-------------------|--------|-------------|
| BankID            | string |             |
| CustomerID        | string |             |
| CustomerAddressID | string |             |

### Body Parameters

| Name              | Type     | Description |
|-------------------|----------|-------------|
| BankID            | string   |             |
| CustomerID        | string   |             |
| CustomerAddressID | string   |             |
| Line1             | string   |             |
| Line2             | string   |             |
| Line3             | string   |             |
| City              | string   |             |
| County            | string   |             |
| State             | string   |             |
| Postcode          | string   |             |
| CountryCode       | string   |             |
| Tags              | []string |             |
| Status            | string   |             |

### Responses

#### Response body

| Name              | Type      | Description |
|-------------------|-----------|-------------|
| CustomerAddressID | string    |             |
| CustomerID        | string    |             |
| Line1             | string    |             |
| Line2             | string    |             |
| Line3             | string    |             |
| City              | string    |             |
| County            | string    |             |
| State             | string    |             |
| Postcode          | string    |             |
| CountryCode       | string    |             |
| Tags              | []string  |             |
| Status            | string    |             |
| InsertDate        | Timestamp |             |

##### Objects

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

## Update customer branch.

Update customer branch information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/branch \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"branch_id": "string"
	}'
```
{{snippet updatecustomerbranch []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/branch`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |
| BranchID   | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer credit limit.

Update customer credit limit information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/credit-limit \
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
{{snippet updatecustomercredit_limit []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/credit-limit`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| CustomerID  | string |             |
| CreditLimit | Amount |             |

##### Objects

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer credut rating and source.

Update customer credit rating and source information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/credit-rating-and-source \
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
{{snippet updatecustomercreditratingand_source []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/credit-rating-and-source`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name         | Type         | Description |
|--------------|--------------|-------------|
| BankID       | string       |             |
| CustomerID   | string       |             |
| CreditRating | CreditRating |             |

##### Objects

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer email.

Update customer email information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/email \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"email": "string"
	}'
```
{{snippet updatecustomeremail []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/email`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |
| Email      | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer identity.

Update customer identity information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/identity \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"legal_name": "string",
		"date_of_birth": "string",
		"title": "string"
	}'
```
{{snippet updatecustomeridentity []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/identity`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| CustomerID  | string |             |
| LegalName   | string |             |
| DateOfBirth | string |             |
| Title       | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer mobile number.

Update customer mobile number information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/mobile-number \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"mobile_phone_number": "string"
	}'
```
{{snippet updatecustomermobile_number []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/mobile-number`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name              | Type   | Description |
|-------------------|--------|-------------|
| BankID            | string |             |
| CustomerID        | string |             |
| MobilePhoneNumber | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer number.

Update customer number information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/number \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"customer_id": "string",
		"customer_number": "string"
	}'
```
{{snippet updatecustomernumber []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/number`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CustomerID     | string |             |
| CustomerNumber | string |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Update customer data.

Update customer data information.

```sh
curl -X PUT \
	/v1/banks/{BankID}/customers/{CustomerID}/data \
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
{{snippet updatecustomerother_data []}}

### HTTP Request

`PUT /v1/banks/{BankID}/customers/{CustomerID}/data`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| BankID     | string |             |
| CustomerID | string |             |

### Body Parameters

| Name                     | Type      | Description |
|--------------------------|-----------|-------------|
| BankID                   | string    |             |
| CustomerID               | string    |             |
| FaceImage                | FaceImage |             |
| RelationshipStatus       | string    |             |
| Dependants               | int32     |             |
| HighestEducationAttained | string    |             |
| EmploymentStatus         | string    |             |

##### Objects

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Customer | Customer |             |

##### Objects

###### Customer

| Name                     | Type         | Description |
|--------------------------|--------------|-------------|
| ID                       | string       |             |
| BankID                   | string       |             |
| CustomerNumber           | string       |             |
| LegalName                | string       |             |
| PhoneNumber              | string       |             |
| Email                    | string       |             |
| FaceImage                | FaceImage    |             |
| DateOfBirth              | Timestamp    |             |
| RelationshipStatus       | string       |             |
| CreditRating             | CreditRating |             |
| CreditLimit              | Amount       |             |
| HighestEducationAttained | string       |             |
| EmploymentStatus         | string       |             |
| KYCStatus                | bool         |             |
| LastOKDate               | Timestamp    |             |
| Title                    | string       |             |
| BranchID                 | string       |             |

###### FaceImage

| Name | Type      | Description |
|------|-----------|-------------|
| URL  | string    |             |
| Date | Timestamp |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
