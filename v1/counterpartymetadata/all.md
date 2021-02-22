# Counterparty Metadata API v1.0.0

Provides the access and availability to the counterparty metadata API.

* Host `https://`

* Base Path ``

## Create a corporate_location {#method-post-createcorporatelocation}

Creates a new corporate_location and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
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

### HTTP Request {#http-request-method-post-createcorporatelocation}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters {#query-parameters-method-post-createcorporatelocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createcorporatelocation}

| Name               | Type              | Description                                                                   |
|--------------------|-------------------|-------------------------------------------------------------------------------|
| corporate_location | CorporateLocation | CorporateLocation is the corporate_location metadata made on a other_account. |
| bank_id            | string            | BankID is the bank identifier.                                                |
| account_id         | string            | AccountID is the account identifier.                                          |
| other_account_id   | string            | OtherAccountID is the other_account identifier.                               |

##### Objects {#objects-CreateCorporateLocationRequest}

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses {#responses-method-post-createcorporatelocation}

#### Response body {#response-body-method-post-createcorporatelocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-CorporateLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-post-createcorporatelocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a image_url {#method-post-createimageurl}

Creates a new image_url and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
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

### HTTP Request {#http-request-method-post-createimageurl}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters {#query-parameters-method-post-createimageurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createimageurl}

| Name             | Type     | Description                                                 |
|------------------|----------|-------------------------------------------------------------|
| image_url        | ImageURL | ImageURL is the image_url metadata made on a other_account. |
| bank_id          | string   | BankID is the bank identifier.                              |
| account_id       | string   | AccountID is the account identifier.                        |
| other_account_id | string   | OtherAccountID is the other_account identifier.             |

##### Objects {#objects-CreateImageURLRequest}

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-post-createimageurl}

#### Response body {#response-body-method-post-createimageurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-post-createimageurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a more_info {#method-post-createmoreinfo}

Creates a new more_info and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
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

### HTTP Request {#http-request-method-post-createmoreinfo}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters {#query-parameters-method-post-createmoreinfo}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createmoreinfo}

| Name             | Type     | Description                                                 |
|------------------|----------|-------------------------------------------------------------|
| more_info        | MoreInfo | MoreInfo is the more_info metadata made on a other_account. |
| bank_id          | string   | BankID is the bank identifier.                              |
| account_id       | string   | AccountID is the account identifier.                        |
| other_account_id | string   | OtherAccountID is the other_account identifier.             |

##### Objects {#objects-CreateMoreInfoRequest}

###### MoreInfo

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

### Responses {#responses-method-post-createmoreinfo}

#### Response body {#response-body-method-post-createmoreinfo}

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

Example:

```json
{
  "more_info": "string"
}
```

#### Response codes {#response-codes-method-post-createmoreinfo}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a open_corporates_url {#method-post-createopencorporatesurl}

Creates a new open_corporates_url and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
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

### HTTP Request {#http-request-method-post-createopencorporatesurl}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters {#query-parameters-method-post-createopencorporatesurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createopencorporatesurl}

| Name                | Type              | Description                                                                    |
|---------------------|-------------------|--------------------------------------------------------------------------------|
| open_corporates_url | OpenCorporatesURL | OpenCorporatesURL is the open_corporates_url metadata made on a other_account. |
| bank_id             | string            | BankID is the bank identifier.                                                 |
| account_id          | string            | AccountID is the account identifier.                                           |
| other_account_id    | string            | OtherAccountID is the other_account identifier.                                |

##### Objects {#objects-CreateOpenCorporatesURLRequest}

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-post-createopencorporatesurl}

#### Response body {#response-body-method-post-createopencorporatesurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-post-createopencorporatesurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a physical_location {#method-post-createphysicallocation}

Creates a new physical_location and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
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

### HTTP Request {#http-request-method-post-createphysicallocation}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters {#query-parameters-method-post-createphysicallocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createphysicallocation}

| Name              | Type             | Description                                                                 |
|-------------------|------------------|-----------------------------------------------------------------------------|
| physical_location | PhysicalLocation | PhysicalLocation is the physical_location metadata made on a other_account. |
| bank_id           | string           | BankID is the bank identifier.                                              |
| account_id        | string           | AccountID is the account identifier.                                        |
| other_account_id  | string           | OtherAccountID is the other_account identifier.                             |

