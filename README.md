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

## Runtime complexity
Since there are always two intergers per interval and we need a
constant amount of three comparisons to decide what to do with two 
merge candidates we can ignore those when looking for the runtime
complexity. Also generally speaking we are improving the runtime complexity by
ordering the already merged intervals after every merge attempt.  
For **n** intervals passed to the program, we 
actually have a worst case of **1+2+3+4...+n** merge attempts. This is the little
gaussian formula  
$$ \sum_{k=1}^n = \frac{n(n+1)}{2} = \frac{n^2+n}{2} $$  
which amounts to
**O($$n^2$$)**.  
That being said this only is true for cases in which the intervals
passed to the program are non overlapping and in ascending order.
E.g. [1,2], [2,3], [3,4], [4,5]. In the two best case scenarios,
the intervals are either passed to the program in descending order
e.g. [4,5], [3,4], [2,3], [1,2] or one of the first two intervals
is one that encases the others. In both cases there will always be
just one merge attempt for every following interval, amounting to a best case runtime
complexity of: **O(n)**  
The average runtime complexity will thus be something between **O(n)** and **O($$n^2$$)**.
In scenarios where the worst case is applied more often a possible solution would be to
pre sort the intervals or even order them randomly (this would be beneficial because it would 
be guraranteed to run in O(n)).

## Robustness
To increase the programs robustness it would be easiest to copy the input parsing approach
presented here in the parser.go and apply it on reading the input bit by bit from file
(consuming the files content bit by bit).
The output could then be written to another file (merge attempt after merge attempt).
This way you would always persist the current state to disk and could even resume it,
in case of unexpected program termination. It would of course slow the execution down,
due to the I/O operations.

## Memory consumption
Since the current state of the programm always loads all the intervalls to memory
and in the worst case scenario copies it to another place in memory the memory required
would be **2n** while **n** would amount to the space required for 
_2\*integer*capacity_of_slice_holding_all_the_intervals_. If the changes descibed in the 
former paragraph _Robustness_ would be applied, the memory consumption would actually
amout to **2** since there would always only be the need to have two intervals in memory.