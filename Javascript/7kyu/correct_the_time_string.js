/*
* https://www.codewars.com/kata/57873ab5e55533a2890000c7
* Correct the time-string
* A new task for you!
*
* You have to create a method, that corrects a given time string.
* There was a problem in addition, so many of the time strings are broken.
* Time is formatted using the 24-hour clock, so from 00:00:00 to 23:59:59.
*
* */

function timeCorrect(timestring) {
    if (!timestring)
        return timestring;
    if (!/(\d{2}):(\d{2}):(\d{2})/g.test(timestring))
        return null;
    return new Date(0, 0, 0, ...timestring.split(':'))
        .toLocaleTimeString('en', { hour12: false })
}

console.log(timeCorrect(null))
console.log(timeCorrect(""))
console.log(timeCorrect("001122"))
console.log(timeCorrect("09:10:01"))
