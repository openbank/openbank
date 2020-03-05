Product Collection API v1.0.0
=============================

Provides CRUD operations on the Product Collection resource.

* Host `https://`

* Base Path ``

Create a product collection {#method-post-createproductcollection}
------------------------------------------------------------------

Creates a new product collection and returns the object.

```sh
curl -X POST \
	https:///v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"collection_code": "string",
		"product_collection": {
			"collection_code": "string",
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
	}'
```

### HTTP Request

`POST https:///v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| collection_code | string |             |

### Body Parameters

| Name               | Type              | Description |
|--------------------|-------------------|-------------|
| bank_id            | string            |             |
| collection_code    | string            |             |
| product_collection | ProductCollection |             |

##### Objects

###### ProductCollection

| Name            | Type       | Description |
|-----------------|------------|-------------|
| collection_code | string     |             |
| products        | \[]Product |             |

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

| Name            | Type       | Description |
|-----------------|------------|-------------|
| collection_code | string     |             |
| products        | \[]Product |             |

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
  "collection_code": "string",
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
| 201    | ProductCollection created successfully.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Delete a product collection {#method-delete-deleteproductcollection}
--------------------------------------------------------------------

Permanently delete a product collection.

```sh
curl -X DELETE \
	https:///v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`DELETE https:///v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| collection_code | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductCollection successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve product collection information {#method-get-getproductcollection}
--------------------------------------------------------------------------

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	https:///v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request

`GET https:///v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| collection_code | string |             |

### Responses

#### Response body

| Name            | Type       | Description |
|-----------------|------------|-------------|
| collection_code | string     |             |
| products        | \[]Product |             |

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
  "collection_code": "string",
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

Update a product collection {#method-put-updateproductcollection}
-----------------------------------------------------------------

Updates a product collection's information

```sh
curl -X PUT \
	https:///v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"collection_code": "string",
		"product_collection": {
			"collection_code": "string",
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
	}'
```

### HTTP Request

`PUT https:///v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name            | Type   | Description |
|-----------------|--------|-------------|
| bank_id         | string |             |
| collection_code | string |             |

### Body Parameters

| Name               | Type              | Description |
|--------------------|-------------------|-------------|
| bank_id            | string            |             |
| collection_code    | string            |             |
| product_collection | ProductCollection |             |

##### Objects

###### ProductCollection

| Name            | Type       | Description |
|-----------------|------------|-------------|
| collection_code | string     |             |
| products        | \[]Product |             |

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

| Name            | Type       | Description |
|-----------------|------------|-------------|
| collection_code | string     |             |
| products        | \[]Product |             |

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
  "collection_code": "string",
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
| 204    | ProductCollection successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
