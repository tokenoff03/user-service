package utils

import "google.golang.org/grpc/metadata"

type MetadataReaderWriter struct {
	metadata.MD
}

func (w MetadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, v := range w.MD {
		if len(v) > 0 {
			if err := handler(k, v[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w MetadataReaderWriter) Set(key, val string) {
	w.MD[key] = append(w.MD[key], val)
}
