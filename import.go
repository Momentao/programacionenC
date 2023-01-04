package main

// This file implements the "import" command

func init() {
	help["import"] = `Import music directory into library`
}

func (api API) Imp