package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/codegangsta/cli"
    "os"
)

//Define constants
const version string = "0.3"
const author string = "Chubbs Solutions"
const email string = "lottoman@chubbs.solutions"
const appName string = "lottoman"
const appDescription string = "Application to get the lucky numbers for the lotto."

//Evaluate options on main
func main() {

app := cli.NewApp()
app.Name = appName
app.Usage = appDescription
app.Email = email
app.Author = author
app.Version = version
app.Commands = []cli.Command{
    {
        Name:      "hoosierlotto",
        ShortName: "hl",
        Usage:     "Get numbers for the Hoosier Lotto",
        Action: func(c *cli.Context) {
            hoosierLottoGet()
        },
    },
    {
        Name:      "powerball",
        ShortName: "pb",
        Usage:     "Get numbers for the powerball",
        Action: func(c *cli.Context) {
            powerballGet()
        },
    },
}

app.Run(os.Args)

}

//IsNumGood Function to make sure the number is unique.
func IsNumGood(slice []int, proposed int) bool {
    for _, ele := range slice {
        if ele == proposed || proposed == 0 {
            return false
        }
    }
    return true
}

//hoosierLottoGet Get numbers for the Hoosier Lotto
func hoosierLottoGet() bool {
    i := 0
    var num []int
    for i < 6 {
        rand.Seed(time.Now().UnixNano())
        temp:=rand.Intn(35)
        if IsNumGood(num,temp) {
            fmt.Print(temp,"\n")
            num=append(num,temp)
            i++
        }
    }
    return true
}

func powerballGet() bool {
    i := 0
    var num []int
    for i < 5 {
        rand.Seed(time.Now().UnixNano())
        temp:=rand.Intn(59)
        if IsNumGood(num,temp) {
             fmt.Print(temp,"\n")
             num=append(num,temp)
             i++
        }
    }
    x := 0
    for x == 0 {
        rand.Seed(time.Now().UnixNano())
        pb := rand.Intn(35)
        if pb != 0 {
            fmt.Println("\nPowerball: ",pb)
            x = 1;
        }                
    }
    return true
}
