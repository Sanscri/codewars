/*
https://www.codewars.com/kata/5601409514fc93442500010b
here was a test in your class and you passed it. Congratulations!
But you're an ambitious person. You want to know if you're better than the average student in your class.

You receive an array with your peers' test scores. Now calculate the average and compare your score!

Return True if you're better, else False!
*/

function betterThanAverage(classPoints, yourPoints) {
    classPoints.push(yourPoints);
    let average = classPoints.reduce((partial_sum, a) => partial_sum + a, 0) / classPoints.length;
    return average < yourPoints;
}
