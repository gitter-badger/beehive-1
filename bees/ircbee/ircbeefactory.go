/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package ircbee

import (
	"github.com/muesli/beehive/bees"
)

// IrcBeeFactory is a factory for IrcBees.
type IrcBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *IrcBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := IrcBee{
		Bee: bees.NewBee(name, factory.Name(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// Name returns the name of this Bee.
func (factory *IrcBeeFactory) Name() string {
	return "ircbee"
}

// Description returns the description of this Bee.
func (factory *IrcBeeFactory) Description() string {
	return "An IRC module for beehive"
}

// Image returns the filename of an image for this Bee.
func (factory *IrcBeeFactory) Image() string {
	return factory.Name() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *IrcBeeFactory) LogoColor() string {
	return "#3274d0"
}

// Options returns the options available to configure this Bee.
func (factory *IrcBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "address",
			Description: "Hostname of IRC server, eg: irc.example.org:6667",
			Type:        "url",
			Mandatory:   true,
		},
		{
			Name:        "nick",
			Description: "Nickname to use for IRC",
			Type:        "string",
			Default:     "beehive",
			Mandatory:   true,
		},
		{
			Name:        "password",
			Description: "Password to use to connect to IRC server",
			Type:        "string",
		},
		{
			Name:        "channels",
			Description: "Which channels to join",
			Type:        "[]string",
			Mandatory:   true,
		},
		{
			Name:        "ssl",
			Description: "Use SSL for IRC connection",
			Default:     false,
			Type:        "bool",
		},
	}
	return opts
}

// Events describes the available events provided by this Bee.
func (factory *IrcBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "message",
			Description: "A message was received over IRC, either in a channel or a private query",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "text",
					Description: "The message that was received",
					Type:        "string",
				},
				{
					Name:        "channel",
					Description: "The channel the message was received in",
					Type:        "string",
				},
				{
					Name:        "user",
					Description: "The user that sent the message",
					Type:        "string",
				},
				{
					Name:        "hostmask",
					Description: "Hostmask of the user that sent the message",
					Type:        "string",
				},
			},
		},
	}
	return events
}

// Actions describes the available actions provided by this Bee.
func (factory *IrcBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "send",
			Description: "Sends a message to a channel or a private query",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "channel",
					Description: "Which channel to send the message to",
					Type:        "string",
					Mandatory:   true,
				},
				{
					Name:        "text",
					Description: "Content of the message",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "join",
			Description: "Joins a channel",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "channel",
					Description: "Channel to join",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "part",
			Description: "Parts a channel",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "channel",
					Description: "Channel to part",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
	}
	return actions
}

func init() {
	f := IrcBeeFactory{}
	bees.RegisterFactory(&f)
}
