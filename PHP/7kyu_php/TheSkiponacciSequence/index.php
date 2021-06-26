/*
https://www.codewars.com/kata/580777ee2e14accd9f000165

Your task is to generate the Fibonacci sequence to n places, with each alternating value as "skip". For example:

"1 skip 2 skip 5 skip 13 skip 34"

Return the result as a string

You can presume that n is always a positive integer between (and including) 1 and 64.
*/

<?php
function skiponacci($n) {
    for($i = 0; $i < $n; $i++) {
        if ($i <= 1) {
            $fibonacci[] = 1;
        } else {
            $fibonacci[] = $fibonacci[$i - 1] + $fibonacci[$i - 2];
        }
    }
    foreach($fibonacci as $key => $val) {
        $fibonacci[$key] = ($key + 1) % 2 === 0 ? 'skip': $val;
    }
    return join(' ',  $fibonacci);
}


