# Transactions API v1.0.0

Provides create and read operations on the transaction resource.

* Host ``

* Base Path ``

## Approve a pending transaction

Approve a pending transaction.

```sh
curl -X POST \
	/v1/transactions/approval \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"transaction_id": "string",
		"approve": "bool",
		"tfa_type": "TFAType"
	}'
```
### HTTP Request

`POST /v1/transactions/approval`

### Body Parameters

| Name          | Type    | Description                                         |
|---------------|---------|-----------------------------------------------------|
| TransactionID | string  | TransactionID is the transaction unique identifier. |
| Approve       | bool    | Approve is Boolean value of approval action.        |
| TFAType       | TFAType | TFAType is type to receive OTP authentication code. |

### Responses

#### Response body

| Name            | Type   | Description                                                                        |
|-----------------|--------|------------------------------------------------------------------------------------|
| AuthorizationID | string | AuthorizationID is the executable code is obtained fromthe payment feedback result |
| SMSCode         | int64  | SMSCode is the OTP code used for testing.                                          |

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
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a transaction

Creates a new transaction and returns its id.

```sh
curl -X POST \
	/v1/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"source_account_id": "string",
		"destination_account_id": "string",
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"remarks": "string"
	}'
```
### HTTP Request

`POST /v1/transactions`

### Body Parameters

| Name                 | Type   | Description                                                                      |
|----------------------|--------|----------------------------------------------------------------------------------|
| SourceAccountID      | string | SourceAccountID is the identifier of the account emitting the transaction.       |
| DestinationAccountID | string | DestinationAccountID is the identifier of the account receiving the transaction. |
| Amount               | Amount | Amount holds the amount value and currency of the transaction.                   |
| Remarks              | string | Remarks is an informational note about the transaction.                          |

##### Objects

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name          | Type   | Description                                              |
|---------------|--------|----------------------------------------------------------|
| TransactionID | string | TransactionID is the unique identifier of a transaction. |
| CratedAt      | string | CreatedAt is the transaction created date.               |

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
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve a transaction

Retrieves all data from a transaction, selected by the transaction_id you supplied.

```sh
curl -X GET \
	/v1/transactions/{TransactionID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/transactions/{TransactionID}`

### Query Parameters

| Name          | Type   | Description                                              |
|---------------|--------|----------------------------------------------------------|
| TransactionID | string | TransactionID is the unique identifier of a transaction. |

### Responses

#### Response body

| Name               | Type            | Description                                                    |
|--------------------|-----------------|----------------------------------------------------------------|
| TransactionID      | string          | TransactionID is the unique identifier of a transaction.       |
| SourceAccount      | BankAccountInfo | SourceAccount is the account emitting the transaction.         |
| DestinationAccount | BankAccountInfo | DestinationAccount is the account receiving the transaction.   |
| Date               | Timestamp       | Date is the date of the transaction.                           |
| Type               | Type            | Type is the type of transaction.                               |
| Status             | Status          | Status is the status of the transaction.                       |
| Amount             | Amount          | Amount holds the amount value and currency of the transaction. |
| Description        | string          | Description describes the transaction.                         |
| UserID             | string          | UserID is the identifier of the issuer of the transaction.     |
| Remarks            | string          | Remarks is an informational note about the transaction.        |

##### Objects

###### BankAccountInfo

| Name      | Type      | Description                                          |
|-----------|-----------|------------------------------------------------------|
| AccountID | string    | AccountID is the identifier of the account.          |
| BankCode  | BankCode  | BankCode is code of the bank the account belongs to. |
| OwnerName | string    | OwnerName is the name of the owner of the account.   |
| MajorType | MajorType | MajorType is the type of account.                    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

Example:

```json
{
  "transaction_id": "string",
  "source_account_id": {
    "account_id": "string",
    "bank_code": "BankCode",
    "owner_name": "string",
    "major_type": "MajorType"
  },
  "destination_account_id": {
    "account_id": "string",
    "bank_code": "BankCode",
    "owner_name": "string",
    "major_type": "MajorType"
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
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all transactions

Returns a list containing up to 20 transactions. You can paginate through transactions by supplying nextstartingindex in your subsequents calls. nextstartingindex contains the transactionid of the last transactionid of the current page.

```sh
curl -X GET \
	/v1/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/transactions`

### Responses

#### Response body

| Name               | Type          | Description                                                 |
|--------------------|---------------|-------------------------------------------------------------|
| Result             | []Transaction | Result is a list containing up to 20 transactions.          |
| HasMore            | bool          | HasMore indicates if there are more transactions available. |
| LastRunningBalance | Amount        | LastRunningBalance is current balance for related account.  |

##### Objects

###### Transaction

| Name               | Type            | Description                                                    |
|--------------------|-----------------|----------------------------------------------------------------|
| TransactionID      | string          | TransactionID is the unique identifier of a transaction.       |
| SourceAccount      | BankAccountInfo | SourceAccount is the account emitting the transaction.         |
| DestinationAccount | BankAccountInfo | DestinationAccount is the account receiving the transaction.   |
| Date               | Timestamp       | Date is the date of the transaction.                           |
| Type               | Type            | Type is the type of transaction.                               |
| Status             | Status          | Status is the status of the transaction.                       |
| Amount             | Amount          | Amount holds the amount value and currency of the transaction. |
| Description        | string          | Description describes the transaction.                         |
| UserID             | string          | UserID is the identifier of the issuer of the transaction.     |
| Remarks            | string          | Remarks is an informational note about the transaction.        |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### BankAccountInfo

| Name      | Type      | Description                                          |
|-----------|-----------|------------------------------------------------------|
| AccountID | string    | AccountID is the identifier of the account.          |
| BankCode  | BankCode  | BankCode is code of the bank the account belongs to. |
| OwnerName | string    | OwnerName is the name of the owner of the account.   |
| MajorType | MajorType | MajorType is the type of account.                    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "result": [
    {
      "transaction_id": "string",
      "source_account_id": {
        "account_id": "string",
        "bank_code": "BankCode",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "destination_account_id": {
        "account_id": "string",
        "bank_code": "BankCode",
        "owner_name": "string",
        "major_type": "MajorType"
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
  "has_more": "bool",
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
| 500    | Returned when an unexpected error occured on the server side.                          |

## List all transactions for given account id

Returns a list containing up to 20 transactions. You can paginate through transactions by supplying nextstartingindex in your subsequents calls. nextstartingindex contains the transactionid of the last transactionid of the current page.

```sh
curl -X GET \
	/v1/accounts/{AccountID}/transactions \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet gettransactionsby_account []}}

