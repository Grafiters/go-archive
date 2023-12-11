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
	Id            string `json:"id"`
	Url           string `json:"url"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	UpVoteCount   uint16 `json:"upVoteCount"`
	DownVoteCount uint16 `json:"downVoteCount"`
	Tags          []TagsPost
}
