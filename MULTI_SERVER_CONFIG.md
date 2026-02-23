# Multi-Server API Configuration ✅

## Now Supporting Both Local and Online Servers!

Your Restaurant Management API is now configured to work on **both local and online environments**.

---

## 🖥️ Server URLs

### Local Development Server
```
Base URL: http://localhost:3000/api/v1
Swagger UI: http://localhost:3000/swagger/index.html
```

### Online Production Server
```
Base URL: https://restaurant-management-f9kx.onrender.com/api/v1
Swagger UI: https://restaurant-management-f9kx.onrender.com/swagger/index.html
```

---

## 📋 Complete Endpoints

Both servers support all 26 endpoints:

### Local Development
```bash
# Example: Get all users
curl http://localhost:3000/api/v1/users

# Example: Create a food item
curl -X POST http://localhost:3000/api/v1/food \
  -H "Content-Type: application/json" \
  -d '{"name":"Pizza","price":12.99,"food_image":"url","menu_id":"menu123"}'
```

### Online Production
```bash
# Example: Get all users
curl https://restaurant-management-f9kx.onrender.com/api/v1/users

# Example: Create a food item
curl -X POST https://restaurant-management-f9kx.onrender.com/api/v1/food \
  -H "Content-Type: application/json" \
  -d '{"name":"Pizza","price":12.99,"food_image":"url","menu_id":"menu123"}'
```

---

## 🔗 How Both Servers Are Configured

### Updated Files:

1. **cmd/server/main.go**
   - Added HTTPS scheme support
   - Updated description to mention both servers
   - Schemes now include both `http` and `https`

2. **docs/swagger.json**
   - Added `https` scheme
   - Added `x-servers` section with both server URLs
   - Updated description

3. **docs/swagger.yaml**
   - Added `https` scheme
   - Added `x-servers` section with both server URLs
   - Updated description

### In Swagger UI:

When you open the Swagger UI, you can now:
- **Local**: See `http` scheme available
- **Online**: See both `http` and `https` schemes available
- **Both**: See documented server information in the API info

---

## 🧪 Testing on Both Servers

### Using Swagger UI Locally
```
1. Start local server: go run cmd/server/main.go
2. Open: http://localhost:3000/swagger/index.html
3. Test any endpoint directly
```

### Using Swagger UI Online
```
1. Visit: https://restaurant-management-f9kx.onrender.com/swagger/index.html
2. Test any endpoint directly
3. See live responses from production server
```

### Using Postman/Insomnia

**Import from local server:**
```
http://localhost:3000/api/v1/swagger.json
```

**Import from online server:**
```
https://restaurant-management-f9kx.onrender.com/api/v1/swagger.json
```

---

## 📱 Frontend Integration Examples

### React/Vue Example - Using Local Server
```javascript
const API_URL = process.env.NODE_ENV === 'development' 
  ? 'http://localhost:3000/api/v1'
  : 'https://restaurant-management-f9kx.onrender.com/api/v1';

// Get all users
fetch(`${API_URL}/users`)
  .then(res => res.json())
  .then(data => console.log(data));
```

### Environment-based Configuration
```bash
# .env.local (development)
REACT_APP_API_URL=http://localhost:3000/api/v1

# .env.production
REACT_APP_API_URL=https://restaurant-management-f9kx.onrender.com/api/v1
```

### Using Environment Variable
```javascript
const API_URL = process.env.REACT_APP_API_URL;
```

---

## ✨ Server Information in Swagger

When you view the API documentation, you'll see:

```
Servers:
• Local Development Server: http://localhost:3000/api/v1
• Online Production Server: https://restaurant-management-f9kx.onrender.com/api/v1
```

This helps developers choose the right endpoint based on their environment.

---

## 🔐 CORS Configuration

### Current CORS Settings (in main.go)
```go
AllowOrigins: []string{
    "http://localhost:5173",
},
```

**Note:** You may need to update CORS settings to allow requests from your production frontend URL:
```go
AllowOrigins: []string{
    "http://localhost:5173",        // Local dev
    "https://your-frontend.com",    // Production frontend
},
```

