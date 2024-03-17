package snowflake

import (
	"time"
	"wild/configs"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func InitSnowFlake() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", configs.Conf.SnowFlakeConfig.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(configs.Conf.SnowFlakeConfig.MachineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2022-05-03", 1); err != nil {
//		fmt.Printf("init failed, err:%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
