package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/vishal21121/movie-stream-app/internals/components"
	"github.com/vishal21121/movie-stream-app/internals/types"
	"github.com/vishal21121/movie-stream-app/internals/utils"
)

func fetchMovies(movieName string) (*types.MovieData, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=en-US&page=1", movieName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_KEY")))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var movieData types.MovieData
	decodeErr := json.NewDecoder(response.Body).Decode(&movieData)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return &movieData, nil
}

func GetMovies(c echo.Context) error {
	movieName := c.QueryParam("movieName")
	fmt.Println("movie", movieName)
	nameSplitted := strings.Split(movieName, " ")
	movieName = strings.Join(nameSplitted, "%20")
	movieDetails, err := fetchMovies(movieName)
	if err != nil {
		return c.JSON(500, map[string]any{
			"statusCode": 500,
			"message":    err.Error(),
		})
	}
	if movieDetails.TotalResults > 0 {
		return utils.Render(c, http.StatusOK, components.MovieList(movieDetails.Results))
	}
	return c.JSON(404, map[string]any{
		"statusCode": 404,
		"message":    "Nothing found",
	})
}
