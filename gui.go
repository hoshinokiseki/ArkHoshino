package main

import (
	g "github.com/AllenDang/giu"
	"strconv"
)

type doctor struct {
	isInitialized bool
	nickName      string
	nickNumber    string
	level         int
	exp           int
	diamond       int // 源石
	orundum       int // 合成玉
}

var (
	doctorInfo = doctor{false, "Blaze", "2021", 120, 0, 0, 0}
)

func loop() {
	g.SingleWindow().Layout(
		g.TabBar().TabItems(
			g.TabItem("主页").Layout(
				g.Row(
					g.ImageWithFile("resources/logo.png").Size(120, 120),
					g.Column(
						g.Row(
							g.Label(doctorInfo.nickName),
							g.Label("Lv. "+strconv.Itoa(doctorInfo.level)),
						),
						g.Label("这里应当有一行个性签名，但不知为何鹰角至今没有开放修改，于是我也就不做了。").Wrapped(true),
						g.ProgressBar(float32(doctorInfo.exp/DoctorLevelExp[doctorInfo.level])).Size(-1, 0).Overlay("EXP"),
					),
				),
				g.Row(
					g.Table().Rows(
						g.TableRow(
							g.Label("源石：0"),
							g.Label("合成玉：0"),
						),
						g.TableRow(
							g.Label("资质凭证：0"),
							g.Label("高级凭证：0"),
						),
						g.TableRow(
							g.Label("采购凭证：0"),
							g.Label("信用点数：0"),
						),
					).Flags(g.TableFlagsNoBordersInBody),
				),
			),
		),
	)
}

func main() {
	g.SetDefaultFont("resources/font.otf", 19)
	wnd := g.NewMasterWindow("ArkHoshino", 400, 720, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
