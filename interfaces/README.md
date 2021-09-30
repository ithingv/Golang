# Interface

Every value has a type
Every function has to specify the type of its arguments

Every function we ever write has to be rewritten to accomodate
different types even if the logic in it is identical?

ex)
func (d deck) shuffle() : can only shuffle a value of type 'deck'
func (s []float64) shuffle() : can only shuffle a value of type '[]float64'
func (s []string) shuffle() : can only shuffle a value of type '[]string'
func (s []int) shuffle() : can only shuffle a value of type '[]int'


Concrete Type
map, struct, int, string, koreanBot

Interface Type
bot

---

1. Interfaces are not generic types
-  Other languages have 'generic' types -go does not
2. Interfaces are 'implicit'
- We don't manually have to say that our custom type satisfies some interface
3. Interfaces are a contract to help us manage types
- Garbage in -> Garbage out. If our custom type's implementation of a function is broken then interfaces won't help us!
4. Interfaces are tough. Step #1 is understanding how to read them
Understand how to read interfaces in the standard lib.
Writing your own interfaces is tough and requires experience
