package err

// 业务错误定义
// ----------
type Err struct {
	error string
}

// 创建错误
func NewErr(error string) *Err {
	return &Err{
		error: error,
	}
}

// 获取错误信息
func (e Err) Error() string {
	return e.error
}
