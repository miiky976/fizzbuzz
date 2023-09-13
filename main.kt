fun main() {
    // create a list of any type
    val list = mutableListOf<String>()
    for (i in 1..100){
        var msg = ""
        if (i%3 == 0) msg += "Fizz"
        if (i%5 == 0) msg += "Buzz"
        if (msg == "") msg = i.toString()
        list.add(msg)
    }
    println(list)
}