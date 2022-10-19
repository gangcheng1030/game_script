package common

import (
	"github.com/gangcheng1030/game_script/utils/robotgoutil"
	"github.com/go-vgo/robotgo"
)

func Xinyangzhiqiang(tm int) {
	for i := 0; i < tm; i++ {
		xinyangzhiqiangHelper()
	}
}

func xinyangzhiqiangHelper() {
	robotgoutil.Press(robotgo.KeyD, 3)
	robotgoutil.Click(440, 757, 2, 3)
	robotgoutil.Click(384, 217, 2, 2)
	robotgoutil.Click(1414, 872, 2, 2)
	robotgoutil.Click(1414, 872, 2, 6)

	robotgoutil.Click(1168, 371, 0, 2)
	robotgoutil.Press(robotgo.KeyD, 10)

	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Press(robotgo.KeyW, 1)

	robotgoutil.Click(1668, 879, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1398, 192, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(1402, 228, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1300, 258, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(557, 273, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1290, 230, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(1325, 237, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1325, 237, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(795, 143, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(613, 187, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)

	robotgoutil.Click(499, 189, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(655, 176, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(480, 371, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(828, 200, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(1344, 206, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1134, 221, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(679, 317, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(649, 365, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(519, 622, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(571, 748, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)

	robotgoutil.Click(429, 687, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(627, 333, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(625, 268, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(980, 128, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(1085, 177, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1608, 672, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Click(1389, 237, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1378, 191, 1, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Click(1374, 193, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(979, 148, 1, 1)
	robotgoutil.Press(robotgo.KeyE, 1)

	robotgoutil.Click(1349, 223, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Click(1392, 187, 1, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyE, 1)
	robotgoutil.Press(robotgo.KeyW, 1)
	robotgoutil.Press(robotgo.KeyR, 1)
	robotgoutil.Press(robotgo.KeyE, 1)

	for i := 0; i < 8; i++ {
		robotgoutil.Press(robotgo.KeyV, 1)
		robotgoutil.Press(robotgo.KeyW, 1)
		robotgoutil.Press(robotgo.KeyR, 1)
		robotgoutil.Press(robotgo.KeyE, 1)
		robotgoutil.Press(robotgo.KeyW, 1)
		robotgoutil.Press(robotgo.KeyR, 1)
		robotgoutil.Press(robotgo.KeyE, 1)
	}

	robotgoutil.Press(robotgo.KeyT, 5)
	robotgoutil.Press(robotgo.KeyD, 10)

	robotgoutil.Click(1320, 276, 1, 2)
	robotgoutil.Click(1290, 302, 1, 2)
}
