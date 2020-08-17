package waktu

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--

		buf[w] = '0'
	} else {
		for v > 0 {
			w--

			const ten = 10
			buf[w] = byte(v%ten) + '0'
			v /= 10
		}
	}

	return w
}
