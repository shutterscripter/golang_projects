# URL Shortener

A high-performance URL shortener service built with Go, Fiber, and Redis. This service provides a RESTful API to shorten URLs with custom short codes, expiration times, and built-in rate limiting.

## Features

- 🔗 **URL Shortening**: Convert long URLs into short, shareable links
- 🎯 **Custom Short URLs**: Support for custom short codes
- ⏱️ **Expiration**: Set custom expiration times for shortened URLs (default: 24 hours)
- 🚦 **Rate Limiting**: IP-based rate limiting to prevent abuse
- ✅ **URL Validation**: Validates URLs before shortening
- 🔒 **SSL Enforcement**: Automatically enforces HTTP/HTTPS protocol
- 📊 **Analytics**: Tracks redirect counts
- 🐳 **Docker Ready**: Fully containerized with Docker Compose

## Tech Stack

- **Language**: Go 1.25.6
- **Web Framework**: [Fiber v3](https://github.com/gofiber/fiber)
- **Database**: Redis
- **Containerization**: Docker & Docker Compose
- **Dependencies**:
  - `gofiber/fiber/v3` - Web framework
  - `redis/go-redis/v9` - Redis client
  - `google/uuid` - UUID generation
  - `asaskevich/govalidator` - URL validation
  - `joho/godotenv` - Environment variable management

## Project Structure

```
url_shortner/
├── api/
│   ├── main.go              # Application entry point
│   ├── go.mod               # Go module dependencies
│   ├── Dockerfile           # API container image
│   ├── database/
│   │   └── database.go      # Redis client setup
│   ├── helpers/
│   │   └── helpers.go       # Utility functions
│   └── routes/
│       ├── shorten.go       # URL shortening endpoint
│       └── resolve.go       # URL resolution endpoint
├── db/
│   └── Dockerfile           # Redis container image
├── docker-compose.yml       # Docker Compose configuration
└── README.md
```

## Environment Variables

Create a `.env` file in the `api/` directory with the following variables:

```env
API_PORT=:3000
DOMAIN=localhost:3000
DB_ADDR=db:6379
DB_PASS=
API_QUOTA=10
```

- `API_PORT`: Port on which the API server runs
- `DOMAIN`: Domain name for generating short URLs
- `DB_ADDR`: Redis server address
- `DB_PASS`: Redis password (leave empty if no password)
- `API_QUOTA`: Number of requests allowed per IP per 30 minutes

## API Endpoints

### 1. Shorten URL

**POST** `/api/v1`

Shortens a given URL with optional custom short code and expiration time.

**Request Body:**

```json
{
  "url": "https://www.example.com/very/long/url",
  "short": "mycustom", // Optional: custom short code
  "expiry": 48 // Optional: expiration in hours (default: 24)
}
```

**Response:**

```json
{
  "url": "https://www.example.com/very/long/url",
  "short": "localhost:3000/mycustom",
  "expiry": 48,
  "rate_limit": 9,
  "rate_limit_reset": 30
}
```

**Error Responses:**

- `400 Bad Request`: Invalid JSON or invalid URL
- `403 Forbidden`: Custom short code already in use
- `503 Service Unavailable`: Rate limit exceeded or domain error

### 2. Resolve URL

**GET** `/:url`

Redirects to the original URL associated with the short code.

**Example:**

```
GET /mycustom
→ Redirects to https://www.example.com/very/long/url
```

**Response:**

- `301 Moved Permanently`: Successful redirect
- `404 Not Found`: Short URL not found
- `500 Internal Server Error`: Database connection error

## Installation & Setup

### Using Docker Compose (Recommended)

1. Clone the repository:

```bash
git clone <repository-url>
cd url_shortner
```

2. Create a `.env` file in the `api/` directory with required environment variables.

3. Build and run with Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:3000`

### Manual Setup

1. Install Go 1.25.6 or higher
2. Install Redis
3. Navigate to the `api/` directory:

```bash
cd api
```

4. Install dependencies:

```bash
go mod download
```

5. Create a `.env` file with required environment variables

6. Run the application:

```bash
go run main.go
```

## Usage Examples

### Shorten a URL

```bash
curl -X POST http://localhost:3000/api/v1 \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://github.com/example/repository",
    "expiry": 72
  }'
```

### Shorten with Custom Code

```bash
curl -X POST http://localhost:3000/api/v1 \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://github.com/example/repository",
    "short": "github-repo",
    "expiry": 168
  }'
```

### Access Shortened URL

```bash
curl http://localhost:3000/github-repo
# or visit in browser
```

## Rate Limiting

- Default quota: 10 requests per IP per 30 minutes
- Configurable via `API_QUOTA` environment variable
- Rate limit info included in API responses
- Automatic reset after time window expires

## Redis Database Structure

The service uses two Redis databases:

- **DB 0**: Stores URL mappings (short code → original URL)
- **DB 1**: Stores rate limiting data (IP → remaining quota) and analytics counters

## Docker Configuration

### Services

- **api**: Go application (port 3000)
- **db**: Redis database (port 6379)

### Volumes

- `.data:/data` - Persists Redis data

## Features in Detail

### URL Validation

- Validates URL format using govalidator
- Prevents shortening of the service's own domain
- Enforces HTTP/HTTPS protocol

### Custom Short Codes

- Allows users to specify custom short codes
- Checks for uniqueness before creation
- Auto-generates 5-character UUID if not provided

### Expiration

- Default expiration: 24 hours
- Customizable per URL
- Automatic cleanup by Redis TTL

### Analytics

- Tracks total redirect counter
- Can be extended for per-URL analytics

## License

This project is open source and available under the MIT License.

## Author

Built by [shutterscripter](https://github.com/shutterscripter)
