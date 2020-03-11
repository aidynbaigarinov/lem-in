# Lem-In

---

**[alem school](http://alem.school/)** project.

The goal is to make a digital version of ant farm and transfer all given ants from **point A** to **point B** using the most optimal paths.

**[The full instructions.](https://github.com/aidynbaigarinov/lem-in/blob/master/instructions.md)**

### Usage

> You need to have the [Go language](https://golang.org/) installed on your PC.

```
$ go build
$ ./lem-in <map.txt>
```

You can build your own map, and place it in `./maps/` folder.

### Usage example

```
$ go build
$ ./lem-in example00.txt

4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1

 497.156µs

```

### Built With

[Go programming language](https://golang.org/)
