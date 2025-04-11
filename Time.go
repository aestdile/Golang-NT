


// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     a := time.Now()

//     b := time.Date(2024, time.April, 16, 10, 30, 0, 0, time.UTC)

//     d := a.Sub(b)

//     days := int(d.Hours() / 24)
//     hours := int(d.Hours() % 24) 
//     minutes := int(d.Minutes()) % 60
//     seconds := int(d.Seconds()) % 60

//     fmt.Printf("Ko'rish uchun: %d kun, %d soat, %d daqiqa, %d soniya\n", days, hours, minutes, seconds)
// }


//---------------------------------------------------------------------



package main

import (
    "fmt"
    "time"
)

type Event struct {
    Name      string    
    Date      time.Time 
    Location  string    
    Duration  time.Duration 
}

type User struct {
    Name    string   
    Events  []Event  
}

func (u *User) AddEvent(name string, date time.Time, location string, duration time.Duration) {
    event := Event{
        Name:     name,
        Date:     date,
        Location: location,
        Duration: duration,
    }
    u.Events = append(u.Events, event)
}

func (u *User) RemoveEvent(eventName string) {
    var updatedEvents []Event
    for _, event := range u.Events {
        if event.Name != eventName {
            updatedEvents = append(updatedEvents, event)
        }
    }
    u.Events = updatedEvents
}

func (u *User) PrintEvents() {
    fmt.Printf("%s ning tadbirlari:\n", u.Name)
    for _, event := range u.Events {
        fmt.Printf("- %s, sanasi: %s, joyi: %s, davomiyligi: %s\n", event.Name, event.Date.Format("2006-Jan-02 15:04"), event.Location, event.Duration)
    }
}

func main() {
    Alice := User{Name: "Alice"}

    Alice.AddEvent("Tadbir 1", time.Now(), "Ortiqcha vaqt: ", 2*time.Hour)
    Alice.AddEvent("Tadbir 2", time.Now().Add(24*time.Hour), "Mehmonxona", 1*time.Hour)
    Alice.AddEvent("Dars 1", time.Now().Add(48*time.Hour), "Universitet", 3*time.Hour)

    Alice.PrintEvents()

    Alice.RemoveEvent("Tadbir 1")

    Alice.PrintEvents()
}


