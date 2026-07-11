# Repository Organization

> **Version:** 1.0
> **Status:** Draft

---

## Overview

This document defines how the ZoneBridge repository is organized.

A consistent repository organization improves maintainability, reduces onboarding time, and provides a predictable development experience for contributors.

The repository organization reflects the architectural principles of the platform and serves as the foundation for all implementation work.

---

## Purpose

The repository is organized to achieve the following goals:

- Encourage modular development.
- Promote clear ownership.
- Reduce coupling.
- Improve discoverability.
- Support long-term maintenance.
- Enable independent evolution of platform contexts.

---

# Repository Philosophy

ZoneBridge is organized around **bounded contexts** rather than technical layers.

Each bounded context owns its implementation from end to end.

This includes:

- Domain models
- Services
- Persistence
- HTTP handlers
- Routes
- Tests

By grouping related code together, contributors can understand a complete capability without navigating unrelated parts of the codebase.

---

# Top-Level Layout

```text
zonebridge/

.github/
backend/
frontend/
docs/
deployments/
scripts/

LICENSE
README.md
CHANGELOG.md
```

Each top-level directory has a single responsibility.

---

# Backend

The backend contains the Go implementation of the platform.

```text
backend/

cmd/
configs/
internal/
migrations/
tests/

Dockerfile
go.mod
go.sum
```

---

## Backend Organization

The backend follows a bounded-context architecture.

```text
internal/

application/
shared/

identity/
community/
knowledge/
collaboration/
audit/
spaces/
notification/

integrations/
```

Each context contains its complete implementation.

Example:

```text
audit/

handlers.go

service.go

repository.go

models.go

routes.go

service_test.go
```

Contexts should not expose internal implementation details to one another.

---

# Frontend

The frontend mirrors the backend architecture wherever possible.

```text
frontend/

public/

src/
```

---

## Frontend Organization

```text
src/

app/

capabilities/

shared/

providers/

hooks/

styles/

assets/
```

Capabilities correspond directly to backend bounded contexts.

Example:

```text
capabilities/

audit/

community/

knowledge/

profile/

spaces/
```

---

# Documentation

Documentation is treated as part of the product.

```text
docs/

platform/

architecture/

backend/

frontend/

deployment/

database/

adr/

standards/
```

Every architectural decision should be documented.

---

# Deployments

Deployment resources are isolated from application code.

```text
deployments/

docker/

cloudflare/

oracle/

github-actions/
```

Deployment configuration should remain reproducible and version-controlled.

---

# Scripts

Development automation belongs in the scripts directory.

Examples include:

- Build scripts
- Development helpers
- Test runners
- Database utilities

Scripts should remain platform-independent where practical.

---

# Naming Conventions

The repository follows the following naming conventions.

## Directories

Use lowercase.

Example:

```text
community
notification
deployment
```

---

## Go Packages

Package names should:

- Be singular where appropriate.
- Be descriptive.
- Avoid abbreviations.
- Follow idiomatic Go naming conventions.

---

## Documentation

Documentation uses lowercase with hyphens.

Examples:

```text
platform-overview.md

system-architecture.md

repository-organization.md
```

---

## Tests

Tests live beside the code they validate.

Example:

```text
service.go

service_test.go
```

---

# Engineering Rules

The following rules apply throughout the repository.

- Every directory has a single responsibility.
- Every bounded context owns its implementation.
- Business logic belongs in services.
- Infrastructure supports business logic.
- Shared code should remain minimal.
- Dependencies point inward.
- Documentation evolves alongside implementation.
- Every feature includes automated tests.

---

# Repository Growth

As the platform evolves, new functionality should extend existing bounded contexts before introducing new ones.

New top-level directories should only be introduced when a clear architectural boundary exists.

Repository organization should remain stable over time.

---

## Related Documents

- [Platform Overview](../platform/platform-overview.md)
- [Architecture Overview](architecture-overview.md)
- [System Architecture](system-architecture.md)
- [Domain Model](domain-model.md)
- [Workflows](workflows.md)