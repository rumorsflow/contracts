// MIT License
//
// Copyright (c) 2022 Spiral Scout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package config

import "time"

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	//
	// func (s *JobsService) Init(cp config.Configurer) error {
	//     s.config := &JobsConfig{}
	//     if err := cp.UnmarshalKey("jobs", s.config); err != nil {
	//         return err
	//     }
	// }
	UnmarshalKey(name string, out any) error

	// Unmarshal unmarshal the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(out any) error

	// Overwrite used to overwrite particular values in the unmarshalled config
	Overwrite(values map[string]any) error

	// Get used to get config section
	Get(name string) any

	// Has checks if config section exists.
	Has(name string) bool

	// GetVersion returns app version
	GetVersion() string

	// GetCmd returns cli command name
	GetCmd() string

	// GracefulTimeout represents timeout for all servers registered in endure
	GracefulTimeout() time.Duration
}
