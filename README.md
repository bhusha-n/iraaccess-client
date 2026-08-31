# Ira Access Client Guide

This guide explains how to integrate and use the `iraaccess-client` package within your Go application to communicate with the access management server.

---

## Access Model

Every registered app can **check** access freely. Only apps explicitly trusted as management apps (e.g. `admin_console`) can **grant** or **delete** access — including granting access for their own app. If your app calls a grant/delete method and isn't a trusted management app, the server returns `PermissionDenied`.

| Method | Who can call it |
|---|---|
| `CheckAccess`, `CheckAdminAccess` | Any registered app |
| `GrantTenantAccessAs`, `GrantAdminAccessAs`, `DeleteAccessAs`, `DeleteAdminAccessAs` | Only trusted management apps |

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
go get github.com/bhusha-n/iraaccess-client@v0.3.1
go mod tidy
```

---

##  Go Implementation

Here is a complete example showing how to initialize the client, manage its connection lifecycle, and interact with the available access control API endpoints.

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
		log.Fatalf("failed to initialize : %v", err)
	}

	// 2. Ensure connection resources are safely closed when main exits
	defer client.CloseConn()

	log.Println("IraAccess Client securely connected to the server!")


	//  Check methods - any registered app can call these


	isAdmin, err := client.CheckAdminAccess(ctx, "target_user_123")
	if err != nil {
		log.Printf(checkAdminAccess failed: %v", err)
	}
	log.Printf("user Admin status: %t", isAdmin)

	isAllowed, err := client.CheckAccess(ctx, "target_user_123", "tenant_xyz")
	if err != nil {
		log.Printf("checkAccess failed: %v", err)
	}
	log.Printf("user access to tenant allowed: %t", isAllowed)


	//  Grant/Delete methods - management apps only
	//  (e.g. admin_console). Calling these from a
	//  non-management app returns PermissionDenied.

	targetAppID := "campaign_manager" // the app whose scope you're managing

	err = client.GrantAdminAccess(ctx, targetAppID, "admin_user_1", "target_user_123")
	if err != nil {
		log.Printf("grantAdminAccess failed: %v", err)
	}

	err = client.GrantTenantAccess(ctx, targetAppID, "admin_user_1", "target_user_123", "tenant_xyz")
	if err != nil {
		log.Printf("grantTenantAccess failed: %v", err)
	}

	err = client.DeleteAccess(ctx, targetAppID, "target_user_123", "tenant_xyz", "admin_user_1")
	if err != nil {
		log.Printf("deleteAccess failed: %v", err)
	}

	err = client.DeleteAdminAccess(ctx, targetAppID, "target_user_123", "admin_user_1")
	if err != nil {
		log.Printf("deleteAdminAccess failed: %v", err)
	}
}
```
