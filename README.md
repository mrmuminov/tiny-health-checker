# Tiny Health Checker

## Usage

1. Copy `config.example.yaml` to `config.yaml`.

2. Edit `config.yaml` to your needs. Like
    ```yaml
    target:
      - name: "target1"
        url: "http://localhost/ping"
        method: "POST"
        headers:
          - name: "Content-Type"
            value: "application/json"
        body: ""
        timeout: 3
    ```
3. Add crontab like this (for per minute):
   ```bash
   * * * * * /path/to/tiny-health-checker
   ```
   

#### `('.')`Enjoy!
