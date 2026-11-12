# IraAccess Client

**Centralized, multi-tenant authorization for Go applications.**

`iraaccess-client` is a Go client library for integrating applications with the IraAccess centralized Access Management Service through gRPC.

It provides a simple API for managing application roles, tenant-level administration, resource permissions, and authorization checks.

## What's New in v1.0.0

Version 1.0.0 is a major architectural update that introduces a new application-scoped, role-based and resource-level authorization model.

### Major Changes

| Feature | Description |
|---|---|
| Application Isolation | Authorization is scoped to individual applications. |
| Super Admin | Application-wide administration. |
| Tenant Admin | Administration restricted to an assigned tenant. |
| Tenant Member | Access to explicitly permitted resources. |
| Resource Permissions | Permissions at the resource-type and resource-ID levels. |
| Resource Scopes | Support for ALL and SELECTED resource scopes. |
| Action-Based Access | VIEW, CREATE, UPDATE and DELETE permissions. |
| Role Management | Assign and remove application and tenant roles. |
| Access Management | Grant and revoke resource permissions. |
| Access Checking | Check whether a user can perform an action on a resource. |
| Access Tree | Retrieve a user's complete access structure. |
| Multiple Configurations | Initialize separate clients for different applications. |

### Breaking Changes from v0.3.1

The previous API has been replaced with a new API.

| Previous API | v1.0.0 API |
|---|---|
| `GrantAdminAccess` | `AssignSuperAdmin` |
| `GrantTenantAccess` | `AssignTenantAdmin` |
| `DeleteAdminAccess` | `RemoveSuperAdmin` |
| `DeleteAccess` | `RemoveTenantAdmin` / `RemoveTenantMember` |
| `CheckAdminAccess` | Use `GetUserAccessTree` or role information from the access tree |
| `CheckAccess` (old signature) | `CheckAccess` (resource-level signature) |
| `CloseConn` | `Close` |

The old API is not compatible with v1.0.0. Update your integration before upgrading.

---

## 1. Access Control Model

IraAccess uses three primary roles.

### Roles

**Super Admin**

- Has application-wide administrative authority.
- Can manage access across tenants within the application.
- Can assign and remove administrative roles.

**Tenant Admin**

- Has administrative authority within an assigned tenant.
- Can manage members and resource permissions within their tenant.
- Cannot administer other tenants or assign application-wide Super Admin privileges.

**Tenant Member**

- Receives resource-level permissions.
- Can access resources according to assigned scopes and actions.
- Does not have administrative privileges.

### Authorization Hierarchy

```text
Application
│
├── Super Admin
│   └── Application-wide administration
│
├── Tenant A
│   ├── Tenant Admin
│   │   └── Tenant-level administration
│   │
│   └── Tenant Members
│       ├── Resource Type: Dialplan
│       │   ├── ALL
│       │   └── SELECTED
│       │       ├── dialplan-001
│       │       └── dialplan-002
│       │
│       └── Resource Type: Extension
│           └── SELECTED
│               ├── extension-001
│               └── extension-002
│
└── Tenant B
    ├── Tenant Admin
    └── Tenant Members
```

### Resource Actions

Permissions are assigned using the following actions:

| Action | Description |
|---|---|
| `ActionView` | View a resource. |
| `ActionCreate` | Create resources. |
| `ActionUpdate` | Update resources. |
| `ActionDelete` | Delete resources. |

### Resource Scopes

Two resource scopes are supported.

**ScopeAll**

Grants the specified actions across all resources of a particular resource type within the tenant.

**ScopeSelected**

Grants the specified actions only on the explicitly supplied resource IDs.

For example:

```go
client.AssignTenantMember(
    ctx,
    "web_shub",
    "tenant-001",
    "user-001",
    "dialplan",
    ira.ScopeSelected,
    []string{"dialplan-001", "dialplan-002"},
    []ira.Action{
        ira.ActionView,
        ira.ActionUpdate,
    },
    "admin-001",
)
```

The member receives VIEW and UPDATE permissions on the two specified dialplans.

---

## 2. Installation

Install the latest major version:

```bash
go get github.com/bhusha-n/iraaccess-client@v1.0.0
go mod tidy
```

Import the package:

```go
import ira "github.com/bhusha-n/iraaccess-client"
```

---

## 3. Configuration Setup

Each registered application must have its own credentials and configuration.

Create a separate JSON configuration file for each application.

### Configuration Template

```json
{
  "app_id": "your-application-id",
  "app_secret": "your-application-secret",
  "iam_url": "localhost:50051"
}
```

| Field | Description |
|---|---|
| `app_id` | Unique registered application identifier. |
| `app_secret` | Application authentication secret. |
| `iam_url` | Host and port of the IraAccess gRPC server. |


## 4. Initialize the Client

Initialize the client by passing the configuration file path.

