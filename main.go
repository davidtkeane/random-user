package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "os"
    "regexp"
    "time"

    "github.com/common-nighthawk/go-figure"
    "github.com/fatih/color"
)

const APIURL string = "https://randomuser.me/api"

var client *http.Client

type RandomUsers struct {
    Results []struct {
        Gender string `json:"gender"`
        Name   struct {
            Title string `json:"title"`
            First string `json:"first"`
            Last  string `json:"last"`
        } `json:"name"`
        Location struct {
            Street struct {
                Number int    `json:"number"`
                Name   string `json:"name"`
            } `json:"street"`
            City        string `json:"city"`
            State       string `json:"state"`
            Country     string `json:"country"`
            Postcode    int    `json:"postcode"`
            Coordinates struct {
                Latitude  string `json:"latitude"`
                Longitude string `json:"longitude"`
            } `json:"coordinates"`
            Timezone struct {
                Offset      string `json:"offset"`
                Description string `json:"description"`
            } `json:"timezone"`
        } `json:"location"`
        Email string `json:"email"`
        Login struct {
            UUID     string `json:"uuid"`
            Username string `json:"username"`
            Password string `json:"password"`
            Salt     string `json:"salt"`
            Md5      string `json:"md5"`
            Sha1     string `json:"sha1"`
            Sha256   string `json:"sha256"`
        } `json:"login"`
        Dob struct {
            Date time.Time `json:"date"`
            Age  int       `json:"age"`
        } `json:"dob"`
        Registered struct {
            Date time.Time `json:"date"`
            Age  int       `json:"age"`
        } `json:"registered"`
        Phone string `json:"phone"`
        Cell  string `json:"cell"`
        ID    struct {
            Name  string `json:"name"`
            Value string `json:"value"`
        } `json:"id"`
        Picture struct {
            Large     string `json:"large"`
            Medium    string `json:"medium"`
            Thumbnail string `json:"thumbnail"`
        } `json:"picture"`
        Nat string `json:"nat"`
    } `json:"results"`
    Info struct {
        Seed    string `json:"seed"`
        Results int    `json:"results"`
        Page    int    `json:"page"`
        Version string `json:"version"`
    } `json:"info"`
}

func GetJson(url string, target interface{}) error {
    resp, err := client.Get(url)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return json.NewDecoder(resp.Body).Decode(target)
}

// stripANSI removes ANSI escape sequences from a string
func stripANSI(input string) string {
    re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return re.ReplaceAllString(input, "")
}

