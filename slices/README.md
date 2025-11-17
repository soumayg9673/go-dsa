# Array/Slices

## Operations
1. (Append): Appending an element to the end of a list, is O(1) time complexity. We go directly to the end and add the element.
2. (Index): Accessing an element by index, is O(1) time complexity. We go directly to the index and return the element.
3. (Delete): Removing an element from the middle of the list, is O(n). We have to shift all the elements after the deleted element down one index.
4. (Search): Searching for an element in a list, is O(n). We have to iterate over the list until we find the element.

## Cons
1. When we need to frequently delete elements from the middle of the list.
2. When we need to frequently search for specific elements in the entire list.