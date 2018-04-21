package bnet

import (
	"fmt"
	"net"
	"reflect"
)

// IsAny 检查一个给定的IP地址是不是IPv4 ANY address 或者 IPv6 ANY address
func IsAny(ip interface{}) bool {
	return IsAnyV4(ip) || IsAnyV6(ip)
}

// IsAnyV4 检查给定的IP地址是不是一个IPv4 ANY address.
func IsAnyV4(ip interface{}) bool {
	return iptos(ip) == "0.0.0.0"
}

// IsAnyV6 检查给定的IP地址是不是一个IPv4 ANY address.
func IsAnyV6(ip interface{}) bool {
	ips := iptos(ip)
	return ips == "::" || ips == "[::]"
}

func iptos(ip interface{}) string {
	if ip == nil || reflect.TypeOf(ip).Kind() == reflect.Ptr && reflect.ValueOf(ip).IsNil() {
		return ""
	}

	switch x := ip.(type) {
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case net.IP:
		return x.String()
	case *net.IP:
		return x.String()
	case *net.IPAddr:
		return x.IP.String()
	case *net.TCPAddr:
		return x.IP.String()
	case *net.UDPAddr:
		return x.IP.String()
	default:
		panic(fmt.Sprintf("invalid type: %T", ip))
	}
}
