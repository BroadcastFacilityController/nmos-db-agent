package watchdog

import "github.com/google/uuid"

func strPtrOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func uuidPtrOrNil(id uuid.UUID) *uuid.UUID {
	if id == uuid.Nil {
		return nil
	}
	return &id
}
