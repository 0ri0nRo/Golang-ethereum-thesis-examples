# Use a base Ubuntu image
FROM ubuntu:latest

# Update package index and install necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    software-properties-common \
    && add-apt-repository -y ppa:ethereum/ethereum \
    && apt-get update \
    && apt-get install -y ethereum

# Run the ghcr.io/edgelesssys/ego-dev image
CMD ["docker", "run", "-it", "ghcr.io/edgelesssys/ego-dev"]
