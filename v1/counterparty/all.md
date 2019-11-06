# Counterpart API v1.0.0

Provides CRUD operations on the counter part resource.

*
Host ``
EOL

*
Base Path ``

## Create an counter party

Creates a new counter party

```sh
curl -X POST \
	/v1/counterparty \
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
{{snippet createcounterparty []}}

### HTTP Request

`POST /v1/counterparty`

### Body Parameters

| Name                                | Type      | Description |
|-------------------------------------|-----------|-------------|
| Name                                | string    |             |
| Description                         | string    |             |
| OtherAccountRoutingScheme           | string    |             |
| OtherAccountRoutingAddress          | string    |             |
| OtherAccountSecondaryRoutingScheme  | string    |             |
| OtherAccountSecondaryRoutingAddress | string    |             |
| OtherBankRoutingScheme              | string    |             |
| OtherBankRoutingAddress             | string    |             |
| OtherBranchRoutingScheme            | string    |             |
| OtherBranchRoutingAddress           | string    |             |
| IsBeneficiary                       | bool      |             |
| Bespoke                             | []Bespoke |             |

##### Objects

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| Key   | string | any info-key you want to add to this counerparty   |
| Value | string | any info-value you want to add to this counerparty |

### Responses

#### Response body

| Name             | Type         | Description |
|------------------|--------------|-------------|
| CounterParty     | CounterParty |             |
| Metadata         | Metadata     |             |
| PhysicalLocation | Location     |             |
| PrivateAlias     | string       |             |

##### Objects

###### CounterParty

| Name                                | Type      | Description                                                                 |
|-------------------------------------|-----------|-----------------------------------------------------------------------------|
| Name                                | string    | The human readable name (e.g. John Bravo)                                   |
| Description                         | string    | The description of the about counter party                                  |
| OtherAccountRoutingScheme           | string    | Account routing schme such as AccountId or AccountNumber or any strings     |
| OtherAccountRoutingAddress          | string    | Account routing address is a valid account identifier                       |
| OtherAccountSecondaryRoutingScheme  | string    | Account secondary routing address such as IBan                              |
| OtherAccountSecondaryRoutingAddress | string    | IBan it should be unique for each counterparty.                             |
| OtherBankRoutingScheme              | string    | Bank routing scheme such as bankId or bankCode or any strings               |
| OtherBankRoutingAddress             | string    | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| OtherBranchRoutingScheme            | string    | Branch routing scheme such as branchId or any other strings                 |
| OtherBranchRoutingAddress           | string    | Branch routing address like branch-id-765 or you can leave it empty         |
| IsBeneficiary                       | bool      | Must be set to true in order to send payments to this counterparty          |
| Bespoke                             | []Bespoke | It support list of key-value, you can add it to the counterarty.            |
| Metadata                            | Metadata  | Metadata about the counter party                                            |

###### Metadata

| Name              | Type     | Description |
|-------------------|----------|-------------|
| PublicAlias       | string   |             |
| MoreInfo          | string   |             |
| URL               | string   |             |
| ImageURL          | string   |             |
| OpenCorporatesURL | string   |             |
| CorporateLocation | Location |             |
| PhysicalLocation  | Location |             |
| PrivateAlias      | string   |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| Key   | string | any info-key you want to add to this counerparty   |
| Value | string | any info-value you want to add to this counerparty |

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

## List all accounts

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/counterparties \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcounterparties []}}

### HTTP Request

`GET /v1/counterparties`

### Responses

#### Response body

| Name    | Type           | Description                                             |
|---------|----------------|---------------------------------------------------------|
| Result  | []CounterParty | Result is a list containing up to 20 Counter parties.   |
| HasMore | bool           | HasMore indicates if there are more accounts available. |

##### Objects

###### CounterParty

| Name                                | Type      | Description                                                                 |
|-------------------------------------|-----------|-----------------------------------------------------------------------------|
| Name                                | string    | The human readable name (e.g. John Bravo)                                   |
| Description                         | string    | The description of the about counter party                                  |
| OtherAccountRoutingScheme           | string    | Account routing schme such as AccountId or AccountNumber or any strings     |
| OtherAccountRoutingAddress          | string    | Account routing address is a valid account identifier                       |
| OtherAccountSecondaryRoutingScheme  | string    | Account secondary routing address such as IBan                              |
| OtherAccountSecondaryRoutingAddress | string    | IBan it should be unique for each counterparty.                             |
| OtherBankRoutingScheme              | string    | Bank routing scheme such as bankId or bankCode or any strings               |
| OtherBankRoutingAddress             | string    | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| OtherBranchRoutingScheme            | string    | Branch routing scheme such as branchId or any other strings                 |
| OtherBranchRoutingAddress           | string    | Branch routing address like branch-id-765 or you can leave it empty         |
| IsBeneficiary                       | bool      | Must be set to true in order to send payments to this counterparty          |
| Bespoke                             | []Bespoke | It support list of key-value, you can add it to the counterarty.            |
| Metadata                            | Metadata  | Metadata about the counter party                                            |

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| Key   | string | any info-key you want to add to this counerparty   |
| Value | string | any info-value you want to add to this counerparty |

###### Metadata

| Name              | Type     | Description |
|-------------------|----------|-------------|
| PublicAlias       | string   |             |
| MoreInfo          | string   |             |
| URL               | string   |             |
| ImageURL          | string   |             |
| OpenCorporatesURL | string   |             |
| CorporateLocation | Location |             |
| PhysicalLocation  | Location |             |
| PrivateAlias      | string   |             |

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
  ],
  "has_more": "bool"
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

## Retrieve an counter party

Retrieves all data from an counter party selected by the supplied counterpartyid.

```sh
curl -X GET \
	/v1/counterparty/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getcounterparty []}}

