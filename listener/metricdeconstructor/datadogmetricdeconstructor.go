package metricdeconstructor

import (
	"strings"
)

type datadogMetricDeconstructor struct {
}

func (parser *datadogMetricDeconstructor) Parse(originalMetric string) (string, map[string]string, error) {
	dimensions := map[string]string{}
	parts := strings.SplitN(originalMetric, "[", 2)
	if len(parts) != 2 {
		return originalMetric, dimensions, nil
	}
	if len(parts[1]) == 0 || parts[1][len(parts[1])-1] != ']' {
		return originalMetric, dimensions, nil
	}
	parts[1] = parts[1][:len(parts[1])-1]
	newMetricName := parts[0]
	tagParts := strings.Split(parts[1], ",")
	for _, tagPart := range tagParts {
		tagSectionsParts := strings.SplitN(tagPart, ":", 2)
		if len(tagSectionsParts) != 2 {
			// Maybe a tag rather than a dimension.  Skip?
			continue
		}
		tagName := tagSectionsParts[0]
		tagValue := tagSectionsParts[1]
		dimensions[tagName] = tagValue
	}
	return newMetricName, dimensions, nil
}

func datadogLoader(options string) (MetricDeconstructor, error) {
	return &datadogMetricDeconstructor{}, nil
}