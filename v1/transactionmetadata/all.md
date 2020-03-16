Transaction Metadata API v1.0.0
===============================

Provides the access and availability to the transaction metadata API.

* Host `https://`

* Base Path ``

Create a comment {#method-post-createcomment}
---------------------------------------------

Creates a new comment and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"comment": {
			"id": "string",
			"value": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createcomment}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters {#query-parameters-method-post-createcomment}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-post-createcomment}

| Name           | Type    | Description                                            |
|----------------|---------|--------------------------------------------------------|
| comment        | Comment | Comment is the comment metadata made on a transaction. |
| bank_id        | string  | BankID is the bank identifier.                         |
| account_id     | string  | AccountID is the account identifier.                   |
| transaction_id | string  | TransactionID is the transaction identifier.           |

##### Objects {#objects-CreateCommentRequest}

###### Comment

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-post-createcomment}

#### Response body {#response-body-method-post-createcomment}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Comment}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-post-createcomment}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Comment created successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a image {#method-post-createimage}
-----------------------------------------

Creates a new image and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"image": {
			"id": "string",
			"label": "string",
			"url": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createimage}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters {#query-parameters-method-post-createimage}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-post-createimage}

| Name           | Type   | Description                                        |
|----------------|--------|----------------------------------------------------|
| image          | Image  | Image is the image metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                     |
| account_id     | string | AccountID is the account identifier.               |
| transaction_id | string | TransactionID is the transaction identifier.       |

##### Objects {#objects-CreateImageRequest}

###### Image

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-post-createimage}

#### Response body {#response-body-method-post-createimage}

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects {#objects-Image}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "label": "string",
  "url": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-post-createimage}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Image created successfully.                                                            |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a narrative {#method-post-createnarrative}
-------------------------------------------------

Creates a new narrative and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"narrative": {
			"narrative": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createnarrative}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters {#query-parameters-method-post-createnarrative}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-post-createnarrative}

| Name           | Type      | Description                                                |
|----------------|-----------|------------------------------------------------------------|
| narrative      | Narrative | Narrative is the narrative metadata made on a transaction. |
| bank_id        | string    | BankID is the bank identifier.                             |
| account_id     | string    | AccountID is the account identifier.                       |
| transaction_id | string    | TransactionID is the transaction identifier.               |

##### Objects {#objects-CreateNarrativeRequest}

###### Narrative

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

### Responses {#responses-method-post-createnarrative}

#### Response body {#response-body-method-post-createnarrative}

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
}
```

#### Response codes {#response-codes-method-post-createnarrative}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Narrative created successfully.                                                        |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a tag {#method-post-createtag}
-------------------------------------

Creates a new tag and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"tag": {
			"id": "string",
			"value": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createtag}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters {#query-parameters-method-post-createtag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-post-createtag}

| Name           | Type   | Description                                    |
|----------------|--------|------------------------------------------------|
| tag            | Tag    | Tag is the tag metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                 |
| account_id     | string | AccountID is the account identifier.           |
| transaction_id | string | TransactionID is the transaction identifier.   |

##### Objects {#objects-CreateTagRequest}

###### Tag

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-post-createtag}

#### Response body {#response-body-method-post-createtag}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Tag}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-post-createtag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Tag created successfully.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a where tag {#method-post-createwheretag}
------------------------------------------------

Creates a new where tag and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"where_tag": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			},
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-post-createwheretag}

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters {#query-parameters-method-post-createwheretag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-post-createwheretag}

| Name           | Type     | Description                                               |
|----------------|----------|-----------------------------------------------------------|
| where_tag      | WhereTag | WhereTag is the where tag metadata made on a transaction. |
| bank_id        | string   | BankID is the bank identifier.                            |
| account_id     | string   | AccountID is the account identifier.                      |
| transaction_id | string   | TransactionID is the transaction identifier.              |

##### Objects {#objects-CreateWhereTagRequest}

###### WhereTag

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-post-createwheretag}

#### Response body {#response-body-method-post-createwheretag}

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects {#objects-WhereTag}

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  },
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-post-createwheretag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | WhereTag created successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a comment {#method-delete-deletecomment}
-----------------------------------------------

Permanently delete a comment.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletecomment}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID}`

### Query Parameters {#query-parameters-method-delete-deletecomment}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| comment        | string | ID is the comment id that will be deleted.   |

### Responses {#responses-method-delete-deletecomment}

#### Response body {#response-body-method-delete-deletecomment}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletecomment}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Comment successfully deleted.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a image {#method-delete-deleteimage}
-------------------------------------------

Permanently delete a image.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deleteimage}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID}`

### Query Parameters {#query-parameters-method-delete-deleteimage}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| image          | string | ID is the image id that will be deleted.     |

### Responses {#responses-method-delete-deleteimage}

#### Response body {#response-body-method-delete-deleteimage}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deleteimage}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Image successfully deleted.                                                            |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a narrative {#method-delete-deletenarrative}
---------------------------------------------------

Permanently delete a narrative.

```sh
curl -X DELETE \
	https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletenarrative}

