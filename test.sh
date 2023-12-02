#!/bin/sh

echo '## ' hey -n 10000 "http://localhost:8080/json"
hey -n 10000 "http://localhost:8080/json"

echo '## ' hey -n 10000 "http://localhost:8080/db"
hey -n 10000 "http://localhost:8080/db"

echo '## ' hey -n 10000 "http://localhost:8080/queries?n=10"
hey -n 10000 "http://localhost:8080/queries?n=10"

echo '## ' hey -n 10000 "http://localhost:8080/fortunes"
hey -n 10000 "http://localhost:8080/fortunes"

echo '## ' hey -n 10000 "http://localhost:8080/updates?n=10"
hey -n 10000 "http://localhost:8080/updates?n=10"

echo '## ' hey -n 10000 "http://localhost:8080/plaintext"
hey -n 10000 "http://localhost:8080/plaintext"
