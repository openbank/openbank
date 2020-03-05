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

### HTTP Request

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type    | Description                                            |
|----------------|---------|--------------------------------------------------------|
| comment        | Comment | Comment is the comment metadata made on a transaction. |
| bank_id        | string  | BankID is the bank identifier.                         |
| account_id     | string  | AccountID is the account identifier.                   |
| transaction_id | string  | TransactionID is the transaction identifier.           |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

#### Response codes

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

### HTTP Request

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type   | Description                                        |
|----------------|--------|----------------------------------------------------|
| image          | Image  | Image is the image metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                     |
| account_id     | string | AccountID is the account identifier.               |
| transaction_id | string | TransactionID is the transaction identifier.       |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects

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

#### Response codes

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

### HTTP Request

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type      | Description                                                |
|----------------|-----------|------------------------------------------------------------|
| narrative      | Narrative | Narrative is the narrative metadata made on a transaction. |
| bank_id        | string    | BankID is the bank identifier.                             |
| account_id     | string    | AccountID is the account identifier.                       |
| transaction_id | string    | TransactionID is the transaction identifier.               |

##### Objects

###### Narrative

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

### Responses

#### Response body

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
}
```

#### Response codes

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

### HTTP Request

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type   | Description                                    |
|----------------|--------|------------------------------------------------|
| tag            | Tag    | Tag is the tag metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                 |
| account_id     | string | AccountID is the account identifier.           |
| transaction_id | string | TransactionID is the transaction identifier.   |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

#### Response codes

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

### HTTP Request

`POST https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type     | Description                                               |
|----------------|----------|-----------------------------------------------------------|
| where_tag      | WhereTag | WhereTag is the where tag metadata made on a transaction. |
| bank_id        | string   | BankID is the bank identifier.                            |
| account_id     | string   | AccountID is the account identifier.                      |
| transaction_id | string   | TransactionID is the transaction identifier.              |

##### Objects

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

### Responses

#### Response body

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects

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

#### Response codes

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

### HTTP Request

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| comment        | string | ID is the comment id that will be deleted.   |

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

### HTTP Request

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| image          | string | ID is the image id that will be deleted.     |

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

### HTTP Request

`DELETE https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| narrative      | string | ID is the narrative id that will be deleted. |

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

### HTTP Request

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| tag            | string | ID is the tag id that will be deleted.       |

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

### HTTP Request

`DELETE https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| where_tag      | string | ID is the where tag id that will be deleted. |

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

### HTTP Request

`GET https:///v1/comments/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comment/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| comment        | string | ID is the comment unique identifier.         |

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

Retrieve all available comments {#method-get-getcomments}
---------------------------------------------------------

Retrieve information regarding all available comments.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses

#### Response body

| Name           | Type       | Description                                  |
|----------------|------------|----------------------------------------------|
| comments       | \[]Comment | Comments is the list of comments.            |
| bank_id        | string     | BankID is the bank identifier.               |
| account_id     | string     | AccountID is the account identifier.         |
| transaction_id | string     | TransactionID is the transaction identifier. |

##### Objects

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

Retrieve image information {#method-get-getimage}
-------------------------------------------------

Retrieve information about the image specified by the ID

```sh
curl -X GET \
	https:///v1/images/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/images/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/image/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| image          | string | ID is the image unique identifier.           |

### Responses

#### Response body

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects

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

