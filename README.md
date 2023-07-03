# go_project

- **SIGNUP FUNCTION API CALL (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
  "first_name": "Sanchar",
  "last_name": "Limbu",
  "email": "limbu@gmail.com",
  "password": "12345678",
  "phone": "9863325642"
}
```

Response :"Successfully Signed Up!!"

- **LOGIN FUNCTION API CALL (POST REQUEST)**

  http://localhost:8000/users/login

```json
{
  "email": "limbu@gmail.com",
  "password": "12345678"
}
```

response will be like this

```json
{
    "_id": "64a2c84bb2c979fdae2974f4",
    "first_name": "Sanchar",
    "last_name": "Limbu",
    "password": "$2a$14$SIU4rS4sqmwQ4p71dCjjHuUz4UO50QyiGxjOU.8DFCL1oHH0anayi",
    "email": "test@gmail.com",
    "phone": "9863325642",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6IlNhbmNoYXIiLCJMYXN0X05hbWUiOiJMaW1idSIsIlVpZCI6IjY0YTJjODRiYjJjOTc5ZmRhZTI5NzRmNCIsImV4cCI6MTY4ODQ3NjEwN30.RUWzhoJhMUVqqNn8usWr1C99yj8d9Fn8ZcgRZ_SujGc",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2ODg5OTQ1MDd9.NGPO-PYfWztVp3RVS_QW7wUXZCM4ffYrEtF_5EOxH68",
    "created_at": "2023-07-03T13:08:27Z",
    "updated_at": "2023-07-03T13:08:27Z",
    "user_id": "64a2c84bb2c979fdae2974f4",
    "user_cart": [],
    "address": [],
    "orders": []
}
```

- **Admin add Product Function (POST REQUEST)**

  http://localhost:8000/admin/addproduct

```json
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
```

Response : "Successfully added our Product Admin!!"

- **View all the Products in db GET REQUEST**

  http://localhost:8000/users/productview

Response

```json
[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "alienwarex15",
    "price": 1500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "giner ale",
    "price": 900,
    "rating": 5,
    "image": "gin.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "whiskey",
    "price": 100,
    "rating": 7,
    "image": "whis.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "acer predator",
    "price": 3000,
    "rating": 10,
    "image": "acer.jpg"
  }
]
```

- **Search Product by regex function (GET REQUEST)**

defines the word search sorting
http://localhost:8000/users/search?name=al

response:

```json
[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 1500,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "ginger Ale",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