```go
package main

import (
    "context"
    "log"
    "time"

    ira "github.com/bhusha-n/iraaccess-client"
)

func main() {
    client, err := ira.InitializeIraAccess("config/web-shub.json")
    if err != nil {
        log.Fatalf("failed to initialize IraAccess: %v", err)
    }
    defer client.Close()

    ctx, cancel := context.WithTimeout(
        context.Background(),
        10*time.Second,
    )
    defer cancel()

    log.Println("IraAccess client initialized")

    _ = ctx
}
```

The client automatically attaches the configured application credentials to outgoing gRPC requests.

Always close the client when your application shuts down.

---

## 5. Role Management

### Assign Super Admin

Assigns application-wide administrative privileges to a user.

```go
err := client.AssignSuperAdmin(
    ctx,
    "web_shub",
    "user-001",
    "system-admin",
)
if err != nil {
    log.Printf("failed to assign super admin: %v", err)
}
```

Parameters:

| Parameter | Description |
|---|---|
| `ctx` | Request context. |
| `targetAppID` | Application receiving the role assignment. |
| `userID` | User receiving the role. |
| `grantedBy` | Identity of the actor granting the role. |

### Assign Tenant Admin

Assigns administrative privileges within a specific tenant.

```go
err := client.AssignTenantAdmin(
    ctx,
    "web_shub",
    "tenant-001",
    "user-002",
    "super-admin",
)
if err != nil {
    log.Printf("failed to assign tenant admin: %v", err)
}
```

### Assign Tenant Member

Assigns a tenant member with specific resource permissions.

```go
err := client.AssignTenantMember(
    ctx,
    "web_shub",
    "tenant-001",
    "user-003",
    "dialplan",
    ira.ScopeAll,
    nil,
    []ira.Action{
        ira.ActionView,
        ira.ActionCreate,
        ira.ActionUpdate,
    },
    "tenant-admin",
)
if err != nil {
    log.Printf("failed to assign tenant member: %v", err)
}
```

For selected resources:

```go
err := client.AssignTenantMember(
    ctx,
    "web_shub",
    "tenant-001",
    "user-003",
    "extension",
    ira.ScopeSelected,
    []string{
        "extension-001",
        "extension-002",
    },
    []ira.Action{
        ira.ActionView,
        ira.ActionUpdate,
    },
    "tenant-admin",
)
```

---

## 6. Remove Roles

### Remove Super Admin

```go
err := client.RemoveSuperAdmin(
    ctx,
    "web_shub",
    "user-001",
    "system-admin",
)
```

### Remove Tenant Admin

```go
err := client.RemoveTenantAdmin(
    ctx,
    "web_shub",
    "tenant-001",
    "user-002",
    "super-admin",
)
```

### Remove Tenant Member

```go
err := client.RemoveTenantMember(
    ctx,
    "web_shub",
    "tenant-001",
    "user-003",
    "tenant-admin",
)
```

These methods return an error if the IAM server rejects the operation or the operation fails.

---

## 7. Grant and Revoke Resource Permissions

### Revoke Resource Access

Use `RevokeResourceAccess` to remove previously granted resource permissions.

```go
err := client.RevokeResourceAccess(
    ctx,
    "web_shub",
    "tenant-001",
    "user-003",
    "dialplan",
    ira.ScopeSelected,
    []string{"dialplan-001"},
    []ira.Action{
        ira.ActionUpdate,
    },
    "tenant-admin",
)
if err != nil {
    log.Printf("failed to revoke resource access: %v", err)
}
```

The grant operation is exposed through `AssignTenantMember`, which creates the tenant member's resource permissions.

---

## 8. Check Resource Access

Use `CheckAccess` to verify whether a user can perform a particular action on a resource.

```go
allowed, err := client.CheckAccess(
    ctx,
    "user-003",
    "tenant-001",
    "dialplan",
    "dialplan-001",
    ira.ActionUpdate,
)
if err != nil {
    log.Printf("access check failed: %v", err)
    return
}

if allowed {
    log.Println("Access granted")
} else {
    log.Println("Access denied")
}
```

Parameters:

| Parameter | Description |
|---|---|
| `ctx` | Request context. |
| `userID` | User whose access is being checked. |
| `tenantID` | Tenant containing the resource. |
| `resourceType` | Type of resource. |
| `resourceID` | Specific resource identifier. |
| `action` | Action being requested. |

The application identity is derived from the authenticated client configuration. The `CheckAccess` method therefore does not require a separate `targetAppID` argument.

**Important:** Always perform authorization checks on your backend before allowing a protected operation.

---

## 9. Get User Access Tree

Retrieve a user's access structure for a particular application.

```go
accessTree, err := client.GetUserAccessTree(
    ctx,
    "user-003",
    "web_shub",
)
if err != nil {
    log.Printf("failed to retrieve user access tree: %v", err)
    return
}

log.Printf("User access tree: %+v", accessTree)
```

This method returns the structured access information provided by the IAM server, including the user's assigned roles and resource permissions.

It is particularly useful for:

- Building access-management dashboards.
- Displaying user permissions.
- Building dynamic navigation menus.
- Managing tenant-level user access.
- Displaying resource-level permissions in an administration UI.


