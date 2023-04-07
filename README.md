# go-local-cache
Implement a local cache in Go that meets the following requirements:

1. Can `Get` a value from the local cache using a key
2. Can `Set` a key/value pair to the local cache
3. Keys are string type and value can be any types
4. Item in the local cache expires in 30 seconds
5. `Set` will overwrite the existed item and reset the timer (when the key is the same)

Also meets other requirements:
* Separate interface & implementation in different files
* Write unit test for each methods in interface
