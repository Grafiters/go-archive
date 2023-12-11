package meme

import (
	"archive/actions/utils/external"
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type MemeType struct {
	Type string `json:"type" binding:"optional" swag:"query"`
}

//  PublicMeme default implementation.
//	@Summary		Meme endpoint configuaration
//	@Description	This endpoint to get list of meme referances on 9gag.com.
//  @Param params query MemeType false "Meme type parameters"
//  @Tags			Public
//	@Produce		json
//	@Router			/public/meme [get]
func PublicMemeIndex(c buffalo.Context) error {
	params := c.Param("type")
	path := fmt.Sprintf("v1/feed-posts/type/%s", params)
	result, err := external.GetRequest(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.Render(http.StatusOK, r.JSON(result))
}
