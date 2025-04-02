# go-project-generator

## Overview

Go Project Generator is a powerful CLI tool designed to streamline the process of creating new Go projects. It provides quick, consistent project scaffolding for various project types, helping developers focus on writing code instead of setting up project structures.

## Features

### Supported Project Types
- CLI Application
- Web Service
- Microservice
- Library
- Simple Command Tool

### Key Capabilities
- Automated project directory creation
- Standardized project structure
- Customizable project templates
- Dependency management
- Git repository initialization (optional)

## Installation

### Prerequisites
- Go 1.21 or higher
- Git (recommended)

### Install from Source
```bash
git clone https://github.com/hassani-jr/go-project-generator.git
cd go-project-generator
go install
```

### Install via Go
```bash
go install github.com/hassani-jr/go-project-generator@latest
```

## Usage

### Basic Command
```bash
go-project-generator create [project-type] [project-name]
```

### Examples
```bash
# Create a CLI application
go-project-generator create cli my-awesome-cli

# Create a web service
go-project-generator create webservice user-api

# Create a library
go-project-generator create lib go-validator
```

### Options
```
--template, -t     Specify a custom template
--output, -o       Specify output directory (default: current directory)
--modules, -m      List of additional modules to include
--git, -g          Initialize git repository (default: true)
--verbose, -v      Enable verbose output
--help, -h         Show help information
```

## Project Structures

Each project type creates a standardized directory structure following Go best practices.

## CLI Application Structure

## Directories and Files
```
project-name/
├── cmd/
│   └── root.go        # Main CLI command structure
├── internal/
│   └── commands/      # Specific command implementations
├── pkg/               # Shared packages if needed
├── configs/           # Configuration files
├── scripts/           # Utility scripts
├── go.mod
├── main.go            # Entry point
└── README.md
```
## Directory Descriptions

### `cmd/`
- Contains the main command-line interface logic
- `root.go`: Defines the primary command structure

### `internal/`
- Houses private application code
- `commands/`: Implements specific CLI commands

### `pkg/`
- Stores packages that can be imported by external projects
- Shared utility functions or types

### `configs/`
- Configuration files for the application
- Environment-specific settings

### `scripts/`
- Utility scripts for development, deployment, etc.

## Getting Started

### Running the Application
```bash
go run main.go
```
## Web Service Project Structure

## Directories and Files
```
project-name/
├── cmd/
│   └── server/        # Server startup and configuration
├── internal/
│   ├── handlers/      # HTTP route handlers
│   ├── middleware/    # Request middleware
│   ├── models/        # Data models
│   └── services/      # Business logic
├── pkg/
│   ├── database/      # Database connection helpers
│   └── config/        # Configuration helpers
├── configs/
│   ├── config.yaml    # Application configuration
│   └── database.yaml  # Database configuration
├── scripts/
│   ├── migrations/    # Database migration scripts
│   └── init.sh        # Startup scripts
├── api/               # OpenAPI/Swagger specs
├── go.mod
├── main.go
└── README.md
```

## Directory Descriptions

### `cmd/server/`
- Server startup and configuration logic
- Main entry point for the web service

### `internal/handlers/`
- HTTP route handlers
- Define API endpoint logic

### `internal/middleware/`
- Request processing middleware
- Authentication, logging, etc.

### `internal/models/`
- Data structures and models
- Database entity representations

### `internal/services/`
- Business logic implementation
- Core application functionality

### `pkg/database/`
- Database connection and interaction helpers
- Connection pooling, query helpers

### `configs/`
- Configuration files
- Environment-specific settings

### `scripts/`
- Development and deployment utility scripts
- Database migrations
- Initialization scripts

### `api/`
- API specification files
- OpenAPI/Swagger documentation

## Getting Started

### Configuration
1. Copy `configs/config.yaml.example` to `configs/config.yaml`
2. Modify configuration as needed

### Running the Service
```bash
go run main.go
```

## Microservice Project Structure

## Directories and Files
```
project-name/
├── cmd/
│   └── server/        # gRPC server startup
├── internal/
│   ├── service/       # Service implementation
│   └── proto/         # Protocol buffer definitions
├── pkg/
│   ├── interceptors/  # gRPC interceptors
│   └── client/        # Service client implementations
├── configs/
│   ├── config.yaml    # Service configuration
│   └── grpc.yaml      # gRPC specific configs
├── scripts/
│   └── proto-gen.sh   # Protobuf generation script
├── api/               # Service definition files
├── go.mod
├── main.go
└── README.md
```

## Directory Descriptions

### `cmd/server/`
- gRPC server startup and configuration
- Main entry point for the microservice

### `internal/service/`
- Core service implementation
- Business logic for the microservice

### `internal/proto/`
- Protocol buffer definitions
- Service contract and data models

### `pkg/interceptors/`
- gRPC request interceptors
- Logging, authentication, etc.

### `pkg/client/`
- Client implementations
- Service-to-service communication helpers

### `configs/`
- Configuration files
- Service and gRPC-specific settings

### `scripts/`
- Utility scripts
- Protobuf code generation

### `api/`
- Service definition files
- OpenAPI/gRPC service specifications

## Getting Started

### Generate Protobuf
```bash
./scripts/proto-gen.sh
```

### Running the Service
```bash
go run main.go
```
### Building
```bash
go build
```

## Go Library Project Structure

## Directories and Files
```
project-name/
├── pkg/
│   └── yourlib/       # Main library code
├── examples/          # Usage examples
├── internal/          # Private implementation details
│   └── helpers/       # Internal helpers
├── scripts/
│   └── generate.sh    # Code generation scripts
├── docs/              # Documentation
├── go.mod
├── README.md
└── LICENSE
```

## Directory Descriptions

### `pkg/yourlib/`
- Main library code
- Public API and core functionality

### `examples/`
- Example usage of the library
- Demonstrates library features

### `internal/`
- Private implementation details
- Not accessible to external users

### `internal/helpers/`
- Internal utility functions
- Supporting library implementation

### `scripts/`
- Utility scripts
- Code generation, build helpers

### `docs/`
- Detailed documentation
- API reference, usage guides

## Getting Started

### Installation
```bash
go get github.com/yourusername/yourlib
```

### Usage Example
```
import "github.com/yourusername/yourlib"

func main() {
    // Use library functionality
}
```

### Running Tests
```bash
go test ./...
```

## Command Tool Project Structure

## Directories and Files

```
project-name/
├── cmd/
│   └── commands/      # Individual command implementations
├── internal/
│   └── utils/         # Utility functions
├── pkg/               # Shared packages
├── configs/           # Configuration files
├── scripts/           # Utility scripts
├── go.mod
├── main.go
└── README.md
```
## Directory Descriptions

### `cmd/commands/`
- Implementations of individual commands
- Specific tool functionalities

### `internal/utils/`
- Utility functions
- Shared helper methods

### `pkg/`
- Packages that can be imported
- Shared functionality

### `configs/`
- Configuration files
- Tool-specific settings

### `scripts/`
- Utility and helper scripts
- Development and deployment aids

## Getting Started

### Running the Tool
```bash
go run main.go [command]
```

### Building
```bash
go build
```

### Installation
```bash
go install
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- The Go team for their amazing language
- All contributors who have helped shape this project
