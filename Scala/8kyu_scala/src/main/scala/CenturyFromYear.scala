/*
* Century From Year
* https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097
*
* Introduction
* The first century spans from the year 1 up to and including the year 100,
*  The second - from the year 101 up to and including the year 200, etc.
*
* Task :
* Given a year, return the century it is in.
*
* The first century spans from the year 1 up to and including the year 100,
* the second - from the year 101 up to and including the year 200, etc.
*
* Input , Output Examples ::
* Let's see some examples:
*
* centuryFromYear(1705)  returns (18)
* centuryFromYear(1900)  returns (19)
* centuryFromYear(1601)  returns (17)
* centuryFromYear(2000)  returns (20)
* Hope you enjoy it .. Awaiting for Best Practice Codes hahaha ..
* Enjoy Learning !!!
* */
object CenturyFromYear {
  def centuryFromYear(year: Int): Int = {
    var century = year / 100
    if(year % 100 > 0){
       century += 1
    }
    century
  }
  def main(args:Array[String]) :Unit = {
    println(centuryFromYear(1705)) // 18
    println(centuryFromYear(1900)) // 19
    println(centuryFromYear(1601)) // 17
    println(centuryFromYear(2000)) // 20
  }
}
