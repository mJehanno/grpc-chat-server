/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/mjehanno/grpc-chat/service/chat"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const port = ":9000"

type Server struct {
	chat.UnimplementedChatServiceServer
	clients map[string]chat.ChatService_ExchangeMessageServer
}

func (s *Server) ExchangeMessage(stream chat.ChatService_ExchangeMessageServer) error {
	uid := uuid.Must(uuid.NewRandom()).String()
	s.clients[uid] = stream

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		for _, s := range s.clients {
			s.SendMsg(msg)
		}
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		chat.RegisterChatServiceServer(grpcServer, &Server{
			clients: make(map[string]chat.ChatService_ExchangeMessageServer),
		})
		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
