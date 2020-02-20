# Practice Round - golang
To add a new strategy, implement the `Strategy` interface and select it in `func main()`.

## Strategies

### Random Strategy
Shuffle the input slices and sequentially sum up all the elements, skipping elements which would overflow the max
slices. Keep shuffling multiple times, and return the trial with the best result.
