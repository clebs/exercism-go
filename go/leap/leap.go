//Package leap provides logic for operating with leap years
package leap

const testVersion = 3

// IsLeapYear returns true if the given year is a leap year, false otherwise
func IsLeapYear(year int) bool {
	return year%4 == 0 && (!(year%100 == 0) || year%400 == 0)
}
