package concat

type (
	Videos []Video
	Video  struct {
		Name       string
		CreateTime int64
	}
)

func (v Videos) Len() int {
	return len(v)
}

func (v Videos) Less(i, j int) bool {
	return v[i].CreateTime < v[j].CreateTime
}

func (v Videos) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
