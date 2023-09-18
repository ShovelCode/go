package main

func main() {
    saySomething()  // No need to import; it's in the same package

    myCar := car{
        make:  "Toyota",
        model: "Camry",
    }
    println(myCar.make, myCar.model)  // No need to import; it's in the same package
}

