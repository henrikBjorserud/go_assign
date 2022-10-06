package main

// Import OS and fmt packages
import (
	"fmt"
    "math/rand"
    "time"
)

type three struct {
    Jk string
    MbAbAg string
    KuGlFb string
}

type four struct {
    Jk string
    MbAbAg string
    Ku string
    GlFb string
}

func newThree(names[]string) *three {
    shift := three{}
    shift.Jk = names[0]
    shift.MbAbAg = names[1]
    shift.KuGlFb = names[2]

    fmt.Println(shift)
    return &shift
}

func newFour(names[]string) *four {
    shift := four{}
    shift.Jk = names[0]
    shift.MbAbAg = names[1]
    shift.Ku = names[2]
    shift.GlFb = names[3]

    fmt.Println(shift)
    return &shift
}
// The main function for making a schedule for verkstan
func name_scan(numPers int) []string {

     names := make([]string, numPers)

    for i := 0; i < numPers;{
        fmt.Println(i)
        fmt.Println("Skriv ett namn och tryck enter: ")
        var name string
        fmt.Scanln(&name)
        names[i] = name
        i++
    }
        return names
}

func shuffle_names(names []string) []string {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(names), func(i, j int) {names[i], names[j] = names[j], names[i]})

    return names
}

func main() {

    fmt.Println("Hur många personal vill du fördela, 3 eller 4?")  // Print simple text on screen
    var numPers int
    fmt.Scanln(&numPers)

    names := name_scan(numPers)

    shuffle_names(names)

    if len(names) == 3 {
        FM := newThree(names)
        fmt.Printf("%+v", FM)


    if len(names) == 4 {
        FM := newFour(names)
        fmt.Printf("%+v", FM)


    }
}
