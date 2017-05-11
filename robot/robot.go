package robot

import (
	"fmt"
	"io"
	"os"
)

type IRobot interface {
	Name() string
	Call(method, params string) ([]byte, error)
	Response(w io.Writer, resp []byte) error
}

//robot.RobotServer包含一个IRobot接口
type RobotClient struct {
	IRobot
}

//新建一个Robot
func NewRobot(r IRobot) *RobotClient {
	return &RobotClient{r}
}

//robot.RobotServer服务方法
func (r *RobotClient) Serve(w io.Writer, method, param string) {
	result, err := r.Call(method, param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s call method:%s,param:%s.Failed.%s\n",
			r.Name(), method, param, err.Error())
	}
	if err := r.Response(w, result); err != nil {
		fmt.Fprintf(os.Stderr, "%s respone Error:%s", r.Name(), err.Error())
	}
}
