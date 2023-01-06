package main

import (
	"log"
	"net/http"

	ingressfactory "github.com/devMarcus21/eventfunnel-azure/pkg/syncreceivers/eventingress/eventingressfactory"
	"github.com/devMarcus21/eventfunnel-azure/pkg/utils/httpwrapper"
)

func main() {
	http.HandleFunc("/ingress", httpwrapper.BuildHttpHandler(ingressfactory.CreateEventIngressHandler()))

	log.Fatal(http.ListenAndServe(":80", nil))
}
