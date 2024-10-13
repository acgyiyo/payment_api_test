# API Payment Application

## Overview
This Go application provides a simple payment processing system that allows users to register new payments, retrieve existing ones, and process refunds. 

## Getting Started

### Prerequisites
* Go programming language (version 1.18 or later)
* A Go module compatible environment
* Docker

### Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/acgyiyo/payment_api_test.git
   ```
2. 
```bash
   cd project-directory
   go mod tidy
   go run cmd/main.go 
   ```

   By default, the application will listen on port 8082. You can customize this by setting the PORT environment variable.

### Available Endpoints

Register a payment:
Endpoint: /payments
Method: POST
Request body: A JSON object containing payment details (e.g., amount, currency, customer information)
Retrieve a payment:
Endpoint: /payments/{paymentId}
Method: GET
Path parameter: paymentId is the unique identifier of the payment
Refund a payment:
Endpoint: /payments/refund
Method: POST

Example

```bash
curl --location 'localhost:8082/payment/refund' \
--header 'Content-Type: application/json' \
--data '{
    "transaction_id": "txn-20241012204528+7"
}'
```

### Additional Notes
Database: The application uses a database to store payment information. You'll need to configure the database connection details in the configuration file and run doccker-compose file
```
internal\config\schemas\local.json

```