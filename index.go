package piscine

func Index2(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}

	return -1
}

/* The Index2 function takes two input strings, s and toFind, and
returns the starting index of the first occurrence of the toFind
string within the s string. If toFind is an empty string, the function
returns 0.

Here's how the function works:

The function first checks if the length of the toFind string is 0.
If it is, it means that there's nothing to find, so it returns 0
as specified in the requirements.
If the length of toFind is not 0, the function enters a for loop
that iterates through the s string using the index i. The loop
runs from 0 to len(s) - len(toFind). This ensures that there are
enough characters remaining in the s string to potentially match
the toFind string.
Inside the loop, the function extracts a substring from the s
string starting at index i and with a length equal to the length
of the toFind string. This is done using the slice notation s[i:i+len(toFind)].
It then compares the extracted substring with the toFind string
using the == operator. If they are equal, it means that a match
has been found, and the function returns the current value of i,
which is the starting index of the match.
If no match is found within the loop, the function returns -1
after completing the loop.
In summary, the Index2 function searches for the first occurrence
of the toFind string within the s string and returns its starting
index. It takes care to handle cases where toFind is an empty string
or when there are not enough characters left in s for a match. */
