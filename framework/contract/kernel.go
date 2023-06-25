package contract

import "net/http"

const KernelKey = "zima:kernel"

type Kernel interface {
	WebEngine() http.Handler
}
