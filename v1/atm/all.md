# ATM API v1.0.0

Provides create and read operations on the ATM resource.

* Host `https://`

* Base Path ``

## Create an ATM {#method-post-createatm}

Creates a new ATM and returns its id.

```sh
curl -X POST \
	https:///v1/atms \
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

### HTTP Request {#http-request-method-post-createatm}

`POST https:///v1/atms`

### Body Parameters {#body-parameters-method-post-createatm}

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| bank_id     | string   | BankID is the bank identifier that owned the ATM              |
| name        | string   | BankID is the identifier of the bank associated with the ATM. |
| 3           | Address  | Address is the ATM's address.                                 |
| location    | Location | Location is the ATM's longitude and latitude                  |
| description | string   | Description is the ATM's description.                         |
| metadata    | string   | Metadata is the ATM's metadata.                               |

##### Objects {#objects-CreateATMRequest}

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses {#responses-method-post-createatm}

#### Response body {#response-body-method-post-createatm}

| Name   | Type   | Description                                 |
|--------|--------|---------------------------------------------|
| atm_id | string | ATM_ID is the unique identifier of the ATM. |

Example:

```json
{
  "atm_id": "string"
}
```

#### Response codes {#response-codes-method-post-createatm}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Transaction created successfully.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve an ATM {#method-get-getatm}

Retrieves information regarding a specific ATM, selected by the supplied ID.

```sh
curl -X GET \
	https:///v1/atms/{ATM_ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getatm}

`GET https:///v1/atms/{ATM_ID}`

### Query Parameters {#query-parameters-method-get-getatm}

| Name   | Type   | Description                               |
|--------|--------|-------------------------------------------|
| atm_id | string | ATM_ID is a unique identifier of the ATM. |

### Responses {#responses-method-get-getatm}

#### Response body {#response-body-method-get-getatm}

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| id          | string   | ID is the unique identifier of an ATM.                        |
| bank_id     | string   | BankID is the identifier of the bank associated with the ATM. |
| name        | string   | Name is the name of ATM.                                      |
| address     | Address  | Address is the ATM's address.                                 |
| location    | Location | Location is the ATM longitude and latitude.                   |
| description | string   | Description is the ATM's description.                         |
| metadata    | string   | Metadata is the ATM's metadata.                               |

##### Objects {#objects-ATM}

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

#### Response codes {#response-codes-method-get-getatm}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available ATMs {#method-get-getatms}

Retrieves information regarding all the available ATMs.

```sh
curl -X GET \
	https:///v1/atms \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getatms}

`GET https:///v1/atms`

### Responses {#responses-method-get-getatms}

#### Response body {#response-body-method-get-getatms}

| Name   | Type   | Description                     |
|--------|--------|---------------------------------|
| result | \[]ATM | Result is the list of the ATMs. |

##### Objects {#objects-GetATMsResponse}

###### ATM

| Name        | Type     | Description                                                   |
|-------------|----------|---------------------------------------------------------------|
| id          | string   | ID is the unique identifier of an ATM.                        |
| bank_id     | string   | BankID is the identifier of the bank associated with the ATM. |
| name        | string   | Name is the name of ATM.                                      |
| address     | Address  | Address is the ATM's address.                                 |
| location    | Location | Location is the ATM longitude and latitude.                   |
| description | string   | Description is the ATM's description.                         |
| metadata    | string   | Metadata is the ATM's metadata.                               |

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

#### Response codes {#response-codes-method-get-getatms}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
