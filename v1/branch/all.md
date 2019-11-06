# Branch API v1.0.0

Provides create and read operations on the branch resource.

*
Host ``
EOL

*
Base Path ``

## Create a branch

Creates a new branch and returns its id.

```sh
curl -X POST \
	/v1/branches \
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
### HTTP Request

`POST /v1/branches`

### Body Parameters

| Name        | Type     | Description                                                         |
|-------------|----------|---------------------------------------------------------------------|
| BankID      | string   | BankID is an identifier for the bank the branch is associated with. |
| Name        | string   | Name is the branch name.                                            |
| PhoneNumber | string   | PhoneNumber is the branch's phone number.                           |
| Address     | Address  | Address is the branch's address.                                    |
| Location    | Location | Location is the branch's longitude and latitude.                    |
| Description | string   | Description is the branch's description.                            |
| Metadata    | string   | Metadata is the branch's metadata.                                  |

##### Objects

###### Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response body

| Name     | Type   | Description                               |
|----------|--------|-------------------------------------------|
| BranchID | string | BranchID is the branch unique identifier. |

Example:

```json
{
  "branch_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Branch created successfully.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a branch

Permanently delete a branch.

```sh
curl -X DELETE \
	/v1/branches/{BranchID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/branches/{BranchID}`

### Query Parameters

| Name     | Type   | Description                               |
|----------|--------|-------------------------------------------|
| BranchID | string | BranchID is the branch unique identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Branch successfully deleted.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve branch information

Retrieve information about the branch specified by the ID

```sh
curl -X GET \
	/v1/branches/{BranchID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/branches/{BranchID}`

### Query Parameters

| Name     | Type   | Description                               |
|----------|--------|-------------------------------------------|
| BranchID | string | BranchID is the branch unique identifier. |

### Responses

#### Response body

| Name        | Type     | Description                                                         |
|-------------|----------|---------------------------------------------------------------------|
| ID          | string   | ID is the unique identifier of the branch.                          |
| BankID      | string   | BankID is an identifier for the bank the branch is associated with. |
| Name        | string   | Name is the branch name.                                            |
| PhoneNumber | string   | PhoneNumber is the branch's phone number.                           |
| Address     | Address  | Address is the branch's address.                                    |
| Location    | Location | Location is the branch's longitude and latitude.                    |
| Description | string   | Description is the branch's description.                            |
| Metadata    | string   | Metadata is the branch's metadata.                                  |

##### Objects

###### Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

## Retrieve all available branches

Retrieve information regarding all available branches.

```sh
curl -X GET \
	/v1/branches \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/branches`

### Responses

#### Response body

| Name   | Type     | Description                       |
|--------|----------|-----------------------------------|
| Result | []Branch | Result is the list of the branch. |

##### Objects

###### Branch

| Name        | Type     | Description                                                         |
|-------------|----------|---------------------------------------------------------------------|
| ID          | string   | ID is the unique identifier of the branch.                          |
| BankID      | string   | BankID is an identifier for the bank the branch is associated with. |
| Name        | string   | Name is the branch name.                                            |
| PhoneNumber | string   | PhoneNumber is the branch's phone number.                           |
| Address     | Address  | Address is the branch's address.                                    |
| Location    | Location | Location is the branch's longitude and latitude.                    |
| Description | string   | Description is the branch's description.                            |
| Metadata    | string   | Metadata is the branch's metadata.                                  |

###### Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

## Update a branch

Updates a branch's information

```sh
curl -X PUT \
	/v1/branches/{BranchID} \
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
### HTTP Request

`PUT /v1/branches/{BranchID}`

### Query Parameters

| Name     | Type   | Description                               |
|----------|--------|-------------------------------------------|
| BranchID | string | BranchID is the branch unique identifier. |

### Body Parameters

| Name        | Type     | Description                                    |
|-------------|----------|------------------------------------------------|
| BranchID    | string   | BranchID is the branch unique identifier.      |
| Name        | string   | Name is the branch name.                       |
| PhoneNumber | string   | PhoneNumber is the branch phone number.        |
| Address     | Address  | Address is the branch address details.         |
| Location    | Location | Location is the branch longitude and latitude. |
| Description | string   | Description is the branch description.         |
| Metadata    | string   | Metadata is the branch metadata.               |

##### Objects

###### Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Branch successfully updated.                                                           |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
