package regulus

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const AsciiArt = `╔═══╗            ╔╗         
║╔═╗║            ║║         
║╚═╝║╔══╗╔══╗╔╗╔╗║║ ╔╗╔╗╔══╗
║╔╗╔╝║╔╗║║╔╗║║║║║║║ ║║║║║══╣
║║║╚╗║║═╣║╚╝║║╚╝║║╚╗║╚╝║╠══║
╚╝╚═╝╚══╝╚═╗║╚══╝╚═╝╚══╝╚══╝
         ╔═╝║               
         ╚══╝               
`

type Regulus struct {
	server *http.Server
}

func New(addr string) *Regulus {
	mux := http.NewServeMux()
	mux.HandleFunc("/ip", getMyIpHandler)

	return &Regulus{
		server: &http.Server{
			Addr:              addr,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,

			Handler: mux,
		},
	}
}

func (r *Regulus) Start() error {
	log.Info().Str("address", r.server.Addr).Msg("Starting server")
	return r.server.ListenAndServe()
}
