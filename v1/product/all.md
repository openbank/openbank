Product API v1.0.0
==================

Provides CRUD operations on the Product resource.

* Host `https://`

* Base Path ``

Create a product {#method-post-createproduct}
---------------------------------------------

Creates a new product and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product": {
			"bank_id": "string",
			"code": "string",
			"parent_product_code": "string",
			"name": "string",
			"category": "string",
			"family": "string",
			"super_family": "string",
			"more_info_url": "string",
			"details": "string",
			"description": "string",
			"meta": {
				"license": "string"
			},
			"product_attributes": [
				{
					"product_code": "string",
					"product_attribute_id": "string",
					"name": "string",
					"type": "AttributeType",
					"value": "string"
				}
			]
		}
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| product_code | string |             |

### Body Parameters

| Name         | Type    | Description |
|--------------|---------|-------------|
| bank_id      | string  |             |
| product_code | string  |             |
| product      | Product |             |

##### Objects

###### Product

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

### Responses

#### Response body

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
    }
  ]
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Product created successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Create a product attribute {#method-post-createproductattribute}
----------------------------------------------------------------

Creates a new product attribute and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product_attribute_id": "string",
		"product_attribute": {
			"product_code": "string",
			"product_attribute_id": "string",
			"name": "string",
			"type": "AttributeType",
			"value": "string"
		}
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name                 | Type   | Description |
|----------------------|--------|-------------|
| bank_id              | string |             |
| product_code         | string |             |
| product_attribute_id | string |             |

### Body Parameters

| Name                 | Type             | Description |
|----------------------|------------------|-------------|
| bank_id              | string           |             |
| product_code         | string           |             |
| product_attribute_id | string           |             |
| product_attribute    | ProductAttribute |             |

##### Objects

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

### Responses

#### Response body

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | ProductAttribute created successfully.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a product {#method-delete-deleteproduct}
-----------------------------------------------

Permanently delete a product.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| product_code | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Product successfully deleted.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a product attribute {#method-delete-deleteproductattribute}
------------------------------------------------------------------

Permanently delete a product attribute.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name                 | Type   | Description |
|----------------------|--------|-------------|
| bank_id              | string |             |
| product_code         | string |             |
| product_attribute_id | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductAttribute successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve product information {#method-get-getproduct}
-----------------------------------------------------

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| product_code | string |             |

### Responses

#### Response body

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
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

Retrieve product attribute information {#method-get-getproductattribute}
------------------------------------------------------------------------

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name                 | Type   | Description |
|----------------------|--------|-------------|
| bank_id              | string |             |
| product_code         | string |             |
| product_attribute_id | string |             |

### Responses

#### Response body

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
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

Retrieve all available products {#method-get-getproducts}
---------------------------------------------------------

Retrieve information regarding all available products.

```sh
curl -X GET \
	https:///v1/banks/{BankID}/products \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/products`

### Query Parameters

| Name    | Type   | Description |
|---------|--------|-------------|
| bank_id | string |             |

### Responses

#### Response body

| Name     | Type       | Description |
|----------|------------|-------------|
| products | \[]Product |             |

##### Objects

###### Product

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "products": [
    {
      "bank_id": "string",
      "code": "string",
      "parent_product_code": "string",
      "name": "string",
      "category": "string",
      "family": "string",
      "super_family": "string",
      "more_info_url": "string",
      "details": "string",
      "description": "string",
      "meta": {
        "license": "string"
      },
      "product_attributes": [
        {
          "product_code": "string",
          "product_attribute_id": "string",
          "name": "string",
          "type": "AttributeType",
          "value": "string"
        }
      ]
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

Update a product {#method-put-updateproduct}
--------------------------------------------

Updates a product's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product": {
			"bank_id": "string",
			"code": "string",
			"parent_product_code": "string",
			"name": "string",
			"category": "string",
			"family": "string",
			"super_family": "string",
			"more_info_url": "string",
			"details": "string",
			"description": "string",
			"meta": {
				"license": "string"
			},
			"product_attributes": [
				{
					"product_code": "string",
					"product_attribute_id": "string",
					"name": "string",
					"type": "AttributeType",
					"value": "string"
				}
			]
		}
	}'
```

### HTTP Request

`PUT https:///v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name         | Type   | Description |
|--------------|--------|-------------|
| bank_id      | string |             |
| product_code | string |             |

### Body Parameters

| Name         | Type    | Description |
|--------------|---------|-------------|
| bank_id      | string  |             |
| product_code | string  |             |
| product      | Product |             |

##### Objects

###### Product

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

### Responses

#### Response body

| Name                | Type                | Description |
|---------------------|---------------------|-------------|
| bank_id             | string              |             |
| code                | string              |             |
| parent_product_code | string              |             |
| name                | string              |             |
| category            | string              |             |
| family              | string              |             |
| super_family        | string              |             |
| more_info_url       | string              |             |
| details             | string              |             |
| description         | string              |             |
| meta                | Metadata            |             |
| product_attributes  | \[]ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| license | string |             |

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
    }
  ]
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Product successfully updated.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Update a product attribute {#method-put-updateproductattribute}
---------------------------------------------------------------

Updates a product attribute's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product_attribute_id": "string",
		"product_attribute": {
			"product_code": "string",
			"product_attribute_id": "string",
			"name": "string",
			"type": "AttributeType",
			"value": "string"
		}
	}'
```

### HTTP Request

`PUT https:///v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name                 | Type   | Description |
|----------------------|--------|-------------|
| bank_id              | string |             |
| product_code         | string |             |
| product_attribute_id | string |             |

### Body Parameters

| Name                 | Type             | Description |
|----------------------|------------------|-------------|
| bank_id              | string           |             |
| product_code         | string           |             |
| product_attribute_id | string           |             |
| product_attribute    | ProductAttribute |             |

##### Objects

###### ProductAttribute

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

### Responses

#### Response body

| Name                 | Type          | Description |
|----------------------|---------------|-------------|
| product_code         | string        |             |
| product_attribute_id | string        |             |
| name                 | string        |             |
| type                 | AttributeType |             |
| value                | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
}
```

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductAttribute successfully updated.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Annex
-----

#### AttributeType

| Value         | Description |
|---------------|-------------|
| UnknownType   |             |
| STRING        |             |
| INTEGER       |             |
| DOUBLE        |             |
| DATE_WITH_DAY |             |
