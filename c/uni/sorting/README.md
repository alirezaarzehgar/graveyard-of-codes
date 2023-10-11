# Practice bubble sort and insertion sort

# How to run
You can build this project using GNU make.

```bash
$ make
```

After compiling source code you can run it with following command:
```bash
$ ./sorting
```

# Help
You can use options to change sorting algorithm. -b for bubble sort and -i for insertion sort.
By default program will use bubble sort.

```bash
$ ./sorting -h
usage: sorting [bih]
```

*Note*: You should press Ctrl+D for ending input data. This key will send signal to kernel
and close stdin for current process.

# Example output
```plaintext
$ ./sorting -b
Sort with bubble sort algorithm.
3 2 56 -3 4 3
-3 2 3 3 4 56 
```

# Tests
This project have a small system test with bats(Bash Automation Tests) system. You can run following command to run tests:
```bash
$ make test
```

That is right! We needn't this complexity for showing two simple sorting algorithms. But doing university exercises is not a real challenge. Actually it's not a challenge. I want to enjoy using some old concepts that I knew and dreaming with my nostalgias with C and UNIX concepts.
