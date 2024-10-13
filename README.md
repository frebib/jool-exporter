# Jool NAT64 translator Prometheus exporter

Exports statistics from `jool stats display` for all instances in a Prometheus metrics format

Takes almost no configuration:
```
Usage:
  jool-exporter

Application Options:
  -l, --listen=          http listen address to serve metrics (default: :9441)
  -p, --metrics-path=    http path at which to serve metrics (default: /metrics)
      --web.config-file= path to web-config file
```

### Example metrics

```
# HELP jool_instance_enabled 1 indicates that the instance is enabled, otherwise 0
# TYPE jool_instance_enabled gauge
jool_instance_enabled{instance="default"} 1
# HELP jool_memory_allocation_failures_total Memory allocation failures
# TYPE jool_memory_allocation_failures_total counter
jool_memory_allocation_failures_total{instance="default"} 0
# HELP jool_received_packets_total Total number of packets received by the translator
# TYPE jool_received_packets_total counter
jool_received_packets_total{instance="default",ipversion="4"} 8.365786e+06
jool_received_packets_total{instance="default",ipversion="6"} 1.294595e+07
# HELP jool_translation_error_total Total number of errors that accounted for a failed/cancelled packet translation
# TYPE jool_translation_error_total counter
jool_translation_error_total{error="64-ttl",instance="default"} 111
jool_translation_error_total{error="bib4-not-found",instance="default"} 7
jool_translation_error_total{error="failed-routes",instance="default"} 19
jool_translation_error_total{error="icmp6err-success",instance="default"} 111
jool_translation_error_total{error="icmpext-big",instance="default"} 1161
jool_translation_error_total{error="pool4-mismatch",instance="default"} 7.174768e+06
jool_translation_error_total{error="pool6-mismatch",instance="default"} 1.2179576e+07
jool_translation_error_total{error="skb-truncated",instance="default"} 874
jool_translation_error_total{error="syn4-expected",instance="default"} 1048
jool_translation_error_total{error="syn6-expected",instance="default"} 2770
jool_translation_error_total{error="type1pkt",instance="default"} 7
jool_translation_error_total{error="unknown",instance="default"} 1294
jool_translation_error_total{error="unknown-icmp4-type",instance="default"} 51
jool_translation_error_total{error="unknown-icmp6-type",instance="default"} 5941
jool_translation_error_total{error="unknown-l4_proto",instance="default"} 24
# HELP jool_translation_success_total Total number of successful translated packets (does not imply successful delivery)
# TYPE jool_translation_success_total counter
jool_translation_success_total{instance="default"} 1.944085e+06
```
