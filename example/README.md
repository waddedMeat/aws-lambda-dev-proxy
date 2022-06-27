# Example Lambda


The lambda does not hot reload.

The docker-compose service does not mount external files into the running container.

## Running the Example

Start the container with
`docker compose up --build`

The `--build` is required to re-build the lambda container if the files have changed


Stop the container with `Ctrl+C`
