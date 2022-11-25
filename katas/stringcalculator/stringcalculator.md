# TDD Code Kata

## TDD Kata - String Calculator
Developed by Roy Osherove - http://osherove.com/tdd-kata-1

The goal of this exercise is to see how far you can get in 30 minutes. Write a test before implementing any code. Remember to solve things as simply as possible so that you force yourself to write tests you did not think about. Remember to refactor after each passing test

### Step 1: Create a simple String calculator with a single method: int Add(string numbers)
The method can take 0, 1 or 2 numbers, and will return their sum. For example:

* Add(“”) should return 0
* Add(“2112”) should return 2112
* Add(“2,3”) should return 5
