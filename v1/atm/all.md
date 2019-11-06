# ATM API v1.0.0

Provides create and read operations on the ATM resource.

*
Host ``
EOL

*
Base Path ``

## Create an ATM

Creates a new ATM and returns its id.

```sh
curl -X POST \
	/v1/atms \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"name": "string",
		"3": {
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
{{snippet createat_m []}}

### HTTP Request

`POST /v1/atms`

### Body Parameters

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| BankID      | string   | BankID is the bank identifier that owned the ATM              |
| Name        | string   | BankID is the identifier of the bank associated with the ATM. |
| Address     | Address  | Address is the ATM's address.                                 |
| Location    | Location | Location is the ATM's longitude and latitude                  |
| Description | string   | Description is the ATM's description.                         |
| Metadata    | string   | Metadata is the ATM's metadata.                               |

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

| Name   | Type   | Description                                 |
|--------|--------|---------------------------------------------|
| ATM_ID | string | ATM_ID is the unique identifier of the ATM. |

Example:

```json
{
  "atm_id": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Transaction created successfully.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve an ATM

Retrieves information regarding a specific ATM, selected by the supplied ID.

```sh
curl -X GET \
	/v1/atms/{ATM_ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getat_m []}}

### HTTP Request

`GET /v1/atms/{ATM_ID}`

### Query Parameters

| Name   | Type   | Description                               |
|--------|--------|-------------------------------------------|
| ATM_ID | string | ATM_ID is a unique identifier of the ATM. |

### Responses

#### Response body

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| ID          | string   | ID is the unique identifier of an ATM.                        |
| BankID      | string   | BankID is the identifier of the bank associated with the ATM. |
| Name        | string   | Name is the name of ATM.                                      |
| Address     | Address  | Address is the ATM's address.                                 |
| Location    | Location | Location is the ATM longitude and latitude.                   |
| Description | string   | Description is the ATM's description.                         |
| Metadata    | string   | Metadata is the ATM's metadata.                               |

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

## Retrieve all available ATMs

Retrieves information regarding all the available ATMs.

```sh
curl -X GET \
	/v1/atms \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getat_ms []}}

### HTTP Request

`GET /v1/atms`

### Responses

#### Response body

| Name   | Type  | Description                     |
|--------|-------|---------------------------------|
| Result | []ATM | Result is the list of the ATMs. |

##### Objects

###### ATM

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| ID          | string   | ID is the unique identifier of an ATM.                        |
| BankID      | string   | BankID is the identifier of the bank associated with the ATM. |
| Name        | string   | Name is the name of ATM.                                      |
| Address     | Address  | Address is the ATM's address.                                 |
| Location    | Location | Location is the ATM longitude and latitude.                   |
| Description | string   | Description is the ATM's description.                         |
| Metadata    | string   | Metadata is the ATM's metadata.                               |

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
