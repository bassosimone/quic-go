package ackhandler

import (
	"github.com/bassosimone/quic-go/internal/protocol"
	"github.com/bassosimone/quic-go/internal/utils"
	"github.com/bassosimone/quic-go/logging"
	"github.com/bassosimone/quic-go/quictrace"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *utils.RTTStats,
	pers protocol.Perspective,
	traceCallback func(quictrace.Event),
	tracer logging.ConnectionTracer,
	logger utils.Logger,
	version protocol.VersionNumber,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, rttStats, pers, traceCallback, tracer, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger, version)
}