### HTTP Request

`GET /v1/counterparty/{ID}`

### Query Parameters

| Name | Type   | Description |
|------|--------|-------------|
| ID   | string |             |

### Responses

#### Response body

| Name                                | Type      | Description                                                                 |
|-------------------------------------|-----------|-----------------------------------------------------------------------------|
| Name                                | string    | The human readable name (e.g. John Bravo)                                   |
| Description                         | string    | The description of the about counter party                                  |
| OtherAccountRoutingScheme           | string    | Account routing schme such as AccountId or AccountNumber or any strings     |
| OtherAccountRoutingAddress          | string    | Account routing address is a valid account identifier                       |
| OtherAccountSecondaryRoutingScheme  | string    | Account secondary routing address such as IBan                              |
| OtherAccountSecondaryRoutingAddress | string    | IBan it should be unique for each counterparty.                             |
| OtherBankRoutingScheme              | string    | Bank routing scheme such as bankId or bankCode or any strings               |
| OtherBankRoutingAddress             | string    | Bank routing address such as eg: testsandbox, must be valid sandbox bankIds |
| OtherBranchRoutingScheme            | string    | Branch routing scheme such as branchId or any other strings                 |
| OtherBranchRoutingAddress           | string    | Branch routing address like branch-id-765 or you can leave it empty         |
| IsBeneficiary                       | bool      | Must be set to true in order to send payments to this counterparty          |
| Bespoke                             | []Bespoke | It support list of key-value, you can add it to the counterarty.            |
| Metadata                            | Metadata  | Metadata about the counter party                                            |

##### Objects

###### Bespoke

| Name  | Type   | Description                                        |
|-------|--------|----------------------------------------------------|
| Key   | string | any info-key you want to add to this counerparty   |
| Value | string | any info-value you want to add to this counerparty |

###### Metadata