Retrieve all available images {#method-get-getimages}
-----------------------------------------------------

Retrieve information regarding all available images.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses

#### Response body

| Name           | Type     | Description                                  |
|----------------|----------|----------------------------------------------|
| images         | \[]Image | Images is the list of images.                |
| bank_id        | string   | BankID is the bank identifier.               |
| account_id     | string   | AccountID is the account identifier.         |
| transaction_id | string   | TransactionID is the transaction identifier. |

##### Objects

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

Retrieve narrative information {#method-get-getnarrative}
---------------------------------------------------------

Retrieve information about the narrative specified by the ID

```sh
curl -X GET \
	https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/narratives/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narrative/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| narrative      | string | ID is the narrative unique identifier.       |

### Responses

#### Response body

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
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

Retrieve all available narratives {#method-get-getnarratives}
-------------------------------------------------------------

Retrieve information regarding all available narratives.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses

#### Response body

| Name           | Type         | Description                                  |
|----------------|--------------|----------------------------------------------|
| narratives     | \[]Narrative | Narratives is the list of narratives.        |
| bank_id        | string       | BankID is the bank identifier.               |
| account_id     | string       | AccountID is the account identifier.         |
| transaction_id | string       | TransactionID is the transaction identifier. |

##### Objects

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

Retrieve tag information {#method-get-gettag}
---------------------------------------------

Retrieve information about the tag specified by the ID

```sh
curl -X GET \
	https:///v1/tags/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/tags/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tag/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| tag            | string | ID is the tag unique identifier.             |

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

Retrieve all available tags {#method-get-gettags}
-------------------------------------------------

Retrieve information regarding all available tags.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses

#### Response body

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| tags           | \[]Tag | Tags is the list of tags.                    |
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

##### Objects

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

Retrieve where tag information {#method-get-getwheretag}
--------------------------------------------------------

Retrieve information about the where tag specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tag/{ID}`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |
| where_tag      | string | ID is the where tag unique identifier.       |

### Responses

#### Response body

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects

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

Retrieve all available where tags {#method-get-getwheretags}
------------------------------------------------------------

Retrieve information regarding all available where tags.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Responses

#### Response body

| Name           | Type        | Description                                  |
|----------------|-------------|----------------------------------------------|
| where_tags     | \[]WhereTag | WhereTags is the list of where tags.         |
| bank_id        | string      | BankID is the bank identifier.               |
| account_id     | string      | AccountID is the account identifier.         |
| transaction_id | string      | TransactionID is the transaction identifier. |

##### Objects

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

### HTTP Request

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/comments`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type    | Description                                            |
|----------------|---------|--------------------------------------------------------|
| comment        | Comment | Comment is the comment metadata made on a transaction. |
| bank_id        | string  | BankID is the bank identifier.                         |
| account_id     | string  | AccountID is the account identifier.                   |
| transaction_id | string  | TransactionID is the transaction identifier.           |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the image identifier.                  |
| value | string    | Value is the comment content/value.          |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

#### Response codes

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

### HTTP Request

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/images`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type   | Description                                        |
|----------------|--------|----------------------------------------------------|
| image          | Image  | Image is the image metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                     |
| account_id     | string | AccountID is the account identifier.               |
| transaction_id | string | TransactionID is the transaction identifier.       |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                    |
|-------|-----------|------------------------------------------------|
| id    | string    | ID is the image identifier.                    |
| label | string    | Label is the label of the image.               |
| url   | string    | URL is the image URL.                          |
| date  | Timestamp | Date is the date the image is created/updated. |
| user  | User      | User is the user information.                  |

##### Objects

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

#### Response codes

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

### HTTP Request

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/narratives`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type      | Description                                                |
|----------------|-----------|------------------------------------------------------------|
| narrative      | Narrative | Narrative is the narrative metadata made on a transaction. |
| bank_id        | string    | BankID is the bank identifier.                             |
| account_id     | string    | AccountID is the account identifier.                       |
| transaction_id | string    | TransactionID is the transaction identifier.               |

##### Objects

###### Narrative

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

### Responses

#### Response body

| Name      | Type   | Description                                |
|-----------|--------|--------------------------------------------|
| narrative | string | Narrative is the content of the narrative. |

Example:

```json
{
  "narrative": "string"
}
```

#### Response codes

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

### HTTP Request

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type   | Description                                    |
|----------------|--------|------------------------------------------------|
| tag            | Tag    | Tag is the tag metadata made on a transaction. |
| bank_id        | string | BankID is the bank identifier.                 |
| account_id     | string | AccountID is the account identifier.           |
| transaction_id | string | TransactionID is the transaction identifier.   |

##### Objects

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

### Responses

#### Response body

| Name  | Type      | Description                                  |
|-------|-----------|----------------------------------------------|
| id    | string    | ID is the tag identifier.                    |
| value | string    | Value is the tag content/value.              |
| date  | Timestamp | Date is the date the tag is created/updated. |
| user  | User      | User is the user information.                |

##### Objects

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

#### Response codes

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

### HTTP Request

`PUT https:///v1/banks/{BankID}/accounts/{AccountID}/transactions/{TransactionID}/metadata/where_tags`

### Query Parameters

| Name           | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| bank_id        | string | BankID is the bank identifier.               |
| account_id     | string | AccountID is the account identifier.         |
| transaction_id | string | TransactionID is the transaction identifier. |

### Body Parameters

| Name           | Type     | Description                                               |
|----------------|----------|-----------------------------------------------------------|
| where_tag      | WhereTag | WhereTag is the where tag metadata made on a transaction. |
| bank_id        | string   | BankID is the bank identifier.                            |
| account_id     | string   | AccountID is the account identifier.                      |
| transaction_id | string   | TransactionID is the transaction identifier.              |

##### Objects

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

### Responses

#### Response body

| Name     | Type      | Description                                                    |
|----------|-----------|----------------------------------------------------------------|
| location | Location  | Location is the latitude and longitude information of the tag. |
| date     | Timestamp | Date is the date the geo information is created/updated.       |
| user     | User      | User is the user information.                                  |

##### Objects

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

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | WhereTag successfully updated.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
