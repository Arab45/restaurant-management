# Restaurant Management API - Complete Endpoint Documentation

## API Base URL
```
http://localhost:3000/api/v1
```

## Overview
The Restaurant Management API is a comprehensive system for managing users, menus, food items, orders, tables, invoices, and notes in a restaurant environment.

---

## Endpoints Summary

### User Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/user` | User registration/sign up |
| GET | `/users` | Get all users with pagination |
| GET | `/user/{id}` | Get a specific user by ID |
| PUT | `/user-update/{id}` | Update user information |

### Food Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/food` | Create a new food item |
| GET | `/foods` | Get all food items with pagination |
| GET | `/food/{id}` | Get a specific food item |
| PUT | `/food-update/{id}` | Update a food item |

### Menu Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/menu` | Create a new menu |
| GET | `/menus` | Get all menus |
| GET | `/menu/{id}` | Get a specific menu |
| PUT | `/menu/{id}` | Update a menu |

### Table Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/table` | Create a new restaurant table |
| GET | `/tables` | Get all tables |
| GET | `/table/{id}` | Get a specific table |
| PUT | `/table/{id}` | Update a table |
| DELETE | `/table/{id}` | Delete a table |

### Order Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/order` | Create a new order |
| GET | `/orders` | Get all orders |
| GET | `/order/{id}` | Get a specific order |
| PUT | `/order/{id}` | Update an order |
| DELETE | `/order/{id}` | Delete an order |

### Order Item Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/orderItem` | Create order items for an order |
| GET | `/orderItems` | Get all order items |
| GET | `/orderItem/{id}` | Get order items by order ID |
| PUT | `/orderItem/{id}` | Update an order item |

### Invoice Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/invoice` | Create a new invoice |
| GET | `/invoices` | Get all invoices |
| GET | `/invoice/{id}` | Get a specific invoice |
| PUT | `/invoice/{id}` | Update an invoice |
| DELETE | `/invoice/{id}` | Delete an invoice |

### Note Management Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/note` | Create a new note |
| GET | `/notes` | Get all notes |
| GET | `/note/{id}` | Get a specific note |
| PUT | `/note/{id}` | Update a note |
| DELETE | `/note/{id}` | Delete a note |

---

## Detailed Endpoint Documentation

### User Endpoints

#### POST `/user` - Sign Up
**Description:** Create a new user account

**Request Body:**
```json
{
  "first_name": "string (required, 2-100 chars)",
  "last_name": "string (required, 2-100 chars)",
  "email": "string (required)",
  "password": "string (required, min 6 chars)",
  "phone": "string (required, 10-15 chars)",
  "avatar": "string (optional)"
}
```

**Success Response:** 200 OK
```json
{
  "InsertedID": "ObjectID"
}
```

**Error Responses:**
- 400: Bad request (validation error)
- 500: Internal server error

---

#### GET `/users` - Get All Users
**Description:** Retrieve a list of all users with pagination

**Query Parameters:**
- `recordPerPage` (optional, default: 10) - Number of records per page
- `page` (optional, default: 1) - Page number

**Success Response:** 200 OK
```json
{
  "total_count": number,
  "user_items": [...]
}
```

---

#### GET `/user/{id}` - Get User by ID
**Description:** Retrieve user details by user ID

**Path Parameters:**
- `id` (string, required) - User ID

**Success Response:** 200 OK
```json
{
  "id": "ObjectID",
  "first_name": "string",
  "last_name": "string",
  "email": "string",
  "phone": "string",
  "avatar": "string",
  "token": "string",
  "refresh_token": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "user_id": "string"
}
```

---

#### PUT `/user-update/{id}` - Update User
**Description:** Update specific user details

**Path Parameters:**
- `id` (string, required) - User ID

**Request Body:** User object with fields to update

**Success Response:** 200 OK

---

### Food Endpoints

#### POST `/food` - Create Food
**Description:** Create a new food item

**Request Body:**
```json
{
  "name": "string (required, 2-100 chars)",
  "price": "number (required)",
  "food_image": "string (required)",
  "menu_id": "string (required)"
}
```

**Success Response:** 201 Created
```json
{
  "InsertedID": "ObjectID"
}
```

---

#### GET `/foods` - Get All Foods
**Description:** Retrieve a paginated list of all food items

**Query Parameters:**
- `recordPerPage` (optional, default: 10)
- `page` (optional, default: 1)

**Success Response:** 200 OK
```json
{
  "total_count": number,
  "food_items": [...]
}
```

---

