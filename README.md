# [hash-challenge](https://github.com/hashlab/hiring/blob/master/challenges/en-us/backend-challenge.md)

**By**: Ednailson Junior

**Recruter**: July Demenjon

### Requirements

* [Docker-compose](https://docs.docker.com/compose/install/)

### Running

    bash init.sh

It will be running the product-list service on the **port** `3333`

### Application

As required the application has 2 microservices. The first one is a REST API that communicates with another microservice
via gRPC asking for the product discount. This last service calculates the discount according to the date (if it is the
user birthday or if it is black friday).

##### Dump

The services do not have a way to create users or products. So I have created a dump for the database.

You can check out the [products](dump_db/products.json) and the [users](dump_db/users.json) created [here](dump_db).

## Product List Service

#### Routes

##### List products

This route lists all the products with discounts. It can receive an `user_id` parameter, and if it is the user's birthday
it will calculate the equivalente discount for all products.

Method `GET`

    /product

Parameters

* **user_id**: string

**Example**

Request

    curl --location --request GET 'http://localhost:3333/product?user_id=201'

Response

```json
[{
	"description": "A great notebook",
	"title": "Notebook Gamer",
	"price_in_cents": 2000,
	"id": "203",
	"discount": {
		"percentage": 5,
		"value_in_cents": 100
	}
}]
```

**Example 2**

Request

    curl --location --request GET 'http://localhost:3333/product'

Response

```json
[{
	"description": "A great notebook",
	"title": "Notebook Gamer",
	"price_in_cents": 2000,
	"id": "203",
	"discount": {
		"percentage": 0,
		"value_in_cents": 0
	}
}]
```

## Discount Calculator Service

This service uses gRPC for communication.

[Here it is the proto](./discount-calculator/server/discount.proto)

#### Arguments


##### Receive

* **product_id**: The product's ID
* **user_id**: The user's ID

##### Return

* **percentage**: The percentage of the discount
* **value_in_cents**: The value of the discount in cents

## Development

### Technologies

The database chose was the [ArangoDB](https://www.arangodb.com/) as a multi-model database it brings a few more
features than the MongoDB.

The discount-calculator service was built in GoLang, and the product-list was built in NodeJS.

It was chose the docker and docker-compose service to join all the microservices and the database. 

# Contact

[LinkedIn](https://www.linkedin.com/in/ednailsonvb/) | [WebSite](ednailson.github.io) | ednailsoncunha@gmail.com