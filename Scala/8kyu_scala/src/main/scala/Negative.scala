/*
* Return Negative
* https://www.codewars.com/kata/55685cd7ad70877c23000102
* In this simple assignment you are given a number and have to make it negative.
*  But maybe the number is already negative?
* Example:
* makeNegative(1); // return -1
* makeNegative(-5); // return -5
* makeNegative(0); // return 0
* Notes:
* The number can be negative already, in which case no change is required.
* Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.
* */
object Negative {

  def makeNegative(n: Int): Int = if(n > 0) -n else n

  def main(args:Array[String]) :Unit = {
    println(makeNegative(1)); // return -1
    println(makeNegative(-5)); // return -5
    println(makeNegative(0)); // return 0
  }
}