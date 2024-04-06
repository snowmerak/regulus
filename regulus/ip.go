package regulus

import (
	"net"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	ResponseFormatString = "str"
	ResponseFormatJSON   = "json"
	ResponseFormatXML    = "xml"
	ResponseFormatYAML   = "yaml"
	ResponseFormatTOML   = "toml"
)

type IpResponse struct {
	Ip string `json:"ip" xml:"ip" yaml:"ip" toml:"ip"`
}

func getMyIpHandler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	switch format {
	case ResponseFormatString:
	case ResponseFormatJSON:
	case ResponseFormatXML:
	case ResponseFormatYAML:
	case ResponseFormatTOML:
	default:
		format = ResponseFormatString
	}

	remoteIp := r.Header.Get("X-Forwarded-For")
	if remoteIp == "" {
		remoteIp = r.Header.Get("X-ProxyUser-Ip")
		if remoteIp == "" {
			remoteIp = r.Header.Get("X-Real-IP")

			if remoteIp == "" {
				ip, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					w.Write([]byte("Cannot get remote IP address"))
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				remoteIp = ip
			}
		}
	}

	remoteIps := strings.Split(remoteIp, ",")
	for i := range remoteIps {
		remoteIps[i] = strings.TrimSpace(remoteIps[i])
	}

	if len(remoteIps) == 0 {
		w.Write([]byte("Cannot get remote IP address"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// only use first IP
	switch format {
	case ResponseFormatString:
		if _, err := w.Write([]byte(remoteIps[0])); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case ResponseFormatJSON:
		if _, err := w.Write([]byte(`{"ip":"` + remoteIps[0] + `"}`)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case ResponseFormatXML:
		if _, err := w.Write([]byte(`<ip>` + remoteIps[0] + `</ip>`)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case ResponseFormatYAML:
		if _, err := w.Write([]byte("ip: " + remoteIps[0])); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case ResponseFormatTOML:
		if _, err := w.Write([]byte("ip = " + remoteIps[0])); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		if _, err := w.Write([]byte(remoteIps[0])); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	log.Info().Strs("remote_ip_list", remoteIps).Msg("Got remote IP address")
}
