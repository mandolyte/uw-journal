
## 2023-04-05

**Not working too well yet**

I cloned libsql and did the quick start to build it. Erred out with this message:
```
gcc  -g -O2 -o mksourceid /home/cecil/Projects/github.com/libsql/libsql/tool/mksourceid.c
tclsh /home/cecil/Projects/github.com/libsql/libsql/tool/mksqlite3h.tcl /home/cecil/Projects/github.com/libsql/libsql >sqlite3.h
/bin/sh: 1: tclsh: not found
make: *** [Makefile:1149: sqlite3.h] Error 127
$ 
```

To install it:
` sudo apt-get install tcl`

Got past that but then had the compiler was throwing warnings and errors for what seemed like every single line of source code. Finally did a control-C to stop it.

They also have releases... I'll try that...

https://github.com/libsql/libsql/releases/tag/libsql-0.2.1

Also failed:
```
$ tar xvf libsql-0.2.1.tar.gz 
sqlite3
libsql
.libs/
.libs/liblibsql-3.42.0.so.0
.libs/tclsqlite.o
.libs/liblibsql-3.42.0.so.0.8.6
.libs/libsqlite3-3.42.0.so.0.8.6
.libs/liblibsql.a
.libs/libsqlite3.so
.libs/libtclsqlite3.a
.libs/liblibsql.lai
.libs/libsqlite3.a
.libs/sqlite3.o
.libs/libtclsqlite3.so
.libs/liblibsql.so
.libs/libtclsqlite3.lai
.libs/libsqlite3-3.42.0.so.0
.libs/libsqlite3.la
.libs/liblibsql.la
.libs/libsqlite3.lai
.libs/libtclsqlite3.la
$ ./sqlite3
./sqlite3: error while loading shared libraries: libedit.so.0: cannot open shared object file: No such file or directory
$ ./libsql 
./libsql: error while loading shared libraries: libedit.so.0: cannot open shared object file: No such file or directory
$ 
```


