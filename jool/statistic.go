package jool

import (
	"strconv"
)

type Statistic uint

// https://github.com/NICMx/Jool/blob/518790de38d8b043326c93b76d770d6f84f26c7c/src/common/stats.h#L14
const (
	StatisticReceived6 Statistic = iota + 1 // JSTAT_RECEIVED6
	StatisticReceived4                      // JSTAT_RECEIVED4
	StatisticSuccess                        // JSTAT_SUCCESS

	StatisticBibEntries // JSTAT_BIB_ENTRIES
	StatisticSessions   // JSTAT_SESSIONS

	StatisticENoMem // JSTAT_ENOMEM

	StatisticXlatorDisabled // JSTAT_XLATOR_DISABLED
	StatisticPool6Unset     // JSTAT_POOL6_UNSET

	StatisticSkbShared    // JSTAT_SKB_SHARED
	StatisticL3HdrOffset  // JSTAT_L3HDR_OFFSET
	StatisticSkbTruncated // JSTAT_SKB_TRUNCATED
	StatisticHdr6         // JSTAT_HDR6
	StatisticHdr4         // JSTAT_HDR4

	StatisticUnknownL4Proto    // JSTAT_UNKNOWN_L4_PROTO
	StatisticUnknownIcmp6Type  // JSTAT_UNKNOWN_ICMP6_TYPE
	StatisticUnknownIcmp4Type  // JSTAT_UNKNOWN_ICMP4_TYPE
	StatisticDoubleIcmp6Error  // JSTAT_DOUBLE_ICMP6_ERROR
	StatisticDoubleIcmp4Error  // JSTAT_DOUBLE_ICMP4_ERROR
	StatisticUnknownProtoInner // JSTAT_UNKNOWN_PROTO_INNER

	StatisticHairpinLoop        // JSTAT_HAIRPIN_LOOP
	StatisticPool6Mismatch      // JSTAT_POOL6_MISMATCH
	StatisticPool4Mismatch      // JSTAT_POOL4_MISMATCH
	StatisticIcmp6Filter        // JSTAT_ICMP6_FILTER
	StatisticUntranslatableDst6 // JSTAT_UNTRANSLATABLE_DST6
	StatisticUntranslatableDst4 // JSTAT_UNTRANSLATABLE_DST4
	Statistic6056F              // JSTAT6056_F
	StatisticMaskDomainNotFound // JSTAT_MASK_DOMAIN_NOT_FOUND
	StatisticBib6NotFound       // JSTAT_BIB6_NOT_FOUND
	StatisticBib4NotFound       // JSTAT_BIB4_NOT_FOUND
	StatisticSessionNotFound    // JSTAT_SESSION_NOT_FOUND
	StatisticAdf                // JSTAT_ADF
	StatisticV4Syn              // JSTAT_V4_SYN
	StatisticSyn6Expected       // JSTAT_SYN6_EXPECTED
	StatisticSyn4Expected       // JSTAT_SYN4_EXPECTED

	StatisticType1Pkt // JSTAT_TYPE1PKT
	StatisticType2Pkt // JSTAT_TYPE2PKT
	StatisticSoExists // JSTAT_SO_EXISTS
	StatisticSoFull   // JSTAT_SO_FULL

	Statistic64Src                        // JSTAT64_SRC
	Statistic64Dst                        // JSTAT64_DST
	Statistic64PskbCopy                   // JSTAT64_PSKB_COPY
	Statistic646791Enoent                 // JSTAT646791_ENOENT
	Statistic64IcmpCsum                   // JSTAT64_ICMP_CSUM
	Statistic64UntranslatableParamProbPtr // JSTAT64_UNTRANSLATABLE_PARAM_PROB_PTR
	Statistic64Ttl                        // JSTAT64_TTL
	Statistic64FragmentedIcmp             // JSTAT64_FRAGMENTED_ICMP
	Statistic642xFrag                     // JSTAT64_2XFRAG
	Statistic64FragThenExt                // JSTAT64_FRAG_THEN_EXT
	Statistic64SegmentsLeft               // JSTAT64_SEGMENTS_LEFT

	Statistic46Src                           // JSTAT46_SRC
	Statistic46Dst                           // JSTAT46_DST
	Statistic46PskbCopy                      // JSTAT46_PSKB_COPY
	Statistic466791ENoEnt                    // JSTAT466791_ENOENT
	Statistic46IcmpCsum                      // JSTAT46_ICMP_CSUM
	Statistic46UntranslatableParamProblemPtr // JSTAT46_UNTRANSLATABLE_PARAM_PROBLEM_PTR
	Statistic46Ttl                           // JSTAT46_TTL
	Statistic46FragmentedIcmp                // JSTAT46_FRAGMENTED_ICMP
	Statistic46SrcRoute                      // JSTAT46_SRC_ROUTE
	Statistic46FragmentedZeroCsum            // JSTAT46_FRAGMENTED_ZERO_CSUM
	Statistic46BadMtu                        // JSTAT46_BAD_MTU

	StatisticFailedRoutes // JSTAT_FAILED_ROUTES
	StatisticPktTooBig    // JSTAT_PKT_TOO_BIG
	StatisticDstOutput    // JSTAT_DST_OUTPUT

	StatisticIcmp6errSuccess // JSTAT_ICMP6ERR_SUCCESS
	StatisticIcmp6errFailure // JSTAT_ICMP6ERR_FAILURE
	StatisticIcmp4errSuccess // JSTAT_ICMP4ERR_SUCCESS
	StatisticIcmp4errFailure // JSTAT_ICMP4ERR_FAILURE

	StatisticIcmpExtBig // JSTAT_ICMPEXT_BIG

	StatisticJooldEmpty      // JSTAT_JOOLD_EMPTY
	StatisticJooldTimeout    // JSTAT_JOOLD_TIMEOUT
	StatisticJooldMissingAck // JSTAT_JOOLD_MISSING_ACK
	StatisticJooldAdOngoing  // JSTAT_JOOLD_AD_ONGOING
	StatisticJooldPktFull    // JSTAT_JOOLD_PKT_FULL
	StatisticJooldQueuing    // JSTAT_JOOLD_QUEUING

	StatisticJooldSssQueued // JSTAT_JOOLD_SSS_QUEUED
	StatisticJooldSssSent   // JSTAT_JOOLD_SSS_SENT
	StatisticJooldSssRcvd   // JSTAT_JOOLD_SSS_RCVD
	StatisticJooldSssEnospc // JSTAT_JOOLD_SSS_ENOSPC
	StatisticJooldPktSent   // JSTAT_JOOLD_PKT_SENT
	StatisticJooldPktRcvd   // JSTAT_JOOLD_PKT_RCVD
	StatisticJooldAds       // JSTAT_JOOLD_ADS
	StatisticJooldAcks      // JSTAT_JOOLD_ACKS

	StatisticUnknown // JSTAT_UNKNOWN
	statisticPadding // JSTAT_PADDING

	StatisticCount
)

