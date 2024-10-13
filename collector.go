package main

import (
	"context"
	"log/slog"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/frebib/jool-exporter/jool"
)

// errorMap lists all jool statistics that should be exported under the
// `jool_translation_error_total` metric, along with the exported error= label
var errorMap = map[jool.Statistic]string{
	jool.StatisticPool6Unset:                      "pool6-unset",
	jool.StatisticSkbShared:                       "skb-shared",
	jool.StatisticL3HdrOffset:                     "l3hdr-offset",
	jool.StatisticSkbTruncated:                    "skb-truncated",
	jool.StatisticHdr6:                            "hdr6",
	jool.StatisticHdr4:                            "hdr4",
	jool.StatisticUnknownL4Proto:                  "unknown-l4_proto",
	jool.StatisticUnknownIcmp6Type:                "unknown-icmp6-type",
	jool.StatisticUnknownIcmp4Type:                "unknown-icmp4-type",
	jool.StatisticDoubleIcmp6Error:                "double-icmp6-error",
	jool.StatisticDoubleIcmp4Error:                "double-icmp4-error",
	jool.StatisticUnknownProtoInner:               "unknown-proto-inner",
	jool.StatisticHairpinLoop:                     "hairpin-loop",
	jool.StatisticPool6Mismatch:                   "pool6-mismatch",
	jool.StatisticPool4Mismatch:                   "pool4-mismatch",
	jool.StatisticIcmp6Filter:                     "icmp6-filter",
	jool.StatisticUntranslatableDst6:              "untranslatable-dst6",
	jool.StatisticUntranslatableDst4:              "untranslatable-dst4",
	jool.Statistic6056F:                           "6056-f",
	jool.StatisticMaskDomainNotFound:              "mask-domain-not-found",
	jool.StatisticBib6NotFound:                    "bib6-not-found",
	jool.StatisticBib4NotFound:                    "bib4-not-found",
	jool.StatisticSessionNotFound:                 "session-not-found",
	jool.StatisticAdf:                             "adf",
	jool.StatisticV4Syn:                           "v4-syn",
	jool.StatisticSyn6Expected:                    "syn6-expected",
	jool.StatisticSyn4Expected:                    "syn4-expected",
	jool.StatisticType1Pkt:                        "type1pkt",
	jool.StatisticType2Pkt:                        "type2pkt",
	jool.StatisticSoExists:                        "so_exists",
	jool.StatisticSoFull:                          "so_full",
	jool.Statistic64Src:                           "64-src",
	jool.Statistic64Dst:                           "64-dst",
	jool.Statistic64PskbCopy:                      "64-pskb-copy",
	jool.Statistic646791Enoent:                    "646791-enoent",
	jool.Statistic64IcmpCsum:                      "64-icmp-csum",
	jool.Statistic64UntranslatableParamProbPtr:    "64-untranslatable-param-prob-ptr",
	jool.Statistic64Ttl:                           "64-ttl",
	jool.Statistic64FragmentedIcmp:                "64-fragmented-icmp",
	jool.Statistic642xFrag:                        "64-2xfrag",
	jool.Statistic64FragThenExt:                   "64-frag-then-ext",
	jool.Statistic64SegmentsLeft:                  "64-segments-left",
	jool.Statistic46Src:                           "46-src",
	jool.Statistic46Dst:                           "46-dst",
	jool.Statistic46PskbCopy:                      "46-pskb-copy",
	jool.Statistic466791ENoEnt:                    "466791-enoent",
	jool.Statistic46IcmpCsum:                      "46-icmp-csum",
	jool.Statistic46UntranslatableParamProblemPtr: "46-untranslatable-param-problem-ptr",
	jool.Statistic46Ttl:                           "46-ttl",
	jool.Statistic46FragmentedIcmp:                "46-fragmented-icmp",
	jool.Statistic46SrcRoute:                      "46-src-route",
	jool.Statistic46FragmentedZeroCsum:            "46-fragmented-zero-csum",
	jool.Statistic46BadMtu:                        "46-bad-mtu",
	jool.StatisticFailedRoutes:                    "failed-routes",
	jool.StatisticPktTooBig:                       "pkt-too-big",
	jool.StatisticDstOutput:                       "dst-output",
	jool.StatisticIcmp6errSuccess:                 "icmp6err-success",
	jool.StatisticIcmp6errFailure:                 "icmp6err-failure",
	jool.StatisticIcmp4errSuccess:                 "icmp4err-success",
	jool.StatisticIcmp4errFailure:                 "icmp4err-failure",
	jool.StatisticIcmpExtBig:                      "icmpext-big",

	// It may seem odd to see this here, but for some reason jool can count
	// unknown errors and tell us about them 🤷
	jool.StatisticUnknown: "unknown",
}

type joolCollector struct {
	ctx       context.Context
	errLogMap sync.Map

	descEnabled    *prometheus.Desc
	descSuccess    *prometheus.Desc
	descXlatErr    *prometheus.Desc
	descRecvd      *prometheus.Desc
	descBIBEntries *prometheus.Desc
	descSessions   *prometheus.Desc
	descENoMem     *prometheus.Desc
}

