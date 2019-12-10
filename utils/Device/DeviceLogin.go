package Device

// 设备登录
type DeviceLogin struct {
	ErrorMessage string             `json:"errorMessage"`
	Place        []DeviceLoginPlace `json:"place"` // 考勤机位置结构
	Error        int                `json:"error"`
	SchoolName   string             `json:"schoolName"` // 学校名称
	DeviceId     string             `json:"deviceId"`   // 设备id
	SchoolId     string             `json:"schoolId"`   // 学校id
}

type DeviceLoginPlace struct {
	Tid   string                  `json:"tid"`  // 类别id
	Name  string                  `json:"name"` // 位置名称
	Child []DeviceLoginPlaceChild // 子类
}

type DeviceLoginPlaceChild struct {
	Tid  string `json:"tid"`  // 类别id
	Name string `json:"name"` // 位置名称
}
