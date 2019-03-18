package kubeadm

import (
	"fmt"
	"github.com/plombardi89/kubeception/pkg/util"
)

func GenerateKubeadmToken() string {
	r := util.NewRandom()

	part1 := r.RandomString(6)
	part2 := r.RandomString(16)

	return fmt.Sprintf("%s.%s", part1, part2)
}

