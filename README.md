# MiniUrl - URL Shortener

MiniUrl is a simple and efficient URL shortening service built with Go. It allows users to shorten long URLs into compact aliases, retrieve the original URLs using those aliases, and delete aliases when they are no longer needed.

## Features

- **Shorten URLs**: Generate a short alias for a given URL.
- **Redirect**: Redirect users to the original URL using the alias.
- **Delete URLs**: Remove aliases when they are no longer needed.
- **Validation**: Ensures that the provided URLs are valid.
- **Random Alias Generation**: Automatically generates unique aliases if none are provided.
- **Basic Authentication**: Secures URL creation and deletion endpoints with username and password.
- **Logging**: Provides detailed logs for debugging and monitoring.

## Project Structure

```
miniUrl/
├── cmd/
│   └── url-shortener/   # Main entry point for the application
├── config/              # Configuration files
├── internal/
│   ├── config/          # Configuration loader
│   ├── http-server/     # HTTP server and handlers
│   │   ├── handlers/
│   │   │   ├── url/
│   │   │   │   ├── save/    # Save URL handler
│   │   │   │   ├── delete/  # Delete URL handler
│   │   │   │   └── redirect/# Redirect handler
│   │   └── middleware/   # Middleware for logging, recovery, etc.
│   ├── lib/              # Utility libraries (e.g., logging, random string generation)
│   └── storage/          # Storage layer (SQLite implementation)
├── mocks/                # Mock implementations for testing
├── README.md             # Project documentation
└── go.mod                # Go module dependencies
```

## Getting Started

### Prerequisites

- Go 1.22.5 or higher
- SQLite3

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/miniUrl.git
   cd miniUrl
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the application:
   Edit the `config/local.yaml` file to set up the environment, storage path, and HTTP server credentials.

4. Run the application:
   ```bash
   go run cmd/url-shortener/main.go
   ```

## Usage

### API Endpoints

#### Shorten URL
- **POST** `/url/`
- **Request Body:**
  ```json
  {
    "url": "https://example.com",
    "alias": "customAlias" // Optional
  }
  ```
- **Response:**
  ```json
  {
    "status": "OK",
    "alias": "customAlias"
  }
  ```

#### Redirect
- **GET** `/{alias}`
- Redirects to the original URL.

#### Delete URL
- **DELETE** `/url/{alias}`
- Deletes the alias.

### Authentication

All URL creation (POST) and deletion (DELETE) endpoints are protected using **Basic Authentication**.

- **Username and Password** must be set in the `config/local.yaml` file.
- When making a request, include the Authorization header:
  ```bash
  -u username:password
  ```
  Example using `curl`:
  ```bash
  curl -X POST -u username:password -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/url/
  ```

### Example Usage with curl

**Shorten a URL:**
```bash
curl -X POST -u admin:password -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/url/
```

**Redirect to original URL:**
```bash
curl -v http://localhost:8080/customAlias
```

**Delete an alias:**
```bash
curl -X DELETE -u admin:password http://localhost:8080/url/customAlias
```

## Future Enhancements
- Rate limiting to prevent abuse
- Admin dashboard for managing URLs
- Expiration time for shortened URLs
- Support for custom domains

## License

This project is licensed under the MIT License.

---

Happy shortening! 🚀

