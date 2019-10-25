# Customer API v1.0.0

Provides create and read operations on the customer resource.

* Host ``

* Base Path ``

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
