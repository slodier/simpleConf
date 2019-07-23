package conf

// 配置文件结构体
type Table struct {
	Server			Server			`json:"server"`
	Database		Database		`json:"mysql"`
	Image			Image			`json:"image"`
}

// 服务
type Server struct {
	Port			int				`json:"port"`
	LogModel		bool			`json:"log_model"`
}

// database
type Database struct{
	Host			string			`json:"host"`
	Port			int				`json:"port"`
	User			string			`json:"user"`
	Pw				string			`json:"pw"`
}

type Image struct {
	
}

func NewTab() *Table {
	return &Table{}
}