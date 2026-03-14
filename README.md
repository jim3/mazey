# Mazey

 for threat triage.


**Mazey** is an early-stage CLI reconnaissance tool for threat triage. It takes *inbound noise* such as automated scans, bots, misconfigured devices and enriches them using various threat intelligence API's like Virus Total, Shodan, etc...

## Why the name?
`Mazey` is named in tribute to my cat. This is a personal project with long-term goals! 🐈


## Current features
- `blacklist [count]` command
- Reads blacklist source from `API_ENDPOINT`
- Enriches each IP through Shodan InternetDB


## Tech stack
- Go
- Cobra + Fang (CLI framework / UX)
- `godotenv` for local env loading

## Quick start

### 1) Set environment variable
Copy the template and edit values:

```bash
cp .env.example .env
```

Then set values in `.env`:

```env
API_ENDPOINT=https://127.0.0.1:8000/blacklist
VT_API_KEY=replace-with-your-virustotal-api-key
```

### 2) Build or run
```bash
make build
mazey filereport 9b97edcbd8099796015c78bbf1723b35
```

## Make targets
- `make help` - list available targets
- `make build` - build binary
- `make run ARGS="..."` - run CLI with args
- `make test` - run tests
- `make fmt` - format Go code
- `make vet` - run static checks
- `make tidy` - clean module dependencies

## Roadmap (high level)

- Add more intelligence sources and richer reporting
- Improve API error handling for 4XX/5XX responses and clearer CLI output
- Improve command coverage and output formatting
- Evolve into a more complete incident triage assistant


## Preview

### Mazey Help & Command Menus

---

![Mazey Help Menu](assets/mazey-cli-help.png)

---
<br>

![Mazey Blacklist](assets/mazey-cli-help-command.png)


