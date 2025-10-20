# String Analyzer Service

This is a RESTful API service that analyzes strings and stores their computed properties.

## Features

- Analyzes strings for various properties:
  - Length
  - Palindrome status
  - Unique character count
  - Word count
  - SHA-256 hash
  - Character frequency map
- Stores and retrieves string analyses.
- Filters strings based on their properties.
- Supports natural language queries for filtering.

## API Endpoints

### Create/Analyze String

- **URL:** `/strings`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "value": "string to analyze"
  }
  ```
- **Success Response (201 Created):**
  ```json
  {
    "id": "sha256_hash_value",
    "value": "string to analyze",
    "properties": {
      "length": 16,
      "is_palindrome": false,
      "unique_characters": 12,
      "word_count": 3,
      "sha256_hash": "abc123...",
      "character_frequency_map": {
        "s": 2,
        "t": 3,
        "r": 2
      }
    },
    "created_at": "2025-08-27T10:00:00Z"
  }
  ```

### Get Specific String

- **URL:** `/strings/{string_value}`
- **Method:** `GET`
- **Success Response (200 OK):**
  ```json
  {
    "id": "sha256_hash_value",
    "value": "requested string",
    "properties": { /* same as above */ },
    "created_at": "2025-08-27T10:00:00Z"
  }
  ```

### Get All Strings with Filtering

- **URL:** `/strings`
- **Method:** `GET`
- **Query Parameters:**
  - `is_palindrome` (boolean)
  - `min_length` (integer)
  - `max_length` (integer)
  - `word_count` (integer)
  - `contains_character` (string)
- **Success Response (200 OK):**
  ```json
  {
    "data": [
      {
        "id": "hash1",
        "value": "string1",
        "properties": { /* ... */ },
        "created_at": "2025-08-27T10:00:00Z"
      }
    ],
    "count": 1,
    "filters_applied": {
      "is_palindrome": true
    }
  }
  ```

### Natural Language Filtering

- **URL:** `/strings/filter-by-natural-language`
- **Method:** `GET`
- **Query Parameters:**
  - `query` (string)
- **Success Response (200 OK):**
  ```json
  {
    "data": [ /* array of matching strings */ ],
    "count": 1,
    "interpreted_query": {
      "original": "all single word palindromic strings",
      "parsed_filters": {
        "word_count": "1",
        "is_palindrome": "true"
      }
    }
  }
  ```

### Delete String

- **URL:** `/strings/{string_value}`
- **Method:** `DELETE`
- **Success Response (204 No Content):** (Empty response body)

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.24.0 or higher)
- [Docker](https://www.docker.com/get-started) (optional)

### Dependencies

The service uses the following Go modules:

- `github.com/gin-gonic/gin`

All other dependencies are indirect and will be downloaded automatically.

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/CyberwizD/String-Analyzer-Service.git
   ```
2. Navigate to the project directory:
   ```sh
   cd String-Analyzer-Service
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application

#### Locally

To run the service locally, execute the following command:

```sh
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

#### With Docker

To run the service using Docker, first build the Docker image:

```sh
docker build -t string-analyzer-service .
```

Then, run the Docker container:

```sh
docker run -p 8080:8080 string-analyzer-service
```

The server will be accessible at `http://localhost:8080`.

### Environment Variables

This service does not require any environment variables to be set.
