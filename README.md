[![Release Go project](https://github.com/theztd/better-uptime-exporter/actions/workflows/release.yml/badge.svg)](https://github.com/theztd/better-uptime-exporter/actions/workflows/release.yml)

# BetterUptime gather

Gather endpoint status from better uptime and convert them to influxdb2/prometheus


## Output

```bash
better_uptime_metrics{url="", domain="Domain from url", monitor_type="", verify_ssl="true/false", method=""} STATUS 

```


### STATUS

  - 0  OK
  - 1  DOWN
  -
  -
  - 9  Unknown
