# Summary
This simple source code recieve two matrixes (A and B) and calculate multiplication of them.

# Build
You can use `Makefile` or compile project manually.

- Build code using `Makefile`:
Source code have two build option `MODE=RELEASE` and `MODE=DEBUG`. By default build system compile with `RELEASE MODE`.
If you compile code using `DEBUG` flag, during program execution you should see some extra information about calculations.
```bash
$ make
```
or
```bash
$ MODE=DEBUG make
```

- Build manually
You can compile project using `gcc` and maybe mingw.
```bash
$ gcc main.c -o main
```
or
```bash
$ gcc main.c -o main -D "DEBUG"
```

# Run
This is example of output:
```plaintext
$ ./main 
Enter n,m of matrix A: 2,3
Enter n,m of matrix B: 3,2
Enter matrix A:
1 2 3
4 5 6
Enter matrix B:
7 8
9 10
11 12
A*B:
58	64	
139	154	
```

If you enable debug mode, you can see following output:
```plaintext
$ MODE=DEBUG make
cc main.c -o main -D "DEBUG"
$ ./main 
Enter n,m of matrix A: 2,3
Enter n,m of matrix B: 3,2
Enter matrix A:
1 2 3
4 5 6
Enter matrix B:
7 8
9 10
11 12
(1*7) + (2*9) + (3*11) = 58
(1*8) + (2*10) + (3*12) = 64
(4*7) + (5*9) + (6*11) = 139
(4*8) + (5*10) + (6*12) = 154
A*B:
58	64	
139	154	
```

# Test application
We prepare minimal test script on `test.sh`. You can run following command to generate same test input and check answer.
```bash
$ ./test.sh | ./main
```

You can check output. After changing code, run this code and ensure changes doesn't break our code.
