# dating-app
Technical Test From Dealls - Create Basic Dating App

The structure Service is DDD in Go software development typically consists of several key concepts:

Domain Layer: The domain layer is the core of the application, containing business logic and domain objects. Domain objects represent important concepts in the business domain, such as entities, value objects, and aggregates. Complex business logic is implemented in this layer.

Infrastructure Layer: The infrastructure layer contains code responsible for interacting with external resources, such as databases, web services, or file systems. This layer provides concrete implementations of contracts defined in the domain layer, such as repositories, external services, and storage mechanisms.

Service Layer: The application layer (also known as the service layer or application service layer) connects the domain layer with the infrastructure layer. This layer contains high-level application logic, such as use cases and application workflows. Its purpose is to orchestrate interactions between domain objects and infrastructure, and to provide an outward-facing interface for use by users or other systems.

Interface Layer: The interface layer contains code related to user interaction, whether through graphical user interfaces, command-line interfaces, or API interfaces. This layer is responsible for presenting data to users and forwarding user requests to the application layer.

## Requirement

Before running the service, make sure the points below are completed:

- [Go](https://golang.org/) (Golang) - Versi 1.18 atau lebih 
- [Postgresql](https://www.postgresql.org/)
- [Redis](https://redis.com/)

## Started

1. clone repository:

    ```bash
    git clone https://github.com/ShadamHarizky/dating-app
    cd dating-app
    ```

2. install dependecies:

    ```bash
    go mod init github.com/ShadamHarizky/dating-app
    go mod tidy
    ```

3. Create .env according to .env-example:

    - Change The Connection Detail Redis on .env file
    - Makesure Postgresql And Rebis completed install on your local computer.

4. Running Project:

    ```bash
    go run cmd/main.go
    ```

5. Testing Using Postman

    - File Already upload in this repository

## Configuration

Update the env file `.env` and adjust it to the configuration on each computer
