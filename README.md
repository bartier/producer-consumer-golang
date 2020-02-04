# producer-consumer

## Steps to run
```
$ wget https://github.com/bartier/producer-consumer-golang/releases/download/v1.0/producer_consumer_amd64_x86_64 && 
$ chmod +x producer_consumer_amd64_x86_64
$ ./producer_consumer_amd64_x86_64
```

## Log example

In this log there was 6 producers (goroutines) and 1 consumer (another goroutine)

```
time="2020-02-03T22:03:58-03:00" level=info msg="consumer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="producer()"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 1 put resource '887'. Buffer now: [887 0 0 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Consumer ID 1 get resource '887'. Buffer now: [0 0 0 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 1 put resource '211'. Buffer now: [211 0 0 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 4 put resource '728'. Buffer now: [728 0 0 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 6 put resource '300'. Buffer now: [728 300 0 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 2 put resource '511'. Buffer now: [728 300 511 0 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 1 put resource '106'. Buffer now: [728 300 511 106 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 3 put resource '540'. Buffer now: [728 300 511 106 540]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 5 buffer full  '318'. Buffer now: [728 300 511 106 540]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Consumer ID 1 get resource '211'. Buffer now: [728 300 511 106 540]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 4 put resource '947'. Buffer now: [728 300 511 106 947]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 5 buffer full  '528'. Buffer now: [728 300 511 106 947]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Consumer ID 1 get resource '540'. Buffer now: [728 300 511 106 947]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Consumer ID 1 get resource '947'. Buffer now: [728 300 511 106 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 6 put resource '541'. Buffer now: [728 300 511 541 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Consumer ID 1 get resource '106'. Buffer now: [728 300 511 541 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 2 put resource '831'. Buffer now: [728 300 511 831 0]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 5 put resource '888'. Buffer now: [728 300 511 831 888]"
time="2020-02-03T22:03:58-03:00" level=info msg="main(): Producer ID 1 buffer full  '737'. Buffer now: [728 300 511 831 888]"
```

## More about producer-consumer problem
- [Wikipedia](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem)
