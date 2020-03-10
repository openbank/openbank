# Accounts API v1.0.0

Provides CRUD operations on the accounts resource.

* Host ``

* Base Path ``

## Verify account existence {#method-get-checkaccount}

Verify whether or not an account exists.

```sh
curl -X GET \
	/v1/accounts/{AccountID}/check \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/accounts/{AccountID}/check`

### Query Parameters

| Name      | Type   | Description                                 |
|-----------|--------|---------------------------------------------|
| AccountID | string | AccountID is the account unique identifier. |

### Responses

#### Response body

| Name        | Type      | Description                                   |
|-------------|-----------|-----------------------------------------------|
| AccountID   | string    | AccountID is the account unique identifier.   |
| BankCode    | string    | BankCode is the code that is related to bank. |
| AccountName | string    | AccountName is the owner name of the account. |
| MajorType   | MajorType | MajorType is the type of account.             |

Example:

```json
{
  "account_id": "string",
  "bank_code": "string",
  "account_name": "string",
  "major_type": "MajorType"
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

## Create an account {#method-post-createaccount}

Creates a new account

```sh
curl -X POST \
	/v1/accounts \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"account_id": "string",
		"description": "string",
		"account_roles": [
			{
				"entity_number": "string",
				"entity_type": "EntityType",
				"role": "string"
			}
		],
		"branch": "string",
		"customer": "string",
		"debit_transaction": {
			"amount": "string",
			"debit_account": "string",
			"exchange_rate": "string",
			"is_fx": "bool"
		},
		"interest_rate": "string",
		"major_type": "MajorType",
		"maturity_date": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"minor": "string"
	}'
