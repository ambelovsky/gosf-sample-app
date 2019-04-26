package main

import (
	f "gosf"
	pluginMessageCounter "gosf-message-counter"
)

// PluginMessageCounter is the message counter GOSF plugin method accessor
var PluginMessageCounter pluginMessageCounter.AppMethods

// RegisterPlugins registers all desired GOSF plugins
func RegisterPlugins() {
	f.RegisterPlugin("message-counter", new(pluginMessageCounter.Plugin))
}

// LoadPluginMethods makes plugin methods available for use throughout the application
func LoadPluginMethods() {
	PluginMessageCounter = f.App["message-counter"].(pluginMessageCounter.AppMethods)
}
