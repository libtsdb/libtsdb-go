# Facebook Gorilla

## Paper

http://www.vldb.org/pvldb/vol8/p1816-teller.pdf

4.1 Time series compression

- timestamps and values are compressed separately using information about previous values
  - but they are put into same byte stream

4.1.1 Compressing time stamps 

- first value is aligned to two hour window
- second value is delta with first value, size is 14 bits because, 14 bits is 16384 seconds, 4.5h
- use a dictionary, the range of dictionary is determined by sample

4.1.2 Compressing values

- first XOR w/ previous value
- variable length encoding
 
## Beringei

https://github.com/facebookarchive/beringei