# Transacton Request API v1.0.0

Provides CRUD operations on the transaction request resource.

* Host `https://`

* Base Path ``

## Answer the transaction reqeust challenge {#method-post-answertransactionrequestchallenge}

Answer the transaction request challenge

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/{TransactionRequestType}/transactionrequest/{TransactionRequestID}/challenge \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"transaction_request_type": "string",
		"transaction_request_id": "string",
		"id": "string",
		"answer": "string"
	}'
```

### HTTP Request {#http-request-method-post-answertransactionrequestchallenge}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/{TransactionRequestType}/transactionrequest/{TransactionRequestID}/challenge`

### Query Parameters {#query-parameters-method-post-answertransactionrequestchallenge}

| Name                     | Type   | Description |
|--------------------------|--------|-------------|
| bank_id                  | string |             |
| account_id               | string |             |
| transaction_request_type | string |             |
| transaction_request_id   | string |             |

### Body Parameters {#body-parameters-method-post-answertransactionrequestchallenge}

| Name                     | Type   | Description |
|--------------------------|--------|-------------|
| bank_id                  | string |             |
| account_id               | string |             |
| transaction_request_type | string |             |
| transaction_request_id   | string |             |
| id                       | string |             |
| answer                   | string |             |

### Responses {#responses-method-post-answertransactionrequestchallenge}

#### Response body {#response-body-method-post-answertransactionrequestchallenge}

| Name            | Type             | Description |
|-----------------|------------------|-------------|
| id              | string           |             |
| type            | string           |             |
| from            | BankAccount      |             |
| details         | ChallengeDetails |             |
| transaction_ids | string           |             |
| status          | string           |             |
| start_date      | Timestamp        |             |
| end_date        | Timestamp        |             |
| challenge       | Challenge        |             |
| charge          | Charge           |             |

##### Objects {#objects-AnswerTransactionRequestChallengeResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### ChallengeDetails

| Name        | Type        | Description |
|-------------|-------------|-------------|
| to          | BankAccount |             |
| amount      | Amount      |             |
| description | string      |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to": {
      "bank_id": "string",
      "account_id": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-answertransactionrequestchallenge}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Answered the transaction request challenge                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an account otp transaction request {#method-post-createaccountotptransaction}

Creates a new account otp transaction request

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT_OTP/transactionrequest \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"to_bank": {
			"bank_id": "string",
			"account_id": "string"
		},
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"description": "string"
	}'
```

### HTTP Request {#http-request-method-post-createaccountotptransaction}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT_OTP/transactionrequest`

### Query Parameters {#query-parameters-method-post-createaccountotptransaction}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Body Parameters {#body-parameters-method-post-createaccountotptransaction}

| Name        | Type        | Description |
|-------------|-------------|-------------|
| bank_id     | string      |             |
| account_id  | string      |             |
| to_bank     | BankAccount |             |
| amount      | Amount      |             |
| description | string      |             |

##### Objects {#objects-TransactionRequest}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-post-createaccountotptransaction}

#### Response body {#response-body-method-post-createaccountotptransaction}

| Name            | Type        | Description |
|-----------------|-------------|-------------|
| id              | string      |             |
| type            | string      |             |
| from            | BankAccount |             |
| details         | Details     |             |
| transaction_ids | \[]string   |             |
| status          | string      |             |
| start_date      | Timestamp   |             |
| end_date        | Timestamp   |             |
| challenge       | Challenge   |             |
| charge          | Charge      |             |

