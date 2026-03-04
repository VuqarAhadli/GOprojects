#  Port Scanner

This CLI application scans open ports on a specified address (e.g., `12.34.56.78`).

Port scanning is performed concurrently using Go workers, and you can configure the number of workers for faster or more controlled scanning.

---

##  Flags

The application supports the following flags:

| Flag        | Description                  | Default                 |
| ----------- | ---------------------------- | ----------------------- |
| `--host`    | Target host to scan          | Auto-detect local IP    |
| `--workers` | Number of concurrent workers | `20`                    |
| `--start`   | Starting port                | `1`                     |
| `--end`     | Ending port                  | `65535`                 |
| `--help`    | Show usage                   | |
---

##  Usage

>  Tip: Add the binary to your `PATH` or create an alias for convenience.

### Example

```bash
# assuming you built the binary as "scanner"
scanner --host 23.23.23.23 --workers 30 --end 32768
```

Or if running directly:

```bash
go run NetworkScanner.go --host 23.23.23.23 --workers 30 --end 32768
```

---

## How It Works

* Uses goroutines for concurrent port scanning
* Worker pool controls parallelism
* Uses TCP dial with timeout
* Reports open ports efficiently

---

##  Disclaimer

Use this tool only on systems you own or have explicit permission to test. Unauthorized port scanning may be illegal if you care enough ;)
