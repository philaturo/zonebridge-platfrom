# Docker Development Environment

This directory contains the containerized development environment for ZoneBridge.

## Purpose

The Docker environment provides reproducible local infrastructure for contributors.

Application code is developed locally while infrastructure services such as PostgreSQL are provided through containers.

## Contents

```text
compose.yaml

Dockerfile.backend

Dockerfile.frontend
```

## Philosophy

Containers exist to support development, testing and deployment.

Application source code remains independent of the container runtime.