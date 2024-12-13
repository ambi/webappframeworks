##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0945 secs
  Slowest:	0.0081 secs
  Fastest:	0.0000 secs
  Average:	0.0005 secs
  Requests/sec:	105794.7766
  
  Total data:	270000 bytes
  Size/request:	27 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [8919]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [965]	|■■■■
  0.002 [51]	|
  0.003 [14]	|
  0.004 [0]	|
  0.005 [0]	|
  0.006 [0]	|
  0.006 [0]	|
  0.007 [6]	|
  0.008 [44]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0003 secs
  75% in 0.0005 secs
  90% in 0.0009 secs
  95% in 0.0011 secs
  99% in 0.0018 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0081 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0004 secs, 0.0000 secs, 0.0030 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2846 secs
  Slowest:	0.0215 secs
  Fastest:	0.0004 secs
  Average:	0.0014 secs
  Requests/sec:	35131.5567
  
  Total data:	307765 bytes
  Size/request:	30 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [9730]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [218]	|■
  0.007 [1]	|
  0.009 [0]	|
  0.011 [0]	|
  0.013 [0]	|
  0.015 [0]	|
  0.017 [0]	|
  0.019 [0]	|
  0.021 [50]	|


Latency distribution:
  10% in 0.0010 secs
  25% in 0.0010 secs
  50% in 0.0012 secs
  75% in 0.0015 secs
  90% in 0.0018 secs
  95% in 0.0022 secs
  99% in 0.0037 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0215 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0014 secs, 0.0004 secs, 0.0196 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.7922 secs
  Slowest:	0.0418 secs
  Fastest:	0.0053 secs
  Average:	0.0090 secs
  Requests/sec:	5579.5890
  

Response time histogram:
  0.005 [1]	|
  0.009 [8175]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.013 [1695]	|■■■■■■■■
  0.016 [48]	|
  0.020 [31]	|
  0.024 [0]	|
  0.027 [0]	|
  0.031 [0]	|
  0.035 [1]	|
  0.038 [19]	|
  0.042 [30]	|


Latency distribution:
  10% in 0.0084 secs
  25% in 0.0085 secs
  50% in 0.0087 secs
  75% in 0.0088 secs
  90% in 0.0092 secs
  95% in 0.0099 secs
  99% in 0.0146 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0053 secs, 0.0418 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0089 secs, 0.0053 secs, 0.0418 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.4919 secs
  Slowest:	0.0088 secs
  Fastest:	0.0004 secs
  Average:	0.0024 secs
  Requests/sec:	20330.0533
  

Response time histogram:
  0.000 [1]	|
  0.001 [16]	|
  0.002 [1477]	|■■■■■■■■
  0.003 [7748]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [430]	|■■
  0.005 [143]	|■
  0.005 [80]	|
  0.006 [48]	|
  0.007 [41]	|
  0.008 [9]	|
  0.009 [7]	|


Latency distribution:
  10% in 0.0020 secs
  25% in 0.0022 secs
  50% in 0.0023 secs
  75% in 0.0025 secs
  90% in 0.0028 secs
  95% in 0.0033 secs
  99% in 0.0055 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0088 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0006 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:	0.0024 secs, 0.0004 secs, 0.0083 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0017 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	4.0445 secs
  Slowest:	0.0556 secs
  Fastest:	0.0164 secs
  Average:	0.0202 secs
  Requests/sec:	2472.5153
  

Response time histogram:
  0.016 [1]	|
  0.020 [7860]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.024 [1652]	|■■■■■■■■
  0.028 [269]	|■
  0.032 [148]	|■
  0.036 [20]	|
  0.040 [0]	|
  0.044 [0]	|
  0.048 [0]	|
  0.052 [0]	|
  0.056 [50]	|


Latency distribution:
  10% in 0.0188 secs
  25% in 0.0191 secs
  50% in 0.0195 secs
  75% in 0.0201 secs
  90% in 0.0217 secs
  95% in 0.0236 secs
  99% in 0.0301 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0164 secs, 0.0556 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0202 secs, 0.0163 secs, 0.0556 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0740 secs
  Slowest:	0.0032 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	135161.6195
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [7140]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1720]	|■■■■■■■■■■
  0.001 [778]	|■■■■
  0.001 [257]	|■
  0.002 [49]	|
  0.002 [23]	|
  0.002 [21]	|
  0.003 [6]	|
  0.003 [1]	|
  0.003 [4]	|


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0004 secs
  90% in 0.0007 secs
  95% in 0.0009 secs
  99% in 0.0013 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0032 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0018 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0019 secs

Status code distribution:
  [200]	10000 responses



