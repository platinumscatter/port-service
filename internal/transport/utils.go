package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/platinumscatter/port-service/internal/domain"
)

func portHttpToDomain(p *Port) (*domain.Port, error) {
	return domain.NewPort(
		p.ID,
		p.Name,
		p.Code,
		p.City,
		p.Country,
		append([]string(nil), p.Alias...),
		append([]string(nil), p.Regions...),
		append([]float64(nil), p.Coordinates...),
		p.Province,
		p.Timezone,
		append([]string(nil), p.Unlocs...),
	)
}

func readPorts(ctx context.Context, r io.Reader, portChan chan Port) error {
	decoder := json.NewDecoder(r)

	t, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimiter: %w", err)
	}
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}
	for decoder.More() {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		t, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read port ID: %w", err)
		}
		portID, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", t)
		}
		var port Port
		if err := decoder.Decode(&port); err != nil {
			return fmt.Errorf("failed to decode port: %w", err)
		}
		port.ID = portID
		portChan <- port
	}
	return nil
}
