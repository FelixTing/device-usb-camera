// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
	deviceUsbCamera "github.com/edgexfoundry/device-usb-camera"
	"github.com/edgexfoundry/device-usb-camera/internal/driver"
)

const (
	serviceName string = "device-usb-camera"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, deviceUsbCamera.Version, d)
}
