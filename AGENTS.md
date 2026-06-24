# AGENTS.md — mlops-k8s-deployment

## What this repo is

A **Python-based MLOps Kubernetes deployment** scaffold. Currently a single initial commit with no application code.

## Repo state

- **Only tracked files:** `.gitignore` (Python template), `README.md` (title-only)
- **Untracked:** `kind-cluster.yaml` — 1 control-plane + 1 worker node local cluster (KinD)
- **No** build system, dependency manager, CI/CD, Dockerfiles, k8s manifests, or instruction files yet

## Known layout

```
.gitignore                  # Python-focused ignore rules
kind-cluster.yaml           # KinD cluster definition (local dev)
README.md                   # Repo title only
```

## Local dev prerequisites

- [kind](https://kind.sigs.k8s.io/) — local Kubernetes cluster (config: `kind-cluster.yaml`)
- `kind create cluster --config kind-cluster.yaml` to bootstrap

## Python conventions

- `.gitignore` expects Python tooling — add `pyproject.toml` and a dependency manager (e.g., pip, poetry, uv) before writing app code
- `.venv/`, `__pycache__/`, `.ruff_cache/`, `.mypy_cache/`, `.pytest_cache/` are already ignored

## What needs to be built

From scratch: Python app(s), Dockerfiles, Helm/k8s manifests, CI/CD, dependency management, testing.
