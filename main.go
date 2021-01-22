package main

import "termgames/fastReader"

func main() {
	fastReader.Gamestart()
	//game := tl.NewGame()
	//
	//level := tl.NewBaseLevel(tl.Cell{
	//	Bg: tl.ColorDefault,
	//	Fg: tl.ColorDefault,
	//	Ch: ' ',
	//})
	//
	////q1num := fastReader.NumberGenerator(float64(2))
	////q1str := strconv.Itoa(q1num)
	////q1answered := true
	//dashboard := Dashboard{
	//	questions: []Question{
	//		{
	//			question: tl.NewText(10, 10, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
	//			visible: false,
	//			answered: false,
	//		},
	//		{
	//			question: tl.NewText(11, 11, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
	//			visible:  false,
	//			answered: false,
	//		},
	//	},
	//	answer:           tl.NewText(8, 6, "answer", tl.ColorGreen|tl.AttrBold, tl.ColorBlack),
	//	rectTangle: tl.NewRectangle(1,1,1,1,tl.ColorBlack),
	//	posX:             2,
	//	posY:             2,
	//	questionsVisible: true,
	//	framesVisible:    50,
	//	tick:             0,
	//	questionMaxLengts: 0,
	//}
	//level.AddEntity(&dashboard)
	//
	//game.Screen().SetLevel(level)
	//game.Screen().SetFps(64)
	//game.Start()
}
