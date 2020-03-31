package analyzers

import "net/url"

func l() (int, []bool, url.URL, error) {
	if 1 == 2 {
		return 0, nil
	}
	return
}
