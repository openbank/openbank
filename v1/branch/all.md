Branch API v1.0.0
=================

Provides create and read operations on the branch resource.

* Host `https://`

* Base Path ``

Create a branch {#method-post-createbranch}
-------------------------------------------

Creates a new branch and returns its id.

```sh
curl -X POST \
	https:///v1/branches \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"name": "string",
		"phone_number": "string",
		"address": {
			"country_name": "string",
			"city_name": "string",
			"state": "string",
			"line_1": "string",
			"postal_code": "string"
		},
		"location": {
			"latitude": "double",
			"longitude": "double"
		},
		"description": "string",
		"metadata": "string"
	}'
```

### HTTP Request {#http-request-method-post-createbranch}

`POST https:///v1/branches`

### Body Parameters {#body-parameters-method-post-createbranch}

| Name         | Type     | Description                                                         |
|--------------|----------|---------------------------------------------------------------------|
| bank_id      | string   | BankID is an identifier for the bank the branch is associated with. |
| name         | string   | Name is the branch name.                                            |
| phone_number | string   | PhoneNumber is the branch's phone number.                           |
| address      | Address  | Address is the branch's address.                                    |
| location     | Location | Location is the branch's longitude and latitude.                    |
| description  | string   | Description is the branch's description.                            |
| metadata     | string   | Metadata is the branch's metadata.                                  |

##### Objects {#objects-CreateBranchRequest}

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

### Responses {#responses-method-post-createbranch}

#### Response body {#response-body-method-post-createbranch}

| Name      | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| branch_id | string | BranchID is the branch unique identifier. |

Example:

```json
{
  "branch_id": "string"
}
```

#### Response codes {#response-codes-method-post-createbranch}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Branch created successfully.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a branch {#method-delete-deletebranch}
---------------------------------------------

Permanently delete a branch.

```sh
curl -X DELETE \
	https:///v1/branches/{BranchID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletebranch}

`DELETE https:///v1/branches/{BranchID}`

### Query Parameters {#query-parameters-method-delete-deletebranch}

| Name      | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| branch_id | string | BranchID is the branch unique identifier. |

### Responses {#responses-method-delete-deletebranch}

#### Response codes {#response-codes-method-delete-deletebranch}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Branch successfully deleted.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve branch information {#method-get-getbranch}
---------------------------------------------------

Retrieve information about the branch specified by the ID

```sh
curl -X GET \
	https:///v1/branches/{BranchID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getbranch}

`GET https:///v1/branches/{BranchID}`

### Query Parameters {#query-parameters-method-get-getbranch}

| Name      | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| branch_id | string | BranchID is the branch unique identifier. |

### Responses {#responses-method-get-getbranch}

#### Response body {#response-body-method-get-getbranch}

| Name         | Type     | Description                                                         |
|--------------|----------|---------------------------------------------------------------------|
| id           | string   | ID is the unique identifier of the branch.                          |
| bank_id      | string   | BankID is an identifier for the bank the branch is associated with. |
| name         | string   | Name is the branch name.                                            |
| phone_number | string   | PhoneNumber is the branch's phone number.                           |
| address      | Address  | Address is the branch's address.                                    |
| location     | Location | Location is the branch's longitude and latitude.                    |
| description  | string   | Description is the branch's description.                            |
| metadata     | string   | Metadata is the branch's metadata.                                  |

##### Objects {#objects-Branch}

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "id": "string",
  "bank_id": "string",
  "name": "string",
  "phone_number": "string",
  "address": {
    "country_name": "string",
    "city_name": "string",
    "state": "string",
    "line_1": "string",
    "postal_code": "string"
  },
  "location": {
    "latitude": "double",
    "longitude": "double"
  },
  "description": "string",
  "metadata": "string"
}
```

#### Response codes {#response-codes-method-get-getbranch}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available branches {#method-get-getbranches}
---------------------------------------------------------

Retrieve information regarding all available branches.

```sh
curl -X GET \
	https:///v1/branches \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getbranches}

`GET https:///v1/branches`

### Responses {#responses-method-get-getbranches}

#### Response body {#response-body-method-get-getbranches}

| Name   | Type      | Description                       |
|--------|-----------|-----------------------------------|
| result | \[]Branch | Result is the list of the branch. |

##### Objects {#objects-GetBranchesResponse}

###### Branch

| Name         | Type     | Description                                                         |
|--------------|----------|---------------------------------------------------------------------|
| id           | string   | ID is the unique identifier of the branch.                          |
| bank_id      | string   | BankID is an identifier for the bank the branch is associated with. |
| name         | string   | Name is the branch name.                                            |
| phone_number | string   | PhoneNumber is the branch's phone number.                           |
| address      | Address  | Address is the branch's address.                                    |
| location     | Location | Location is the branch's longitude and latitude.                    |
| description  | string   | Description is the branch's description.                            |
| metadata     | string   | Metadata is the branch's metadata.                                  |

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "result": [
    {
      "id": "string",
      "bank_id": "string",
      "name": "string",
      "phone_number": "string",
      "address": {
        "country_name": "string",
        "city_name": "string",
        "state": "string",
        "line_1": "string",
        "postal_code": "string"
      },
      "location": {
        "latitude": "double",
        "longitude": "double"
      },
      "description": "string",
      "metadata": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getbranches}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a branch {#method-put-updatebranch}
------------------------------------------

Updates a branch's information

```sh
curl -X PUT \
	https:///v1/branches/{BranchID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"branch_id": "string",
		"name": "string",
		"phone_number": "string",
		"address": {
			"country_name": "string",
			"city_name": "string",
			"state": "string",
			"line_1": "string",
			"postal_code": "string"
		},
		"location": {
			"latitude": "double",
			"longitude": "double"
		},
		"description": "string",
		"metadata": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatebranch}

`PUT https:///v1/branches/{BranchID}`

### Query Parameters {#query-parameters-method-put-updatebranch}

| Name      | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| branch_id | string | BranchID is the branch unique identifier. |

### Body Parameters {#body-parameters-method-put-updatebranch}

| Name         | Type     | Description                                    |
|--------------|----------|------------------------------------------------|
| branch_id    | string   | BranchID is the branch unique identifier.      |
| name         | string   | Name is the branch name.                       |
| phone_number | string   | PhoneNumber is the branch phone number.        |
| address      | Address  | Address is the branch address details.         |
| location     | Location | Location is the branch longitude and latitude. |
| description  | string   | Description is the branch description.         |
| metadata     | string   | Metadata is the branch metadata.               |

##### Objects {#objects-UpdateBranchRequest}

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

### Responses {#responses-method-put-updatebranch}

#### Response codes {#response-codes-method-put-updatebranch}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Branch successfully updated.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
