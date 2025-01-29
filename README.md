# Task Manager App Backend

## Description
The Task Manager App Backend is a robust and scalable backend service built with GoLang. It provides a comprehensive API for managing tasks and users, leveraging GraphQL for flexible and efficient data querying. The backend is designed with a focus on security, performance, and maintainability, utilizing JWT for user authentication, PostgreSQL for data storage, and Docker and Kubernetes for containerization and orchestration. Continuous Integration and Continuous Deployment (CI/CD) are managed with Jenkins.

## Technologies Used
- **GoLang**: Backend development
- **GraphQL**: API development
- **PostgreSQL**: Database
- **JWT**: User authentication
- **Docker**: Containerization
- **Kubernetes**: Orchestration and deployment
- **Jenkins**: CI/CD
- **Domain-Driven Design (DDD)**: Code organization

## Architecture
The backend follows the principles of Domain-Driven Design (DDD), structured into distinct layers:
- **Domain**: Contains the core business logic and domain entities.
- **Application**: Handles the application logic and orchestrates the use cases.
- **Infrastructure**: Manages the interaction with external systems like databases and third-party services.

## Features
- **GraphQL Endpoints**: Create, read, update, and delete tasks.
- **User Authentication**: Secure user authentication using JWT.
- **Database Integration**: Seamless integration with PostgreSQL for data persistence.

## Installation Instructions
1. **Clone the repository**:
   ```sh
   git clone https://github.com/your-repo/task-manager-app.git
   cd task-manager-app/backend
   ```
2. **Install dependencies**:
   ```sh
   go mod download
   ```
3. **Configure PostgreSQL**:
   - Update the database configuration in `config/config.go`.
4. **Run the application**:
   ```sh
   go run cmd/main.go
   ```
5. **Run with Docker**:
   ```sh
   docker-compose up --build
   ```

## Deployment Instructions
1. **Kubernetes Setup**:
   - Ensure you have a Kubernetes cluster running.
   - Apply the Kubernetes manifests:
     ```sh
     kubectl apply -f k8s/
     ```
2. **Kubernetes Manifests**:
   - `backend-deployment.yaml`: Deployment configuration for the backend.
   - `backend-service.yaml`: Service configuration for the backend.
   - `postgres-deployment.yaml`: Deployment configuration for PostgreSQL.
   - `postgres-service.yaml`: Service configuration for PostgreSQL.

## CI/CD with Jenkins
1. **Jenkins Setup**:
   - Install Jenkins and required plugins (Docker, Kubernetes, Git, etc.).
   - Configure Jenkins credentials for Docker and Kubernetes.
2. **Jenkins Pipeline**:
   - Example Jenkinsfile for CI/CD:
     ```groovy
     pipeline {
         agent any
         stages {
             stage('Build') {
                 steps {
                     sh 'docker build -t your-repo/task-manager-app:latest .'
                 }
             }
             stage('Test') {
                 steps {
                     sh 'go test ./...'
                 }
             }
             stage('Deploy') {
                 steps {
                     sh 'kubectl apply -f k8s/'
                 }
             }
         }
     }
     ```

## Running Tests
- **Unit Tests**:
  ```sh
  go test ./internal/tests/unit/...
  ```
- **Integration Tests**:
  ```sh
  go test ./internal/tests/integration/...
  ```

## Contributing
We welcome contributions from the community. To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and push the branch to your fork.
4. Open a pull request with a detailed description of your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
