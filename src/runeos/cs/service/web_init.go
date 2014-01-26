package service

var (
	ws map[string]interface{}
	wr map[string]interface{}
)

func init() {
	ws = make(map[string]interface{})
	wr = make(map[string]interface{})
}
