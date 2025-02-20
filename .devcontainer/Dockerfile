FROM mcr.microsoft.com/devcontainers/go:1-bullseye

# Install PostgreSQL client
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client