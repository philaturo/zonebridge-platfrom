# ZoneBridge

Collaborative platform for the Zone01 Kisumu community.

ZoneBridge is an open-source platform that extends the Zone01 Kisumu learning experience through peer collaboration, developer communities, mentorship, audit coordination, knowledge sharing, hackathons, bootcamps, and open-source initiatives.

Rather than replacing the official Zone01 platform, ZoneBridge complements it by providing the collaborative platform that surrounds project-based education while preserving the principles of peer learning and independent problem solving.

---

## Documentation

[Architecture](docs/architecture/system-architecture.md) •
[Platform](docs/platform/platform-overview.md) •
[Contributing](CONTRIBUTING.md) •
[Security](SECURITY.md) •
[Changelog](CHANGELOG.md) •
[License](LICENSE)

---

## Why ZoneBridge?

The Zone01 learning model is built around projects, peer learning, audits, and collaboration.

Every project generates valuable implementation knowledge, technical decisions, debugging experiences, and lessons learned. Much of this knowledge becomes fragmented across repositories, messaging platforms, and personal notes, making it difficult for future apprentices to benefit from previous work.

ZoneBridge exists to strengthen the collaborative side of the Zone01 ecosystem by making knowledge easier to discover, communities easier to build, and collaboration easier to sustain.

The platform is designed to complement—not replace—the official Zone01 systems.

---

# Platform

ZoneBridge brings together the collaborative aspects of the Zone01 Kisumu community into a single platform.

## Communities

Technology-focused communities where apprentices collaborate around programming languages, frameworks, cloud infrastructure, cybersecurity, artificial intelligence, blockchain, DevOps, open source, and emerging technologies.

Communities provide discussion spaces, shared resources, announcements, collaborative learning opportunities, and technical events.

---

## Knowledge Hub

A centralized knowledge base where apprentices publish technical articles, implementation notes, project write-ups, debugging experiences, architectural discussions, and learning resources.

Knowledge remains discoverable long after projects have been completed.

---

## Mentorship

A structured peer-learning environment that enables apprentices to discover experienced members, request mentorship, pair program, exchange technical knowledge, and support one another throughout their learning journey.

---

## Audit Coordination

Community-driven audit coordination that simplifies requesting audits, discovering available auditors, managing availability, and organizing peer evaluations while remaining complementary to the official Zone01 evaluation workflow.

---

## Bootcamps

Dedicated collaborative spaces for technical bootcamps and intensive learning programs where participants can communicate, access shared resources, collaborate with mentors, and continue discussions beyond scheduled sessions.

---

## Hackathons

Support for community hackathons through team formation, announcements, project collaboration, technical discussions, mentor support, and project showcasing.

---

## Open Source

A collaborative environment where apprentices discover repositories, contribute to community projects, participate in open-source initiatives, and gain experience through real-world software development.

---

# Architecture

ZoneBridge is implemented as a **modular monolith** organized around **bounded contexts**.

Each context owns its business logic, services, persistence, HTTP handlers, routes, and tests while sharing a common platform foundation.

This architecture promotes maintainability, scalability, and clear ownership without introducing unnecessary operational complexity.

For more information, see the Architecture documentation.

→ **docs/architecture/**

---

# Engineering Principles

Every engineering decision within ZoneBridge is guided by a shared set of principles.

- People before complexity.
- Documentation is part of the product.
- Modular architecture.
- Security by design.
- Performance through simplicity.
- Production-ready engineering.
- Open collaboration.

---

# Technology Stack

| Layer | Technology | Purpose |
|--------|------------|---------|
| Frontend | React · TypeScript | User Interface |
| Backend | Go | Platform Services |
| Database | PostgreSQL | Persistent Storage |
| Cache | Redis | Background Processing & Caching |
| Realtime | WebSockets | Presence & Collaboration |
| Infrastructure | Docker | Local Development |
| Deployment | Cloudflare · Oracle Cloud | Platform Hosting |
| CI/CD | GitHub Actions | Continuous Integration |

---

# Repository

```text
zonebridge-platform/

├── .github/
├── backend/
├── frontend/
├── docs/
├── deployments/
├── scripts/
│
├── README.md
├── LICENSE
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
├── CHANGELOG.md
├── Makefile
├── Taskfile.yml
├── .editorconfig
├── .gitignore
└── .gitattributes
```

The repository follows a documentation-first engineering approach where architectural decisions, implementation, testing, and deployment evolve together.

---

# Development

Clone the repository.

```bash
git clone https://github.com/philaturo/zonebridge-platform.git

cd zonebridge-platform
```

Install project dependencies.

```bash
task setup
```

Start the development environment.

```bash
task dev
```

See the [Contributing Guide](CONTRIBUTING.md) for the complete development workflow.

---

# Documentation

Additional documentation is available throughout the repository.

| Document | Description |
|-----------|-------------|
| [Platform](docs/platform/platform-overview.md) | Platform vision and design |
| [Architecture](docs/architecture/system-architecture.md) | System architecture and engineering decisions |
| [Backend](docs/backend/README.md) | Backend implementation |
| [Frontend](docs/frontend/README.md) | Frontend architecture |
| [Database](docs/database/README.md) | Database design |
| [Deployment](docs/deployment/README.md) | Deployment strategy |
| [Contributing Guide](CONTRIBUTING.md) | Contribution workflow |
| [Security Policy](SECURITY.md) | Security reporting |
| [Code of Conduct](CODE_OF_CONDUCT.md) | Community guidelines |
| [Changelog](CHANGELOG.md) | Release history |

---

# Contributing

ZoneBridge is an open-source project built by and for the Zone01 Kisumu community.

Whether you're improving documentation, fixing bugs, proposing ideas, reviewing pull requests, or building new functionality, your contributions help strengthen the platform for the entire community.

Please read the [Contributing Guide](CONTRIBUTING.md) before submitting issues or pull requests.

---

# License

ZoneBridge is released under the MIT License.

See the [LICENSE](LICENSE) file for details.

---

Built for the Zone01 Kisumu community.

Designed with a documentation-first and production-grade engineering approach.