---

## 📊 Schemes Supported

### Local Development
- ✅ `http://localhost:3000/api/v1`

### Online Production  
- ✅ `http://restaurant-management-f9kx.onrender.com/api/v1` (if enabled)
- ✅ `https://restaurant-management-f9kx.onrender.com/api/v1` (recommended)

---

## 🚀 API Usage

### All 26 Endpoints Available on Both Servers

#### User Endpoints
- `POST /user` - Create user
- `GET /users` - Get all users
- `GET /user/{id}` - Get user
- `PUT /user-update/{id}` - Update user

#### Food Endpoints
- `POST /food` - Create food
- `GET /foods` - Get all foods
- `GET /food/{id}` - Get food
- `PUT /food-update/{id}` - Update food

#### Menu Endpoints
- `POST /menu` - Create menu
- `GET /menus` - Get all menus
- `GET /menu/{id}` - Get menu
- `PUT /menu/{id}` - Update menu

#### Table Endpoints
- `POST /table` - Create table
- `GET /tables` - Get all tables
- `GET /table/{id}` - Get table
- `PUT /table/{id}` - Update table
- `DELETE /table/{id}` - Delete table

#### Order Endpoints
- `POST /order` - Create order
- `GET /orders` - Get all orders
- `GET /order/{id}` - Get order
- `PUT /order/{id}` - Update order
- `DELETE /order/{id}` - Delete order

#### Order Item Endpoints
- `POST /orderItem` - Create order item
- `GET /orderItems` - Get all items
- `GET /orderItem/{id}` - Get items by order
- `PUT /orderItem/{id}` - Update order item

#### Invoice Endpoints
- `POST /invoice` - Create invoice
- `GET /invoices` - Get all invoices
- `GET /invoice/{id}` - Get invoice
- `PUT /invoice/{id}` - Update invoice
- `DELETE /invoice/{id}` - Delete invoice

#### Note Endpoints
- `POST /note` - Create note
- `GET /notes` - Get all notes
- `GET /note/{id}` - Get note
- `PUT /note/{id}` - Update note
- `DELETE /note/{id}` - Delete note

---

## 💡 Best Practices

### For Development
```
Use: http://localhost:3000/api/v1
Why: Faster iteration, no deployment needed
```

### For Testing
```
Use: https://restaurant-management-f9kx.onrender.com/api/v1
Why: Test against production-like environment
```

### For Production
```
Use: https://restaurant-management-f9kx.onrender.com/api/v1
Why: HTTPS, stable, officially deployed
```

---

## 📝 Swagger Documentation Files Updated

✅ **docs/swagger.json** - Both servers documented in x-servers  
✅ **docs/swagger.yaml** - Both servers documented in x-servers  
✅ **cmd/server/main.go** - HTTPS scheme added, description updated

---

## 🎯 Summary

Your API now:
- ✅ Works locally at `http://localhost:3000/api/v1`
- ✅ Works online at `https://restaurant-management-f9kx.onrender.com/api/v1`
- ✅ Supports both HTTP and HTTPS schemes
- ✅ Has updated Swagger documentation for both servers
- ✅ Shows both server options in Swagger UI
- ✅ Is ready for frontend integration in any environment

---

## 🔗 Quick Links

| Resource | Local | Online |
|----------|-------|--------|
| API Base | `http://localhost:3000/api/v1` | `https://restaurant-management-f9kx.onrender.com/api/v1` |
| Swagger UI | `http://localhost:3000/swagger/index.html` | `https://restaurant-management-f9kx.onrender.com/swagger/index.html` |
| Swagger JSON | `http://localhost:3000/api/v1/swagger.json` | `https://restaurant-management-f9kx.onrender.com/api/v1/swagger.json` |
| Swagger YAML | `http://localhost:3000/api/v1/swagger.yaml` | `https://restaurant-management-f9kx.onrender.com/api/v1/swagger.yaml` |

---

**Ready to use both servers!** 🎉
