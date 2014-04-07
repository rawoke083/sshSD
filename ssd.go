// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example shows a basic multi-column menu similar to that found in the
 * ncurses examples from TLDP */
package main

import (
gc "code.google.com/p/goncurses"
 "os/exec"   
    "os"
    //"fmt"
    )

const (
	HEIGHT = 70
	WIDTH  = 90
)

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()


	yMAX, xMAX := stdscr.MaxYX()
	yMAX = yMAX - 5
	xMAX = xMAX - 5
	
	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)
	//gc.InitPair(1, gc.C_RED, gc.C_BLACK)
	//gc.InitPair(2, gc.C_CYAN, gc.C_BLACK)

	// build the menu items
	menu_items := []string{
		" P  - 152.111.192.51",
		" W  - 152.111.192.52",
		" E  - 152.111.192.53",
		" R  - 152.111.192.54",
		" T  - 152.111.192.55",		
		"Exit"}
	items := make([]*gc.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = gc.NewItem(val, "")
		defer items[i].Free()
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

	menuwin, _ := gc.NewWindow(yMAX, xMAX, 1, 1)
	menuwin.Keypad(true)

	menu.SetWindow(menuwin)
	//dwin := menuwin.Derived(6, 38, 3, 1)
	
	
	//menu.SubWindow(dwin)
	menu.Option(gc.O_SHOWDESC, true)
	menu.Format(5, 2)
	menu.Mark("*")

	// MovePrint centered menu title
	title := "My Menu"
	menuwin.Box(0, 0)
	menuwin.ColorOn(1)
	menuwin.MovePrint(1, (WIDTH/2)-(len(title)/2), title)
	menuwin.ColorOff(1)
	menuwin.MoveAddChar(2, 0, gc.ACS_LTEE)
	
	//menuwin.HLine(4, 1, gc.ACS_HLINE, WIDTH-2)
  //  menuwin.HLine(12, 10, gc.ACS_HLINE, WIDTH-2)


	//menuwin.MoveAddChar(2, WIDTH-1, gc.ACS_RTEE)

	y, _ := stdscr.MaxYX()
	stdscr.ColorOn(2)
	stdscr.MovePrint(y-3, 1,
		"Use up/down arrows or page up/down to navigate. 'q' to exit")
	stdscr.ColorOff(2)
	stdscr.Refresh()

	menu.Post()
	defer menu.UnPost()
	menuwin.Refresh()

	for {
		gc.Update()

		 ch := menuwin.GetChar(); 
		 
		 if(ch== 'p'){
			 	
				cmd2 := exec.Command("clear")
				cmd2.Stdout = os.Stdout
				cmd2.Stdin = os.Stdin				
				//cmd2.Stderr = os.Stderr
				cmd2.Run();
				
				cmd := exec.Command("ssh","root@152.111.192.51")
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin				
				cmd.Stderr = os.Stderr
			
				cmd.Run()
				return
						
				
		 } else if (ch=='q' ) {
			return
		} else if (ch==27 ) {
			return
		} else {
			menu.Driver(gc.DriverActions[ch])
		}
	}
}
