package main

// This file uses mplayer as a playback back-end.

import (
	"log"
	"os/exec"
	//"fmt"
)

type MPlayer struct {
	cmd *exec.Cmd
}

func (m *MPlayer) Play(file string) {
	m.cmd = exec.Command("mplayer", "-really-quiet", file)
	out, err := m.cmd.CombinedOutput()

	// If the command was killed with signal 9,
	// the player was just stopped, so we don'