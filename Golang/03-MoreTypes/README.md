# Section 3 - More Data Types

Run from root using the following command (UNIX/Linux):
```bash
go run Golang/02-Basics/main.go
```

The output should look as follows:
```
i: 21
j: 54
v: {4 2}
v': {4 5}
v1: {1 2} v2: {1 0} v3: {0 0} p4: &{1 2}
array: [Hello World]
array length: 2
primes: [2 3 5 7 11 13]
primes slice: [3 5 7]

names array: [John Paul George Ringo]
b slice: [John Paul]
c slice: [Paul George]

after modification...
names array: [John XXX George Ringo]
b slice: [John XXX]
c slice: [XXX George]

slice literal q: [2 3 5 7 11 13]
slice literal r: [true false true true false true]

d: [0 0 0 0 0 0 0 0 0 0]
d1: [0 0 0 0 0 0 0 0 0 0]
d2: [0 0 0 0 0 0 0 0 0 0]
d3: [0 0 0 0 0 0 0 0 0 0]
d4: [0 0 0 0 0 0 0 0 0 0]

len = 6, cap = 6 [2 3 5 7 11 13]
len = 0, cap = 6 []
len = 4, cap = 6 [2 3 5 7]
len = 2, cap = 4 [5 7]

nilSlice: [] 0 0

e: [0 0 0 0 0] 5 5
f: [] 0 5
f: [0 0 0 0 0] 5 5
f: [0 0 0 0] 4 4 

X _ X
O _ X
_ _ O

len = 0, cap = 0 []
len = 1, cap = 1 [0]
len = 5, cap = 6 [0 1 2 3 4]

2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128

1
2
4
8
16
32
64
128
256
512
[0 1 4 9 16 25 36 49]
[1 2 5 10 17 26 37 50]
[4 5 8 13 20 29 40 53]
[9 10 13 18 25 34 45 58]
[16 17 20 25 32 41 52 65]
[25 26 29 34 41 50 61 74]
[36 37 40 45 52 61 72 85]
[49 50 53 58 65 74 85 98]

map: map[Bell Labs:{40.68433 -74.39967}]
map literal: map[Apple:{37.33182 -122.03118} Google:{37.42202 -122.08408}]
map literal 2: map[Chicago:{41.87811 -87.6298} Houston:{29.76043 -95.3698} Los Angeles:{34.05223 -118.24368} New York City:{40.71278 -74.00594} Philadelphia:{39.95233 -75.16379} Pittsburgh:{40.44062 -79.99589} San Francisco:{37.77493 -122.41942} Washington D.C.:{38.90719 -77.03687}]

The value: 42
The value: 48
The value: 0
The value: 0 Present? false 

13
5
81 

0 -> pos: 0 neg: 0
1 -> pos: 1 neg: -1
2 -> pos: 3 neg: -3
3 -> pos: 6 neg: -6
4 -> pos: 10 neg: -10
5 -> pos: 15 neg: -15
6 -> pos: 21 neg: -21
7 -> pos: 28 neg: -28
8 -> pos: 36 neg: -36
9 -> pos: 45 neg: -45

1
1
2
3
5
8
13
21
34
55
```