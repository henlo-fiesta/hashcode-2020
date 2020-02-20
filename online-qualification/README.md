# Google Books ?

## Henlo I am live?WTF

### WTF?

## Problems

Book:
  ID
  Score

Library:
  Set of books
  Time it take to scan
  Number of books that can be scanned each day

Time:
  Days

## Algorithm

- Sturcts of Books and Lib
- Create a scoring mechanism
  - Boundary Optimisation - Remove those that were unable do within remaining time
  - Rescoring every iteration
  - Scoring based on points then time
- Sort

## FAQ

https://codingcompetitions.withgoogle.com/hashcode/faq

## Solution

![img][img/pic.png]

**Greedy**

signup days 2 

2 books per day
book per days


libs book = 0,1,2,3,4
Sort by book score
Sort by book score
3  , 9 ,



lib 0 5days = 10 score
lib{
  sign up = 5days
  bookperdays = 10 books
}

lib 1 7days = 6 score

---------------------------------------------------------
10 days
books[1,2,3,4,5]

D->0  1  2  3  4  5  6  7  8  9
L  
1                    s  s any any
2  s  s  s  B4 B3 B1
3           s  s  s  B2 any any any

lib 0 - 2 signup, 1 book
  - B1 = 2
  - B3 = 4
----first iteration----
10 days
book al avail
2 days signup 
1 books perdays
4 days = 6 scores
6/4 = 1.5
----secon iteration----
7 days
B0,B2 left
check books, check deadline
no book left skip
= 0 scores

lib 1 - 3 signup, 1 book
  - B1 = 2
  - B4 = 5
  - B3 = 4
----first iteration----
10 days
book al avail
3 days signup 
2 books perdays
6 days = 11 scores
11/6 = 1.83
Add THIS LIB because highest output scores per day

lib 2 - 4 signup, 2 book
  - B2 = 3
  - B4 = 5
----first iteration----
10 days
book al avail
4 days signup 
2 books perdays
5 days = 8 scores
8/5 = 1.6
----second iteration----
7 days
B0,B2 left
check books, check deadline against signupday
has b2
5 days = 3 scores
Add THIS LIB because highest output scores per day


then third iteration cant do shit be cause it generate 0 scores

----

---DONE----
use 8 days got 14 scores