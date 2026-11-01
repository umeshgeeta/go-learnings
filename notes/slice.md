Slice 's' is nothing but what is exposed through 'aparture' on the underlying array. Natually, we need to know how big is the aparture - legth of the slice and where to start. For the first 3 slices, we still start from the 0th element while for the 4th slice we start from the 3rd cell.

Is that narration, or way of 'understanding Slice' is valid/correct? Anything missing or you would like to correct?

>>
The aperture picture is right. One part of it does not match the fourth print: the slice’s capacity does not stay 6.

`s` holds a slice header, three fields pointing at an array created by the literal on line 6:

- pointer: where the aperture starts
- length: how many elements are visible
- capacity: how far this header is allowed to extend, counting from that pointer to the end of the array

The array itself stays six cells: `[2, 3, 5, 7, 11, 13]`. Nothing in this program allocates a new one. The header’s capacity is a separate number, and it changes when the pointer moves.

| assignment | start | len | cap |
|---|---|---|---|
| line 6, the literal | cell 0 (`2`) | 6 | 6 |
| line 10, `s[:0]` | cell 0 | 0 | 6 |
| line 14, `s[:4]` | cell 0 | 4 | 6 |
| line 18, `s[2:]` | cell 2 (`5`), the 3rd cell | 2 | 4 |

The first three headers start at cell 0, so the capacity stays 6. Line 18 moves the start two cells forward. Capacity is measured only forward from the new start, so it becomes `6 - 2 = 4`. The cells holding `2` and `3` are still in the array, and this header can no longer see them. A slice cannot move its aperture backward.

Each of those assignments stores a new header in the variable `s`. The visible elements are `s[0]` through `s[len-1]`. The capacity is the extra room in front of the pointer: line 14 can reopen `[:4]` after a length of 0 because `4 <= cap`. Appending past that capacity is the point where Go allocates a new, larger array and the old six cells stop being the backing store.

"Capacity is measured only forward from the new start"

I guess that is the 'key' correction to my understanding!

>>

Yes. That is the key correction. The backing array still has six cells after `s = s[2:]`, but the new header starts at `5`, so its capacity is only the four cells from there to the end: `[5, 7, 11, 13]`.

