package model

import "time"

//operation es el modelo para cada operacion
type Operation struct {
	Id         int       `json:"id"`
	Time       time.Time `json:"time"`
	State      string    `json:"state"`
	Origin_acc int       `json:"origin_acc"`
	Amount     float64   `json:"amount"`
	Dest_acc   int       `json:"dest_acc"`
}

//operations es donde se almacenaran varias operaciones
type operations []Operation