### HTTP Request

`GET /v1/accounts/{AccountID}/transactions`

### Query Parameters

| Name      | Type   | Description                                       |
|-----------|--------|---------------------------------------------------|
| AccountID | string | AccountID is the unique identifier of an account. |

### Responses

#### Response body

| Name               | Type          | Description                                                 |
|--------------------|---------------|-------------------------------------------------------------|
| Result             | []Transaction | Result is a list containing up to 20 transactions.          |
| HasMore            | bool          | HasMore indicates if there are more transactions available. |
| LastRunningBalance | Amount        | LastRunningBalance is current balance for related account.  |

##### Objects

###### Transaction

| Name               | Type            | Description                                                    |
|--------------------|-----------------|----------------------------------------------------------------|
| TransactionID      | string          | TransactionID is the unique identifier of a transaction.       |
| SourceAccount      | BankAccountInfo | SourceAccount is the account emitting the transaction.         |
| DestinationAccount | BankAccountInfo | DestinationAccount is the account receiving the transaction.   |
| Date               | Timestamp       | Date is the date of the transaction.                           |
| Type               | Type            | Type is the type of transaction.                               |
| Status             | Status          | Status is the status of the transaction.                       |
| Amount             | Amount          | Amount holds the amount value and currency of the transaction. |
| Description        | string          | Description describes the transaction.                         |
| UserID             | string          | UserID is the identifier of the issuer of the transaction.     |
| Remarks            | string          | Remarks is an informational note about the transaction.        |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### BankAccountInfo

| Name      | Type      | Description                                          |
|-----------|-----------|------------------------------------------------------|
| AccountID | string    | AccountID is the identifier of the account.          |
| BankCode  | BankCode  | BankCode is code of the bank the account belongs to. |
| OwnerName | string    | OwnerName is the name of the owner of the account.   |
| MajorType | MajorType | MajorType is the type of account.                    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "result": [
    {
      "transaction_id": "string",
      "source_account_id": {
        "account_id": "string",
        "bank_code": "BankCode",
        "owner_name": "string",
        "major_type": "MajorType"
      },
      "destination_account_id": {
        "account_id": "string",
        "bank_code": "BankCode",
        "owner_name": "string",
        "major_type": "MajorType"
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
  "has_more": "bool",
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
| 500    | Returned when an unexpected error occured on the server side.                          |

## Authorize a transaction with 2FA

Authorization allows execution of transactions with 2-factor authentication (2FA).

```sh
curl -X POST \
	/v1/transactions/confirmation \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"transaction_id": "string",
		"code": "string",
		"authorization_id": "string",
		"tfa_type": "TFAType"
	}'
```
{{snippet tfa []}}

### HTTP Request

`POST /v1/transactions/confirmation`

### Body Parameters

| Name            | Type    | Description                                                                          |
|-----------------|---------|--------------------------------------------------------------------------------------|
| TransactionID   | string  | TransactionID is transaction / payment identification code requires approval.        |
| Code            | string  | Code is 2-digit authentication code is sent via SMS.                                 |
| AuthorizationID | string  | AuthorizationID is the executable code is obtained from the payment feedback result. |
| TFAType         | TFAType | TFAType is type to receive OTP authentication code.                                  |

### Responses

#### Response body

| Name        | Type   | Description                                    |
|-------------|--------|------------------------------------------------|
| TraceNumber | string | TraceNumber is the transaction reference code. |

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
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  Status

Status defines the status of a transaction.

| Value         | Description                                               |
|---------------|-----------------------------------------------------------|
| UnknownStatus |                                                           |
| Success       | Status_Success is the value for a successful transaction. |
| Pending       | Status_Pending is the value for a pending transaction.    |
| Rejected      | Status_Rejected is the value for a rejected transaction.  |

####  TFAType

TFAType is available type of TFA.

| Value   | Description                  |
|---------|------------------------------|
| SMS     | TFAType_SMS Message.         |
| SAFEKEY | TFAType_SAFEKEY Application. |

####  Type

Type defines the type of a transaction.

| Value       | Description                                        |
|-------------|----------------------------------------------------|
| UnknownType |                                                    |
| Credit      | Type_Credit is the value for a credit transaction. |
| Debit       | Type_Debit is the value for a debit transaction.   |
