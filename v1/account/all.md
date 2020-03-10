Accounts API v1.0.0
===================

Provides CRUD operations on the accounts resource.

* Host `https://`

* Base Path ``

Verify account existence {#method-get-checkaccount}
---------------------------------------------------

Verify whether or not an account exists.

```sh
curl -X GET \
	https:///v1/accounts/{AccountID}/check \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}/check`

### Query Parameters

| Name       | Type   | Description                                 |
|------------|--------|---------------------------------------------|
| account_id | string | AccountID is the account unique identifier. |

### Responses

#### Response body

| Name         | Type      | Description                                   |
|--------------|-----------|-----------------------------------------------|
| account_id   | string    | AccountID is the account unique identifier.   |
| bank_code    | string    | BankCode is the code that is related to bank. |
| account_name | string    | AccountName is the owner name of the account. |
| major_type   | MajorType | MajorType is the type of account.             |

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

Create an account {#method-post-createaccount}
----------------------------------------------

Creates a new account

```sh
curl -X POST \
	https:///v1/accounts \
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

`POST https:///v1/accounts`

### Body Parameters

| Name              | Type             | Description                                                             |
|-------------------|------------------|-------------------------------------------------------------------------|
| account_id        | string           | AccountID is the identifier of the account.                             |
| description       | string           | Description about the account                                           |
| account_roles     | \[]AccountRole   | Roles for the account                                                   |
| branch            | string           | Branch is the branch code of the associated branch.                     |
| customer          | string           | Customer                                                                |
| debit_transaction | DebitTransaction | DebitTransaction debited to account                                     |
| interest_rate     | string           | Rate of interest for an account                                         |
| major_type        | MajorType        | MajorType is the type of the account.                                   |
| maturity_date     | Timestamp        | The maturity date is the date on which the principal amount becomes due |
| minor             | string           | Minor account                                                           |

##### Objects

###### AccountRole

| Name          | Type       | Description                                      |
|---------------|------------|--------------------------------------------------|
| entity_number | string     | The identification number assigned to an account |
| entity_type   | EntityType | Type of entity                                   |
| role          | string     | Name of the Role                                 |

###### DebitTransaction

| Name          | Type   | Description                                                                     |
|---------------|--------|---------------------------------------------------------------------------------|
| amount        | string | Amount is the amount debited.                                                   |
| debit_account | string | DebitAccount is the account number to debit from.                               |
| exchange_rate | string | Exchange rate is the exchange rate for the transaction (if applicable).         |
| is_fx         | bool   | IsFx is a boolean indicating whether the transaction required foreign exchange. |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name                      | Type      | Description                                                                |
|---------------------------|-----------|----------------------------------------------------------------------------|
| account_id                | string    | AccountID is the unique identifier of the newly created account.           |
| credit_transaction_number | string    | CreditTransactionNumber is the transaction number of the credited account. |
| debit_transaction_number  | string    | DebitTransactionNumber is the transaction number of the debited account.   |
| major_type                | MajorType | MajorType is the type of the account.                                      |
| minor                     | string    | Minor account                                                              |

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

Delete an account {#method-delete-deleteaccount}
------------------------------------------------

Permanently delete an account.

```sh
curl -X DELETE \
	https:///v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/accounts/{AccountID}`

### Query Parameters

| Name       | Type   | Description                                 |
|------------|--------|---------------------------------------------|
| account_id | string | AccountID is the account unique identifier. |

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

Retrieve an account {#method-get-getaccount}
--------------------------------------------

Retrieves all data from an account selected by the supplied account_id.

```sh
curl -X GET \
	https:///v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}`

### Query Parameters

| Name       | Type   | Description |
|------------|--------|-------------|
| account_id | string |             |

### Responses

#### Response body

| Name                              | Type          | Description                                                                           |
|-----------------------------------|---------------|---------------------------------------------------------------------------------------|
| account_id                        | string        | AccountID is the unique identifier of an account.                                     |
| branch                            | string        | Branch is the branch code for the branch associated with the account.                 |
| branch_name                       | string        | BranchName is the long-form name of the branch associated with the account.           |
| status                            | string        | Status is the status of the account.                                                  |
| accrued_interest_at_maturity_date | Timestamp     | Interest accrues at an annual rate of interest that is fixed                          |
| amount_due                        | Amount        | Specify when payments are due on money borrowed                                       |
| available_balance                 | Amount        | AvailableBalance is the available balance of the account.                             |
| available_credit_limit            | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| checking_interest_rate            | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| contract_date                     | Timestamp     | ContractDate is the date of the contract initialization.                              |
| credit_limit                      | string        | CreditLimit is the allowed credit limit.                                              |
| current_accrued_interest          | string        | Interest earned but not received                                                      |
| current_balance                   | Amount        | CurrentBalance is the current balance of the account.                                 |
| current_term                      | string        | CurrentTerm is the account validity period.                                           |
| due_date                          | Timestamp     | DueDate is the loan maturity date.                                                    |
| interest_rate                     | string        | InterestRate is the interest rate for the account.                                    |
| major_type                        | MajorType     | MajorType is the account type.                                                        |
| major_category                    | MajorCategory | MajorCategory is the account category.                                                |
| maturity_date                     | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| next_payment_due_date             | Timestamp     | Specify when payments are due on money borrowed                                       |
| owner_name                        | string        | OwnerName is the name of the account's owner.                                         |
| start_date                        | Timestamp     | Account opening date                                                                  |

##### Objects

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

Retrieve an account's status {#method-get-getaccountstatus}
-----------------------------------------------------------

Retrieves status of the account, selected by the account_id you supplied.

```sh
curl -X GET \
	https:///v1/accounts/{AccountID}/status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts/{AccountID}/status`

### Query Parameters

| Name       | Type   | Description                                        |
|------------|--------|----------------------------------------------------|
| account_id | string | AccountID is the unique identifier of the account. |

### Responses

#### Response body

| Name               | Type      | Description                                  |
|--------------------|-----------|----------------------------------------------|
| account_id         | string    | AccountID is the account unique identifier.  |
| effetive_date      | Timestamp | EffectiveDate is the date of the request.    |
| status             | string    | Status is the current status of the account. |
| status_description | string    | StatusDescription describe about the Status. |

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

List all accounts {#method-get-getaccounts}
-------------------------------------------

Returns a list containing up to 20 accounts. `next_starting_index` can be used for pagination.

```sh
curl -X GET \
	https:///v1/accounts \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/accounts`

### Responses

#### Response body

| Name   | Type       | Description                           |
|--------|------------|---------------------------------------|
| result | \[]Account | Result is the paginated query result. |

##### Objects

###### Account

| Name                              | Type          | Description                                                                           |
|-----------------------------------|---------------|---------------------------------------------------------------------------------------|
| account_id                        | string        | AccountID is the unique identifier of an account.                                     |
| branch                            | string        | Branch is the branch code for the branch associated with the account.                 |
| branch_name                       | string        | BranchName is the long-form name of the branch associated with the account.           |
| status                            | string        | Status is the status of the account.                                                  |
| accrued_interest_at_maturity_date | Timestamp     | Interest accrues at an annual rate of interest that is fixed                          |
| amount_due                        | Amount        | Specify when payments are due on money borrowed                                       |
| available_balance                 | Amount        | AvailableBalance is the available balance of the account.                             |
| available_credit_limit            | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| checking_interest_rate            | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| contract_date                     | Timestamp     | ContractDate is the date of the contract initialization.                              |
| credit_limit                      | string        | CreditLimit is the allowed credit limit.                                              |
| current_accrued_interest          | string        | Interest earned but not received                                                      |
| current_balance                   | Amount        | CurrentBalance is the current balance of the account.                                 |
| current_term                      | string        | CurrentTerm is the account validity period.                                           |
| due_date                          | Timestamp     | DueDate is the loan maturity date.                                                    |
| interest_rate                     | string        | InterestRate is the interest rate for the account.                                    |
| major_type                        | MajorType     | MajorType is the account type.                                                        |
| major_category                    | MajorCategory | MajorCategory is the account category.                                                |
| maturity_date                     | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| next_payment_due_date             | Timestamp     | Specify when payments are due on money borrowed                                       |
| owner_name                        | string        | OwnerName is the name of the account's owner.                                         |
| start_date                        | Timestamp     | Account opening date                                                                  |

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

Update an account {#method-put-updateaccount}
---------------------------------------------

Updates an account with all the fields supplied.

```sh
curl -X PUT \
	https:///v1/accounts/{AccountID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"account_id": "string",
		"description": "string"
	}'
```

### HTTP Request

`PUT https:///v1/accounts/{AccountID}`

### Query Parameters

| Name       | Type   | Description                                                  |
|------------|--------|--------------------------------------------------------------|
| account_id | string | AccountID is the unique identifier of the account to update. |

### Body Parameters

| Name        | Type   | Description                                                  |
|-------------|--------|--------------------------------------------------------------|
| account_id  | string | AccountID is the unique identifier of the account to update. |
| description | string | Description to update                                        |

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

Update an account status {#method-put-updateaccountstatus}
----------------------------------------------------------

Updates the status of an account.

```sh
curl -X PUT \
	https:///v1/accounts/{AccountID}/status \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"account_id": "string",
		"status": "string"
	}'
```

### HTTP Request

`PUT https:///v1/accounts/{AccountID}/status`

### Query Parameters

| Name       | Type   | Description                                 |
|------------|--------|---------------------------------------------|
| account_id | string | AccountID is the account unique identifier. |

### Body Parameters

| Name       | Type   | Description                                 |
|------------|--------|---------------------------------------------|
| account_id | string | AccountID is the account unique identifier. |
| status     | string | Status is the status to be updated          |

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

Annex
-----

#### EntityType

EntityType describes the type of the entity.

| Value             | Description                                 |
|-------------------|---------------------------------------------|
| UnknownEntityType |                                             |
| Pers              | EntityType_Pers is an individual customer.  |
| Org               | EntityType_Org is an organisation customer. |