##### Objects {#objects-TransactionResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to_sandbox_tan": {
      "bank_id": "string",
      "account_id": "string"
    },
    "to_sepa": {
      "iban": "string"
    },
    "to_counterparty": {
      "counterparty_id": "string"
    },
    "to_transfer_to_phone": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_atm": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_account": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "transfer_type": "string",
      "future_date": "string",
      "to": {
        "legal_name": "string",
        "date_of_birth": "string",
        "mobile_phone_number": "string",
        "kyc_document": {
          "type": "string",
          "number": "string"
        }
      }
    },
    "to_sepa_credit_transfers": {
      "debtor_account": {
        "iban": "string"
      },
      "instructed_amount": {
        "cur": "string",
        "num": "string"
      },
      "creditor_account": {
        "iban": "string"
      },
      "creditor_name": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "[]string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-createaccountotptransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account OTP Transaction request created successfully.                                  |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an account transaction request {#method-post-createaccounttransaction}

Creates a new account transaction request

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT/transactionrequest \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"to_bank": {
			"bank_id": "string",
			"account_id": "string"
		},
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"description": "string"
	}'
```

### HTTP Request {#http-request-method-post-createaccounttransaction}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT/transactionrequest`

### Query Parameters {#query-parameters-method-post-createaccounttransaction}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Body Parameters {#body-parameters-method-post-createaccounttransaction}

| Name        | Type        | Description |
|-------------|-------------|-------------|
| bank_id     | string      |             |
| account_id  | string      |             |
| to_bank     | BankAccount |             |
| amount      | Amount      |             |
| description | string      |             |

##### Objects {#objects-TransactionRequest}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-post-createaccounttransaction}

#### Response body {#response-body-method-post-createaccounttransaction}

| Name            | Type        | Description |
|-----------------|-------------|-------------|
| id              | string      |             |
| type            | string      |             |
| from            | BankAccount |             |
| details         | Details     |             |
| transaction_ids | \[]string   |             |
| status          | string      |             |
| start_date      | Timestamp   |             |
| end_date        | Timestamp   |             |
| challenge       | Challenge   |             |
| charge          | Charge      |             |

##### Objects {#objects-TransactionResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to_sandbox_tan": {
      "bank_id": "string",
      "account_id": "string"
    },
    "to_sepa": {
      "iban": "string"
    },
    "to_counterparty": {
      "counterparty_id": "string"
    },
    "to_transfer_to_phone": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_atm": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_account": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "transfer_type": "string",
      "future_date": "string",
      "to": {
        "legal_name": "string",
        "date_of_birth": "string",
        "mobile_phone_number": "string",
        "kyc_document": {
          "type": "string",
          "number": "string"
        }
      }
    },
    "to_sepa_credit_transfers": {
      "debtor_account": {
        "iban": "string"
      },
      "instructed_amount": {
        "cur": "string",
        "num": "string"
      },
      "creditor_account": {
        "iban": "string"
      },
      "creditor_name": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "[]string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-createaccounttransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account Transaction request created successfully.                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an counter party transaction request {#method-post-createcounterpartytransaction}

Creates a new counter party transaction request

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/COUNTERPARTY/transactionrequest \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"to": {
			"counterparty_id": "string"
		},
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"description": "string",
		"charge_policy": "string",
		"future_date": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcounterpartytransaction}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/COUNTERPARTY/transactionrequest`

### Query Parameters {#query-parameters-method-post-createcounterpartytransaction}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Body Parameters {#body-parameters-method-post-createcounterpartytransaction}

| Name          | Type           | Description |
|---------------|----------------|-------------|
| bank_id       | string         |             |
| account_id    | string         |             |
| to            | ToCounterparty |             |
| amount        | Amount         |             |
| description   | string         |             |
| charge_policy | string         |             |
| future_date   | string         |             |

##### Objects {#objects-CounterPartyTransactionRequest}

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-post-createcounterpartytransaction}

#### Response body {#response-body-method-post-createcounterpartytransaction}

| Name            | Type        | Description |
|-----------------|-------------|-------------|
| id              | string      |             |
| type            | string      |             |
| from            | BankAccount |             |
| details         | Details     |             |
| transaction_ids | \[]string   |             |
| status          | string      |             |
| start_date      | Timestamp   |             |
| end_date        | Timestamp   |             |
| challenge       | Challenge   |             |
| charge          | Charge      |             |

##### Objects {#objects-TransactionResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to_sandbox_tan": {
      "bank_id": "string",
      "account_id": "string"
    },
    "to_sepa": {
      "iban": "string"
    },
    "to_counterparty": {
      "counterparty_id": "string"
    },
    "to_transfer_to_phone": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_atm": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_account": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "transfer_type": "string",
      "future_date": "string",
      "to": {
        "legal_name": "string",
        "date_of_birth": "string",
        "mobile_phone_number": "string",
        "kyc_document": {
          "type": "string",
          "number": "string"
        }
      }
    },
    "to_sepa_credit_transfers": {
      "debtor_account": {
        "iban": "string"
      },
      "instructed_amount": {
        "cur": "string",
        "num": "string"
      },
      "creditor_account": {
        "iban": "string"
      },
      "creditor_name": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "[]string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-createcounterpartytransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Counter Party Transaction request created successfully.                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an free form transaction request {#method-post-createfreeformtransaction}

Creates a new free form transaction request

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/FREE_FORM/transactionrequest \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"description": "string"
	}'
```

