## Simple Example

In AWS SDK, you can see that most of the `struct` and `functions` receives and `pointer`(e.g `*string` and `*int`).
```golang
func main() {
  // some code here

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("my-bucket"),
	})

  // the rest of it below
}
```

You can see in the _code snippet_ above that there's a function called `aws.String()`, it basically receives a `string` and returns a `*string`, and with that you have `aws.Int()`, `aws.Uint()`, ... how could we improve that using *Generics*?

Simple! Just writing a generic solution.
```golang
func Ptr[T any](value T) *T {
  return &value
}
```

Let's breakdown the function above:
- `func`: The keyword to define a function
- `Ptr`: The function name
- `[T any]`: Says that `T` is any type.
- `(value T) *T`: The function will receive 1 param of type `T` named `value`, and return a `*T` value.