`DELETE https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID}`

### Query Parameters {#query-parameters-method-delete-deletenarrative}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| narrative      | string | ID is the narrative id that will be deleted. |

### Responses {#responses-method-delete-deletenarrative}

#### Response body {#response-body-method-delete-deletenarrative}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletenarrative}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Narrative successfully deleted.                                                        |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a tag {#method-delete-deletetag}
---------------------------------------

Permanently delete a tag.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletetag}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID}`

### Query Parameters {#query-parameters-method-delete-deletetag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| tag            | string | ID is the tag id that will be deleted.       |

### Responses {#responses-method-delete-deletetag}

#### Response body {#response-body-method-delete-deletetag}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletetag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Tag successfully deleted.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a where tag {#method-delete-deletewheretag}
--------------------------------------------------

Permanently delete a where tag.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-delete-deletewheretag}

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID}`

### Query Parameters {#query-parameters-method-delete-deletewheretag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| where_tag      | string | ID is the where tag id that will be deleted. |

### Responses {#responses-method-delete-deletewheretag}

#### Response body {#response-body-method-delete-deletewheretag}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-delete-deletewheretag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | WhereTag successfully deleted.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve comment information {#method-get-getcomment}
-----------------------------------------------------

Retrieve information about the comment specified by the ID

```sh
curl -X GET \
	https:///v1/comments/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcomment}

`GET https:///v1/comments/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID}`

### Query Parameters {#query-parameters-method-get-getcomment}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| comment        | string | ID is the comment unique identifier.         |

### Responses {#responses-method-get-getcomment}

#### Response body {#response-body-method-get-getcomment}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Comment}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-get-getcomment}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available comments {#method-get-getcomments}
---------------------------------------------------------