#### GET `/food/{id}` - Get Food by ID
**Description:** Retrieve a food item by its ID

**Path Parameters:**
- `id` (string, required) - Food ID

**Success Response:** 200 OK

---

#### PUT `/food-update/{id}` - Update Food
**Description:** Update food item details

**Path Parameters:**
- `id` (string, required) - Food ID

**Request Body:** Food object with fields to update
```json
{
  "name": "string (optional)",
  "price": "number (optional)",
  "food_image": "string (optional)",
  "menu_id": "string (optional)"
}
```

**Success Response:** 200 OK

---

### Menu Endpoints

#### POST `/menu` - Create Menu
**Description:** Create a new menu

**Request Body:**
```json
{
  "name": "string (required)",
  "category": "string (required)",
  "start_date": "timestamp (optional)",
  "end_date": "timestamp (optional)"
}
```

**Success Response:** 200 OK

---

#### GET `/menus` - Get All Menus
**Description:** Retrieve a list of all menus

**Success Response:** 200 OK
```json
[...]
```

---

#### GET `/menu/{id}` - Get Menu by ID
**Description:** Retrieve menu details by menu ID

**Path Parameters:**
- `id` (string, required) - Menu ID

**Success Response:** 200 OK

---

#### PUT `/menu/{id}` - Update Menu
**Description:** Update menu details

**Path Parameters:**
- `id` (string, required) - Menu ID

**Request Body:** Menu object with fields to update

**Success Response:** 200 OK

---

### Table Endpoints

#### POST `/table` - Create Table
**Description:** Create a new restaurant table

**Request Body:**
```json
{
  "table_number": "integer (required)",
  "number_of_guests": "integer (required)"
}
```

**Success Response:** 200 OK

---

#### GET `/tables` - Get All Tables
**Description:** Retrieve a list of all restaurant tables

**Success Response:** 200 OK

---

#### GET `/table/{id}` - Get Table by ID
**Description:** Retrieve table details by table ID

**Path Parameters:**
- `id` (string, required) - Table ID

**Success Response:** 200 OK

---

#### PUT `/table/{id}` - Update Table
**Description:** Update table details

**Path Parameters:**
- `id` (string, required) - Table ID

**Request Body:** Table object with fields to update

**Success Response:** 200 OK

---

#### DELETE `/table/{id}` - Delete Table
**Description:** Delete a table by table ID

**Path Parameters:**
- `id` (string, required) - Table ID

**Success Response:** 200 OK

---

### Order Endpoints

#### POST `/order` - Create Order
**Description:** Create a new order for a table

**Request Body:**
```json
{
  "order_date": "timestamp (required)",
  "table_id": "string (required)"
}
```

**Success Response:** 200 OK

---

#### GET `/orders` - Get All Orders
**Description:** Retrieve a list of all orders

**Success Response:** 200 OK

---

#### GET `/order/{id}` - Get Order by ID
**Description:** Retrieve order details by order ID

**Path Parameters:**
- `id` (string, required) - Order ID

**Success Response:** 200 OK

---

#### PUT `/order/{id}` - Update Order
**Description:** Update order details

**Path Parameters:**
- `id` (string, required) - Order ID

**Request Body:** Order object with fields to update

**Success Response:** 200 OK

---

#### DELETE `/order/{id}` - Delete Order
**Description:** Delete an order by order ID

**Path Parameters:**
- `id` (string, required) - Order ID

**Success Response:** 200 OK

---

### Order Item Endpoints

#### POST `/orderItem` - Create Order Items
**Description:** Create one or more order items associated with an order

**Request Body:**
```json
{
  "table_id": "string",
  "order_item": [
    {
      "quantity": "string (required, S|M|L)",
      "unit_price": "number (required)",
      "food_id": "string (required)"
    }
  ]
}
```

**Success Response:** 200 OK

---

#### GET `/orderItems` - Get All Order Items
**Description:** Retrieve a list of all order items

**Success Response:** 200 OK

---

#### GET `/orderItem/{id}` - Get Order Items by Order ID
**Description:** Retrieve all items for a specific order with details

**Path Parameters:**
- `id` (string, required) - Order ID

**Success Response:** 200 OK

---

#### PUT `/orderItem/{id}` - Update Order Item
**Description:** Update order item details

**Path Parameters:**
- `id` (string, required) - Order Item ID

**Request Body:** Order item object with fields to update

**Success Response:** 200 OK

---

### Invoice Endpoints

#### POST `/invoice` - Create Invoice
**Description:** Create a new invoice for an order

