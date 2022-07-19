package config

import "github.com/gangcheng1030/game_script/cmd/superrobot/model"

var AccountCache = map[string]*model.AccountConfig{
	"gangcheng11223@163.com": {
		AccountName: "gangcheng11223@163.com",
		Password:    "cjdmj092301",
		JunTuanName: "天与地cg",
	},
	"diandian092301@163.com": {
		AccountName: "diandian092301@163.com",
		Password:    "cjdmj092301",
		JunTuanName: "TMR点点",
	},
	"gangcheng11224@163.com": {
		AccountName: "gangcheng11224@163.com",
		Password:    "Cjdmj112233",
		JunTuanName: "TMR刚神2号",
	},
	"gangcheng11225@163.com": {
		AccountName: "gangcheng11225@163.com",
		Password:    "Cjdmj112233",
		JunTuanName: "TMR刚神3号",
	},
	"gangcheng11226@163.com": {
		AccountName: "gangcheng11226@163.com",
		Password:    "Cjdmj112233",
		JunTuanName: "TMR刚神4号",
	},
	"gangcheng11227@163.com": {
		AccountName: "gangcheng11227@163.com",
		Password:    "Cjdmj112233",
		JunTuanName: "TMR刚神5号",
	},
	"yan7955@126.com": {
		AccountName: "yan7955@126.com",
		Password:    "yan7955@",
		JunTuanName: "TMR军哥",
	},
	"yjh179614936@163.com": {
		AccountName: "yjh179614936@163.com",
		Password:    "277426225bb",
		JunTuanName: "他十二姥爷",
	},
	"chenggang092101@163.com": {
		AccountName: "chenggang092101@163.com",
		Password:    "cjdmj092101",
		JunTuanName: "TMR",
	},
}

var NodeConfigCache = map[string]*model.NodeConfig{
	"node1": {
		Group:          "1",
		CanNotBeLeader: false,
		FullScreenMode: 0,
		DirName:        "D:\\Netease\\超激斗梦境",
	},
	"node2": {
		Group:          "1",
		CanNotBeLeader: false,
		FullScreenMode: 0,
		DirName:        "G:\\Netease\\CJDMJ",
	},
	"node3": {
		Group:          "1",
		CanNotBeLeader: true,
		FullScreenMode: 0,
		DirName:        "D:\\Netease\\超激斗梦境",
	},
	"node4": {
		Group:          "1",
		CanNotBeLeader: false,
		FullScreenMode: 1,
		DirName:        "D:\\Netease\\CJDMJ",
	},
	"node5": {
		Group:          "1",
		CanNotBeLeader: false,
		FullScreenMode: 0,
		DirName:        "D:\\Netease\\CJDMJ",
	},
}
