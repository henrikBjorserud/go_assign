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

    three := make(map[string]string)
    three["JK"] = ""
    three["MB_AB_AG"] = ""
    three["KU_GL_FB"] = ""

    four := make(map[string]string)
    four["JK"] = ""
    four["MB_AB_AG"] = ""
    four["KU"] = ""
    four["GL_FB"] = ""

    fmt.Println("Hur många personal vill du fördela, 3 eller 4?")  // Print simple text on screen
    var numPers int
    fmt.Scanln(&numPers)

    names := name_scan(numPers)

    shuffle_names(names)

    if len(names) == 3 {
        FM := newThree(names)
        fmt.Println(FM)

    }
}
