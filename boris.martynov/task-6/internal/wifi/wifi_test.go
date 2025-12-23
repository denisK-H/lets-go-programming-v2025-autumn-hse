package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifiImpl "github.com/JingolBong/task-6/internal/wifi"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .
var errMocker = errors.New("error mock interface")

const errGettingInterface = "getting interfaces: "

func TestGetAddressesSuccess(t *testing.T) {
	t.Parallel()

	address, _ := net.ParseMAC("00:11:22:33:44:55")

	mock := NewWiFiHandle(t)
	mock.On("Interfaces").Return([]*wifi.Interface{
		{Name: "wlan", HardwareAddr: address},
	}, nil)

	service := myWifiImpl.New(mock)
	addressGot, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{address}, addressGot)
}

func TestGetAddressesError(t *testing.T) {
	t.Parallel()

	mock := NewWiFiHandle(t)
	mock.On("Interfaces").Return(nil, errMocker)

	service := myWifiImpl.New(mock)
	addrs, err := service.GetAddresses()

	require.Nil(t, addrs)
	require.ErrorContains(t, err, errGettingInterface)
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mock := NewWiFiHandle(t)
	mock.On("Interfaces").Return([]*wifi.Interface{
		{Name: "wlan"},
	}, nil)

	service := myWifiImpl.New(mock)
	nameGot, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan"}, nameGot)
}

func TestGetNamesError(t *testing.T) {
	t.Parallel()

	mock := NewWiFiHandle(t)
	mock.On("Interfaces").Return(nil, errMocker)

	service := myWifiImpl.New(mock)
	names, err := service.GetNames()

	require.Nil(t, names)
	require.ErrorContains(t, err, errGettingInterface)
}
