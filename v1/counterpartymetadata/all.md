# Counterparty Metadata API v1.0.0

Provides the access and availability to the counterparty metadata API.

* Host ``

* Base Path ``

## Create a corporate_location

Creates a new corporate_location and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"corporate_location": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createcorporatelocation []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name              | Type              | Description                                     |
|-------------------|-------------------|-------------------------------------------------|
| CorporateLocation | CorporateLocation | account.                                        |
| BankID            | string            | BankID is the bank identifier.                  |
| AccountID         | string            | AccountID is the account identifier.            |
| OtherAccountID    | string            | OtherAccountID is the other_account identifier. |

##### Objects

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a image_url

Creates a new image_url and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"image_url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createimageurl []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type     | Description                                     |
|----------------|----------|-------------------------------------------------|
| ImageURL       | ImageURL | account.                                        |
| BankID         | string   | BankID is the bank identifier.                  |
| AccountID      | string   | AccountID is the account identifier.            |
| OtherAccountID | string   | OtherAccountID is the other_account identifier. |

##### Objects

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a more_info

Creates a new more_info and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"more_info": {
			"more_info": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createmoreinfo []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type     | Description                                     |
|----------------|----------|-------------------------------------------------|
| MoreInfo       | MoreInfo | account.                                        |
| BankID         | string   | BankID is the bank identifier.                  |
| AccountID      | string   | AccountID is the account identifier.            |
| OtherAccountID | string   | OtherAccountID is the other_account identifier. |

##### Objects

###### MoreInfo

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

### Responses

#### Response body

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

Example:

```json
{
  "more_info": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a open## corporates## url

Creates a new opencorporatesurl and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"open_corporates_url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createopencorporates_url []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name              | Type              | Description                                     |
|-------------------|-------------------|-------------------------------------------------|
| OpenCorporatesURL | OpenCorporatesURL | url metadata made on a other_account.           |
| BankID            | string            | BankID is the bank identifier.                  |
| AccountID         | string            | AccountID is the account identifier.            |
| OtherAccountID    | string            | OtherAccountID is the other_account identifier. |

##### Objects

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a physical_location

Creates a new physical_location and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"physical_location": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createphysicallocation []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name             | Type             | Description                                     |
|------------------|------------------|-------------------------------------------------|
| PhysicalLocation | PhysicalLocation | account.                                        |
| BankID           | string           | BankID is the bank identifier.                  |
| AccountID        | string           | AccountID is the account identifier.            |
| OtherAccountID   | string           | OtherAccountID is the other_account identifier. |

##### Objects

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a private_alias

Creates a new private_alias and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"alias": {
			"alias": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createprivatealias []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type        | Description                                     |
|----------------|-------------|-------------------------------------------------|
| PrivateAlias   | PublicAlias | account.                                        |
| BankID         | string      | BankID is the bank identifier.                  |
| AccountID      | string      | AccountID is the account identifier.            |
| OtherAccountID | string      | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a public_alias

Creates a new public_alias and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"alias": {
			"alias": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet createpublicalias []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type        | Description                                     |
|----------------|-------------|-------------------------------------------------|
| PublicAlias    | PublicAlias | account.                                        |
| BankID         | string      | BankID is the bank identifier.                  |
| AccountID      | string      | AccountID is the account identifier.            |
| OtherAccountID | string      | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a url

Creates a new url and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type   | Description                                      |
|----------------|--------|--------------------------------------------------|
| URL            | URL    | URL is the url metadata made on a other_account. |
| BankID         | string | BankID is the bank identifier.                   |
| AccountID      | string | AccountID is the account identifier.             |
| OtherAccountID | string | OtherAccountID is the other_account identifier.  |

##### Objects

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a corporate location

Permanently delete a corporate location.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletecorporatelocation []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | CorporateLocation successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a image url

Permanently delete a image url.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteimageurl []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ImageURL successfully deleted.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a more info

Permanently delete a more info.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletemoreinfo []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | MoreInfo successfully deleted.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a open corporates url

Permanently delete a open corporates url.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteopencorporates_url []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | OpenCorporatesURL successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a physical location

Permanently delete a physical location.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletephysicallocation []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PhysicalLocation successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a private alias

Permanently delete a private alias.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteprivatealias []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PrivateAlias successfully deleted.                                                     |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a public alias

Permanently delete a public alias.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deletepublicalias []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PublicAlias successfully deleted.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a url

Permanently delete a url.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | URL successfully deleted.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve corporate location information

Retrieve information about the corporate location specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcorporatelocation []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
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

## Retrieve all available corporate locations

Retrieve information regarding all available corporate locations.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_locations \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcorporatelocations []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_locations`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name               | Type                | Description                                            |
|--------------------|---------------------|--------------------------------------------------------|
| CorporateLocations | []CorporateLocation | CorporateLocations is the list of corporate_locations. |
| BankID             | string              | BankID is the bank identifier.                         |
| AccountID          | string              | AccountID is the account identifier.                   |
| OtherAccountID     | string              | OtherAccountID is the other_account identifier.        |

##### Objects

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "corporate_locations": [
    {
      "location": {
        "latitude": "double",
        "longitude": "double"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve image url information

Retrieve information about the image url specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getimageurl []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
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

## Retrieve all available image urls

Retrieve information regarding all available image urls.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getimageurl_s []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_urls`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name           | Type       | Description                                     |
|----------------|------------|-------------------------------------------------|
| ImageURLs      | []ImageURL | ImageURLs is the list of image_urls.            |
| BankID         | string     | BankID is the bank identifier.                  |
| AccountID      | string     | AccountID is the account identifier.            |
| OtherAccountID | string     | OtherAccountID is the other_account identifier. |

##### Objects

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "image_urls": [
    {
      "url": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve more info information

Retrieve information about the more info specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getmoreinfo []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

Example:

```json
{
  "more_info": "string"
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

## Retrieve all available more infos

Retrieve information regarding all available more infos.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_infos \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getmoreinfos []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_infos`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name           | Type       | Description                                     |
|----------------|------------|-------------------------------------------------|
| MoreInfos      | []MoreInfo | MoreInfos is the list of more_infos.            |
| BankID         | string     | BankID is the bank identifier.                  |
| AccountID      | string     | AccountID is the account identifier.            |
| OtherAccountID | string     | OtherAccountID is the other_account identifier. |

##### Objects

###### MoreInfo

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

Example:

```json
{
  "more_infos": [
    {
      "more_info": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve open corporates url information

Retrieve information about the open corporates url specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getopencorporates_url []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
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

## Retrieve all available open corporates urls

Retrieve information regarding all available open corporates urls.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getopencorporatesurls []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_urls`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name               | Type                | Description                                     |
|--------------------|---------------------|-------------------------------------------------|
| OpenCorporatesURLs | []OpenCorporatesURL | urls.                                           |
| BankID             | string              | BankID is the bank identifier.                  |
| AccountID          | string              | AccountID is the account identifier.            |
| OtherAccountID     | string              | OtherAccountID is the other_account identifier. |

##### Objects

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "open_corporates_urls": [
    {
      "url": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve other account metadata information

Retrieve information about the other account metadata specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getotheraccount_metadata []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name              | Type              | Description |
|-------------------|-------------------|-------------|
| PublicAlias       | string            |             |
| PrivateAlias      | string            |             |
| MoreInfo          | string            |             |
| URL               | string            |             |
| ImageURL          | string            |             |
| OpenCorporatesURL | string            |             |
| CorporateLocation | CorporateLocation |             |
| PhysicalLocation  | PhysicalLocation  |             |

##### Objects

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "public_alias": "string",
  "private_alias": "string",
  "more_info": "string",
  "url": "string",
  "image_url": "string",
  "open_corporates_url": "string",
  "corporate_location": {
    "location": {
      "latitude": "double",
      "longitude": "double"
    }
  },
  "physical_location": {
    "location": {
      "latitude": "double",
      "longitude": "double"
    }
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

## Retrieve physical location information

Retrieve information about the physical location specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getphysicallocation []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
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

## Retrieve all available physical locations

Retrieve information regarding all available physical locations.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_locations \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getphysicallocations []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_locations`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name              | Type               | Description                                          |
|-------------------|--------------------|------------------------------------------------------|
| PhysicalLocations | []PhysicalLocation | PhysicalLocations is the list of physical_locations. |
| BankID            | string             | BankID is the bank identifier.                       |
| AccountID         | string             | AccountID is the account identifier.                 |
| OtherAccountID    | string             | OtherAccountID is the other_account identifier.      |

##### Objects

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "physical_locations": [
    {
      "location": {
        "latitude": "double",
        "longitude": "double"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve private alias information

Retrieve information about the private alias specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getprivatealias []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
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

## Retrieve all available private aliases

Retrieve information regarding all available private aliases.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_aliases \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getprivatealiases []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_aliases`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name           | Type          | Description                                     |
|----------------|---------------|-------------------------------------------------|
| PrivateAliases | []PublicAlias | PrivateAliases is the list of private_aliases.  |
| BankID         | string        | BankID is the bank identifier.                  |
| AccountID      | string        | AccountID is the account identifier.            |
| OtherAccountID | string        | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "aliases": [
    {
      "alias": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve public alias information

Retrieve information about the public alias specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getpublicalias []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
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

## Retrieve all available public aliases

Retrieve information regarding all available public aliases.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_aliases \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getpublicaliases []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_aliases`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name           | Type          | Description                                     |
|----------------|---------------|-------------------------------------------------|
| PublicAliases  | []PublicAlias | PublicAliases is the list of public_aliases.    |
| BankID         | string        | BankID is the bank identifier.                  |
| AccountID      | string        | AccountID is the account identifier.            |
| OtherAccountID | string        | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "aliases": [
    {
      "alias": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Retrieve url information

Retrieve information about the url specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
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

## Retrieve all available urls

Retrieve information regarding all available urls.

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet geturls []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/urls`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Responses

#### Response body

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| URLs           | []URL  | URLs is the list of urls.                       |
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

##### Objects

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "urls": [
    {
      "url": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "other_account_id": "string"
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

## Update a corporate location

Updates a corporate location's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"corporate_location": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updatecorporatelocation []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name              | Type              | Description                                     |
|-------------------|-------------------|-------------------------------------------------|
| CorporateLocation | CorporateLocation | account.                                        |
| BankID            | string            | BankID is the bank identifier.                  |
| AccountID         | string            | AccountID is the account identifier.            |
| OtherAccountID    | string            | OtherAccountID is the other_account identifier. |

##### Objects

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | CorporateLocation successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a image url

Updates a image url's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"image_url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updateimageurl []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type     | Description                                     |
|----------------|----------|-------------------------------------------------|
| ImageURL       | ImageURL | account.                                        |
| BankID         | string   | BankID is the bank identifier.                  |
| AccountID      | string   | AccountID is the account identifier.            |
| OtherAccountID | string   | OtherAccountID is the other_account identifier. |

##### Objects

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ImageURL successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a more info

Updates a more info's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"more_info": {
			"more_info": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updatemoreinfo []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type     | Description                                     |
|----------------|----------|-------------------------------------------------|
| MoreInfo       | MoreInfo | account.                                        |
| BankID         | string   | BankID is the bank identifier.                  |
| AccountID      | string   | AccountID is the account identifier.            |
| OtherAccountID | string   | OtherAccountID is the other_account identifier. |

##### Objects

###### MoreInfo

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

### Responses

#### Response body

| Name     | Type   | Description |
|----------|--------|-------------|
| MoreInfo | string |             |

Example:

```json
{
  "more_info": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | MoreInfo successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a open corporates url

Updates a open corporates url's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"open_corporates_url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updateopencorporates_url []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name              | Type              | Description                                     |
|-------------------|-------------------|-------------------------------------------------|
| OpenCorporatesURL | OpenCorporatesURL | url metadata made on a other_account.           |
| BankID            | string            | BankID is the bank identifier.                  |
| AccountID         | string            | AccountID is the account identifier.            |
| OtherAccountID    | string            | OtherAccountID is the other_account identifier. |

##### Objects

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | OpenCorporatesURL successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a physical location

Updates a physical location's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"physical_location": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updatephysicallocation []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name             | Type             | Description                                     |
|------------------|------------------|-------------------------------------------------|
| PhysicalLocation | PhysicalLocation | account.                                        |
| BankID           | string           | BankID is the bank identifier.                  |
| AccountID        | string           | AccountID is the account identifier.            |
| OtherAccountID   | string           | OtherAccountID is the other_account identifier. |

##### Objects

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses

#### Response body

| Name     | Type     | Description |
|----------|----------|-------------|
| Location | Location |             |

##### Objects

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PhysicalLocation successfully updated.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a private alias

Updates a private alias's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"alias": {
			"alias": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updateprivatealias []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type        | Description                                     |
|----------------|-------------|-------------------------------------------------|
| PrivateAlias   | PublicAlias | account.                                        |
| BankID         | string      | BankID is the bank identifier.                  |
| AccountID      | string      | AccountID is the account identifier.            |
| OtherAccountID | string      | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PrivateAlias successfully updated.                                                     |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a public alias

Updates a public alias's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"alias": {
			"alias": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
{{snippet updatepublicalias []}}

### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type        | Description                                     |
|----------------|-------------|-------------------------------------------------|
| PublicAlias    | PublicAlias | account.                                        |
| BankID         | string      | BankID is the bank identifier.                  |
| AccountID      | string      | AccountID is the account identifier.            |
| OtherAccountID | string      | OtherAccountID is the other_account identifier. |

##### Objects

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

### Responses

#### Response body

| Name  | Type   | Description |
|-------|--------|-------------|
| Alias | string |             |

Example:

```json
{
  "alias": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PublicAlias successfully updated.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a url

Updates a url's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"url": {
			"url": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"other_account_id": "string"
	}'
```
### HTTP Request

`PUT /v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters

| Name           | Type   | Description                                     |
|----------------|--------|-------------------------------------------------|
| BankID         | string | BankID is the bank identifier.                  |
| AccountID      | string | AccountID is the account identifier.            |
| OtherAccountID | string | OtherAccountID is the other_account identifier. |

### Body Parameters

| Name           | Type   | Description                                      |
|----------------|--------|--------------------------------------------------|
| URL            | URL    | URL is the url metadata made on a other_account. |
| BankID         | string | BankID is the bank identifier.                   |
| AccountID      | string | AccountID is the account identifier.             |
| OtherAccountID | string | OtherAccountID is the other_account identifier.  |

##### Objects

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

### Responses

#### Response body

| Name | Type   | Description |
|------|--------|-------------|
| URL  | string |             |

Example:

```json
{
  "url": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | URL successfully updated.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
