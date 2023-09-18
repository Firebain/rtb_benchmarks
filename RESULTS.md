Node: https://mgaididei.grafana.net/a/k6-app/runs/1948343
Go: https://mgaididei.grafana.net/a/k6-app/runs/1948338
Rust: https://mgaididei.grafana.net/a/k6-app/runs/1948330

|      | Requests | Req/s    | Latency | Node diff | Go diff |
| ---- | -------- | -------- | ------- | --------- | ------- |
| Node | 135K     | 1 845.67 | 59      | 0         | -40%    |
| Go   | 220.2K   | 3 060.67 | 34      | +66%      | 0       |
| Rust | 275.1K   | 3 973.67 | 26      | +115%     | +30%    |

Node: https://mgaididei.grafana.net/a/k6-app/runs/1950394
Go std: https://mgaididei.grafana.net/a/k6-app/runs/1950400
Go fiber: https://mgaididei.grafana.net/a/k6-app/runs/1950406
C#: https://mgaididei.grafana.net/a/k6-app/runs/1950421
Rust: https://mgaididei.grafana.net/a/k6-app/runs/1950430

|          | Requests | Req/s    | Latency |
| -------- | -------- | -------- | ------- |
| Node     | 138.5K   | 1 858.33 | 56      |
| Go std   | 219K     | 3 012.67 | 34      |
| Go fiber | 221.9K   | 3 082.67 | 33      |
| C#       | 170.7K   | 2 338    | 46      |
| Rust     | 254.9K   | 3 641.67 | 29      |
