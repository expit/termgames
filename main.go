package main

import (
	"fastReader/fastReader"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

type Dashboard struct {
	questions []Question
	answer *tl.Text
	rectTangle *tl.Rectangle
	posX int
	posY int
	questionsVisible bool
	framesVisible	int
	tick int
	questionMaxLengts int
	answered int
}

type Question struct {
	question *tl.Text
	visible bool
	answered bool

}

func maxLen(questions []Question) (max int){
	max = 0
	for _,question := range questions {
		if len(question.question.Text()) > max {
			max = len(question.question.Text())
		}
	}
	return
}

func (question *Question) isAnswered(answer *tl.Text) bool{
	if question.question.Text() == answer.Text() || question.answered {
		question.answered = true
		question.visible = true
		return true
	}
	return question.answered
}

func (dashboard *Dashboard) Draw (screen *tl.Screen){

	if dashboard.tick < 1 {
		dashboard.answer.SetText("")
	}
	// check if number of ticks passed with visible state
	// change to invisible if ticks greater that framesVisible
	if dashboard.tick < dashboard.framesVisible {
		dashboard.questionsVisible = true
	} else {
		dashboard.questionsVisible = false
	}

	dashboard.rectTangle.Draw(screen)

	dashboard.answered = 0
	for i, question := range dashboard.questions{
		if dashboard.questionsVisible ||  question.isAnswered(dashboard.answer) || question.visible{
			if  question.isAnswered(dashboard.answer) || question.visible {
				dashboard.answered += 1
				dashboard.questions[i].visible = true
			}
			if  question.visible {
				dashboard.questions[i].visible = true
			}
			question.question.SetPosition(dashboard.posX + 1, dashboard.posY +  2 * i )
			question.question.Draw(screen)
		}
	}

	if dashboard.answered == len(dashboard.questions) {
		dashboard.questions = append(dashboard.questions, Question{
			question: tl.NewText(5, 5, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
			visible:  false,
			answered: false,
		})
		dashboard.questions = append(dashboard.questions, Question{
			question: tl.NewText(5, 5, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
			visible:  false,
			answered: false,
		})
	}

	dashboard.answer.SetPosition(dashboard.posX + maxLen(dashboard.questions) + 2, dashboard.posX + 2 * dashboard.answered)
	dashboard.answer.Draw(screen)
}

func (dashboard *Dashboard) Tick(event tl.Event){
	// Increment tick if not reached max
	if dashboard.tick < dashboard.framesVisible {
		dashboard.tick = dashboard.tick +1
	}

	dashboard.questionMaxLengts = maxLen(dashboard.questions)
	questions := len(dashboard.questions)
	dashboard.rectTangle.SetSize(dashboard.questionMaxLengts * 2 + 5, questions * 2 + 3)
	if event.Type == tl.EventKey {
		k := event.Key
		r := event.Ch
		switch k {
		case tl.KeySpace:
			if len(dashboard.answer.Text()) ==5 {
				dashboard.answer.SetText("")
			} else {

			dashboard.tick = 0
			}
		case tl.KeyBackspace:
			fallthrough
		case tl.KeyBackspace2:
			dashboard.answer.SetText(backspace(dashboard.answer.Text()))
		case tl.KeyCtrlR:
			dashboard.answer.SetText("")
		case tl.KeyCtrlO:
			dashboard.questions = []Question{dashboard.questions[1]}
			for i,_ := range dashboard.questions		{
				newRiddle := strconv.Itoa(fastReader.NumberGenerator(float64(5)))
				dashboard.questions[i].question.SetText( newRiddle)
				dashboard.questions[i].answered = false
				dashboard.questions[i].visible = false
				dashboard.tick = 0
			}
		default:
			dashboard.answer.SetText(fmt.Sprintf("%s%s", dashboard.answer.Text(), string(r)))
		}
	}

}
func backspace(s1 string)(s2 string){
	length := len(s1)
	if length == 0{
		return ""
	}
	return s1[:len(s1)-1]
}
func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorDefault,
		Ch: ' ',
	})

	//q1num := fastReader.NumberGenerator(float64(2))
	//q1str := strconv.Itoa(q1num)
	//q1answered := true
	dashboard := Dashboard{
		questions: []Question{
			{
				question: tl.NewText(10, 10, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
				visible: false,
				answered: false,
			},
			{
				question: tl.NewText(11, 11, strconv.Itoa(fastReader.NumberGenerator(float64(5))), tl.ColorGreen, tl.ColorBlack),
				visible:  false,
				answered: false,
			},
		},
		answer:           tl.NewText(8, 6, "answer", tl.ColorGreen|tl.AttrBold, tl.ColorBlack),
		rectTangle: tl.NewRectangle(1,1,1,1,tl.ColorBlack),
		posX:             2,
		posY:             2,
		questionsVisible: true,
		framesVisible:    50,
		tick:             0,
		questionMaxLengts: 0,
	}
	level.AddEntity(&dashboard)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(64)
	game.Start()
}