---

## 10. Complete Example

The following example demonstrates initialization, assigning roles, granting permissions, checking access, retrieving the access tree, and closing the client.

```go
package main

import (
    "context"
    "log"
    "time"

    ira "github.com/bhusha-n/iraaccess-client"
)

func main() {
    client, err := ira.InitializeIraAccess("config/web-shub.json")
    if err != nil {
        log.Fatalf("initialization failed: %v", err)
    }
    defer client.Close()

    ctx, cancel := context.WithTimeout(
        context.Background(),
        15*time.Second,
    )
    defer cancel()

    // 1. Assign Super Admin
    err = client.AssignSuperAdmin(
        ctx,
        "web_shub",
        "user-001",
        "system-admin",
    )
    if err != nil {
        log.Printf("AssignSuperAdmin: %v", err)
        return
    }

    // 2. Assign Tenant Admin
    err = client.AssignTenantAdmin(
        ctx,
        "web_shub",
        "tenant-001",
        "user-002",
        "user-001",
    )
    if err != nil {
        log.Printf("AssignTenantAdmin: %v", err)
        return
    }

    // 3. Assign Tenant Member
    err = client.AssignTenantMember(
        ctx,
        "web_shub",
        "tenant-001",
        "user-003",
        "dialplan",
        ira.ScopeSelected,
        []string{"dialplan-001", "dialplan-002"},
        []ira.Action{
            ira.ActionView,
            ira.ActionUpdate,
        },
        "user-002",
    )
    if err != nil {
        log.Printf("AssignTenantMember: %v", err)
        return
    }

    // 4. Check Resource Access
    allowed, err := client.CheckAccess(
        ctx,
        "user-003",
        "tenant-001",
        "dialplan",
        "dialplan-001",
        ira.ActionUpdate,
    )
    if err != nil {
        log.Printf("CheckAccess: %v", err)
        return
    }

    log.Printf("Update permission: %t", allowed)

    // 5. Retrieve User Access Tree
    accessTree, err := client.GetUserAccessTree(
        ctx,
        "user-003",
        "web_shub",
    )
    if err != nil {
        log.Printf("GetUserAccessTree: %v", err)
        return
    }

    log.Printf("Access tree: %+v", accessTree)

    log.Println("IraAccess operations completed")
}
```

---

## 11. Error Handling

The client returns errors from the underlying gRPC service.

Handle errors appropriately in your application.

Common gRPC error categories include:

| gRPC Code | Meaning |
|---|---|
| `InvalidArgument` | Invalid request parameters. |
| `Unauthenticated` | Invalid or missing application credentials. |
| `PermissionDenied` | Application or actor is not authorized to perform the operation. |
| `NotFound` | Requested entity does not exist. |
| `AlreadyExists` | Requested assignment already exists, where applicable. |
| `Internal` | Unexpected server-side error. |
| `Unavailable` | IAM service is unavailable. |

Do not treat an authorization error as permission granted.

---

## 12. Client API Reference

### Initialization and Lifecycle

| Method | Description |
|---|---|
| `InitializeIraAccess(jsonPath)` | Initialize the client from a JSON configuration. |
| `Close()` | Close the underlying gRPC connection. |

### Role Management

| Method | Description |
|---|---|
| `AssignSuperAdmin(...)` | Assign application Super Admin. |
| `AssignTenantAdmin(...)` | Assign Tenant Admin. |
| `AssignTenantMember(...)` | Assign Tenant Member with resource permissions. |
| `RemoveSuperAdmin(...)` | Remove Super Admin. |
| `RemoveTenantAdmin(...)` | Remove Tenant Admin. |
| `RemoveTenantMember(...)` | Remove Tenant Member. |

### Resource Permissions

| Method | Description |
|---|---|
| `AssignTenantMember(...)` | Grant tenant member resource permissions. |
| `RevokeResourceAccess(...)` | Revoke resource permissions. |

### Authorization

| Method | Description |
|---|---|
| `CheckAccess(...)` | Check a user's resource-level permission. |
| `GetUserAccessTree(...)` | Retrieve a user's application access structure. |

---

## 13. Important Security Considerations

- Keep `app_secret` confidential.
- Use a separate configuration for each registered application.
- Do not expose application credentials to frontend applications.
- Perform authorization checks on the backend.
- Supply `grantedBy` and `deletedBy` using the authenticated identity from your application's authentication layer. Do not blindly trust user-supplied values.
- Use TLS for gRPC communication in production. The current example client uses insecure transport credentials and should be updated to TLS before production deployment.

---

## 14. Version

**Current Version:** `v1.0.0`

**Previous Version:** `v0.3.1`

**Repository:** https://github.com/bhusha-n/iraaccess-client

**Installation:**

```bash
go get github.com/bhusha-n/iraaccess-client@v4.0.0
```

---

**IraAccess — Centralized Authorization. Application Isolation. Fine-Grained Access Control.**