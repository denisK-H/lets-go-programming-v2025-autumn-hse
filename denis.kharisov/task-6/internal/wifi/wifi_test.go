package wifi_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/mdlayher/wifi"

	myWiFi "github.com/denisK-H/task-6/internal/wifi"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .

const (
	gettingInterfacesError = "getting interfaces: "
	testWifi               = "wlan0"
	testMAC                = "00:11:22:33:44:55"
)

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)

	mockHandle.On("Interfaces").Return([]*wifi.Interface{
		{Name: testWifi},
	}, nil)

	service := myWiFi.New(mockHandle)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{testWifi}, names)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return(nil, assert.AnError)

	service := myWiFi.New(mockHandle)
	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, gettingInterfacesError)
}

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mac, err := net.ParseMAC(testMAC)
	require.NoError(t, err)

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifi.Interface{
		{HardwareAddr: mac},
	}, nil)

	service := myWiFi.New(mockHandle)
	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{mac}, addrs)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return(nil, assert.AnError)

	service := myWiFi.New(mockHandle)
	addrs, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, addrs)
	require.ErrorContains(t, err, gettingInterfacesError)
}