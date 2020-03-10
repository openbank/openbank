Statements API v1.0.0
=====================

Provides CRUD operations on the statement resource.

* Host `https://`

* Base Path ``

Retrieve a statement {#method-get-getstatement}
-----------------------------------------------

GetStatement retrieves a specific statement, specified by its ID.

```sh
curl -X GET \
	https:///v1/statement/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/statement/{ID}`

### Query Parameters

| Name | Type   | Description                                  |
|------|--------|----------------------------------------------|
| id   | string | ID is the unique identifier of the statement |

### Responses

#### Response body

| Name        | Type   | Description                                                           |
|-------------|--------|-----------------------------------------------------------------------|
| id          | string | ID is the unique identifier of a statement.                           |
| status      | Status | Status is the status of the statement.                                |
| date        | string | Date is the date of the statement.                                    |
| description | string | Description is the description of the statement.                      |
| amount      | Amount | Amount is the amount if the transcation that writes on the statement. |
| balance     | Amount | Balance is remaining balance after related transaction.               |

##### Objects

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

Example:

```json
{
  "id": "string",
  "status": "Status",
  "date": "string",
  "description": "string",
  "amount": {
    "cur": "string",
    "num": "string"
  },
  "balance": {
    "cur": "string",
    "num": "string"
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

Annex
-----

#### Status

Status define the status of a statment.

| Value         | Description                                           |
|---------------|-------------------------------------------------------|
| UnknownStatus |                                                       |
| Received      | Status_Received is the value for a received statment. |
| Pending       | Status_Pending is the value for a pending statement.  |
