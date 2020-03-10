Counterpart API v1.0.0
======================

Provides CRUD operations on the counter part resource.

* Host `https://`

* Base Path ``

Create an counter party {#method-post-createcounterparty}
---------------------------------------------------------

Creates a new counter party

```sh
curl -X POST \
	https:///v1/counterparty \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"name": "string",
		"description": "string",
		"other_account_routing_scheme": "string",
		"other_account_routing_address": "string",
		"other_account_secondary_routing_scheme": "string",
		"other_account_secondary_routing_address": "string",
		"other_bank_routing_scheme": "string",
		"other_bank_routing_address": "string",
		"other_branch_routing_scheme": "string",
		"other_branch_routing_address": "string",
		"is_beneficiary": "bool",
		"bespoke": [
			{
				"key": "string",
				"value": "string"
			}
		]
	}'
```

### HTTP Request

`POST https:///v1/counterparty`

### Body Parameters

| Name                                    | Type       | Description |
|-----------------------------------------|------------|-------------|
| name                                    | string     |             |
| description                             | string     |             |
| other_account_routing_scheme            | string     |             |
| other_account_routing_address           | string     |             |
| other_account_secondary_routing_scheme  | string     |             |
| other_account_secondary_routing_address | string     |             |
| other_bank_routing_scheme               | string     |             |
| other_bank_routing_address              | string     |             |
| other_branch_routing_scheme             | string     |             |
| other_branch_routing_address            | string     |             |
| is_beneficiary                          | bool       |             |
| bespoke                                 | \[]Bespoke |             |

##### Objects

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| key   | string | any info-key you want to add to this counerparty   |
| value | string | any info-value you want to add to this counerparty |

### Responses

#### Response body

| Name              | Type         | Description |
|-------------------|--------------|-------------|
| counter_party     | CounterParty |             |
| metadata          | Metadata     |             |
| physical_location | Location     |             |
| private_alias     | string       |             |

##### Objects

###### CounterParty

| Name                                    | Type       | Description                                                                 |
|-----------------------------------------|------------|-----------------------------------------------------------------------------|
| name                                    | string     | The human readable name (e.g. John Bravo)                                   |
| description                             | string     | The description of the about counter party                                  |
| other_account_routing_scheme            | string     | Account routing schme such as AccountId or AccountNumber or any strings     |
| other_account_routing_address           | string     | Account routing address is a valid account identifier                       |
| other_account_secondary_routing_scheme  | string     | Account secondary routing address such as IBan                              |
| other_account_secondary_routing_address | string     | IBan it should be unique for each counterparty.                             |
| other_bank_routing_scheme               | string     | Bank routing scheme such as bankId or bankCode or any strings               |
| other_bank_routing_address              | string     | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| other_branch_routing_scheme             | string     | Branch routing scheme such as branchId or any other strings                 |
| other_branch_routing_address            | string     | Branch routing address like branch-id-765 or you can leave it empty         |
| is_beneficiary                          | bool       | Must be set to true in order to send payments to this counterparty          |
| bespoke                                 | \[]Bespoke | It support list of key-value, you can add it to the counterarty.            |
| metadata                                | Metadata   | Metadata about the counter party                                            |

###### Metadata

| Name                | Type     | Description |
|---------------------|----------|-------------|
| public_alias        | string   |             |
| more_info           | string   |             |
| url                 | string   |             |
| image_url           | string   |             |
| open_corporates_url | string   |             |
| corporate_location  | Location |             |
| physical_location   | Location |             |
| private_alias       | string   |             |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| key   | string | any info-key you want to add to this counerparty   |
| value | string | any info-value you want to add to this counerparty |

Example:

```json
{
  "counter_party": {
    "name": "string",
    "description": "string",
    "other_account_routing_scheme": "string",
    "other_account_routing_address": "string",
    "other_account_secondary_routing_scheme": "string",
    "other_account_secondary_routing_address": "string",
    "other_bank_routing_scheme": "string",
    "other_bank_routing_address": "string",
    "other_branch_routing_scheme": "string",
    "other_branch_routing_address": "string",
    "is_beneficiary": "bool",
    "bespoke": [
      {
        "key": "string",
        "value": "string"
      }
    ],
    "metadata": {
      "public_alias": "string",
      "more_info": "string",
      "url": "string",
      "image_url": "string",
      "open_corporates_url": "string",
      "corporate_location": {
        "latitude": "double",
        "longitude": "double"
      },
      "physical_location": {
        "latitude": "double",
        "longitude": "double"
      },
      "private_alias": "string"
    }
  },
  "metadata": {
    "public_alias": "string",
    "more_info": "string",
    "url": "string",
    "image_url": "string",
    "open_corporates_url": "string",
    "corporate_location": {
      "latitude": "double",
      "longitude": "double"
    },
    "physical_location": {
      "latitude": "double",
      "longitude": "double"
    },
    "private_alias": "string"
  },
  "physical_location": {
    "latitude": "double",
    "longitude": "double"
  },
  "private_alias": "string"
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Counter Party created successfully.                                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

List all accounts {#method-get-getcounterparties}
-------------------------------------------------

Returns a list containing up to 20 accounts. `after_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/counterparties \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/counterparties`

### Responses

#### Response body

| Name   | Type            | Description                           |
|--------|-----------------|---------------------------------------|
| result | \[]CounterParty | Result is the paginated query result. |

##### Objects

###### CounterParty

| Name                                    | Type       | Description                                                                 |
|-----------------------------------------|------------|-----------------------------------------------------------------------------|
| name                                    | string     | The human readable name (e.g. John Bravo)                                   |
| description                             | string     | The description of the about counter party                                  |
| other_account_routing_scheme            | string     | Account routing schme such as AccountId or AccountNumber or any strings     |
| other_account_routing_address           | string     | Account routing address is a valid account identifier                       |
| other_account_secondary_routing_scheme  | string     | Account secondary routing address such as IBan                              |
| other_account_secondary_routing_address | string     | IBan it should be unique for each counterparty.                             |
| other_bank_routing_scheme               | string     | Bank routing scheme such as bankId or bankCode or any strings               |
| other_bank_routing_address              | string     | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| other_branch_routing_scheme             | string     | Branch routing scheme such as branchId or any other strings                 |
| other_branch_routing_address            | string     | Branch routing address like branch-id-765 or you can leave it empty         |
| is_beneficiary                          | bool       | Must be set to true in order to send payments to this counterparty          |
| bespoke                                 | \[]Bespoke | It support list of key-value, you can add it to the counterarty.            |
| metadata                                | Metadata   | Metadata about the counter party                                            |

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| key   | string | any info-key you want to add to this counerparty   |
| value | string | any info-value you want to add to this counerparty |

###### Metadata

| Name                | Type     | Description |
|---------------------|----------|-------------|
| public_alias        | string   |             |
| more_info           | string   |             |
| url                 | string   |             |
| image_url           | string   |             |
| open_corporates_url | string   |             |
| corporate_location  | Location |             |
| physical_location   | Location |             |
| private_alias       | string   |             |

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
      "name": "string",
      "description": "string",
      "other_account_routing_scheme": "string",
      "other_account_routing_address": "string",
      "other_account_secondary_routing_scheme": "string",
      "other_account_secondary_routing_address": "string",
      "other_bank_routing_scheme": "string",
      "other_bank_routing_address": "string",
      "other_branch_routing_scheme": "string",
      "other_branch_routing_address": "string",
      "is_beneficiary": "bool",
      "bespoke": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "metadata": {
        "public_alias": "string",
        "more_info": "string",
        "url": "string",
        "image_url": "string",
        "open_corporates_url": "string",
        "corporate_location": {
          "latitude": "double",
          "longitude": "double"
        },
        "physical_location": {
          "latitude": "double",
          "longitude": "double"
        },
        "private_alias": "string"
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

Retrieve an counter party {#method-get-getcounterparty}
-------------------------------------------------------

Retrieves all data from an counter party selected by the supplied counter_party_id.

```sh
curl -X GET \
	https:///v1/counterparty/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/counterparty/{ID}`

### Query Parameters

| Name | Type   | Description |
|------|--------|-------------|
| id   | string |             |

### Responses

#### Response body

| Name                                    | Type       | Description                                                                 |
|-----------------------------------------|------------|-----------------------------------------------------------------------------|
| name                                    | string     | The human readable name (e.g. John Bravo)                                   |
| description                             | string     | The description of the about counter party                                  |
| other_account_routing_scheme            | string     | Account routing schme such as AccountId or AccountNumber or any strings     |
| other_account_routing_address           | string     | Account routing address is a valid account identifier                       |
| other_account_secondary_routing_scheme  | string     | Account secondary routing address such as IBan                              |
| other_account_secondary_routing_address | string     | IBan it should be unique for each counterparty.                             |
| other_bank_routing_scheme               | string     | Bank routing scheme such as bankId or bankCode or any strings               |
| other_bank_routing_address              | string     | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| other_branch_routing_scheme             | string     | Branch routing scheme such as branchId or any other strings                 |
| other_branch_routing_address            | string     | Branch routing address like branch-id-765 or you can leave it empty         |
| is_beneficiary                          | bool       | Must be set to true in order to send payments to this counterparty          |
| bespoke                                 | \[]Bespoke | It support list of key-value, you can add it to the counterarty.            |
| metadata                                | Metadata   | Metadata about the counter party                                            |

##### Objects

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| key   | string | any info-key you want to add to this counerparty   |
| value | string | any info-value you want to add to this counerparty |

###### Metadata

| Name                | Type     | Description |
|---------------------|----------|-------------|
| public_alias        | string   |             |
| more_info           | string   |             |
| url                 | string   |             |
| image_url           | string   |             |
| open_corporates_url | string   |             |
| corporate_location  | Location |             |
| physical_location   | Location |             |
| private_alias       | string   |             |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "name": "string",
  "description": "string",
  "other_account_routing_scheme": "string",
  "other_account_routing_address": "string",
  "other_account_secondary_routing_scheme": "string",
  "other_account_secondary_routing_address": "string",
  "other_bank_routing_scheme": "string",
  "other_bank_routing_address": "string",
  "other_branch_routing_scheme": "string",
  "other_branch_routing_address": "string",
  "is_beneficiary": "bool",
  "bespoke": [
    {
      "key": "string",
      "value": "string"
    }
  ],
  "metadata": {
    "public_alias": "string",
    "more_info": "string",
    "url": "string",
    "image_url": "string",
    "open_corporates_url": "string",
    "corporate_location": {
      "latitude": "double",
      "longitude": "double"
    },
    "physical_location": {
      "latitude": "double",
      "longitude": "double"
    },
    "private_alias": "string"
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

Retrieve an account with other account id {#method-get-getotheraccountbyid}
---------------------------------------------------------------------------

Retrieves other account (counter party) by other account id

```sh
curl -X GET \
	https:///v1/accounts/{AccountID}/otheraccounts/{OtherAccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}/otheraccounts/{OtherAccountID}`

### Query Parameters

| Name             | Type   | Description |
|------------------|--------|-------------|
| account_id       | string |             |
| other_account_id | string |             |

### Responses

#### Response body

| Name          | Type         | Description |
|---------------|--------------|-------------|
| other_account | OtherAccount |             |

##### Objects

###### OtherAccount

| Name             | Type               | Description |
|------------------|--------------------|-------------|
| id               | string             |             |
| holder           | Holder             |             |
| bank_routing     | BankRouting        |             |
| account_routings | \[]AccountRoutings |             |
| metadata         | Metadata           |             |

###### Holder

| Name     | Type   | Description |
|----------|--------|-------------|
| name     | string |             |
| is_alias | bool   |             |

###### BankRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| scheme  | string |             |
| address | string |             |

###### AccountRoutings

| Name    | Type   | Description |
|---------|--------|-------------|
| scheme  | string |             |
| address | string |             |

###### Metadata

| Name                | Type     | Description |
|---------------------|----------|-------------|
| public_alias        | string   |             |
| more_info           | string   |             |
| url                 | string   |             |
| image_url           | string   |             |
| open_corporates_url | string   |             |
| corporate_location  | Location |             |
| physical_location   | Location |             |
| private_alias       | string   |             |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "other_account": {
    "id": "string",
    "holder": {
      "name": "string",
      "is_alias": "bool"
    },
    "bank_routing": {
      "scheme": "string",
      "address": "string"
    },
    "account_routings": [
      {
        "scheme": "string",
        "address": "string"
      }
    ],
    "metadata": {
      "public_alias": "string",
      "more_info": "string",
      "url": "string",
      "image_url": "string",
      "open_corporates_url": "string",
      "corporate_location": {
        "latitude": "double",
        "longitude": "double"
      },
      "physical_location": {
        "latitude": "double",
        "longitude": "double"
      },
      "private_alias": "string"
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

Retrieve all other account with account id {#method-get-getotheraccounts}
-------------------------------------------------------------------------

Retrieves all other accounts (counter party) for an account id

```sh
curl -X GET \
	https:///v1/accounts/{AccountID}/otheraccounts \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}/otheraccounts`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| account_id | string |             |

### Responses

#### Response body

| Name           | Type            | Description |
|----------------|-----------------|-------------|
| other_accounts | \[]OtherAccount |             |

##### Objects

###### OtherAccount

| Name             | Type               | Description |
|------------------|--------------------|-------------|
| id               | string             |             |
| holder           | Holder             |             |
| bank_routing     | BankRouting        |             |
| account_routings | \[]AccountRoutings |             |
| metadata         | Metadata           |             |

###### Holder

| Name     | Type   | Description |
|----------|--------|-------------|
| name     | string |             |
| is_alias | bool   |             |

###### BankRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| scheme  | string |             |
| address | string |             |

###### AccountRoutings

| Name    | Type   | Description |
|---------|--------|-------------|
| scheme  | string |             |
| address | string |             |

###### Metadata

| Name                | Type     | Description |
|---------------------|----------|-------------|
| public_alias        | string   |             |
| more_info           | string   |             |
| url                 | string   |             |
| image_url           | string   |             |
| open_corporates_url | string   |             |
| corporate_location  | Location |             |
| physical_location   | Location |             |
| private_alias       | string   |             |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "other_accounts": [
    {
      "id": "string",
      "holder": {
        "name": "string",
        "is_alias": "bool"
      },
      "bank_routing": {
        "scheme": "string",
        "address": "string"
      },
      "account_routings": [
        {
          "scheme": "string",
          "address": "string"
        }
      ],
      "metadata": {
        "public_alias": "string",
        "more_info": "string",
        "url": "string",
        "image_url": "string",
        "open_corporates_url": "string",
        "corporate_location": {
          "latitude": "double",
          "longitude": "double"
        },
        "physical_location": {
          "latitude": "double",
          "longitude": "double"
        },
        "private_alias": "string"
      }
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

Annex
-----

#### TransactionRequestType

Entity type defines the type of counterparty

| Value                         | Description                                  |
|-------------------------------|----------------------------------------------|
| UnknownTransactionRequestType |                                              |
| COUNTERPARTY                  | Counter party type                           |
| SEPA                          | TransactionRequestType_SEPA transaction type |