func NewJoolCollector(ctx context.Context, namespace string) prometheus.Collector {
	return &joolCollector{
		ctx: ctx,

		descEnabled: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "instance", "enabled"),
			"1 indicates that the instance is enabled, otherwise 0",
			[]string{"instance"}, nil,
		),
		descSuccess: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "translation", "success_total"),
			"Total number of successful translated packets (does not imply successful delivery)",
			[]string{"instance"}, nil,
		),
		descRecvd: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "received_packets_total"),
			"Total number of packets received by the translator",
			[]string{"instance", "ipversion"}, nil,
		),
		descBIBEntries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "bib", "entries_total"),
			"Total number of entries in the Binding Information Base: https://datatracker.ietf.org/doc/html/rfc6146#section-3.1",
			[]string{"instance"}, nil,
		),
		descSessions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "bib", "sessions_total"),
			"Total number of session entries in the Binding Information Base",
			[]string{"instance"}, nil,
		),
		descENoMem: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "memory_allocation", "failures_total"),
			"Memory allocation failures",
			[]string{"instance"}, nil,
		),

		descXlatErr: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "translation", "error_total"),
			"Total number of errors that accounted for a failed/cancelled packet translation",
			[]string{"instance", "error"}, nil,
		),
	}
}

func (j *joolCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- j.descEnabled
	descs <- j.descSuccess
	descs <- j.descRecvd
	descs <- j.descBIBEntries
	descs <- j.descSessions
	descs <- j.descENoMem

	descs <- j.descXlatErr
}

func (j *joolCollector) Collect(metrics chan<- prometheus.Metric) {
	instances, err := jool.Instances(j.ctx)
	if err != nil {
		slog.Error("Error querying jool instances", "error", err)
		desc := prometheus.NewDesc("jool_collect_error", "Error querying jool instances", nil, nil)
		metrics <- prometheus.NewInvalidMetric(desc, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(instances))
	for _, instance := range instances {
		go func() {
			defer wg.Done()
			err := j.collectInstance(instance.Name, metrics)
			if err != nil {
				slog.Error("Error querying jool instance", "error", err, "instance", instance.Name)
				desc := prometheus.NewDesc(
					"jool_collect_error", "Error querying jool instance", nil,
					prometheus.Labels{"instance": instance.Name},
				)
				metrics <- prometheus.NewInvalidMetric(desc, err)
			}
		}()
	}

	wg.Wait()
}

func (j *joolCollector) collectInstance(inst string, metrics chan<- prometheus.Metric) error {
	stats, err := jool.Stats(j.ctx, inst)
	if err != nil {
		return err
	}

	// Unconditionally emit these counters even if the stats didn't return them
	// to us (as zero) so they're always initialised
	metrics <- prometheus.MustNewConstMetric(j.descEnabled, prometheus.GaugeValue, float64(1-stats[jool.StatisticXlatorDisabled]), inst)
	metrics <- prometheus.MustNewConstMetric(j.descSuccess, prometheus.CounterValue, float64(stats[jool.StatisticSuccess]), inst)
	metrics <- prometheus.MustNewConstMetric(j.descRecvd, prometheus.CounterValue, float64(stats[jool.StatisticReceived4]), inst, "4")
	metrics <- prometheus.MustNewConstMetric(j.descRecvd, prometheus.CounterValue, float64(stats[jool.StatisticReceived6]), inst, "6")
	metrics <- prometheus.MustNewConstMetric(j.descBIBEntries, prometheus.CounterValue, float64(stats[jool.StatisticBibEntries]), inst)
	metrics <- prometheus.MustNewConstMetric(j.descSessions, prometheus.CounterValue, float64(stats[jool.StatisticSessions]), inst)
	metrics <- prometheus.MustNewConstMetric(j.descENoMem, prometheus.CounterValue, float64(stats[jool.StatisticENoMem]), inst)

	for stat, value := range stats {
		switch stat {
		// We already handled these stats above
		case jool.StatisticXlatorDisabled,
			jool.StatisticSuccess,
			jool.StatisticReceived4,
			jool.StatisticReceived6,
			jool.StatisticBibEntries,
			jool.StatisticSessions,
			jool.StatisticENoMem:
			continue
		default:
			// Assume that everything else is an error
			errcode, exists := errorMap[stat]
			if !exists {
				// This convoluted mess atomically ensures we only log errors
				// for each (instance, error) pair exactly once
				var emptyMap sync.Map
				atomicMap, _ := j.errLogMap.LoadOrStore(inst, &emptyMap)
				_, isLogged := atomicMap.(*sync.Map).LoadOrStore(stat, struct{}{})
				if !isLogged {
					slog.Warn("Unknown error statistic found", "instance", inst, "stat", stat)
				}

				// It's not an error so we should skip it
				continue
			}
			metrics <- prometheus.MustNewConstMetric(j.descXlatErr, prometheus.CounterValue, float64(value), inst, errcode)
		}
	}
	return nil
}

var _ prometheus.Collector = &joolCollector{}
