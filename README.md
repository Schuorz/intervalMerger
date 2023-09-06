# Interval Merger
The Interval Merge is a simple tool for merging
overlapping intervals. Intervals need to be of the form [a,b]
with a and b being integers.  
Example:  
[1,2], [2,5], [3,4], [4,6], [8,12] => [1,2], [2,6], [8,12]

## Prerequisites
You need to have Go version 1.21 installed. You can than either
build the binary yourself or simply run `make` to build a binary
called `prog`.

## Running the program
After building the program, you can run it by calling the 
program by its build name (`prog` if you ran `make`) and pass
the intervals you want to be merged as one string, like so 
`"[1,2],[3,6],[4,7]"`. Note that whitespaces are ignored and
the only seperator between the intervals allowes is `,` which
can also be omitted entirely.  
A complete valid function call should look somewhat like this:  
`./prog "[1,2],[3,6],[4,7]"`