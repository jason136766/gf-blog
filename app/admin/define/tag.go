package define

type TagStore struct {
	TagName string `p:"tag_name" v:"required|length:2,10|not-exists:tags,tag_name"`
}