##### Objects {#objects-CreatePhysicalLocationRequest}

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses {#responses-method-post-createphysicallocation}

#### Response body {#response-body-method-post-createphysicallocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-PhysicalLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-post-createphysicallocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a private_alias {#method-post-createprivatealias}

Creates a new private_alias and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
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

### HTTP Request {#http-request-method-post-createprivatealias}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters {#query-parameters-method-post-createprivatealias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createprivatealias}

| Name             | Type        | Description                                                         |
|------------------|-------------|---------------------------------------------------------------------|
| alias            | PublicAlias | PrivateAlias is the private_alias metadata made on a other_account. |
| bank_id          | string      | BankID is the bank identifier.                                      |
| account_id       | string      | AccountID is the account identifier.                                |
| other_account_id | string      | OtherAccountID is the other_account identifier.                     |

##### Objects {#objects-CreatePrivateAliasRequest}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

### Responses {#responses-method-post-createprivatealias}

#### Response body {#response-body-method-post-createprivatealias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-post-createprivatealias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a public_alias {#method-post-createpublicalias}

Creates a new public_alias and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
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

### HTTP Request {#http-request-method-post-createpublicalias}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters {#query-parameters-method-post-createpublicalias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createpublicalias}

| Name             | Type        | Description                                                       |
|------------------|-------------|-------------------------------------------------------------------|
| alias            | PublicAlias | PublicAlias is the public_alias metadata made on a other_account. |
| bank_id          | string      | BankID is the bank identifier.                                    |
| account_id       | string      | AccountID is the account identifier.                              |
| other_account_id | string      | OtherAccountID is the other_account identifier.                   |

##### Objects {#objects-CreatePublicAliasRequest}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

### Responses {#responses-method-post-createpublicalias}

#### Response body {#response-body-method-post-createpublicalias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-post-createpublicalias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a url {#method-post-createurl}

Creates a new url and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
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

### HTTP Request {#http-request-method-post-createurl}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters {#query-parameters-method-post-createurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-post-createurl}

| Name             | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| url              | URL    | URL is the url metadata made on a other_account. |
| bank_id          | string | BankID is the bank identifier.                   |
| account_id       | string | AccountID is the account identifier.             |
| other_account_id | string | OtherAccountID is the other_account identifier.  |

##### Objects {#objects-CreateURLRequest}

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-post-createurl}

#### Response body {#response-body-method-post-createurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-post-createurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Corporate Location created successfully.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a corporate location {#method-delete-deletecorporatelocation}

Permanently delete a corporate location.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletecorporatelocation}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters {#query-parameters-method-delete-deletecorporatelocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deletecorporatelocation}

#### Response codes {#response-codes-method-delete-deletecorporatelocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | CorporateLocation successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a image url {#method-delete-deleteimageurl}

Permanently delete a image url.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deleteimageurl}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters {#query-parameters-method-delete-deleteimageurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deleteimageurl}

#### Response codes {#response-codes-method-delete-deleteimageurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ImageURL successfully deleted.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a more info {#method-delete-deletemoreinfo}

Permanently delete a more info.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletemoreinfo}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters {#query-parameters-method-delete-deletemoreinfo}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deletemoreinfo}

#### Response codes {#response-codes-method-delete-deletemoreinfo}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | MoreInfo successfully deleted.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a open corporates url {#method-delete-deleteopencorporatesurl}

Permanently delete a open corporates url.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deleteopencorporatesurl}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters {#query-parameters-method-delete-deleteopencorporatesurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deleteopencorporatesurl}

#### Response codes {#response-codes-method-delete-deleteopencorporatesurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | OpenCorporatesURL successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a physical location {#method-delete-deletephysicallocation}

Permanently delete a physical location.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletephysicallocation}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters {#query-parameters-method-delete-deletephysicallocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deletephysicallocation}

#### Response codes {#response-codes-method-delete-deletephysicallocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PhysicalLocation successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a private alias {#method-delete-deleteprivatealias}

Permanently delete a private alias.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deleteprivatealias}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters {#query-parameters-method-delete-deleteprivatealias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deleteprivatealias}

#### Response codes {#response-codes-method-delete-deleteprivatealias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PrivateAlias successfully deleted.                                                     |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a public alias {#method-delete-deletepublicalias}

Permanently delete a public alias.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletepublicalias}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters {#query-parameters-method-delete-deletepublicalias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deletepublicalias}

#### Response codes {#response-codes-method-delete-deletepublicalias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PublicAlias successfully deleted.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a url {#method-delete-deleteurl}

Permanently delete a url.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deleteurl}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters {#query-parameters-method-delete-deleteurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-delete-deleteurl}

#### Response codes {#response-codes-method-delete-deleteurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | URL successfully deleted.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve corporate location information {#method-get-getcorporatelocation}

Retrieve information about the corporate location specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcorporatelocation}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters {#query-parameters-method-get-getcorporatelocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getcorporatelocation}

#### Response body {#response-body-method-get-getcorporatelocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-CorporateLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-get-getcorporatelocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available corporate locations {#method-get-getcorporatelocations}

Retrieve information regarding all available corporate locations.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_locations \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcorporatelocations}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_locations`

### Query Parameters {#query-parameters-method-get-getcorporatelocations}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getcorporatelocations}

#### Response body {#response-body-method-get-getcorporatelocations}

| Name                | Type                 | Description                                            |
|---------------------|----------------------|--------------------------------------------------------|
| corporate_locations | \[]CorporateLocation | CorporateLocations is the list of corporate_locations. |
| bank_id             | string               | BankID is the bank identifier.                         |
| account_id          | string               | AccountID is the account identifier.                   |
| other_account_id    | string               | OtherAccountID is the other_account identifier.        |

##### Objects {#objects-GetCorporateLocationsResponse}

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

#### Response codes {#response-codes-method-get-getcorporatelocations}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve image url information {#method-get-getimageurl}

Retrieve information about the image url specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getimageurl}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters {#query-parameters-method-get-getimageurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getimageurl}

#### Response body {#response-body-method-get-getimageurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-get-getimageurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available image urls {#method-get-getimageurls}

Retrieve information regarding all available image urls.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getimageurls}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_urls`

