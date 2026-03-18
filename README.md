# 🚀 VenusX: Motor TPV & Sistema Logístico Backend

VenusX es una API RESTful desarrollada en **Go (Golang)** diseñada para solucionar ineficiencias logísticas reales en el sector del retail y estaciones de servicio de alto volumen. 

Este proyecto nace de la observación directa ("Domain Knowledge") de los cuellos de botella en los sistemas TPV tradicionales, proponiendo una arquitectura backend moderna, modular y preparada para alta concurrencia.

## 🧠 Arquitectura y Diseño

El sistema está construido siguiendo los principios de **Clean Architecture**, separando estrictamente las responsabilidades para garantizar un código escalable, mantenible y altamente testeable:

- **Capa de Rutas y Controladores:** Manejo de peticiones HTTP con el framework `Gin`.
- **Capa de Casos de Uso (Lógica de Negocio):** Reglas del sistema de inventario y ventas.
- **Capa de Repositorios (Persistencia):** Interacción con la base de datos PostgreSQL utilizando el ORM `GORM`.

## 🛠️ Stack Tecnológico

* **Lenguaje:** Go (Golang)
* **Framework Web:** Gin Gonic
* **Base de Datos:** PostgreSQL
* **ORM:** GORM (con Auto-Migrations)
* **Testing:** Paquete nativo `testing` + `Testify` (Mocks e Inyección de Dependencias)
* **Frontend Integrado:** React / Electron (Desarrollo colaborativo)

## 🔥 Características Principales

* **Búsqueda Optimizada:** Endpoints de baja latencia para la búsqueda de artículos por código de barras.
* **Gestión de Carritos:** Lógica transaccional para evitar inconsistencias de stock durante cobros paralelos.
* **Seguridad:** Autenticación de usuarios, gestión de variables de entorno seguras y configuración estricta de CORS.
* **Testing Automatizado:** Arquitectura preparada para cobertura de pruebas unitarias y de integración, simulando flujos reales de caja.

## ⚙️ Instalación y Uso Local

1. Clona el repositorio:
   ```bash
   git clone https://github.com/bryanjue/VenusX-api.git
