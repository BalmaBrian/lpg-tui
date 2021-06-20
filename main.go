package main

import (
	"log"
	"os/exec"

	ui "github.com/VladimirMarkelov/clui"
)

func createView() {
	// Creating window
	view1 := ui.AddWindow(0, 0, 30, 21, "On Pannel")
	view2 := ui.AddWindow(50, 0, 30, 21, "Off Pannel")

	// Creating main invis Parent Frame
	frameMainOn := ui.CreateFrame(view1, 24, 15, ui.BorderNone, ui.Fixed)
	frameMainOn.SetPack(ui.Horizontal)
	frameMainOff := ui.CreateFrame(view2, 24, 15, ui.BorderNone, ui.Fixed)
	frameMainOff.SetPack(ui.Horizontal)

	// Creating child frame for LOX buttons
	frameLOXOn := ui.CreateFrame(frameMainOn, 24, 15, ui.BorderThin, ui.Fixed)
	frameLOXOn.SetPack(ui.Vertical)
	frameLOXOn.SetTitle("LOX")
	frameLOXOff := ui.CreateFrame(frameMainOff, 24, 15, ui.BorderThin, ui.Fixed)
	frameLOXOff.SetPack(ui.Vertical)
	frameLOXOff.SetTitle("LOX")

	// Creating child frame for FUEL buttons
	frameFUELOn := ui.CreateFrame(frameMainOn, 24, 15, ui.BorderThin, ui.Fixed)
	frameFUELOn.SetPack(ui.Vertical)
	frameFUELOn.SetTitle("FUEL")
	frameFUELOff := ui.CreateFrame(frameMainOff, 24, 15, ui.BorderThin, ui.Fixed)
	frameFUELOff.SetPack(ui.Vertical)
	frameFUELOff.SetTitle("FUEL")

	// Creating LOX buttons
	btnMainLOXOn := ui.CreateButton(frameLOXOn, ui.AutoSize, 4, "Main On", ui.Fixed)
	btnPressLOXOn := ui.CreateButton(frameLOXOn, ui.AutoSize, 4, "Press On", ui.Fixed)
	btnPurgeLOXOn := ui.CreateButton(frameLOXOn, ui.AutoSize, 4, "Purge On", ui.Fixed)
	btnVentLOXOn := ui.CreateButton(frameLOXOn, ui.AutoSize, 4, "Vent On", ui.Fixed)
	btnMainLOXOff := ui.CreateButton(frameLOXOff, ui.AutoSize, 4, "Main Off", ui.Fixed)
	btnPressLOXOff := ui.CreateButton(frameLOXOff, ui.AutoSize, 4, "Press Off", ui.Fixed)
	btnPurgeLOXOff := ui.CreateButton(frameLOXOff, ui.AutoSize, 4, "Purge Off", ui.Fixed)
	btnVentLOXOff := ui.CreateButton(frameLOXOff, ui.AutoSize, 4, "Vent Off", ui.Fixed)

	// Setting LOX button shadow
	btnMainLOXOn.SetShadowType(ui.ShadowHalf)
	btnPressLOXOn.SetShadowType(ui.ShadowHalf)
	btnPurgeLOXOn.SetShadowType(ui.ShadowHalf)
	btnVentLOXOn.SetShadowType(ui.ShadowHalf)
	btnMainLOXOff.SetShadowType(ui.ShadowHalf)
	btnPressLOXOff.SetShadowType(ui.ShadowHalf)
	btnPurgeLOXOff.SetShadowType(ui.ShadowHalf)
	btnVentLOXOff.SetShadowType(ui.ShadowHalf)

	// Creating FUEL buttons
	btnMainFUELOn := ui.CreateButton(frameFUELOn, ui.AutoSize, 4, "Main On", ui.Fixed)
	btnPressFUELOn := ui.CreateButton(frameFUELOn, ui.AutoSize, 4, "Press On", ui.Fixed)
	btnPurgeFUELOn := ui.CreateButton(frameFUELOn, ui.AutoSize, 4, "Purge On", ui.Fixed)
	btnVentFUELOn := ui.CreateButton(frameFUELOn, ui.AutoSize, 4, "Vent On", ui.Fixed)
	btnMainFUELOff := ui.CreateButton(frameFUELOff, ui.AutoSize, 4, "Main Off", ui.Fixed)
	btnPressFUELOff := ui.CreateButton(frameFUELOff, ui.AutoSize, 4, "Press Off", ui.Fixed)
	btnPurgeFUELOff := ui.CreateButton(frameFUELOff, ui.AutoSize, 4, "Purge Off", ui.Fixed)
	btnVentFUELOff := ui.CreateButton(frameFUELOff, ui.AutoSize, 4, "Vent Off", ui.Fixed)

	// Setting FUEL button shadow
	btnMainFUELOn.SetShadowType(ui.ShadowHalf)
	btnPressFUELOn.SetShadowType(ui.ShadowHalf)
	btnPurgeFUELOn.SetShadowType(ui.ShadowHalf)
	btnVentFUELOn.SetShadowType(ui.ShadowHalf)
	btnMainFUELOff.SetShadowType(ui.ShadowHalf)
	btnPressFUELOff.SetShadowType(ui.ShadowHalf)
	btnPurgeFUELOff.SetShadowType(ui.ShadowHalf)
	btnVentFUELOff.SetShadowType(ui.ShadowHalf)

	// Setting LOX button actions
	btnMainLOXOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-MAIN-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPressLOXOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PRESS-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPurgeLOXOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PURGE-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnVentLOXOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-VENT-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnMainLOXOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-MAIN-CLOSED")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPressLOXOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PRESS-CLOSED")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPurgeLOXOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PURGE-CLOSED")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnVentLOXOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-VENT-CLOSED")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})

	// Setting FUEL button actions
	btnMainFUELOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-MAIN-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPressFUELOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PRESS-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPurgeFUELOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PURGE-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnVentFUELOn.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-VENT-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnMainFUELOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-MAIN-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPressFUELOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PRESS-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnPurgeFUELOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-PURGE-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
	btnVentFUELOff.OnClick(func(ev ui.Event) {
		cmd := exec.Command("lpg", "--valve", "/main/FUEL-VENT-OPEN")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
}

func mainLoop() {
	// Every application must create a single Composer and
	// call its intialize method
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	// Windows
	createView()

	// start event processing loop - the main core of the library
	ui.MainLoop()
}

func main() {
	mainLoop()
}