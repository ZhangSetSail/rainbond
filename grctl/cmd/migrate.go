package cmd

import (
	"context"
	"fmt"
	rainbondv1alpha1 "github.com/goodrain/rainbond-operator/api/v1alpha1"
	"github.com/goodrain/rainbond/grctl/clients"
	"github.com/urfave/cli"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func NewCmdMigrateConsole() cli.Command {
	c := cli.Command{
		Name:  "migrate",
		Usage: "migrate the console to the cluster deployment",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "env,e",
				Usage: "configure console env. For example <grctl migrate -e MYSQL_HOST=127.0.0.1>",
			},
			cli.StringSliceFlag{
				Name:  "label,l",
				Usage: "configure console label. For example <grctl migrate -l label=label>",
			},
			cli.StringSliceFlag{
				Name:  "arg,a",
				Usage: "configure console arg. For example <grctl migrate -a ./test.sh>",
			},
			cli.StringFlag{
				Name:  "image,i",
				Usage: "configure console image. For example <grctl migrate -i console:test>",
				Value: "registry.cn-hangzhou.aliyuncs.com/goodrain/rainbond:v5.10.1-release-allinone",
			},
			cli.IntFlag{
				Name:  "replicas,r",
				Usage: "configure console replicas. For example <grctl migrate -r 1>",
				Value: 1,
			},
			cli.StringFlag{
				Name:  "port,p",
				Usage: "configure console port. For example <grctl migrate -p 7070>",
				Value: "7070",
			},
		},
		Action: func(c *cli.Context) error {
			Common(c)
			return initConsoleYaml(c)
		},
	}
	return c
}

func initConsoleYaml(ctx *cli.Context) error {
	labels := make(map[string]string)
	labels["belongTo"] = "rainbond-operator"
	labels["creator"] = "Rainbond"
	labels["name"] = "rbd-app-ui"
	labels["port"] = ctx.String("port")
	for _, label := range ctx.StringSlice("label") {
		labelkv := strings.Split(label, "=")
		if len(labelkv) != 2 {
			return fmt.Errorf("label format is incorrect %v", label)
		}
		labels[labelkv[0]] = labelkv[1]
	}
	envs := []corev1.EnvVar{corev1.EnvVar{
		Name:  "DB_TYPE",
		Value: "mysql",
	}}
	for _, env := range ctx.StringSlice("env") {
		envkv := strings.Split(env, "=")
		if len(envkv) != 2 {
			return fmt.Errorf("env format is incorrect %v", envkv)
		}
		envs = append(envs, corev1.EnvVar{
			Name:  envkv[0],
			Value: envkv[1],
		})
	}
	replicas := int32(ctx.Int("replicas"))
	consoleObject := rainbondv1alpha1.RbdComponent{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RbdComponent",
			APIVersion: "rainbond.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "rbd-app-ui",
			Namespace: "rbd-system",
			Labels:    labels,
		},
		Spec: rainbondv1alpha1.RbdComponentSpec{
			Replicas:        &replicas,
			Image:           ctx.String("image"),
			ImagePullPolicy: "IfNotPresent",
			Args:            ctx.StringSlice("arg"),
			Env:             envs,
		},
	}
	err := clients.RainbondKubeClient.Create(context.Background(), &consoleObject)
	if err != nil {
		return fmt.Errorf("create rbdcomponent rbd-app-ui error:%v", err)
	}
	return nil
}