### HTTP Request {#http-request-method-post-createfreeformtransaction}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/FREE_FORM/transactionrequest`

### Query Parameters {#query-parameters-method-post-createfreeformtransaction}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Body Parameters {#body-parameters-method-post-createfreeformtransaction}

| Name        | Type   | Description |
|-------------|--------|-------------|
| bank_id     | string |             |
| account_id  | string |             |
| amount      | Amount |             |
| description | string |             |

##### Objects {#objects-FreeFormTransactionRequest}

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

### Responses {#responses-method-post-createfreeformtransaction}

#### Response body {#response-body-method-post-createfreeformtransaction}

| Name            | Type        | Description |
|-----------------|-------------|-------------|
| id              | string      |             |
| type            | string      |             |
| from            | BankAccount |             |
| details         | Details     |             |
| transaction_ids | \[]string   |             |
| status          | string      |             |
| start_date      | Timestamp   |             |
| end_date        | Timestamp   |             |
| challenge       | Challenge   |             |
| charge          | Charge      |             |

##### Objects {#objects-TransactionResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to_sandbox_tan": {
      "bank_id": "string",
      "account_id": "string"
    },
    "to_sepa": {
      "iban": "string"
    },
    "to_counterparty": {
      "counterparty_id": "string"
    },
    "to_transfer_to_phone": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_atm": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_account": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "transfer_type": "string",
      "future_date": "string",
      "to": {
        "legal_name": "string",
        "date_of_birth": "string",
        "mobile_phone_number": "string",
        "kyc_document": {
          "type": "string",
          "number": "string"
        }
      }
    },
    "to_sepa_credit_transfers": {
      "debtor_account": {
        "iban": "string"
      },
      "instructed_amount": {
        "cur": "string",
        "num": "string"
      },
      "creditor_account": {
        "iban": "string"
      },
      "creditor_name": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "[]string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-createfreeformtransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Free Form Transaction request created successfully.                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an sepa transaction request {#method-post-createsepatransaction}

Creates a new sepa transaction request

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/SEPA/transactionrequest \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"account_id": "string",
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"to": {
			"iban": "string"
		},
		"description": "string",
		"charge_policy": "string",
		"future_date": "string"
	}'
```

### HTTP Request {#http-request-method-post-createsepatransaction}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/SEPA/transactionrequest`

### Query Parameters {#query-parameters-method-post-createsepatransaction}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Body Parameters {#body-parameters-method-post-createsepatransaction}

| Name          | Type   | Description |
|---------------|--------|-------------|
| bank_id       | string |             |
| account_id    | string |             |
| amount        | Amount |             |
| to            | ToSepa |             |
| description   | string |             |
| charge_policy | string |             |
| future_date   | string |             |