func GetRandomUser(nationality string, gender string, exportToFile bool) {
    apiURL := APIURL
    params := ""
    if nationality != "" {
        params += fmt.Sprintf("nat=%s", nationality)
    }
    if gender != "" {
        if params != "" {
            params += "&"
        }
        params += fmt.Sprintf("gender=%s", gender)
    }

    if params != "" {
        apiURL += "?" + params
    }

    var randomUsers RandomUsers

    err := GetJson(apiURL, &randomUsers)

    if err != nil {
        fmt.Printf("Error getting random user: %s\n", err.Error())
        return
    }

    user := randomUsers.Results[0]

    // Define color functions for terminal output
    titleColor := color.New(color.FgGreen).SprintFunc()
    infoColor := color.New(color.FgBlue).SprintFunc()
    detailColor := color.New(color.FgYellow).SprintFunc()
    passwordColor := color.New(color.FgRed, color.Bold).SprintFunc()

    // ASCII Art Avatar
    avatar := figure.NewFigure(user.Name.First, "thin", true)

    // Prepare output for terminal
    output := ""
    output += fmt.Sprintf("%s\n", titleColor("--- User Information ---"))
    output += fmt.Sprintf("%s\n\n", avatar.String())
    output += fmt.Sprintf("%s %s\n", infoColor("Gender:"), detailColor(user.Gender))
    output += fmt.Sprintf("%s %s %s %s\n\n", infoColor("Name:"), detailColor(user.Name.Title), detailColor(user.Name.First), detailColor(user.Name.Last))
    output += fmt.Sprintf("%s\n\n", titleColor("--- Location ---"))
    output += fmt.Sprintf("%s %d %s\n", infoColor("Street:"), user.Location.Street.Number, detailColor(user.Location.Street.Name))
    output += fmt.Sprintf("%s %s\n", infoColor("City:"), detailColor(user.Location.City))
    output += fmt.Sprintf("%s %s\n", infoColor("State:"), detailColor(user.Location.State))
    output += fmt.Sprintf("%s %s\n", infoColor("Country:"), detailColor(user.Location.Country))
    output += fmt.Sprintf("%s %d\n", infoColor("Postcode:"), detailColor(user.Location.Postcode))
    output += fmt.Sprintf("%s %s %s\n", infoColor("Coordinates:"), detailColor(user.Location.Coordinates.Latitude), detailColor(user.Location.Coordinates.Longitude))
    output += fmt.Sprintf("%s %s %s\n\n", infoColor("Timezone:"), detailColor(user.Location.Timezone.Offset), detailColor(user.Location.Timezone.Description))
    output += fmt.Sprintf("%s\n\n", titleColor("--- Login Information ---"))
    output += fmt.Sprintf("%s %s\n", infoColor("Email:"), detailColor(user.Email))
    output += fmt.Sprintf("%s %s\n", infoColor("Username:"), detailColor(user.Login.Username))
    output += fmt.Sprintf("%s %s\n", infoColor("Password:"), passwordColor(user.Login.Password))
    output += fmt.Sprintf("  Password Strength: Unbreakable (as long as you don't tell anyone!)\n\n")
    output += fmt.Sprintf("%s %s\n", infoColor("UUID:"), detailColor(user.Login.UUID))

    if user.ID.Name != "" && user.ID.Value != "" {
        output += fmt.Sprintf("%s %s %s\n", infoColor("ID:"), detailColor(user.ID.Name), detailColor(user.ID.Value))
    } else {
        output += fmt.Sprintf("%s %s\n", infoColor("ID:"), detailColor("Not available"))
    }

    output += fmt.Sprintf("\n%s\n\n", titleColor("--- Login Hashes ---"))
    output += fmt.Sprintf("%s %s\n", infoColor("Salt:"), detailColor(user.Login.Salt))
    output += fmt.Sprintf("%s %s\n", infoColor("MD5:"), detailColor(user.Login.Md5))
    output += fmt.Sprintf("%s %s\n", infoColor("SHA1:"), detailColor(user.Login.Sha1))
    output += fmt.Sprintf("%s %s\n\n", infoColor("SHA256:"), detailColor(user.Login.Sha256))
    output += fmt.Sprintf("%s\n\n", titleColor("--- Personal Details ---"))
    output += fmt.Sprintf("%s %s (Age: %d)\n", infoColor("Date of Birth:"), detailColor(user.Dob.Date.Format(time.RFC3339)), user.Dob.Age)
    output += fmt.Sprintf("%s %s (Age: %d)\n", infoColor("Registered:"), detailColor(user.Registered.Date.Format(time.RFC3339)), user.Registered.Age)
    output += fmt.Sprintf("%s %s\n", infoColor("Phone:"), detailColor(user.Phone))
    output += fmt.Sprintf("%s %s\n\n", infoColor("Cell:"), detailColor(user.Cell))
    output += fmt.Sprintf("%s\n\n", titleColor("--- Picture ---"))
    output += fmt.Sprintf("%s %s\n", infoColor("Large:"), detailColor(user.Picture.Large))
    output += fmt.Sprintf("%s %s\n", infoColor("Medium:"), detailColor(user.Picture.Medium))
    output += fmt.Sprintf("%s %s\n\n", infoColor("Thumbnail:"), detailColor(user.Picture.Thumbnail))
    output += fmt.Sprintf("%s %s\n", infoColor("Nationality:"), detailColor(user.Nat))

    // Display output in terminal
    fmt.Println(output)

    // Prepare plain text output for file
    plainOutput := stripANSI(output) // Remove ANSI color codes

    // Save to file if requested
    if exportToFile {
        file, err := os.Create("random_user.txt")
        if err != nil {
            fmt.Printf("Error creating file: %s\n", err.Error())
            return
        }
        defer file.Close()

        file.WriteString(plainOutput)
        fmt.Println("Output exported to random_user.txt")
    }
}

func main() {
    client = &http.Client{Timeout: 10 * time.Second}

    nationality := flag.String("nat", "", "Nationality to filter by (e.g., us, gb, fr)")
    gender := flag.String("gender", "", "Gender to filter by (male, female)")
    exportToFile := flag.Bool("text", false, "Export output to a text file")
    help := flag.Bool("help", false, "Show help message")

    flag.Parse()

    if *help {
        fmt.Println("Usage: go run randomusercolor.go [options]")
        fmt.Println("Usage: go run randomusercolor.go -nat ie -gender male --text")
        flag.PrintDefaults()
        os.Exit(0)
    }

    GetRandomUser(*nationality, *gender, *exportToFile)
}