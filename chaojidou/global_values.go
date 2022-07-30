package chaojidou

var Follwers []string

var Leader string

var DefaultCaptain ChaoJiDou

func init() {
	DefaultCaptain, _ = Build(CLIENT_TYPE_OFFICIAL)
}
