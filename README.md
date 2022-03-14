# TOP X numbers

This program reads `a file` which contains one number each line and prints out the `top X numbers` as you wish.

## How to build and run

I use Docker to build and run this little project, so please make sure you have Docker installed.

Then all you need to do are:

```
# Build
docker build -t topx .

# Then quick run with default parameters
docker run -it --rm topx

# Run tests
docker run -it --rm topx sh -c "TOPX=4 go test ./... -v --cover"
```

## Available options for running the program

We use environment variables to override program parameters just for simplicity. You can change the following two environment variables:

- `INPUT ` The relative file path which contains the numbers
- `TOPX` Indicates who many largest numbers you want to pick up.

There are some sample commands below:

```

docker run -it --rm topx sh -c "INPUT=./test1.txt go run main.go"

docker run -it --rm topx sh -c "INPUT=./test1.txt TOPX=1 go run main.go"

docker run -it --rm topx sh -c "TOPX=1 go run main.go"

```

## Complexities

- Space complexity: We only have a heap with size `x` to store the top numbers, and we don't read the entire file into memory , so the space complexity is `O(x)`.

- Time complexity: As we use Priority Heap to store the top `x` numbers, the time complexity for both `Pop` and `Push` are `O(logx)`, and we need to do this for every number, given we have `n` numbers in total. So the complexity of finding top x numbers is `O(nlogx)`, and lastly we do another sorting for the top `x` numbers to make sure numbers are following the descending order, which the time complexity is `O(logx)` as well. So the total time complexity should still be `O(nlogx)`

## Future improvements

- Refactor the `calculator.CalculateLargestNumbers` function to extract file reading, so that this function doesn't need to reply on a actual file in unit and can be tested better, and even to be mocked out when unit testing.
- Still in the `calculator.CalculateLargestNumbers` function, it would be better to just call `heap.Init()` once no matter how many top numbers we need to fetch.
- Allow program to parse options for parameters, such as `./app -input x.txt -topx 5`
- Add benchmark tests to reveal the actual performance of this program, and improve the program from there.
- Better error handling. Such as custom error types, more specific error messages and so on.
