# Tiny Health Checker

## Usage

1. Create `config.yaml` file from the `config.example.yaml` example.

2. Edit `config.yaml` to your needs. Like
    ```yaml
    target:                           # Target
    - name: "target1"               # string
      url: "http://localhost/ping"  # string
      method: "POST"                # string
      headers:                      # array<map<string, string>>
        - name: "Content-Type"      # string
          value: "application/json" # string
      body: ""                      # string (default: "")
      timeout: 3                    # int (seconds, default: 1)
      retry:                        # Retry
        count: 3                    # int (second, default: 1)
        interval: 1                 # int (second, default: 1)
      ssl_verify: false             # bool (default: true)
      status: 200                   # int - Check response status code
    alert:                            # Alert
    - type: "telegram"              # enum<telegram, std> 
      name: "Telegram Alert Bot"    # string (any name)
      token: "1234:AAE0J"           # string (telegram bot token)
      chat_id: 1                    # bigint (telegram chat id)
    - type: "std"                   # enum<telegram, std>
      name: "STD Log"               # string (any name)
    
    ```

3. Add crontab like this (for per minute):
    ```bash
    * * * * * /path/to/tiny-health-checker -config /path/to/config.yaml
    ```

#### `('.')` Enjoy!
