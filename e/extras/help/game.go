package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi/gi"
	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"
// 	"math/rand"

	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"

	"github.com/goki/gi/units"
	"github.com/goki/ki"

	"github.com/goki/ki/kit"
	// 	"github.com/goki/gi/svg"
	//"strconv"
	//"math
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/svg"
// 	"time"
)

func main() {

	gimain.Main(func() {
		mainrun()
	})
}

type GameFrame struct {
	Row *gi.Layout
	gi.Frame
}

var KiT_GameFrame = kit.Types.AddType(&GameFrame{}, nil)

func (gf *GameFrame) ConnectEvents2D() {
	// 	fmt.Printf("Hi \n")
	gf.ConnectEvent(oswin.KeyChordEvent, gi.HiPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		// fvv := recv.Embed(KiT_DomFrame).(*DomFrame)
		kt := d.(*key.ChordEvent)
		ch := kt.Chord()

		// fmt.Printf("HI2 \n")
		switch ch {
		case "z":
			kt.SetProcessed()
			gf.UpAction()
		}
	})

}

func (gf *GameFrame) HasFocus2D() bool {
	return true // always.. we're typically a dialog anyway
}
func (gf *GameFrame) UpAction() {

	// 	fmt.Printf("Up action!!\n")
// 	up, _ := gf.Row.ChildByName("upAction", 0)
// 	up.(*gi.Action).Trigger()
}



var SvgGame *svg.SVG
var SvgPeople *svg.Group
var SvgMap *svg.Group

var gmin, gmax, gsz, ginc gi.Vec2D
var GameSize float32 = 200


var trow *gi.Layout

func mainrun() {
	width := 1024
	height := 768

	// turn these on to see a traces of various stages of processing..
	// gi.Update2DTrace = true
	// gi.Render2DTrace = true
	// gi.Layout2DTrace = true
	// ki.SignalTrace = true

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	oswin.TheApp.SetName("Game")
	oswin.TheApp.SetAbout("Game")

	win := gi.NewWindow2D("game", "Game", width, height, true) // true = pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	// style sheet
	var css = ki.Props{
		"Action": ki.Props{
			"background-color": gi.Prefs.Colors.Control, // gi.Color{255, 240, 240, 255},
		},
		"#combo": ki.Props{
			"background-color": gi.Color{240, 255, 240, 255},
		},
		".hslides": ki.Props{
			"background-color": gi.Color{240, 225, 255, 255},
		},
		"kbd": ki.Props{
			"color": "blue",
		},
	}
	vp.CSS = css

	mfr := win.SetMainFrame()
	// 	dfr := mfr.AddNewChild(KiT_DomFrame, "domframe").(*DomFrame)
	mfr.SetProp("spacing", units.NewValue(1, units.Ex))
	// mfr.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// dfr.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// dfr.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// dfr.SetProp("background-color", "radial-gradient(red, lighter-80)")

	// vars in here :

	// end of vars

	trow = mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>Game - Play me now!</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()

	trow.AddNewChild(gi.KiT_Space, "spc1")
	
	gfr := mfr.AddNewChild(KiT_GameFrame, "gameframe").(*GameFrame)
	gfr.SetProp("background-color", "white")

	gfr.Row = mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	gfr.Row.Lay = gi.LayoutHoriz

	doSomethingOne := gfr.Row.AddNewChild(gi.KiT_Action, "doSomethingOne").(*gi.Action)
	doSomethingOne.Text = "Do something One"

	doSomethingOne.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
	

	})

	doSomethingTwo := gfr.Row.AddNewChild(gi.KiT_Action, "doSomethingTwo").(*gi.Action)
	doSomethingTwo.Text = "Do something Two"

	doSomethingTwo.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		

	})

	gfr.AddNewChild(gi.KiT_Space, "spc2")

	SvgGame = gfr.AddNewChild(svg.KiT_SVG, "SvgGame").(*svg.SVG)
	SvgGame.SetProp("min-width", GameSize)
	SvgGame.SetProp("min-height", GameSize)
	SvgGame.SetStretchMaxWidth()
	SvgGame.SetStretchMaxHeight()

	
	SvgPeople = SvgGame.AddNewChild(svg.KiT_Group, "SvgPeople").(*svg.Group)
	SvgMap = SvgGame.AddNewChild(svg.KiT_Group, "SvgMap").(*svg.Group)

	gmin = gi.Vec2D{-10, -10}
	gmax = gi.Vec2D{10, 10}
	gsz = gmax.Sub(gmin)
	ginc = gsz.DivVal(GameSize)

	SvgGame.ViewBox.Min = gmin
	SvgGame.ViewBox.Size = gsz
	SvgGame.Norm = true
	SvgGame.InvertY = true
	SvgGame.Fill = true
	SvgGame.SetProp("background-color", "white")
	SvgGame.SetProp("stroke-width", ".8pct")

	InitMap()
	InitPlayer()

	//////////////////////////////////////////
	//      Main Menu

	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.KnownChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.KnownChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}

func InitMap() {
	
}

 var player *svg.Rect

func InitPlayer() {
	updt := SvgGame.UpdateStart()
	SvgPeople.DeleteChildren(true)

	player = SvgPeople.AddNewChild(svg.KiT_Rect, "player").(*svg.Rect)

	player.SetProp("fill", "red")
	player.SetProp("stroke", "darkred")
	player.Size = gi.Vec2D{2, 2}
	player.Pos = gi.Vec2D{-5, -10}

	SvgGame.UpdateEnd(updt)

}



// func JumpLoop() {
//   fmt.Printf("HIII \n")

//   for y := -9.9; y > -10; y++ {
//     updt := SvgGame.UpdateStart()
//     if y < 10 {

//       if VertSpeed == 1 {
//       player.Pos.Y = float32(y)
//       } else {
//         player.Pos.Y = float32(y) - 2
//       }

//     } else {
//       fmt.Printf("Coming down \n")
//           SvgGame.UpdateEnd(updt)

//       break
//     }
//     SvgGame.UpdateEnd(updt)
//     time.Sleep(1 * time.Millisecond)

//   }
//   JumpLoopDown()

// }

// func JumpLoopDown() {
//   fmt.Printf("Coming down func \n")

//   for y := player.Pos.Y; y >= -10; y-- {
//     updt := SvgGame.UpdateStart()
//     player.Pos.Y = float32(y)
//     SvgGame.UpdateEnd(updt)
//     fmt.Printf("Updated before this! \n")
//     time.Sleep(1 * time.Millisecond)

//   }

// }

var obstacle *svg.Rect

func MainLoop() {
	
	for i := 0; i > -1; i++ {

		updt := SvgGame.UpdateStart()

		SvgGame.UpdateEnd(updt)

	}
}

