/*

https://www.codewars.com/kata/55f73be6e12baaa5900000d4

Messi goals function
Messi is a soccer player with goals in three leagues:
LaLiga
Copa del Rey
Champions
Complete the function to return his total number of goals in all three leagues.
Note: the input will always be valid.
For example:
5, 10, 2  -->  17 */

<?php
function goals (int $laLigaGoals, int $copaDelReyGoals, int $championsLeagueGoals) : int {
    return $laLigaGoals + $copaDelReyGoals + $championsLeagueGoals;
}
echo goals(0,0,0); // 0
echo goals(43, 10, 5); // 58