// Code generated by protoc-gen-go.
// source: github.com/google/battery-historian/pb/usagestats_proto/android_battery_usage.proto
// DO NOT EDIT!

/*
Package usagestats is a generated protocol buffer package.

It is generated from these files:
	github.com/google/battery-historian/pb/usagestats_proto/android_battery_usage.proto
	github.com/google/battery-historian/pb/usagestats_proto/android_package_info.proto

It has these top-level messages:
	SystemInfo
*/
package usagestats

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SystemInfo struct {
	// Installed Package List
	InstalledPackages []*PackageInfo `protobuf:"bytes,1,rep,name=installed_packages" json:"installed_packages,omitempty"`
	// e.g. hammerhead-userdebug L MASTER eng.xx... dev-keys Build
	BuildDisplay *string `protobuf:"bytes,2,opt,name=build_display" json:"build_display,omitempty"`
	// e.g. google/hammerhead/hammerhead:L/MASTER/
	//        eng.abcd.20140304.184123:userdebug/dev-keys
	BuildFingerprint *string `protobuf:"bytes,3,opt,name=build_fingerprint" json:"build_fingerprint,omitempty"`
	// Bootloader version e.g. HHZ11k
	Bootloader *string `protobuf:"bytes,4,opt,name=bootloader" json:"bootloader,omitempty"`
	// e.g. T-Mobile
	NetworkOperator *string `protobuf:"bytes,5,opt,name=network_operator" json:"network_operator,omitempty"`
	// e.g. userdebug
	BuildType *string `protobuf:"bytes,6,opt,name=build_type" json:"build_type,omitempty"`
	// e.g. hammerhead
	Device *string `protobuf:"bytes,7,opt,name=device" json:"device,omitempty"`
	// e.g. M8974A-2.0.50.1.02
	BasebandRadio *string `protobuf:"bytes,8,opt,name=baseband_radio" json:"baseband_radio,omitempty"`
	// e.g. hammerhead
	Hardware *string `protobuf:"bytes,9,opt,name=hardware" json:"hardware,omitempty"`
	// Checkin/OTA groups for the device
	//    e.g. auto.droidfood, auto.googlefood.lmp
	Groups []string `protobuf:"bytes,10,rep,name=groups" json:"groups,omitempty"`
	// e.g. 22
	SdkVersion *int32 `protobuf:"varint,11,opt,name=sdk_version" json:"sdk_version,omitempty"`
	// e.g. ES
	CountryCode *string `protobuf:"bytes,12,opt,name=country_code" json:"country_code,omitempty"`
	// e.g. Europe/London
	TimeZone         *string `protobuf:"bytes,13,opt,name=time_zone" json:"time_zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}

func (m *SystemInfo) GetInstalledPackages() []*PackageInfo {
	if m != nil {
		return m.InstalledPackages
	}
	return nil
}

func (m *SystemInfo) GetBuildDisplay() string {
	if m != nil && m.BuildDisplay != nil {
		return *m.BuildDisplay
	}
	return ""
}

func (m *SystemInfo) GetBuildFingerprint() string {
	if m != nil && m.BuildFingerprint != nil {
		return *m.BuildFingerprint
	}
	return ""
}

func (m *SystemInfo) GetBootloader() string {
	if m != nil && m.Bootloader != nil {
		return *m.Bootloader
	}
	return ""
}

func (m *SystemInfo) GetNetworkOperator() string {
	if m != nil && m.NetworkOperator != nil {
		return *m.NetworkOperator
	}
	return ""
}

func (m *SystemInfo) GetBuildType() string {
	if m != nil && m.BuildType != nil {
		return *m.BuildType
	}
	return ""
}

func (m *SystemInfo) GetDevice() string {
	if m != nil && m.Device != nil {
		return *m.Device
	}
	return ""
}

func (m *SystemInfo) GetBasebandRadio() string {
	if m != nil && m.BasebandRadio != nil {
		return *m.BasebandRadio
	}
	return ""
}

func (m *SystemInfo) GetHardware() string {
	if m != nil && m.Hardware != nil {
		return *m.Hardware
	}
	return ""
}

func (m *SystemInfo) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *SystemInfo) GetSdkVersion() int32 {
	if m != nil && m.SdkVersion != nil {
		return *m.SdkVersion
	}
	return 0
}

func (m *SystemInfo) GetCountryCode() string {
	if m != nil && m.CountryCode != nil {
		return *m.CountryCode
	}
	return ""
}

func (m *SystemInfo) GetTimeZone() string {
	if m != nil && m.TimeZone != nil {
		return *m.TimeZone
	}
	return ""
}

func init() {
}
