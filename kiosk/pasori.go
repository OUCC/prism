package main

// #cgo LDFLAGS: -lpafe
// #include <libpafe/libpafe.h>
// #include <libpafe/pasori_command.h>
// #include <libpafe/felica_command.h>
import "C"
import (
	. "github.com/OUCC/prism/logger"

	"fmt"
	"time"
)

var (
	pasori     *C.pasori
	pasoriCode = make(chan string, 0)
)

func setupPaSoRi() {
	if pasori = C.pasori_open(); pasori == nil {
		Log.Fatal("failed to open pasori")
	}

	switch C.pasori_type(pasori) {
	case C.PASORI_TYPE_S310:
		Log.Debug("PaSoRi S310")
	case C.PASORI_TYPE_S320:
		Log.Debug("PaSoRi S320")
	case C.PASORI_TYPE_S330:
		Log.Debug("PaSoRi S330")
	}

	C.pasori_init(pasori)
}

func pasoriLoop() {
	for {
		if f := C.felica_polling(pasori, C.FELICA_POLLING_ANY, 0, 1); f != nil {
			var idm [C.FELICA_IDM_LENGTH]uint8
			C.felica_get_idm(f, (*C.uint8)(&idm[0]))
			idmStr := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x:%02x:%02x",
				idm[0], idm[1], idm[2], idm[3],
				idm[4], idm[5], idm[6], idm[7])
			Log.Debug(idmStr)

			pasoriCode <- idmStr
		}
		time.Sleep(3 * time.Second)
	}
}
