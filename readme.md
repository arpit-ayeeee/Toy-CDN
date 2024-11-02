#  Toy CDN :

This Go project is a simple CDN service with caching capabilities. It forwards client requests to an origin server, caches the responses locally, and serves cached responses on subsequent requests. This reduces load on the origin server and improves response times for frequently accessed resources.

## Features

- **Proxy Server**: Forwards HTTP requests from clients to a specified origin server.
- **Caching**: Stores responses in a local cache directory to improve performance on repeated requests.
- **Dynamic Origin Mapping**: Configurable mapping for host-based origin redirection.
- **Cache Expiry**: No expiry is implemented in this version; entries are cached indefinitely.

## Setup

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or newer)
- MySQL Database (optional, if needed for custom database usage)

### Configuration

1. **Origin Mapping**: The `originMap` variable in `main.go` maps local hosts to specific origin servers. Adjust as needed.

    ```go
    var originMap = map[string]string{
        "localhost:9000": "http://google.com",
    }
    ```

2. **Cache Directory**: This project uses a local `./cache` directory to store cached responses. The directory is created automatically on startup if it does not exist.

## Running the Server

1. Clone the repository and navigate to the project directory:

    ```bash
    git clone <repository-url>
    cd <project-directory>
    ```

2. Build and run the Go application:

    ```bash
    go run main.go
    ```

3. The server listens on port 9000 by default. You can access it by navigating to `http://localhost:9000`.

## Code Structure

- **`main.go`**: Contains the main proxy server code with caching functionality.
  - **Functions**:
    - `proxy`: Handles incoming requests, checks for cached responses, and proxies requests to the origin if needed.
    - `cacheGet`: Reads from the cache if the response is available.
    - `cachePut`: Writes the response to the cache.

## Usage

Send HTTP requests to `http://localhost:9000` or any mapped host. For example:

## Sample Response
Fetching from origin: http://google.com/?
Cache hit: f22be6712ce1385865805856627b5ccf
Cache hit: f22be6712ce1385865805856627b5ccf
Fetching from origin: http://google.com/helloworld?
Fetching from origin: http://google.com/helloworld?
Fetching from origin: http://google.com/helloworld?
Fetching from origin: http://google.com/search?q=tes%5C&sca_esv=6bf062aa074a1fec&sxsrf=ADLYWIJrKSArIR8n3aReDsAHkEo6sO2PIg%3A1730565312892&source=hp&ei=wFQmZ_uVNK6cseMP89KNiQ8&iflsig=AL9hbdgAAAAAZyZi0MbZbmk8O1Ci7bTxWc37ouvfRS91&ved=0ahUKEwj79e6Vir6JAxUuTmwGHXNpI_EQ4dUDCBg&uact=5&oq=tes%5C&gs_lp=Egdnd3Mtd2l6IgR0ZXNcMg0QABiABBixAxhDGIoFMgoQABiABBhDGIoFMgUQABiABDIFEAAYgAQyBRAAGIAEMgUQABiABDIFEAAYgAQyBRAAGIAEMgUQABiABDIFEAAYgARImQ9QAFjUCnABeACQAQCYAWOgAcIDqgEBNbgBA8gBAPgBAZgCBqAC3wOoAgrCAgoQIxiABBgnGIoFwgIREC4YgAQYsQMY0QMYgwEYxwHCAgsQABiABBixAxiDAcICCBAuGIAEGLEDwgIFEC4YgATCAggQABiABBixA8ICBxAjGCcY6gLCAhAQLhiABBjRAxhDGMcBGIoFwgIQEAAYgAQYsQMYQxiDARiKBcICFhAuGIAEGLEDGNEDGEMYxwEYyQMYigXCAhMQLhiABBixAxjRAxhDGMcBGIoFwgILEAAYgAQYkgMYigWYAwOSBwM1LjGgB7ou&sclient=gws-wiz
Cache hit: d65b27acb73d5ec7ef902f5bee88f9b0

```bash
curl http://localhost:9000/path

