package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

type Ticket struct {
    Number      int    `json:"number"`
    Description string `json:"description"`
    State       string `json:"state"`
}

func ReadStatus(p martini.Params) Ticket {

    ticket := Ticket{
        Number:      645,
        Description: "A dummy customer ticket " + p["id"],
        State:       "resolved",
    }
    return ticket
}


// $curl -i -H "Accept: application/json" -H "Content-Type:application/json" -X GET http://localhost:3000/status/id:12345
func main() {

    m := martini.Classic()
    m.Use(render.Renderer())

    m.Get("/status/:id", func(r render.Render, params martini.Params) { r.JSON(200, ReadStatus(params)) })

    m.Run()

}

