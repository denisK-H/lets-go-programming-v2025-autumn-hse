//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .

package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/kef1rch1k/task-6/internal/wifi"
	mwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var errIntf = errors.New("interfaces error")

const (
	name1 = "Biba"
	name2 = "Boba"

	BroadcastMAC = "ff:ff:ff:ff:ff:ff"
	ZeroMAC      = "00:00:00:00:00:00"
)

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	macBroadcast, err := net.ParseMAC(BroadcastMAC)
	require.NoError(t, err)

	macZero, err := net.ParseMAC(ZeroMAC)
	require.NoError(t, err)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{
			Name:         name1,
			HardwareAddr: macBroadcast,
		},
		{
			Name:         name2,
			HardwareAddr: macZero,
		},
	}, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)

	require.Len(t, addrs, 2)
	require.Equal(t, macBroadcast, addrs[0])
	require.Equal(t, macZero, addrs[1])
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetAddresses()
	require.ErrorContains(t, err, "interfaces error")
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{Name: name1},
		{Name: name2},
	}, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Len(t, names, 2)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetNames()
	require.ErrorContains(t, err, "interfaces error")
}
