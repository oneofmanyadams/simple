package sconfiguration

import (
	"fmt"
)

type Sconf struct {
	file string
	settings map[string]string
	flags map[string]string
	testerFile string
}

////////////////////////////////
// General methods
////////////////////////////////

func NewSconf(file_location string) (new_sconf Sconf) {
	new_sconf.file = file_location
	new_sconf.settings = make(map[string]string)
	new_sconf.flags = make(map[string]string)
	return
}

func (sconf *Sconf) Load() {
	sconf.settings, sconf.flags = LoadConfigFromFileWithFlags(sconf.file)
}

func (sconf *Sconf) Save() {
	SaveConfigToFileWithFlags(sconf.file, sconf.settings, sconf.flags)
}

func (sconf *Sconf) Display() {
	fmt.Println("CONFIGURATION DISPLAY")
	fmt.Println("===================================================")
	fmt.Println("Conf File:", sconf.file)
	fmt.Println("Test File:", sconf.testerFile)
	fmt.Println("===================================================")	
	for setting_name, setting_value := range sconf.settings {
		fmt.Println("---------------------------------------------------")
		fmt.Println("Setting:", setting_name)
		fmt.Println("	-", "Value:", setting_value)
		fmt.Println("	-", "Flag:", sconf.flags[setting_name])
	}
}

////////////////////////////////
// Main methods
////////////////////////////////
func (sconf *Sconf) InfoFor(setting string) (value string, flag string) {
	value = sconf.ValueFor(setting)
	flag = sconf.FlagFor(setting)
	return
}

func (sconf *Sconf) ChangeInfoFor(setting string, value string, flag string) {
	sconf.ChangeSettingValue(setting, value)
	sconf.ChangeFlagValue(setting, flag)
}

////////////////////////////////
// Value only methods
////////////////////////////////

func (sconf *Sconf) ValueFor(setting string) (value string) {
	if found_value, setting_exists := sconf.settings[setting]; setting_exists {
		value = found_value
	}
	return
}

func (sconf *Sconf) ChangeSettingValue(setting string, new_value string) {
	if _, setting_exists := sconf.settings[setting]; setting_exists {
		sconf.settings[setting] = new_value
	}	
}

////////////////////////////////
// Flag only methods
////////////////////////////////

func (sconf *Sconf) FlagFor(setting string) (value string) {
	if found_flag, flag_exists := sconf.flags[setting]; flag_exists {
		value = found_flag
	}
	return
}

func (sconf *Sconf) ChangeFlagValue(setting string, new_value string) {
	if _, flag_exists := sconf.flags[setting]; flag_exists {
		sconf.flags[setting] = new_value
	}	
}
