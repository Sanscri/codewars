/*
Expression matter
https://www.codewars.com/kata/5ae62fcf252e66d44d00008e
Given three integers a ,b ,c, return the largest number obtained after inserting the following operators and brackets: +, *, ()
In other words , try every combination of a,b,c with [*+()] , and return the Maximum Obtained
* */
function expressionsMatter(a, b, c) {
    return Math.max(
        a + b + c,
        a * b * c,
        a + b * c,
        a * b + c,
        a * (b + c),
        (a + b) * c,
    );
}

console.log(expressionsMatter(1,1,1)) // 3
console.log(expressionsMatter(1,2,3)) // 9
