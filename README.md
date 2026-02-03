# ✅ Login API (MVC Architecture)

A simple login API endpoint built using a **scalable MVC architecture**.
It only handles **login**, not signup.

Upon successful login, the endpoint returns a **JWT token** signed with a secret key and valid for **60 minutes**.

---

## ✅ Features

### **Architecture**

* Scalable MVC structure with clear separation of concerns.
* Modular design.

### **Dependency Injection**

* Uses DI for clean dependency management and testability.

### **Interfaces**

* Uses interfaces for loose coupling between layers and services.

### **Error Handling**

* Proper error handling using a custom error structure `AppError`.

---

## ✅ Requirements

* Go **1.24.x**

---

## ✅ Setup

### 1. Clone the repository

```bash
git clone https://github.com/ashishDevv/login-task.git
```

### 2. Install dependencies

```bash
go download
```

### 3. Create `.env` file

```env
APP_PORT=8080
DB_URL=""
AUTH_SECRET=supersecretkey
AUTH_EXPIRY_MIN=60
```

### 4. Build the project

```bash
go build cmd/main.go
```

### 5. Run the server

```bash
./main
```

---

## ✅ Routes

### **POST /login**

#### Request

```http
POST /login
Content-Type: application/json

{
  "email": "ashu@gmail.com",
  "password": "12345"
}
```

#### Response

```json
{
  "id": 1,
  "token": "JWT token with 60min expiry"
}
```

---