func (s Statistic) String() string {
	switch s {
	case StatisticReceived6:
		return "Received6"
	case StatisticReceived4:
		return "Received4"
	case StatisticSuccess:
		return "Success"
	case StatisticBibEntries:
		return "BibEntries"
	case StatisticSessions:
		return "Sessions"
	case StatisticENoMem:
		return "ENoMem"
	case StatisticXlatorDisabled:
		return "XlatorDisabled"

	case StatisticPool6Unset:
		return "Pool6Unset"
	case StatisticSkbShared:
		return "SkbShared"
	case StatisticL3HdrOffset:
		return "L3HdrOffset"
	case StatisticSkbTruncated:
		return "SkbTruncated"
	case StatisticHdr6:
		return "Hdr6"
	case StatisticHdr4:
		return "Hdr4"
	case StatisticUnknownL4Proto:
		return "UnknownL4Proto"
	case StatisticUnknownIcmp6Type:
		return "UnknownIcmp6Type"
	case StatisticUnknownIcmp4Type:
		return "UnknownIcmp4Type"
	case StatisticDoubleIcmp6Error:
		return "DoubleIcmp6Error"
	case StatisticDoubleIcmp4Error:
		return "DoubleIcmp4Error"
	case StatisticUnknownProtoInner:
		return "UnknownProtoInner"
	case StatisticHairpinLoop:
		return "HairpinLoop"
	case StatisticPool6Mismatch:
		return "Pool6Mismatch"
	case StatisticPool4Mismatch:
		return "Pool4Mismatch"
	case StatisticIcmp6Filter:
		return "Icmp6Filter"
	case StatisticUntranslatableDst6:
		return "UntranslatableDst6"
	case StatisticUntranslatableDst4:
		return "UntranslatableDst4"
	case Statistic6056F:
		return "6056F"
	case StatisticMaskDomainNotFound:
		return "MaskDomainNotFound"
	case StatisticBib6NotFound:
		return "Bib6NotFound"
	case StatisticBib4NotFound:
		return "Bib4NotFound"
	case StatisticSessionNotFound:
		return "SessionNotFound"
	case StatisticAdf:
		return "Adf"
	case StatisticV4Syn:
		return "V4Syn"
	case StatisticSyn6Expected:
		return "Syn6Expected"
	case StatisticSyn4Expected:
		return "Syn4Expected"
	case StatisticType1Pkt:
		return "Type1Pkt"
	case StatisticType2Pkt:
		return "Type2Pkt"
	case StatisticSoExists:
		return "SoExists"
	case StatisticSoFull:
		return "SoFull"
	case Statistic64Src:
		return "64Src"
	case Statistic64Dst:
		return "64Dst"
	case Statistic64PskbCopy:
		return "64PskbCopy"
	case Statistic646791Enoent:
		return "646791Enoent"
	case Statistic64IcmpCsum:
		return "64IcmpCsum"
	case Statistic64UntranslatableParamProbPtr:
		return "64UntranslatableParamProbPtr"
	case Statistic64Ttl:
		return "64Ttl"
	case Statistic64FragmentedIcmp:
		return "64FragmentedIcmp"
	case Statistic642xFrag:
		return "642xFrag"
	case Statistic64FragThenExt:
		return "64FragThenExt"
	case Statistic64SegmentsLeft:
		return "64SegmentsLeft"
	case Statistic46Src:
		return "46Src"
	case Statistic46Dst:
		return "46Dst"
	case Statistic46PskbCopy:
		return "46PskbCopy"
	case Statistic466791ENoEnt:
		return "466791ENoEnt"
	case Statistic46IcmpCsum:
		return "46IcmpCsum"
	case Statistic46UntranslatableParamProblemPtr:
		return "46UntranslatableParamProblemPtr"
	case Statistic46Ttl:
		return "46Ttl"
	case Statistic46FragmentedIcmp:
		return "46FragmentedIcmp"
	case Statistic46SrcRoute:
		return "46SrcRoute"
	case Statistic46FragmentedZeroCsum:
		return "46FragmentedZeroCsum"
	case Statistic46BadMtu:
		return "46BadMtu"
	case StatisticFailedRoutes:
		return "FailedRoutes"
	case StatisticPktTooBig:
		return "PktTooBig"
	case StatisticDstOutput:
		return "DstOutput"
	case StatisticIcmp6errSuccess:
		return "Icmp6errSuccess"
	case StatisticIcmp6errFailure:
		return "Icmp6errFailure"
	case StatisticIcmp4errSuccess:
		return "Icmp4errSuccess"
	case StatisticIcmp4errFailure:
		return "Icmp4errFailure"
	case StatisticIcmpExtBig:
		return "IcmpExtBig"

	case StatisticJooldEmpty:
		return "JooldEmpty"
	case StatisticJooldTimeout:
		return "JooldTimeout"
	case StatisticJooldMissingAck:
		return "JooldMissingAck"
	case StatisticJooldAdOngoing:
		return "JooldAdOngoing"
	case StatisticJooldPktFull:
		return "JooldPktFull"
	case StatisticJooldQueuing:
		return "JooldQueuing"
	case StatisticJooldSssQueued:
		return "JooldSssQueued"
	case StatisticJooldSssSent:
		return "JooldSssSent"
	case StatisticJooldSssRcvd:
		return "JooldSssRcvd"
	case StatisticJooldSssEnospc:
		return "JooldSssEnospc"
	case StatisticJooldPktSent:
		return "JooldPktSent"
	case StatisticJooldPktRcvd:
		return "JooldPktRcvd"
	case StatisticJooldAds:
		return "JooldAds"
	case StatisticJooldAcks:
		return "JooldAcks"

	case StatisticUnknown:
		return "Unknown"
	default:
		panic("Invalid statistic: " + strconv.Itoa(int(s)))
	}
}

