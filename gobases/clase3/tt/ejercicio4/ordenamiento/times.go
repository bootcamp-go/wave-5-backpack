package ordenamiento

import (
	"log"
	"time"
)

// Track times
func comienzo(s string) (string, time.Time) {
	log.Printf("Comienza: %s\n", s)

	return s, time.Now()
}

func track(s string, comienzo time.Time) time.Duration {
	final := time.Now()

	log.Printf("Finaliza: %s\n", s)
	log.Printf("Duration: %s\n",final.Sub(final))

	return final.Sub(final)
}
