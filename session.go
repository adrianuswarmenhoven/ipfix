package ipfix

type Session struct {

	//AssociatedTemplates Templates points to the list of active templates. Without a template record a data record can not be encoded or decoded
	AssociatedTemplates *ActiveTemplates
}

/*

 Achieving this goal is complicated somewhat by two factors: 1) the
   need to support the reuse of Template IDs within a Transport Session
   and 2) the need to support unreliable transmission for Templates when
   UDP is used as the transport protocol for IPFIX Messages.

   Sequencing Template Management Actions

   Since there is no guarantee of the ordering of exported IPFIX
   Messages across SCTP Streams or over UDP, an Exporting Process MUST
   sequence all Template management actions (i.e., Template Records
   defining new Templates and Template Withdrawals withdrawing them)
   using the Export Time field in the IPFIX Message Header.

   An Exporting Process MUST NOT export a Data Set described by a new
   Template in an IPFIX Message with an Export Time before the Export
   Time of the IPFIX Message containing that Template.  If a new
   Template and a Data Set described by it appear in the same IPFIX
   Message, the Template Set containing the Template MUST appear before
   the Data Set in the Message.

   An Exporting Process MUST NOT export any Data Sets described by a
   withdrawn Template in IPFIX Messages with an Export Time after the
   Export Time of the IPFIX Message containing the Template Withdrawal
   withdrawing that Template.




Claise, et al.               Standards Track                   [Page 42]

RFC 7011              IPFIX Protocol Specification        September 2013


   Put another way, a Template describes Data Records contained in IPFIX
   Messages when the Export Time of such messages is between a specific
   start and end time, inclusive.  The start time is the Export Time of
   the IPFIX Message containing the Template Record.  The end time is
   one of two times: if the template is withdrawn during the session,
   then it is the Export Time of the IPFIX Message containing the
   Template Withdrawal for the template; otherwise, it is the end of the
   Transport Session.

   Even if sent in order, IPFIX Messages containing Template management
   actions could arrive at the Collecting Process out of order, i.e., if
   sent via UDP or via different SCTP Streams.  Given this, Template
   Withdrawals and subsequent reuse of Template IDs can significantly
   complicate the problem of determining Template lifetimes at the
   Collecting Process.  A Collecting Process MAY implement a buffer and
   use Export Time information to disambiguate the order of Template
   management actions.  This buffer, if implemented, SHOULD be
   configurable to impart a delay on the order of the maximum reordering
   delay experienced at the Collecting Process.  Note, in this case,
   that the Collecting Process's clock is irrelevant: it is only
   comparing the Export Times of Messages to each other.

8.3.  Additional Considerations for Template Management over SCTP

   The specifications in this section apply only to SCTP; in cases of
   contradiction with specifications in Section 8 or Section 8.1, this
   section takes precedence.

   Template Sets and Options Template Sets MAY be sent on any SCTP
   Stream.  Data Sets sent on a given SCTP Stream MAY be represented by
   Template Records exported on any SCTP Stream.

   Template Sets and Options Template Sets MUST be sent reliably, using
   SCTP ordered delivery.

   Template Withdrawals MAY be sent on any SCTP Stream.  Template
   Withdrawals MUST be sent reliably, using SCTP ordered delivery.
   Template IDs MAY be reused by sending a Template Withdrawal and/or a
   new Template Record on a different SCTP Stream than the stream on
   which the original Template was sent.

   Additional Template Management considerations are provided in
   [RFC6526], which specifies an extension to explicitly link Templates
   with SCTP Streams.  In exchange for more restrictive rules on the
   assignment of Template Records to SCTP Streams, this extension allows
   fast, reliable reuse of Template IDs and estimation of Data Record
   loss per Template.

Additional Considerations for Template Management over UDP

   The specifications in this section apply only to UDP; in cases of
   contradiction with specifications in Section 8 or Section 8.1, this
   section takes precedence.

   Since UDP provides no method for reliable transmission of Templates,
   Exporting Processes using UDP as the transport protocol MUST
   periodically retransmit each active Template at regular intervals.
   The Template retransmission interval MUST be configurable via, for
   example, the templateRefreshTimeout and optionsTemplateRefreshTimeout
   parameters as defined in [RFC6728].  Default settings for these
   values are deployment- and application-specific.

   Before exporting any Data Records described by a given Template
   Record or Options Template Record, especially in the case of Template
   ID reuse as described in Section 8.1, the Exporting Process SHOULD
   send multiple copies of the Template Record in a separate IPFIX
   Message, in order to help ensure that the Collecting Process has
   received it.

   In order to minimize resource requirements for Templates that are no
   longer being used by the Exporting Process, the Collecting Process
   MAY associate a lifetime with each Template received in a Transport
   Session.  Templates not refreshed by the Exporting Process within the
   lifetime can then be discarded by the Collecting Process.  The
   Template lifetime at the Collecting Process MAY be exposed by a
   configuration parameter or MAY be derived from observation of the
   interval of periodic Template retransmissions from the Exporting
   Process.  In this latter case, the Template lifetime SHOULD default
   to at least 3 times the observed retransmission rate.

   Template Withdrawals (Section 8.1) MUST NOT be sent by Exporting
   Processes exporting via UDP and MUST be ignored by Collecting
   Processes collecting via UDP.  Template IDs MAY be reused by
   Exporting Processes by exporting a new Template for the Template ID
   after waiting at least 3 times the retransmission delay.  Note that
   Template ID reuse may lead to incorrect interpretation of Data
   Records if the retransmission and lifetime are not properly
   configured.

   When a Collecting Process receives a new Template Record or Options
   Template Record via UDP for an already-allocated Template ID, and
   that Template or Options Template is identical to the already-
   received Template or Options Template, it SHOULD NOT log the
   retransmission, as this is the normal operation of Template refresh
   over UDP.




Claise, et al.               Standards Track                   [Page 44]

RFC 7011              IPFIX Protocol Specification        September 2013


   When a Collecting Process receives a new Template Record or Options
   Template Record for an already-allocated Template ID, and that
   Template or Options Template is different from the already-received
   Template or Options Template, the Collecting Process MUST replace the
   Template or Options Template for that Template ID with the newly
   received Template or Options Template.  This is the normal operation
   of Template ID reuse over UDP.

   As Template IDs are unique per UDP session and per Observation
   Domain, at any given time, the Collecting Process SHOULD maintain the
   following for all the current Template Records and Options Template
   Records: <IPFIX Device, Exporter source UDP port, Collector IP
   address, Collector destination UDP port, Observation Domain ID,
   Template ID, Template Definition, Last Received>.

9.  The Collecting Process's Side

   This section describes the handling of the IPFIX protocol at the
   Collecting Process common to all transport protocols.  Additional
   considerations for SCTP and UDP are provided in Sections 9.2 and 9.3,
   respectively.  Template management at Collecting Processes is covered
   in Section 8.

   The Collecting Process MUST listen for association requests /
   connections to start new Transport Sessions from the Exporting
   Process.

   The Collecting Process MUST note the Information Element identifier
   of any Information Element that it does not understand and MAY
   discard that Information Element from received Data Records.

   The Collecting Process MUST accept padding in Data Records and
   Template Records.  The padding size is the Set Length minus the size
   of the Set Header (4 octets for the Set ID and the Set Length),
   modulo the minimum Record size deduced from the Template Record.

   The IPFIX protocol has a Sequence Number field in the Export header
   that increases with the number of IPFIX Data Records in the IPFIX
   Message.  A Collector can detect out-of-sequence, dropped, or
   duplicate IPFIX Messages by tracking the Sequence Number.  A
   Collector SHOULD provide a logging mechanism for tracking out-of-
   sequence IPFIX Messages.  Such out-of-sequence IPFIX Messages may be
   due to Exporter resource exhaustion where it cannot transmit messages
   at their creation rate, an Exporting Process reset, congestion on the
   network link between the Exporter and Collector, Collector resource
   exhaustion where it cannot process the IPFIX Messages at their
   arrival rate, out-of-order packet reception, duplicate packet
   reception, or an attacker injecting false messages.



Claise, et al.               Standards Track                   [Page 45]

RFC 7011              IPFIX Protocol Specification        September 2013


9.1.  Collecting Process Handling of Malformed IPFIX Messages

   If the Collecting Process receives a malformed IPFIX Message, it MUST
   discard the IPFIX Message and SHOULD log the error.  A malformed
   IPFIX Message is one that cannot be interpreted due to nonsensical
   length values (e.g., a variable-length Information Element longer
   than its enclosing Set, a Set longer than its enclosing IPFIX
   Message, or an IPFIX Message shorter than an IPFIX Message Header) or
   a reserved Version value (which may indicate that a future version of
   IPFIX is being used for export but in practice occurs most often when
   non-IPFIX data is sent to an IPFIX Collecting Process).  Note that
   non-zero Set padding does not constitute a malformed IPFIX Message.

   As the most likely cause of malformed IPFIX Messages is a poorly
   implemented Exporting Process, or the sending of non-IPFIX data to an
   IPFIX Collecting Process, human intervention is likely necessary to
   correct the issue.  In the meantime, the Collecting Process MAY
   attempt to rectify the situation any way it sees fit, including:

   - terminating the TCP connection or SCTP connection

   - using the receiver window to reduce network load from the
     malfunctioning Exporting Process

   - buffering and saving malformed IPFIX Message(s) to assist in
     diagnosis

   - attempting to resynchronize the stream, e.g., as described in
     Section 10.3 of [RFC5655]

   Resynchronization should only be attempted if the Collecting Process
   has reason to believe that the error is transient.  On the other
   hand, the Collecting Process SHOULD stop processing IPFIX Messages
   from clearly malfunctioning Exporting Processes (e.g., those from
   which the last few IPFIX Messages have been malformed).

9.2.  Additional Considerations for SCTP Collecting Processes

   As an Exporting Process may request and support more than one stream
   per SCTP association, the Collecting Process MUST support the opening
   of multiple SCTP Streams.

9.3.  Additional Considerations for UDP Collecting Processes

   A Transport Session for IPFIX Messages transported over UDP is
   defined from the point of view of the Exporting Process and roughly
   corresponds to the time during which a given Exporting Process sends
   IPFIX Messages over UDP to a given Collecting Process.  Since this is



Claise, et al.               Standards Track                   [Page 46]

RFC 7011              IPFIX Protocol Specification        September 2013


   difficult to detect at the Collecting Process, the Collecting Process
   MAY discard all Transport Session state after no IPFIX Messages are
   received from a given Exporting Process within a given Transport
   Session during a configurable idle timeout.

   The Collecting Process SHOULD accept Data Records without the
   associated Template Record (or other definitions such as Common
   Properties) required to decode the Data Record.  If the Template
   Records or other definitions have not been received at the time Data
   Records are received, the Collecting Process MAY store the Data
   Records for a short period of time and decode them after the Template
   Records or other definitions are received, comparing Export Times of
   IPFIX Messages containing the Template Records with those containing
   the Data Records as discussed in Section 8.2.  Note that this
   mechanism may lead to incorrectly interpreted records in the presence
   of Template ID reuse or other identifiers with limited lifetimes.

   10.  Transport Protocol

   The IPFIX Protocol Specification has been designed to be transport
   protocol independent.  Note that the Exporter can export to multiple
   Collecting Processes using independent transport protocols.

   The IPFIX Message Header 16-bit Length field limits the length of an
   IPFIX Message to 65535 octets, including the header.  A Collecting
   Process MUST be able to handle IPFIX Message lengths of up to
   65535 octets.

   While an Exporting Process or Collecting Process may support multiple
   transport protocols, Transport Sessions are bound to a transport
   protocol.  Transport Session state MUST NOT be migrated by an
   Exporting Process or Collecting Process among Transport Sessions
   using different transport protocols between the same Exporting
   Process and Collecting Process pair.  In other words, an Exporting
   Process supporting multiple transport protocols is conceptually
   multiple Exporting Processes, one per supported transport protocol.
   Likewise, a Collecting Process supporting multiple transport
   protocols is conceptually multiple Collecting Processes, one per
   supported transport protocol.

10.1.  Transport Compliance and Transport Usage

   SCTP [RFC4960] using the Partially Reliable SCTP (PR-SCTP) extension
   as specified in [RFC3758] MUST be implemented by all compliant
   implementations.  UDP [UDP] MAY also be implemented by compliant
   implementations.  TCP [TCP] MAY also be implemented by compliant
   implementations.




Claise, et al.               Standards Track                   [Page 47]

RFC 7011              IPFIX Protocol Specification        September 2013


   SCTP should be used in deployments where Exporters and Collectors are
   communicating over links that are susceptible to congestion.  SCTP is
   capable of providing any required degree of reliability when used
   with the PR-SCTP extension.

   TCP may be used in deployments where Exporters and Collectors
   communicate over links that are susceptible to congestion, but SCTP
   is preferred, due to its ability to limit back pressure on Exporters
   and its message-versus-stream orientation.

   UDP may be used, although it is not a congestion-aware protocol.
   However, in this case the IPFIX traffic between the Exporter and
   Collector must be separately contained or provisioned to minimize the
   risk of congestion-related loss.

   By default, the Collecting Process listens for connections on SCTP,
   TCP, and/or UDP port 4739.  By default, the Collecting Process
   listens for secure connections on SCTP, TCP, and/or UDP port 4740
   (refer to the Security Considerations section).  By default, the
   Exporting Process attempts to connect to one of these ports.  It MUST
   be possible to configure both the Exporting and Collecting Processes
   to use different ports than the default.

10.2.  SCTP

   This section describes how IPFIX is transported over SCTP [RFC4960]
   using the PR-SCTP [RFC3758] extension.

10.2.1.  Congestion Avoidance

   SCTP provides the required level of congestion avoidance by design.

   SCTP detects congestion in the end-to-end path between the IPFIX
   Exporting Process and the IPFIX Collecting Process, and limits the
   transfer rate accordingly.  When an IPFIX Exporting Process has
   records to export but detects that transmission by SCTP is
   temporarily impossible, it can either wait until sending is possible
   again or decide to drop the record.  In the latter case, the dropped
   export data SHOULD be accounted for, so that the amount of dropped
   export data can be reported using the mechanism described in
   Section 4.3.










Claise, et al.               Standards Track                   [Page 48]

RFC 7011              IPFIX Protocol Specification        September 2013


10.2.2.  Reliability

   The SCTP transport protocol is by default reliable but has the
   capability to deliver messages with partial reliability [RFC3758].

   Using reliable SCTP messages for IPFIX export is not in itself a
   guarantee that all Data Records will be delivered.  If there is
   congestion on the link from the Exporting Process to the Collecting
   Process, or if a significant number of retransmissions are required,
   the send queues on the Exporting Process may fill up; the Exporting
   Process MAY either suspend, export, or discard the IPFIX Messages.
   If Data Records are discarded, the IPFIX Sequence Numbers used for
   export MUST reflect the loss of data.

10.2.3.  MTU

   SCTP provides the required IPFIX Message fragmentation service based
   on Path MTU (PMTU) discovery.

10.2.4.  Association Establishment and Shutdown

   The IPFIX Exporting Process initiates an SCTP association with the
   IPFIX Collecting Process.  The Exporting Process MAY establish more
   than one association (connection "bundle" in SCTP terminology) to the
   Collecting Process.

   An Exporting Process MAY support more than one active association to
   different Collecting Processes (including the case of different
   Collecting Processes on the same host).

   When an Exporting Process is shut down, it SHOULD shut down the SCTP
   association.

   When a Collecting Process no longer wants to receive IPFIX Messages,
   it SHOULD shut down its end of the association.  The Collecting
   Process SHOULD continue to receive and process IPFIX Messages until
   the Exporting Process has closed its end of the association.

   When a Collecting Process detects that the SCTP association has been
   abnormally terminated, it MUST continue to listen for a new
   association establishment.

   When an Exporting Process detects that the SCTP association to the
   Collecting Process is abnormally terminated, it SHOULD try to
   re-establish the association.

   Association timeouts SHOULD be configurable.




Claise, et al.               Standards Track                   [Page 49]

RFC 7011              IPFIX Protocol Specification        September 2013


10.2.5.  Failover

   If the Collecting Process does not acknowledge an attempt by the
   Exporting Process to establish an association, SCTP will
   automatically retry association establishment using exponential
   backoff.  The Exporter MAY log an alarm if the underlying SCTP
   association establishment times out; this timeout should be
   configurable on the Exporter.

   The Exporting Process MAY open a backup SCTP association to a
   Collecting Process in advance, if it supports Collecting Process
   failover.

10.2.6.  Streams

   An Exporting Process MAY request more than one SCTP Stream per
   association.  Each of these streams may be used for the transmission
   of IPFIX Messages containing Data Sets, Template Sets, and/or Options
   Template Sets.

   Depending on the requirements of the application, the Exporting
   Process may send Data Sets with full or partial reliability, using
   ordered or out-of-order delivery, over any SCTP Stream established
   during SCTP association setup.

   An IPFIX Exporting Process MAY use any PR-SCTP service definition as
   per Section 4 of the PR-SCTP specification [RFC3758] when using
   partial reliability to transmit IPFIX Messages containing only
   Data Sets.

   However, Exporting Processes SHOULD mark such IPFIX Messages for
   retransmission for as long as resource or other constraints allow.

10.3.  UDP

   This section describes how IPFIX is transported over UDP [UDP].

10.3.1.  Congestion Avoidance

   UDP has no integral congestion-avoidance mechanism.  Its use over
   congestion-sensitive network paths is therefore not recommended.  UDP
   MAY be used in deployments where Exporters and Collectors always
   communicate over dedicated links that are not susceptible to
   congestion, i.e., links that are over-provisioned compared to the
   maximum export rate from the Exporters.






Claise, et al.               Standards Track                   [Page 50]

RFC 7011              IPFIX Protocol Specification        September 2013


10.3.2.  Reliability

   UDP is not a reliable transport protocol and cannot guarantee
   delivery of messages.  IPFIX Messages sent from the Exporting Process
   to the Collecting Process using UDP may therefore be lost.  UDP MUST
   NOT be used unless the application can tolerate some loss of IPFIX
   Messages.

   The Collecting Process SHOULD deduce the loss and reordering of IPFIX
   Data Records by looking at the discontinuities in the IPFIX Sequence
   Number.  In the case of UDP, the IPFIX Sequence Number contains the
   total number of IPFIX Data Records sent for the Transport Session
   prior to the receipt of this IPFIX Message, modulo 2^32.  A Collector
   SHOULD detect out-of-sequence, dropped, or duplicate IPFIX Messages
   by tracking the Sequence Number.

   Exporting Processes exporting IPFIX Messages via UDP MUST include a
   valid UDP checksum [UDP] in UDP datagrams including IPFIX Messages.

10.3.3.  MTU

   The maximum size of exported messages MUST be configured such that
   the total packet size does not exceed the PMTU.  If the PMTU is
   unknown, a maximum packet size of 512 octets SHOULD be used.

10.3.4.  Session Establishment and Shutdown

   As UDP is a connectionless protocol, there is no real session
   establishment or shutdown for IPFIX over UDP.  An Exporting Process
   starts sending IPFIX Messages to a Collecting Process at one point in
   time and stops sending them at another point in time.  This can lead
   to some complications in Template management, as outlined in
   Section 8.4 above.

10.3.5.  Failover and Session Duplication

   Because UDP is not a connection-oriented protocol, the Exporting
   Process is unable to determine from the transport protocol that the
   Collecting Process is no longer able to receive the IPFIX Messages.
   Therefore, it cannot invoke a failover mechanism.  However, the
   Exporting Process MAY duplicate the IPFIX Message to several
   Collecting Processes.









Claise, et al.               Standards Track                   [Page 51]

RFC 7011              IPFIX Protocol Specification        September 2013


10.4.  TCP

   This section describes how IPFIX is transported over TCP [TCP].

10.4.1.  Congestion Avoidance

   TCP controls the rate at which data can be sent from the Exporting
   Process to the Collecting Process, using a mechanism that takes into
   account both congestion in the network and the capabilities of the
   receiver.

   Therefore, an IPFIX Exporting Process may not be able to send IPFIX
   Messages at the rate that the Metering Process generates them, either
   because of congestion in the network or because the Collecting
   Process cannot handle IPFIX Messages fast enough.  As long as
   congestion is transient, the Exporting Process can buffer IPFIX
   Messages for transmission.  But such buffering is necessarily
   limited, both because of resource limitations and because of
   timeliness requirements, so ongoing and/or severe congestion may lead
   to a situation where the Exporting Process is blocked.

   When an Exporting Process has Data Records to export but the
   transmission buffer is full, and it wants to avoid blocking, it can
   decide to drop some Data Records.  The dropped Data Records MUST be
   accounted for, so that the number of lost records can later be
   reported as described in Section 4.3.

10.4.2.  Reliability

   TCP ensures reliable delivery of data from the Exporting Process to
   the Collecting Process.

10.4.3.  MTU

   As TCP offers a stream service instead of a datagram or sequential
   packet service, IPFIX Messages transported over TCP are instead
   separated using the Length field in the IPFIX Message Header.  The
   Exporting Process can choose any valid length for exported IPFIX
   Messages, as TCP handles segmentation.

   Exporting Processes may choose IPFIX Message lengths lower than the
   maximum in order to ensure timely export of Data Records.









Claise, et al.               Standards Track                   [Page 52]

RFC 7011              IPFIX Protocol Specification        September 2013


10.4.4.  Connection Establishment and Shutdown

   The IPFIX Exporting Process initiates a TCP connection to the
   Collecting Process.  An Exporting Process MAY support more than one
   active connection to different Collecting Processes (including the
   case of different Collecting Processes on the same host).  An
   Exporting Process MAY support more than one active connection to the
   same Collecting Process to avoid head-of-line blocking across
   Observation Domains.

   The Exporter MAY log an alarm if the underlying TCP connection
   establishment times out; this timeout should be configurable on the
   Exporter.

   When an Exporting Process is shut down, it SHOULD shut down the TCP
   connection.

   When a Collecting Process no longer wants to receive IPFIX Messages,
   it SHOULD close its end of the connection.  The Collecting Process
   SHOULD continue to read IPFIX Messages until the Exporting Process
   has closed its end.

   When a Collecting Process detects that the TCP connection to the
   Exporting Process has terminated abnormally, it MUST continue to
   listen for a new connection.

   When an Exporting Process detects that the TCP connection to the
   Collecting Process has terminated abnormally, it SHOULD try to
   re-establish the connection.  Connection timeouts and retry schedules
   SHOULD be configurable.  In the default configuration, an Exporting
   Process MUST NOT attempt to establish a connection more frequently
   than once per minute.

10.4.5.  Failover

   If the Collecting Process does not acknowledge an attempt by the
   Exporting Process to establish a connection, TCP will automatically
   retry connection establishment using exponential backoff.  The
   Exporter MAY log an alarm if the underlying TCP connection
   establishment times out; this timeout should be configurable on the
   Exporter.

   The Exporting Process MAY open a backup TCP connection to a
   Collecting Process in advance, if it supports Collecting Process
   failover.






Claise, et al.               Standards Track                   [Page 53]

RFC 7011              IPFIX Protocol Specification        September 2013


11.  Security Considerations

   The security considerations for the IPFIX protocol have been derived
   from an analysis of potential security threats, as discussed in the
   Security Considerations section of the IPFIX requirements document
   [RFC3917].  The requirements for IPFIX security are as follows:

   1. IPFIX must provide a mechanism to ensure the confidentiality of
      IPFIX data transferred from an Exporting Process to a Collecting
      Process, in order to prevent disclosure of Flow Records
      transported via IPFIX.

   2. IPFIX must provide a mechanism to ensure the integrity of IPFIX
      data transferred from an Exporting Process to a Collecting
      Process, in order to prevent the injection of incorrect data or
      control information (e.g., Templates), or the duplication of
      Messages, in an IPFIX Message stream.

   3. IPFIX must provide a mechanism to authenticate IPFIX Collecting
      and Exporting Processes, to prevent the collection of data from an
      unauthorized Exporting Process or the export of data to an
      unauthorized Collecting Process.

   Because IPFIX can be used to collect information for network
   forensics and billing purposes, attacks designed to confuse, disable,
   or take information from an IPFIX collection system may be seen as a
   prime objective during a sophisticated network attack.

   An attacker in a position to inject false messages into an IPFIX
   Message stream can affect either the application using IPFIX (by
   falsifying data) or the IPFIX Collecting Process itself (by modifying
   or revoking Templates, or changing options); for this reason, IPFIX
   Message integrity is important.

   The IPFIX Messages themselves may also contain information of value
   to an attacker, including information about the configuration of the
   network as well as end-user traffic and payload data, so care must be
   taken to confine their visibility to authorized users.  When an
   Information Element containing end-user payload information is
   exported, it SHOULD be transmitted to the Collecting Process using a
   means that secures its contents against eavesdropping.  Suitable
   mechanisms include the use of either a direct point-to-point
   connection assumed to be unavailable to attackers, or the use of an
   encryption mechanism.  It is the responsibility of the Collecting
   Process to provide a satisfactory degree of security for this
   collected data, including, if necessary, encryption and/or
   anonymization of any reported data; see Section 11.8.

*/