### Query Parameters {#query-parameters-method-get-getimageurls}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getimageurls}

#### Response body {#response-body-method-get-getimageurls}

| Name             | Type        | Description                                     |
|------------------|-------------|-------------------------------------------------|
| image_urls       | \[]ImageURL | ImageURLs is the list of image_urls.            |
| bank_id          | string      | BankID is the bank identifier.                  |
| account_id       | string      | AccountID is the account identifier.            |
| other_account_id | string      | OtherAccountID is the other_account identifier. |

##### Objects {#objects-GetImageURLsResponse}

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

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

#### Response codes {#response-codes-method-get-getimageurls}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve more info information {#method-get-getmoreinfo}

Retrieve information about the more info specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getmoreinfo}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters {#query-parameters-method-get-getmoreinfo}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getmoreinfo}

#### Response body {#response-body-method-get-getmoreinfo}

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

Example:

```json
{
  "more_info": "string"
}
```

#### Response codes {#response-codes-method-get-getmoreinfo}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available more infos {#method-get-getmoreinfos}

Retrieve information regarding all available more infos.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_infos \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getmoreinfos}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_infos`

### Query Parameters {#query-parameters-method-get-getmoreinfos}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getmoreinfos}

#### Response body {#response-body-method-get-getmoreinfos}

| Name             | Type        | Description                                     |
|------------------|-------------|-------------------------------------------------|
| more_infos       | \[]MoreInfo | MoreInfos is the list of more_infos.            |
| bank_id          | string      | BankID is the bank identifier.                  |
| account_id       | string      | AccountID is the account identifier.            |
| other_account_id | string      | OtherAccountID is the other_account identifier. |

##### Objects {#objects-GetMoreInfosResponse}

###### MoreInfo

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

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

#### Response codes {#response-codes-method-get-getmoreinfos}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve open corporates url information {#method-get-getopencorporatesurl}

Retrieve information about the open corporates url specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getopencorporatesurl}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters {#query-parameters-method-get-getopencorporatesurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getopencorporatesurl}

#### Response body {#response-body-method-get-getopencorporatesurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-get-getopencorporatesurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available open corporates urls {#method-get-getopencorporatesurls}

Retrieve information regarding all available open corporates urls.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getopencorporatesurls}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_urls`

### Query Parameters {#query-parameters-method-get-getopencorporatesurls}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getopencorporatesurls}

#### Response body {#response-body-method-get-getopencorporatesurls}

| Name                 | Type                 | Description                                             |
|----------------------|----------------------|---------------------------------------------------------|
| open_corporates_urls | \[]OpenCorporatesURL | OpenCorporatesURLs is the list of open_corporates_urls. |
| bank_id              | string               | BankID is the bank identifier.                          |
| account_id           | string               | AccountID is the account identifier.                    |
| other_account_id     | string               | OtherAccountID is the other_account identifier.         |

