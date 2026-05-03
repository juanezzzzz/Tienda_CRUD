# 🛒 API REST - Tienda (Go + PostgreSQL)

API REST desarrollada en Go para la gestión completa de una tienda online. Permite administrar productos, carritos de compra, items, pagos, facturación, cupones y transacciones.

---

## 🚀 Descripción

Este proyecto implementa un backend tipo e-commerce utilizando arquitectura modular. Está diseñado para simular un entorno real de negocio, donde se gestionan múltiples entidades relacionadas como productos, métodos de pago y procesos de compra.

---

## 🧱 Arquitectura del proyecto

El proyecto sigue una estructura organizada por capas:

- **config/** → conexión y configuración de la base de datos  
- **models/** → definición de entidades y estructuras  
- **controllers/** → manejo de solicitudes HTTP (CRUD)  
- **routes/** → definición y agrupación de endpoints  
- **main.go** → punto de entrada del servidor  

---

## 🛠️ Tecnologías utilizadas

- Go (Golang)
- Gorilla Mux
- PostgreSQL
- Git & GitHub

---

## 🔐 Variables de entorno

Antes de ejecutar el proyecto, configura las siguientes variables:


DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=tienda_db


---

## ▶️ Ejecución del proyecto

1. Clonar el repositorio:


git clone https://github.com/juanezzzzz/Tienda_CRUD.git

cd Tienda_CRUD


2. Instalar dependencias:


go mod tidy


3. Ejecutar el servidor:


go run main.go


El servidor se iniciará en:


http://localhost:8082


---

## 📡 Endpoints principales

### 🛍️ Productos
- GET /productos  
- POST /productos  
- PUT /productos/{id}  
- DELETE /productos/{id}  

### 🛒 Carritos
- GET /carritos  
- POST /carritos  

### 📦 Items de carrito
- GET /carrito-items  
- POST /carrito-items  

### 💳 Pagos
- POST /pagos  

### 🧾 Facturación
- POST /facturacion  

### 🎟️ Cupones
- GET /cupones  

### 💰 Transacciones
- GET /transacciones  

---

## 🧠 Características

- API REST modular y escalable  
- Manejo de múltiples entidades relacionadas  
- Conexión a base de datos PostgreSQL  
- Estructura preparada para futuras mejoras (autenticación, servicios, etc.)

---

## 📅 Desarrollo

Proyecto desarrollado entre el **1 y el 3 de mayo de 2026**, aplicando un flujo de trabajo basado en Git:

- **develop** → rama de desarrollo  
- **feature/** → desarrollo por módulos  
- **main** → versión estable  

---

## 👨‍💻 Autor

**Juan Valencia**

---

## 🔮 Mejoras futuras

- Implementación de autenticación JWT  
- Separación en capas (services / repositories)  
- Validación de datos y manejo de errores avanzado  
- Dockerización del proyecto  