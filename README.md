# golang

Repository to practice things about a golang.
We will write different programs and slowly become expert.

### November 7, 2020
--------------------

Command Line Arguments:
- Command-line arguments are available to a program in a variable named Args that is part of 
  the os package;
- The variable os.Args is a slice of strings.
- For now, think of a slice as a dynamically sized array where individual elements can be accessed 
  as s[i] and a contiguous subsequence as s[m:n]. The number of elements is given by len(s).
- all indexing in Go half-open intervals that include the first index but exclude the last, 
  because it simplifies logic

Variable Declaration:
 - A variable can be initialized as part of its declaration.
- If it is not explicitly initialized, it is implicitly initialized to the zero value for its type, 
  which is 0 for numeric types and the empty string "" for strings.
- the + operator concatenates the values
- s += sep + os.Args[i]
  is an assignment statement that concatenates the old value of s with sep and os.Args[i] and
  assigns it back to s; it is e quivalent to 
  s = s + sep + os.Args[i]
- The := symbol is part of a short variable declaration, a statement that declares one or more 
  variables and gives them appropriate types based on the initi alizer values;
- The increment statement i++ adds 1 to i; i += 1 or i = i+1
  In go they are postfix only 