##### Objects {#objects-GetOpenCorporatesURLsResponse}

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

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

#### Response codes {#response-codes-method-get-getopencorporatesurls}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve other account metadata information {#method-get-getotheraccountmetadata}

Retrieve information about the other account metadata specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getotheraccountmetadata}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata`

### Query Parameters {#query-parameters-method-get-getotheraccountmetadata}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getotheraccountmetadata}

#### Response body {#response-body-method-get-getotheraccountmetadata}

| Name                | Type              | Description |
|---------------------|-------------------|-------------|
| public_alias        | string            |             |
| private_alias       | string            |             |
| more_info           | string            |             |
| url                 | string            |             |
| image_url           | string            |             |
| open_corporates_url | string            |             |
| corporate_location  | CorporateLocation |             |
| physical_location   | PhysicalLocation  |             |

##### Objects {#objects-GetOtherAccountMetadataResponse}

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

#### Response codes {#response-codes-method-get-getotheraccountmetadata}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve physical location information {#method-get-getphysicallocation}

Retrieve information about the physical location specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getphysicallocation}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters {#query-parameters-method-get-getphysicallocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getphysicallocation}

#### Response body {#response-body-method-get-getphysicallocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-PhysicalLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-get-getphysicallocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available physical locations {#method-get-getphysicallocations}

Retrieve information regarding all available physical locations.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_locations \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getphysicallocations}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_locations`

### Query Parameters {#query-parameters-method-get-getphysicallocations}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getphysicallocations}

#### Response body {#response-body-method-get-getphysicallocations}

| Name               | Type                | Description                                          |
|--------------------|---------------------|------------------------------------------------------|
| physical_locations | \[]PhysicalLocation | PhysicalLocations is the list of physical_locations. |
| bank_id            | string              | BankID is the bank identifier.                       |
| account_id         | string              | AccountID is the account identifier.                 |
| other_account_id   | string              | OtherAccountID is the other_account identifier.      |

##### Objects {#objects-GetPhysicalLocationsResponse}

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

#### Response codes {#response-codes-method-get-getphysicallocations}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve private alias information {#method-get-getprivatealias}

Retrieve information about the private alias specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getprivatealias}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters {#query-parameters-method-get-getprivatealias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getprivatealias}

#### Response body {#response-body-method-get-getprivatealias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-get-getprivatealias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available private aliases {#method-get-getprivatealiases}

Retrieve information regarding all available private aliases.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_aliases \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getprivatealiases}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_aliases`

### Query Parameters {#query-parameters-method-get-getprivatealiases}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getprivatealiases}

#### Response body {#response-body-method-get-getprivatealiases}

| Name             | Type           | Description                                     |
|------------------|----------------|-------------------------------------------------|
| aliases          | \[]PublicAlias | PrivateAliases is the list of private_aliases.  |
| bank_id          | string         | BankID is the bank identifier.                  |
| account_id       | string         | AccountID is the account identifier.            |
| other_account_id | string         | OtherAccountID is the other_account identifier. |

##### Objects {#objects-GetPrivateAliasesResponse}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

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

#### Response codes {#response-codes-method-get-getprivatealiases}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve public alias information {#method-get-getpublicalias}

Retrieve information about the public alias specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getpublicalias}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters {#query-parameters-method-get-getpublicalias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getpublicalias}

#### Response body {#response-body-method-get-getpublicalias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-get-getpublicalias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available public aliases {#method-get-getpublicaliases}

Retrieve information regarding all available public aliases.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_aliases \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getpublicaliases}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_aliases`

### Query Parameters {#query-parameters-method-get-getpublicaliases}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-getpublicaliases}

#### Response body {#response-body-method-get-getpublicaliases}

| Name             | Type           | Description                                     |
|------------------|----------------|-------------------------------------------------|
| aliases          | \[]PublicAlias | PublicAliases is the list of public_aliases.    |
| bank_id          | string         | BankID is the bank identifier.                  |
| account_id       | string         | AccountID is the account identifier.            |
| other_account_id | string         | OtherAccountID is the other_account identifier. |

##### Objects {#objects-GetPublicAliasesResponse}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

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

#### Response codes {#response-codes-method-get-getpublicaliases}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve url information {#method-get-geturl}

