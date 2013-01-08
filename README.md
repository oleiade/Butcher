Butcher
=======

Slice fat files into shinny slim ones in just a command

```
butcher -input bigfatfile -output dir -sample linecountbysubfile
```

It's goal is to help splitting very large files into smaller ones based on line count, or size, or the number of
output files you'd want to see generated. Supporting input and ouput compression (tar, gzip, bz2, zip, lz4), and 
delivering a live progress estimation.

Butcher is written in Go, and intend to be efficient, portable, and more importantly : extandable.

In the pipe (still under development):

*quite everything at the moment*

* Compression support
* Split based on size
* Split by output files count
* Tests
* Clean packaging and extendability
