package wheel

var wheelFactory = make([]Wheel, 10)

type ModeType int

const (
	EXEC_ONE ModeType = iota //执行一次
	EXEC_FOR                 //循环执行
)

type RunType int

const (
	QUEUE RunType = iota //顺序执行
	PARA                 //并行执行
)

var AllWheel []Wheel

func init() {
	AllWheel = make([]Wheel, 0)
}

type Wheel struct {
	Tasks         []Task
	WheelExecTime int64 //轮子开始执行时间
	WheelEndTime  int64 //轮子结束执行时间
	Mode          ModeType
}

type Task struct {
	Num    int               //任务编号（递增）
	Url    string            //执行的链接
	Params map[string]string //每次执行的的参数
	//TaskExecTime int64           //任务执行时间
	//TaskEndTime  int64           //任务结束执行时间
	Mode       ModeType
	ModeSecond int //type1 循环时间
	//Strategy   RunType //执行顺序 1:并发执行 2:顺序执行
}
