package main

// Import random, time (to improve pseudo-random) and fmt packages
import (
	"fmt"
    "math/rand"
    "time"
)
// The struct that is used to create a schedule for three worrkers
type three struct {
    Jk string
    MbAbAg string
    KuGlFb string
}
// The struct that is used to create a schedule for four workers
type four struct {
    Jk string
    MbAbAg string
    Ku string
    GlFb string
}
// Rotate names
func rotate_names(names []string) []string {
    var newdata []string
    first := names[:1]
    newdata = append(newdata,names[1:]...)
    newdata = append(newdata,first...)
    return newdata
}

// Assign names to three empty slots
func newThree(names[]string) *three {
    shift := three{}
    shift.Jk = names[0]
    shift.MbAbAg = names[1]
    shift.KuGlFb = names[2]

    return &shift
}
// Assign names to four empty slots
func newFour(names[]string) *four {
    shift := four{}
    shift.Jk = names[0]
    shift.MbAbAg = names[1]
    shift.Ku = names[2]
    shift.GlFb = names[3]

    return &shift
}

// The main function for making a schedule for verkstan
func name_scan(numPers int) []string {

     names := make([]string, numPers)

    for i := 0; i < numPers;{
        fmt.Println("Skriv ett namn och tryck enter: ")
        var name string
        fmt.Scanln(&name)
        names[i] = name
        i++
    }
        return names
}

// Put names in random order
func shuffle_names(names []string) []string {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(names), func(i, j int) {names[i], names[j] = names[j], names[i]})

    return names
}

func main() {
    // Ask for number of people working
    fmt.Println("Hur många personal vill du fördela, 3 eller 4?")
    var numPers int
    fmt.Scanln(&numPers)
    // Collect names of workers
    names := name_scan(numPers)
    // Randomize names
    shuffle_names(names)
    if len(names) == 3 {
        // Assign three workers before lunch
        FM := newThree(names)
        fmt.Println("Before Lunch:")
        fmt.Printf("%+v", FM)
        // Rotate names and assign after lunch
        rotated_names := rotate_names(names)
        fmt.Println("\nAfter Lunch")
        EM := newThree(rotated_names)
        fmt.Printf("%+v", EM)
        fmt.Println("\n")
    }

    if len(names) == 4 {
        // Assign four workers before lunch
        FM := newFour(names)
        fmt.Println("Before Lunch:")
        fmt.Printf("%+v", FM)
        // Rotate names and assign after lunch
        rotated_names := rotate_names(names)
        EM := newFour(rotated_names)
        fmt.Println("\nAfter Lunch:")
        fmt.Printf("%+v", EM)
        fmt.Println("\n")
    }
}
