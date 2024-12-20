package form

const DEFAULT_PAGE = 1
const DEFAULT_SIZE = 10

type PageForm struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Status int    `json:"status"`
	TaskId string `json:"taskId"`
}

func (form PageForm) GetPage() int {
	if form.Page <= 0 {
		return DEFAULT_PAGE
	} else {
		return form.Page
	}
}

func (form PageForm) GetSize() int {
	if form.Size <= 0 {
		return DEFAULT_SIZE
	} else {
		return form.Size
	}
}

func (form PageForm) PageOffset() int {
	return (form.GetPage() - 1) * form.GetSize()
}
