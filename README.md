# Tiny Health Checker

## Usage

1. Copy `config.example.yaml` to `config.yaml`.

2. Edit `config.yaml` to your needs. Like
 ```yaml
 target:                           # Target
   - name: "target1"               # string
     url: "http://localhost/ping"  # string
     method: "POST"                # string
     headers:                      # array<map<string, string>>
       - name: "Content-Type"      # string
         value: "application/json" # string
     body: ""                      # string
     timeout: 3                    # int (seconds)
     retry:                        # Retry
       count: 3                    # int
       interval: 1                 # int (seconds)
 ```
3. Add crontab like this (for per minute):
```bash
* * * * * /path/to/tiny-health-checker
```
   

#### `('.')` Enjoy!
