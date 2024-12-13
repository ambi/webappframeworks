##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.5236 secs
  Slowest:	0.0994 secs
  Fastest:	0.0001 secs
  Average:	0.0025 secs
  Requests/sec:	19096.7462
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.010 [9950]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.020 [9]	|
  0.030 [8]	|
  0.040 [5]	|
  0.050 [6]	|
  0.060 [6]	|
  0.070 [4]	|
  0.080 [4]	|
  0.089 [3]	|
  0.099 [4]	|


Latency distribution:
  10% in 0.0017 secs
  25% in 0.0020 secs
  50% in 0.0023 secs
  75% in 0.0025 secs
  90% in 0.0030 secs
  95% in 0.0034 secs
  99% in 0.0051 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0994 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:	0.0025 secs, 0.0000 secs, 0.0945 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.6776 secs
  Slowest:	0.0557 secs
  Fastest:	0.0001 secs
  Average:	0.0033 secs
  Requests/sec:	14756.9584
  
  Total data:	307847 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.006 [9914]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [35]	|
  0.017 [1]	|
  0.022 [14]	|
  0.028 [8]	|
  0.033 [4]	|
  0.039 [7]	|
  0.045 [5]	|
  0.050 [5]	|
  0.056 [6]	|


Latency distribution:
  10% in 0.0028 secs
  25% in 0.0028 secs
  50% in 0.0030 secs
  75% in 0.0035 secs
  90% in 0.0041 secs
  95% in 0.0045 secs
  99% in 0.0054 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0557 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0033 secs, 0.0001 secs, 0.0543 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.9000 secs
  Slowest:	0.0220 secs
  Fastest:	0.0015 secs
  Average:	0.0095 secs
  Requests/sec:	5263.2366
  
  Total data:	3187762 bytes
  Size/request:	318 bytes

Response time histogram:
  0.001 [1]	|
  0.004 [6]	|
  0.006 [18]	|
  0.008 [31]	|
  0.010 [7801]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [2021]	|■■■■■■■■■■
  0.014 [93]	|
  0.016 [6]	|
  0.018 [6]	|
  0.020 [8]	|
  0.022 [9]	|


Latency distribution:
  10% in 0.0090 secs
  25% in 0.0092 secs
  50% in 0.0094 secs
  75% in 0.0097 secs
  90% in 0.0099 secs
  95% in 0.0100 secs
  99% in 0.0128 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0015 secs, 0.0220 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0095 secs, 0.0015 secs, 0.0207 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.8816 secs
  Slowest:	0.0370 secs
  Fastest:	0.0004 secs
  Average:	0.0044 secs
  Requests/sec:	11342.3775
  
  Total data:	14970000 bytes
  Size/request:	1497 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [4066]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [5832]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [51]	|
  0.015 [31]	|
  0.019 [3]	|
  0.022 [3]	|
  0.026 [4]	|
  0.030 [3]	|
  0.033 [3]	|
  0.037 [3]	|


Latency distribution:
  10% in 0.0039 secs
  25% in 0.0040 secs
  50% in 0.0042 secs
  75% in 0.0045 secs
  90% in 0.0049 secs
  95% in 0.0051 secs
  99% in 0.0078 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0370 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0044 secs, 0.0004 secs, 0.0354 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	4.2372 secs
  Slowest:	0.1095 secs
  Fastest:	0.0145 secs
  Average:	0.0212 secs
  Requests/sec:	2360.0391
  
  Total data:	3187903 bytes
  Size/request:	318 bytes

Response time histogram:
  0.015 [1]	|
  0.024 [9148]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.034 [593]	|■■■
  0.043 [8]	|
  0.053 [78]	|
  0.062 [122]	|■
  0.072 [0]	|
  0.081 [0]	|
  0.091 [0]	|
  0.100 [4]	|
  0.109 [46]	|


Latency distribution:
  10% in 0.0186 secs
  25% in 0.0189 secs
  50% in 0.0195 secs
  75% in 0.0205 secs
  90% in 0.0234 secs
  95% in 0.0274 secs
  99% in 0.0539 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0145 secs, 0.1095 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0212 secs, 0.0145 secs, 0.1095 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.4413 secs
  Slowest:	0.0497 secs
  Fastest:	0.0000 secs
  Average:	0.0021 secs
  Requests/sec:	22661.4258
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [9908]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [61]	|
  0.015 [6]	|
  0.020 [5]	|
  0.025 [4]	|
  0.030 [3]	|
  0.035 [3]	|
  0.040 [3]	|
  0.045 [3]	|
  0.050 [3]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0018 secs
  50% in 0.0019 secs
  75% in 0.0023 secs
  90% in 0.0028 secs
  95% in 0.0031 secs
  99% in 0.0048 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0497 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0021 secs, 0.0000 secs, 0.0480 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200]	10000 responses