func ParseStatistic(s string) Statistic {
	switch s {
	case "JSTAT_RECEIVED6":
		return StatisticReceived6
	case "JSTAT_RECEIVED4":
		return StatisticReceived4
	case "JSTAT_SUCCESS":
		return StatisticSuccess
	case "JSTAT_BIB_ENTRIES":
		return StatisticBibEntries
	case "JSTAT_SESSIONS":
		return StatisticSessions
	case "JSTAT_ENOMEM":
		return StatisticENoMem
	case "JSTAT_XLATOR_DISABLED":
		return StatisticXlatorDisabled
	case "JSTAT_POOL6_UNSET":
		return StatisticPool6Unset
	case "JSTAT_SKB_SHARED":
		return StatisticSkbShared
	case "JSTAT_L3HDR_OFFSET":
		return StatisticL3HdrOffset
	case "JSTAT_SKB_TRUNCATED":
		return StatisticSkbTruncated
	case "JSTAT_HDR6":
		return StatisticHdr6
	case "JSTAT_HDR4":
		return StatisticHdr4
	case "JSTAT_UNKNOWN_L4_PROTO":
		return StatisticUnknownL4Proto
	case "JSTAT_UNKNOWN_ICMP6_TYPE":
		return StatisticUnknownIcmp6Type
	case "JSTAT_UNKNOWN_ICMP4_TYPE":
		return StatisticUnknownIcmp4Type
	case "JSTAT_DOUBLE_ICMP6_ERROR":
		return StatisticDoubleIcmp6Error
	case "JSTAT_DOUBLE_ICMP4_ERROR":
		return StatisticDoubleIcmp4Error
	case "JSTAT_UNKNOWN_PROTO_INNER":
		return StatisticUnknownProtoInner
	case "JSTAT_HAIRPIN_LOOP":
		return StatisticHairpinLoop
	case "JSTAT_POOL6_MISMATCH":
		return StatisticPool6Mismatch
	case "JSTAT_POOL4_MISMATCH":
		return StatisticPool4Mismatch
	case "JSTAT_ICMP6_FILTER":
		return StatisticIcmp6Filter
	case "JSTAT_UNTRANSLATABLE_DST6":
		return StatisticUntranslatableDst6
	case "JSTAT_UNTRANSLATABLE_DST4":
		return StatisticUntranslatableDst4
	case "JSTAT6056_F":
		return Statistic6056F
	case "JSTAT_MASK_DOMAIN_NOT_FOUND":
		return StatisticMaskDomainNotFound
	case "JSTAT_BIB6_NOT_FOUND":
		return StatisticBib6NotFound
	case "JSTAT_BIB4_NOT_FOUND":
		return StatisticBib4NotFound
	case "JSTAT_SESSION_NOT_FOUND":
		return StatisticSessionNotFound
	case "JSTAT_ADF":
		return StatisticAdf
	case "JSTAT_V4_SYN":
		return StatisticV4Syn
	case "JSTAT_SYN6_EXPECTED":
		return StatisticSyn6Expected
	case "JSTAT_SYN4_EXPECTED":
		return StatisticSyn4Expected
	case "JSTAT_TYPE1PKT":
		return StatisticType1Pkt
	case "JSTAT_TYPE2PKT":
		return StatisticType2Pkt
	case "JSTAT_SO_EXISTS":
		return StatisticSoExists
	case "JSTAT_SO_FULL":
		return StatisticSoFull
	case "JSTAT64_SRC":
		return Statistic64Src
	case "JSTAT64_DST":
		return Statistic64Dst
	case "JSTAT64_PSKB_COPY":
		return Statistic64PskbCopy
	case "JSTAT646791_ENOENT":
		return Statistic646791Enoent
	case "JSTAT64_ICMP_CSUM":
		return Statistic64IcmpCsum
	case "JSTAT64_UNTRANSLATABLE_PARAM_PROB_PTR":
		return Statistic64UntranslatableParamProbPtr
	case "JSTAT64_TTL":
		return Statistic64Ttl
	case "JSTAT64_FRAGMENTED_ICMP":
		return Statistic64FragmentedIcmp
	case "JSTAT64_2XFRAG":
		return Statistic642xFrag
	case "JSTAT64_FRAG_THEN_EXT":
		return Statistic64FragThenExt
	case "JSTAT64_SEGMENTS_LEFT":
		return Statistic64SegmentsLeft
	case "JSTAT46_SRC":
		return Statistic46Src
	case "JSTAT46_DST":
		return Statistic46Dst
	case "JSTAT46_PSKB_COPY":
		return Statistic46PskbCopy
	case "JSTAT466791_ENOENT":
		return Statistic466791ENoEnt
	case "JSTAT46_ICMP_CSUM":
		return Statistic46IcmpCsum
	case "JSTAT46_UNTRANSLATABLE_PARAM_PROBLEM_PTR":
		return Statistic46UntranslatableParamProblemPtr
	case "JSTAT46_TTL":
		return Statistic46Ttl
	case "JSTAT46_FRAGMENTED_ICMP":
		return Statistic46FragmentedIcmp
	case "JSTAT46_SRC_ROUTE":
		return Statistic46SrcRoute
	case "JSTAT46_FRAGMENTED_ZERO_CSUM":
		return Statistic46FragmentedZeroCsum
	case "JSTAT46_BAD_MTU":
		return Statistic46BadMtu
	case "JSTAT_FAILED_ROUTES":
		return StatisticFailedRoutes
	case "JSTAT_PKT_TOO_BIG":
		return StatisticPktTooBig
	case "JSTAT_DST_OUTPUT":
		return StatisticDstOutput
	case "JSTAT_ICMP6ERR_SUCCESS":
		return StatisticIcmp6errSuccess
	case "JSTAT_ICMP6ERR_FAILURE":
		return StatisticIcmp6errFailure
	case "JSTAT_ICMP4ERR_SUCCESS":
		return StatisticIcmp4errSuccess
	case "JSTAT_ICMP4ERR_FAILURE":
		return StatisticIcmp4errFailure
	case "JSTAT_ICMPEXT_BIG":
		return StatisticIcmpExtBig
	case "JSTAT_JOOLD_EMPTY":
		return StatisticJooldEmpty
	case "JSTAT_JOOLD_TIMEOUT":
		return StatisticJooldTimeout
	case "JSTAT_JOOLD_MISSING_ACK":
		return StatisticJooldMissingAck
	case "JSTAT_JOOLD_AD_ONGOING":
		return StatisticJooldAdOngoing
	case "JSTAT_JOOLD_PKT_FULL":
		return StatisticJooldPktFull
	case "JSTAT_JOOLD_QUEUING":
		return StatisticJooldQueuing
	case "JSTAT_JOOLD_SSS_QUEUED":
		return StatisticJooldSssQueued
	case "JSTAT_JOOLD_SSS_SENT":
		return StatisticJooldSssSent
	case "JSTAT_JOOLD_SSS_RCVD":
		return StatisticJooldSssRcvd
	case "JSTAT_JOOLD_SSS_ENOSPC":
		return StatisticJooldSssEnospc
	case "JSTAT_JOOLD_PKT_SENT":
		return StatisticJooldPktSent
	case "JSTAT_JOOLD_PKT_RCVD":
		return StatisticJooldPktRcvd
	case "JSTAT_JOOLD_ADS":
		return StatisticJooldAds
	case "JSTAT_JOOLD_ACKS":
		return StatisticJooldAcks
	case "JSTAT_UNKNOWN":
		return StatisticUnknown
	default:
		return 0
	}
}
