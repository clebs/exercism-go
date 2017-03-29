package secret

const testVersion = 1

const reverse = 16

var masks = []uint{1, 2, 4, 8}
var handshakes = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake converts a numeric code into its secret handshake sequence of moves
func Handshake(code uint) []string {
	secret := make([]string, 0)

	for _, mask := range masks {
		if mask&code == mask {
			secret = append(secret, handshakes[mask])
		}
	}

	if code&reverse == reverse {
		for i, j := 0, len(secret)-1; i < j; i, j = i+1, j-1 {
			secret[i], secret[j] = secret[j], secret[i]
		}
	}
	return secret
}