##### Objects {#objects-SEPATransactionRequest}

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

### Responses {#responses-method-post-createsepatransaction}

#### Response body {#response-body-method-post-createsepatransaction}

| Name            | Type        | Description |
|-----------------|-------------|-------------|
| id              | string      |             |
| type            | string      |             |
| from            | BankAccount |             |
| details         | Details     |             |
| transaction_ids | \[]string   |             |
| status          | string      |             |
| start_date      | Timestamp   |             |
| end_date        | Timestamp   |             |
| challenge       | Challenge   |             |
| charge          | Charge      |             |

##### Objects {#objects-TransactionResponse}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "id": "string",
  "type": "string",
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "details": {
    "to_sandbox_tan": {
      "bank_id": "string",
      "account_id": "string"
    },
    "to_sepa": {
      "iban": "string"
    },
    "to_counterparty": {
      "counterparty_id": "string"
    },
    "to_transfer_to_phone": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_atm": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "message": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "to": {
        "mobile_phone_number": "string"
      }
    },
    "to_transfer_to_account": {
      "amount": {
        "cur": "string",
        "num": "string"
      },
      "description": "string",
      "transfer_type": "string",
      "future_date": "string",
      "to": {
        "legal_name": "string",
        "date_of_birth": "string",
        "mobile_phone_number": "string",
        "kyc_document": {
          "type": "string",
          "number": "string"
        }
      }
    },
    "to_sepa_credit_transfers": {
      "debtor_account": {
        "iban": "string"
      },
      "instructed_amount": {
        "cur": "string",
        "num": "string"
      },
      "creditor_account": {
        "iban": "string"
      },
      "creditor_name": "string"
    },
    "amount": {
      "cur": "string",
      "num": "string"
    },
    "description": "string"
  },
  "transaction_ids": "[]string",
  "status": "string",
  "start_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "end_date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "challenge": {
    "id": "string",
    "allowed_attempts": "int32",
    "challenge_type": "string"
  },
  "charge": {
    "summary": "string",
    "amount": {
      "cur": "string",
      "num": "string"
    }
  }
}
```

#### Response codes {#response-codes-method-post-createsepatransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | SEPA Transaction request created successfully.                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve transaction types {#method-get-getsupportedtransactionrequesttypes}

Retrieves list of transaction request types

```sh
curl -X GET \
	https:///v1/banks/{BankID}/transaction-request-types \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getsupportedtransactionrequesttypes}

`GET https:///v1/banks/{BankID}/transaction-request-types`

