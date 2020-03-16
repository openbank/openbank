Account Application API v1.0.0
==============================

Provides create and read operations on the account application resource.

* Host `https://`

* Base Path ``

Create an account application {#method-post-createaccountapplication}
---------------------------------------------------------------------

Creates a new account application

```sh
curl -X POST \
	https:///v1/account-applications \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"product_code": "string",
		"user_id": "string",
		"profile_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createaccountapplication}

`POST https:///v1/account-applications`

### Body Parameters {#body-parameters-method-post-createaccountapplication}

| Name         | Type   | Description |
|--------------|--------|-------------|
| product_code | string |             |
| user_id      | string |             |
| profile_id   | string |             |

### Responses {#responses-method-post-createaccountapplication}

#### Response body {#response-body-method-post-createaccountapplication}

| Name                | Type                     | Description |
|---------------------|--------------------------|-------------|
| id                  | string                   |             |
| product_code        | string                   |             |
| user                | User                     |             |
| customer_profile    | Profile                  |             |
| date_of_application | Timestamp                |             |
| status              | AccountApplicationStatus |             |

##### Objects {#objects-CreateAccountApplicationResponse}

###### User

| Name     | Type   | Description                   |
|----------|--------|-------------------------------|
| id       | string | User identifier used to login |
| email    | string | Email of the user             |
| username | string | Display name of the user      |

###### Profile

| Name                       | Type         | Description                                                |
|----------------------------|--------------|------------------------------------------------------------|
| profile_id                 | string       | ProfileID is the unique identifier of a profile.           |
| fullname                   | string       | Full name                                                  |
| username                   | string       | User name                                                  |
| birthdate                  | string       | Birth date                                                 |
| language                   | string       | Language code used                                         |
| country                    | string       | User country code (VN, US, ID, SG, ...).                   |
| email                      | string       | User email address                                         |
| email_verified             | bool         | True if email is verified, otherwise False                 |
| mobile                     | string       | Mobile number                                              |
| photo                      | string       | User profile photo url                                     |
| title                      | string       | Title                                                      |
| permanent_address          | Address      | Permanent address                                          |
| contact_address            | Address      | Contact address                                            |
| contact_number             | string       | Contact number                                             |
| profile_number             | string       | profile number                                             |
| face_image_url             | string       | Face image of the customer                                 |
| face_image_date            | string       | Date when the face image was added/updated                 |
| relationship_status        | string       | RelationshipStatus. Ex: Single                             |
| dependents                 | int32        | Number of dependents                                       |
| dob_of_dependents          | \[]Timestamp | Date of birth of dependents                                |
| credit_rating              | CreditRating | Credit rating                                              |
| credit_limit               | Amount       | Credit Limit                                               |
| highest_education_attained | string       | Highest education such as bachelor, masters etc            |
| employment_status          | string       | Current employment status                                  |
| kyc_status                 | bool         | Know Your Customer status                                  |
| branchId                   | string       | Branch Identifier                                          |
| nameSuffix                 | string       | Name suffix                                                |
| first_name                 | string       | FirstName of the person                                    |
| middle_name                | string       | MiddleName or middle names (space separated) of the person |
| last_name                  | string       | LastName or last names (space separated) of the person     |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

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

##### Enums {#enums-CreateAccountApplicationResponse}

###### AccountApplicationStatus

Entity type defines the type of account application

| Value                    | Description                     |
|--------------------------|---------------------------------|
| UnknownApplicationStatus |                                 |
| Requested                | Account application is created  |
| Accepted                 | Account application is accepted |

Example:

```json
{
  "id": "string",
  "product_code": "string",
  "user": {
    "id": "string",
    "email": "string",
    "username": "string"
  },
  "customer_profile": {
    "profile_id": "string",
    "fullname": "string",
    "username": "string",
    "birthdate": "string",
    "language": "string",
    "country": "string",
    "email": "string",
    "email_verified": "bool",
    "mobile": "string",
    "photo": "string",
    "title": "string",
    "permanent_address": {
      "country_name": "string",
      "city_name": "string",
      "state": "string",
      "line_1": "string",
      "postal_code": "string"
    },
    "contact_address": {
      "country_name": "string",
      "city_name": "string",
      "state": "string",
      "line_1": "string",
      "postal_code": "string"
    },
    "contact_number": "string",
    "profile_number": "string",
    "face_image_url": "string",
    "face_image_date": "string",
    "relationship_status": "string",
    "dependents": "int32",
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
    "kyc_status": "bool",
    "branchId": "string",
    "nameSuffix": "string",
    "first_name": "string",
    "middle_name": "string",
    "last_name": "string"
  },
  "date_of_application": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "AccountApplicationStatus"
}
```

#### Response codes {#response-codes-method-post-createaccountapplication}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account application created successfully.                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve an account application {#method-get-getaccountapplication}
-------------------------------------------------------------------

Retrieves all data from an account application selected by the supplied account_application_id.

```sh
curl -X GET \
	https:///v1/account-applications/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getaccountapplication}

`GET https:///v1/account-applications/{ID}`

### Query Parameters {#query-parameters-method-get-getaccountapplication}

| Name | Type   | Description |
|------|--------|-------------|
| id   | string |             |

### Responses {#responses-method-get-getaccountapplication}

#### Response body {#response-body-method-get-getaccountapplication}

| Name                | Type                     | Description |
|---------------------|--------------------------|-------------|
| id                  | string                   |             |
| product_code        | string                   |             |
| user_id             | string                   |             |
| customer_profile_id | string                   |             |
| date_of_application | Timestamp                |             |
| status              | AccountApplicationStatus |             |

##### Objects {#objects-AccountApplication}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

##### Enums {#enums-AccountApplication}

###### AccountApplicationStatus

Entity type defines the type of account application

| Value                    | Description                     |
|--------------------------|---------------------------------|
| UnknownApplicationStatus |                                 |
| Requested                | Account application is created  |
| Accepted                 | Account application is accepted |

Example:

```json
{
  "id": "string",
  "product_code": "string",
  "user_id": "string",
  "customer_profile_id": "string",
  "date_of_application": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "AccountApplicationStatus"
}
```

#### Response codes {#response-codes-method-get-getaccountapplication}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

List all accounts {#method-get-getaccountapplications}
------------------------------------------------------

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/account-applications \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getaccountapplications}

`GET https:///v1/account-applications`

### Responses {#responses-method-get-getaccountapplications}

#### Response body {#response-body-method-get-getaccountapplications}

| Name   | Type                  | Description                           |
|--------|-----------------------|---------------------------------------|
| result | \[]AccountApplication | Result is the paginated query result. |

##### Objects {#objects-GetAccountApplicationsResponse}

###### AccountApplication

| Name                | Type                     | Description |
|---------------------|--------------------------|-------------|
| id                  | string                   |             |
| product_code        | string                   |             |
| user_id             | string                   |             |
| customer_profile_id | string                   |             |
| date_of_application | Timestamp                |             |
| status              | AccountApplicationStatus |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

##### Enums {#enums-GetAccountApplicationsResponse}

###### AccountApplicationStatus

Entity type defines the type of account application

| Value                    | Description                     |
|--------------------------|---------------------------------|
| UnknownApplicationStatus |                                 |
| Requested                | Account application is created  |
| Accepted                 | Account application is accepted |

Example:

```json
{
  "result": [
    {
      "id": "string",
      "product_code": "string",
      "user_id": "string",
      "customer_profile_id": "string",
      "date_of_application": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "status": "AccountApplicationStatus"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getaccountapplications}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update an account application {#method-put-updateaccountapplicationstatus}
--------------------------------------------------------------------------

Updates an account application with the given status

```sh
curl -X PUT \
	https:///v1/account-applications/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"id": "string",
		"status": "AccountApplicationStatus"
	}'
```

### HTTP Request {#http-request-method-put-updateaccountapplicationstatus}

`PUT https:///v1/account-applications/{ID}`

### Query Parameters {#query-parameters-method-put-updateaccountapplicationstatus}

| Name | Type   | Description                                                       |
|------|--------|-------------------------------------------------------------------|
| id   | string | ID is the unique identifier of the account application to update. |

### Body Parameters {#body-parameters-method-put-updateaccountapplicationstatus}

| Name   | Type                     | Description                                                       |
|--------|--------------------------|-------------------------------------------------------------------|
| id     | string                   | ID is the unique identifier of the account application to update. |
| status | AccountApplicationStatus | Status of the account application                                 |

##### Enums {#enums-UpdateAccountApplicationStatusRequest}

###### AccountApplicationStatus

Entity type defines the type of account application

| Value                    | Description                     |
|--------------------------|---------------------------------|
| UnknownApplicationStatus |                                 |
| Requested                | Account application is created  |
| Accepted                 | Account application is accepted |

### Responses {#responses-method-put-updateaccountapplicationstatus}

#### Response body {#response-body-method-put-updateaccountapplicationstatus}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-put-updateaccountapplicationstatus}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Account Application status updated successfully.                                       |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