Retrieve information regarding all available comments.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getcomments}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters {#query-parameters-method-get-getcomments}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses {#responses-method-get-getcomments}

#### Response body {#response-body-method-get-getcomments}

| Name           | Type       | Description                                  |
|----------------|------------|----------------------------------------------|
| comments       | \[]Comment | Comments is the list of comments.            |
| bank_id        | string     | BankID is the bank identifier.               |
| account_id     | string     | AccountID is the account identifier.         |
| transaction_id | string     | TransactionID is the transaction identifier. |

##### Objects {#objects-GetCommentsResponse}

###### Comment

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "comments": [
    {
      "id": "string",
      "value": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "user": {
        "id": "string",
        "provider": "string",
        "display_name": "string"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "transaction_id": "string"
}
```

#### Response codes {#response-codes-method-get-getcomments}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve image information {#method-get-getimage}
-------------------------------------------------

Retrieve information about the image specified by the ID

```sh
curl -X GET \
	https:///v1/images/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getimage}

`GET https:///v1/images/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID}`

### Query Parameters {#query-parameters-method-get-getimage}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| image          | string | ID is the image unique identifier.           |

### Responses {#responses-method-get-getimage}

#### Response body {#response-body-method-get-getimage}

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects {#objects-Image}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "label": "string",
  "url": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-get-getimage}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available images {#method-get-getimages}
-----------------------------------------------------

Retrieve information regarding all available images.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getimages}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters {#query-parameters-method-get-getimages}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses {#responses-method-get-getimages}

#### Response body {#response-body-method-get-getimages}

| Name           | Type     | Description                                  |
|----------------|----------|----------------------------------------------|
| images         | \[]Image | Images is the list of images.                |
| bank_id        | string   | BankID is the bank identifier.               |
| account_id     | string   | AccountID is the account identifier.         |
| transaction_id | string   | TransactionID is the transaction identifier. |

##### Objects {#objects-GetImagesResponse}

###### Image

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "images": [
    {
      "id": "string",
      "label": "string",
      "url": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "user": {
        "id": "string",
        "provider": "string",
        "display_name": "string"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "transaction_id": "string"
}
```

#### Response codes {#response-codes-method-get-getimages}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve narrative information {#method-get-getnarrative}
---------------------------------------------------------

Retrieve information about the narrative specified by the ID

```sh
curl -X GET \
	https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getnarrative}

`GET https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID}`

### Query Parameters {#query-parameters-method-get-getnarrative}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| narrative      | string | ID is the narrative unique identifier.       |

### Responses {#responses-method-get-getnarrative}

#### Response body {#response-body-method-get-getnarrative}

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
}
```

#### Response codes {#response-codes-method-get-getnarrative}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available narratives {#method-get-getnarratives}
-------------------------------------------------------------

Retrieve information regarding all available narratives.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getnarratives}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters {#query-parameters-method-get-getnarratives}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses {#responses-method-get-getnarratives}

#### Response body {#response-body-method-get-getnarratives}

| Name           | Type         | Description                                  |
|----------------|--------------|----------------------------------------------|
| narratives     | \[]Narrative | Narratives is the list of narratives.        |
| bank_id        | string       | BankID is the bank identifier.               |
| account_id     | string       | AccountID is the account identifier.         |
| transaction_id | string       | TransactionID is the transaction identifier. |

##### Objects {#objects-GetNarrativesResponse}

###### Narrative

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narratives": [
    {
      "narrative": "string"
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "transaction_id": "string"
}
```

#### Response codes {#response-codes-method-get-getnarratives}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve tag information {#method-get-gettag}
---------------------------------------------

Retrieve information about the tag specified by the ID

```sh
curl -X GET \
	https:///v1/tags/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-gettag}

`GET https:///v1/tags/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID}`

### Query Parameters {#query-parameters-method-get-gettag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| tag            | string | ID is the tag unique identifier.             |

### Responses {#responses-method-get-gettag}

#### Response body {#response-body-method-get-gettag}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Tag}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-get-gettag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available tags {#method-get-gettags}
-------------------------------------------------

Retrieve information regarding all available tags.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-gettags}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters {#query-parameters-method-get-gettags}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses {#responses-method-get-gettags}

#### Response body {#response-body-method-get-gettags}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| tags           | \[]Tag | Tags is the list of tags.                    |
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

##### Objects {#objects-GetTagsResponse}

###### Tag

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "tags": [
    {
      "id": "string",
      "value": "string",
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "user": {
        "id": "string",
        "provider": "string",
        "display_name": "string"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "transaction_id": "string"
}
```

#### Response codes {#response-codes-method-get-gettags}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve where tag information {#method-get-getwheretag}
--------------------------------------------------------

Retrieve information about the where tag specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getwheretag}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID}`

### Query Parameters {#query-parameters-method-get-getwheretag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| where_tag      | string | ID is the where tag unique identifier.       |

### Responses {#responses-method-get-getwheretag}

#### Response body {#response-body-method-get-getwheretag}

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects {#objects-WhereTag}

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  },
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-get-getwheretag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve all available where tags {#method-get-getwheretags}
------------------------------------------------------------

Retrieve information regarding all available where tags.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getwheretags}

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters {#query-parameters-method-get-getwheretags}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses {#responses-method-get-getwheretags}

#### Response body {#response-body-method-get-getwheretags}

| Name           | Type        | Description                                  |
|----------------|-------------|----------------------------------------------|
| where_tags     | \[]WhereTag | WhereTags is the list of where tags.         |
| bank_id        | string      | BankID is the bank identifier.               |
| account_id     | string      | AccountID is the account identifier.         |
| transaction_id | string      | TransactionID is the transaction identifier. |

##### Objects {#objects-GetWhereTagsResponse}

###### WhereTag

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "where_tags": [
    {
      "location": {
        "latitude": "double",
        "longitude": "double"
      },
      "date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "user": {
        "id": "string",
        "provider": "string",
        "display_name": "string"
      }
    }
  ],
  "bank_id": "string",
  "account_id": "string",
  "transaction_id": "string"
}
```

#### Response codes {#response-codes-method-get-getwheretags}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a comment {#method-put-updatecomment}
--------------------------------------------

Updates a comment's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"comment": {
			"id": "string",
			"value": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatecomment}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters {#query-parameters-method-put-updatecomment}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-put-updatecomment}

| Name           | Type    | Description                                            |
|----------------|---------|--------------------------------------------------------|
| comment        | Comment | Comment is the comment metadata made on a transaction. |
| bank_id        | string  | BankID is the bank identifier.                         |
| account_id     | string  | AccountID is the account identifier.                   |
| transaction_id | string  | TransactionID is the transaction identifier.           |

##### Objects {#objects-UpdateCommentRequest}

###### Comment

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-put-updatecomment}

#### Response body {#response-body-method-put-updatecomment}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Comment}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatecomment}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Comment successfully updated.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a image {#method-put-updateimage}
----------------------------------------

Updates a image's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"image": {
			"id": "string",
			"label": "string",
			"url": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updateimage}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters {#query-parameters-method-put-updateimage}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-put-updateimage}

| Name           | Type   | Description                                        |
|----------------|--------|----------------------------------------------------|
| image          | Image  | Image is the image metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                     |
| account_id     | string | AccountID is the account identifier.               |
| transaction_id | string | TransactionID is the transaction identifier.       |

##### Objects {#objects-UpdateImageRequest}

###### Image

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-put-updateimage}

#### Response body {#response-body-method-put-updateimage}

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects {#objects-Image}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "label": "string",
  "url": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updateimage}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Image successfully updated.                                                            |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a narrative {#method-put-updatenarrative}
------------------------------------------------

Updates a narrative's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"narrative": {
			"narrative": "string"
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatenarrative}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters {#query-parameters-method-put-updatenarrative}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-put-updatenarrative}

| Name           | Type      | Description                                                |
|----------------|-----------|------------------------------------------------------------|
| narrative      | Narrative | Narrative is the narrative metadata made on a transaction. |
| bank_id        | string    | BankID is the bank identifier.                             |
| account_id     | string    | AccountID is the account identifier.                       |
| transaction_id | string    | TransactionID is the transaction identifier.               |

##### Objects {#objects-UpdateNarrativeRequest}

###### Narrative

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

### Responses {#responses-method-put-updatenarrative}

#### Response body {#response-body-method-put-updatenarrative}

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
}
```

#### Response codes {#response-codes-method-put-updatenarrative}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Narrative successfully updated.                                                        |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a tag {#method-put-updatetag}
------------------------------------

Updates a tag's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"tag": {
			"id": "string",
			"value": "string",
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatetag}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters {#query-parameters-method-put-updatetag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-put-updatetag}

| Name           | Type   | Description                                    |
|----------------|--------|------------------------------------------------|
| tag            | Tag    | Tag is the tag metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                 |
| account_id     | string | AccountID is the account identifier.           |
| transaction_id | string | TransactionID is the transaction identifier.   |

##### Objects {#objects-UpdateTagRequest}

###### Tag

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-put-updatetag}

#### Response body {#response-body-method-put-updatetag}

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects {#objects-Tag}

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "id": "string",
  "value": "string",
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatetag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Tag successfully updated.                                                              |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a where tag {#method-put-updatewheretag}
-----------------------------------------------

Updates a where tag's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"where_tag": {
			"location": {
				"latitude": "double",
				"longitude": "double"
			},
			"date": {
				"seconds": "int64",
				"nanos": "int32"
			},
			"user": {
				"id": "string",
				"provider": "string",
				"display_name": "string"
			}
		},
		"bank_id": "string",
		"account_id": "string",
		"transaction_id": "string"
	}'
```

### HTTP Request {#http-request-method-put-updatewheretag}

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters {#query-parameters-method-put-updatewheretag}

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters {#body-parameters-method-put-updatewheretag}

| Name           | Type     | Description                                               |
|----------------|----------|-----------------------------------------------------------|
| where_tag      | WhereTag | WhereTag is the where tag metadata made on a transaction. |
| bank_id        | string   | BankID is the bank identifier.                            |
| account_id     | string   | AccountID is the account identifier.                      |
| transaction_id | string   | TransactionID is the transaction identifier.              |

##### Objects {#objects-UpdateWhereTagRequest}

###### WhereTag

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

### Responses {#responses-method-put-updatewheretag}

#### Response body {#response-body-method-put-updatewheretag}

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects {#objects-WhereTag}

###### Location

| Name      | Type   | Description                                                         |
|-----------|--------|---------------------------------------------------------------------|
| latitude  | double | The latitude in degrees. It must be in the range [-90.0, +90.0].    |
| longitude | double | The longitude in degrees. It must be in the range [-180.0, +180.0\] |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### User

| Name         | Type   | Description                             |
|--------------|--------|-----------------------------------------|
| id           | string | ID is the user identifier.              |
| provider     | string | Provider is the provider of the user.   |
| display_name | string | DisplayName is the user's display name. |

Example:

```json
{
  "location": {
    "latitude": "double",
    "longitude": "double"
  },
  "date": {
    "seconds": "int64",
    "nanos": "int32"
  },
  "user": {
    "id": "string",
    "provider": "string",
    "display_name": "string"
  }
}
```

#### Response codes {#response-codes-method-put-updatewheretag}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | WhereTag successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
