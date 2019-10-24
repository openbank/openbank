# Customer API v1.0.0

Provides create and read operations on the customer resource.

* Host ``

* Base Path ``

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
