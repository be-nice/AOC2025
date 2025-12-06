# Advent of Code 2025

These are my GO solutions for this years AoC challenges

## Usage

Create data folder inside the Day_x folders  
and place your input files there labeled  
"part1.txt", "part2.txt", "test_part1.txt", "test_part2.txt"

**Input files are not part of the repository,  
since they are unique and because AoC creator  
has asked not to include them (don´t ask me why, there isn´t
a reasonable nor sensible answer for that**

### Project structure

AOC2025  
├── config.go  
├── Day_1  
│   ├── data  
│   │   ├── part1.txt  
│   │   ├── part2.txt  
│   │   ├── test_part1.txt  
│   │   └── test_part2.txt  
│   ├── part1.go  
│   ├── part2.go  
│   └── type.go  
├── go.mod  
├── main.go  
├── README.md  
├── Types  
│   └── types.go  
└── Utils  
 └── read_input.go

### Running

```go
go run . <n> <s>
```

**Where**  
**n** is the **day** number for challenge **1-12**  
**s** is the **part** identifier **a | b | at | bt**

a - run part1  
b - run part2  
at - run test part1  
bt - run test part2

#### Builtin help

```go
go run . h | help
or
go run .
```

### Final

All runs are timed, timing starts with call to the days part function.  
That includes reading and parsing the input file.
