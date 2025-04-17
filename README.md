# 📚 School Management API

![School API Logo](https://miro.medium.com/v2/resize:fit:2000/format:webp/1*OcmVkcsM5BWRHrg8GC17iw.png)

A simple RESTful API for managing students in a school database using **Go (Golang)**, **PostgreSQL**, and **Go chi**.

## 🚀 Features

- **🔄 CRUD Operations:** Create, Read, Update, Delete students.
- **📜 Pagination:** Efficiently handle large datasets.
- **📦 JSON-based API Responses:** Standardized data format for easy consumption.
- **✅ Input Validation:** Ensure data integrity before processing.
- **⚡ Bulk Insert:** Efficiently insert multiple records in one request.
- **🔒 Environment Variables:** Securely manage database connection details.
- **🗄️ PostgreSQL Database Connection:** Persistent data storage with PostgreSQL.
- **🚯 IP-Based Rate Limiting:** Prevent abuse by limiting requests per IP.
- **🔑 API Key Authentication:** Protect endpoints with API Key middleware.
- **🛑 Graceful Shutdown:** Ensure smooth termination of the API.
- **🐳 Docker Support:** Easily deploy the application using Docker.
- **☸️ Kubernetes Deployment:** Scalable and manageable deployment in Kubernetes.
- **🌐 Docker Hub Integration:** Uploaded images are available on Docker Hub.

---

## 🏗️ Tech Stack

- **Backend:** Go (Golang), GO chi
- **Database:** PostgreSQL
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **Logging:** Log Package
- **API Format:** RESTful, JSON

---

## 📂 Project Structure

```
handlers
  ├── handlers.go     # API Endpoints
models
  ├── models.go       # Student struct
validation
  ├── validation.go   # Input Validation

main.go               # Entry Point
.env                  # Environment Variables
README.md             # Project Documentation
LICENSE               # License Information
go.mod                # Module Dependencies
go.sum                # Dependency Checksums
dockerfile            # Dockerfile for building the image
docker-compose.yaml   # Docker Compose file for multi-container setup
dockerfile_web-server:1.0 
dockerfile_web-server-ubuntu:1.0
kubernetes/           # Basic k8s configuration files
  ├── postgres.yaml
  ├── postgres-config.yaml
  ├── postgres-secret.yaml
  ├── server.yaml

kubernetes_updated/   # Kubernetes configuration updated files
  ├── auth-secret-volume.yaml  # Volume for API key secrets
  ├── ingress-nginx.yaml       # Ingress controller
  ├── postgres.yaml            # PostgreSQL deployment
  ├── persistent-volume.yaml   # Persistent volume for PostgreSQL
  ├── postgres-config.yaml     # ConfigMap for PostgreSQL
  ├── postgres-secret.yaml     # Secrets for PostgreSQL credentials
  ├── server.yaml              # Go server deployment with replicas
  ├── server-ingress.yaml      # Ingress for Go server

```

---

## 📦 Installation & Setup

### Prerequisites
- Install **Go (v1.24 or latest)**
- Install and set up **PostgreSQL**
- Install **Docker** and **Kubernetes**

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/nsshohag/school_api_postgres.git
cd school_api_postgres
```

### 2️⃣ Configure Environment Variables
If you don't use docker-compose or kubernetes then, create a `.env` file in the root directory add your PostgreSQL credentials:
```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=sadat
DB_PASSWORD=11235813
DB_NAME=school_db
API_KEY=sadat-api-key-1123
```

> **Note**: Update `.env` with your actual database credentials and API key. For Docker or Kubernetes, these are managed via secrets.

### 3️⃣ Install Dependencies
```sh
go mod tidy
```

### 4️⃣ Set Up the Database
#### Prerequisites
- PostgreSQL installed
- Command-line access

#### Steps
1. **Start PostgreSQL**  
   ```bash
   sudo service postgresql start
   ```

2. **Log in**  
   ```bash
   psql -h localhost -p 5433 -U sadat -d postgres
   ```
   Password: `11235813`

3. **Create Database**  
   ```sql
   CREATE DATABASE school_db;
   ```

4. **Connect**  
   ```sql
   \c school_db
   ```

5. **Create Table**  
   ```sql
   CREATE TABLE students (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       age INT NOT NULL,
       class INTEGER NOT NULL CHECK (class BETWEEN 1 AND 10)
   );
   ```

#### Troubleshooting
- **Connection Issues**: Check service, credentials, and port.
- **Permissions**: Ensure `sadat` has proper privileges.
- **Port Conflict**: Update `DB_PORT` if `5433` is in use.

#### Next Steps
- Insert data into `students`.
- Connect application to database.

### 4️⃣ Run the Application
```sh
go run main.go
```

The server will start at `http://localhost:8080`

---

### 6️⃣ Run Docker Compose
For multi-container setup, you can use Docker Compose:
```sh
docker-compose up -d
```

---

## ☸️ Kubernetes Deployment

The project includes comprehensive Kubernetes configuration for robust, scalable deployment in production environments.

### 🧩 Kubernetes Components

- **ConfigMap:** Store non-sensitive configuration data with `postgres-config.yaml`
- **Secrets:** Securely manage sensitive data like database credentials with `postgres-secret.yaml`
- **Persistent Volume:** Ensure data persistence across pod restarts with `persistent-volume.yaml`
- **Deployment:** Deploy the PostgreSQL database and API server with replicas
- **Services:** Expose the API and database internally and externally
- **Ingress:** Route external traffic to services with `server-ingress.yaml` and `ingress-nginx.yaml`
- **Volume Mounts:** Mount configuration and secrets securely with `auth-secret-volume.yaml`

All resources are in the `school-system` namespace.

### 🚢 Deployment Steps

#### Prerequisites
- **Kubernetes Cluster**: Use Minikube/Kind for local testing or a cloud provider (e.g., GKE, EKS, AKS).

- **kubectl**: Installed and configured to interact with your cluster.

- **Docker Hub Access**: Ensure the `nsshohag/web-server-without-dot-env-auth:1.0` image is accessible.


#### 1️⃣ Create Namespace

```sh
kubectl create namespace school-system
```

#### 2️⃣ Apply Persistent Volume Configuration

```sh
kubectl apply -f kubernetes_updated/persistent-volume.yaml
```

#### 3️⃣ Apply ConfigMap and Secrets

```sh
kubectl apply -f kubernetes_updated/postgres-config.yaml
kubectl apply -f kubernetes_updated/postgres-secret.yaml
kubectl apply -f kubernetes_updated/auth-secret-volume.yaml
```

#### 4️⃣ Deploy PostgreSQL Database

```sh
kubectl apply -f kubernetes_updated/postgres.yaml
```

#### 5️⃣ Create Database Table

To create the `students` table in PostgreSQL:

1. **Find the PostgreSQL Pod**:

   ```sh
   kubectl get pods -n school-system -l app=postgres
   ```

   Example output:

   ```
   NAME                                 READY   STATUS    RESTARTS   AGE
   postgres-deployment-5f8b6c7d8-xyz   1/1     Running   0          5m
   ```

2. **Access the Pod**:

   ```sh
   kubectl exec -it postgres-deployment-5f8b6c7d8-xyz -n school-system -- bash
   ```

3. **Connect to PostgreSQL**:

   ```sh
   psql -U sadat -d school_db
   ```

   > **Note**: Use values from `postgres-secret` for username and database.

4. **Run SQL Query**:

   ```sql
   CREATE TABLE students (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       age INT NOT NULL,
       class INTEGER NOT NULL CHECK (class BETWEEN 1 AND 10)
   );
   ```

5. **Verify Table**:

   ```sql
   \dt
   ```

6. **Exit**:

   ```sql
   \q
   exit
   ```

#### 6️⃣ Deploy API Server with 3 Replicas

```sh
kubectl apply -f kubernetes_updated/server.yaml
```

#### 7️⃣ Set Up Services and Ingress

```sh
kubectl apply -f kubernetes_updated/ingress-nginx.yaml
kubectl apply -f kubernetes_updated/server-ingress.yaml
```

### 📡 Sending Requests to the API

After deployment, send requests to the API using the `X-API-Key` header:

```
-H "X-API-Key: sadat-api-key-1123"
```

#### Option 1: Via Ingress (Production)

If using `sadat.com` (from `server-ingress.yaml`):

1. **Configure DNS**: Point `sadat.com` to the Ingress IP. For Minikube:

   ```sh
   minikube ip
   ```

2. **Send Request**:

   Get all students:

   ```sh
   curl -H "X-API-Key: sadat-api-key-1123" http://sadat.com/api/v1/students
   ``` 
   **Response:**
```json
[
  {
    "id": 1,
    "name": "Nazmus Sadat Shohag",
    "age": 24,
    "class": 10
  },
    {
    "id": 2,
    "name": "SH Rony",
    "age": 24,
    "class": 10
  }
]
```   
3. **Troubleshooting**: If `sadat.com` doesn’t resolve, get the Ingress IP:

   ```sh
   kubectl get ingress -n school-system
   ```

   Use: `http://<ingress-ip>/api/v1/students`.

#### Option 2: Via NodePort (Testing)

Use the `server-service-node` NodePort (30100):

1. **Get Node IP**:

   For Minikube:

   ```sh
   minikube ip
   ```

   For cloud clusters:

   ```sh
   kubectl get nodes -o wide
   ```

2. **Send Request**:

   ```sh
   curl -H "X-API-Key: sadat-api-key-1123" http://<node-ip>:30100/api/v1/students
   ```

#### Option 3: Port-Forwarding (Local Testing)

For local development:

1. **Port-Forward**:

   ```sh
   kubectl port-forward svc/server-service 8080:8080 -n school-system
   ```

2. **Send Request**:

   ```sh
   curl -H "X-API-Key: sadat-api-key-1123" http://localhost:8080/api/v1/students
   ```

3. **Stop**: Press `Ctrl+C`.

### 📊 Monitoring and Scaling

- **Check Deployment Status**:

  ```sh
  kubectl get deployments -n school-system
  kubectl get pods -n school-system
  kubectl get services -n school-system
  ```

- **Scale Replicas**:

  ```sh
  kubectl scale deployment server-deployment --replicas=5 -n school-system
  ```

- **View Logs**:

  ```sh
  kubectl logs <pod-name> -n school-system
  ```

  Get pod names:

  ```sh
  kubectl get pods -n school-system
  ```

- **Troubleshoot PostgreSQL**:

  ```sh
  kubectl describe svc postgres-service -n school-system
  ```

### 🛠️ Notes for Cloners

- **Clone**: Ensure `kubernetes_updated/` contains all YAML files.
- **Secrets**: Set up `postgres-secret` and `auth-volume` with correct values.
- **ConfigMap**: Verify `postgres-config.yaml` sets `postgres-url` to `postgres-service`.

- **API Key**: Use `sadat-api-key-1123` for requests.
- **Database**: Create the `students` table manually after PostgreSQL deployment.


---

## 📖 API Endpoints

### Base URL

`http://localhost:8080/api/v1`

### Authentication

Include the API key in the request header:

```
X-API-Key: sadat-api-key-1123
```

### Student Routes

| Method | Endpoint                       | Description                   |
|--------|--------------------------------|-------------------------------|
| GET    | `/api/v1/students`            | Get All Students              |
| POST   | `/api/v1/students`            | Create Student                |
| GET    | `/api/v1/students/{id}`       | Get Student by ID             |
| PUT    | `/api/v1/students/{id}`       | Update Student                |
| PATCH  | `/api/v1/students/{id}`       | Patch Student                 |
| DELETE | `/api/v1/students/{id}`       | Delete Student                |
| POST   | `/api/v1/students/bulk`       | Bulk Insert Students          |

### 🔍 Get All Students
```http
GET api/v1/students
```
**Response:**
```json
[
  {
    "id": 1,
    "name": "Nazmus Sadat Shohag",
    "age": 24,
    "class": 10
  },
    {
    "id": 2,
    "name": "SH Rony",
    "age": 24,
    "class": 10
  }
]
```

### 📌 Get Student by ID
```http
GET api/v1/students/{id}
```

### ➕ Create Student
```http
POST api/v1/students
```
**Request Body:**
```json
{
  "name": "Preity",
  "age": 24,
  "class": 9
}
```

### ✏️ Update Student
```http
PUT api/v1/students/{id}
```
**Request Body:**
```json
{
  "name": "Preety",
  "age": 25,
  "class": 10
}
```


### 🔄 Patch Student
```http
PATCH api/v1/students/{id}
```
**Request Body:**
```json
{
  "age": 26
}
```

### 🗑️ Delete Student
```http
DELETE api/v1/students/{id}
```

### 🔄 Bulk Insert Students
```http
POST /api/v1/students/bulk
```
**Request Body:**
```json
[
  { "name": "Abir", "age": 10, "class": 4 },
  { "name": "Dristy", "age": 9, "class": 3 }
]
```

---
<!-- 
## 📸 Screenshots

![API Example](https://via.placeholder.com/800x400?text=API+Example)

---
-->

## 🔍 Additional Features

### 🗄️ PostgreSQL Database Connection
The API connects to a **PostgreSQL** database for persistent data storage. Connection parameters are managed through environment variables to ensure security and flexibility.

### 🚯 IP-Based Rate Limiting
To prevent abuse, the API enforces **IP-based rate limiting**, restricting excessive requests from the same IP within a specific timeframe.
### 🔐 API Key Authentication
The API uses **API Key Authentication** to protect endpoints. Each request must include a valid API key in the header:

-H "X-API-Key: sadat-api-key-1123"
### 🛑 Graceful Shutdown
The API supports **graceful shutdown**, ensuring proper cleanup of resources when the server is stopped, preventing issues like lingering database connections. When the server gets a shutdown request, it finishes ongoing requests for a specific time, and during that time, it does not take any new requests.

### 📜 Pagination
The API implements **pagination** for efficient handling of large datasets, allowing clients to request data in smaller chunks for better performance.

### ⚡ Bulk Insert with Batching
Bulk insert allows efficiently adding multiple records in a single request, reducing the number of database transactions and improving performance. Also did batching so that query does not exceed.

### 🌐 Docker Hub Integration
Images for the API and web server have been uploaded to Docker Hub, allowing for easy deployment and version management.

### ☸️ Kubernetes Configuration
The project includes Kubernetes configuration files for deploying the application in a cluster, ensuring scalability and manageability.

---

## 📜 License

MIT License. See `LICENSE` for more details.

---

## ⭐ Contributing

Pull requests are welcome! For major changes, please open an issue first.

---

## 🏆 Author

Developed by **Nazmus Sadat Shohag**

🔗 Connect with me on [LinkedIn](https://www.linkedin.com/in/nazmus-sadat-492bba291/)