```
### HTTP Request

`POST /v1/accounts`

### Body Parameters

| Name             | Type             | Description                                                             |
|------------------|------------------|-------------------------------------------------------------------------|
| AccountID        | string           | AccountID is the identifier of the account.                             |
| Description      | string           | Description about the account                                           |
| AccountRoles     | []AccountRole    | Roles for the account                                                   |
| Branch           | string           | Branch is the branch code of the associated branch.                     |
| Customer         | string           | Customer                                                                |
| DebitTransaction | DebitTransaction | DebitTransaction debited to account                                     |
| InterestRate     | string           | Rate of interest for an account                                         |
| MajorType        | MajorType        | MajorType is the type of the account.                                   |
| MaturityDate     | Timestamp        | The maturity date is the date on which the principal amount becomes due |
| Minor            | string           | Minor account                                                           |

##### Objects

###### AccountRole

| Name         | Type       | Description                                      |
|--------------|------------|--------------------------------------------------|
| EntityNumber | string     | The identification number assigned to an account |
| EntityType   | EntityType | Type of entity                                   |
| Role         | string     | Name of the Role                                 |

###### DebitTransaction

| Name         | Type   | Description                                                                     |
|--------------|--------|---------------------------------------------------------------------------------|
| Amount       | string | Amount is the amount debited.                                                   |
| DebitAccount | string | DebitAccount is the account number to debit from.                               |
| ExchangeRate | string | Exchange rate is the exchange rate for the transaction (if applicable).         |
| IsFx         | bool   | IsFx is a boolean indicating whether the transaction required foreign exchange. |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name                    | Type      | Description                                                                |
|-------------------------|-----------|----------------------------------------------------------------------------|
| AccountID               | string    | AccountID is the unique identifier of the newly created account.           |
| CreditTransactionNumber | string    | CreditTransactionNumber is the transaction number of the credited account. |
| DebitTransactionNumber  | string    | DebitTransactionNumber is the transaction number of the debited account.   |
| MajorType               | MajorType | MajorType is the type of the account.                                      |
| Minor                   | string    | Minor account                                                              |

Example:

```json
{
  "account_id": "string",
  "credit_transaction_number": "string",
  "debit_transaction_number": "string",
  "major_type": "MajorType",
  "minor": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account created successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete an account {#method-delete-deleteaccount}

Permanently delete an account.

```sh
curl -X DELETE \
	/v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/accounts/{AccountID}`

### Query Parameters

| Name      | Type   | Description                                 |
|-----------|--------|---------------------------------------------|
| AccountID | string | AccountID is the account unique identifier. |

### Responses

#### Response body

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Account deleted successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve an account {#method-get-getaccount}

Retrieves all data from an account selected by the supplied account_id.

```sh
curl -X GET \
	/v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/accounts/{AccountID}`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| AccountID | string |             |

### Responses

#### Response body

| Name                          | Type          | Description                                                                           |
|-------------------------------|---------------|---------------------------------------------------------------------------------------|
| AccountID                     | string        | AccountID is the unique identifier of an account.                                     |
| Branch                        | string        | Branch is the branch code for the branch associated with the account.                 |
| BranchName                    | string        | BranchName is the long-form name of the branch associated with the account.           |
| Status                        | string        | Status is the status of the account.                                                  |
| AccruedInterestAtMaturityDate | Timestamp     | Interest accrues at an annual rate of interest that is fixed                          |
| AmountDue                     | Amount        | Specify when payments are due on money borrowed                                       |
| AvailableBalance              | Amount        | AvailableBalance is the available balance of the account.                             |
| AvailableCreditLimit          | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| CheckingInterestRate          | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| ContractDate                  | Timestamp     | ContractDate is the date of the contract initialization.                              |
| CreditLimit                   | string        | CreditLimit is the allowed credit limit.                                              |
| CurrentAccruedInterest        | string        | Interest earned but not received                                                      |
| CurrentBalance                | Amount        | CurrentBalance is the current balance of the account.                                 |
| CurrentTerm                   | string        | CurrentTerm is the account validity period.                                           |
| DueDate                       | Timestamp     | DueDate is the loan maturity date.                                                    |
| InterestRate                  | string        | InterestRate is the interest rate for the account.                                    |
| MajorType                     | MajorType     | MajorType is the account type.                                                        |
| MajorCategory                 | MajorCategory | MajorCategory is the account category.                                                |
| MaturityDate                  | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| NextPaymentDueDate            | Timestamp     | Specify when payments are due on money borrowed                                       |
| OwnerName                     | string        | OwnerName is the name of the account's owner.                                         |
| StartDate                     | Timestamp     | Account opening date                                                                  |

##### Objects

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
  "account_id": "string",
  "branch": "string",
  "branch_name": "string",
  "status": "string",
  "accrued_interest_at_maturity_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "amount_due": {
    "cur": "string",
    "num": "string"
  },
  "available_balance": {
    "cur": "string",
    "num": "string"
  },
  "available_credit_limit": "string",
  "checking_interest_rate": "string",
  "contract_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "credit_limit": "string",
  "current_accrued_interest": "string",
  "current_balance": {
    "cur": "string",
    "num": "string"
  },
  "current_term": "string",
  "due_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "interest_rate": "string",
  "major_type": "MajorType",
  "major_category": "MajorCategory",
  "maturity_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "next_payment_due_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "owner_name": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
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

## Retrieve an account's status {#method-get-getaccountstatus}

Retrieves status of the account, selected by the account_id you supplied.

```sh
curl -X GET \
	/v1/accounts/{AccountID}/status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getaccountstatus []}}

### HTTP Request

`GET /v1/accounts/{AccountID}/status`

### Query Parameters

| Name      | Type   | Description                                        |
|-----------|--------|----------------------------------------------------|
| AccountID | string | AccountID is the unique identifier of the account. |

### Responses

#### Response body

| Name              | Type      | Description                                  |
|-------------------|-----------|----------------------------------------------|
| AccountID         | string    | AccountID is the account unique identifier.  |
| EffectiveDate     | Timestamp | EffectiveDate is the date of the request.    |
| Status            | string    | Status is the current status of the account. |
| StatusDescription | string    | StatusDescription describe about the Status. |

##### Objects

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "account_id": "string",
  "effetive_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "status": "string",
  "status_description": "string"
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

## List all accounts {#method-get-getaccounts}

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	/v1/accounts \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/accounts`

### Responses

#### Response body

| Name   | Type      | Description                           |
|--------|-----------|---------------------------------------|
| Result | []Account | Result is the paginated query result. |

##### Objects

###### Account

| Name                          | Type          | Description                                                                           |
|-------------------------------|---------------|---------------------------------------------------------------------------------------|
| AccountID                     | string        | AccountID is the unique identifier of an account.                                     |
| Branch                        | string        | Branch is the branch code for the branch associated with the account.                 |
| BranchName                    | string        | BranchName is the long-form name of the branch associated with the account.           |
| Status                        | string        | Status is the status of the account.                                                  |
| AccruedInterestAtMaturityDate | Timestamp     | Interest accrues at an annual rate of interest that is fixed                          |
| AmountDue                     | Amount        | Specify when payments are due on money borrowed                                       |
| AvailableBalance              | Amount        | AvailableBalance is the available balance of the account.                             |
| AvailableCreditLimit          | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| CheckingInterestRate          | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| ContractDate                  | Timestamp     | ContractDate is the date of the contract initialization.                              |
| CreditLimit                   | string        | CreditLimit is the allowed credit limit.                                              |
| CurrentAccruedInterest        | string        | Interest earned but not received                                                      |
| CurrentBalance                | Amount        | CurrentBalance is the current balance of the account.                                 |
| CurrentTerm                   | string        | CurrentTerm is the account validity period.                                           |
| DueDate                       | Timestamp     | DueDate is the loan maturity date.                                                    |
| InterestRate                  | string        | InterestRate is the interest rate for the account.                                    |
| MajorType                     | MajorType     | MajorType is the account type.                                                        |
| MajorCategory                 | MajorCategory | MajorCategory is the account category.                                                |
| MaturityDate                  | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| NextPaymentDueDate            | Timestamp     | Specify when payments are due on money borrowed                                       |
| OwnerName                     | string        | OwnerName is the name of the account's owner.                                         |
| StartDate                     | Timestamp     | Account opening date                                                                  |

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
  "result": [
    {
      "account_id": "string",
      "branch": "string",
      "branch_name": "string",
      "status": "string",
      "accrued_interest_at_maturity_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "amount_due": {
        "cur": "string",
        "num": "string"
      },
      "available_balance": {
        "cur": "string",
        "num": "string"
      },
      "available_credit_limit": "string",
      "checking_interest_rate": "string",
      "contract_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "credit_limit": "string",
      "current_accrued_interest": "string",
      "current_balance": {
        "cur": "string",
        "num": "string"
      },
      "current_term": "string",
      "due_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "interest_rate": "string",
      "major_type": "MajorType",
      "major_category": "MajorCategory",
      "maturity_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "next_payment_due_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "owner_name": "string",
      "start_date": {
        "seconds": "int64",
        "nanos": "int32"
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

## Update an account {#method-put-updateaccount}

Updates an account with all the fields supplied.

```sh
curl -X PUT \
	/v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"account_id": "string",
		"description": "string"
	}'
```
### HTTP Request

`PUT /v1/accounts/{AccountID}`

### Query Parameters

| Name      | Type   | Description                                                  |
|-----------|--------|--------------------------------------------------------------|
| AccountID | string | AccountID is the unique identifier of the account to update. |

### Body Parameters

| Name        | Type   | Description                                                  |
|-------------|--------|--------------------------------------------------------------|
| AccountID   | string | AccountID is the unique identifier of the account to update. |
| Description | string | Description to update                                        |

### Responses

#### Response body

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Account updated successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update an account status {#method-put-updateaccountstatus}

Updates the status of an account.

```sh
curl -X PUT \
	/v1/accounts/{AccountID}/status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"account_id": "string",
		"status": "string"
	}'
```
{{snippet updateaccountstatus []}}

### HTTP Request

`PUT /v1/accounts/{AccountID}/status`

### Query Parameters

| Name      | Type   | Description                                 |
|-----------|--------|---------------------------------------------|
| AccountID | string | AccountID is the account unique identifier. |

### Body Parameters

| Name      | Type   | Description                                 |
|-----------|--------|---------------------------------------------|
| AccountID | string | AccountID is the account unique identifier. |
| Status    | string | Status is the status to be updated          |

### Responses

#### Response body

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | AccountStatus updated successfully.                                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  EntityType

EntityType describes the type of the entity.

| Value             | Description                                 |
|-------------------|---------------------------------------------|
| UnknownEntityType |                                             |
| Pers              | EntityType_Pers is an individual customer.  |
| Org               | EntityType_Org is an organisation customer. |