Retrieve information about the url specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-geturl}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters {#query-parameters-method-get-geturl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-geturl}

#### Response body {#response-body-method-get-geturl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-get-geturl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve all available urls {#method-get-geturls}

Retrieve information regarding all available urls.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/urls \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-geturls}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/urls`

### Query Parameters {#query-parameters-method-get-geturls}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Responses {#responses-method-get-geturls}

#### Response body {#response-body-method-get-geturls}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| urls             | \[]URL | URLs is the list of urls.                       |
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

##### Objects {#objects-GetURLsResponse}

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

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

#### Response codes {#response-codes-method-get-geturls}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a corporate location {#method-put-updatecorporatelocation}

Updates a corporate location's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location \
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

### HTTP Request {#http-request-method-put-updatecorporatelocation}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/corporate_location`

### Query Parameters {#query-parameters-method-put-updatecorporatelocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updatecorporatelocation}

| Name               | Type              | Description                                                                   |
|--------------------|-------------------|-------------------------------------------------------------------------------|
| corporate_location | CorporateLocation | CorporateLocation is the corporate_location metadata made on a other_account. |
| bank_id            | string            | BankID is the bank identifier.                                                |
| account_id         | string            | AccountID is the account identifier.                                          |
| other_account_id   | string            | OtherAccountID is the other_account identifier.                               |

##### Objects {#objects-UpdateCorporateLocationRequest}

###### CorporateLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses {#responses-method-put-updatecorporatelocation}

#### Response body {#response-body-method-put-updatecorporatelocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-CorporateLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-put-updatecorporatelocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | CorporateLocation successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a image url {#method-put-updateimageurl}

Updates a image url's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url \
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

### HTTP Request {#http-request-method-put-updateimageurl}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/image_url`

### Query Parameters {#query-parameters-method-put-updateimageurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updateimageurl}

| Name             | Type     | Description                                                 |
|------------------|----------|-------------------------------------------------------------|
| image_url        | ImageURL | ImageURL is the image_url metadata made on a other_account. |
| bank_id          | string   | BankID is the bank identifier.                              |
| account_id       | string   | AccountID is the account identifier.                        |
| other_account_id | string   | OtherAccountID is the other_account identifier.             |

##### Objects {#objects-UpdateImageURLRequest}

###### ImageURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-put-updateimageurl}

#### Response body {#response-body-method-put-updateimageurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-put-updateimageurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ImageURL successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a more info {#method-put-updatemoreinfo}

Updates a more info's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info \
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

### HTTP Request {#http-request-method-put-updatemoreinfo}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/more_info`

### Query Parameters {#query-parameters-method-put-updatemoreinfo}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updatemoreinfo}

| Name             | Type     | Description                                                 |
|------------------|----------|-------------------------------------------------------------|
| more_info        | MoreInfo | MoreInfo is the more_info metadata made on a other_account. |
| bank_id          | string   | BankID is the bank identifier.                              |
| account_id       | string   | AccountID is the account identifier.                        |
| other_account_id | string   | OtherAccountID is the other_account identifier.             |

##### Objects {#objects-UpdateMoreInfoRequest}

###### MoreInfo

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

### Responses {#responses-method-put-updatemoreinfo}

#### Response body {#response-body-method-put-updatemoreinfo}

| Name      | Type   | Description |
|-----------|--------|-------------|
| more_info | string |             |

Example:

```json
{
  "more_info": "string"
}
```

#### Response codes {#response-codes-method-put-updatemoreinfo}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | MoreInfo successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a open corporates url {#method-put-updateopencorporatesurl}

Updates a open corporates url's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url \
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

### HTTP Request {#http-request-method-put-updateopencorporatesurl}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/open_corporates_url`

### Query Parameters {#query-parameters-method-put-updateopencorporatesurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updateopencorporatesurl}

| Name                | Type              | Description                                                                    |
|---------------------|-------------------|--------------------------------------------------------------------------------|
| open_corporates_url | OpenCorporatesURL | OpenCorporatesURL is the open_corporates_url metadata made on a other_account. |
| bank_id             | string            | BankID is the bank identifier.                                                 |
| account_id          | string            | AccountID is the account identifier.                                           |
| other_account_id    | string            | OtherAccountID is the other_account identifier.                                |

##### Objects {#objects-UpdateOpenCorporatesURLRequest}

###### OpenCorporatesURL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-put-updateopencorporatesurl}

#### Response body {#response-body-method-put-updateopencorporatesurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-put-updateopencorporatesurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | OpenCorporatesURL successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a physical location {#method-put-updatephysicallocation}