### Query Parameters {#query-parameters-method-get-getsupportedtransactionrequesttypes}

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses {#responses-method-get-getsupportedtransactionrequesttypes}

#### Response body {#response-body-method-get-getsupportedtransactionrequesttypes}

| Name                      | Type                      | Description |
|---------------------------|---------------------------|-------------|
| transaction_request_types | \[]TransactionRequestType |             |

##### Objects {#objects-GetSupportedTransactionRequestTypesResponse}

###### TransactionRequestType

| Name   | Type   | Description |
|--------|--------|-------------|
| value  | string |             |
| charge | Charge |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "transaction_request_types": [
    {
      "value": "string",
      "charge": {
        "summary": "string",
        "amount": {
          "cur": "string",
          "num": "string"
        }
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getsupportedtransactionrequesttypes}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve transaction types {#method-get-gettransactionrequesttypes}

Retrieves list of transaction request types for an account_id

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-gettransactionrequesttypes}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types`

### Query Parameters {#query-parameters-method-get-gettransactionrequesttypes}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Responses {#responses-method-get-gettransactionrequesttypes}

#### Response body {#response-body-method-get-gettransactionrequesttypes}

| Name                      | Type                      | Description |
|---------------------------|---------------------------|-------------|
| transaction_request_types | \[]TransactionRequestType |             |

##### Objects {#objects-GetTransactionRequestTypesResponse}

###### TransactionRequestType

| Name   | Type   | Description |
|--------|--------|-------------|
| value  | string |             |
| charge | Charge |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "transaction_request_types": [
    {
      "value": "string",
      "charge": {
        "summary": "string",
        "amount": {
          "cur": "string",
          "num": "string"
        }
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-gettransactionrequesttypes}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve transaction requests {#method-get-gettransactionrequests}

Retrieves list of transaction requests for an account_id

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactionrequests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-gettransactionrequests}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactionrequests`

### Query Parameters {#query-parameters-method-get-gettransactionrequests}

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

### Responses {#responses-method-get-gettransactionrequests}

#### Response body {#response-body-method-get-gettransactionrequests}

| Name                              | Type                              | Description |
|-----------------------------------|-----------------------------------|-------------|
| transaction_requests_with_charges | \[]TransactionRequestsWithCharges |             |

##### Objects {#objects-GetTransactionRequestsResponse}

###### TransactionRequestsWithCharges

| Name            | Type      | Description |
|-----------------|-----------|-------------|
| id              | string    |             |
| type            | string    |             |
| from            | FromPhone |             |
| details         | Details   |             |
| transaction_ids | \[]string |             |
| status          | string    |             |
| start_date      | Timestamp |             |
| end_date        | Timestamp |             |
| challenge       | Challenge |             |
| charge          | Charge    |             |

###### FromPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |
| nickname            | string |             |

###### Details

| Name                     | Type                  | Description |
|--------------------------|-----------------------|-------------|
| to_sandbox_tan           | BankAccount           |             |
| to_sepa                  | ToSepa                |             |
| to_counterparty          | ToCounterparty        |             |
| to_transfer_to_phone     | ToTransferToPhone     |             |
| to_transfer_to_atm       | ToTransferToAtm       |             |
| to_transfer_to_account   | ToTransferToAccount   |             |
| to_sepa_credit_transfers | ToSepaCreditTransfers |             |
| amount                   | Amount                |             |
| description              | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name             | Type   | Description |
|------------------|--------|-------------|
| id               | string |             |
| allowed_attempts | int32  |             |
| challenge_type   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| summary | string |             |
| amount  | Amount |             |

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### ToCounterparty

| Name            | Type   | Description |
|-----------------|--------|-------------|
| counterparty_id | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| amount      | Amount    |             |
| description | string    |             |
| message     | string    |             |
| from        | FromPhone |             |
| to          | ToPhone   |             |

###### ToTransferToAccount

| Name          | Type   | Description |
|---------------|--------|-------------|
| amount        | Amount |             |
| description   | string |             |
| transfer_type | string |             |
| future_date   | string |             |
| to            | To     |             |

###### ToSepaCreditTransfers

| Name              | Type            | Description |
|-------------------|-----------------|-------------|
| debtor_account    | DebtorAccount   |             |
| instructed_amount | Amount          |             |
| creditor_account  | CreditorAccount |             |
| creditor_name     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### ToPhone

| Name                | Type   | Description |
|---------------------|--------|-------------|
| mobile_phone_number | string |             |

###### To

| Name                | Type        | Description |
|---------------------|-------------|-------------|
| legal_name          | string      |             |
| date_of_birth       | string      |             |
| mobile_phone_number | string      |             |
| kyc_document        | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| type   | string |             |
| number | string |             |

Example:

```json
{
  "transaction_requests_with_charges": [
    {
      "id": "string",
      "type": "string",
      "from": {
        "mobile_phone_number": "string",
        "nickname": "string"
      },
      "details": {
        "to_sandbox_tan": {
          "bank_id": "string",
          "account_id": "string"
        },
        "to_sepa": {
          "iban": "string"
        },
        "to_counterparty": {
          "counterparty_id": "string"
        },
        "to_transfer_to_phone": {
          "amount": {
            "cur": "string",
            "num": "string"
          },
          "description": "string",
          "message": "string",
          "from": {
            "mobile_phone_number": "string",
            "nickname": "string"
          },
          "to": {
            "mobile_phone_number": "string"
          }
        },
        "to_transfer_to_atm": {
          "amount": {
            "cur": "string",
            "num": "string"
          },
          "description": "string",
          "message": "string",
          "from": {
            "mobile_phone_number": "string",
            "nickname": "string"
          },
          "to": {
            "mobile_phone_number": "string"
          }
        },
        "to_transfer_to_account": {
          "amount": {
            "cur": "string",
            "num": "string"
          },
          "description": "string",
          "transfer_type": "string",
          "future_date": "string",
          "to": {
            "legal_name": "string",
            "date_of_birth": "string",
            "mobile_phone_number": "string",
            "kyc_document": {
              "type": "string",
              "number": "string"
            }
          }
        },
        "to_sepa_credit_transfers": {
          "debtor_account": {
            "iban": "string"
          },
          "instructed_amount": {
            "cur": "string",
            "num": "string"
          },
          "creditor_account": {
            "iban": "string"
          },
          "creditor_name": "string"
        },
        "amount": {
          "cur": "string",
          "num": "string"
        },
        "description": "string"
      },
      "transaction_ids": "[]string",
      "status": "string",
      "start_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "end_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "challenge": {
        "id": "string",
        "allowed_attempts": "int32",
        "challenge_type": "string"
      },
      "charge": {
        "summary": "string",
        "amount": {
          "cur": "string",
          "num": "string"
        }
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-gettransactionrequests}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Import a historic transaction {#method-post-savehistorictransaction}

Import a historic transaction

```sh
curl -X POST \
	https:///v1/transactionrequest/import \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"from": {
			"bank_id": "string",
			"account_id": "string"
		},
		"to": {
			"bank_id": "string",
			"account_id": "string"
		},
		"amount": {
			"cur": "string",
			"num": "string"
		},
		"description": "string",
		"posted": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"completed": {
			"seconds": "int64",
			"nanos": "int32"
		},
		"transaction_request_type": "string",
		"charge_policy": "string"
	}'
```

### HTTP Request {#http-request-method-post-savehistorictransaction}

`POST https:///v1/transactionrequest/import`

### Body Parameters {#body-parameters-method-post-savehistorictransaction}

| Name                     | Type        | Description |
|--------------------------|-------------|-------------|
| from                     | BankAccount |             |
| to                       | BankAccount |             |
| amount                   | Amount      |             |
| description              | string      |             |
| posted                   | Timestamp   |             |
| completed                | Timestamp   |             |
| transaction_request_type | string      |             |
| charge_policy            | string      |             |

##### Objects {#objects-HistoricTransaction}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses {#responses-method-post-savehistorictransaction}

#### Response body {#response-body-method-post-savehistorictransaction}

| Name                     | Type        | Description |
|--------------------------|-------------|-------------|
| from                     | BankAccount |             |
| to                       | BankAccount |             |
| amount                   | Amount      |             |
| description              | string      |             |
| posted                   | Timestamp   |             |
| completed                | Timestamp   |             |
| transaction_request_type | string      |             |
| charge_policy            | string      |             |

##### Objects {#objects-HistoricTransaction}

###### BankAccount

| Name       | Type   | Description |
|------------|--------|-------------|
| bank_id    | string |             |
| account_id | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

Example:

```json
{
  "from": {
    "bank_id": "string",
    "account_id": "string"
  },
  "to": {
    "bank_id": "string",
    "account_id": "string"
  },
  "amount": {
    "cur": "string",
    "num": "string"
  },
  "description": "string",
  "posted": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "completed": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "transaction_request_type": "string",
  "charge_policy": "string"
}
```

#### Response codes {#response-codes-method-post-savehistorictransaction}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | A Historic Transaction has been imported successfully.                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
