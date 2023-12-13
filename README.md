## Introduction

This project aims to develop a robust and scalable API for an e-commerce website using the Go programming language, MongoDB as the database, Docker for containerization, and the Gin web framework to streamline the development process.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Signup](#signup)
  - [Login](#login)
  - [Admin Add Product](#admin-add-product)
  - [View All Products](#view-all-products)
  - [Search Products](#search-products)
  - [Search Products by Category](#search-products-by-category)
  - [Add to Cart](#add-to-cart)
  - [List Cart](#list-cart)
  - [Remove Item from Cart](#remove-item-from-cart)
  - [Add Address](#add-address)
  - [Edit Address](#edit-address)
  - [Delete Address](#delete-address)
  - [Cart Checkout](#cart-checkout)
  - [Instantly Buy Products](#instantly-buy-products)
- [Contributing](#contributing)
- [Acknowledgments](#acknowledgments)


## Key Features:

1. Go Programming Language:
   - Leveraging the power and efficiency of Go for high-performance and concurrent API development.
   - Utilizing Go's statically typed nature and extensive standard library for efficient coding.
3. MongoDB Integration:
   - Implementing MongoDB as the database to store and manage product information, user data, and other relevant e-commerce details.
   - Leveraging MongoDB's flexible schema to accommodate evolving business requirements.
5. Docker Containerization:
   - Utilizing Docker for containerization to ensure consistency in development, testing, and deployment environments.
   - Employing Docker Compose for orchestrating the MongoDB container and the Go application, simplifying the setup process.
7. Gin Web Framework:
   - Using the Gin web framework to streamline the development of the API with its fast performance and minimalistic design.
   - Taking advantage of Gin's middleware support for implementing features like authentication, logging, and error handling.
9. RESTful API Design:
    - Designing a RESTful API that adheres to best practices for e-commerce, including endpoints for product listing, user authentication, shopping cart management, and order processing.
    - Implementing CRUD operations to manage products, users, and orders efficiently.

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [MongoDB](https://www.mongodb.com/)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Infamous-Ironman/go_project   ```

2. Navigate into the project directory**

   ```bash
   cd gp_project
   ```


3. You can start the project with below commands

    ```bash
    docker-compose up -d
    go run main.go
    ```

### Usage

- #### SIGNUP

<http://localhost:8000/users/signup>

```json
{
  "first_name": "Sanchar",
  "last_name": "Limbu",
  "email": "test@gmail.com",
  "password": "12345678",
  "phone": "9841111111"
}
```

Response: "Successfully Signed Up!!"

- #### LOGIN

  <http://localhost:8000/users/login>

```json
{
  "email": "test@gmail.com",
  "password": "12345678"
}
```

Response:

```json
{
    "_id": "64c283e8593e19b180b38f20",
    "first_name": "Sanchar",
    "last_name": "Limbu",
    "password": "$2a$14$XfMCoZ9Tdk4F8HeRDaP1FO4VVQWySHgnwQxvRjo.w/0AeQQ6FD5/y",
    "email": "test@gmail.com",
    "phone": "9841111111",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6IlNhbmNoYXIiLCJMYXN0X05hbWUiOiJMaW1idSIsIlVpZCI6IjY0YzI4M2U4NTkzZTE5YjE4MGIzOGYyMCIsImV4cCI6MTY5MDU1NTc1Nn0.TgS9RlTcOmMT-LyPFOqdFswu4fbqkWS3CngVNFmW6sk",
    "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2OTEwNzQxNTZ9.JTJKeD1jNdHBmkeNy778aBqZB8IUzpjoVm0FC2CgSyI",
    "created_at": "2023-07-27T14:49:12Z",
    "updtaed_at": "2023-07-27T14:49:12Z",
    "user_id": "64c283e8593e19b180b38f20",
    "usercart": [],
    "address": [],
    "orders": []
}
```

- #### Admin add Product

  <http://localhost:8000/admin/addproduct>

```json
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg",
  "category": ["Laptop", "Gaming"]
}
```

Response : "Successfully added our Product Admin!!"

- #### View all Products

  <http://localhost:8000/users/productview>

Response

```json
[
    {
        "Product_ID": "64c215ae975cd1706d1492c5",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg",
        "category": [
            "Laptop"
        ]
    },
    {
        "Product_ID": "64c281df593e19b180b38f1f",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg",
        "category": [
            "Laptop",
            "Gaming"
        ]
    }
]
```

- #### Search Products

<http://localhost:8000/users/search?name=Al>

Response:

```json
[
    {
        "Product_ID": "64c215ae975cd1706d1492c5",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg",
        "category": [
            "Laptop"
        ]
    },
    {
        "Product_ID": "64c281df593e19b180b38f1f",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg",
        "category": [
            "Laptop",
            "Gaming"
        ]
    }
]
```

- #### Search Products by Category

  <http://localhost:8000/users/search-category?name=Gaming>

Response:

```json
[
    {
        "Product_ID": "64c281df593e19b180b38f1f",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg",
        "category": [
            "Laptop",
            "Gaming"
        ]
    }
]
```

- #### Add to Cart

  <http://localhost:8000/addtocart?id=64a3b75a227c961f850ef099&userID=64a3b752227c961f850ef098>

Response: "Successfully Added to the cart"

- #### List Cart

  <http://localhost:8000/listcart?id=64c283e8593e19b180b38f20>

Response:

```json
2500[
    {
        "Product_ID": "64c215ae975cd1706d1492c5",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg"
    }
]
```

- #### Remove Item from Cart

  <http://localhost:8000/removeitem?id=64c215ae975cd1706d1492c5&userID=64c283e8593e19b180b38f20>

Response: "Successfully removed from cart"

- #### Add Address

  <http://localhost:8000/addaddress?id=64c283e8593e19b180b38f20>

  The Address array is limited to a single value more addresses is not acceptable

```json
{
  "house_name": "house",
  "street_name": "street",
  "city_name": "bhaktapur",
  "pin_code": "69420"
}
```

Response: "Successfully Added the Address"

- #### Edit Address

  <http://localhost:8000/edithomeaddress?id=64c283e8593e19b180b38f20>

```json
{
    "house_name": "Lab",
    "street_name": "One",
    "city_name": "Crocas",
    "pin_code": "69420"
}
```

Response: "Successfully Updated the address"

- #### Delete Address

  <http://localhost:8000/deleteaddresses?id=64c283e8593e19b180b38f20>

Response: "Successfully Deleted!"

- #### Cart Checkout

  <http://localhost:8000/cartcheckout?id=64c283e8593e19b180b38f20>

Response: "Successfully Placed the order"

- #### Instantly Buy Products
  <http://localhost:8000/instantbuy?userid=64c283e8593e19b180b38f20&pid=64c215ae975cd1706d1492c5>

Response: "Successully placed the order"


### Contributing

You are free to use to this project as you desire.


### Acknowledgments

Special Thanks to [Prajwal Khatiwada](https://github.com/Prajwal-khatiwada) for frontend implementation.
