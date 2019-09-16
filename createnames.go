// A litte program to generate some random names. Could be modified to take the number of names desired.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {


    firstnames := [40]string{
        "Adam",
        "Ann",
        "Beth",
        "Bhupinder",
        "Charles",
        "Clare",
        "Dennis",
        "Donna",
        "Elvis",
        "Eve",
        "Faith",
        "Fred",
        "Georgina",
        "Guresh",
        "Henry",
        "Heather",
        "Ike",
        "Isobel",
        "John",
        "Jane",
        "Kevin",
        "Karen",
        "Li",
        "Laura",
        "Mark",
        "Mary",
        "Norm",
        "Nancy",
        "Owen",
        "Oda",
        "Paul",
        "Pam",
        "Quentin",
        "Quinn",
        "Ralph",
        "Rose",
        "Steve",
        "Sally",
        "Tom",
        "Tamara",
    }

    lastnames := [23]string{
        "Adams",
        "Brown",
        "Clarkson",
        "Ditson",
        "Eby",
        "Frank",
        "George",
        "Howard",
        "Ivy",
        "Jacobs",
        "Kemp",
        "Levert",
        "Markson",
        "Nobbert",
        "Owens",
        "Peters",
        "Quince",
        "Roberts",
        "Stevens",
        "Tamber",
        "Unh",
        "Vermenger",
        "Wagner",
    }

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    s2 := rand.NewSource(time.Now().UnixNano())
    r2 := rand.New(s2)

    // Create 20 random names
    for i:= 0; i < 20; i++ {
        first := firstnames[r1.Intn(len(firstnames))]
        last := lastnames[r2.Intn(len(lastnames))]

        fmt.Println(first, last)
    }
}