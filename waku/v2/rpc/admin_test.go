package rpc

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"testing"

	"github.com/status-im/go-waku/tests"
	"github.com/status-im/go-waku/waku/v2/node"
	"github.com/stretchr/testify/require"
)

func makeAdminService(t *testing.T) *AdminService {
	options := node.WithWakuRelay()
	n, err := node.New(context.Background(), options)
	require.NoError(t, err)
	err = n.Start()
	require.NoError(t, err)
	return &AdminService{n}
}

func TestV1Peers(t *testing.T) {
	port, err := tests.FindFreePort(t, "", 5)
	require.NoError(t, err)

	host, err := tests.MakeHost(context.Background(), port, rand.Reader)
	require.NoError(t, err)

	var reply PeersReply

	request, err := http.NewRequest(http.MethodPost, "url", bytes.NewReader([]byte("")))
	require.NoError(t, err)

	a := makeAdminService(t)

	err = a.GetV1Peers(request, &GetPeersArgs{}, &reply)
	require.NoError(t, err)
	require.Len(t, reply.Peers, 0)

	var reply2 SuccessReply

	addr := fmt.Sprintf("/p2p/%s", host.ID().Pretty())
	err = a.PostV1Peers(request, &PeersArgs{Peers: []string{addr}}, &reply2)
	require.NoError(t, err)
	require.True(t, reply2.Success)
}
