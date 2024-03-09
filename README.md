# Enigma Laundry Submission Project

This is a submission project for Enigma Laundry.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project is a Go application that manages customers, products, and transactions for Enigma Laundry. It uses the Gin web framework and PostgreSQL as the database.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed on your machine
- PostgreSQL database server
- Git (optional)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/submission-project-enigma-laundry.git
   ```

2. Change into the project directory:

   ```bash
   cd submission-project-enigma-laundry
   ```

3. Install dependencies:

   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/lib/pq
   ```

4. Set up the database:

   - Create a PostgreSQL database.
   - Update the database connection details in the `config/config.go` file.

## Usage

To run the application, use the following command:

```bash
go run main.go
```

The application will start, and you can access it at [http://localhost:8080](http://localhost:8080).

## API Endpoints

- **Create Customer:**
  ```
  POST /customers
  ```

- **Get Customer:**
  ```
  GET /customers/:id
  ```

- **Update Customer:**
  ```
  PUT /customers/:id
  ```

- **Delete Customer:**
  ```
  DELETE /customers/:id
  ```

- **Create Product:**
  ```
  POST /products
  ```

- **List Products:**
  ```
  GET /products
  ```

- **Get Product:**
  ```
  GET /products/:id
  ```

- **Update Product:**
  ```
  PUT /products/:id
  ```

- **Delete Product:**
  ```
  DELETE /products/:id
  ```

- **Create Transaction:**
  ```
  POST /transactions
  ```

- **Get Transaction:**
  ```
  GET /transactions/:id_bill
  ```

- **List Transactions:**
  ```
  GET /transactions
  ```

## Contributing

To contribute to this project, follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add your feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Replace the placeholder URLs, usernames, and other details with your actual project information. This template provides a basic structure, and you can expand or modify it based on your project's specific requirements.