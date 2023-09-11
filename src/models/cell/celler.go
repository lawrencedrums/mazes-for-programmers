package models

type Celler interface {
    Link(Celler, bool)
    Unlink(Celler, bool)
    Links() []Celler
    IsLinked(Celler) bool

    Row() int
    Col() int
    Neighbors() []Celler
}
