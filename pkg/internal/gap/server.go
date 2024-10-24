package gap

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/rs/zerolog/log"
	"strings"

	"github.com/spf13/viper"
)

var Nx *nex.Conn

func InitializeToNexus() error {
	grpcBind := strings.SplitN(viper.GetString("grpc_bind"), ":", 2)

	outboundIp, _ := nex.GetOutboundIP()

	grpcOutbound := fmt.Sprintf("%s:%s", outboundIp, grpcBind[1])

	var err error
	Nx, err = nex.NewNexusConn(viper.GetString("nexus_addr"), &proto.ServiceInfo{
		Id:       viper.GetString("id"),
		Type:     nex.ServiceTypePusher,
		Label:    "Pusher",
		GrpcAddr: grpcOutbound,
	})
	if err == nil {
		go func() {
			err := Nx.RunRegistering()
			if err != nil {
				log.Error().Err(err).Msg("An error occurred while registering service...")
			}
		}()
	}

	return err
}