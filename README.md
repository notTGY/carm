# carm [cat + rm]

Why?

Often when running batch jobs, such as MPI experiments you need
to quickly check outputs of stdout and stderr files. You only need
to read them once. So you just repeat `cat` and `rm` commands on
the same file. With `carm` it will take half the time.

## Usage

```
carm task.o001
```

## Future features

`carm --restore` to restore latest deleted file


## Build

1. `go build -o bin/carm`
2. `dpkg-deb --build . carm.deb`
