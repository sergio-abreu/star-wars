package apparitions

func NewMovie(title string) Movie {
	return Movie{Title: title}
}

type Movie struct {
	Title string
}