Updates a physical location's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location \
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

### HTTP Request {#http-request-method-put-updatephysicallocation}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/physical_location`

### Query Parameters {#query-parameters-method-put-updatephysicallocation}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updatephysicallocation}

| Name              | Type             | Description                                                                 |
|-------------------|------------------|-----------------------------------------------------------------------------|
| physical_location | PhysicalLocation | PhysicalLocation is the physical_location metadata made on a other_account. |
| bank_id           | string           | BankID is the bank identifier.                                              |
| account_id        | string           | AccountID is the account identifier.                                        |
| other_account_id  | string           | OtherAccountID is the other_account identifier.                             |

##### Objects {#objects-UpdatePhysicalLocationRequest}

###### PhysicalLocation

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

### Responses {#responses-method-put-updatephysicallocation}

#### Response body {#response-body-method-put-updatephysicallocation}

| Name     | Type     | Description |
|----------|----------|-------------|
| location | Location |             |

##### Objects {#objects-PhysicalLocation}

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  }
}
```

#### Response codes {#response-codes-method-put-updatephysicallocation}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PhysicalLocation successfully updated.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a private alias {#method-put-updateprivatealias}

Updates a private alias's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias \
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

### HTTP Request {#http-request-method-put-updateprivatealias}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/private_alias`

### Query Parameters {#query-parameters-method-put-updateprivatealias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updateprivatealias}

| Name             | Type        | Description                                                         |
|------------------|-------------|---------------------------------------------------------------------|
| alias            | PublicAlias | PrivateAlias is the private_alias metadata made on a other_account. |
| bank_id          | string      | BankID is the bank identifier.                                      |
| account_id       | string      | AccountID is the account identifier.                                |
| other_account_id | string      | OtherAccountID is the other_account identifier.                     |

##### Objects {#objects-UpdatePrivateAliasRequest}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

### Responses {#responses-method-put-updateprivatealias}

#### Response body {#response-body-method-put-updateprivatealias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-put-updateprivatealias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PrivateAlias successfully updated.                                                     |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a public alias {#method-put-updatepublicalias}

Updates a public alias's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias \
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

### HTTP Request {#http-request-method-put-updatepublicalias}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/public_alias`

### Query Parameters {#query-parameters-method-put-updatepublicalias}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updatepublicalias}

| Name             | Type        | Description                                                       |
|------------------|-------------|-------------------------------------------------------------------|
| alias            | PublicAlias | PublicAlias is the public_alias metadata made on a other_account. |
| bank_id          | string      | BankID is the bank identifier.                                    |
| account_id       | string      | AccountID is the account identifier.                              |
| other_account_id | string      | OtherAccountID is the other_account identifier.                   |

##### Objects {#objects-UpdatePublicAliasRequest}

###### PublicAlias

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

### Responses {#responses-method-put-updatepublicalias}

#### Response body {#response-body-method-put-updatepublicalias}

| Name  | Type   | Description |
|-------|--------|-------------|
| alias | string |             |

Example:

```json
{
  "alias": "string"
}
```

#### Response codes {#response-codes-method-put-updatepublicalias}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | PublicAlias successfully updated.                                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a url {#method-put-updateurl}

Updates a url's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url \
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

### HTTP Request {#http-request-method-put-updateurl}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/other_accounts/{OtherAccountID}/metadata/url`

### Query Parameters {#query-parameters-method-put-updateurl}

| Name             | Type   | Description                                     |
|------------------|--------|-------------------------------------------------|
| bank_id          | string | BankID is the bank identifier.                  |
| account_id       | string | AccountID is the account identifier.            |
| other_account_id | string | OtherAccountID is the other_account identifier. |

### Body Parameters {#body-parameters-method-put-updateurl}

| Name             | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| url              | URL    | URL is the url metadata made on a other_account. |
| bank_id          | string | BankID is the bank identifier.                   |
| account_id       | string | AccountID is the account identifier.             |
| other_account_id | string | OtherAccountID is the other_account identifier.  |

##### Objects {#objects-UpdateURLRequest}

###### URL

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

### Responses {#responses-method-put-updateurl}

#### Response body {#response-body-method-put-updateurl}

| Name | Type   | Description |
|------|--------|-------------|
| url  | string |             |

Example:

```json
{
  "url": "string"
}
```

#### Response codes {#response-codes-method-put-updateurl}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | URL successfully updated.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
