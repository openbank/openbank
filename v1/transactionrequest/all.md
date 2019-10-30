# Transacton Request API v1.0.0

Provides CRUD operations on the transaction request resource.

*
Host ``
EOL

*
Base Path ``

## Answer the transaction reqeust challenge

Answer the transaction request challenge

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/{TransactionRequestType}/transactionrequest/{TransactionRequestID}/challenge \
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
{{snippet answertransactionrequest_challenge []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/{TransactionRequestType}/transactionrequest/{TransactionRequestID}/challenge`

### Query Parameters

| Name                   | Type   | Description |
|------------------------|--------|-------------|
| BankID                 | string |             |
| AccountID              | string |             |
| TransactionRequestType | string |             |
| TransactionRequestID   | string |             |

### Body Parameters

| Name                   | Type   | Description |
|------------------------|--------|-------------|
| BankID                 | string |             |
| AccountID              | string |             |
| TransactionRequestType | string |             |
| TransactionRequestID   | string |             |
| ID                     | string |             |
| Answer                 | string |             |

### Responses

#### Response body

| Name           | Type             | Description |
|----------------|------------------|-------------|
| ID             | string           |             |
| Type           | string           |             |
| From           | BankAccount      |             |
| Details        | ChallengeDetails |             |
| TransactionIds | string           |             |
| Status         | string           |             |
| StartDate      | Timestamp        |             |
| EndDate        | Timestamp        |             |
| Challenge      | Challenge        |             |
| Charge         | Charge           |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### ChallengeDetails

| Name        | Type        | Description |
|-------------|-------------|-------------|
| To          | BankAccount |             |
| Amount      | Amount      |             |
| Description | string      |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Answered the transaction request challenge                                             |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an account otp transaction request

Creates a new account otp transaction request

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT_OTP/transactionrequest \
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
{{snippet createaccountotp_transaction []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT_OTP/transactionrequest`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Body Parameters

| Name        | Type        | Description |
|-------------|-------------|-------------|
| BankID      | string      |             |
| AccountID   | string      |             |
| To          | BankAccount |             |
| Amount      | Amount      |             |
| Description | string      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name           | Type        | Description |
|----------------|-------------|-------------|
| ID             | string      |             |
| Type           | string      |             |
| From           | BankAccount |             |
| Details        | Details     |             |
| TransactionIds | []string    |             |
| Status         | string      |             |
| StartDate      | Timestamp   |             |
| EndDate        | Timestamp   |             |
| Challenge      | Challenge   |             |
| Charge         | Charge      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account OTP Transaction request created successfully.                                  |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an account transaction request

Creates a new account transaction request

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT/transactionrequest \
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
{{snippet createaccounttransaction []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/ACCOUNT/transactionrequest`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Body Parameters

| Name        | Type        | Description |
|-------------|-------------|-------------|
| BankID      | string      |             |
| AccountID   | string      |             |
| To          | BankAccount |             |
| Amount      | Amount      |             |
| Description | string      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name           | Type        | Description |
|----------------|-------------|-------------|
| ID             | string      |             |
| Type           | string      |             |
| From           | BankAccount |             |
| Details        | Details     |             |
| TransactionIds | []string    |             |
| Status         | string      |             |
| StartDate      | Timestamp   |             |
| EndDate        | Timestamp   |             |
| Challenge      | Challenge   |             |
| Charge         | Charge      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Account Transaction request created successfully.                                      |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an counter party transaction request

Creates a new counter party transaction request

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/COUNTERPARTY/transactionrequest \
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
{{snippet createcounterparty_transaction []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/COUNTERPARTY/transactionrequest`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Body Parameters

| Name         | Type           | Description |
|--------------|----------------|-------------|
| BankID       | string         |             |
| AccountID    | string         |             |
| To           | ToCounterparty |             |
| Amount       | Amount         |             |
| Description  | string         |             |
| ChargePolicy | string         |             |
| FutureDate   | string         |             |

##### Objects

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name           | Type        | Description |
|----------------|-------------|-------------|
| ID             | string      |             |
| Type           | string      |             |
| From           | BankAccount |             |
| Details        | Details     |             |
| TransactionIds | []string    |             |
| Status         | string      |             |
| StartDate      | Timestamp   |             |
| EndDate        | Timestamp   |             |
| Challenge      | Challenge   |             |
| Charge         | Charge      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Counter Party Transaction request created successfully.                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an free form transaction request

Creates a new free form transaction request

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/FREE_FORM/transactionrequest \
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
{{snippet createfreeform_transaction []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/FREE_FORM/transactionrequest`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Body Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| AccountID   | string |             |
| Amount      | Amount |             |
| Description | string |             |

##### Objects

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

### Responses

#### Response body

| Name           | Type        | Description |
|----------------|-------------|-------------|
| ID             | string      |             |
| Type           | string      |             |
| From           | BankAccount |             |
| Details        | Details     |             |
| TransactionIds | []string    |             |
| Status         | string      |             |
| StartDate      | Timestamp   |             |
| EndDate        | Timestamp   |             |
| Challenge      | Challenge   |             |
| Charge         | Charge      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Free Form Transaction request created successfully.                                    |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create an sepa transaction request

Creates a new sepa transaction request

```sh
curl -X POST \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/SEPA/transactionrequest \
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
{{snippet createsepa_transaction []}}

### HTTP Request

`POST /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types/SEPA/transactionrequest`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Body Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| BankID       | string |             |
| AccountID    | string |             |
| Amount       | Amount |             |
| To           | ToSepa |             |
| Description  | string |             |
| ChargePolicy | string |             |
| FutureDate   | string |             |

##### Objects

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

### Responses

#### Response body

| Name           | Type        | Description |
|----------------|-------------|-------------|
| ID             | string      |             |
| Type           | string      |             |
| From           | BankAccount |             |
| Details        | Details     |             |
| TransactionIds | []string    |             |
| Status         | string      |             |
| StartDate      | Timestamp   |             |
| EndDate        | Timestamp   |             |
| Challenge      | Challenge   |             |
| Charge         | Charge      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | SEPA Transaction request created successfully.                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve transaction types

Retrieves list of transaction request types

```sh
curl -X GET \
	/v1/banks/{BankID}/transaction-request-types \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getsupportedtransactionrequesttypes []}}

### HTTP Request

`GET /v1/banks/{BankID}/transaction-request-types`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name                    | Type                     | Description |
|-------------------------|--------------------------|-------------|
| TransactionRequestTypes | []TransactionRequestType |             |

##### Objects

###### TransactionRequestType

| Name   | Type   | Description |
|--------|--------|-------------|
| Value  | string |             |
| Charge | Charge |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Retrieve transaction types

Retrieves list of transaction request types for an account_id

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet gettransactionrequest_types []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/transaction-request-types`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Responses

#### Response body

| Name                    | Type                     | Description |
|-------------------------|--------------------------|-------------|
| TransactionRequestTypes | []TransactionRequestType |             |

##### Objects

###### TransactionRequestType

| Name   | Type   | Description |
|--------|--------|-------------|
| Value  | string |             |
| Charge | Charge |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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

## Retrieve transaction requests

Retrieves list of transaction requests for an account_id

```sh
curl -X GET \
	/v1/banks/{BankID}/accounts/{AccountID}/transactionrequests \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet gettransactionrequests []}}

### HTTP Request

`GET /v1/banks/{BankID}/accounts/{AccountID}/transactionrequests`

### Query Parameters

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

### Responses

#### Response body

| Name                           | Type                             | Description |
|--------------------------------|----------------------------------|-------------|
| TransactionRequestsWithCharges | []TransactionRequestsWithCharges |             |

##### Objects

###### TransactionRequestsWithCharges

| Name           | Type      | Description |
|----------------|-----------|-------------|
| ID             | string    |             |
| Type           | string    |             |
| From           | FromPhone |             |
| Details        | Details   |             |
| TransactionIds | []string  |             |
| Status         | string    |             |
| StartDate      | Timestamp |             |
| EndDate        | Timestamp |             |
| Challenge      | Challenge |             |
| Charge         | Charge    |             |

###### FromPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |
| Nickname          | string |             |

###### Details

| Name                  | Type                  | Description |
|-----------------------|-----------------------|-------------|
| ToSandboxTan          | BankAccount           |             |
| ToSepa                | ToSepa                |             |
| ToCounterparty        | ToCounterparty        |             |
| ToTransferToPhone     | ToTransferToPhone     |             |
| ToTransferToAtm       | ToTransferToAtm       |             |
| ToTransferToAccount   | ToTransferToAccount   |             |
| ToSepaCreditTransfers | ToSepaCreditTransfers |             |
| Amount                | Amount                |             |
| Description           | string                |             |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### Challenge

| Name            | Type   | Description |
|-----------------|--------|-------------|
| ID              | string |             |
| AllowedAttempts | int32  |             |
| ChallengeType   | string |             |

###### Charge

| Name    | Type   | Description |
|---------|--------|-------------|
| Summary | string |             |
| Amount  | Amount |             |

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### ToSepa

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### ToCounterparty

| Name           | Type   | Description |
|----------------|--------|-------------|
| CounterpartyID | string |             |

###### ToTransferToPhone

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAtm

| Name        | Type      | Description |
|-------------|-----------|-------------|
| Amount      | Amount    |             |
| Description | string    |             |
| Message     | string    |             |
| From        | FromPhone |             |
| To          | ToPhone   |             |

###### ToTransferToAccount

| Name         | Type   | Description |
|--------------|--------|-------------|
| Amount       | Amount |             |
| Description  | string |             |
| TransferType | string |             |
| FutureDate   | string |             |
| To           | To     |             |

###### ToSepaCreditTransfers

| Name             | Type            | Description |
|------------------|-----------------|-------------|
| DebtorAccount    | DebtorAccount   |             |
| InstructedAmount | Amount          |             |
| CreditorAccount  | CreditorAccount |             |
| CreditorName     | string          |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### ToPhone

| Name              | Type   | Description |
|-------------------|--------|-------------|
| MobilePhoneNumber | string |             |

###### To

| Name              | Type        | Description |
|-------------------|-------------|-------------|
| LegalName         | string      |             |
| DateOfBirth       | string      |             |
| MobilePhoneNumber | string      |             |
| KycDocument       | KycDocument |             |

###### DebtorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### CreditorAccount

| Name | Type   | Description |
|------|--------|-------------|
| Iban | string |             |

###### KycDocument

| Name   | Type   | Description |
|--------|--------|-------------|
| Type   | string |             |
| Number | string |             |

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

## Import a historic transaction

Import a historic transaction

```sh
curl -X POST \
	/v1/transactionrequest/import \
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
{{snippet savehistorictransaction []}}

### HTTP Request

`POST /v1/transactionrequest/import`

### Body Parameters

| Name                   | Type        | Description |
|------------------------|-------------|-------------|
| From                   | BankAccount |             |
| To                     | BankAccount |             |
| Amount                 | Amount      |             |
| Description            | string      |             |
| Posted                 | Timestamp   |             |
| Completed              | Timestamp   |             |
| TransactionRequestType | string      |             |
| ChargePolicy           | string      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

### Responses

#### Response body

| Name                   | Type        | Description |
|------------------------|-------------|-------------|
| From                   | BankAccount |             |
| To                     | BankAccount |             |
| Amount                 | Amount      |             |
| Description            | string      |             |
| Posted                 | Timestamp   |             |
| Completed              | Timestamp   |             |
| TransactionRequestType | string      |             |
| ChargePolicy           | string      |             |

##### Objects

###### BankAccount

| Name      | Type   | Description |
|-----------|--------|-------------|
| BankID    | string |             |
| AccountID | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | A Historic Transaction has been imported successfully.                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
