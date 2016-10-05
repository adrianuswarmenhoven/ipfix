//Generated, do not edit
package ipfixmessage

import (
	"fmt"
)

// NewFieldValueByID returns an empty FieldValue that matches the enterprise id and element id
func NewFieldValueByID(enterpriseid int, elementid int) (FieldValue, error) {
	switch enterpriseid {

	case 0: // IANA - https://www.ietf.org/assignments/ipfix/ipfix.xml

		switch elementid {
		case 1:
			return &FieldValueUnsigned64{}, nil // octetDeltaCount
		case 2:
			return &FieldValueUnsigned64{}, nil // packetDeltaCount
		case 3:
			return &FieldValueUnsigned64{}, nil // deltaFlowCount
		case 4:
			return &FieldValueUnsigned8{}, nil // protocolIdentifier
		case 5:
			return &FieldValueUnsigned8{}, nil // ipClassOfService
		case 6:
			return &FieldValueUnsigned16{}, nil // tcpControlBits
		case 7:
			return &FieldValueUnsigned16{}, nil // sourceTransportPort
		case 8:
			return &FieldValueIPv4Address{}, nil // sourceIPv4Address
		case 9:
			return &FieldValueUnsigned8{}, nil // sourceIPv4PrefixLength
		case 10:
			return &FieldValueUnsigned32{}, nil // ingressInterface
		case 11:
			return &FieldValueUnsigned16{}, nil // destinationTransportPort
		case 12:
			return &FieldValueIPv4Address{}, nil // destinationIPv4Address
		case 13:
			return &FieldValueUnsigned8{}, nil // destinationIPv4PrefixLength
		case 14:
			return &FieldValueUnsigned32{}, nil // egressInterface
		case 15:
			return &FieldValueIPv4Address{}, nil // ipNextHopIPv4Address
		case 16:
			return &FieldValueUnsigned32{}, nil // bgpSourceAsNumber
		case 17:
			return &FieldValueUnsigned32{}, nil // bgpDestinationAsNumber
		case 18:
			return &FieldValueIPv4Address{}, nil // bgpNextHopIPv4Address
		case 19:
			return &FieldValueUnsigned64{}, nil // postMCastPacketDeltaCount
		case 20:
			return &FieldValueUnsigned64{}, nil // postMCastOctetDeltaCount
		case 21:
			return &FieldValueUnsigned32{}, nil // flowEndSysUpTime
		case 22:
			return &FieldValueUnsigned32{}, nil // flowStartSysUpTime
		case 23:
			return &FieldValueUnsigned64{}, nil // postOctetDeltaCount
		case 24:
			return &FieldValueUnsigned64{}, nil // postPacketDeltaCount
		case 25:
			return &FieldValueUnsigned64{}, nil // minimumIpTotalLength
		case 26:
			return &FieldValueUnsigned64{}, nil // maximumIpTotalLength
		case 27:
			return &FieldValueIPv6Address{}, nil // sourceIPv6Address
		case 28:
			return &FieldValueIPv6Address{}, nil // destinationIPv6Address
		case 29:
			return &FieldValueUnsigned8{}, nil // sourceIPv6PrefixLength
		case 30:
			return &FieldValueUnsigned8{}, nil // destinationIPv6PrefixLength
		case 31:
			return &FieldValueUnsigned32{}, nil // flowLabelIPv6
		case 32:
			return &FieldValueUnsigned16{}, nil // icmpTypeCodeIPv4
		case 33:
			return &FieldValueUnsigned8{}, nil // igmpType
		case 34:
			return &FieldValueUnsigned32{}, nil // samplingInterval
		case 35:
			return &FieldValueUnsigned8{}, nil // samplingAlgorithm
		case 36:
			return &FieldValueUnsigned16{}, nil // flowActiveTimeout
		case 37:
			return &FieldValueUnsigned16{}, nil // flowIdleTimeout
		case 38:
			return &FieldValueUnsigned8{}, nil // engineType
		case 39:
			return &FieldValueUnsigned8{}, nil // engineId
		case 40:
			return &FieldValueUnsigned64{}, nil // exportedOctetTotalCount
		case 41:
			return &FieldValueUnsigned64{}, nil // exportedMessageTotalCount
		case 42:
			return &FieldValueUnsigned64{}, nil // exportedFlowRecordTotalCount
		case 43:
			return &FieldValueIPv4Address{}, nil // ipv4RouterSc
		case 44:
			return &FieldValueIPv4Address{}, nil // sourceIPv4Prefix
		case 45:
			return &FieldValueIPv4Address{}, nil // destinationIPv4Prefix
		case 46:
			return &FieldValueUnsigned8{}, nil // mplsTopLabelType
		case 47:
			return &FieldValueIPv4Address{}, nil // mplsTopLabelIPv4Address
		case 48:
			return &FieldValueUnsigned8{}, nil // samplerId
		case 49:
			return &FieldValueUnsigned8{}, nil // samplerMode
		case 50:
			return &FieldValueUnsigned32{}, nil // samplerRandomInterval
		case 51:
			return &FieldValueUnsigned8{}, nil // classId
		case 52:
			return &FieldValueUnsigned8{}, nil // minimumTTL
		case 53:
			return &FieldValueUnsigned8{}, nil // maximumTTL
		case 54:
			return &FieldValueUnsigned32{}, nil // fragmentIdentification
		case 55:
			return &FieldValueUnsigned8{}, nil // postIpClassOfService
		case 56:
			return &FieldValueMacAddress{}, nil // sourceMacAddress
		case 57:
			return &FieldValueMacAddress{}, nil // postDestinationMacAddress
		case 58:
			return &FieldValueUnsigned16{}, nil // vlanId
		case 59:
			return &FieldValueUnsigned16{}, nil // postVlanId
		case 60:
			return &FieldValueUnsigned8{}, nil // ipVersion
		case 61:
			return &FieldValueUnsigned8{}, nil // flowDirection
		case 62:
			return &FieldValueIPv6Address{}, nil // ipNextHopIPv6Address
		case 63:
			return &FieldValueIPv6Address{}, nil // bgpNextHopIPv6Address
		case 64:
			return &FieldValueUnsigned32{}, nil // ipv6ExtensionHeaders
		case 70:
			return &FieldValueOctetArray{}, nil // mplsTopLabelStackSection
		case 71:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection2
		case 72:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection3
		case 73:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection4
		case 74:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection5
		case 75:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection6
		case 76:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection7
		case 77:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection8
		case 78:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection9
		case 79:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection10
		case 80:
			return &FieldValueMacAddress{}, nil // destinationMacAddress
		case 81:
			return &FieldValueMacAddress{}, nil // postSourceMacAddress
		case 82:
			return &FieldValueString{}, nil // interfaceName
		case 83:
			return &FieldValueString{}, nil // interfaceDescription
		case 84:
			return &FieldValueString{}, nil // samplerName
		case 85:
			return &FieldValueUnsigned64{}, nil // octetTotalCount
		case 86:
			return &FieldValueUnsigned64{}, nil // packetTotalCount
		case 87:
			return &FieldValueUnsigned32{}, nil // flagsAndSamplerId
		case 88:
			return &FieldValueUnsigned16{}, nil // fragmentOffset
		case 89:
			return &FieldValueUnsigned32{}, nil // forwardingStatus
		case 90:
			return &FieldValueOctetArray{}, nil // mplsVpnRouteDistinguisher
		case 91:
			return &FieldValueUnsigned8{}, nil // mplsTopLabelPrefixLength
		case 92:
			return &FieldValueUnsigned32{}, nil // srcTrafficIndex
		case 93:
			return &FieldValueUnsigned32{}, nil // dstTrafficIndex
		case 94:
			return &FieldValueString{}, nil // applicationDescription
		case 95:
			return &FieldValueOctetArray{}, nil // applicationId
		case 96:
			return &FieldValueString{}, nil // applicationName
		case 98:
			return &FieldValueUnsigned8{}, nil // postIpDiffServCodePoint
		case 99:
			return &FieldValueUnsigned32{}, nil // multicastReplicationFactor
		case 100:
			return &FieldValueString{}, nil // className
		case 101:
			return &FieldValueUnsigned8{}, nil // classificationEngineId
		case 102:
			return &FieldValueUnsigned16{}, nil // layer2packetSectionOffset
		case 103:
			return &FieldValueUnsigned16{}, nil // layer2packetSectionSize
		case 104:
			return &FieldValueOctetArray{}, nil // layer2packetSectionData
		case 128:
			return &FieldValueUnsigned32{}, nil // bgpNextAdjacentAsNumber
		case 129:
			return &FieldValueUnsigned32{}, nil // bgpPrevAdjacentAsNumber
		case 130:
			return &FieldValueIPv4Address{}, nil // exporterIPv4Address
		case 131:
			return &FieldValueIPv6Address{}, nil // exporterIPv6Address
		case 132:
			return &FieldValueUnsigned64{}, nil // droppedOctetDeltaCount
		case 133:
			return &FieldValueUnsigned64{}, nil // droppedPacketDeltaCount
		case 134:
			return &FieldValueUnsigned64{}, nil // droppedOctetTotalCount
		case 135:
			return &FieldValueUnsigned64{}, nil // droppedPacketTotalCount
		case 136:
			return &FieldValueUnsigned8{}, nil // flowEndReason
		case 137:
			return &FieldValueUnsigned64{}, nil // commonPropertiesId
		case 138:
			return &FieldValueUnsigned64{}, nil // observationPointId
		case 139:
			return &FieldValueUnsigned16{}, nil // icmpTypeCodeIPv6
		case 140:
			return &FieldValueIPv6Address{}, nil // mplsTopLabelIPv6Address
		case 141:
			return &FieldValueUnsigned32{}, nil // lineCardId
		case 142:
			return &FieldValueUnsigned32{}, nil // portId
		case 143:
			return &FieldValueUnsigned32{}, nil // meteringProcessId
		case 144:
			return &FieldValueUnsigned32{}, nil // exportingProcessId
		case 145:
			return &FieldValueUnsigned16{}, nil // templateId
		case 146:
			return &FieldValueUnsigned8{}, nil // wlanChannelId
		case 147:
			return &FieldValueString{}, nil // wlanSSID
		case 148:
			return &FieldValueUnsigned64{}, nil // flowId
		case 149:
			return &FieldValueUnsigned32{}, nil // observationDomainId
		case 150:
			return &FieldValueDateTimeSeconds{}, nil // flowStartSeconds
		case 151:
			return &FieldValueDateTimeSeconds{}, nil // flowEndSeconds
		case 152:
			return &FieldValueDateTimeMilliseconds{}, nil // flowStartMilliseconds
		case 153:
			return &FieldValueDateTimeMilliseconds{}, nil // flowEndMilliseconds
		case 154:
			return &FieldValueDateTimeMicroseconds{}, nil // flowStartMicroseconds
		case 155:
			return &FieldValueDateTimeMicroseconds{}, nil // flowEndMicroseconds
		case 156:
			return &FieldValueDateTimeNanoseconds{}, nil // flowStartNanoseconds
		case 157:
			return &FieldValueDateTimeNanoseconds{}, nil // flowEndNanoseconds
		case 158:
			return &FieldValueUnsigned32{}, nil // flowStartDeltaMicroseconds
		case 159:
			return &FieldValueUnsigned32{}, nil // flowEndDeltaMicroseconds
		case 160:
			return &FieldValueDateTimeMilliseconds{}, nil // systemInitTimeMilliseconds
		case 161:
			return &FieldValueUnsigned32{}, nil // flowDurationMilliseconds
		case 162:
			return &FieldValueUnsigned32{}, nil // flowDurationMicroseconds
		case 163:
			return &FieldValueUnsigned64{}, nil // observedFlowTotalCount
		case 164:
			return &FieldValueUnsigned64{}, nil // ignoredPacketTotalCount
		case 165:
			return &FieldValueUnsigned64{}, nil // ignoredOctetTotalCount
		case 166:
			return &FieldValueUnsigned64{}, nil // notSentFlowTotalCount
		case 167:
			return &FieldValueUnsigned64{}, nil // notSentPacketTotalCount
		case 168:
			return &FieldValueUnsigned64{}, nil // notSentOctetTotalCount
		case 169:
			return &FieldValueIPv6Address{}, nil // destinationIPv6Prefix
		case 170:
			return &FieldValueIPv6Address{}, nil // sourceIPv6Prefix
		case 171:
			return &FieldValueUnsigned64{}, nil // postOctetTotalCount
		case 172:
			return &FieldValueUnsigned64{}, nil // postPacketTotalCount
		case 173:
			return &FieldValueUnsigned64{}, nil // flowKeyIndicator
		case 174:
			return &FieldValueUnsigned64{}, nil // postMCastPacketTotalCount
		case 175:
			return &FieldValueUnsigned64{}, nil // postMCastOctetTotalCount
		case 176:
			return &FieldValueUnsigned8{}, nil // icmpTypeIPv4
		case 177:
			return &FieldValueUnsigned8{}, nil // icmpCodeIPv4
		case 178:
			return &FieldValueUnsigned8{}, nil // icmpTypeIPv6
		case 179:
			return &FieldValueUnsigned8{}, nil // icmpCodeIPv6
		case 180:
			return &FieldValueUnsigned16{}, nil // udpSourcePort
		case 181:
			return &FieldValueUnsigned16{}, nil // udpDestinationPort
		case 182:
			return &FieldValueUnsigned16{}, nil // tcpSourcePort
		case 183:
			return &FieldValueUnsigned16{}, nil // tcpDestinationPort
		case 184:
			return &FieldValueUnsigned32{}, nil // tcpSequenceNumber
		case 185:
			return &FieldValueUnsigned32{}, nil // tcpAcknowledgementNumber
		case 186:
			return &FieldValueUnsigned16{}, nil // tcpWindowSize
		case 187:
			return &FieldValueUnsigned16{}, nil // tcpUrgentPointer
		case 188:
			return &FieldValueUnsigned8{}, nil // tcpHeaderLength
		case 189:
			return &FieldValueUnsigned8{}, nil // ipHeaderLength
		case 190:
			return &FieldValueUnsigned16{}, nil // totalLengthIPv4
		case 191:
			return &FieldValueUnsigned16{}, nil // payloadLengthIPv6
		case 192:
			return &FieldValueUnsigned8{}, nil // ipTTL
		case 193:
			return &FieldValueUnsigned8{}, nil // nextHeaderIPv6
		case 194:
			return &FieldValueUnsigned32{}, nil // mplsPayloadLength
		case 195:
			return &FieldValueUnsigned8{}, nil // ipDiffServCodePoint
		case 196:
			return &FieldValueUnsigned8{}, nil // ipPrecedence
		case 197:
			return &FieldValueUnsigned8{}, nil // fragmentFlags
		case 198:
			return &FieldValueUnsigned64{}, nil // octetDeltaSumOfSquares
		case 199:
			return &FieldValueUnsigned64{}, nil // octetTotalSumOfSquares
		case 200:
			return &FieldValueUnsigned8{}, nil // mplsTopLabelTTL
		case 201:
			return &FieldValueUnsigned32{}, nil // mplsLabelStackLength
		case 202:
			return &FieldValueUnsigned32{}, nil // mplsLabelStackDepth
		case 203:
			return &FieldValueUnsigned8{}, nil // mplsTopLabelExp
		case 204:
			return &FieldValueUnsigned32{}, nil // ipPayloadLength
		case 205:
			return &FieldValueUnsigned16{}, nil // udpMessageLength
		case 206:
			return &FieldValueUnsigned8{}, nil // isMulticast
		case 207:
			return &FieldValueUnsigned8{}, nil // ipv4IHL
		case 208:
			return &FieldValueUnsigned32{}, nil // ipv4Options
		case 209:
			return &FieldValueUnsigned64{}, nil // tcpOptions
		case 210:
			return &FieldValueOctetArray{}, nil // paddingOctets
		case 211:
			return &FieldValueIPv4Address{}, nil // collectorIPv4Address
		case 212:
			return &FieldValueIPv6Address{}, nil // collectorIPv6Address
		case 213:
			return &FieldValueUnsigned32{}, nil // exportInterface
		case 214:
			return &FieldValueUnsigned8{}, nil // exportProtocolVersion
		case 215:
			return &FieldValueUnsigned8{}, nil // exportTransportProtocol
		case 216:
			return &FieldValueUnsigned16{}, nil // collectorTransportPort
		case 217:
			return &FieldValueUnsigned16{}, nil // exporterTransportPort
		case 218:
			return &FieldValueUnsigned64{}, nil // tcpSynTotalCount
		case 219:
			return &FieldValueUnsigned64{}, nil // tcpFinTotalCount
		case 220:
			return &FieldValueUnsigned64{}, nil // tcpRstTotalCount
		case 221:
			return &FieldValueUnsigned64{}, nil // tcpPshTotalCount
		case 222:
			return &FieldValueUnsigned64{}, nil // tcpAckTotalCount
		case 223:
			return &FieldValueUnsigned64{}, nil // tcpUrgTotalCount
		case 224:
			return &FieldValueUnsigned64{}, nil // ipTotalLength
		case 225:
			return &FieldValueIPv4Address{}, nil // postNATSourceIPv4Address
		case 226:
			return &FieldValueIPv4Address{}, nil // postNATDestinationIPv4Address
		case 227:
			return &FieldValueUnsigned16{}, nil // postNAPTSourceTransportPort
		case 228:
			return &FieldValueUnsigned16{}, nil // postNAPTDestinationTransportPort
		case 229:
			return &FieldValueUnsigned8{}, nil // natOriginatingAddressRealm
		case 230:
			return &FieldValueUnsigned8{}, nil // natEvent
		case 231:
			return &FieldValueUnsigned64{}, nil // initiatorOctets
		case 232:
			return &FieldValueUnsigned64{}, nil // responderOctets
		case 233:
			return &FieldValueUnsigned8{}, nil // firewallEvent
		case 234:
			return &FieldValueUnsigned32{}, nil // ingressVRFID
		case 235:
			return &FieldValueUnsigned32{}, nil // egressVRFID
		case 236:
			return &FieldValueString{}, nil // VRFname
		case 237:
			return &FieldValueUnsigned8{}, nil // postMplsTopLabelExp
		case 238:
			return &FieldValueUnsigned16{}, nil // tcpWindowScale
		case 239:
			return &FieldValueUnsigned8{}, nil // biflowDirection
		case 240:
			return &FieldValueUnsigned8{}, nil // ethernetHeaderLength
		case 241:
			return &FieldValueUnsigned16{}, nil // ethernetPayloadLength
		case 242:
			return &FieldValueUnsigned16{}, nil // ethernetTotalLength
		case 243:
			return &FieldValueUnsigned16{}, nil // dot1qVlanId
		case 244:
			return &FieldValueUnsigned8{}, nil // dot1qPriority
		case 245:
			return &FieldValueUnsigned16{}, nil // dot1qCustomerVlanId
		case 246:
			return &FieldValueUnsigned8{}, nil // dot1qCustomerPriority
		case 247:
			return &FieldValueString{}, nil // metroEvcId
		case 248:
			return &FieldValueUnsigned8{}, nil // metroEvcType
		case 249:
			return &FieldValueUnsigned32{}, nil // pseudoWireId
		case 250:
			return &FieldValueUnsigned16{}, nil // pseudoWireType
		case 251:
			return &FieldValueUnsigned32{}, nil // pseudoWireControlWord
		case 252:
			return &FieldValueUnsigned32{}, nil // ingressPhysicalInterface
		case 253:
			return &FieldValueUnsigned32{}, nil // egressPhysicalInterface
		case 254:
			return &FieldValueUnsigned16{}, nil // postDot1qVlanId
		case 255:
			return &FieldValueUnsigned16{}, nil // postDot1qCustomerVlanId
		case 256:
			return &FieldValueUnsigned16{}, nil // ethernetType
		case 257:
			return &FieldValueUnsigned8{}, nil // postIpPrecedence
		case 258:
			return &FieldValueDateTimeMilliseconds{}, nil // collectionTimeMilliseconds
		case 259:
			return &FieldValueUnsigned16{}, nil // exportSctpStreamId
		case 260:
			return &FieldValueDateTimeSeconds{}, nil // maxExportSeconds
		case 261:
			return &FieldValueDateTimeSeconds{}, nil // maxFlowEndSeconds
		case 262:
			return &FieldValueOctetArray{}, nil // messageMD5Checksum
		case 263:
			return &FieldValueUnsigned8{}, nil // messageScope
		case 264:
			return &FieldValueDateTimeSeconds{}, nil // minExportSeconds
		case 265:
			return &FieldValueDateTimeSeconds{}, nil // minFlowStartSeconds
		case 266:
			return &FieldValueOctetArray{}, nil // opaqueOctets
		case 267:
			return &FieldValueUnsigned8{}, nil // sessionScope
		case 268:
			return &FieldValueDateTimeMicroseconds{}, nil // maxFlowEndMicroseconds
		case 269:
			return &FieldValueDateTimeMilliseconds{}, nil // maxFlowEndMilliseconds
		case 270:
			return &FieldValueDateTimeNanoseconds{}, nil // maxFlowEndNanoseconds
		case 271:
			return &FieldValueDateTimeMicroseconds{}, nil // minFlowStartMicroseconds
		case 272:
			return &FieldValueDateTimeMilliseconds{}, nil // minFlowStartMilliseconds
		case 273:
			return &FieldValueDateTimeNanoseconds{}, nil // minFlowStartNanoseconds
		case 274:
			return &FieldValueOctetArray{}, nil // collectorCertificate
		case 275:
			return &FieldValueOctetArray{}, nil // exporterCertificate
		case 276:
			return &FieldValueBoolean{}, nil // dataRecordsReliability
		case 277:
			return &FieldValueUnsigned8{}, nil // observationPointType
		case 278:
			return &FieldValueUnsigned32{}, nil // newConnectionDeltaCount
		case 279:
			return &FieldValueUnsigned64{}, nil // connectionSumDurationSeconds
		case 280:
			return &FieldValueUnsigned64{}, nil // connectionTransactionId
		case 281:
			return &FieldValueIPv6Address{}, nil // postNATSourceIPv6Address
		case 282:
			return &FieldValueIPv6Address{}, nil // postNATDestinationIPv6Address
		case 283:
			return &FieldValueUnsigned32{}, nil // natPoolId
		case 284:
			return &FieldValueString{}, nil // natPoolName
		case 285:
			return &FieldValueUnsigned16{}, nil // anonymizationFlags
		case 286:
			return &FieldValueUnsigned16{}, nil // anonymizationTechnique
		case 287:
			return &FieldValueUnsigned16{}, nil // informationElementIndex
		case 288:
			return &FieldValueString{}, nil // p2pTechnology
		case 289:
			return &FieldValueString{}, nil // tunnelTechnology
		case 290:
			return &FieldValueString{}, nil // encryptedTechnology
		case 291:
			return &FieldValueBasicList{}, nil // basicList
		case 292:
			return &FieldValueSubTemplateList{}, nil // subTemplateList
		case 293:
			return &FieldValueSubTemplateMultiList{}, nil // subTemplateMultiList
		case 294:
			return &FieldValueUnsigned8{}, nil // bgpValidityState
		case 295:
			return &FieldValueUnsigned32{}, nil // IPSecSPI
		case 296:
			return &FieldValueUnsigned32{}, nil // greKey
		case 297:
			return &FieldValueUnsigned8{}, nil // natType
		case 298:
			return &FieldValueUnsigned64{}, nil // initiatorPackets
		case 299:
			return &FieldValueUnsigned64{}, nil // responderPackets
		case 300:
			return &FieldValueString{}, nil // observationDomainName
		case 301:
			return &FieldValueUnsigned64{}, nil // selectionSequenceId
		case 302:
			return &FieldValueUnsigned64{}, nil // selectorId
		case 303:
			return &FieldValueUnsigned16{}, nil // informationElementId
		case 304:
			return &FieldValueUnsigned16{}, nil // selectorAlgorithm
		case 305:
			return &FieldValueUnsigned32{}, nil // samplingPacketInterval
		case 306:
			return &FieldValueUnsigned32{}, nil // samplingPacketSpace
		case 307:
			return &FieldValueUnsigned32{}, nil // samplingTimeInterval
		case 308:
			return &FieldValueUnsigned32{}, nil // samplingTimeSpace
		case 309:
			return &FieldValueUnsigned32{}, nil // samplingSize
		case 310:
			return &FieldValueUnsigned32{}, nil // samplingPopulation
		case 311:
			return &FieldValueFloat64{}, nil // samplingProbability
		case 312:
			return &FieldValueUnsigned16{}, nil // dataLinkFrameSize
		case 313:
			return &FieldValueOctetArray{}, nil // ipHeaderPacketSection
		case 314:
			return &FieldValueOctetArray{}, nil // ipPayloadPacketSection
		case 315:
			return &FieldValueOctetArray{}, nil // dataLinkFrameSection
		case 316:
			return &FieldValueOctetArray{}, nil // mplsLabelStackSection
		case 317:
			return &FieldValueOctetArray{}, nil // mplsPayloadPacketSection
		case 318:
			return &FieldValueUnsigned64{}, nil // selectorIdTotalPktsObserved
		case 319:
			return &FieldValueUnsigned64{}, nil // selectorIdTotalPktsSelected
		case 320:
			return &FieldValueFloat64{}, nil // absoluteError
		case 321:
			return &FieldValueFloat64{}, nil // relativeError
		case 322:
			return &FieldValueDateTimeSeconds{}, nil // observationTimeSeconds
		case 323:
			return &FieldValueDateTimeMilliseconds{}, nil // observationTimeMilliseconds
		case 324:
			return &FieldValueDateTimeMicroseconds{}, nil // observationTimeMicroseconds
		case 325:
			return &FieldValueDateTimeNanoseconds{}, nil // observationTimeNanoseconds
		case 326:
			return &FieldValueUnsigned64{}, nil // digestHashValue
		case 327:
			return &FieldValueUnsigned64{}, nil // hashIPPayloadOffset
		case 328:
			return &FieldValueUnsigned64{}, nil // hashIPPayloadSize
		case 329:
			return &FieldValueUnsigned64{}, nil // hashOutputRangeMin
		case 330:
			return &FieldValueUnsigned64{}, nil // hashOutputRangeMax
		case 331:
			return &FieldValueUnsigned64{}, nil // hashSelectedRangeMin
		case 332:
			return &FieldValueUnsigned64{}, nil // hashSelectedRangeMax
		case 333:
			return &FieldValueBoolean{}, nil // hashDigestOutput
		case 334:
			return &FieldValueUnsigned64{}, nil // hashInitialiserValue
		case 335:
			return &FieldValueString{}, nil // selectorName
		case 336:
			return &FieldValueFloat64{}, nil // upperCILimit
		case 337:
			return &FieldValueFloat64{}, nil // lowerCILimit
		case 338:
			return &FieldValueFloat64{}, nil // confidenceLevel
		case 339:
			return &FieldValueUnsigned8{}, nil // informationElementDataType
		case 340:
			return &FieldValueString{}, nil // informationElementDescription
		case 341:
			return &FieldValueString{}, nil // informationElementName
		case 342:
			return &FieldValueUnsigned64{}, nil // informationElementRangeBegin
		case 343:
			return &FieldValueUnsigned64{}, nil // informationElementRangeEnd
		case 344:
			return &FieldValueUnsigned8{}, nil // informationElementSemantics
		case 345:
			return &FieldValueUnsigned16{}, nil // informationElementUnits
		case 346:
			return &FieldValueUnsigned32{}, nil // privateEnterpriseNumber
		case 347:
			return &FieldValueOctetArray{}, nil // virtualStationInterfaceId
		case 348:
			return &FieldValueString{}, nil // virtualStationInterfaceName
		case 349:
			return &FieldValueOctetArray{}, nil // virtualStationUUID
		case 350:
			return &FieldValueString{}, nil // virtualStationName
		case 351:
			return &FieldValueUnsigned64{}, nil // layer2SegmentId
		case 352:
			return &FieldValueUnsigned64{}, nil // layer2OctetDeltaCount
		case 353:
			return &FieldValueUnsigned64{}, nil // layer2OctetTotalCount
		case 354:
			return &FieldValueUnsigned64{}, nil // ingressUnicastPacketTotalCount
		case 355:
			return &FieldValueUnsigned64{}, nil // ingressMulticastPacketTotalCount
		case 356:
			return &FieldValueUnsigned64{}, nil // ingressBroadcastPacketTotalCount
		case 357:
			return &FieldValueUnsigned64{}, nil // egressUnicastPacketTotalCount
		case 358:
			return &FieldValueUnsigned64{}, nil // egressBroadcastPacketTotalCount
		case 359:
			return &FieldValueDateTimeMilliseconds{}, nil // monitoringIntervalStartMilliSeconds
		case 360:
			return &FieldValueDateTimeMilliseconds{}, nil // monitoringIntervalEndMilliSeconds
		case 361:
			return &FieldValueUnsigned16{}, nil // portRangeStart
		case 362:
			return &FieldValueUnsigned16{}, nil // portRangeEnd
		case 363:
			return &FieldValueUnsigned16{}, nil // portRangeStepSize
		case 364:
			return &FieldValueUnsigned16{}, nil // portRangeNumPorts
		case 365:
			return &FieldValueMacAddress{}, nil // staMacAddress
		case 366:
			return &FieldValueIPv4Address{}, nil // staIPv4Address
		case 367:
			return &FieldValueMacAddress{}, nil // wtpMacAddress
		case 368:
			return &FieldValueUnsigned32{}, nil // ingressInterfaceType
		case 369:
			return &FieldValueUnsigned32{}, nil // egressInterfaceType
		case 370:
			return &FieldValueUnsigned16{}, nil // rtpSequenceNumber
		case 371:
			return &FieldValueString{}, nil // userName
		case 372:
			return &FieldValueString{}, nil // applicationCategoryName
		case 373:
			return &FieldValueString{}, nil // applicationSubCategoryName
		case 374:
			return &FieldValueString{}, nil // applicationGroupName
		case 375:
			return &FieldValueUnsigned64{}, nil // originalFlowsPresent
		case 376:
			return &FieldValueUnsigned64{}, nil // originalFlowsInitiated
		case 377:
			return &FieldValueUnsigned64{}, nil // originalFlowsCompleted
		case 378:
			return &FieldValueUnsigned64{}, nil // distinctCountOfSourceIPAddress
		case 379:
			return &FieldValueUnsigned64{}, nil // distinctCountOfDestinationIPAddress
		case 380:
			return &FieldValueUnsigned32{}, nil // distinctCountOfSourceIPv4Address
		case 381:
			return &FieldValueUnsigned32{}, nil // distinctCountOfDestinationIPv4Address
		case 382:
			return &FieldValueUnsigned64{}, nil // distinctCountOfSourceIPv6Address
		case 383:
			return &FieldValueUnsigned64{}, nil // distinctCountOfDestinationIPv6Address
		case 384:
			return &FieldValueUnsigned8{}, nil // valueDistributionMethod
		case 385:
			return &FieldValueUnsigned32{}, nil // rfc3550JitterMilliseconds
		case 386:
			return &FieldValueUnsigned32{}, nil // rfc3550JitterMicroseconds
		case 387:
			return &FieldValueUnsigned32{}, nil // rfc3550JitterNanoseconds
		case 388:
			return &FieldValueBoolean{}, nil // dot1qDEI
		case 389:
			return &FieldValueBoolean{}, nil // dot1qCustomerDEI
		case 390:
			return &FieldValueUnsigned16{}, nil // flowSelectorAlgorithm
		case 391:
			return &FieldValueUnsigned64{}, nil // flowSelectedOctetDeltaCount
		case 392:
			return &FieldValueUnsigned64{}, nil // flowSelectedPacketDeltaCount
		case 393:
			return &FieldValueUnsigned64{}, nil // flowSelectedFlowDeltaCount
		case 394:
			return &FieldValueUnsigned64{}, nil // selectorIDTotalFlowsObserved
		case 395:
			return &FieldValueUnsigned64{}, nil // selectorIDTotalFlowsSelected
		case 396:
			return &FieldValueUnsigned64{}, nil // samplingFlowInterval
		case 397:
			return &FieldValueUnsigned64{}, nil // samplingFlowSpacing
		case 398:
			return &FieldValueUnsigned64{}, nil // flowSamplingTimeInterval
		case 399:
			return &FieldValueUnsigned64{}, nil // flowSamplingTimeSpacing
		case 400:
			return &FieldValueUnsigned16{}, nil // hashFlowDomain
		case 401:
			return &FieldValueUnsigned64{}, nil // transportOctetDeltaCount
		case 402:
			return &FieldValueUnsigned64{}, nil // transportPacketDeltaCount
		case 403:
			return &FieldValueIPv4Address{}, nil // originalExporterIPv4Address
		case 404:
			return &FieldValueIPv6Address{}, nil // originalExporterIPv6Address
		case 405:
			return &FieldValueUnsigned32{}, nil // originalObservationDomainId
		case 406:
			return &FieldValueUnsigned32{}, nil // intermediateProcessId
		case 407:
			return &FieldValueUnsigned64{}, nil // ignoredDataRecordTotalCount
		case 408:
			return &FieldValueUnsigned16{}, nil // dataLinkFrameType
		case 409:
			return &FieldValueUnsigned16{}, nil // sectionOffset
		case 410:
			return &FieldValueUnsigned16{}, nil // sectionExportedOctets
		case 411:
			return &FieldValueOctetArray{}, nil // dot1qServiceInstanceTag
		case 412:
			return &FieldValueUnsigned32{}, nil // dot1qServiceInstanceId
		case 413:
			return &FieldValueUnsigned8{}, nil // dot1qServiceInstancePriority
		case 414:
			return &FieldValueMacAddress{}, nil // dot1qCustomerSourceMacAddress
		case 415:
			return &FieldValueMacAddress{}, nil // dot1qCustomerDestinationMacAddress
		case 417:
			return &FieldValueUnsigned64{}, nil // postLayer2OctetDeltaCount
		case 418:
			return &FieldValueUnsigned64{}, nil // postMCastLayer2OctetDeltaCount
		case 420:
			return &FieldValueUnsigned64{}, nil // postLayer2OctetTotalCount
		case 421:
			return &FieldValueUnsigned64{}, nil // postMCastLayer2OctetTotalCount
		case 422:
			return &FieldValueUnsigned64{}, nil // minimumLayer2TotalLength
		case 423:
			return &FieldValueUnsigned64{}, nil // maximumLayer2TotalLength
		case 424:
			return &FieldValueUnsigned64{}, nil // droppedLayer2OctetDeltaCount
		case 425:
			return &FieldValueUnsigned64{}, nil // droppedLayer2OctetTotalCount
		case 426:
			return &FieldValueUnsigned64{}, nil // ignoredLayer2OctetTotalCount
		case 427:
			return &FieldValueUnsigned64{}, nil // notSentLayer2OctetTotalCount
		case 428:
			return &FieldValueUnsigned64{}, nil // layer2OctetDeltaSumOfSquares
		case 429:
			return &FieldValueUnsigned64{}, nil // layer2OctetTotalSumOfSquares
		case 430:
			return &FieldValueUnsigned64{}, nil // layer2FrameDeltaCount
		case 431:
			return &FieldValueUnsigned64{}, nil // layer2FrameTotalCount
		case 432:
			return &FieldValueIPv4Address{}, nil // pseudoWireDestinationIPv4Address
		case 433:
			return &FieldValueUnsigned64{}, nil // ignoredLayer2FrameTotalCount
		case 434:
			return &FieldValueSigned64{}, nil // mibObjectValueInteger
		case 435:
			return &FieldValueOctetArray{}, nil // mibObjectValueOctetString
		case 436:
			return &FieldValueOctetArray{}, nil // mibObjectValueOID
		case 437:
			return &FieldValueOctetArray{}, nil // mibObjectValueBits
		case 438:
			return &FieldValueIPv4Address{}, nil // mibObjectValueIPAddress
		case 439:
			return &FieldValueUnsigned64{}, nil // mibObjectValueCounter
		case 440:
			return &FieldValueUnsigned32{}, nil // mibObjectValueGauge
		case 441:
			return &FieldValueUnsigned32{}, nil // mibObjectValueTimeTicks
		case 442:
			return &FieldValueUnsigned64{}, nil // mibObjectValueUnsigned
		case 443:
			return &FieldValueSubTemplateList{}, nil // mibObjectValueTable
		case 444:
			return &FieldValueSubTemplateList{}, nil // mibObjectValueRow
		case 445:
			return &FieldValueOctetArray{}, nil // mibObjectIdentifier
		case 446:
			return &FieldValueUnsigned32{}, nil // mibSubIdentifier
		case 447:
			return &FieldValueUnsigned64{}, nil // mibIndexIndicator
		case 448:
			return &FieldValueUnsigned8{}, nil // mibCaptureTimeSemantics
		case 449:
			return &FieldValueOctetArray{}, nil // mibContextEngineID
		case 450:
			return &FieldValueString{}, nil // mibContextName
		case 451:
			return &FieldValueString{}, nil // mibObjectName
		case 452:
			return &FieldValueString{}, nil // mibObjectDescription
		case 453:
			return &FieldValueString{}, nil // mibObjectSyntax
		case 454:
			return &FieldValueString{}, nil // mibModuleName
		case 455:
			return &FieldValueString{}, nil // mobileIMSI
		case 456:
			return &FieldValueString{}, nil // mobileMSISDN
		case 457:
			return &FieldValueUnsigned16{}, nil // httpStatusCode
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 8057: // IPFIXColStyle - https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return &FieldValueUnsigned8{}, nil // DNSRCode
		case 2:
			return &FieldValueString{}, nil // DNSName
		case 3:
			return &FieldValueUnsigned16{}, nil // DNSQType
		case 4:
			return &FieldValueUnsigned16{}, nil // DNSClass
		case 5:
			return &FieldValueUnsigned32{}, nil // DNSRRTTL
		case 6:
			return &FieldValueUnsigned16{}, nil // DNSRDataLength
		case 7:
			return &FieldValueString{}, nil // DNSRData
		case 8:
			return &FieldValueUnsigned16{}, nil // DNSPSIZE
		case 9:
			return &FieldValueUnsigned8{}, nil // DNSRDO
		case 10:
			return &FieldValueUnsigned16{}, nil // DNSTransactionID
		case 700:
			return &FieldValueUnsigned8{}, nil // HBType
		case 701:
			return &FieldValueUnsigned8{}, nil // HBDir
		case 702:
			return &FieldValueUnsigned16{}, nil // HBSizeMsg
		case 703:
			return &FieldValueUnsigned16{}, nil // HBSizePayload
		case 800:
			return &FieldValueUnsigned32{}, nil // HTTPRequestMethod
		case 801:
			return &FieldValueString{}, nil // HTTPRequestHost
		case 802:
			return &FieldValueString{}, nil // HTTPRequestURL
		case 803:
			return &FieldValueString{}, nil // HTTPRequestReferer
		case 804:
			return &FieldValueString{}, nil // HTTPRequestAgent
		case 805:
			return &FieldValueUnsigned32{}, nil // HTTPResponseCode
		case 806:
			return &FieldValueString{}, nil // HTTPResponseType
		case 807:
			return &FieldValueUnsigned64{}, nil // HTTPResponseTime
		case 808:
			return &FieldValueString{}, nil // HTTPSHost
		case 809:
			return &FieldValueUnsigned64{}, nil // HTTPSResponseTime
		case 810:
			return &FieldValueUnsigned32{}, nil // SMTPCommands
		case 811:
			return &FieldValueUnsigned32{}, nil // SMTPMailCount
		case 812:
			return &FieldValueUnsigned32{}, nil // SMTPRcptCount
		case 813:
			return &FieldValueString{}, nil // SMTPSender
		case 814:
			return &FieldValueString{}, nil // SMTPRecipient
		case 815:
			return &FieldValueUnsigned32{}, nil // SMTPStatusCodes
		case 816:
			return &FieldValueUnsigned32{}, nil // SMTPCode2XXCount
		case 817:
			return &FieldValueUnsigned32{}, nil // SMTPCode3XXCount
		case 818:
			return &FieldValueUnsigned32{}, nil // SMTPCode4XXCount
		case 819:
			return &FieldValueUnsigned32{}, nil // SMTPCode5XXCount
		case 820:
			return &FieldValueString{}, nil // SMTPDomain
		case 821:
			return &FieldValueString{}, nil // HTTPRequestRange
		case 830:
			return &FieldValueUnsigned32{}, nil // SIPMethod
		case 831:
			return &FieldValueUnsigned32{}, nil // SIPStatusCode
		case 832:
			return &FieldValueString{}, nil // SIPRequestURI
		case 833:
			return &FieldValueString{}, nil // SIPFrom
		case 834:
			return &FieldValueString{}, nil // SIPTo
		case 835:
			return &FieldValueString{}, nil // SIPContact
		case 836:
			return &FieldValueString{}, nil // SIPVia
		case 837:
			return &FieldValueString{}, nil // SIPRoute
		case 838:
			return &FieldValueString{}, nil // SIPRecordRoute
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 16982: // IPFIXColStyle - https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return &FieldValueUnsigned8{}, nil // destinationGeo
		case 100:
			return &FieldValueUnsigned8{}, nil // HTTPUserAgent
		case 101:
			return &FieldValueUnsigned8{}, nil // HTTPMethod
		case 102:
			return &FieldValueString{}, nil // HTTPDomain
		case 103:
			return &FieldValueString{}, nil // HTTPReferer
		case 104:
			return &FieldValueUnsigned8{}, nil // HTTPContentType
		case 105:
			return &FieldValueString{}, nil // HTTPUrl
		case 106:
			return &FieldValueUnsigned16{}, nil // HTTPStatus
		case 107:
			return &FieldValueUnsigned16{}, nil // HTTPHeaderCount
		case 200:
			return &FieldValueUnsigned32{}, nil // appPID
		case 201:
			return &FieldValueString{}, nil // appName
		case 202:
			return &FieldValueUnsigned32{}, nil // appUID
		case 203:
			return &FieldValueUnsigned8{}, nil // appMatchLevel
		case 400:
			return &FieldValueIPv6Address{}, nil // tunnelSrcIPv6
		case 401:
			return &FieldValueIPv6Address{}, nil // tunnelDstIPv6
		case 402:
			return &FieldValueUnsigned16{}, nil // tunnelSrcPort
		case 403:
			return &FieldValueUnsigned16{}, nil // tunnelDstPort
		case 404:
			return &FieldValueUnsigned8{}, nil // tunnelProtocol
		case 405:
			return &FieldValueUnsigned8{}, nil // tunnelType
		case 406:
			return &FieldValueUnsigned8{}, nil // tunnelTCPFlags
		case 407:
			return &FieldValueUnsigned16{}, nil // tunnelICMPcode
		case 408:
			return &FieldValueUnsigned8{}, nil // tunnelTeredoHeaders
		case 409:
			return &FieldValueUnsigned8{}, nil // tunnelTeredoTrailers
		case 410:
			return &FieldValueString{}, nil // tunnelSourceGeo
		case 411:
			return &FieldValueString{}, nil // tunnelDestinationGeo
		case 412:
			return &FieldValueString{}, nil // tunnelOuterSourceGeo
		case 413:
			return &FieldValueString{}, nil // tunnelOuterDestinationGeo
		case 414:
			return &FieldValueUnsigned8{}, nil // tunnelHOPLimit
		case 500:
			return &FieldValueUnsigned32{}, nil // HTTPRequestType
		case 501:
			return &FieldValueString{}, nil // HTTPRequestHost
		case 502:
			return &FieldValueString{}, nil // HTTPRequestURL
		case 503:
			return &FieldValueUnsigned32{}, nil // HTTPRequestAgentID
		case 504:
			return &FieldValueString{}, nil // HTTPRequestAgent
		case 505:
			return &FieldValueString{}, nil // HTTPRequestReferer
		case 506:
			return &FieldValueUnsigned32{}, nil // HTTPResponseCode
		case 507:
			return &FieldValueString{}, nil // HTTPResponseType
		case 508:
			return &FieldValueUnsigned32{}, nil // HTTPResponseTime
		case 800:
			return &FieldValueUnsigned16{}, nil // tlsCliVer
		case 801:
			return &FieldValueUnsigned16{}, nil // tlsSerVer
		case 802:
			return &FieldValueUnsigned16{}, nil // tlsSerCips
		case 803:
			return &FieldValueOctetArray{}, nil // tlsCliCips
		case 804:
			return &FieldValueString{}, nil // tlsCertSubject
		case 805:
			return &FieldValueString{}, nil // tlsCertIssuer
		case 806:
			return &FieldValueUnsigned64{}, nil // tlsCertNotBefore
		case 807:
			return &FieldValueUnsigned64{}, nil // tlsCerNotAfter
		case 808:
			return &FieldValueUnsigned16{}, nil // tlsPkeyLength
		case 809:
			return &FieldValueUnsigned32{}, nil // tlsPkeyExponent
		case 810:
			return &FieldValueString{}, nil // tlsPkeyAlgorithm
		case 811:
			return &FieldValueString{}, nil // tlsServerName
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 35632: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 180:
			return &FieldValueString{}, nil // HTTPUrl
		case 187:
			return &FieldValueString{}, nil // HTTPHost
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 39499: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return &FieldValueString{}, nil // HTTPRequestHost
		case 2:
			return &FieldValueString{}, nil // HTTPRequestURL
		case 3:
			return &FieldValueString{}, nil // HTTPRequestReferer
		case 4:
			return &FieldValueUnsigned32{}, nil // HTTPRequestType
		case 10:
			return &FieldValueString{}, nil // HTTPResponseType
		case 12:
			return &FieldValueUnsigned32{}, nil // HTTPResponseCode
		case 20:
			return &FieldValueString{}, nil // HTTPRequestAgent
		case 21:
			return &FieldValueUnsigned32{}, nil // HTTPRequestAgentID
		case 22:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentOS
		case 23:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentOSMajor
		case 24:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentOSMinor
		case 25:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentOSBuild
		case 26:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentApp
		case 27:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentAppMajor
		case 28:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentAppMinor
		case 29:
			return &FieldValueUnsigned16{}, nil // HTTPRequestAgentAppBuild
		case 32:
			return &FieldValueUnsigned8{}, nil // voipPacketType
		case 33:
			return &FieldValueString{}, nil // sipCallId
		case 34:
			return &FieldValueString{}, nil // sipCallingParty
		case 35:
			return &FieldValueString{}, nil // sipCalledParty
		case 36:
			return &FieldValueString{}, nil // sipVia
		case 37:
			return &FieldValueDateTimeNanoseconds{}, nil // sipInviteRingingTime
		case 38:
			return &FieldValueDateTimeNanoseconds{}, nil // sipOkTime
		case 39:
			return &FieldValueDateTimeNanoseconds{}, nil // sipByeTime
		case 40:
			return &FieldValueIPv4Address{}, nil // sipRtpIp4
		case 41:
			return &FieldValueIPv6Address{}, nil // sipRtpIp6
		case 42:
			return &FieldValueUnsigned16{}, nil // sipRtpAudio
		case 43:
			return &FieldValueUnsigned16{}, nil // sipRtpVideo
		case 44:
			return &FieldValueUnsigned64{}, nil // sipStats
		case 45:
			return &FieldValueUnsigned8{}, nil // rtpCodec
		case 46:
			return &FieldValueSigned32{}, nil // rtpJitter
		case 47:
			return &FieldValueUnsigned32{}, nil // rtcpLost
		case 48:
			return &FieldValueUnsigned64{}, nil // rtcpPackets
		case 49:
			return &FieldValueUnsigned64{}, nil // rtcpOctets
		case 50:
			return &FieldValueUnsigned8{}, nil // rtcpSourceCount
		case 51:
			return &FieldValueString{}, nil // sipUserAgent
		case 52:
			return &FieldValueString{}, nil // sipRequestUri
		case 53:
			return &FieldValueString{}, nil // sipCSeq
		case 61:
			return &FieldValueUnsigned32{}, nil // NPMJitterDev
		case 62:
			return &FieldValueUnsigned32{}, nil // NPMJitterAvg
		case 63:
			return &FieldValueUnsigned32{}, nil // NPMJitterMin
		case 64:
			return &FieldValueUnsigned32{}, nil // NPMJitterMax
		case 65:
			return &FieldValueUnsigned32{}, nil // NPMDelayDev
		case 66:
			return &FieldValueUnsigned32{}, nil // NPMDelayAvg
		case 67:
			return &FieldValueUnsigned32{}, nil // NPMDelayMin
		case 68:
			return &FieldValueUnsigned32{}, nil // NPMDelayMax
		case 69:
			return &FieldValueUnsigned32{}, nil // NPMRoundTripTime
		case 70:
			return &FieldValueUnsigned32{}, nil // NPMServerResponseTime
		case 71:
			return &FieldValueUnsigned32{}, nil // NPMTCPRetransmission
		case 72:
			return &FieldValueUnsigned32{}, nil // NPMTCPOutOfOrder
		case 110:
			return &FieldValueUnsigned16{}, nil // DNSID
		case 111:
			return &FieldValueUnsigned16{}, nil // DNSFlagsCodes
		case 112:
			return &FieldValueUnsigned16{}, nil // DNSQuestionCount
		case 113:
			return &FieldValueUnsigned16{}, nil // DNSAnswRecCount
		case 114:
			return &FieldValueUnsigned16{}, nil // DNSAuthRecCount
		case 115:
			return &FieldValueUnsigned16{}, nil // DNSAddtRecCount
		case 116:
			return &FieldValueString{}, nil // DNSCRRName
		case 117:
			return &FieldValueUnsigned16{}, nil // DNSCRRType
		case 118:
			return &FieldValueUnsigned16{}, nil // DNSCRRClass
		case 119:
			return &FieldValueUnsigned32{}, nil // DNSCRRTTL
		case 120:
			return &FieldValueString{}, nil // DNSCRRRDATA
		case 121:
			return &FieldValueString{}, nil // DNSQNAME
		case 122:
			return &FieldValueUnsigned16{}, nil // DNSQTYPE
		case 123:
			return &FieldValueUnsigned16{}, nil // DNSQCLASS
		case 124:
			return &FieldValueUnsigned16{}, nil // DNSCRRRDATALen
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 44913: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 10:
			return &FieldValueUnsigned16{}, nil // origSourceTransportPort
		case 11:
			return &FieldValueIPv4Address{}, nil // origSourceIPv4Address
		case 12:
			return &FieldValueUnsigned16{}, nil // origDestinationTransportPort
		case 13:
			return &FieldValueIPv4Address{}, nil // origDestinationIPv4Address
		case 14:
			return &FieldValueIPv6Address{}, nil // origSourceIPv6Address
		case 15:
			return &FieldValueIPv6Address{}, nil // origDestinationIPv6Address
		case 20:
			return &FieldValueString{}, nil // HTTPRequestHost
		case 21:
			return &FieldValueString{}, nil // HTTPRequestURL
		case 22:
			return &FieldValueString{}, nil // HTTPRequestUserAgent
		case 12345:
			return &FieldValueString{}, nil // Unknown
		default:
			return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	default:
		return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
	}
}

// FieldLengthByID returns the default length that matches the enterprise id and element id
func FieldLengthByID(enterpriseid int, elementid int) (uint16, error) {
	switch enterpriseid {

	case 0: // IANA - https://www.ietf.org/assignments/ipfix/ipfix.xml

		switch elementid {
		case 1:
			return 8, nil // octetDeltaCount
		case 2:
			return 8, nil // packetDeltaCount
		case 3:
			return 8, nil // deltaFlowCount
		case 4:
			return 1, nil // protocolIdentifier
		case 5:
			return 1, nil // ipClassOfService
		case 6:
			return 2, nil // tcpControlBits
		case 7:
			return 2, nil // sourceTransportPort
		case 8:
			return 4, nil // sourceIPv4Address
		case 9:
			return 1, nil // sourceIPv4PrefixLength
		case 10:
			return 4, nil // ingressInterface
		case 11:
			return 2, nil // destinationTransportPort
		case 12:
			return 4, nil // destinationIPv4Address
		case 13:
			return 1, nil // destinationIPv4PrefixLength
		case 14:
			return 4, nil // egressInterface
		case 15:
			return 4, nil // ipNextHopIPv4Address
		case 16:
			return 4, nil // bgpSourceAsNumber
		case 17:
			return 4, nil // bgpDestinationAsNumber
		case 18:
			return 4, nil // bgpNextHopIPv4Address
		case 19:
			return 8, nil // postMCastPacketDeltaCount
		case 20:
			return 8, nil // postMCastOctetDeltaCount
		case 21:
			return 4, nil // flowEndSysUpTime
		case 22:
			return 4, nil // flowStartSysUpTime
		case 23:
			return 8, nil // postOctetDeltaCount
		case 24:
			return 8, nil // postPacketDeltaCount
		case 25:
			return 8, nil // minimumIpTotalLength
		case 26:
			return 8, nil // maximumIpTotalLength
		case 27:
			return 16, nil // sourceIPv6Address
		case 28:
			return 16, nil // destinationIPv6Address
		case 29:
			return 1, nil // sourceIPv6PrefixLength
		case 30:
			return 1, nil // destinationIPv6PrefixLength
		case 31:
			return 4, nil // flowLabelIPv6
		case 32:
			return 2, nil // icmpTypeCodeIPv4
		case 33:
			return 1, nil // igmpType
		case 34:
			return 4, nil // samplingInterval
		case 35:
			return 1, nil // samplingAlgorithm
		case 36:
			return 2, nil // flowActiveTimeout
		case 37:
			return 2, nil // flowIdleTimeout
		case 38:
			return 1, nil // engineType
		case 39:
			return 1, nil // engineId
		case 40:
			return 8, nil // exportedOctetTotalCount
		case 41:
			return 8, nil // exportedMessageTotalCount
		case 42:
			return 8, nil // exportedFlowRecordTotalCount
		case 43:
			return 4, nil // ipv4RouterSc
		case 44:
			return 4, nil // sourceIPv4Prefix
		case 45:
			return 4, nil // destinationIPv4Prefix
		case 46:
			return 1, nil // mplsTopLabelType
		case 47:
			return 4, nil // mplsTopLabelIPv4Address
		case 48:
			return 1, nil // samplerId
		case 49:
			return 1, nil // samplerMode
		case 50:
			return 4, nil // samplerRandomInterval
		case 51:
			return 1, nil // classId
		case 52:
			return 1, nil // minimumTTL
		case 53:
			return 1, nil // maximumTTL
		case 54:
			return 4, nil // fragmentIdentification
		case 55:
			return 1, nil // postIpClassOfService
		case 56:
			return 6, nil // sourceMacAddress
		case 57:
			return 6, nil // postDestinationMacAddress
		case 58:
			return 2, nil // vlanId
		case 59:
			return 2, nil // postVlanId
		case 60:
			return 1, nil // ipVersion
		case 61:
			return 1, nil // flowDirection
		case 62:
			return 16, nil // ipNextHopIPv6Address
		case 63:
			return 16, nil // bgpNextHopIPv6Address
		case 64:
			return 4, nil // ipv6ExtensionHeaders
		case 70:
			return 65535, nil // mplsTopLabelStackSection
		case 71:
			return 65535, nil // mplsLabelStackSection2
		case 72:
			return 65535, nil // mplsLabelStackSection3
		case 73:
			return 65535, nil // mplsLabelStackSection4
		case 74:
			return 65535, nil // mplsLabelStackSection5
		case 75:
			return 65535, nil // mplsLabelStackSection6
		case 76:
			return 65535, nil // mplsLabelStackSection7
		case 77:
			return 65535, nil // mplsLabelStackSection8
		case 78:
			return 65535, nil // mplsLabelStackSection9
		case 79:
			return 65535, nil // mplsLabelStackSection10
		case 80:
			return 6, nil // destinationMacAddress
		case 81:
			return 6, nil // postSourceMacAddress
		case 82:
			return 65535, nil // interfaceName
		case 83:
			return 65535, nil // interfaceDescription
		case 84:
			return 65535, nil // samplerName
		case 85:
			return 8, nil // octetTotalCount
		case 86:
			return 8, nil // packetTotalCount
		case 87:
			return 4, nil // flagsAndSamplerId
		case 88:
			return 2, nil // fragmentOffset
		case 89:
			return 4, nil // forwardingStatus
		case 90:
			return 65535, nil // mplsVpnRouteDistinguisher
		case 91:
			return 1, nil // mplsTopLabelPrefixLength
		case 92:
			return 4, nil // srcTrafficIndex
		case 93:
			return 4, nil // dstTrafficIndex
		case 94:
			return 65535, nil // applicationDescription
		case 95:
			return 65535, nil // applicationId
		case 96:
			return 65535, nil // applicationName
		case 98:
			return 1, nil // postIpDiffServCodePoint
		case 99:
			return 4, nil // multicastReplicationFactor
		case 100:
			return 65535, nil // className
		case 101:
			return 1, nil // classificationEngineId
		case 102:
			return 2, nil // layer2packetSectionOffset
		case 103:
			return 2, nil // layer2packetSectionSize
		case 104:
			return 65535, nil // layer2packetSectionData
		case 128:
			return 4, nil // bgpNextAdjacentAsNumber
		case 129:
			return 4, nil // bgpPrevAdjacentAsNumber
		case 130:
			return 4, nil // exporterIPv4Address
		case 131:
			return 16, nil // exporterIPv6Address
		case 132:
			return 8, nil // droppedOctetDeltaCount
		case 133:
			return 8, nil // droppedPacketDeltaCount
		case 134:
			return 8, nil // droppedOctetTotalCount
		case 135:
			return 8, nil // droppedPacketTotalCount
		case 136:
			return 1, nil // flowEndReason
		case 137:
			return 8, nil // commonPropertiesId
		case 138:
			return 8, nil // observationPointId
		case 139:
			return 2, nil // icmpTypeCodeIPv6
		case 140:
			return 16, nil // mplsTopLabelIPv6Address
		case 141:
			return 4, nil // lineCardId
		case 142:
			return 4, nil // portId
		case 143:
			return 4, nil // meteringProcessId
		case 144:
			return 4, nil // exportingProcessId
		case 145:
			return 2, nil // templateId
		case 146:
			return 1, nil // wlanChannelId
		case 147:
			return 65535, nil // wlanSSID
		case 148:
			return 8, nil // flowId
		case 149:
			return 4, nil // observationDomainId
		case 150:
			return 4, nil // flowStartSeconds
		case 151:
			return 4, nil // flowEndSeconds
		case 152:
			return 4, nil // flowStartMilliseconds
		case 153:
			return 4, nil // flowEndMilliseconds
		case 154:
			return 4, nil // flowStartMicroseconds
		case 155:
			return 4, nil // flowEndMicroseconds
		case 156:
			return 4, nil // flowStartNanoseconds
		case 157:
			return 4, nil // flowEndNanoseconds
		case 158:
			return 4, nil // flowStartDeltaMicroseconds
		case 159:
			return 4, nil // flowEndDeltaMicroseconds
		case 160:
			return 4, nil // systemInitTimeMilliseconds
		case 161:
			return 4, nil // flowDurationMilliseconds
		case 162:
			return 4, nil // flowDurationMicroseconds
		case 163:
			return 8, nil // observedFlowTotalCount
		case 164:
			return 8, nil // ignoredPacketTotalCount
		case 165:
			return 8, nil // ignoredOctetTotalCount
		case 166:
			return 8, nil // notSentFlowTotalCount
		case 167:
			return 8, nil // notSentPacketTotalCount
		case 168:
			return 8, nil // notSentOctetTotalCount
		case 169:
			return 16, nil // destinationIPv6Prefix
		case 170:
			return 16, nil // sourceIPv6Prefix
		case 171:
			return 8, nil // postOctetTotalCount
		case 172:
			return 8, nil // postPacketTotalCount
		case 173:
			return 8, nil // flowKeyIndicator
		case 174:
			return 8, nil // postMCastPacketTotalCount
		case 175:
			return 8, nil // postMCastOctetTotalCount
		case 176:
			return 1, nil // icmpTypeIPv4
		case 177:
			return 1, nil // icmpCodeIPv4
		case 178:
			return 1, nil // icmpTypeIPv6
		case 179:
			return 1, nil // icmpCodeIPv6
		case 180:
			return 2, nil // udpSourcePort
		case 181:
			return 2, nil // udpDestinationPort
		case 182:
			return 2, nil // tcpSourcePort
		case 183:
			return 2, nil // tcpDestinationPort
		case 184:
			return 4, nil // tcpSequenceNumber
		case 185:
			return 4, nil // tcpAcknowledgementNumber
		case 186:
			return 2, nil // tcpWindowSize
		case 187:
			return 2, nil // tcpUrgentPointer
		case 188:
			return 1, nil // tcpHeaderLength
		case 189:
			return 1, nil // ipHeaderLength
		case 190:
			return 2, nil // totalLengthIPv4
		case 191:
			return 2, nil // payloadLengthIPv6
		case 192:
			return 1, nil // ipTTL
		case 193:
			return 1, nil // nextHeaderIPv6
		case 194:
			return 4, nil // mplsPayloadLength
		case 195:
			return 1, nil // ipDiffServCodePoint
		case 196:
			return 1, nil // ipPrecedence
		case 197:
			return 1, nil // fragmentFlags
		case 198:
			return 8, nil // octetDeltaSumOfSquares
		case 199:
			return 8, nil // octetTotalSumOfSquares
		case 200:
			return 1, nil // mplsTopLabelTTL
		case 201:
			return 4, nil // mplsLabelStackLength
		case 202:
			return 4, nil // mplsLabelStackDepth
		case 203:
			return 1, nil // mplsTopLabelExp
		case 204:
			return 4, nil // ipPayloadLength
		case 205:
			return 2, nil // udpMessageLength
		case 206:
			return 1, nil // isMulticast
		case 207:
			return 1, nil // ipv4IHL
		case 208:
			return 4, nil // ipv4Options
		case 209:
			return 8, nil // tcpOptions
		case 210:
			return 65535, nil // paddingOctets
		case 211:
			return 4, nil // collectorIPv4Address
		case 212:
			return 16, nil // collectorIPv6Address
		case 213:
			return 4, nil // exportInterface
		case 214:
			return 1, nil // exportProtocolVersion
		case 215:
			return 1, nil // exportTransportProtocol
		case 216:
			return 2, nil // collectorTransportPort
		case 217:
			return 2, nil // exporterTransportPort
		case 218:
			return 8, nil // tcpSynTotalCount
		case 219:
			return 8, nil // tcpFinTotalCount
		case 220:
			return 8, nil // tcpRstTotalCount
		case 221:
			return 8, nil // tcpPshTotalCount
		case 222:
			return 8, nil // tcpAckTotalCount
		case 223:
			return 8, nil // tcpUrgTotalCount
		case 224:
			return 8, nil // ipTotalLength
		case 225:
			return 4, nil // postNATSourceIPv4Address
		case 226:
			return 4, nil // postNATDestinationIPv4Address
		case 227:
			return 2, nil // postNAPTSourceTransportPort
		case 228:
			return 2, nil // postNAPTDestinationTransportPort
		case 229:
			return 1, nil // natOriginatingAddressRealm
		case 230:
			return 1, nil // natEvent
		case 231:
			return 8, nil // initiatorOctets
		case 232:
			return 8, nil // responderOctets
		case 233:
			return 1, nil // firewallEvent
		case 234:
			return 4, nil // ingressVRFID
		case 235:
			return 4, nil // egressVRFID
		case 236:
			return 65535, nil // VRFname
		case 237:
			return 1, nil // postMplsTopLabelExp
		case 238:
			return 2, nil // tcpWindowScale
		case 239:
			return 1, nil // biflowDirection
		case 240:
			return 1, nil // ethernetHeaderLength
		case 241:
			return 2, nil // ethernetPayloadLength
		case 242:
			return 2, nil // ethernetTotalLength
		case 243:
			return 2, nil // dot1qVlanId
		case 244:
			return 1, nil // dot1qPriority
		case 245:
			return 2, nil // dot1qCustomerVlanId
		case 246:
			return 1, nil // dot1qCustomerPriority
		case 247:
			return 65535, nil // metroEvcId
		case 248:
			return 1, nil // metroEvcType
		case 249:
			return 4, nil // pseudoWireId
		case 250:
			return 2, nil // pseudoWireType
		case 251:
			return 4, nil // pseudoWireControlWord
		case 252:
			return 4, nil // ingressPhysicalInterface
		case 253:
			return 4, nil // egressPhysicalInterface
		case 254:
			return 2, nil // postDot1qVlanId
		case 255:
			return 2, nil // postDot1qCustomerVlanId
		case 256:
			return 2, nil // ethernetType
		case 257:
			return 1, nil // postIpPrecedence
		case 258:
			return 4, nil // collectionTimeMilliseconds
		case 259:
			return 2, nil // exportSctpStreamId
		case 260:
			return 4, nil // maxExportSeconds
		case 261:
			return 4, nil // maxFlowEndSeconds
		case 262:
			return 65535, nil // messageMD5Checksum
		case 263:
			return 1, nil // messageScope
		case 264:
			return 4, nil // minExportSeconds
		case 265:
			return 4, nil // minFlowStartSeconds
		case 266:
			return 65535, nil // opaqueOctets
		case 267:
			return 1, nil // sessionScope
		case 268:
			return 4, nil // maxFlowEndMicroseconds
		case 269:
			return 4, nil // maxFlowEndMilliseconds
		case 270:
			return 4, nil // maxFlowEndNanoseconds
		case 271:
			return 4, nil // minFlowStartMicroseconds
		case 272:
			return 4, nil // minFlowStartMilliseconds
		case 273:
			return 4, nil // minFlowStartNanoseconds
		case 274:
			return 65535, nil // collectorCertificate
		case 275:
			return 65535, nil // exporterCertificate
		case 276:
			return 1, nil // dataRecordsReliability
		case 277:
			return 1, nil // observationPointType
		case 278:
			return 4, nil // newConnectionDeltaCount
		case 279:
			return 8, nil // connectionSumDurationSeconds
		case 280:
			return 8, nil // connectionTransactionId
		case 281:
			return 16, nil // postNATSourceIPv6Address
		case 282:
			return 16, nil // postNATDestinationIPv6Address
		case 283:
			return 4, nil // natPoolId
		case 284:
			return 65535, nil // natPoolName
		case 285:
			return 2, nil // anonymizationFlags
		case 286:
			return 2, nil // anonymizationTechnique
		case 287:
			return 2, nil // informationElementIndex
		case 288:
			return 65535, nil // p2pTechnology
		case 289:
			return 65535, nil // tunnelTechnology
		case 290:
			return 65535, nil // encryptedTechnology
		case 291:
			return 65535, nil // basicList
		case 292:
			return 65535, nil // subTemplateList
		case 293:
			return 65535, nil // subTemplateMultiList
		case 294:
			return 1, nil // bgpValidityState
		case 295:
			return 4, nil // IPSecSPI
		case 296:
			return 4, nil // greKey
		case 297:
			return 1, nil // natType
		case 298:
			return 8, nil // initiatorPackets
		case 299:
			return 8, nil // responderPackets
		case 300:
			return 65535, nil // observationDomainName
		case 301:
			return 8, nil // selectionSequenceId
		case 302:
			return 8, nil // selectorId
		case 303:
			return 2, nil // informationElementId
		case 304:
			return 2, nil // selectorAlgorithm
		case 305:
			return 4, nil // samplingPacketInterval
		case 306:
			return 4, nil // samplingPacketSpace
		case 307:
			return 4, nil // samplingTimeInterval
		case 308:
			return 4, nil // samplingTimeSpace
		case 309:
			return 4, nil // samplingSize
		case 310:
			return 4, nil // samplingPopulation
		case 311:
			return 8, nil // samplingProbability
		case 312:
			return 2, nil // dataLinkFrameSize
		case 313:
			return 65535, nil // ipHeaderPacketSection
		case 314:
			return 65535, nil // ipPayloadPacketSection
		case 315:
			return 65535, nil // dataLinkFrameSection
		case 316:
			return 65535, nil // mplsLabelStackSection
		case 317:
			return 65535, nil // mplsPayloadPacketSection
		case 318:
			return 8, nil // selectorIdTotalPktsObserved
		case 319:
			return 8, nil // selectorIdTotalPktsSelected
		case 320:
			return 8, nil // absoluteError
		case 321:
			return 8, nil // relativeError
		case 322:
			return 4, nil // observationTimeSeconds
		case 323:
			return 4, nil // observationTimeMilliseconds
		case 324:
			return 4, nil // observationTimeMicroseconds
		case 325:
			return 4, nil // observationTimeNanoseconds
		case 326:
			return 8, nil // digestHashValue
		case 327:
			return 8, nil // hashIPPayloadOffset
		case 328:
			return 8, nil // hashIPPayloadSize
		case 329:
			return 8, nil // hashOutputRangeMin
		case 330:
			return 8, nil // hashOutputRangeMax
		case 331:
			return 8, nil // hashSelectedRangeMin
		case 332:
			return 8, nil // hashSelectedRangeMax
		case 333:
			return 1, nil // hashDigestOutput
		case 334:
			return 8, nil // hashInitialiserValue
		case 335:
			return 65535, nil // selectorName
		case 336:
			return 8, nil // upperCILimit
		case 337:
			return 8, nil // lowerCILimit
		case 338:
			return 8, nil // confidenceLevel
		case 339:
			return 1, nil // informationElementDataType
		case 340:
			return 65535, nil // informationElementDescription
		case 341:
			return 65535, nil // informationElementName
		case 342:
			return 8, nil // informationElementRangeBegin
		case 343:
			return 8, nil // informationElementRangeEnd
		case 344:
			return 1, nil // informationElementSemantics
		case 345:
			return 2, nil // informationElementUnits
		case 346:
			return 4, nil // privateEnterpriseNumber
		case 347:
			return 65535, nil // virtualStationInterfaceId
		case 348:
			return 65535, nil // virtualStationInterfaceName
		case 349:
			return 65535, nil // virtualStationUUID
		case 350:
			return 65535, nil // virtualStationName
		case 351:
			return 8, nil // layer2SegmentId
		case 352:
			return 8, nil // layer2OctetDeltaCount
		case 353:
			return 8, nil // layer2OctetTotalCount
		case 354:
			return 8, nil // ingressUnicastPacketTotalCount
		case 355:
			return 8, nil // ingressMulticastPacketTotalCount
		case 356:
			return 8, nil // ingressBroadcastPacketTotalCount
		case 357:
			return 8, nil // egressUnicastPacketTotalCount
		case 358:
			return 8, nil // egressBroadcastPacketTotalCount
		case 359:
			return 4, nil // monitoringIntervalStartMilliSeconds
		case 360:
			return 4, nil // monitoringIntervalEndMilliSeconds
		case 361:
			return 2, nil // portRangeStart
		case 362:
			return 2, nil // portRangeEnd
		case 363:
			return 2, nil // portRangeStepSize
		case 364:
			return 2, nil // portRangeNumPorts
		case 365:
			return 6, nil // staMacAddress
		case 366:
			return 4, nil // staIPv4Address
		case 367:
			return 6, nil // wtpMacAddress
		case 368:
			return 4, nil // ingressInterfaceType
		case 369:
			return 4, nil // egressInterfaceType
		case 370:
			return 2, nil // rtpSequenceNumber
		case 371:
			return 65535, nil // userName
		case 372:
			return 65535, nil // applicationCategoryName
		case 373:
			return 65535, nil // applicationSubCategoryName
		case 374:
			return 65535, nil // applicationGroupName
		case 375:
			return 8, nil // originalFlowsPresent
		case 376:
			return 8, nil // originalFlowsInitiated
		case 377:
			return 8, nil // originalFlowsCompleted
		case 378:
			return 8, nil // distinctCountOfSourceIPAddress
		case 379:
			return 8, nil // distinctCountOfDestinationIPAddress
		case 380:
			return 4, nil // distinctCountOfSourceIPv4Address
		case 381:
			return 4, nil // distinctCountOfDestinationIPv4Address
		case 382:
			return 8, nil // distinctCountOfSourceIPv6Address
		case 383:
			return 8, nil // distinctCountOfDestinationIPv6Address
		case 384:
			return 1, nil // valueDistributionMethod
		case 385:
			return 4, nil // rfc3550JitterMilliseconds
		case 386:
			return 4, nil // rfc3550JitterMicroseconds
		case 387:
			return 4, nil // rfc3550JitterNanoseconds
		case 388:
			return 1, nil // dot1qDEI
		case 389:
			return 1, nil // dot1qCustomerDEI
		case 390:
			return 2, nil // flowSelectorAlgorithm
		case 391:
			return 8, nil // flowSelectedOctetDeltaCount
		case 392:
			return 8, nil // flowSelectedPacketDeltaCount
		case 393:
			return 8, nil // flowSelectedFlowDeltaCount
		case 394:
			return 8, nil // selectorIDTotalFlowsObserved
		case 395:
			return 8, nil // selectorIDTotalFlowsSelected
		case 396:
			return 8, nil // samplingFlowInterval
		case 397:
			return 8, nil // samplingFlowSpacing
		case 398:
			return 8, nil // flowSamplingTimeInterval
		case 399:
			return 8, nil // flowSamplingTimeSpacing
		case 400:
			return 2, nil // hashFlowDomain
		case 401:
			return 8, nil // transportOctetDeltaCount
		case 402:
			return 8, nil // transportPacketDeltaCount
		case 403:
			return 4, nil // originalExporterIPv4Address
		case 404:
			return 16, nil // originalExporterIPv6Address
		case 405:
			return 4, nil // originalObservationDomainId
		case 406:
			return 4, nil // intermediateProcessId
		case 407:
			return 8, nil // ignoredDataRecordTotalCount
		case 408:
			return 2, nil // dataLinkFrameType
		case 409:
			return 2, nil // sectionOffset
		case 410:
			return 2, nil // sectionExportedOctets
		case 411:
			return 65535, nil // dot1qServiceInstanceTag
		case 412:
			return 4, nil // dot1qServiceInstanceId
		case 413:
			return 1, nil // dot1qServiceInstancePriority
		case 414:
			return 6, nil // dot1qCustomerSourceMacAddress
		case 415:
			return 6, nil // dot1qCustomerDestinationMacAddress
		case 417:
			return 8, nil // postLayer2OctetDeltaCount
		case 418:
			return 8, nil // postMCastLayer2OctetDeltaCount
		case 420:
			return 8, nil // postLayer2OctetTotalCount
		case 421:
			return 8, nil // postMCastLayer2OctetTotalCount
		case 422:
			return 8, nil // minimumLayer2TotalLength
		case 423:
			return 8, nil // maximumLayer2TotalLength
		case 424:
			return 8, nil // droppedLayer2OctetDeltaCount
		case 425:
			return 8, nil // droppedLayer2OctetTotalCount
		case 426:
			return 8, nil // ignoredLayer2OctetTotalCount
		case 427:
			return 8, nil // notSentLayer2OctetTotalCount
		case 428:
			return 8, nil // layer2OctetDeltaSumOfSquares
		case 429:
			return 8, nil // layer2OctetTotalSumOfSquares
		case 430:
			return 8, nil // layer2FrameDeltaCount
		case 431:
			return 8, nil // layer2FrameTotalCount
		case 432:
			return 4, nil // pseudoWireDestinationIPv4Address
		case 433:
			return 8, nil // ignoredLayer2FrameTotalCount
		case 434:
			return 8, nil // mibObjectValueInteger
		case 435:
			return 65535, nil // mibObjectValueOctetString
		case 436:
			return 65535, nil // mibObjectValueOID
		case 437:
			return 65535, nil // mibObjectValueBits
		case 438:
			return 4, nil // mibObjectValueIPAddress
		case 439:
			return 8, nil // mibObjectValueCounter
		case 440:
			return 4, nil // mibObjectValueGauge
		case 441:
			return 4, nil // mibObjectValueTimeTicks
		case 442:
			return 8, nil // mibObjectValueUnsigned
		case 443:
			return 65535, nil // mibObjectValueTable
		case 444:
			return 65535, nil // mibObjectValueRow
		case 445:
			return 65535, nil // mibObjectIdentifier
		case 446:
			return 4, nil // mibSubIdentifier
		case 447:
			return 8, nil // mibIndexIndicator
		case 448:
			return 1, nil // mibCaptureTimeSemantics
		case 449:
			return 65535, nil // mibContextEngineID
		case 450:
			return 65535, nil // mibContextName
		case 451:
			return 65535, nil // mibObjectName
		case 452:
			return 65535, nil // mibObjectDescription
		case 453:
			return 65535, nil // mibObjectSyntax
		case 454:
			return 65535, nil // mibModuleName
		case 455:
			return 65535, nil // mobileIMSI
		case 456:
			return 65535, nil // mobileMSISDN
		case 457:
			return 2, nil // httpStatusCode
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 8057: // IPFIXColStyle - https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return 1, nil // DNSRCode
		case 2:
			return 65535, nil // DNSName
		case 3:
			return 2, nil // DNSQType
		case 4:
			return 2, nil // DNSClass
		case 5:
			return 4, nil // DNSRRTTL
		case 6:
			return 2, nil // DNSRDataLength
		case 7:
			return 65535, nil // DNSRData
		case 8:
			return 2, nil // DNSPSIZE
		case 9:
			return 1, nil // DNSRDO
		case 10:
			return 2, nil // DNSTransactionID
		case 700:
			return 1, nil // HBType
		case 701:
			return 1, nil // HBDir
		case 702:
			return 2, nil // HBSizeMsg
		case 703:
			return 2, nil // HBSizePayload
		case 800:
			return 4, nil // HTTPRequestMethod
		case 801:
			return 65535, nil // HTTPRequestHost
		case 802:
			return 65535, nil // HTTPRequestURL
		case 803:
			return 65535, nil // HTTPRequestReferer
		case 804:
			return 65535, nil // HTTPRequestAgent
		case 805:
			return 4, nil // HTTPResponseCode
		case 806:
			return 65535, nil // HTTPResponseType
		case 807:
			return 8, nil // HTTPResponseTime
		case 808:
			return 65535, nil // HTTPSHost
		case 809:
			return 8, nil // HTTPSResponseTime
		case 810:
			return 4, nil // SMTPCommands
		case 811:
			return 4, nil // SMTPMailCount
		case 812:
			return 4, nil // SMTPRcptCount
		case 813:
			return 65535, nil // SMTPSender
		case 814:
			return 65535, nil // SMTPRecipient
		case 815:
			return 4, nil // SMTPStatusCodes
		case 816:
			return 4, nil // SMTPCode2XXCount
		case 817:
			return 4, nil // SMTPCode3XXCount
		case 818:
			return 4, nil // SMTPCode4XXCount
		case 819:
			return 4, nil // SMTPCode5XXCount
		case 820:
			return 65535, nil // SMTPDomain
		case 821:
			return 65535, nil // HTTPRequestRange
		case 830:
			return 4, nil // SIPMethod
		case 831:
			return 4, nil // SIPStatusCode
		case 832:
			return 65535, nil // SIPRequestURI
		case 833:
			return 65535, nil // SIPFrom
		case 834:
			return 65535, nil // SIPTo
		case 835:
			return 65535, nil // SIPContact
		case 836:
			return 65535, nil // SIPVia
		case 837:
			return 65535, nil // SIPRoute
		case 838:
			return 65535, nil // SIPRecordRoute
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 16982: // IPFIXColStyle - https://raw.githubusercontent.com/CESNET/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return 1, nil // destinationGeo
		case 100:
			return 1, nil // HTTPUserAgent
		case 101:
			return 1, nil // HTTPMethod
		case 102:
			return 65535, nil // HTTPDomain
		case 103:
			return 65535, nil // HTTPReferer
		case 104:
			return 1, nil // HTTPContentType
		case 105:
			return 65535, nil // HTTPUrl
		case 106:
			return 2, nil // HTTPStatus
		case 107:
			return 2, nil // HTTPHeaderCount
		case 200:
			return 4, nil // appPID
		case 201:
			return 65535, nil // appName
		case 202:
			return 4, nil // appUID
		case 203:
			return 1, nil // appMatchLevel
		case 400:
			return 16, nil // tunnelSrcIPv6
		case 401:
			return 16, nil // tunnelDstIPv6
		case 402:
			return 2, nil // tunnelSrcPort
		case 403:
			return 2, nil // tunnelDstPort
		case 404:
			return 1, nil // tunnelProtocol
		case 405:
			return 1, nil // tunnelType
		case 406:
			return 1, nil // tunnelTCPFlags
		case 407:
			return 2, nil // tunnelICMPcode
		case 408:
			return 1, nil // tunnelTeredoHeaders
		case 409:
			return 1, nil // tunnelTeredoTrailers
		case 410:
			return 65535, nil // tunnelSourceGeo
		case 411:
			return 65535, nil // tunnelDestinationGeo
		case 412:
			return 65535, nil // tunnelOuterSourceGeo
		case 413:
			return 65535, nil // tunnelOuterDestinationGeo
		case 414:
			return 1, nil // tunnelHOPLimit
		case 500:
			return 4, nil // HTTPRequestType
		case 501:
			return 65535, nil // HTTPRequestHost
		case 502:
			return 65535, nil // HTTPRequestURL
		case 503:
			return 4, nil // HTTPRequestAgentID
		case 504:
			return 65535, nil // HTTPRequestAgent
		case 505:
			return 65535, nil // HTTPRequestReferer
		case 506:
			return 4, nil // HTTPResponseCode
		case 507:
			return 65535, nil // HTTPResponseType
		case 508:
			return 4, nil // HTTPResponseTime
		case 800:
			return 2, nil // tlsCliVer
		case 801:
			return 2, nil // tlsSerVer
		case 802:
			return 2, nil // tlsSerCips
		case 803:
			return 65535, nil // tlsCliCips
		case 804:
			return 65535, nil // tlsCertSubject
		case 805:
			return 65535, nil // tlsCertIssuer
		case 806:
			return 8, nil // tlsCertNotBefore
		case 807:
			return 8, nil // tlsCerNotAfter
		case 808:
			return 2, nil // tlsPkeyLength
		case 809:
			return 4, nil // tlsPkeyExponent
		case 810:
			return 65535, nil // tlsPkeyAlgorithm
		case 811:
			return 65535, nil // tlsServerName
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 35632: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 180:
			return 65535, nil // HTTPUrl
		case 187:
			return 65535, nil // HTTPHost
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 39499: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 1:
			return 65535, nil // HTTPRequestHost
		case 2:
			return 65535, nil // HTTPRequestURL
		case 3:
			return 65535, nil // HTTPRequestReferer
		case 4:
			return 4, nil // HTTPRequestType
		case 10:
			return 65535, nil // HTTPResponseType
		case 12:
			return 4, nil // HTTPResponseCode
		case 20:
			return 65535, nil // HTTPRequestAgent
		case 21:
			return 4, nil // HTTPRequestAgentID
		case 22:
			return 2, nil // HTTPRequestAgentOS
		case 23:
			return 2, nil // HTTPRequestAgentOSMajor
		case 24:
			return 2, nil // HTTPRequestAgentOSMinor
		case 25:
			return 2, nil // HTTPRequestAgentOSBuild
		case 26:
			return 2, nil // HTTPRequestAgentApp
		case 27:
			return 2, nil // HTTPRequestAgentAppMajor
		case 28:
			return 2, nil // HTTPRequestAgentAppMinor
		case 29:
			return 2, nil // HTTPRequestAgentAppBuild
		case 32:
			return 1, nil // voipPacketType
		case 33:
			return 65535, nil // sipCallId
		case 34:
			return 65535, nil // sipCallingParty
		case 35:
			return 65535, nil // sipCalledParty
		case 36:
			return 65535, nil // sipVia
		case 37:
			return 4, nil // sipInviteRingingTime
		case 38:
			return 4, nil // sipOkTime
		case 39:
			return 4, nil // sipByeTime
		case 40:
			return 4, nil // sipRtpIp4
		case 41:
			return 16, nil // sipRtpIp6
		case 42:
			return 2, nil // sipRtpAudio
		case 43:
			return 2, nil // sipRtpVideo
		case 44:
			return 8, nil // sipStats
		case 45:
			return 1, nil // rtpCodec
		case 46:
			return 4, nil // rtpJitter
		case 47:
			return 4, nil // rtcpLost
		case 48:
			return 8, nil // rtcpPackets
		case 49:
			return 8, nil // rtcpOctets
		case 50:
			return 1, nil // rtcpSourceCount
		case 51:
			return 65535, nil // sipUserAgent
		case 52:
			return 65535, nil // sipRequestUri
		case 53:
			return 65535, nil // sipCSeq
		case 61:
			return 4, nil // NPMJitterDev
		case 62:
			return 4, nil // NPMJitterAvg
		case 63:
			return 4, nil // NPMJitterMin
		case 64:
			return 4, nil // NPMJitterMax
		case 65:
			return 4, nil // NPMDelayDev
		case 66:
			return 4, nil // NPMDelayAvg
		case 67:
			return 4, nil // NPMDelayMin
		case 68:
			return 4, nil // NPMDelayMax
		case 69:
			return 4, nil // NPMRoundTripTime
		case 70:
			return 4, nil // NPMServerResponseTime
		case 71:
			return 4, nil // NPMTCPRetransmission
		case 72:
			return 4, nil // NPMTCPOutOfOrder
		case 110:
			return 2, nil // DNSID
		case 111:
			return 2, nil // DNSFlagsCodes
		case 112:
			return 2, nil // DNSQuestionCount
		case 113:
			return 2, nil // DNSAnswRecCount
		case 114:
			return 2, nil // DNSAuthRecCount
		case 115:
			return 2, nil // DNSAddtRecCount
		case 116:
			return 65535, nil // DNSCRRName
		case 117:
			return 2, nil // DNSCRRType
		case 118:
			return 2, nil // DNSCRRClass
		case 119:
			return 4, nil // DNSCRRTTL
		case 120:
			return 65535, nil // DNSCRRRDATA
		case 121:
			return 65535, nil // DNSQNAME
		case 122:
			return 2, nil // DNSQTYPE
		case 123:
			return 2, nil // DNSQCLASS
		case 124:
			return 2, nil // DNSCRRRDATALen
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	case 44913: // IPFIXColStyle - https://raw.githubusercontent.com/SecDorks/ipfixcol/master/base/config/ipfix-elements.xml

		switch elementid {
		case 10:
			return 2, nil // origSourceTransportPort
		case 11:
			return 4, nil // origSourceIPv4Address
		case 12:
			return 2, nil // origDestinationTransportPort
		case 13:
			return 4, nil // origDestinationIPv4Address
		case 14:
			return 16, nil // origSourceIPv6Address
		case 15:
			return 16, nil // origDestinationIPv6Address
		case 20:
			return 65535, nil // HTTPRequestHost
		case 21:
			return 65535, nil // HTTPRequestURL
		case 22:
			return 65535, nil // HTTPRequestUserAgent
		case 12345:
			return 65535, nil // Unknown
		default:
			return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
		}

	default:
		return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
	}
	return 0, nil
}

