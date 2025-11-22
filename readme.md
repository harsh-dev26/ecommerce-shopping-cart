# E-Commerce Shopping Cart Application

A full-stack e-commerce shopping cart application built with Go (Gin + GORM) backend and React frontend.

## Features

- User authentication (signup/login)
- Browse items
- Add items to cart
- Checkout and create orders
- View cart and order history

## Tech Stack

**Backend:**
- Go
- Gin Web Framework
- GORM (SQLite)
- JWT Authentication

**Frontend:**
- React
- Axios
- CSS3

## Prerequisites

- Go 1.19 or higher
- Node.js 14 or higher
- npm or yarn

## Installation & Setup

### Backend Setup

1. Navigate to backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The backend server will start on `http://localhost:8080`

4. (Optional) Seed sample data:
```bash
go run seed.go
```

### Frontend Setup

1. Navigate to frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm start
```

The frontend will open at `http://localhost:3000`

## API Endpoints

### User Endpoints
- `POST /users` - Create a new user
- `GET /users` - List all users
- `POST /users/login` - Login user

### Item Endpoints
- `POST /items` - Create an item
- `GET /items` - List all items

### Cart Endpoints (Protected)
- `POST /carts` - Create/update cart
- `GET /carts` - List all carts

### Order Endpoints (Protected)
- `POST /orders` - Create order from cart
- `GET /orders` - List all orders

## Usage

1. **Create a user** (use Postman or the API):
```bash
POST http://localhost:8080/users
{
  "username": "testuser",
  "password": "password123"
}
```

2. **Login** via the web interface with your credentials

3. **Browse items** and add them to your cart

4. **Checkout** to convert your cart to an order

5. **View cart** and **order history** using the buttons

## Project Structure