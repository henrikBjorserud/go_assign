package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
    "math/rand"
    "time"
)

type Worker struct {
    Name string
}

type Three struct {
    Jk Worker
    MbAbAg Worker
    KuGlFb Worker
}

type Four struct {
    Jk Worker
    MbAbAg Worker
    Ku Worker
    GlFb Worker
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

    fmt.Println(names)
    shuffle_names(names)
    fmt.Println(names)


    fmt.Println(os.Getenv("USER"), ", Let's be friends!") // Read Linux $USER environment variable 
}
