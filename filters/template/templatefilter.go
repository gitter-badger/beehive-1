/*
 *    Copyright (C) 2017 Christian Muehlhaeuser
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

// Package templatefilter provides a template-based filter.
package templatefilter

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/muesli/beehive/filters"
)

// TemplateFilter is a template-based filter.
type TemplateFilter struct {
}

// Name returns the name of this Filter.
func (filter *TemplateFilter) Name() string {
	return "template"
}

// Description returns the description of this Filter.
func (filter *TemplateFilter) Description() string {
	return "This filter passes when a template-if returns true"
}

// Passes returns true when the Filter matched the data.
func (filter *TemplateFilter) Passes(data interface{}, value interface{}) bool {
	switch v := value.(type) {
	case string:
		var res bytes.Buffer

		funcMap := template.FuncMap{
			"Left": func(values ...interface{}) string {
				return values[0].(string)[:values[1].(int)]
			},
			"Mid": func(values ...interface{}) string {
				if len(values) > 2 {
					return values[0].(string)[values[1].(int):values[2].(int)]
				}
				return values[0].(string)[values[1].(int):]
			},
			"Right": func(values ...interface{}) string {
				return values[0].(string)[len(values[0].(string))-values[1].(int):]
			},
			"Last": func(values ...interface{}) string {
				return values[0].([]string)[len(values[0].([]string))-1]
			},
			// strings functions
			// "Compare":      strings.Compare, // 1.5+ only
			"Contains":     strings.Contains,
			"ContainsAny":  strings.ContainsAny,
			"Count":        strings.Count,
			"EqualFold":    strings.EqualFold,
			"HasPrefix":    strings.HasPrefix,
			"HasSuffix":    strings.HasSuffix,
			"Index":        strings.Index,
			"IndexAny":     strings.IndexAny,
			"Join":         strings.Join,
			"LastIndex":    strings.LastIndex,
			"LastIndexAny": strings.LastIndexAny,
			"Repeat":       strings.Repeat,
			"Replace":      strings.Replace,
			"Split":        strings.Split,
			"SplitAfter":   strings.SplitAfter,
			"SplitAfterN":  strings.SplitAfterN,
			"SplitN":       strings.SplitN,
			"Title":        strings.Title,
			"ToLower":      strings.ToLower,
			"ToTitle":      strings.ToTitle,
			"ToUpper":      strings.ToUpper,
			"Trim":         strings.Trim,
			"TrimLeft":     strings.TrimLeft,
			"TrimPrefix":   strings.TrimPrefix,
			"TrimRight":    strings.TrimRight,
			"TrimSpace":    strings.TrimSpace,
			"TrimSuffix":   strings.TrimSuffix,
		}

		if strings.Index(v, "{{test") >= 0 {
			v = strings.Replace(v, "{{test", "{{if", -1)
			v += "true{{end}}"
		}

		tmpl, err := template.New("_" + v).Funcs(funcMap).Parse(v)
		if err == nil {
			err = tmpl.Execute(&res, data)
		}
		if err != nil {
			panic(err)
		}

		return res.String() == "true"
	}

	return false
}

func init() {
	f := TemplateFilter{}

	filters.RegisterFilter(&f)
}
