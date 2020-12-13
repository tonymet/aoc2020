## Refreshers & Lessons
* define unique error with ErrorOOB = errors.New("oob") -- then test with if err == ErrorOOB
* fmt.scanf is powerful, but can't scan multiple words into one string 
* fmt.scanXXX returns io.EOF 
* slice[start:finish] references start to (finish - 1)
* small unit suites payoff
* be faster to consider memoization / dynamic approach if problem is n+1 style solution

## Mistakes
* d8 -- visit count not cleared --  passing struct copies struct but passes elements by reference 
* bag of bags pt 2 -- jumping into recursion too quickly without clarifying method. 
* combinatorics / for d10