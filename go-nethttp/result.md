##  hey -n 10000 http://localhost:8080/json

Summary:
  Total:	0.0890 secs
  Slowest:	0.0060 secs
  Fastest:	0.0000 secs
  Average:	0.0004 secs
  Requests/sec:	112413.3890
  
  Total data:	280000 bytes
  Size/request:	28 bytes

Response time histogram:
  0.000 [1]	|
  0.001 [7825]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [1676]	|■■■■■■■■■
  0.002 [436]	|■■
  0.002 [11]	|
  0.003 [1]	|
  0.004 [0]	|
  0.004 [0]	|
  0.005 [0]	|
  0.005 [0]	|
  0.006 [50]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0006 secs
  90% in 0.0010 secs
  95% in 0.0012 secs
  99% in 0.0015 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0060 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0020 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0014 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/db

Summary:
  Total:	0.2255 secs
  Slowest:	0.0229 secs
  Fastest:	0.0000 secs
  Average:	0.0011 secs
  Requests/sec:	44349.2495
  
  Total data:	317776 bytes
  Size/request:	31 bytes

Response time histogram:
  0.000 [1]	|
  0.002 [9926]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [29]	|
  0.007 [3]	|
  0.009 [8]	|
  0.011 [5]	|
  0.014 [4]	|
  0.016 [7]	|
  0.018 [5]	|
  0.021 [5]	|
  0.023 [7]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0005 secs
  50% in 0.0013 secs
  75% in 0.0015 secs
  90% in 0.0016 secs
  95% in 0.0016 secs
  99% in 0.0020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0229 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0009 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:	0.0011 secs, 0.0000 secs, 0.0214 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/queries?n=10

Summary:
  Total:	1.1268 secs
  Slowest:	0.4310 secs
  Fastest:	0.0003 secs
  Average:	0.0052 secs
  Requests/sec:	8874.3422
  
  Total data:	3197723 bytes
  Size/request:	319 bytes

Response time histogram:
  0.000 [1]	|
  0.043 [9832]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.086 [26]	|
  0.129 [85]	|
  0.173 [52]	|
  0.216 [1]	|
  0.259 [2]	|
  0.302 [0]	|
  0.345 [0]	|
  0.388 [0]	|
  0.431 [1]	|


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0024 secs
  50% in 0.0029 secs
  75% in 0.0040 secs
  90% in 0.0058 secs
  95% in 0.0064 secs
  99% in 0.1282 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.4310 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0010 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.0052 secs, 0.0002 secs, 0.4310 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0001 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/fortunes

Summary:
  Total:	0.2909 secs
  Slowest:	0.1403 secs
  Fastest:	0.0001 secs
  Average:	0.0012 secs
  Requests/sec:	34378.4674
  
  Total data:	19750000 bytes
  Size/request:	1975 bytes

Response time histogram:
  0.000 [1]	|
  0.014 [9957]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [1]	|
  0.042 [0]	|
  0.056 [36]	|
  0.070 [3]	|
  0.084 [0]	|
  0.098 [0]	|
  0.112 [0]	|
  0.126 [1]	|
  0.140 [1]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0005 secs
  75% in 0.0016 secs
  90% in 0.0019 secs
  95% in 0.0031 secs
  99% in 0.0072 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.1403 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0003 secs
  resp wait:	0.0012 secs, 0.0001 secs, 0.1394 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/updates?n=10

Summary:
  Total:	3.0644 secs
  Slowest:	0.0417 secs
  Fastest:	0.0017 secs
  Average:	0.0152 secs
  Requests/sec:	3263.3117
  
  Total data:	3197754 bytes
  Size/request:	319 bytes

Response time histogram:
  0.002 [1]	|
  0.006 [26]	|
  0.010 [107]	|■
  0.014 [1394]	|■■■■■■■
  0.018 [7842]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.022 [494]	|■■■
  0.026 [121]	|■
  0.030 [10]	|
  0.034 [0]	|
  0.038 [2]	|
  0.042 [3]	|


Latency distribution:
  10% in 0.0132 secs
  25% in 0.0143 secs
  50% in 0.0152 secs
  75% in 0.0160 secs
  90% in 0.0170 secs
  95% in 0.0181 secs
  99% in 0.0227 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0017 secs, 0.0417 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0002 secs
  resp wait:	0.0152 secs, 0.0017 secs, 0.0417 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200]	10000 responses



##  hey -n 10000 http://localhost:8080/plaintext

Summary:
  Total:	0.0706 secs
  Slowest:	0.0028 secs
  Fastest:	0.0000 secs
  Average:	0.0003 secs
  Requests/sec:	141607.0379
  
  Total data:	130000 bytes
  Size/request:	13 bytes

Response time histogram:
  0.000 [1]	|
  0.000 [5982]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.001 [2355]	|■■■■■■■■■■■■■■■■
  0.001 [965]	|■■■■■■
  0.001 [477]	|■■■
  0.001 [114]	|■
  0.002 [60]	|
  0.002 [24]	|
  0.002 [16]	|
  0.003 [4]	|
  0.003 [2]	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0001 secs
  50% in 0.0002 secs
  75% in 0.0004 secs
  90% in 0.0008 secs
  95% in 0.0009 secs
  99% in 0.0015 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0000 secs, 0.0028 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0007 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0019 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0018 secs

Status code distribution:
  [200]	10000 responses



