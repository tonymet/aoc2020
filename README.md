## Refreshers
* define unique error with ErrorOOB = errors.New("oob") -- then test with if err == ErrorOOB
* fmt.scanf is powerful, but can't scan multiple words into one string 
* fmt.scanXXX returns io.EOF 
* slice[start:finish] references start to (finish - 1)

## Mistakes
* d8 -- visit count not cleared --  passing struct copies struct but passes elements by reference 