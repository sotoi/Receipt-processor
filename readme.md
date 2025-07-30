# Receipt Processing Service Take-Home

## Overview

Go-based receipt processing service designed to validate receipts, calculate reward points based on predefined rules, and manage receipts in storage. The service exposes HTTP endpoints for saving receipts and retrieving calculated points.

## Getting Started

### Prerequisites

- **Go:** Ensure you have Go installed (version 1.23 or later)

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/fetch.git
   cd fetch
   ```

2. **Install Dependencies:**
   ```bash
   go mod download
   ```

### Running the Application

Start the server using the following command:

```bash
go run main.go
```

or

```bash
make run
```

The server will start and listen on the configured port (default is `8080`).

### API Endpoints

#### 1. Save Receipt

- **Endpoint:** `/receipts/process`
- **Method:** `POST`
- **Description:** Validates and saves a receipt.
- **Request Body:**

  ```json
  {
      "retailer": "Walgreens",
      "purchaseDate": "2022-01-02",
      "purchaseTime": "08:13",
      "total": "2.65",
      "items": [
          {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
          {"shortDescription": "Dasani", "price": "1.40"}
      ]
  }
  ```

- **Response:**
  - **Success (200):**

    ```json
    {
      "id": "generated-receipt-uuid"
    }
    ```

  - **Failure (400):**
    - Invalid method or payload.

#### 2. Get Points

- **Endpoint:** `/receipts/{id}/points`
- **Method:** `GET`
- **Description:** Retrieves the calculated points for a specific receipt.
- **Path Parameter:** id(string) - The unique identifier of the receipt.
- **Response:**
  - **Success (200):**

    ```json
    {
      "points": 28
    }
    ```

  - **Failure (404):**
    - Receipt not found.

  - **Failure (400):**
    - Invalid receipt ID format.