| Name              | Type     | Description |
|-------------------|----------|-------------|
| PublicAlias       | string   |             |
| MoreInfo          | string   |             |
| URL               | string   |             |
| ImageURL          | string   |             |
| OpenCorporatesURL | string   |             |
| CorporateLocation | Location |             |
| PhysicalLocation  | Location |             |
| PrivateAlias      | string   |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

## Retrieve an account with other account id

Retrieves other account (counter party) by other account id

```sh
curl -X GET \
	/v1/accounts/{AccountID}/otheraccounts/{OtherAccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getotheraccountbyid []}}

### HTTP Request

`GET /v1/accounts/{AccountID}/otheraccounts/{OtherAccountID}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| AccountID      | string |             |
| OtherAccountID | string |             |

### Responses

#### Response body

| Name         | Type         | Description |
|--------------|--------------|-------------|
| OtherAccount | OtherAccount |             |

##### Objects

###### OtherAccount

| Name            | Type              | Description |
|-----------------|-------------------|-------------|
| ID              | string            |             |
| Holder          | Holder            |             |
| BankRouting     | BankRouting       |             |
| AccountRoutings | []AccountRoutings |             |
| Metadata        | Metadata          |             |

###### Holder

| Name    | Type   | Description |
|---------|--------|-------------|
| Name    | string |             |
| IsAlias | bool   |             |

###### BankRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| Scheme  | string |             |
| Address | string |             |

###### AccountRoutings

| Name    | Type   | Description |
|---------|--------|-------------|
| Scheme  | string |             |
| Address | string |             |

###### Metadata

| Name              | Type     | Description |
|-------------------|----------|-------------|
| PublicAlias       | string   |             |
| MoreInfo          | string   |             |
| URL               | string   |             |
| ImageURL          | string   |             |
| OpenCorporatesURL | string   |             |
| CorporateLocation | Location |             |
| PhysicalLocation  | Location |             |
| PrivateAlias      | string   |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

## Retrieve all other account with account id

Retrieves all other accounts (counter party) for an account id

```sh
curl -X GET \
	/v1/accounts/{AccountID}/otheraccounts \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getotheraccounts []}}

### HTTP Request

`GET /v1/accounts/{AccountID}/otheraccounts`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| AccountID | string |             |

### Responses

#### Response body

| Name          | Type           | Description |
|---------------|----------------|-------------|
| OtherAccounts | []OtherAccount |             |

##### Objects

###### OtherAccount

| Name            | Type              | Description |
|-----------------|-------------------|-------------|
| ID              | string            |             |
| Holder          | Holder            |             |
| BankRouting     | BankRouting       |             |
| AccountRoutings | []AccountRoutings |             |
| Metadata        | Metadata          |             |

###### Holder

| Name    | Type   | Description |
|---------|--------|-------------|
| Name    | string |             |
| IsAlias | bool   |             |

###### BankRouting

| Name    | Type   | Description |
|---------|--------|-------------|
| Scheme  | string |             |
| Address | string |             |

###### AccountRoutings

| Name    | Type   | Description |
|---------|--------|-------------|
| Scheme  | string |             |
| Address | string |             |

###### Metadata

| Name              | Type     | Description |
|-------------------|----------|-------------|
| PublicAlias       | string   |             |
| MoreInfo          | string   |             |
| URL               | string   |             |
| ImageURL          | string   |             |
| OpenCorporatesURL | string   |             |
| CorporateLocation | Location |             |
| PhysicalLocation  | Location |             |
| PrivateAlias      | string   |             |

###### Location

| Name      | Type   | Description                                                        |
|-----------|--------|--------------------------------------------------------------------|
| Latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].   |
| Longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0] |

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

## Annex

####  TransactionRequestType

Entity type defines the type of counterparty

| Value                         | Description                                  |
|-------------------------------|----------------------------------------------|
| UnknownTransactionRequestType |                                              |
| COUNTERPARTY                  | Counter party type                           |
| SEPA                          | TransactionRequestType_SEPA transaction type |