**Request Body:**
```json
{
  "order_id": "string (required)",
  "payment_method": "string (optional, CARD|CASH)",
  "payment_status": "string (optional, PENDING|PAID)"
}
```

**Success Response:** 200 OK

---

#### GET `/invoices` - Get All Invoices
**Description:** Retrieve a list of all invoices

**Success Response:** 200 OK

---

#### GET `/invoice/{id}` - Get Invoice by ID
**Description:** Retrieve invoice details with order items

**Path Parameters:**
- `id` (string, required) - Invoice ID

**Success Response:** 200 OK
```json
{
  "invoice_id": "string",
  "order_id": "string",
  "payment_method": "string",
  "payment_status": "string",
  "payment_due": "number",
  "payment_due_date": "timestamp",
  "table_number": "number",
  "order_details": [...]
}
```

---

#### PUT `/invoice/{id}` - Update Invoice
**Description:** Update invoice payment status and method

**Path Parameters:**
- `id` (string, required) - Invoice ID

**Request Body:**
```json
{
  "payment_method": "string (CARD|CASH)",
  "payment_status": "string (PENDING|PAID)"
}
```

**Success Response:** 200 OK

---

#### DELETE `/invoice/{id}` - Delete Invoice
**Description:** Delete an invoice by invoice ID

**Path Parameters:**
- `id` (string, required) - Invoice ID

**Success Response:** 200 OK

---

### Note Endpoints

#### POST `/note` - Create Note
**Description:** Create a new note

**Request Body:**
```json
{
  "title": "string (required)",
  "text": "string (required)"
}
```

**Success Response:** 200 OK

---

#### GET `/notes` - Get All Notes
**Description:** Retrieve a list of all notes

**Success Response:** 200 OK

---

#### GET `/note/{id}` - Get Note by ID
**Description:** Retrieve note details by note ID

**Path Parameters:**
- `id` (string, required) - Note ID

**Success Response:** 200 OK

---

#### PUT `/note/{id}` - Update Note
**Description:** Update note details

**Path Parameters:**
- `id` (string, required) - Note ID

**Request Body:** Note object with fields to update

**Success Response:** 200 OK

---

#### DELETE `/note/{id}` - Delete Note
**Description:** Delete a note by note ID

**Path Parameters:**
- `id` (string, required) - Note ID

**Success Response:** 200 OK

---

## Data Models

### User Model
```json
{
  "_id": "ObjectID",
  "first_name": "string",
  "last_name": "string",
  "email": "string",
  "password": "string (hashed)",
  "phone": "string",
  "avatar": "string",
  "token": "string (JWT)",
  "refresh_token": "string (JWT)",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "user_id": "string"
}
```

### Food Model
```json
{
  "_id": "ObjectID",
  "name": "string",
  "price": "number",
  "food_image": "string (URL)",
  "menu_id": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "food_id": "string"
}
```

### Menu Model
```json
{
  "_id": "ObjectID",
  "name": "string",
  "category": "string",
  "start_date": "timestamp",
  "end_date": "timestamp",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "menu_id": "string"
}
```

### Table Model
```json
{
  "_id": "ObjectID",
  "table_number": "integer",
  "number_of_guests": "integer",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "table_id": "string"
}
```

### Order Model
```json
{
  "_id": "ObjectID",
  "order_date": "timestamp",
  "table_id": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "order_id": "string"
}
```

### OrderItem Model
```json
{
  "_id": "ObjectID",
  "quantity": "string (S|M|L)",
  "unit_price": "number",
  "food_id": "string",
  "order_id": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "order_item_id": "string"
}
```

### Invoice Model
```json
{
  "_id": "ObjectID",
  "invoice_id": "string",
  "order_id": "string",
  "payment_method": "string (CARD|CASH)",
  "payment_status": "string (PENDING|PAID)",
  "payment_due_date": "timestamp",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Note Model
```json
{
  "_id": "ObjectID",
  "title": "string",
  "text": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "note_id": "string"
}
```

---

## Swagger UI
Access the interactive Swagger UI documentation at:
```
http://localhost:3000/swagger/index.html
```

## Generated Files
- `docs/swagger.json` - JSON format API specification
- `docs/swagger.yaml` - YAML format API specification
- `docs/docs.go` - Go package containing embedded Swagger documentation

---

## Notes
- All timestamps are in RFC3339 format
- All IDs are MongoDB ObjectID hex strings
- Pagination defaults: recordPerPage=10, page=1
- Authentication tokens are JWT format
- Password is hashed using bcrypt
