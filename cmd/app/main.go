package main

import (
	"log"
	v1 "my-go-api/api/v1"
	"my-go-api/configs"
	"net/http"
)

func main() {
	// Load configuration
	configs.LoadConfig()

	// Initialize routes
	router := v1.SetupRoutes()

	// Use the port from the configuration
	port := configs.AppConfig.Server.Port
	log.Printf("Starting server on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

/*
├── cmd/
│   └── app/
│       ├── main.go         # Entry point of the application
├── configs/
│   ├── config.yaml         # Configuration files (YAML, JSON, TOML, etc.)
│   └── config.go           # Config loading logic
├── internal/
│   ├── handlers/           # HTTP handlers (controllers)
│   │   ├── user_handler.go
│   ├── services/           # Business logic (services)
│   │   ├── user_service.go
│   ├── repositories/       # Data access logic (repositories)
│   │   ├── user_repo.go
│   ├── models/             # Data models (structs)
│   │   ├── user.go
│   ├── middleware/         # Middleware (e.g., authentication)
│   │   ├── auth_middleware.go
│   └── utils/              # Utility functions (helpers, etc.)
│       ├── hash.go
│       └── jwt.go
├── pkg/                    # Public packages for reuse across projects
│   └── logger/             # Example: Logging utility
│       └── logger.go
├── api/                    # API versioning
│   └── v1/
│       ├── routes.go       # Route definitions (e.g., HTTP endpoints)
├── db/                     # Database setup (migrations, seeds)
│   ├── migrations/
│   └── seed.go
├── scripts/                # Helper scripts (e.g., for local setup)
├── .env                    # Environment variables
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksum file
├── Dockerfile              # Dockerfile for containerization
├── Makefile                # Automate common tasks (build, run, etc.)
└── README.md               # Project documentation
*/
