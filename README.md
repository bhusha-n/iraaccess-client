# Ira Access Client Guide

This guide explains how to integrate and use the `iraaccess-client` package within your Go application to communicate with the access management server.

---

##  Configuration Setup

The library requires a JSON configuration file to authenticate your application identity. You can name this file anything (e.g., `ira-config.json`) and place it anywhere, provided you pass its path to the initializer.

### Configuration Template
Create your JSON file with the following layout:

```json
{
  "app_id": "your-assigned-application-id",
  "app_secret": "your-assigned-application-secret",
  "iam_url": "host-and-port-of-grpc-access-manager"
}
```

*   **app_id**: The unique identifier for your application.
*   **app_secret**: The secure credentials matching your application registration.
*   **iam_url**: The host and port where your centralized gRPC Access Manager is running.

---

##  Installation

Run the following commands in your project terminal to install the client package:

```bash
go get github.com/bhusha-n/iraaccess-client@v0.2.1
go mod tidy
```

---

##  Go Implementation 

Here is a complete, example showing how to initialize the client, manage its connection lifecycle, and interact with the available access control API endpoints.

```go
package main

import (
	"context"
	"log"
	"time"

	ira "github.com/bhusha-n/iraaccess-client"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Initialize the client by passing the path to your JSON config file
	client, err := ira.InitializeIraAccess("ira-config.json")
	if err != nil {
		log.Fatalf("Failed to initialize : %v", err)
	}
	
	// 2. Ensure connection resources are safely closed when main exits
    defer client.CloseConn()

	log.Println("IraAccess Client securely connected to the server!")

	// ==========================================
	//  Available API Usage Methods
	// ==========================================

	// Grant Admin Access
	err = client.GrantAdminAccess(ctx, "admin_user_1", "target_user_123")
	if err != nil {
		log.Printf("GrantAdminAccess failed: %v", err)
	}

	// Check Admin Access
	isAdmin, err := client.CheckAdminAccess(ctx, "target_user_123")
	if err != nil {
		log.Printf("CheckAdminAccess failed: %v", err)
	}
	log.Printf("User Admin status: %t", isAdmin)

	// Grant Tenant Access
	err = client.GrantTenantAccess(ctx, "admin_user_1", "target_user_123", "tenant_xyz")
	if err != nil {
		log.Printf("GrantTenantAccess failed: %v", err)
	}

	// Check Tenant Access
	isAllowed, err := client.CheckAccess(ctx, "target_user_123", "tenant_xyz")
	if err != nil {
		log.Printf("CheckAccess failed: %v", err)
	}
	log.Printf("User access to tenant allowed: %t", isAllowed)

	// Delete Tenant Access
	err = client.DeleteAccess(ctx, "target_user_123", "tenant_xyz", "admin_user_1")
	if err != nil {
		log.Printf("DeleteAccess failed: %v", err)
	}

	// Delete Admin Access
	err = client.DeleteAdminAccess(ctx, "target_user_123", "admin_user_1")
	if err != nil {
		log.Printf("DeleteAdminAccess failed: %v", err)
	}
}
```
