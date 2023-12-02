##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.1564 secs
  Slowest:	0.0123 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	63941.3041
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8459]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [1135]	|■■■■■
  0.004 [315]	|■
  0.005 [36]	|
  0.006 [6]	|
  0.007 [14]	|
  0.009 [15]	|
  0.010 [14]	|
  0.011 [1]	|
  0.012 [4]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0005 secs
  75% in 0.0009 secs
  90% in 0.0017 secs
  95% in 0.0023 secs
  99% in 0.0036 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0123 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0022 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0039 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0039 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.6974 secs
  Slowest:	0.1479 secs
  Fastest:	0.0003 secs
  Average:	0.0033 secs
  Requests/sec:	14338.5660
  
  Total data:	307707 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.015 [9385]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [570]	|■■
  0.045 [3]	|
  0.059 [6]	|
  0.074 [4]	|
  0.089 [2]	|
  0.104 [11]	|
  0.118 [8]	|
  0.133 [1]	|
  0.148 [9]	|


Latency distribution:
  10% in 0.0008 secs
  25% in 0.0013 secs
  50% in 0.0019 secs
  75% in 0.0025 secs
  90% in 0.0037 secs
  95% in 0.0181 secs
  99% in 0.0200 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.1479 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0016 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0028 secs
  resp wait:	0.0031 secs, 0.0002 secs, 0.1419 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	3.2367 secs
  Slowest:	0.2725 secs
  Fastest:	0.0011 secs
  Average:	0.0156 secs
  Requests/sec:	3089.5364
  
  Total data:	3187272 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.028 [9294]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.055 [119]	|■
  0.082 [7]	|
  0.110 [30]	|
  0.137 [383]	|■■
  0.164 [161]	|■
  0.191 [0]	|
  0.218 [0]	|
  0.245 [0]	|
  0.272 [5]	|


Latency distribution:
  10% in 0.0035 secs
  25% in 0.0043 secs
  50% in 0.0055 secs
  75% in 0.0095 secs
  90% in 0.0235 secs
  95% in 0.1291 secs
  99% in 0.1464 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0011 secs, 0.2725 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0155 secs, 0.0010 secs, 0.2724 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.8656 secs
  Slowest:	0.2528 secs
  Fastest:	0.0003 secs
  Average:	0.0039 secs
  Requests/sec:	11553.0713
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.026 [9959]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.051 [4]	|
  0.076 [18]	|
  0.101 [6]	|
  0.127 [7]	|
  0.152 [3]	|
  0.177 [0]	|
  0.202 [0]	|
  0.228 [0]	|
  0.253 [2]	|


Latency distribution:
  10% in 0.0010 secs
  25% in 0.0019 secs
  50% in 0.0025 secs
  75% in 0.0032 secs
  90% in 0.0043 secs
  95% in 0.0190 secs
  99% in 0.0214 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.2528 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:	0.0038 secs, 0.0002 secs, 0.2527 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	12.5355 secs
  Slowest:	0.2500 secs
  Fastest:	0.0068 secs
  Average:	0.0618 secs
  Requests/sec:	797.7346
  
  Total data:	3187826 bytes
  Size/request:	318 bytes

Response time histogram:
  0.007 [1]	|
  0.031 [44]	|
  0.055 [4233]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.080 [4787]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.104 [746]	|■■■■■■
  0.128 [74]	|■
  0.153 [43]	|
  0.177 [24]	|
  0.201 [31]	|
  0.226 [6]	|
  0.250 [11]	|


Latency distribution:
  10% in 0.0494 secs
  25% in 0.0525 secs
  50% in 0.0567 secs
  75% in 0.0674 secs
  90% in 0.0790 secs
  95% in 0.0879 secs
  99% in 0.1464 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0068 secs, 0.2500 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0013 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:	0.0617 secs, 0.0067 secs, 0.2500 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.1362 secs
  Slowest:	0.0100 secs
  Fastest:	0.0001 secs
  Average:	0.0007 secs
  Requests/sec:	73403.9218
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8359]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [1140]	|■■■■■
  0.003 [349]	|■■
  0.004 [87]	|
  0.005 [15]	|
  0.006 [27]	|
  0.007 [12]	|
  0.008 [0]	|
  0.009 [5]	|
  0.010 [5]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0004 secs
  75% in 0.0007 secs
  90% in 0.0015 secs
  95% in 0.0020 secs
  99% in 0.0035 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0100 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0032 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0043 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0028 secs
  resp read:	0.0002 secs, 0.0000 secs, 0.0057 secs

Status code distribution:
  [200]	10000 responses



