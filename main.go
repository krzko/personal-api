package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hako/durafmt"
)

func main() {
	// export GO111MODULE=on
	port := "127.0.0.1:31337"

	donations := Donations{
		Currencies: []Project{
			{
				Name:    "Bitcoin",
				Address: "175w9BSN6TjWHbVQKwUWkaZh13j8YX4zgA",
			},
			{
				Name:    "Ethereum",
				Address: "0x63395cCD3C9889B19909813607b07c7a526a250E",
			},
			{
				Name:    "Litecoin",
				Address: "ltc1qjyn0hxxx06k827yvyusp94l23qg9m6wa736u0k",
			},
		},
	}

	person := Person{
		FirstName: "Kristof",
		LastName:  "Kowalski",
	}

	projects := Projects{
		Projects: []Project{
			{
				Name:    "Cashmere Bot",
				Address: "http://cashmere.cash",
			},
			{
				Name:    "Goat Cash",
				Address: "https://goat.cash",
			},
			{
				Name:    "VolAir",
				Address: "https://volair.io",
			},
		},
	}

	social := Social{
		GitHub:    "https://github.com/krzko",
		Instagram: "https://instagram.com/krzko",
		LinkedIn:  "https://www.linkedin.com/in/kristofkowalski/",
		Strava:    "https://www.strava.com/athletes/2028490",
		Twitter:   "https://twitter.com/krzko",
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"FirstName:":  person.FirstName,
			"LastName":    person.LastName,
			"Uptime":      person.uptime(1975, time.January, 15, 03, 15, 00, 211),
			"Message":     "Made with 🍺.",
			"Social":      social,
			"X-Projects":  projects,
			"X-Donations": donations,
		})
	})

	r.Run(port)
}

func (p *Person) uptime(year int, month time.Month, day, hour, min, sec, nsec int) (u string) {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	birthday := time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
	diff := now.Sub(birthday)
	duration := durafmt.Parse(diff).String()
	return duration
}
