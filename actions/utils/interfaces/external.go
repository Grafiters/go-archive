package interfaces

type Meta struct {
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
	Sid       string `json:"sid"`
}

type Data struct {
	NextCursor string      `json:"nextCursor"`
	Posts      []PostsData `json:"posts"`
	Tags       []TagsPost  `json:"tags"`
}

type TagsPost struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

type PostsData struct {
	Id            string    `json:"id"`
	Url           string    `json:"url"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	UpVoteCount   uint16    `json:"upVoteCount"`
	DownVoteCount uint16    `json:"downVoteCount"`
	Interest      []string  `json:"interests"`
	Image         ImageList `json:"images"`
	Tags          []TagsPost
}

type Images struct {
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
	Url    string `json:"url"`
}

type ImageList struct {
	Image700         Images `json:"image700"`
	Image460         Images `json:"image460"`
	ImageFbThumbnail Images `json:"imageFbThumbnail"`
	Image460Sv       Images `json:"image460sv"`
}
