package pdf

const (
	UPLOADS_BASE_URL = "/uploads"
	UPLOADS_BASE_DIR = "./uploads"
)

type PDF struct {
}

func New() *PDF {
	return &PDF{}
}
