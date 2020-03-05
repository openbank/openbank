Transactions API v1.0.0
=======================

Provides create and read operations on the transaction resource.

* Host `https://`

* Base Path ``

Approve a pending transaction {#method-post-approvepayment}
-----------------------------------------------------------

Approve a pending transaction.

```sh
curl -X POST \
	https:///v1/transactions/approval \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"transaction_id": "string",
		"approve": "bool",
		"tfa_type": "TFAType"
	}'
```

### HTTP Request

`POST https:///v1/transactions/approval`

### Body Parameters

| Name           | Type    | Description                                         |
|----------------|---------|-----------------------------------------------------|
| transaction_id | string  | TransactionID is the transaction unique identifier. |
| approve        | bool    | Approve is Boolean value of approval action.        |
| tfa_type       | TFAType | TFAType is type to receive OTP authentication code. |

### Responses

#### Response body

| Name             | Type   | Description                                                                        |
|------------------|--------|------------------------------------------------------------------------------------|
| authorization_id | string | AuthorizationID is the executable code is obtained fromthe payment feedback result |
| sms_code         | int64  | SMSCode is the OTP code used for testing.                                          |

Example:

```json
{
  "authorization_id": "string",
  "sms_code": "int64"
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Transaction successfully approved.                                                     |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a transaction {#method-post-createtransaction}
-----------------------------------------------------

Creates a new transaction and returns its id.

```sh
curl -X POST \
	https:///v1/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"source_account_id": "string",
		"source_offline_user": {
			"user_id": "string",
			"first_name": "string",
			"middle_name": "string",
			"last_name": "string",
			"mobile_no": "string",
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"destination_account": {
			"account_id": "string",
			"bank_code": "string",
			"owner_name": "string",
			"major_type": "MajorType"
		},
		"destination_offline_user": {
			"user_id": "string",
			"first_name": "string",
			"middle_name": "string",
			"last_name": "string",
			"mobile_no": "string",
			"location": {
				"latitude": "double",
				"longitude": "double"
			}
		},
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"remarks": "string"
	}'
```

### HTTP Request

`POST https:///v1/transactions`

### Body Parameters

| Name                     | Type            | Description                                                                |
|--------------------------|-----------------|----------------------------------------------------------------------------|
| source_account_id        | string          | SourceAccountID is the identifier of the account emitting the transaction. |
| source_offline_user      | OfflineUserInfo | SourceOfflineUser is the contact information for an offline user.          |
| destination_account      | BankAccountInfo | DestinationAccount is the account receiving the transaction.               |
| destination_offline_user | OfflineUserInfo | DestinationOfflineUser is the contact information for an offline user.     |
| amount                   | Amount          | Amount holds the amount value and currency of the transaction.             |
| remarks                  | string          | Remarks is an informational note about the transaction.                    |

##### Objects

###### OfflineUserInfo

| Name        | Type     | Description                                                |
|-------------|----------|------------------------------------------------------------|
| user_id     | string   | UserID                                                     |
| first_name  | string   | FirstName of the person                                    |
| middle_name | string   | MiddleName or middle names (space separated) of the person |
| last_name   | string   | LastName or last names (space separated) of the person     |
| mobile_no   | string   | MobileNo contact of the person                             |
| location    | Location | Location is the physical location of the interaction.      |

###### BankAccountInfo

| Name       | Type      | Description                                          |
|------------|-----------|------------------------------------------------------|
| account_id | string    | AccountID is the identifier of the account.          |
| bank_code  | string    | BankCode is code of the bank the account belongs to. |
| owner_name | string    | OwnerName is the name of the owner of the account.   |
| major_type | MajorType | MajorType is the type of account.                    |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

### Responses

#### Response body

| Name           | Type   | Description                                              |
|----------------|--------|----------------------------------------------------------|
| transaction_id | string | TransactionID is the unique identifier of a transaction. |
| created_at     | string | CreatedAt is the transaction created date.               |

Example:

```json
{
  "transaction_id": "string",
  "created_at": "string"
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

Retrieve a transaction {#method-get-gettransaction}
---------------------------------------------------

Retrieves all data from a transaction, selected by the transaction_id you supplied.

```sh
curl -X GET \
	https:///v1/transactions/{TransactionID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/transactions/{TransactionID}`

### Query Parameters

| Name           | Type   | Description                                              |
|----------------|--------|----------------------------------------------------------|
| transaction_id | string | TransactionID is the unique identifier of a transaction. |

### Responses

#### Response body

| Name                     | Type            | Description                                                            |
|--------------------------|-----------------|------------------------------------------------------------------------|
| transaction_id           | string          | TransactionID is the unique identifier of a transaction.               |
| source_account           | BankAccountInfo | SourceAccount is the account emitting the transaction.                 |
| source_offline_user      | OfflineUserInfo | SourceOfflineUser is the contact information for an offline user.      |
| destination_account      | BankAccountInfo | DestinationAccount is the account receiving the transaction.           |
| destination_offline_user | OfflineUserInfo | DestinationOfflineUser is the contact information for an offline user. |
| date                     | Timestamp       | Date is the date of the transaction.                                   |
| type                     | Type            | Type is the type of transaction.                                       |
| status                   | Status          | Status is the status of the transaction.                               |
| amount                   | Amount          | Amount holds the amount value and currency of the transaction.         |
| description              | string          | Description describes the transaction.                                 |
| user_id                  | string          | UserID is the identifier of the issuer of the transaction.             |
| remarks                  | string          | Remarks is an informational note about the transaction.                |

##### Objects

###### BankAccountInfo

| Name       | Type      | Description                                          |
|------------|-----------|------------------------------------------------------|
| account_id | string    | AccountID is the identifier of the account.          |
| bank_code  | string    | BankCode is code of the bank the account belongs to. |
| owner_name | string    | OwnerName is the name of the owner of the account.   |
| major_type | MajorType | MajorType is the type of account.                    |

###### OfflineUserInfo

| Name        | Type     | Description                                                |
|-------------|----------|------------------------------------------------------------|
| user_id     | string   | UserID                                                     |
| first_name  | string   | FirstName of the person                                    |
| middle_name | string   | MiddleName or middle names (space separated) of the person |
| last_name   | string   | LastName or last names (space separated) of the person     |
| mobile_no   | string   | MobileNo contact of the person                             |
| location    | Location | Location is the physical location of the interaction.      |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

Example:

```json
{
  "transaction_id": "string",
  "source_account": {
    "account_id": "string",
    "bank_code": "string",
    "owner_name": "string",
    "major_type": "MajorType"
  },
  "source_offline_user": {
    "user_id": "string",
    "first_name": "string",
    "middle_name": "string",
    "last_name": "string",
    "mobile_no": "string",
    "location": {
      "latitude": "double",
      "longitude": "double"
    }
  },
  "destination_account": {
    "account_id": "string",
    "bank_code": "string",
    "owner_name": "string",
    "major_type": "MajorType"
  },
  "destination_offline_user": {
    "user_id": "string",
    "first_name": "string",
    "middle_name": "string",
    "last_name": "string",
    "mobile_no": "string",
    "location": {
      "latitude": "double",
      "longitude": "double"
    }
  },
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "type": "Type",
  "status": "Status",
  "amount": {
    "cur": "string",
    "num": "string"
  },
  "description": "string",
  "user_id": "string",
  "remarks": "string"
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

List all transactions {#method-get-gettransactions}
---------------------------------------------------

Returns a list containing up to 20 transactions. You can paginate through transactions by supplying next_starting_index in your subsequents calls. next_starting_index contains the transaction_id of the last transaction_id of the current page.

```sh
curl -X GET \
	https:///v1/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/transactions`

### Responses

#### Response body

| Name                 | Type           | Description                                                |
|----------------------|----------------|------------------------------------------------------------|
| result               | \[]Transaction | Result is the paginated query result.                      |
| last_running_balance | Amount         | LastRunningBalance is current balance for related account. |

##### Objects

###### Transaction

| Name                     | Type            | Description                                                            |
|--------------------------|-----------------|------------------------------------------------------------------------|
| transaction_id           | string          | TransactionID is the unique identifier of a transaction.               |
| source_account           | BankAccountInfo | SourceAccount is the account emitting the transaction.                 |
| source_offline_user      | OfflineUserInfo | SourceOfflineUser is the contact information for an offline user.      |
| destination_account      | BankAccountInfo | DestinationAccount is the account receiving the transaction.           |
| destination_offline_user | OfflineUserInfo | DestinationOfflineUser is the contact information for an offline user. |
| date                     | Timestamp       | Date is the date of the transaction.                                   |
| type                     | Type            | Type is the type of transaction.                                       |
| status                   | Status          | Status is the status of the transaction.                               |
| amount                   | Amount          | Amount holds the amount value and currency of the transaction.         |
| description              | string          | Description describes the transaction.                                 |
| user_id                  | string          | UserID is the identifier of the issuer of the transaction.             |
| remarks                  | string          | Remarks is an informational note about the transaction.                |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### BankAccountInfo

| Name       | Type      | Description                                          |
|------------|-----------|------------------------------------------------------|
| account_id | string    | AccountID is the identifier of the account.          |
| bank_code  | string    | BankCode is code of the bank the account belongs to. |
| owner_name | string    | OwnerName is the name of the owner of the account.   |
| major_type | MajorType | MajorType is the type of account.                    |

###### OfflineUserInfo

| Name        | Type     | Description                                                |
|-------------|----------|------------------------------------------------------------|
| user_id     | string   | UserID                                                     |
| first_name  | string   | FirstName of the person                                    |
| middle_name | string   | MiddleName or middle names (space separated) of the person |
| last_name   | string   | LastName or last names (space separated) of the person     |
| mobile_no   | string   | MobileNo contact of the person                             |
| location    | Location | Location is the physical location of the interaction.      |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

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
      "transaction_id": "string",
      "source_account": {
        "account_id": "string",
        "bank_code": "string",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "source_offline_user": {
        "user_id": "string",
        "first_name": "string",
        "middle_name": "string",
        "last_name": "string",
        "mobile_no": "string",
        "location": {
          "latitude": "double",
          "longitude": "double"
        }
      },
      "destination_account": {
        "account_id": "string",
        "bank_code": "string",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "destination_offline_user": {
        "user_id": "string",
        "first_name": "string",
        "middle_name": "string",
        "last_name": "string",
        "mobile_no": "string",
        "location": {
          "latitude": "double",
          "longitude": "double"
        }
      },
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "type": "Type",
      "status": "Status",
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "user_id": "string",
      "remarks": "string"
    }
  ],
  "last_running_balance": {
    "cur": "string",
    "num": "string"
  }
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

List all transactions for given account id {#method-get-gettransactionsbyaccount}
---------------------------------------------------------------------------------

Returns a list containing up to 20 transactions. You can paginate through transactions by supplying next_starting_index in your subsequents calls. next_starting_index contains the transaction_id of the last transaction_id of the current page.

```sh
curl -X GET \
	https:///v1/accounts/{AccountID}/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}/transactions`

### Query Parameters

| Name       | Type   | Description                                       |
|------------|--------|---------------------------------------------------|
| account_id | string | AccountID is the unique identifier of an account. |

### Responses

#### Response body

| Name                 | Type           | Description                                                |
|----------------------|----------------|------------------------------------------------------------|
| result               | \[]Transaction | Result is a list containing up to 20 transactions.         |
| last_running_balance | Amount         | LastRunningBalance is current balance for related account. |

##### Objects

###### Transaction

| Name                     | Type            | Description                                                            |
|--------------------------|-----------------|------------------------------------------------------------------------|
| transaction_id           | string          | TransactionID is the unique identifier of a transaction.               |
| source_account           | BankAccountInfo | SourceAccount is the account emitting the transaction.                 |
| source_offline_user      | OfflineUserInfo | SourceOfflineUser is the contact information for an offline user.      |
| destination_account      | BankAccountInfo | DestinationAccount is the account receiving the transaction.           |
| destination_offline_user | OfflineUserInfo | DestinationOfflineUser is the contact information for an offline user. |
| date                     | Timestamp       | Date is the date of the transaction.                                   |
| type                     | Type            | Type is the type of transaction.                                       |
| status                   | Status          | Status is the status of the transaction.                               |
| amount                   | Amount          | Amount holds the amount value and currency of the transaction.         |
| description              | string          | Description describes the transaction.                                 |
| user_id                  | string          | UserID is the identifier of the issuer of the transaction.             |
| remarks                  | string          | Remarks is an informational note about the transaction.                |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### BankAccountInfo

| Name       | Type      | Description                                          |
|------------|-----------|------------------------------------------------------|
| account_id | string    | AccountID is the identifier of the account.          |
| bank_code  | string    | BankCode is code of the bank the account belongs to. |
| owner_name | string    | OwnerName is the name of the owner of the account.   |
| major_type | MajorType | MajorType is the type of account.                    |

###### OfflineUserInfo

| Name        | Type     | Description                                                |
|-------------|----------|------------------------------------------------------------|
| user_id     | string   | UserID                                                     |
| first_name  | string   | FirstName of the person                                    |
| middle_name | string   | MiddleName or middle names (space separated) of the person |
| last_name   | string   | LastName or last names (space separated) of the person     |
| mobile_no   | string   | MobileNo contact of the person                             |
| location    | Location | Location is the physical location of the interaction.      |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

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
      "transaction_id": "string",
      "source_account": {
        "account_id": "string",
        "bank_code": "string",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "source_offline_user": {
        "user_id": "string",
        "first_name": "string",
        "middle_name": "string",
        "last_name": "string",
        "mobile_no": "string",
        "location": {
          "latitude": "double",
          "longitude": "double"
        }
      },
      "destination_account": {
        "account_id": "string",
        "bank_code": "string",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "destination_offline_user": {
        "user_id": "string",
        "first_name": "string",
        "middle_name": "string",
        "last_name": "string",
        "mobile_no": "string",
        "location": {
          "latitude": "double",
          "longitude": "double"
        }
      },
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "type": "Type",
      "status": "Status",
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "user_id": "string",
      "remarks": "string"
    }
  ],
  "last_running_balance": {
    "cur": "string",
    "num": "string"
  }
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

Authorize a transaction with 2FA {#method-post-tfa}
---------------------------------------------------

Authorization allows execution of transactions with 2-factor authentication (2FA).

```sh
curl -X POST \
	https:///v1/transactions/confirmation \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"transaction_id": "string",
		"code": "string",
		"authorization_id": "string",
		"tfa_type": "TFAType"
	}'
```

### HTTP Request

`POST https:///v1/transactions/confirmation`

### Body Parameters

| Name             | Type    | Description                                                                          |
|------------------|---------|--------------------------------------------------------------------------------------|
| transaction_id   | string  | TransactionID is transaction / payment identification code requires approval.        |
| code             | string  | Code is 2-digit authentication code is sent via SMS.                                 |
| authorization_id | string  | AuthorizationID is the executable code is obtained from the payment feedback result. |
| tfa_type         | TFAType | TFAType is type to receive OTP authentication code.                                  |

### Responses

#### Response body

| Name         | Type   | Description                                    |
|--------------|--------|------------------------------------------------|
| trace_number | string | TraceNumber is the transaction reference code. |

Example:

```json
{
  "trace_number": "string"
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Transaction authorized.                                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Annex
-----

#### Status

Status defines the status of a transaction.

| Value         | Description                                               |
|---------------|-----------------------------------------------------------|
| UnknownStatus |                                                           |
| Success       | Status_Success is the value for a successful transaction. |
| Pending       | Status_Pending is the value for a pending transaction.    |
| Rejected      | Status_Rejected is the value for a rejected transaction.  |

#### TFAType

TFAType is available type of TFA.

| Value   | Description                  |
|---------|------------------------------|
| SMS     | TFAType_SMS Message.         |
| SAFEKEY | TFAType_SAFEKEY Application. |

#### Type

Type defines the type of a transaction.

| Value       | Description                                        |
|-------------|----------------------------------------------------|
| UnknownType |                                                    |
| Credit      | Type_Credit is the value for a credit transaction. |
| Debit       | Type_Debit is the value for a debit transaction.   |
