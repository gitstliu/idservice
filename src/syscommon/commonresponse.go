package syscommon

type CommonPageResponse struct {
	CommonResponse
	PageNum  int
	PageSize int
	Sort     string
}

type CommonListResponse struct {
	CommonResponse
}

type CommonResponse struct {
	Code    int
	Message string
	Result  interface{}
}
