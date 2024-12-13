#!/bin/sh

# benchmark_tool=bombardier
benchmark_tool=hey

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/json"
$benchmark_tool -n 10000 "http://localhost:8080/json"

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/db"
$benchmark_tool -n 10000 "http://localhost:8080/db"

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/queries?n=10"
$benchmark_tool -n 10000 "http://localhost:8080/queries?n=10"

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/fortunes"
$benchmark_tool -n 10000 "http://localhost:8080/fortunes"

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/updates?n=10"
$benchmark_tool -n 10000 "http://localhost:8080/updates?n=10"

echo '## ' $benchmark_tool -n 10000 "http://localhost:8080/plaintext"
$benchmark_tool -n 10000 "http://localhost:8080/plaintext"
