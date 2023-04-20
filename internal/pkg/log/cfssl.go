/*
Copyright 2023 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

import (
	cfssllog "github.com/cloudflare/cfssl/log"
	"github.com/sirupsen/logrus"
)

type cfsslAdapter logrus.Entry

var _ cfssllog.SyslogWriter = (*cfsslAdapter)(nil)

// Debug implements log.SyslogWriter
func (a *cfsslAdapter) Debug(msg string) {
	(*logrus.Entry)(a).Debug(msg)
}

// Info implements log.SyslogWriter
func (a *cfsslAdapter) Info(msg string) {
	(*logrus.Entry)(a).Info(msg)
}

// Warning implements log.SyslogWriter
func (a *cfsslAdapter) Warning(msg string) {
	(*logrus.Entry)(a).Warn(msg)
}

// Err implements log.SyslogWriter
func (a *cfsslAdapter) Err(msg string) {
	(*logrus.Entry)(a).Error(msg)
}

// Crit implements log.SyslogWriter
func (a *cfsslAdapter) Crit(msg string) {
	(*logrus.Entry)(a).Error(msg)
}

// Emerg implements log.SyslogWriter
func (a *cfsslAdapter) Emerg(msg string) {
	(*logrus.Entry)(a).Error(msg)
}
