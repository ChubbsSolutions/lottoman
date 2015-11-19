package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/mailgun/mailgun-go"
	"github.com/ttacon/chalk"
)

//Define constants
const version string = "0.5.1"
const author string = "Chubbs Solutions"
const email string = "lottoman@chubbs.solutions"
const appName string = "lottoman"
const appDescription string = "Application to get the lucky numbers for the lotto."

//MailgunPublicAPIKey key for the mail service.
var MailgunPublicAPIKey = os.Getenv("MAILGUN_PUBLIC_API_KEY")

//MailgunDomain key for the mail service.
var MailgunDomain = os.Getenv("MAILGUN_DOMAIN")

//Evaluate options on main
func main() {
	t := time.Now()
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
				luckyNumbers := hoosierLottoGet()

				err := displayNumbers(luckyNumbers, nil)
				if err != nil {
					fmt.Println(chalk.Red, "Error displaying numbers.")
					return
				}

				if len(c.Args()) == 1 {
					recipient := c.Args()[0]
					subject := fmt.Sprintf("Hoosier Lotto Numbers for %s", t.Format("Jan 02 2006"))
					err = emailNumbers(luckyNumbers, nil, subject, recipient)
					if err != nil {
						fmt.Println(chalk.Red, "Error emailing numbers.")
						return
					}
				}
			},
		},
		{
			Name:      "powerball",
			ShortName: "pb",
			Usage:     "Get numbers for the powerball",
			Action: func(c *cli.Context) {
				luckyNumbers, pb := powerballGet()

				err := displayNumbers(luckyNumbers, &pb)
				if err != nil {
					fmt.Println(chalk.Red, "Error displaying numbers.")
					return
				}

				if len(c.Args()) == 1 {
					recipient := c.Args()[0]
					subject := fmt.Sprintf("Powerball Numbers for %s", t.Format("Jan 02 2006"))
					err := emailNumbers(luckyNumbers, &pb, subject, recipient)
					if err != nil {
						fmt.Println(chalk.Red, "Error emailing numbers.")
						return
					}
				}

			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(chalk.Red, "Error running app.")
		os.Exit(-1)
	}

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
func hoosierLottoGet() []int {
	i := 0
	var num []int
	for i < 6 {
		temp := generateRandomNumber(1, 48)
		if IsNumGood(num, temp) {
			num = append(num, temp)
			i++
		}
	}
	return num
}

func powerballGet() ([]int, int) {
	i := 0
	var num []int
	var pb int
	for i < 5 {
		temp := generateRandomNumber(1, 69)
		if IsNumGood(num, temp) {
			num = append(num, temp)
			i++
		}
	}
	pb = generateRandomNumber(1, 26)
	return num, pb
}

func displayNumbers(numbers []int, pb *int) error {
	for _, num := range numbers {
		fmt.Print(chalk.Cyan, num, "\n")
	}
	if pb != nil {
		fmt.Print(chalk.Yellow, "\nPowerball: ", *pb, "\n")
	}
	return nil
}

func emailNumbers(numbers []int, pb *int, subject, recipient string) error {

	if MailgunPublicAPIKey == "" || MailgunDomain == "" {
		fmt.Println(chalk.Red, "Please set the MAILGUN_PUBLIC_API_KEY and MAILGUN_DOMAIN variables.")
		return errors.New("Please set the MAILGUN_PUBLIC_API_KEY and MAILGUN_DOMAIN variables.")
	}

	sender := fmt.Sprintf("donotreply@%s", MailgunDomain)
	var body string

	for _, number := range numbers {
		n := strconv.Itoa(number)
		body += n + "\n"
	}

	if pb != nil {
		p := strconv.Itoa(*pb)
		body += "\nPowerball: " + p
	}

	gun := mailgun.NewMailgun(MailgunDomain, MailgunPublicAPIKey, "")
	m := mailgun.NewMessage(sender, subject, body, recipient)

	_, _, err := gun.Send(m)
	if err != nil {
		fmt.Println(chalk.Red, err)
		return errors.New("Error sending.")
	}
	// fmt.Printf("Response ID: %s\n", id)
	// fmt.Printf("Message from server: %s\n", response)

	return nil
}

//generateRandomNumber Get random number given the min and the max
func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	proposed := rand.Intn(max-min) + min
	return proposed
}
