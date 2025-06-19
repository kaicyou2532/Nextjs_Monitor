# Next.js Monitor

This repository provides a small Go program that watches a Next.js development server.
It periodically checks a URL and, if the server is down, starts `npm run dev` in the
configured directory.

## Prerequisites

- Go 1.23 or later
- Node.js and npm
- A Next.js project that is launched with `npm run dev`

## Building

```bash
# clone the repository and change into it
cd Nextjs_Monitor

# install dependencies and build
go build
```

## Running

The monitor is configured using command line flags:

- `-dir` – directory where `npm run dev` should be executed (default: current directory)
- `-url` – URL used for the health check (default: `http://localhost:3000`)
- `-interval` – how often to run the check (default: `1m`)
- `-pattern` – pattern used with `pgrep` to detect the running `npm` process (default: `npm.*run.*dev`)

Example:

```bash
./monitor -dir /path/to/nextjs/app -url http://localhost:3000 -interval 30s
```

The program will check the provided URL every 30 seconds. If the server is not
responding and no `npm run dev` process is found, the monitor starts the server in
the given directory.

You can run the monitor as a background service (e.g., using `systemd`) to keep
your Next.js development server alive.
