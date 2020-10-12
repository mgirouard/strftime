// Package strftime is a "good enough" implementation of the venerable strftime.
//
// Good enough means:
//
// 1) fast enough for most use cases;
// 2) most format chars are supported;
// 3) locale handling is missing;
// 4) %G is the same as %Y;
// 4) %U, %V, and %W are all the same;
//
//
// Supported Format Chars
// 
// %a The abbreviated name of the day of the week. Eg: Mon.
//
// %A The full name of the day of the week. Eg: Monday.
//
// %b The abbreviated month name. Eg: Feb.
//
// %B The full month name. Eg: February.
//
// %c Equivalent to %F %T.
//
// %C The century number (year/100) as a 2-digit integer. Eg 20.
//
// %d The day of the month as a zero-padded decimal number (01 to 31). Eg: 09.
//
// %D Equivalent to %m/%d/%y. Rarely used outside the US due to ambiguous
// ordering.
//
// %e Like %d, the day of the month as a decimal number, but a leading zero
// is replaced by a space.
// 
// %F Equivalent to %Y-%m-%d (the ISO 8601 date format).
// 
// %g Equivalent to %y.
// 
// %G Equivalent to %Y.
// 
// %h Equivalent to %b.
// 
// %% A literal '%' character.
// 
// %H The hour as a decimal number using a 24-hour clock (range 00 to 23).
// 
// %I The hour as a decimal number using a 12-hour clock (range 01 to 12).
// 
// %l The hour (12-hour  clock) as a decimal number (range 1 to 12); single
// digits are preceded by a blank.
// 
// %m The month as a decimal number (range 01 to 12).
// 
// %M The minute as a decimal number (range 00 to 59).
// 
// %n A newline character.
//
// %R The time in 24-hour notation (%H:%M).
// 
// %s The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC)
// 
// %S The second as a zero-padded decimal number (range 00 to 60).
// 
// %t A tab character.
// 
// %T The time in 24-hour notation (%H:%M:%S).
// 
// %u The day of the week as a decimal, range 1 to 7, Monday being 1.
// 
// %w The day of the week as a decimal, range 0 to 6, Sunday being 0.
// 
// %x Equivalent to %F.
// 
// %X Equivalent to %T.
// 
// %y The year as a decimal number without a century (range 00 to 99)
//
// %Y The year as a decimal number including the century.
//
//
// Future Support is Planned For
// 
// %j The day of the year as a decimal number (range 001 to 366).
// 
// %k The hour (24-hour clock) as a decimal number (range 0 to 23); single
// digits are preceded by a blank.
// 
// %p Either "AM" or "PM" according to the given time value.
// 
// %P Like %p but in lowercase: "am" or "pm".
// 
// %r The time in a.m. or p.m. notation.
// 
// %U The week number of the current year as a decimal number, range 00 to 53.
//
// %V Equivalent to %U.
// 
// %W Equivalent to %U.
// 
// %z The +hhmm or -hhmm numeric timezone (that is, the hour and minute offset from UTC).
// 
// %Z The timezone name or abbreviation.
// 
// %+ The date and time in date(1) format. (TZ) (Not supported in glibc2.)
package strftime
