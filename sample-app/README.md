# Allscreenshots Demo Application

A sample web application demonstrating the Allscreenshots Go SDK.

## Prerequisites

- Go 1.21 or later
- Allscreenshots API key

## Setup

1. Set your API key:

```bash
export ALLSCREENSHOTS_API_KEY="your-api-key"
```

2. Run the application:

```bash
go run main.go
```

3. Open your browser and navigate to [http://localhost:8080](http://localhost:8080)

## Usage

1. Enter a URL to capture (e.g., `https://github.com`)
2. Select a device preset:
   - Desktop HD (1920x1080)
   - iPhone 14 (390x844)
   - iPad (820x1180)
3. Optionally enable "Full page" to capture the entire scrollable page
4. Click "Take Screenshot"
5. The captured screenshot will appear in the result area

## Configuration

| Environment Variable | Description | Default |
|---------------------|-------------|---------|
| `ALLSCREENSHOTS_API_KEY` | API key for authentication (required) | - |
| `PORT` | HTTP server port | `8080` |

## Project structure

```
sample-app/
├── main.go           # Application entry point and handlers
├── templates/
│   └── index.html    # Main page template
├── static/           # Static assets (if any)
├── go.mod            # Go module file
└── README.md         # This file
```

## License

Apache License 2.0. See [LICENSE](LICENSE) for details.
