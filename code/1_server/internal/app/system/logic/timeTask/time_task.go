/*
* @desc:定时任务栈
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/13 17:41
 */

package timeTask

import (
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.RegisterTaskList(New())
}

func New() *sTaskList {
	return &sTaskList{
		mu: gmutex.New(),
	}
}

type sTaskList struct {
	taskList []*model.TimeTask
	mu       *gmutex.Mutex
}

// AddTask 添加任务
func (s *sTaskList) AddTask(task *model.TimeTask) {
	if task.FuncName == "" || task.Run == nil {
		return
	}
	s.taskList = append(s.taskList, task)
}

// GetByName 通过方法名获取对应task信息
func (s *sTaskList) GetByName(funcName string) *model.TimeTask {
	var result *model.TimeTask
	for _, item := range s.taskList {
		if item.FuncName == funcName {
			result = item
			break
		}
	}
	return result
}

// EditParams 修改参数
func (s *sTaskList) EditParams(funcName string, params []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range s.taskList {
		if item.FuncName == funcName {
			item.Param = params
			break
		}
	}
}