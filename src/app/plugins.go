package main

import (
	f "github.com/ambelovsky/gosf"
	//pluginMessageCounter "github.com/ambelovsky/gosf-message-counter"
	pluginSamplePlugin "github.com/ambelovsky/gosf-sample-plugin"
)

// PluginMessageCounter is the message counter GOSF plugin method accessor
//var PluginMessageCounter pluginMessageCounter.AppMethods

// PluginSamplePlugin is the sample GOSF plugin method accessor
var PluginSamplePlugin pluginSamplePlugin.AppMethods

// RegisterPlugins registers all desired GOSF plugins
func RegisterPlugins() {
	//f.RegisterPlugin("message-counter", new(pluginMessageCounter.Plugin))
	f.RegisterPlugin("sample-plugin", new(pluginSamplePlugin.Plugin))
}

// LoadPluginMethods makes plugin methods available for use throughout the application
func LoadPluginMethods() {
	//PluginMessageCounter = f.App["message-counter"].(pluginMessageCounter.AppMethods)
	PluginSamplePlugin = f.App["message-counter"].(pluginSamplePlugin.AppMethods)
}
