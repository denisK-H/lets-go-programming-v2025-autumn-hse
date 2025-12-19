package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/VlasfimosY/task-6/internal/wifi"
	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .

var errMock = errors.New("mock error")

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface(nil), errMock)

	wifiService := wifi.New(mockHandle)

	_, err := wifiService.GetAddresses()
	require.ErrorContains(t, err, "getting interfaces:")
}

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	t.Parallel()

	addr1, _ := net.ParseMAC("00:11:22:33:44:55")
	addr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: addr1, Name: "wlan0"},
		{HardwareAddr: addr2, Name: "wlan1"},
	}, nil)

	wifiService := wifi.New(mockHandle)

	addrs, err := wifiService.GetAddresses()
	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{addr1, addr2}, addrs)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface(nil), errMock)

	wifiService := wifi.New(mockHandle)

	_, err := wifiService.GetNames()
	require.ErrorContains(t, err, "getting interfaces:")
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	t.Parallel()

	addr1, _ := net.ParseMAC("00:11:22:33:44:55")

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifilib.Interface{
		{HardwareAddr: addr1, Name: "wlan0"},
		{HardwareAddr: nil, Name: "lo"},
	}, nil)

	wifiService := wifi.New(mockHandle)

	names, err := wifiService.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "lo"}, names)
}
