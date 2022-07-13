package domain

const carSound string = "wrrrum"

type Car struct {
	Id                 int    `json:"id"`
	Model              string `json:"model"`
	RegistrationPlates string `json:"registration_plates"`
}

func (car *Car) DriveSound() string {
	return carSound
}
