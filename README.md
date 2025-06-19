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

# install dependencies and build the monitor binary
go build -o monitor
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

You can run the monitor as a background service using `systemd` to keep your
Next.js development server running indefinitely.

## Running with systemd

The repository includes an example service file `nextjs-monitor.service`.
Update the `ExecStart` line to point to the `monitor` binary and the directory
containing your Next.js app. Then copy the file to `/etc/systemd/system/`:

```bash
sudo cp nextjs-monitor.service /etc/systemd/system/
```

Reload systemd, enable the service, and start it:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now nextjs-monitor.service
```

The monitor will now run in the background and automatically restart if it
exits.
