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

### CLI Application
```
my-cli/
├── cmd/
│   └── my-cli/
│       └── main.go
├── internal/
│   └── app/
│       └── app.go
├── pkg/
│   └── utils/
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

### Web Service
```
my-service/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers.go
│   ├── middleware/
│   │   └── middleware.go
│   └── models/
│       └── models.go
├── pkg/
│   └── utils/
├── configs/
│   └── config.yaml
├── .gitignore
├── go.mod
├── go.sum
└── README.md
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
