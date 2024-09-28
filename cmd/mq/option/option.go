// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package option

import (
	"github.com/goodrain/rainbond/config/configs"
	"github.com/goodrain/rainbond/config/configs/rbdcomponent"
	"github.com/spf13/pflag"
)
import "github.com/sirupsen/logrus"
import "fmt"

// Config config server
type Config struct {
}

// MQServer lb worker server
type MQServer struct {
	MQConfig  *rbdcomponent.MQConfig
	LogConfig *configs.LogConfig
}

// NewMQServer new server
func NewMQServer() *MQServer {
	return &MQServer{}
}

// AddFlags config
func (a *MQServer) AddFlags(fs *pflag.FlagSet) {
	configs.AddLogFlags(fs, a.LogConfig)
	rbdcomponent.AddMQFlags(fs, a.MQConfig)
}

// SetLog 设置log
func (a *MQServer) SetLog() {
	level, err := logrus.ParseLevel(a.LogConfig.LogLevel)
	if err != nil {
		fmt.Println("set log level error." + err.Error())
		return
	}
	logrus.SetLevel(level)
}
